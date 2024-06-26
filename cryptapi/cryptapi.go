/*
 * go lang wrapper over cryptapi endpoints
 * @author - Joseph Folayan
 * @email - folayanjoey@gmail.com
 * @github - github.com\joey1123455
 */
package cryptapi

import (
	"errors"
	"fmt"
	"net/url"

	"github.com/AnnonaOrg/pkg/cryptapi/utils"
)

/*
 * Crypt - an instance representing the connection to the crypt api
 * @Coin - the currency to transact in
 * @OwnAddress - the wallet to recieve payment
 * @CallBack - url to send the request status
 * @Params - url querry parrams
 * @CaParams - querry params
 * @PaymentAddrs - the wallet to send payment to
 */
type Crypt struct {
	Coin         string
	OwnAddress   string
	CallBack     string
	Params       map[string]string
	CaParams     map[string]string
	PaymentAddrs string
}

/*
 * InitCryptWrapper - creates the crypt request instance
 * @coin - the currency to transact in
 * @ownAddress - the wallet to recieve payment
 * @callBack - url to send the request status
 * @paymentAddrs - the wallet to send payment to
 * @params - url querry parrams
 * @caParams - querry params
 * returns - ptr to crypt instance
 */
func InitCryptWrapper(coin, ownAddress, callBack string, params, caParams map[string]string) *Crypt {
	// TODO: check if coin is valid
	return &Crypt{
		Coin:         coin,
		OwnAddress:   ownAddress,
		CallBack:     callBack,
		PaymentAddrs: "",
		Params:       params,
		CaParams:     caParams,
	}
}

/*
 * EstTransactionFee - returns the estimated value for a crypto transaction
 * @coin - the coin being transacted in
 * @adress - no of address being credited
 * @priority - credit priority settings
 * returns - response data or error
 */
func EstTransactionFee(coin string, address string, priority string) (map[string]any, error) {
	res, err := utils.Request(coin, "estimate", map[string]string{
		"address":  address,
		"priority": priority,
	})
	if err != nil {
		return nil, err
	}
	if res["status"] == "success" {
		return res, nil
	}
	return nil, errors.New(res["error"].(string))
}

/*
 * Convert - This method allows you to easily convert prices from FIAT to Crypto or even between cryptocurrencies
 * @coin - currency to convert to
 * @value -  the amount to convert
 * @from - currency to convert from
 * returns - json response and error
 */
func Convert(coin string, value string, from string) (map[string]any, error) {
	param := map[string]string{
		"value": value,
		"from":  from,
	}
	res, err := utils.Request(coin, "convert", param)
	if err != nil {
		return nil, err
	}
	fmt.Println(res["error"])
	if res["status"] == "success" {
		return res, nil
	}
	return nil, errors.New(res["error"].(string))
}

/*
 * CryptWrapper - an interface defining the crypt api library
 * @GenPaymentAdress - returns a payment wallet address
 * @CheckLogs - checks payment logs for requets
 * @GenQR - generates a qr code for payment
 */
type CryptWrapper interface {
}

/*
 * GenPaymentAdress - creates the address for customer to pay too
 * @w - crypt instance (reciever method)
 * returns - payment wallet or error
 */
func (w *Crypt) GenPaymentAdress() (string, error) {
	if w.Coin == "" || w.CallBack == "" || w.OwnAddress == "" {
		return "", errors.New("incomplte information coin, callback url and own address must be specified")
	}
	callBackUrl, err := url.Parse(w.CallBack)
	query := callBackUrl.Query()
	if err != nil {
		return "", err
	}
	for key, val := range w.CaParams {
		query.Add(key, val)
	}
	params := make(map[string]string)
	for key, val := range w.Params {
		params[key] = val
	}
	callBackUrl.RawQuery = query.Encode()
	callBack := callBackUrl.String()
	params["callback"] = callBack
	params["address"] = w.OwnAddress

	res, err := utils.Request(w.Coin, "create", params)
	if err != nil {
		return "", err
	}
	if res["status"] == "success" {
		add := res["address_in"].(string)
		w.PaymentAddrs = add
		return add, nil
	}
	return "", errors.New(res["error"].(string))
}

/*
 * CheckLogs - provides logs for transactions sent to a payment wallet
 * @w - ptr to crypt instance (reciever method)
 * returns - logs or error
 */
func (w *Crypt) CheckLogs() (map[string]any, error) {
	if w.Coin == "" || w.CallBack == "" {
		return nil, errors.New("incomplete data")
	}
	callBackUrl, err := url.Parse(w.CallBack)
	query := callBackUrl.Query()
	if err != nil {
		return nil, err
	}
	for key, val := range w.CaParams {
		query.Add(key, val)
	}

	callBackUrl.RawQuery = query.Encode()
	callBack := callBackUrl.String()
	res, err := utils.Request(w.Coin, "logs", map[string]string{
		"callback": callBack,
	})
	fmt.Println(res)
	if err != nil {
		return nil, err
	}
	if res["status"] == "success" {
		return res, nil
	}
	return nil, errors.New(res["error"].(string))
}

/*
 * GenQR - generates the qr code for the transaction
 * @value - qr value
 * @w - crypt instance
 * @size - qr code size
 * returns qr code or error
 */
func (w *Crypt) GenQR(value string, size string) (map[string]any, error) {
	if size == "0" {
		size = "512"
	}
	params := map[string]string{
		"address": w.PaymentAddrs,
	}
	if value != "" {
		params["value"] = value
	}
	params["size"] = size
	res, err := utils.Request(w.Coin, "qrcode", params)
	if err != nil {
		return nil, err
	}
	if res["status"] == "success" {
		return res, nil
	}
	return nil, errors.New(res["error"].(string))
}
