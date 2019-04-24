package launcher

import (
	"database/sql"
	"fmt"
	"github.com/HydroProtocol/hydro-sdk-backend/sdk/ethereum"
	"github.com/HydroProtocol/hydro-sdk-backend/test"
	"github.com/HydroProtocol/hydro-sdk-backend/utils"
	"github.com/shopspring/decimal"
	"os"
	"testing"
	"time"
)

func TestLauncher_Run(t *testing.T) {
	test.PreTest()

	launchLog := LaunchLog{
		ItemType:    "hydroTrade",
		ItemID:      1,
		Status:      "created",
		Hash:        sql.NullString{},
		BlockNumber: sql.NullInt64{},

		From:     "0x93388b4efe13b9b18ed480783c05462409851547",
		To:       "0x179fd00c328d4ecdb5043c8686d377a24ede9d11",
		Value:    decimal.Zero,
		GasLimit: 1000000,

		// use a reasonable gasPrice instead of static one
		// There are several options
		//   1. use eth gas station api (https://ethgasstation.info/)
		//   2. use eth_gasPrice Ethereum json rpc call
		//   3. develop your own eth gasPrice tracker
		GasPrice: decimal.NullDecimal{Decimal: utils.StringToDecimal("10000000000"), Valid: true},
		Data:     "0x884dad2e0000000000000000000000003870b6f2c0b723f4855d8ad53ab7599b02d4df840000000000000000000000000000000000000000000000001bc16d674ec8000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000572255eb17edfc402010102540be3ff006400c80064000000000007f130000000000000000000001c0100000000000000000000000000000000000000000000000000000000000019cef14892021d56d31b8f6ca6ed99ab89ac918a2d8e2e9d034b14ccf1dfa17f27601ed6b0d1a7fd64f6e62c2fea27580f36521d2609baaf2f9921ecd6cb761b00000000000000000000000000000000000000000000000000000000000001a000000000000000000000000000000000000000000000000000000000000003c0000000000000000000000000fe1e07852eb0fa0df66843e84a41da212b455e980000000000000000000000009712e6cadf82d1902088ef858502ca17261bb89300000000000000000000000093388b4efe13b9b18ed480783c05462409851547000000000000000000000000000000000000000000000000000000000000000200000000000000000000000085cf54dd216997bcf324c72aa1c845be2f0592990000000000000000000000000000000000000000000000000de0b6b3a76400000000000000000000000000000000000000000000000000000c7d713b49da00000000000000000000000000000000000000000000000000000572255eb17edfc402000002540be3ff006400c80064000000000002cc82000000000000000000001b01000000000000000000000000000000000000000000000000000000000000a096a43f547bd361d79f965723604965cf45fb9e67754c67accb100eef344804571c23cc135c501eb72d137eeb6f430bb3a62da2122610223e3d80b99cecdff900000000000000000000000085cf54dd216997bcf324c72aa1c845be2f0592990000000000000000000000000000000000000000000000000de0b6b3a76400000000000000000000000000000000000000000000000000000b1a2bc2ec5000000000000000000000000000000000000000000000000000000572255eb17edfc402000002540be3ff006400c8006400000000000c8a85000000000000000000001b01000000000000000000000000000000000000000000000000000000000000583d5bece6c5b0ef5b1807f681270c864e987ffeed66578d39343ab6996851c5371013591a920621cd32c8bfd3f7d89d8aef36dfed7c636e08a0b0871491e76500000000000000000000000000000000000000000000000000000000000000020000000000000000000000000000000000000000000000000de0b6b3a76400000000000000000000000000000000000000000000000000000de0b6b3a7640000",

		ExecutedAt: time.Now(),
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
	}

	hydro := ethereum.NewEthereumHydro(os.Getenv("HSK_BLOCKCHAIN_RPC_URL"), os.Getenv("HSK_HYBRID_EXCHANGE_ADDRESS"))
	signService := NewDefaultSignService(os.Getenv("HSK_RELAYER_PK"), hydro.GetTransactionCount)

	signedTransaction := signService.Sign(&launchLog)

	hash, err := hydro.SendRawTransaction(signedTransaction)

	fmt.Println(hash, err)
	// if err != nil {
	// 	t.Log(err)
	// 	t.Errorf("sendRawTransaction err")
	// }
}
