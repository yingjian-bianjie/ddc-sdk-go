package test

import (
	"fmt"
	"github.com/bianjieai/ddc-sdk-go/ddc-sdk-platform-go/pkg/log"
	"testing"

	"github.com/ethereum/go-ethereum/common"
)

var ddc721Service = client.GetDDC721Service()

func TestMint(t *testing.T) {
	opts.From = common.HexToAddress(pl)
	tx, err := ddc721Service.Mint(opts, genV1, "www.123.com")
	if err != nil {
		log.Error.Println(err.Error())
		return
	}
	fmt.Println(tx.Hash())
}

func TestSafeMint(t *testing.T) {
	opts.From = common.HexToAddress(platform)
	transaction, err := ddc721Service.SafeMint(opts, "0x83F3B42F4E1795999089FAB8F117DAB18A496577", "www.123.com", []byte{})
	if err != nil {
		return
	}
	fmt.Println(transaction.Hash().Hex())
}

func TestDDC721BalanceOf(t *testing.T) {
	fmt.Println(ddc721Service.BalanceOf("0x83F3B42F4E1795999089FAB8F117DAB18A496577"))
}

func TestApprove(t *testing.T) {
	fmt.Println(ddc721Service.Approve(opts, v1, 1))
}

func TestGetApprove(t *testing.T) {
	fmt.Println(ddc721Service.GetApprove(3))
}

func TestIsApprovedForAll(t *testing.T) {
	fmt.Println(ddc721Service.IsApprovedForAll(platform, v1))
}

func TestTransferFrom(t *testing.T) {
	opts.From = common.HexToAddress(genV1)
	opts.GasLimit = 1e6
	fmt.Println(ddc721Service.TransferFrom(opts, genV1, pl, 102))
}

func TestSafeTransferFrom(t *testing.T) {
	opts.From = common.HexToAddress(platform)
	fmt.Println(ddc721Service.SafeTransferFrom(opts, platform, v2, 5, []byte{}))
}

func TestOwnerOf(t *testing.T) {
	fmt.Println(ddc721Service.OwnerOf(102))
}

func TestDDCURI(t *testing.T) {
	fmt.Println(ddc721Service.DdcURI(1))
}

func TestSymbol(t *testing.T) {
	fmt.Println(ddc721Service.Symbol())
}
