package service

import (
	"context"
	"math/big"
	"strconv"

	abi2 "github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"

	"github.com/bianjieai/ddc-sdk-go/evm-ddc-sdk-platform/app/constant"
	"github.com/bianjieai/ddc-sdk-go/evm-ddc-sdk-platform/app/handler"
	"github.com/bianjieai/ddc-sdk-go/evm-ddc-sdk-platform/app/models/dto"
	"github.com/bianjieai/ddc-sdk-go/evm-ddc-sdk-platform/config"
	"github.com/bianjieai/ddc-sdk-go/evm-ddc-sdk-platform/pkg/contracts"
	"github.com/bianjieai/ddc-sdk-go/evm-ddc-sdk-platform/pkg/log"
	types2 "github.com/bianjieai/ddc-sdk-go/evm-ddc-sdk-platform/pkg/types"
)

type BlockService struct {
}

// GetBlockByNumber
// @Description: 运营方或平台方根据区块高度对区块信息进行查询，并解析区块数据返回给运营方或平台方
// @receiver b
// @param blockNumber： 区块高度
// @return *types2.Block： 区块信息
// @return error
func (b *BlockService) GetBlockByNumber(blockNumber int64) (*types.Block, error) {

	block, err := config.Info.Conn().BlockByNumber(context.Background(), big.NewInt(blockNumber))
	if err != nil {
		log.Error.Printf("failed to execute GetBlockByNumber: %v", err.Error())
		return nil, types2.NewSDKError(types2.QueryError.Error(), err.Error())
	}

	return block, nil
}

// GetBlockEvents
// @Description: 获取指定区块内的所有events
// @receiver b
// @param blockNumber 块高
// @return *dto.BlockEventBean 事件和时间戳的实体
// @return error
func (b *BlockService) GetBlockEvents(blockNumber int64) (*dto.BlockEventBean, error) {

	//查找区块中所有交易
	block, err := config.Info.Conn().BlockByNumber(context.Background(), big.NewInt(blockNumber))
	if err != nil {
		log.Error.Printf("failed to execute BlockByNumber: %v", err.Error())
		return nil, types2.NewSDKError(types2.QueryError.Error(), err.Error())
	}
	var blockEvents []interface{}
	//查找每笔交易中的event
	for _, tx := range block.Transactions() {
		txEvents, err := b.GetTxEvents(tx.Hash())
		if err != nil {
			log.Error.Printf("failed to get events of Tx: %v,error: %v", tx.Hash(), err.Error())
			return nil, types2.NewSDKError(types2.QueryError.Error(), err.Error())
		}
		blockEvents = append(blockEvents, txEvents)
	}
	time := block.Time()

	return &dto.BlockEventBean{Events: blockEvents, Timestamp: strconv.FormatUint(time, 10)}, nil
}

// GetBlockEvents
// @Description: 获取指定区块内的所有ddcId和回执
// @receiver b
// @param blockNumber 块高
// @return *dto.BlockDdcInfoBean ddc信息和时间戳的实体
// @return error
func (b *BlockService) GetBlockDdcs(blockNumber int64) (*dto.BlockDdcInfoBean, error) {

	//查找区块中所有交易
	block, err := config.Info.Conn().BlockByNumber(context.Background(), big.NewInt(blockNumber))
	if err != nil {
		log.Error.Printf("failed to execute BlockByNumber: %v", err.Error())
		return nil, types2.NewSDKError(types2.QueryError.Error(), err.Error())
	}
	var blockDdcs []dto.DdcInfoBean
	//查找每笔交易中的event
	for _, tx := range block.Transactions() {
		_, txReceipt, ddcId, err := b.GetTxEventsWithReceiptAndDdcId(tx.Hash())
		if err != nil {
			log.Error.Printf("failed to get events,receipt,ddcId of Tx: %v,error: %v", tx.Hash(), err.Error())
			return nil, types2.NewSDKError(types2.QueryError.Error(), err.Error())
		}
		blockDdcs = append(blockDdcs, dto.DdcInfoBean{
			Receipt: txReceipt,
			DdcId:   ddcId,
		})
	}
	time := block.Time()

	return &dto.BlockDdcInfoBean{DdcInfos: blockDdcs, Timestamp: strconv.FormatUint(time, 10)}, nil
}

