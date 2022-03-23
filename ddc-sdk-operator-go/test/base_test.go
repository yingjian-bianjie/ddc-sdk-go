package test

import (
	"fmt"
	"testing"

	"github.com/ethereum/go-ethereum/common"

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
