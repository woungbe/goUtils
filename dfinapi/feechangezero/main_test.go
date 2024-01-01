package checkbot

/*
일단 보관 용으로 가지고 있음 .

*/

import (
	"context"
	adminconfig "dfinForBinance/dfinAdmin/adminConfig"
	"dfinForBinance/dfinLIB/dfinUtil"
	"fmt"
	"os"
	"strings"
	"testing"
	"time"

	"github.com/adshao/go-binance/v2"
)

var binapi *binance.Client

func DefaultCheck() {
	dfinUtil.RandInitSeed()
	cnf := adminconfig.GetConfig()
	er := cnf.InitConfig(true) //실서비스일때는 fales 로
	if er != nil {
		fmt.Println("설정정보로드에러")
		fmt.Println(er)
		os.Exit(1)
	}
	binapi = GetFutuClient()
}

func GetFutuClient() *binance.Client {
	client := binance.NewClient(adminconfig.GetBrokerApiKey(), adminconfig.GetBrokerSecretKey())
	return client
}

func TestCallCheckUser(t *testing.T) {
	DefaultCheck()
	sql := "SELECT * FROM accountMasterReCycle"
	rows, err := adminconfig.GetMainMysqlDB().DBQuerySelect(sql)
	if err != nil {
		fmt.Println("err : ", err)
	}

	marketList := GetMarketList()

	for _, v := range rows {
		flg := true
		acc := v["subaccountId"].(string)
		for _, val := range marketList {
			_, err := ChangeFee(acc, val)
			if err != nil {
				fmt.Println("ChangeFee , ", err)
				// flg = false
			}
			fmt.Println(acc, " : ", val)
			time.Sleep(time.Millisecond * 100)
		}

		fmt.Println("flg : , ", flg)
		SaveDBUpdate(acc)

		time.Sleep(time.Second)
	}
}

func GetMarketList() []string {

	res, err := binapi.NewExchangeInfoService().Do(context.Background())
	if err != nil {
		fmt.Println("err : ", err)
	}

	var send []string
	for _, v := range res.Symbols {
		if strings.HasSuffix(v.Symbol, "USDT") {
			if v.Status == "TRADING" {
				send = append(send, v.Symbol)
			}
		}
	}

	return send
}

func ChangeFee(acc, symbol string) (*binance.BrokerChangeUSDTfuturesCommissionREQ, error) {
	ret, er := binapi.NewBrokerChangeUSDTfuturesCommission().SubAccountID(acc).Symbol(symbol).MakerAdjustment(0).TakerAdjustment(0).Do(context.Background())
	return ret, er
}

func SaveDBUpdate(acc string) {
	sql := fmt.Sprintf("update accountMasterReCycle set feeFlg=1 where subaccountId ='%s'", acc)
	_, err := adminconfig.GetMainMysqlDB().DBQueryExec(sql)
	if err != nil {
		fmt.Println("err : ", err)
	}
}
