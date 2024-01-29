package main

import (
	"context"
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"crypto/sha1"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"golang.org/x/crypto/pbkdf2"
	"google.golang.org/grpc"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
	"io"
	"log"
	"net/http"
	"sync"
	tms_proto_licensing "testtms/tms.proto.licensing"
	"time"
)

var count int = 0
var countLock sync.Mutex

func CallTMSAPIGetCustomer(url string, wg *sync.WaitGroup) {
	defer wg.Done()
	method := "GET"
	client := &http.Client{}
	client.Timeout = time.Minute
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	_ = body
	countLock.Lock()
	defer countLock.Unlock()
	count++
	fmt.Println(count)
	fmt.Println(string(body))
}

func CallCheckLicense(conn *grpc.ClientConn, customerId string, atmId string) error {
	// Call grpc CheckLicense
	req := tms_proto_licensing.CheckLicenseRequest{
		AtmId:      atmId,
		CustomerId: customerId,
	}
	client := tms_proto_licensing.NewLicenseServiceClient(conn)
	resp, err := client.CheckLicense(context.Background(), &req)
	if err != nil {
		log.Println(fmt.Sprintf("[%s][%s]Error: %v", customerId, atmId, err))
		return err
	}
	log.Println(fmt.Sprintf("[%s][%s]Status: %v", customerId, atmId, resp.Status))
	return nil
}

type dbConfig struct {
	host     string
	port     int
	user     string
	password string
	dbname   string
}

func CallCheckLicenseCustomer(url string, customerId string, dbConf dbConfig, wgParam *sync.WaitGroup) error {
	// Get all atmId from db
	wgParam.Add(1)
	defer wgParam.Done()
	conn, err := gorm.Open(sqlserver.Open(fmt.Sprintf("sqlserver://%s:%s@%s:%d?database=%s",
		dbConf.user, dbConf.password, dbConf.host, dbConf.port, dbConf.dbname)))
	if err != nil {
		log.Println(err)
		return err
	}
	var atmIds []string
	conn.Table("tblLicenseAtm").Select("AtmId").Find(&atmIds)
	log.Println(fmt.Sprintf("[%s]Total atmIds: %d", customerId, len(atmIds)))
	log.Println(fmt.Sprintf("[%s]List atmIds: %+v", customerId, atmIds))
	grpcConn, err := grpc.Dial(url, grpc.WithInsecure())
	if err != nil {
		log.Println(fmt.Sprintf("[%s]Error: %v", customerId, err))
		return err
	}
	// Call grpc
	var wg sync.WaitGroup
	for _, atmId := range atmIds {
		go func(customer string, atm string) {
			wg.Add(1)
			CallCheckLicense(grpcConn, customer, atm)
			wg.Done()
		}(customerId, atmId)
	}
	wg.Wait()
	return nil
}

type CustomerInfo struct {
	Id          string      `json:"id"`
	Code        string      `json:"code"`
	Name        string      `json:"name"`
	DbServer    string      `json:"dbServer"`
	DbName      string      `json:"dbName"`
	DbPort      int         `json:"dbPort"`
	DbUser      string      `json:"dbUser"`
	DbPassword  string      `json:"dbPassword"`
	Transaction bool        `json:"transaction"`
	AtmMgmt     bool        `json:"atmMgmt"`
	Cash        bool        `json:"cash"`
	Software    bool        `json:"software"`
	Security    bool        `json:"security"`
	CreatedDate time.Time   `json:"createdDate"`
	IsActive    bool        `json:"isActive"`
	Roles       interface{} `json:"roles"`
}

func GetCustomerList(url string) ([]CustomerInfo, error) {
	body, err := http.Get(url)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	defer body.Body.Close()
	var customerList []CustomerInfo
	err = json.NewDecoder(body.Body).Decode(&customerList)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return customerList, nil
}

func DecryptAes(encryptedData string) (string, error) {
	// Decrypt DB user and password
	key := string([]byte{'4', 'a', 's', 'm', '$', 'm', 'c', 'r', 't', '@', '1', '9', '8', 't', 'q', 'k', '$'})
	if encryptedData != "" {
		salt := md5.Sum([]byte(key))
		keyArray := pbkdf2.Key([]byte(key), salt[:], 1000, 48, sha1.New)
		iv := keyArray[32:]
		keyArray = keyArray[:32]
		rijndael, err := aes.NewCipher(keyArray)
		if err != nil {
			return "", err
		}
		// Decrypt
		mode := cipher.NewCBCDecrypter(rijndael, iv)
		byteData, err := hex.DecodeString(encryptedData)
		if err != nil {
			return "", err
		}
		plainByte := make([]byte, len(byteData))
		mode.CryptBlocks(plainByte, byteData)
		// Unpad plaintext
		plainByte = unpadPKCS7(plainByte)
		return string(plainByte), nil
	}
	return "", errors.New("Empty encrypted data")
}

// PKCS7 Unpadding
func unpadPKCS7(data []byte) []byte {
	padding := int(data[len(data)-1])
	return data[:len(data)-padding]
}

type aesKeys struct {
	keyBytes []byte
	ivBytes  []byte
}
