package models

import (
	mssql "github.com/microsoft/go-mssqldb"
	"time"
)

// TerminalHub represents the structure of the "tblTerminalHub" table in SQL Server.
type TerminalHub struct {
	ID                      string    `json:"Id"`                      // Unique Identifier
	TerminalHubID           string    `json:"TerminalHubID"`           // Terminal Hub ID
	ProcessorId             int64     `json:"ProcessorId"`             // Processor ID
	Cassette1Denom          int64     `json:"Cassette1Denom"`          // Denomination for Cassette 1
	Cassette2Denom          int64     `json:"Cassette2Denom"`          // Denomination for Cassette 2
	Cassette1Remain         int64     `json:"Cassette1Remain"`         // Remaining Amount in Cassette 1
	Cassette2Remain         int64     `json:"Cassette2Remain"`         // Remaining Amount in Cassette 2
	TotalLowAmount          int64     `json:"TotalLowAmount"`          // Total Low Amount
	WorkingKey              string    `json:"WorkingKey"`              // Working Key
	WorkingKeyInterval      int64     `json:"WorkingKeyInterval"`      // Working Key Interval
	HealthCheckInterval     int64     `json:"HealthCheckInterval"`     // Health Check Interval
	TxnDelayMin             int64     `json:"TxnDelayMin"`             // Minimum Transaction Delay
	TxnDelayMax             int64     `json:"TxnDelayMax"`             // Maximum Transaction Delay
	Surcharge               float64   `json:"Surcharge"`               // Surcharge Amount
	Status                  bool      `json:"Status"`                  // Status (true - Active, false - Inactive)
	Deleted                 bool      `json:"Deleted"`                 // Deleted (true - Deleted, false - Not Deleted)
	CreatedDate             time.Time `json:"CreatedDate"`             // Creation Date and Time
	RequestTimeoutProcessor int64     `json:"RequestTimeoutProcessor"` // Request Timeout for Processor
	RequestTimeoutTerminal  int64     `json:"RequestTimeoutTerminal"`  // Request Timeout for Terminal
	CurrentTSQ              int64     `json:"CurrentTSQ"`
}

// Terminal represents the structure of the "tblTerminal" table in SQL Server.
type Terminal struct {
	ID               string    `json:"Id"`               // Unique Identifier
	ShopID           string    `json:"ShopId"`           // Shop ID (nullable)
	TerminalID       string    `json:"TerminalID"`       // Terminal ID (Not Null)
	TerminalName     string    `json:"TerminalName"`     // Terminal Name
	SerieNo          string    `json:"SerieNo"`          // Serial Number
	TerminalModelID  string    `json:"TerminalModelId"`  // Terminal Model ID (Not Null)
	OsVersionID      string    `json:"OsVersionId"`      // OS Version ID (nullable)
	Address1         string    `json:"Address1"`         // Address Line 1
	Address2         string    `json:"Address2"`         // Address Line 2
	City             string    `json:"City"`             // City
	State            string    `json:"State"`            // State
	Location         string    `json:"Location"`         // Location
	NetworkType      int64     `json:"NetworkType"`      // Network Type (Not Null)
	CreatedDate      time.Time `json:"CreatedDate"`      // Creation Date and Time (Not Null)
	ActivedDate      time.Time `json:"ActivedDate"`      // Activation Date and Time
	LastModifiedDate time.Time `json:"LastModifiedDate"` // Last Modified Date and Time
	Deleted          bool      `json:"Deleted"`          // Deleted (Not Null)
	Status           int64     `json:"Status"`           // Status (Not Null)
	IsActive         bool      `json:"IsActive"`         // Is Active (Not Null)
	Zipcode          string    `json:"Zipcode"`          // Zipcode
	ProcessorTID     string    `json:"ProcessorTID"`     // Processor Terminal ID
	ProcessorID      int64     `json:"ProcessorId"`      // Processor ID (nullable)
	DeviceID         string    `json:"DeviceID"`         // Device ID
	TerminalHubID    string    `json:"TerminalHubId"`    // Terminal Hub ID (nullable)
}

