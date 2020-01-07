package wechatpay

import (
	"bytes"
	"crypto/tls"
	"encoding/xml"
	"io/ioutil"
	"net/http"
	"strings"
)

// Pay 统一下单
func (w *WechatPay) Pay(param UnitOrder) (*UnifyOrderResult, error) {
	if param.AppID == "" {
		param.AppID = w.AppID
	}
	param.MchID = w.MchID
	param.NonceStr = randomNonceStr()

	var m map[string]interface{}
	m = make(map[string]interface{}, 0)
	m["appid"] = param.AppID
	m["body"] = param.Body
	m["mch_id"] = param.MchID
	m["notify_url"] = param.NotifyURL
	m["trade_type"] = param.TradeType
	m["spbill_create_ip"] = param.SpbillCreateIP
	m["total_fee"] = param.TotalFee
	m["out_trade_no"] = param.OutTradeNo
	m["nonce_str"] = param.NonceStr
	if param.TradeType == "MWEB" {
		m["scene_info"] = param.SceneInfo
	}
	if param.TradeType == "JSAPI" {
		m["openid"] = param.Openid
	}
	if param.TimeExpire != "" {
		m["time_expire"] = param.TimeExpire
	}
	if param.TimeStart != "" {
		m["time_start"] = param.TimeStart
	}
	if param.Attach != "" {
		m["attach"] = param.Attach
	}
	param.Sign = GetSign(m, w.APIKey)

	bytesReq, err := xml.Marshal(param)
	if err != nil {
		return nil, err
	}
	strReq := string(bytesReq)
	strReq = strings.Replace(strReq, "UnitOrder", "xml", -1)
	req, err := http.NewRequest("POST", UNITORDERURL, bytes.NewReader([]byte(strReq)))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Accept", "application/xml")
	req.Header.Set("Content-Type", "application/xml;charset=utf-8")
	if param.TradeType == "MWEB" {
		req.Header.Set("Referer", param.Referer)
	}
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	wReq := http.Client{Transport: tr}
	resp, err := wReq.Do(req)
	if err != nil {
		return nil, err
	}
	body, _ := ioutil.ReadAll(resp.Body)

	var payResult UnifyOrderResult
	err = xml.Unmarshal(body, &payResult)
	if err != nil {
		return nil, err
	}
	return &payResult, nil
}

// Transfers 企业付款
func (w *WechatPay) Transfers(p *EnterpriseTransfers) (*EnterpriseTransfersResult, error) {
	p.MchID = w.MchID
	if p.MchAppID == "" {
		p.MchAppID = w.AppID
	}
	p.NonceStr = randomNonceStr()
	var m = map[string]interface{}{
		"mch_appid":        p.MchAppID,
		"mchid":            p.MchID,
		"device_info":      p.DeviceInfo,
		"nonce_str":        p.NonceStr,
		"sign":             p.Sign,
		"partner_trade_no": p.PartnerTradeNo,
		"openid":           p.Openid,
		"check_name":       p.CheckName,
		"re_user_name":     p.ReUserName,
		"amount":           p.Amount,
		"desc":             p.Desc,
		"spbill_create_ip": p.SpBillCreateIP,
	}
	p.Sign = GetSign(m, w.APIKey)
	bytesReq, err := xml.Marshal(p)
	if err != nil {
		return nil, err
	}
	strReq := string(bytesReq)
	strReq = strings.Replace(strReq, "EnterpriseTransfers", "xml", -1)
	req, err := http.NewRequest("POST", TRANSFERSURL, bytes.NewReader([]byte(strReq)))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Accept", "application/xml")
	req.Header.Set("Content-Type", "application/xml;charset=utf-8")
	//tr := &http.Transport{
	//	TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	//}
	wReq := http.Client{Transport: WithCertBytes(w.ApiclientCert, w.ApiclientKey)}
	resp, err := wReq.Do(req)
	if err != nil {
		return nil, err
	}
	body, _ := ioutil.ReadAll(resp.Body)

	var payResult EnterpriseTransfersResult
	err = xml.Unmarshal(body, &payResult)
	if err != nil {
		return nil, err
	}
	return &payResult, nil
}
