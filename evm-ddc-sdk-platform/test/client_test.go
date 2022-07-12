package test

import (
	"github.com/bianjieai/ddc-sdk-go/evm-ddc-sdk-platform/app"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
)

var (
	clientBuilder = app.DDCSdkClientBuilder{}
	client        = clientBuilder.
		//SetSignEventListener(new(SignListener)). //注册实现了签名接口的结构体
		//SetGasPrice(1).                          //设置gasPrice
		//SetGatewayURL("http://192.168.150.42:8545").
		//SetAuthorityAddress("0xBcE9AA1924D7197C9C945e43638Bf589f91bcB71").
		//SetChargeAddress("0xF41b6185bFB22E2EFC5fB8395Fa3B952951E2d0b").
		//SetDDC721Address("0x74b6114d011891Ac21FD1d586bc7F3407c63c216").
		//SetDDC1155Address("0x9f7388e114DfDFAbAF8e4b881894E4C7e1b52C17").
		//RegisterLog("./log.log").
		//Build()
		SetSignEventListener(new(SignListener)). //注册实现了签名接口的结构体
		SetGasPrice(1).                          //设置gasPrice
		SetGatewayURL("https://opbningxia.bsngate.com:18602/api/9979c62195e54c43a7cfa4935a2c075e/evmrpc").
		SetAuthorityAddress("0xFa1d2d3EEd20C4E4F5b927D9730d9F4D56314B29").
		SetChargeAddress("0x0B8ae0e1b4a4Eb0a0740A250220eE3642d92dc4D").
		SetDDC721Address("0x354c6aF2cB870BEFEA8Ea0284C76e4A46B8F2870").
		SetDDC1155Address("0x0E762F4D11439B1130D402995328b634cB9c9973").
		RegisterLog("./log.log").
		Build()
		//SetGatewayURL("https://opbningxia.bsngate.com:18602/api/[ProjectId]/evmrpc")
		//SetGatewayAPIKey("x-api-key").
		//SetGatewayAPIValue("[ProjectKey]").
		//SetAuthorityAddress("0xa7FC5B0F4A0085c5Ce689b919a866675Ce37B66b").
		//SetChargeAddress("0x3BBb01B38958d4dbF1e004611EbB3c65979B0511").
		//SetDDC721Address("0x3B09b7A00271C5d9AE84593850dE3A526b8BF96e").
		//SetDDC1155Address("0xe5d3b9E7D16E03A4A1060c72b5D1cb7806DD9070").
		//RegisterLog("./log.log"). //设置日志输出的文件路径
		//Build()

	opts = new(bind.TransactOpts)

	platform    = "0x7FAF93F524FFDD1FB36BEC0ED6A167E8CE55BC4E"
	platformPri = "0xED43B9686AB520C896BC33A1461BAF163EDBF0DBC4D3199E77793A49B9BB2568" //对应的私钥
	//无gas
	v1    = "0x32B6034EEBF55CEEF908816BE6709B234A2376D2"
	v1Pri = "0xB97BCE72ED2DFBAC1501605591ABDFA3D1DBAE4EB3ED990B7EEAB003478C3CD0"
	//无gas
	v2    = "0x0DF4E53A8A273E7813F0D7D0B2C8510C1D546E06" //ddcID:5
	v2Pri = "0x1933ABA4A5525F648FEA7EB7D8E368D5E23F482BF7CD31A9FF9DEA48F607EC6D"

	//账户和私钥的映射
	senderPri = map[string]string{
		platform: platformPri,
		v1:       v1Pri,
		v2:       v2Pri,
	}
)
