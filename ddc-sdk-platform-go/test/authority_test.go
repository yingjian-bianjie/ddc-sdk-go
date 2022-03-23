package test

import (
	"fmt"
	"testing"

	"github.com/ethereum/go-ethereum/common"
)

var authorityService = client.GetAuthorityService()

func TestGetAccount(t *testing.T) {
	fmt.Println(authorityService.GetAccount(platform))

	fmt.Println(authorityService.GetAccount(genV1))
}
func TestUpdateAccState(t *testing.T) {
	opts.From = common.HexToAddress(platform)
	fmt.Println(authorityService.UpdateAccState(opts, v1, 1, true))
}
