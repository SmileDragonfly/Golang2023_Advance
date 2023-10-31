package main

import (
	"github.com/google/uuid"
	mssql "github.com/microsoft/go-mssqldb"
	"gorm.io/driver/sqlserver"
	"gorm.io/gen"
	"gorm.io/gorm"
	"testgorm/models"
	"testgorm/query"
	"time"
)

// Dynamic SQL
type Querier interface {
	// SELECT * FROM @@table WHERE name = @name{{if role !=""}} AND role = @role{{end}}
	FilterWithNameAndRole(name, role string) ([]gen.T, error)
}

func main() {
	gormdb, err := gorm.Open(sqlserver.Open("server=192.168.66.51;user id=paygreen_dev;password=123@123A;port=1433;database=PayGreen_Dev_Hub"), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	//g := gen.NewGenerator(gen.Config{
	//	OutPath: "../query",
	//	Mode:    gen.WithoutContext | gen.WithDefaultQuery | gen.WithQueryInterface, // generate mode
	//})
	//g.UseDB(gormdb)
	//g.ApplyBasic(models.HubTxnLock{})
	//g.ApplyInterface(func(Querier) {}, models.HubTxnLock{})
	//g.Execute()
	query.SetDefault(gormdb)
	var id mssql.UniqueIdentifier
	id.Scan(uuid.New().String())
	err = query.HubTxnLock.Create(&models.HubTxnLock{
		ID:            id,
		TerminalHubID: "GORM HUB ID",
		HubTID:        "GORM HUBTID",
		TerminalID:    "GORM Terminal ID",
		CreatedDate:   time.Now(),
		ExpiredTime:   time.Now(),
	})
	if err != nil {
		panic(err)
	}
}
