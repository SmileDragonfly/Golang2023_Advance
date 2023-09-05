package main

import "time"

type Transaction struct {
	ID                  string
	TraceNumber         string
	UserID              string
	MerchantID          string
	Amount              int64
	Fee                 int64
	TotalAmount         int64
	OpenBalance         int64
	CloseBalance        int64
	TransactionType     string // TopupWaitAuthen, Topup, Withdraw, TransferP2P
	SourceRef           string // Topup: LinkId
	DestinationRef      string // Topup: AccountId
	DestinationName     string // Topup: Mobile
	Remark              string // Topup: BankSwiftCode
	State               string // Success, Fail
	ProviderCode        string // Topup: BankSwiftCode
	ChannelCode         string // Topup: BankSwiftCode
	ChannelRef          string // Topup: Url
	ChannelResponseCode string // Topup: response code
	ChannelData         string // Topup: response data
	CreateTime          time.Time
	UpdateTime          time.Time
}
