package test

import (
	"fmt"
	"testing"
)

func TestGetTransByHash(t *testing.T) {
	fmt.Println(client.GetTransByHash("0xd5cb1c52265379bf0d3a3e900f9bb54e45ce5a7188041e32136aa9bcb3488308"))
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
