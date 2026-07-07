package dmdb

import (
	"database/sql/driver"
	"fmt"
	"os"
)

type DMStmt struct {
	//dm.DmStatement
	driver.Stmt
}

func (dmStmt DMStmt) Exec(args []driver.Value) (driver.Result, error) {
	for i, v := range args {
		if _, ok := v.([]byte); ok {
			args[i] = string(v.([]byte))
		}
	}
	if os.Getenv("RDS_SDK_DEBUG") == "true" {
		fmt.Println("stmt exec: ", args)
	}
	return dmStmt.Stmt.Exec(args)
}
