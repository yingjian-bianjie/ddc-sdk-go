package test

import (
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"math/big"
	"testing"
)

var ddc1155Service = client.GetDDC1155Service()

func TestDDC1155_SafeMint(t *testing.T) {
	opts.From = common.HexToAddress(platform)
	hash, err := ddc1155Service.SafeMint(opts, platform, 2, "www.123.com", []byte{})
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(hash)
}

func TestSafeMintBatch(t *testing.T) {
	opts.From = common.HexToAddress(platform)
	amounts := []*big.Int{big.NewInt(2)}
	ddcURIs := []string{"123"}
	tx, err := ddc1155Service.SafeMintBatch(opts, platform, amounts, ddcURIs, []byte{})
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(tx.Hash())
}
func TestDDC1155_BalanceOf(t *testing.T) {
	fmt.Println(ddc1155Service.BalanceOf(platform, 10))
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