// GetTxEvents
// @Description: 获取指定交易中的所有events
// @receiver b
// @param txHash 交易哈希
// @return events 查询出的事件
// @return err
func (b *BlockService) GetTxEvents(txHash common.Hash) (events []interface{}, err error) {
	//获取对应的交易回执
	receipt, err := config.Info.Conn().TransactionReceipt(context.Background(), txHash)
	if err != nil {
		log.Error.Printf("failed to execute TransactionReceipt: %v", err.Error())
		return nil, types2.NewSDKError(types2.QueryError.Error(), err.Error())
	}
	var event interface{}
	var abi *abi2.ABI
	//获取交易的logs中的所有log对应的event
	for _, l := range receipt.Logs {
		//1.匹配对应的合约
		switch l.Address {
		case config.Info.AuthorityAddress():
			{
				authority := handler.GetAuthority()

				abi, err = contracts.AuthorityMetaData.GetAbi()
				if err != nil {
					log.Error.Printf("failed to get Authority abi: %v", err.Error())
					return nil, err
				}
				//2.匹配对应的事件
				switch l.Topics[0] {
				case abi.Events[constant.EventAddAccount].ID:
					event, err = authority.ParseAddAccount(*l)
				case abi.Events[constant.EventUpdateAccountState].ID:
					event, err = authority.ParseUpdateAccountState(*l)
				case abi.Events[constant.EventCrossPlatformApproval].ID:
					event, err = authority.ParseCrossPlatformApproval(*l)
				}
			}
		case config.Info.ChargeAddress():
			{
				charge := handler.GetCharge()
				abi, err = contracts.ChargeMetaData.GetAbi()
				if err != nil {
					log.Error.Printf("failed to get Charge abi: %v", err.Error())
					return nil, err
				}
				switch l.Topics[0] {
				case abi.Events[constant.EventRecharge].ID:
					event, err = charge.ParseRecharge(*l)
				case abi.Events[constant.EventPay].ID:
					event, err = charge.ParsePay(*l)
				case abi.Events[constant.EventSetFee].ID:
					event, err = charge.ParseSetFee(*l)
				case abi.Events[constant.EventDelFee].ID:
					event, err = charge.ParseDelFee(*l)
				case abi.Events[constant.EventDelDDC].ID:
					event, err = charge.ParseDelDDC(*l)
				}
			}
		case config.Info.Ddc721Address():
			{
				ddc721 := handler.GetDDC721()
				abi, err = contracts.DDC721MetaData.GetAbi()
				if err != nil {
					log.Error.Printf("failed to get DDC721 abi: %v", err.Error())
					return nil, err
				}
				switch l.Topics[0] {
				case abi.Events[constant.EventTransfer].ID:
					event, err = ddc721.ParseTransfer(*l)
				case abi.Events[constant.EventPay].ID:
					event, err = ddc721.ParseOwnershipTransferred(*l)
				case abi.Events[constant.EventEnterBlacklist].ID:
					event, err = ddc721.ParseEnterBlacklist(*l)
				case abi.Events[constant.EventExitBlacklist].ID:
					event, err = ddc721.ParseExitBlacklist(*l)
				case abi.Events[constant.EventSetURI].ID:
					event, err = ddc721.ParseSetURI(*l)
				}
			}
		case config.Info.Ddc1155Address():
			{
				ddc1155 := handler.GetDDC1155()
				abi, err = contracts.DDC1155MetaData.GetAbi()
				if err != nil {
					log.Error.Printf("failed to get DDC1155 abi: %v", err.Error())
					return nil, err
				}
				switch l.Topics[0] {
				case abi.Events[constant.EventTransferSingle].ID:
					event, err = ddc1155.ParseTransferSingle(*l)
				case abi.Events[constant.EventTransferBatch].ID:
					event, err = ddc1155.ParseTransferBatch(*l)
				case abi.Events[constant.EventEnterBlacklist].ID:
					event, err = ddc1155.ParseEnterBlacklist(*l)
				case abi.Events[constant.EventExitBlacklist].ID:
					event, err = ddc1155.ParseExitBlacklist(*l)
				case abi.Events[constant.EventSetURI].ID:
					event, err = ddc1155.ParseSetURI(*l)
				}
			}
		}
		if err != nil {
			log.Error.Printf("failed to parse event: %v", err.Error())
			return nil, err
		}
		events = append(events, event)
	}

	return
}

