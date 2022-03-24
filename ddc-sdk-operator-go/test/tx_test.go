package test

import (
	"fmt"
	"testing"
)

func TestGetTransByHash(t *testing.T) {
	tx, _, err := client.GetTransByHash("0xb3b4a5075c237885685bb276177a9783f5364e647af4a2cee92e10144239dfd5")
	if err != nil {
		return
	}
	fmt.Println(tx.Gas(), tx.Cost(), tx.GasFeeCap(),tx.Value(),tx.GasPrice())
}
func TestGetTransReceipt(t *testing.T) {
	receipt, err := client.GetTransReceipt("0xb3b4a5075c237885685bb276177a9783f5364e647af4a2cee92e10144239dfd5")
	if err != nil {
		return
	}
	fmt.Println(receipt.BlockNumber)
}
func TestGetTransStatus(t *testing.T) {
	//736001537064
	//735901636239
	fmt.Println(client.GetTransStatus("0xb3b4a5075c237885685bb276177a9783f5364e647af4a2cee92e10144239dfd5"))
}
