package test

import (
	"fmt"
	"testing"

	"github.com/ethereum/go-ethereum/common"

	"github.com/bianjieai/ddc-sdk-go/ddc-sdk-operator-go/config"
)

var chargeService = client.GetChargeService()

func TestSelfRecharge(t *testing.T) {
	opts.From = common.HexToAddress(operator)
	fmt.Println(chargeService.SelfRecharge(opts, 1e1))
}
func TestRecharge(t *testing.T) {
	opts.From = common.HexToAddress(operator)
	opts.GasLimit = 1e7
	tx, err := chargeService.Recharge(opts, platform, 1e11)
	if err != nil {
		return
	}
	fmt.Println(tx.Hash())
}
func TestBalanceOf(t *testing.T) {

	balance, _ := chargeService.BalanceOf(operator)
	fmt.Println("balance:", balance)
}
func TestQueryFee(t *testing.T) {
	fmt.Println(chargeService.QueryFee(config.Info.Ddc721Address().Hex(), [4]byte{0x23, 0xb8, 0x72, 0xdd}))
}
func TestSetFee(t *testing.T) {
	opts.From = common.HexToAddress(operator)
	fmt.Println(chargeService.SetFee(opts, config.Info.AuthorityAddress().String(), [4]byte{0xe3, 0xf0, 0x0c, 0x3a}, 10))
}
