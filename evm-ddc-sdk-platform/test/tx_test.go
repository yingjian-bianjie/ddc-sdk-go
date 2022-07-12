package test

import (
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"testing"
)

func TestGetTransByHash(t *testing.T) {
	fmt.Println(client.GetTxByHash("0xd5cb1c52265379bf0d3a3e900f9bb54e45ce5a7188041e32136aa9bcb3488308"))
}
func TestGetTransReceipt(t *testing.T) {
	fmt.Println(client.GetTxReceipt("0x10e8d4e5f30b05d268b1b61ba89d85d8c8010153b21db00efe14f4656bf0b92b"))
}
func TestGetTransStatus(t *testing.T) {
	fmt.Println(client.GetTxStatus("0xc848f2cd49658cca7963c13d9218872bc667d3d0c4f8d1fc6ffb14df3cd029c7"))
}

func TestGetTxEventsWithReceiptAndDdcId(t *testing.T) {
	_, receipt, id, err := client.GetTxEventsWithReceiptAndDdcId(common.HexToHash("0x19a8fed4e4a487d551f8fde8f959bcb3e691d01bc0ffac4dc40b246c6f2c2665"))
	if err != nil {
		t.Fatal(err.Error())
	}
	t.Log(receipt)
	t.Log(id)
}

func TestGetBlockDdcs(t *testing.T) {
	data, err := client.GetBlockDdcs(15057363)
	if err != nil {
		t.Fatal(err.Error())
	}
	t.Log(data)
}
