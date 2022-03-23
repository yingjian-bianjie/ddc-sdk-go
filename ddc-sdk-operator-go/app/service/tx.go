package service

import (
	"context"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"

	"github.com/bianjieai/ddc-sdk-go/ddc-sdk-operator-go/app/handler"
	"github.com/bianjieai/ddc-sdk-go/ddc-sdk-operator-go/pkg/log"
	types2 "github.com/bianjieai/ddc-sdk-go/ddc-sdk-operator-go/pkg/types"
)

type TxService struct {
}

// GetTransByHash
// @Description: 运营方或平台方根据交易哈希对交易信息进行查询
// @receiver t
// @param txHash: 交易哈希
// @return *types.Transaction: 交易信息
// @return bool： pending状态
// @return error
func (t TxService) GetTransByHash(txHash string) (*types.Transaction, bool, error) {
	transaction, isPending, err := handler.GetConn().TransactionByHash(context.Background(), common.HexToHash(txHash))
	if err != nil {
		log.Error.Printf("failed to execute GetTransByHash: %v", err.Error())
		return transaction, isPending, types2.NewSDKError(types2.QueryError.Error(), err.Error())
	}
	transaction.Data()

	return transaction, isPending, nil
}

// GetTransReceipt
// @Description: 运营方或平台方根据交易哈希对交易回执信息进行查询。
// @receiver t
// @param txHash: 交易哈希
// @return string： 交易回执
// @return error
func (t TxService) GetTransReceipt(txHash string) (*types.Receipt, error) {
	receipt, err := handler.GetConn().TransactionReceipt(context.Background(), common.HexToHash(txHash))
	if err != nil {
		log.Error.Printf("failed to execute GetTransReceipt: %v", err.Error())
		return nil, types2.NewSDKError(types2.QueryError.Error(), err.Error())
	}

	return receipt, nil
}

// GetTransStatus
// @Description: 运营方或平台方根据交易哈希查询交易状态是否成功
// @receiver t
// @param txHash: 交易哈希
// @return bool：交易是否成功
// @return error
func (t TxService) GetTransStatus(txHash string) (bool, error) {
	receipt, err := handler.GetConn().TransactionReceipt(context.Background(), common.HexToHash(txHash))
	if err != nil {
		log.Error.Printf("failed to execute GetTransStatus: %v", err.Error())
		return false, types2.NewSDKError(types2.QueryError.Error(), err.Error())
	}

	if receipt.Status == 0 {
		return false, nil
	} else {
		return true, nil
	}
}

// GetTimeByTxHash
// @Description: 通过txHash查询所在块的时间
// @receiver t
// @param txHash 交易哈希
// @return uint64 所在块的时间
// @return error
func (t TxService) GetTimeByTxHash(txHash string) (uint64, error) {
	receipt, err := handler.GetConn().TransactionReceipt(context.Background(), common.HexToHash(txHash))
	if err != nil {
		log.Error.Printf("failed to get TransactionReceipt: %v", err.Error())
		return 0, types2.NewSDKError(types2.QueryError.Error(), err.Error())
	}
	block, err := handler.GetConn().BlockByNumber(context.Background(), receipt.BlockNumber)
	if err != nil {
		log.Error.Printf("failed to execute BlockByNumber: %v", err.Error())
		return 0, types2.NewSDKError(types2.QueryError.Error(), err.Error())
	}

	return block.Time(), nil
}
