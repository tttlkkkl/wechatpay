package wechatpay

import (
	"fmt"
	"testing"
)

func TestWechat_Refund(t *testing.T) {
	var refundData OrderRefund

	refundData.TotalFee = 1
	refundData.OutTradeNo = "1234567"
	refundData.OutRefundNo = "r122121"
	refundData.RefundFee = 1
	fmt.Println(wechatClient.Refund(refundData))
}
