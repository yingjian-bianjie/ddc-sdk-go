package test

import (
	"fmt"
	"github.com/bianjieai/ddc-sdk-go/evm-ddc-sdk-platform/pkg/log"
	"testing"

	"github.com/ethereum/go-ethereum/common"
)

var ddc721Service = client.GetDDC721Service()

func TestMint(t *testing.T) {
	opts.From = common.HexToAddress(platform)
	tx, err := ddc721Service.Mint(opts, platform, "www.123.com")
	if err != nil {
		log.Error.Println(err.Error())
		return
	}
	fmt.Println(tx.Hash())
}

func TestSafeMint(t *testing.T) {
	opts.From = common.HexToAddress(platform)
	//fmt.Println("0xAAAA:",opts.From.String())
	address, _ := ddc721Service.HexToBech32(opts.From.String())
	fmt.Println("iaaxxx:", address)
	//oXaddr,_ := ddc721Service.Bech32ToHex(address)
	//fmt.Println("0xXXX:",oXaddr)
	tx, err := ddc721Service.SafeMint(opts, platform, "www.haha.com", []byte{})
	if err != nil {
		log.Error.Println(err.Error())
		return
	}
	fmt.Println(tx.Hash())
}

func TestDDC721BalanceOf(t *testing.T) {
	fmt.Println(ddc721Service.BalanceOf(platform))
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
	opts.From = common.HexToAddress(platform)
	fmt.Println(ddc721Service.TransferFrom(opts, platform, v2, 5))
}

func TestSafeTransferFrom(t *testing.T) {
	opts.From = common.HexToAddress(platform)
	fmt.Println(ddc721Service.SafeTransferFrom(opts, platform, v2, 519, []byte{}))
}

func TestOwnerOf(t *testing.T) {
	fmt.Println(ddc721Service.OwnerOf(1809))
}

func TestSetURI(t *testing.T) {
	opts.From = common.HexToAddress(platform)
	fmt.Println(ddc721Service.SetURI(opts, 1809, "www.wahaha.com"))
}

func TestName(t *testing.T) {
	fmt.Println(ddc721Service.Name())
}

func TestDDCURI(t *testing.T) {
	fmt.Println(ddc721Service.DdcURI(1809))
}

func TestSymbol(t *testing.T) {
	fmt.Println(ddc721Service.Symbol())
}

func TestParseTransfer(t *testing.T) {
	receipt, err := client.GetTxReceipt("0x3eb1854f20dbd35405adee6f72e1fba1576b1cce245dbab279d969ca5788aa12")
	if err != nil {
		t.Fatal(err)
	}
	for _, vlog := range receipt.Logs {
		fmt.Println(ddc721Service.ParseTransfer(*vlog))
	}

}
