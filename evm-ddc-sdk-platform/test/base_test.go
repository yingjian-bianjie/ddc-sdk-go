package test

import (
	"fmt"
	"github.com/bianjieai/ddc-sdk-go/evm-ddc-sdk-platform/app/handler"
	"testing"

	"github.com/ethereum/go-ethereum/common"

	"github.com/bianjieai/ddc-sdk-go/evm-ddc-sdk-platform/app/constant"
)

func TestEstimateGasLimit(t *testing.T) {
	opts.From = common.HexToAddress(platform)
	fmt.Println(client.EstimateGasLimit(opts, constant.ContrDDC721, constant.FuncMint, common.HexToAddress(platform), "www.123.com"))
}
func TestGetDDCId(t *testing.T) {
	//fmt.Println(client.DDCIdByHash("0x8ae229265cb29093dc01c7039aadaf901be8582660c691befbb352b6f5e61b9f"))
	//fmt.Println(client.DDCIdByHash("0x13bfa73224df9fcd422a210ee6d5b97885107c7b4b44e9e2247ee1a589b58e34"))
	fmt.Println(client.DDCIdByHash("0x557d89460de90b3e5e39b987f46cd20f0f3b175e4b1f86a83c667fb4f63795f4"))
}
func TestCreateAccount(t *testing.T) {
	fmt.Println(client.CreateAccount())
}

func TestHexToBech32(t *testing.T) {
	fmt.Println(client.HexToBech32("0x10B8FF6CB3C71A67838BEF96300528206DF4F884"))
}

func TestConn(t *testing.T) {
	handler.GetConn()
}
