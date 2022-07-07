package test

import (
	"crypto/ecdsa"
	"log"
	"strings"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
)

type SignListener struct {
}

// SignEvent
// @Description: 自定义的签名方法
// @receiver s
// @param sender 调用者的账户地址
// @param tx 待签名的交易
// @return *types.Transaction 签名好的交易
// @return error
func (s *SignListener) SignEvent(sender common.Address, tx *types.Transaction) (*types.Transaction, error) {

	// 提取私钥
	privateKey, err := StringToPrivateKey(senderPri["0x"+strings.ToUpper(sender.Hex()[2:])])
	if err != nil {
		log.Fatalf("StringToPrivateKey failed:%v", err)
	}
	// 签名,m
	signTx, err := types.SignTx(tx, types.HomesteadSigner{}, privateKey)

	return signTx, err
}

// StringToPrivateKey
// @Description: 从明文的私钥字符串转换成ECDSA格式的私钥
// @param privateKeyStr 明文的私钥字符串
// @return *ecdsa.PrivateKey
// @return error
func StringToPrivateKey(privateKeyStr string) (*ecdsa.PrivateKey, error) {
	privateKeyByte, err := hexutil.Decode(privateKeyStr)
	if err != nil {
		return nil, err
	}
	privateKey, err := crypto.ToECDSA(privateKeyByte)
	if err != nil {
		return nil, err
	}
	return privateKey, nil
}
