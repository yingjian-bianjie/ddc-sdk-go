package test

import (
	"fmt"
	"testing"

	"github.com/ethereum/go-ethereum/common"

	"github.com/bianjieai/ddc-sdk-go/ddc-sdk-platform-go/pkg/contracts"
)

func TestGetBlockByNumber(t *testing.T) {

	fmt.Println(client.GetBlockByNumber(301751))
}
func TestGetEvents(t *testing.T) {

	fmt.Println(client.GetBlockEvents(301751))
}
func TestGetEvent(t *testing.T) {

	events, _ := client.GetTxEvents(common.HexToHash("0xfff913ba0f63fc3b5ead83453d8cd8ad95c142735ef8f71c285e267112fe1676"))

	for _, e := range events {
		res, ok := e.(*contracts.DDC721Transfer)
		if !ok {
			continue
		}
		fmt.Println(res.DdcId)
	}
}
