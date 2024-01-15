package send2server

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

// var msg feedmsg.FeedRichMsgModel
// var msgType, msgID, msgTime, textContent, textContentEx, textContentExPic, imagePicURL, imageFilePath string
// msgType = "text"
// textContent = "Hello,world223"
// msg = feedmsg.NewFeedRichMsg(msgType, msgID, msgTime, textContent, textContentEx, textContentExPic, imagePicURL, imageFilePath)
// sendMsgToServer(msg, "", "","","") //http://127.0.0.1:8080/ws/push/xxx
func SendMsgToServer(msgStruct interface{}, serverRouter, serverChannel, serverToken, serverPath string) (ret string, err error) {
	// senderTokenUuid:=
	jsonParams, err := json.Marshal(msgStruct)
	if err != nil {
		return
	}
	jsonString := string(jsonParams)
	// fmt.Println(jsonString)
	payload := strings.NewReader(jsonString)
	if len(serverRouter) == 0 {
		serverRouter = "http://127.0.0.1:8080"
	}
	if len(serverChannel) == 0 {
		serverChannel = "xxx"
	}
	if len(serverPath) == 0 {
		serverPath = "/ws/push/"
	}
	apiUrl := fmt.Sprintf("%s%s%s", serverRouter, serverPath, serverChannel)
	req, err := http.NewRequest("POST", apiUrl, payload)
	if err != nil {
		return
	}
	req.Header.Add("content-type", "application/json")
	req.Header.Add("cache-control", "no-cache")
	req.Header.Add("apiclient", serverToken) //接收 c.Request.Header["Apiclient"]
	req.Header.Add("User-Agent", "apisender")
	res, err := http.DefaultClient.Do(req)
	if res == nil || err != nil {
		return
	}
	defer func() {
		if res != nil && res.Body != nil {
			_ = res.Body.Close()
		}
	}()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return
	}
	// fmt.Println(res)
	// fmt.Println(string(body))
	ret = string(body)
	return
}
