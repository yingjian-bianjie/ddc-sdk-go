package test

import (
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"testing"
	"time"

	"github.com/bianjieai/ddc-sdk-go/ddc-sdk-operator-go/app/constant"
	"github.com/bianjieai/ddc-sdk-go/ddc-sdk-operator-go/app/service"
)

func TestEstimateGasLimit(t *testing.T) {
	opts.From = common.HexToAddress(platform)
	base := service.Base{}
	fmt.Println(base.EstimateGasLimit(opts, constant.ContrDDC721, constant.FuncMint, common.HexToAddress(platform), "www.123.com"))
}
func TestGetDDCId(t *testing.T) {
	base := service.Base{}
	fmt.Println(base.DDCIdByHash("0xc806461847abf17c899a027edddeaa8c0ea57b64f10580f63c1c385f9aa3b282"))
}
func TestCreateAccount(t *testing.T) {
	base := service.Base{}
	fmt.Println(base.CreateAccount())
}
func TestAddrToBeach32(t *testing.T) {
	bech32, err := client.HexToBech32(common.Hex2Bytes("02CEB40D892061D457E7FA346988D0FF329935DF"))
	if err != nil {
		return
	}
	fmt.Println(bech32)
}
func TestConn(t *testing.T) {
	base := service.Base{}
	fmt.Println(base.DDCIdByHash("0xc806461847abf17c899a027edddeaa8c0ea57b64f10580f63c1c385f9aa3b282"))

	time.Sleep(time.Minute * 10)
	fmt.Println(ddc721Service.BalanceOf(platform))
}
