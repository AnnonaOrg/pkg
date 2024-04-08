package cryptapi

import (
	"testing"

	"github.com/AnnonaOrg/pkg/cryptapi/utils"
)

const callBackUrl = "https://webhook.site/b02cd43e-3c21-498a-9541-331fc5e65fe2"

// logs error or success
func checkErr(t *testing.T, err error, res any) {
	if err != nil {
		t.Log(err)
		t.Error("fail")
	} else {
		t.Log("success")
	}
}

// test get coins util
func TestGetCoins(t *testing.T) {
	res, err := utils.GetCoins()
	checkErr(t, err, res)
}

// test gen address method
func TestGenPaymentAdress(t *testing.T) {
	var ci = InitCryptWrapper("btc", "bc1q8ahrlrnm82f5xvljkdc36ufr934ycse5nrqjq8", callBackUrl, map[string]string{
		"multi_chain": "1",
	}, map[string]string{
		"user_id":  "10",
		"order_id": "12345",
	})
	res, err := ci.GenPaymentAdress()
	checkErr(t, err, res)
}

// test get logs
func TestCheckLogs(t *testing.T) {
	var ci = InitCryptWrapper("btc", "bc1q8ahrlrnm82f5xvljkdc36ufr934ycse5nrqjq8", callBackUrl, map[string]string{
		"multi_chain": "1",
	}, map[string]string{
		"user_id":  "10",
		"order_id": "12345",
	})
	res, err := ci.CheckLogs()
	checkErr(t, err, res)
}

// // test getting QR code
func TestGenQR(t *testing.T) {
	var ci = InitCryptWrapper("btc", "bc1q8ahrlrnm82f5xvljkdc36ufr934ycse5nrqjq8", callBackUrl, map[string]string{
		"convert":     "1",
		"multi_chain": "1",
	}, map[string]string{})
	ci.GenPaymentAdress()
	res, err := ci.GenQR("1", "300")
	checkErr(t, err, res)
}

// test getting estimate
func TestEstTransactionFee(t *testing.T) {
	res, err := EstTransactionFee("polygon_matic", "1", "default")
	checkErr(t, err, res)
}

// test convert
func TestConvert(t *testing.T) {
	res, err := Convert("btc", "1", "USD")
	checkErr(t, err, res)
}
