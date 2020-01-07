package wechatpay

import (
	"bytes"
	"crypto/tls"
	"encoding/xml"
	log "github.com/sdbaiguanghe/glog"
	"io/ioutil"
	"net/http"
	"strings"
)

// Refund 退款
func (w *WechatPay) Refund(param OrderRefund) (*OrderRefundResult, error) {

	if param.AppID == "" {
		param.AppID = w.AppID
	}
	param.MchID = w.MchID
	param.NonceStr = randomNonceStr()

	var m map[string]interface{}
	m = make(map[string]interface{}, 0)
	m["appid"] = param.AppID
	m["mch_id"] = param.MchID
	m["total_fee"] = param.TotalFee
	m["out_trade_no"] = param.OutTradeNo
	m["nonce_str"] = param.NonceStr
	m["refund_fee"] = param.RefundFee
	m["out_refund_no"] = param.OutRefundNo
	param.Sign = GetSign(m, w.APIKey)

	bytesReq, err := xml.Marshal(param)
	if err != nil {
		log.Error(err, "xml marshal failed")
		return nil, err
	}

	strReq := string(bytesReq)
	strReq = strings.Replace(strReq, "Refund", "xml", -1)
	bytesReq = []byte(strReq)

	//发送unified order请求.
	req, err := http.NewRequest("POST", REFUNDURL, bytes.NewReader(bytesReq))
	if err != nil {
		log.Error(err, "new http request failed,err :"+err.Error())
		return nil, err
	}
	req.Header.Set("Accept", "application/xml")
	req.Header.Set("Content-Type", "application/xml;charset=utf-8")

	wReq := http.Client{
		Transport: WithCertBytes(w.ApiclientCert, w.ApiclientKey),
	}

	resp, _err := wReq.Do(req)
	if _err != nil {
		log.Error(err, "http request failed! err :"+_err.Error())
		return nil, _err
	}
	body, _ := ioutil.ReadAll(resp.Body)

	var refundResp OrderRefundResult

	_err = xml.Unmarshal(body, &refundResp)
	if _err != nil {
		log.Error(err, "http request failed! err :"+_err.Error())
		return nil, _err
	}
	return &refundResp, nil
}

// RefundQuery 退款查询
func (w *WechatPay) RefundQuery(refundStatus OrderRefundQuery) (*OrderRefundQueryResult, error) {

	refundStatus.AppID = w.AppID
	refundStatus.MchID = w.MchID
	refundStatus.NonceStr = randomNonceStr()

	var m map[string]interface{}
	m = make(map[string]interface{}, 0)
	m["appid"] = refundStatus.AppID
	m["mch_id"] = refundStatus.MchID
	m["out_trade_no"] = refundStatus.OutTradeNo
	m["nonce_str"] = refundStatus.NonceStr
	refundStatus.Sign = GetSign(m, w.APIKey)

	bytesReq, err := xml.Marshal(refundStatus)
	if err != nil {
		log.Error(err, "xml marshal failed,err:"+err.Error())
		return nil, err
	}

	strReq := string(bytesReq)
	strReq = strings.Replace(strReq, "RefundQuery", "xml", -1)
	bytesReq = []byte(strReq)

	req, err := http.NewRequest("POST", REFUNDQUERYURL, bytes.NewReader(bytesReq))
	if err != nil {
		log.Error(err, "new http request failed,err :"+err.Error())
		return nil, err
	}
	req.Header.Set("Accept", "application/xml")
	req.Header.Set("Content-Type", "application/xml;charset=utf-8")

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	wReq := http.Client{Transport: tr}
	resp, _err := wReq.Do(req)
	if _err != nil {
		log.Error(err, "http request failed! err :"+_err.Error())
		return nil, err
	}
	var refundResp OrderRefundQueryResult
	body, _ := ioutil.ReadAll(resp.Body)

	_err = xml.Unmarshal(body, &refundResp)
	if _err != nil {
		log.Error(err, "http request failed! err :"+_err.Error())
		return nil, err
	}
	return &refundResp, nil
}
