package test

import (
	"fmt"
	"testing"

	"github.com/ethereum/go-ethereum/common"

	"github.com/bianjieai/ddc-sdk-go/ddc-sdk-platform-go/app/constant"
	"github.com/bianjieai/ddc-sdk-go/ddc-sdk-platform-go/app/service"
)

func TestEstimateGasLimit(t *testing.T) {
	opts.From = common.HexToAddress(platform)
	base := service.Base{}
	fmt.Println(base.EstimateGasLimit(opts, constant.ContrDDC721, constant.FuncMint, common.HexToAddress(platform), "www.123.com"))
}
func TestGetDDCId(t *testing.T) {
	base := service.Base{}
	fmt.Println(base.DDCIdByHash("0xa3f568301f4853d8cb1e9b81d91271b2c973d60206f4bc2759ca761ebe115bf3"))
}
func TestCreateAccount(t *testing.T){
	base := service.Base{}
	fmt.Println(base.CreateAccount())
//	&{0xa0e06e3e3cC56cc863Dd518C945205682dF3A5a2 98da12a135fb8957d567e5b31f80c9c24791a3dda35401a3c4ed091aee63008a48dcfc7e8bcc20af29ff56b3f95d6d5fea2bf781ba8d7464352df495256d741b 099a6c5066c6d5cf73be09367b81419cdfe1a79029e28786f9227693cc73043a property bounce fitness author soap baby cotton runway jealous leopard tongue artwork} <nil>
}