// HubTransaction represents the structure of the "tblHubTransaction" table in SQL Server.
type HubTransaction struct {
	ID                 string    `json:"Id"`                 // Unique Identifier
	TerminalHubID      string    `json:"TerminalHubID"`      // Terminal Hub ID
	TerminalID         string    `json:"TerminalID"`         // Terminal ID
	SerialNumber       string    `json:"SerialNumber"`       // Serial Number
	TranCode           string    `json:"TranCode"`           // Transaction Code: CWD - Payment | INQ - Balance Inquiry
	ServerDate         time.Time `json:"ServerDate"`         // Server Date and Time
	TsqHub             string    `json:"TsqHub"`             // TSQ (Transaction Sequence Number) for the Hub
	CardMasked         string    `json:"CardMasked"`         // Masked Card Number
	EncryptedPAN       string    `json:"EncryptedPAN"`       // Encrypted PAN (Primary Account Number)
	WithdrawalAmount   float64   `json:"WithdrawalAmount"`   // Withdrawal Amount: Total Amount in json request
	OriginalAmount     float64   `json:"OriginalAmount"`     // Original Transaction Amount
	CashbackAmount     float64   `json:"CashbackAmount"`     // Cashback Amount
	TipAmount          float64   `json:"TipAmount"`          // Tip Amount
	SurchargeAmount    float64   `json:"SurchargeAmount"`    // Surcharge Amount
	TranStatus         int64     `json:"TranStatus"`         // Transaction Status (1 - Success, 0 - Failed, 2 - Create)
	ResponseCode       string    `json:"ResponseCode"`       // Response Code
	Description        string    `json:"Description"`        // Description
	SettlementDate     time.Time `json:"SettlementDate"`     // Settlement Date and Time: From processor
	AccountBalance     float64   `json:"AccountBalance"`     // Account Balance
	AvailableBalance   float64   `json:"AvailableBalance"`   // Available Balance
	SeqNum             string    `json:"SeqNum"`             // Sequence Number
	TxnId              string    `json:"TxnId"`              // Transaction ID (Formatted as 'yyyyMMddHHmmss|TerminalHubID|TSQ')
	SystemTraceNumber  string    `json:"SystemTraceNumber"`  // System Trace Number
	TransactionCounter string    `json:"TransactionCounter"` // EMV: 9F36 Value
}

// HubTxnLock represents the structure of the "tblHubTxnLock" table in SQL Server.
type HubTxnLock struct {
	ID            mssql.UniqueIdentifier `json:"Id" gorm:"column:Id;primaryKey"`            // Unique Identifier
	TerminalHubID string                 `json:"TerminalHubId" gorm:"column:TerminalHubID"` // Terminal Hub ID (nullable)
	HubTID        string                 `json:"HubTID" gorm:"column:HubTID"`               // Hub Terminal ID
	TerminalID    string                 `json:"TerminalID" gorm:"column:TerminalID"`       // Terminal ID
	CreatedDate   time.Time              `json:"CreatedDate" gorm:"column:CreatedDate"`     // Creation Date and Time
	ExpiredTime   time.Time              `json:"ExpiredTime" gorm:"column:ExpiredTime"`     // Expiration Time
}

func (HubTxnLock) TableName() string {
	return "dbo.tblHubTxnLock"
}

// Processor represents the structure of the "tblProcessor" table in SQL Server.
type Processor struct {
	ID             int64  `json:"Id"`             // Unique Identifier
	ProcessorName  string `json:"ProcessorName"`  // Processor Name
	BIN            string `json:"BIN"`            // BIN (Bank Identification Number)
	IpAddress      string `json:"IpAddress"`      // IP Address
	Port           int64  `json:"Port"`           // Port
	ProtocolTypeID int64  `json:"ProtocolTypeId"` // Protocol Type ID
	EnableTLS      bool   `json:"EnableTLS"`      // Enable TLS (true - Enabled, false - Disabled)
	Description    string `json:"Description"`    // Description
}
