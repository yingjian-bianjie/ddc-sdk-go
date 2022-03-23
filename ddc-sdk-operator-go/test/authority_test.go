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
	fmt.Println(authorityService.GetAccount("0x2F4AB3D3D827E61D71BDCF33A5C52090BBE7F0C9"))
}
func TestAddFunction(t *testing.T) {
	//safeMint
	fmt.Println(authorityService.AddFunction(opts, 2, config.Info.Ddc721Address().Hex(), [4]byte{0xf6, 0xdd, 0xa9, 0x36}))
}
func TestUpdateAccState(t *testing.T) {
	opts.From = common.HexToAddress(operator)
	fmt.Println(authorityService.UpdateAccState(opts, platform, 1, true))
}
