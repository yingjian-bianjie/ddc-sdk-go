package test

import (
	"fmt"
	"testing"
)

func TestGetTransByHash(t *testing.T) {
	fmt.Println(client.GetTransByHash("0xd5cb1c52265379bf0d3a3e900f9bb54e45ce5a7188041e32136aa9bcb3488308"))
}
func TestGetTransReceipt(t *testing.T) {
	receipt, err := client.GetTransReceipt("0x66058dd25516834fcefd57f150411016ddf015e892923df8d5f17db13480b048")
	if err != nil {
		return
	}
	fmt.Println(receipt.BlockNumber)
}
func TestGetTransStatus(t *testing.T) {
	//736001537064
	//735901636239
	fmt.Println(client.GetTransStatus("0xc0ae24065bc23af7de30a22c58ed3bb82936dc9d17334ff742e784b172f439eb"))
}
