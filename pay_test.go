package wechatpay

import (
	"fmt"
	"os"
	"testing"
)

var (
	wechatCert   = "111111111111232121321311"
	wechatKey    = "12123222222222223232332323"
	wechatAppID  = "102801212"
	wechatMchID  = "232312123"
	wechatAPIKey = "121212"
)
var wechatClient *WechatPay

func TestMain(m *testing.M) {
	wechatClient = New(wechatAppID, wechatMchID,
		wechatAPIKey, []byte(wechatCert), []byte(wechatKey))
	exitCode := m.Run()
	os.Exit(exitCode)
}

func TestWechat_Pay(t *testing.T) {
	var payData UnitOrder
	payData.NotifyURL = "http://47.98.87.189"
	payData.TradeType = "NATIVE"
	payData.Body = "测试支付"
	payData.SpbillCreateIP = "47.98.87.189"
	payData.TotalFee = 1
	payData.OutTradeNo = "123456789"

	fmt.Println(wechatClient.Pay(payData))
}
