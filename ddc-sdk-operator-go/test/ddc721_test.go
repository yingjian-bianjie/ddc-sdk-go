package test

import (
	"fmt"
	"testing"

	"github.com/ethereum/go-ethereum/common"
)

var ddc721Service = client.GetDDC721Service()

func TestDDC721BalanceOf(t *testing.T) {
	fmt.Println(ddc721Service.BalanceOf(platform))
}

func TestApprove(t *testing.T) {
	opts.From = common.HexToAddress(platform)
	fmt.Println(ddc721Service.Approve(opts, v1, 17))
}

func TestGetApprove(t *testing.T) {
	fmt.Println(ddc721Service.GetApprove(3))
}

func TestIsApprovedForAll(t *testing.T) {
	fmt.Println(ddc721Service.IsApprovedForAll(platform, v1))
}

func TestTransferFrom(t *testing.T) {
	opts.From = common.HexToAddress(platform)
	transaction, err := ddc721Service.TransferFrom(opts, platform, platform, 17)
	if err != nil {
		return
	}
	fmt.Println(transaction.Hash())
}

func TestSafeTransferFrom(t *testing.T) {
	opts.From = common.HexToAddress(platform)
	fmt.Println(ddc721Service.SafeTransferFrom(opts, platform, v2, 5, []byte{}))
}

func TestFreeze(t *testing.T) {
	opts.From = common.HexToAddress(platform)
	freeze, err := ddc721Service.Freeze(opts, 5)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(freeze)
}

func TestOwnerOf(t *testing.T) {
	owner, err := ddc721Service.OwnerOf(17)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(owner)
}

func TestDDCURI(t *testing.T) {
	fmt.Println(ddc721Service.DdcURI(1))
}

func TestSymbol(t *testing.T) {
	fmt.Println(ddc721Service.Symbol())
}
