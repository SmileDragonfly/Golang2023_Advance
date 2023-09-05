package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
)

func main() {
	log.Println("Start")
	db, err := sql.Open("postgres", "host=localhost port=5432 user=postgres password=123@123A dbname=ewallet_dev sslmode=disable")
	if err != nil {
		log.Panic(err)
	}
	defer db.Close()
	// Gọi procedure và lấy kết quả truy vấn
	userId := "f1caff32-3150-408a-983a-396e1dee94d9"
	query := "CALL ewallet.get_transactions_by_user($1, $2)"
	var rsName string
	tx, _ := db.Begin()
	defer tx.Commit()
	err = tx.QueryRow(query, userId, "my_refcursor").Scan(&rsName)
	if err != nil {
		panic(err)
	}
	var trxns []Transaction
	if rsName != "" {
		// Sử dụng kết quả từ refcursor
		query := fmt.Sprintf(`FETCH FORWARD 50 FROM "%s"`, rsName)
		rows, err := tx.Query(query)
		if err != nil {
			panic(err)
		}
		defer rows.Close()
		for rows.Next() {
			var trxn Transaction
			err := rows.Scan(&trxn.ID,
				&trxn.TraceNumber,
				&trxn.UserID,
				&trxn.MerchantID,
				&trxn.Amount,
				&trxn.Fee,
				&trxn.OpenBalance,
				&trxn.CloseBalance,
				&trxn.TransactionType,
				&trxn.SourceRef,
				&trxn.DestinationRef,
				&trxn.DestinationName,
				&trxn.Remark,
				&trxn.State,
				&trxn.ProviderCode,
				&trxn.ChannelCode,
				&trxn.ChannelRef,
				&trxn.ChannelResponseCode,
				&trxn.ChannelData,
				&trxn.CreateTime,
				&trxn.UpdateTime,
				&trxn.TotalAmount,
			)
			if err != nil {
				panic(err)
			}
			trxns = append(trxns, trxn)
		}
	} else {
		fmt.Printf("Result: %s", rsName)
	}
	fmt.Printf("TRXNS: %+v", len(trxns))
}
