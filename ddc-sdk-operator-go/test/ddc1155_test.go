package test

import (
	"fmt"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
)

var ddc1155Service = client.GetDDC1155Service()

func TestDDC1155_BalanceOf(t *testing.T) {
	fmt.Println(ddc1155Service.BalanceOf(platform, 10))
}
func TestDDC1155Freeze(t *testing.T) {
	opts.From = common.HexToAddress(platform)
	fmt.Println(ddc1155Service.Freeze(opts, 5))
}
func TestDDC1155DDCURI(t *testing.T) {
	fmt.Println(ddc1155Service.DdcURI(1))
}
func TestDDC1155(t *testing.T) {
	owners := []common.Address{common.HexToAddress(platform)}
	ddcIds := []*big.Int{big.NewInt(1)}

	fmt.Println(ddc1155Service.BalanceOfBatch(owners, ddcIds))
}
func TestDDC1155SetURI(t *testing.T) {
	opts.From = common.HexToAddress(platform)
	fmt.Println(ddc1155Service.SetURI(opts, platform, 2, "www.123.com"))
}
