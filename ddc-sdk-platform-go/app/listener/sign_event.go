package listener

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

// SignEventListener 签名接口
type SignEventListener interface {
	SignEvent(addr common.Address, tx *types.Transaction) (*types.Transaction, error)
}
