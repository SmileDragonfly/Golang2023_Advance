package main

import "time"

type APIToken struct {
	Id            int64     `gorm:"column:Id;primary_key"`
	TerminalID    string    `gorm:"column:TerminalID;type:uniqueidentifier,not null"`
	UserID        string    `gorm:"column:UserID;type:uniqueidentifier"`
	Token         string    `gorm:"column:Token;type:varchar(250)"`
	FirebaseToken string    `gorm:"column:FirebaseToken;type:varchar(250)"`
	IP            string    `gorm:"column:IP;type:varchar(50)"`
	CreateDate    time.Time `gorm:"column:CreateDate;type:datetime,not null"`
	ExpiryDate    time.Time `gorm:"column:ExpiryDate;type:datetime,not null"`
}

func (a APIToken) TableName() string {
	return "tblAPIToken"
}

type HubTxnLock struct {
	Id            string    `gorm:"column:Id;primary_key;type:uniqueidentifier"`
	TerminalHubId string    `gorm:"column:TerminalHubId;type:uniqueidentifier"`
	HubTID        string    `gorm:"column:HubTID;type:varchar(50)"`
	TerminalID    string    `gorm:"column:TerminalID;type:varchar(50)"`
	CreatedDate   time.Time `gorm:"column:CreatedDate;type:datetime"`
	ExpiredTime   time.Time `gorm:"column:ExpiredTime;type:datetime"`
	LockType      int64     `gorm:"column:LockType;type:int"`
}

func (h HubTxnLock) TableName() string {
	return "tblHubTxnLock"
}