// GetTxEventsWithReceipt
// @Description: 获取指定交易中的所有events、回执信息、ddcID
// @receiver b
// @param txHash 交易哈希
// @return events 查询出的事件
// @return err
func (b *BlockService) GetTxEventsWithReceiptAndDdcId(txHash common.Hash) (events []interface{},
	txReceipt *types.Receipt, ddcId uint64, err error) {
	//获取对应的交易回执
	receipt, err := config.Info.Conn().TransactionReceipt(context.Background(), txHash)
	if err != nil {
		log.Error.Printf("failed to execute TransactionReceipt: %v", err.Error())
		return nil, nil, 0, types2.NewSDKError(types2.QueryError.Error(), err.Error())
	}
	txReceipt = receipt
	var event interface{}
	//获取交易的logs中的所有log对应的event
	for _, l := range receipt.Logs {
		//1.匹配对应的合约
		switch l.Address {
		case config.Info.AuthorityAddress():
			{
				authority := handler.GetAuthority()
				abi, err := contracts.AuthorityMetaData.GetAbi()
				if err != nil {
					log.Error.Printf("failed to get Authoriy abi: %v", err.Error())
					return nil, txReceipt, 0, err
				}
				//2.匹配对应的事件
				switch l.Topics[0] {
				case abi.Events[constant.EventAddAccount].ID:
					event, err = authority.ParseAddAccount(*l)
				case abi.Events[constant.EventUpdateAccountState].ID:
					event, err = authority.ParseUpdateAccountState(*l)
				case abi.Events[constant.EventCrossPlatformApproval].ID:
					event, err = authority.ParseCrossPlatformApproval(*l)
				}
			}
		case config.Info.ChargeAddress():
			{
				charge := handler.GetCharge()
				abi, err := contracts.ChargeMetaData.GetAbi()
				if err != nil {
					log.Error.Printf("failed to get Charge abi: %v", err.Error())
					return nil, txReceipt, 0, err
				}
				switch l.Topics[0] {
				case abi.Events[constant.EventRecharge].ID:
					event, err = charge.ParseRecharge(*l)
				case abi.Events[constant.EventPay].ID:
					event, err = charge.ParsePay(*l)
				case abi.Events[constant.EventSetFee].ID:
					event, err = charge.ParseSetFee(*l)
				case abi.Events[constant.EventDelFee].ID:
					event, err = charge.ParseDelFee(*l)
				case abi.Events[constant.EventDelDDC].ID:
					event, err = charge.ParseDelDDC(*l)
				}
			}
		case config.Info.Ddc721Address():
			{
				ddc721 := handler.GetDDC721()
				abi, err := contracts.DDC721MetaData.GetAbi()
				if err != nil {
					log.Error.Printf("failed to get DDC721 abi: %v", err.Error())
					return nil, txReceipt, 0, err
				}
				switch l.Topics[0] {
				case abi.Events[constant.EventTransfer].ID:
					event, err = ddc721.ParseTransfer(*l)
				case abi.Events[constant.EventPay].ID:
					event, err = ddc721.ParseOwnershipTransferred(*l)
				case abi.Events[constant.EventEnterBlacklist].ID:
					event, err = ddc721.ParseEnterBlacklist(*l)
				case abi.Events[constant.EventExitBlacklist].ID:
					event, err = ddc721.ParseExitBlacklist(*l)
				case abi.Events[constant.EventSetURI].ID:
					event, err = ddc721.ParseSetURI(*l)
				}
			}
		case config.Info.Ddc1155Address():
			{
				ddc1155 := handler.GetDDC1155()
				abi, err := contracts.DDC1155MetaData.GetAbi()
				if err != nil {
					log.Error.Printf("failed to get DDC1155 abi: %v", err.Error())
					return nil, txReceipt, 0, err
				}
				switch l.Topics[0] {
				case abi.Events[constant.EventTransferSingle].ID:
					event, err = ddc1155.ParseTransferSingle(*l)
				case abi.Events[constant.EventTransferBatch].ID:
					event, err = ddc1155.ParseTransferBatch(*l)
				case abi.Events[constant.EventEnterBlacklist].ID:
					event, err = ddc1155.ParseEnterBlacklist(*l)
				case abi.Events[constant.EventExitBlacklist].ID:
					event, err = ddc1155.ParseExitBlacklist(*l)
				case abi.Events[constant.EventSetURI].ID:
					event, err = ddc1155.ParseSetURI(*l)
				}
			}
		}
		if err != nil {
			log.Error.Printf("failed to parse event: %v", err.Error())
			return nil, txReceipt, 0, err
		}
		if res, ok := event.(*contracts.DDC721Transfer); ok {
			ddcId = res.DdcId.Uint64()
		}
		events = append(events, event)
	}
	return
}
