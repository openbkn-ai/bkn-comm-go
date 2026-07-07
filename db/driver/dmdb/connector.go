package dmdb

import (
	"context"
	"database/sql/driver"
)

type DMConnector struct {
	driver.Connector
}

func (dmConnector *DMConnector) Connect(ctx context.Context) (driver.Conn, error) {
	conn, err := dmConnector.Connector.Connect(ctx)
	if err != nil {
		return nil, err
	}
	return &DMConn{conn}, err
}
