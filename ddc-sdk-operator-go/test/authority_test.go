package test

import (
	"fmt"
	"testing"

	"github.com/ethereum/go-ethereum/common"

	"github.com/bianjieai/ddc-sdk-go/ddc-sdk-operator-go/config"
)

var authorityService = client.GetAuthorityService()

func TestAddConsumerByOperator(t *testing.T) {
	opts.From = common.HexToAddress(operator)
	fmt.Println(authorityService.AddAccountByOperator(opts, platform, "ddcplatform", "did:ddcplatform", ""))
}
func TestGetAccount(t *testing.T) {
	fmt.Println(authorityService.GetAccount(platform))
}
func TestAddFunction(t *testing.T) {
	opts.From = common.HexToAddress(operator)
	transaction, err := authorityService.AddFunction(opts, 1, config.Info.Ddc1155Address().Hex(), [4]byte{0xd7, 0xa7, 0x8d, 0xb8})
	if err != nil {
		return
	}
	//safeMint
	fmt.Println(transaction.Hash())
}
func TestUpdateAccState(t *testing.T) {
	opts.From = common.HexToAddress(operator)
	fmt.Println(authorityService.UpdateAccState(opts, platform, 1, true))
}
