package test

import (
	"github.com/bianjieai/ddc-sdk-go/ddc-sdk-platform-go/app"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
)

var (
	clientBuilder = app.DDCSdkClientBuilder{}
	//client        = clientBuilder.
	//		SetSignEventListener(new(SignListener)). //注册实现了签名接口的结构体
	//		SetGasPrice(1e10).                       //建议设置gasPrice>=1e10
	//		SetGatewayUrl("https://opbtest.bsngate.com:18602/api/0e346e1fb134477cafb6c6c2583ce3c4/evmrpc").
	//		SetAuthorityAddress("0xa7FC5B0F4A0085c5Ce689b919a866675Ce37B66b").
	//		SetChargeAddress("0x3BBb01B38958d4dbF1e004611EbB3c65979B0511").
	//		SetDDC721Address("0x3B09b7A00271C5d9AE84593850dE3A526b8BF96e").
	//		SetDDC1155Address("0xe5d3b9E7D16E03A4A1060c72b5D1cb7806DD9070").
	//		RegisterLog("./log.log"). //设置日志输出的文件路径
	//		Build()
	client = clientBuilder.
		SetSignEventListener(new(SignListener)).
		SetGasPrice(1).
		SetGatewayUrl("http://192.168.150.42:8545").
		SetAuthorityAddress("0x6a3B24042dA7Bb5F2CBF1BCB2ABE0C632590C580").
		SetChargeAddress("0x95aDFbA9050C5D886419334Ae478b9844f413eF2").
		//SetDDC721Address("0x1C917baf05863417391acCfe85d305Eae41401Ec").
		SetDDC721Address("0x02A25C69843E197e3063Ed848f6FEA512633CB8E").
		SetDDC1155Address("0x02A25C69843E197e3063Ed848f6FEA512633CB8E").
		RegisterLog("./log.log").
		Build()
	opts = new(bind.TransactOpts)

	genV1  = "0x4F3B9C50A29562E3B71144A62E8447AC0DE7EC33"
	genV1P = "0x199a46b9bb8a68a2c1c9bc160cf1ee9614585d088826c48635c707ef18775366"

	pl  = "0xEDEDA0F1C1FAF1F34C9126975267FC6F95BE75F9"
	plP = "0xDCF313CD386BCF3787ECF9968F603729B278C2A0DCA18E940E3EF59EA600FD8E"

	operator    = "0x02CEB40D892061D457E7FA346988D0FF329935DF"
	operatorPri = "0xE253AB375A5806FA331E7DB32EDE524BD7D998475A60C957806066F14F479C25"
	platform    = "0x7FAF93F524FFDD1FB36BEC0ED6A167E8CE55BC4E" //ddcId：2、4、6
	platformPri = "0xED43B9686AB520C896BC33A1461BAF163EDBF0DBC4D3199E77793A49B9BB2568"
	//无gas
	v1    = "0x07B7BE76ED588CCEFB4C4A573CB28A7D2A1403CC"
	v1Pri = "0x1B8C36A57CB8D7FA20594283498EF310DCA9DFECDF6E9FDD04E992A5DA164E0B"
	//无gas
	v2    = "0x918F7F275A6C2D158E5B76F769D3F1678958A334" //ddcId:5
	v2Pri = "0x2F6976C530CFD2D0CC19EFC1868BD6A0AA1886D0BFCFA5A59D9B8899BE9B7241"

	//账户和私钥的映射
	senderPri = map[string]string{
		genV1:    genV1P,
		pl:       plP,
		operator: operatorPri,
		platform: platformPri,
		v1:       v1Pri,
		v2:       v2Pri,
	}
)
