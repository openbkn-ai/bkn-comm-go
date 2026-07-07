package kingbase

import "database/sql/driver"

type Driver struct{}

func (d Driver) Open(name string) (driver.Conn, error) {
	return Open(name)
}
