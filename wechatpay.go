package wechatpay

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"

	"crypto/tls"
	log "github.com/sdbaiguanghe/glog"
	"math/rand"
	"net/http"
	"sort"
	"strings"
	"time"
)

// WechatPay 微信支付
type WechatPay struct {
	AppID         string
	MchID         string
	APIKey        string
	ApiclientCert []byte
	ApiclientKey  []byte
}

// New 实例化一个支付客户端
func New(appID, mchID, apiKey string, apiclientCert, apiclientKey []byte) (client *WechatPay) {
	client = &WechatPay{}
	client.AppID = appID
	client.MchID = mchID
	client.APIKey = apiKey
	client.ApiclientCert = apiclientCert
	client.ApiclientKey = apiclientKey
	return client
}

// GetSign wxpay计算签名的函数
func GetSign(mReq map[string]interface{}, key string) (sign string) {

	sortedKeys := make([]string, 0)
	for k := range mReq {
		sortedKeys = append(sortedKeys, k)
	}
	sort.Strings(sortedKeys)
	var signStrings string
	for _, k := range sortedKeys {
		value := fmt.Sprintf("%v", mReq[k])
		if value != "" {
			signStrings = signStrings + k + "=" + value + "&"
		}
	}
	if key != "" {
		signStrings = signStrings + "key=" + key
	}
	md5Ctx := md5.New()
	md5Ctx.Write([]byte(signStrings))
	cipherStr := md5Ctx.Sum(nil)
	upperSign := strings.ToUpper(hex.EncodeToString(cipherStr))
	return upperSign
}

// VerifySign 微信支付签名验证函数
func (w *WechatPay) VerifySign(needVerifyM map[string]interface{}, sign string) bool {
	delete(needVerifyM, "sign")
	signCalc := GetSign(needVerifyM, w.APIKey)
	if sign == signCalc {
		log.Info("wechat verify success!")
		return true
	}
	log.Info("wechat vertify failed!")
	return false
}

// WithCertBytes 证书
func WithCertBytes(cert, key []byte) *http.Transport {
	tlsCert, err := tls.X509KeyPair(cert, key)
	if err != nil {
		log.Error("----通信证书错误---", err.Error())
		return nil
	}
	conf := &tls.Config{
		Certificates:       []tls.Certificate{tlsCert},
		InsecureSkipVerify: true,
	}
	trans := &http.Transport{
		TLSClientConfig: conf,
	}
	return trans
}

func randomNonceStr() string {
	str := "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	bytes := []byte(str)
	result := []byte{}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < 32; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	return string(result)
}
