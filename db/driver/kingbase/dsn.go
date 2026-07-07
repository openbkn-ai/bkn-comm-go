package kingbase

import (
	"fmt"
	"time"

	"github.com/openbkn-ai/bkn-comm-go/db/driver/common"
)

func FormatDSN(cfg common.DSNConfig) (string, error) {
	dsn := ""
	if cfg.Username != "" {
		dsn += fmt.Sprintf("user=%s ", cfg.Username)
	}
	if cfg.Password != "" {
		dsn += fmt.Sprintf("password=%s ", cfg.Password)
	}
	if cfg.Host != "" {
		dsn += fmt.Sprintf("host=%s ", cfg.Host)
	}
	if cfg.Port != "" {
		dsn += fmt.Sprintf("port=%s ", cfg.Port)
	}
	if cfg.DBName != "" {
		dsn += fmt.Sprintf("search_path=%s ", cfg.DBName)
	}

	if timeoutStr, exist := cfg.Props.Get("timeout"); exist {
		timeout, err := time.ParseDuration(timeoutStr.(string))
		if err != nil {
			return "", err
		}

		dsn += fmt.Sprintf("connect_timeout=%d ", timeout/(1000*1000*1000))
	}

	dsn += "sslmode=disable dbname=proton"
	return dsn, nil
}
