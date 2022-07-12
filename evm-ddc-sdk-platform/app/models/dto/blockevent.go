package dto

import "github.com/ethereum/go-ethereum/core/types"

type BlockEventBean struct {
	Events    []interface{}
	Timestamp string
}

type BlockDdcInfoBean struct {
	DdcInfos  []DdcInfoBean
	Timestamp string
}

type DdcInfoBean struct {
	DdcId   uint64
	Receipt *types.Receipt
}
