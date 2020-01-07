package wechatpay

const (
	UNIT_ORDER_URL   = "https://api.mch.weixin.qq.com/pay/unifiedorder"
	REFUND_URL       = "https://api.mch.weixin.qq.com/secapi/pay/refund"
	REFUND_QUERY_URL = "https://api.mch.weixin.qq.com/pay/refundquery"
	TRANSFERS_URL    = "https://api.mch.weixin.qq.com/mmpaymkttransfers/promotion/transfers"
)

type Base struct {
	AppId    string `xml:"appid"`
	MchId    string `xml:"mch_id"`
	NonceStr string `xml:"nonce_str"`
	Sign     string `xml:"sign"`
}

//统一下单请求参数
type UnitOrder struct {
	Base
	Body           string `xml:"body"`
	NotifyUrl      string `xml:"notify_url"`
	TradeType      string `xml:"trade_type"`
	SpbillCreateIp string `xml:"spbill_create_ip"`
	TotalFee       int    `xml:"total_fee"`
	OutTradeNo     string `xml:"out_trade_no"`
	SceneInfo      string `xml:"scene_info"`
	Openid         string `xml:"openid"`
	TimeStart      string `xml:"time_start"`
	TimeExpire     string `xml:"time_expire"`
	Attach         string `xml:"attach"`
	Referer        string
}

//统一下单返回参数
type UnifyOrderResult struct {
	Base
	ReturnCode string `xml:"return_code"`
	ReturnMsg  string `xml:"return_msg"`
	Attach     string `xml:"attach"`
	ResultCode string `xml:"result_code"`
	PrepayId   string `xml:"prepay_id"`
	TradeType  string `xml:"trade_type"`
	CodeUrl    string `xml:"code_url"`
	MwebUrl    string `xml:"mweb_url"`
}

//订单查询
type OrderQuery struct {
	Base
	Transaction_id string `xml:"transaction_id"`
}

type OrderQueryResult struct {
	Base
	ReturnCode     string `xml:"return_code"`
	ReturnMsg      string `xml:"return_msg"`
	ResultCode     string `xml:"result_code"`
	OpenId         string `xml:"prepay_id"`
	TradeType      string `xml:"trade_type"`
	TradeState     string `xml:"trade_state"`
	BankType       string `xml:"bank_type"`
	TotalTee       string `xml:"total_fee"`
	CashFee        int    `xml:"cash_fee"`
	TransactionId  string `xml:"transaction_id"`
	OutTradeNo     string `xml:"out_trade_no"`
	TimeEnd        string `xml:"time_end"`
	TradeStateDesc string `xml:"trade_state_desc"`
}

//下单回调
type PayNotifyResult struct {
	Base
	ReturnCode    string `xml:"return_code"`
	ReturnMsg     string `xml:"return_msg"`
	ResultCode    string `xml:"result_code"`
	OpenId        string `xml:"openid"`
	IsSubscribe   string `xml:"is_subscribe"`
	TradeType     string `xml:"trade_type"`
	BankType      string `xml:"bank_type"`
	TotalFee      int    `xml:"total_fee"`
	FeeType       string `xml:"fee_type"`
	CashFee       int    `xml:"cash_fee"`
	CashFeeType   string `xml:"cash_fee_type"`
	TransactionId string `xml:"transaction_id"`
	OutTradeNo    string `xml:"out_trade_no"`
	Attach        string `xml:"attach"`
	TimeEnd       string `xml:"time_end"`
}

//下单回调返回值
type PayNotifyResp struct {
	ReturnCode string `xml:"return_code"`
	ReturnMsg  string `xml:"return_msg"`
}

//订单退款
type OrderRefund struct {
	Base
	TotalFee    int    `xml:"total_fee"`
	OutTradeNo  string `xml:"out_trade_no"`
	OutRefundNo string `xml:"out_refund_no"`
	RefundFee   int    `xml:"refund_fee"`
}

//订单退款结果
type OrderRefundResult struct {
	Base
	ReturnCode    string `xml:"return_code"`
	ReturnMsg     string `xml:"return_msg"`
	ResultCode    string `xml:"result_code"`
	TransactionId string `xml:"transaction_id"`
	OutRefundNo   string `xml:"out_refund_no"`
	OutTradeNo    string `xml:"out_trade_no"`
	RefundFee     int    `xml:"refund_fee"`
	TotalFee      int    `xml:"total_fee"`
	CashFee       int    `xml:"cash_fee"`
}

//退款查询
type OrderRefundQuery struct {
	Base
	OutTradeNo string `xml:"out_trade_no"`
}

//退款结果查询
type OrderRefundQueryResult struct {
	Base
	ReturnCode            string `xml:"return_code"`
	ReturnMsg             string `xml:"return_msg"`
	ResultCode            string `xml:"result_code"`
	OutTradeNo            string `xml:"out_trade_no"`
	RefundStatus_0        string `xml:"refund_status_0"`
	SettlementRefundFee_0 string `xml:"settlement_refund_fee_0"`
	ErrCodeDes            string `xml:"err_code_des"`
}

// EnterpriseTransfers 企业付款参数
type EnterpriseTransfers struct {
	MchAppID       string `xml:"mch_appid"`
	MchID          string `xml:"mchid"`
	DeviceInfo     string `xml:"device_info"`
	NonceStr       string `xml:"nonce_str"`
	Sign           string `xml:"sign"`
	PartnerTradeNo string `xml:"partner_trade_no"`
	Openid         string `xml:"openid"`
	CheckName      string `xml:"check_name"`
	ReUserName     string `xml:"re_user_name"`
	Amount         int    `xml:"amount"`
	Desc           string `xml:"desc"`
	SpBillCreateIP string `xml:"spbill_create_ip"`
}

// EnterpriseTransfersResult 企业付款返回结果
type EnterpriseTransfersResult struct {
	ReturnCode     string `xml:"return_code"`
	ReturnMsg      string `xml:"return_msg"`
	MchAppID       string `xml:"mch_appid"`
	MchID          string `xml:"mchid"`
	DeviceInfo     string `xml:"device_info"`
	NonceStr       string `xml:"nonce_str"`
	ResultCode     string `xml:"result_code"`
	ErrCode        string `xml:"err_code"`
	ErrCodeDes     string `xml:"err_code_des"`
	PartnerTradeNo string `xml:"partner_trade_no"`
	PaymentNo      string `xml:"payment_no"`
	PaymentTime    string `xml:"payment_time"`
}
