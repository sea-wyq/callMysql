package data

import (
	"callMysql/internal/conf"
	"database/sql"

	"github.com/go-kratos/kratos/v2/log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/google/wire"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewMysqlController)

// Data .
type Data struct {
	db *sql.DB
}

// NewData .
func NewData(c *conf.Data, logger log.Logger) (*Data, func(), error) {
	// db, err := sql.Open("mysql", "root:root@tcp(mysql-service:3306)/demo")
	db, err := sql.Open(c.Database.Driver, c.Database.Source)
	if err != nil {
		logger.Log(log.LevelError, "mysql connect fail。。。")
	}
	cleanup := func() {
		logger.Log(log.LevelInfo, "closing the data resources")
		db.Close()
	}
	return &Data{db}, cleanup, err
}
