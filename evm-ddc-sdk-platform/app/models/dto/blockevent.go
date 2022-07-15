package dto

import (
	"github.com/ethereum/go-ethereum/common"
)

type BlockEventBean struct {
	Events    []interface{}
	Timestamp string
}

type BlockDdcInfoBean struct {
	DdcInfos  []DdcInfoBean
	Timestamp string
}

type DdcInfoBean struct {
	Hash   common.Hash
	DdcIds []uint64
}
