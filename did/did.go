// Copyright The kweaver.ai Authors.
//
// Licensed under the Apache License, Version 2.0.
// See the LICENSE file in the project root for details.

package did

import (
	"errors"
	"net"
	"os"

	"github.com/sony/sonyflake"

	"github.com/openbkn-ai/bkn-comm-go/logger"
)

var sf *sonyflake.Sonyflake

// 初始化 sonyflake 对象
func init() {
	var st sonyflake.Settings
	st.MachineID = getMachineID
	st.CheckMachineID = checkMachineID

	sf = sonyflake.NewSonyflake(st)

	if sf == nil {
		panic("sonyflake not created")
	}
}

// 将 POD_IP 作为 MachineID, k8s 保证 POD_IP 唯一, POD_IP作为环境变量传入
func getMachineID() (uint16, error) {
	ipString := os.Getenv("POD_IP")
	if ipString == "" {
		logger.Error("Failed to get pod ip from env")
		return 0, errors.New("failed to get pod ip from env")
	}

	ip := net.ParseIP(ipString)
	if ip.IsLoopback() {
		logger.Error("IP is a loopback address")
		return 0, errors.New("IP is a loopback address")
	}

	if ip.IsUnspecified() {
		logger.Error("IP is an unspecified address, either the IPv4 address '0.0.0.0' or the IPv6 address '::'")
		return 0, errors.New("IP is an unspecified address")
	}

	if len(ip) == net.IPv6len {
		ip = ip[12:16]
	} else {
		ip = ip.To4()
	}

	return uint16(ip[2])<<8 + uint16(ip[3]), nil
}

// 校验 machineID, 若ip地址的后两段都为 0 (x.y.0.0), 则返回 false
func checkMachineID(machineID uint16) bool {
	return machineID != 0
}

// 生成分布式 ID
func GenerateDistributedID() (id uint64, err error) {
	id, err = sf.NextID()
	if err != nil {
		logger.Error(err.Error())
		return 0, err
	}

	return id, nil
}
