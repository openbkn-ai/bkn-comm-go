package tidb

import (
	"database/sql/driver"

	"github.com/go-sql-driver/mysql"
)

func Open(dsn string) (driver.Conn, error) {
	return mysql.MySQLDriver{}.Open(dsn)
}

func OpenConnector(dsn string) (driver.Connector, error) {
	return mysql.MySQLDriver{}.OpenConnector(dsn)
}
