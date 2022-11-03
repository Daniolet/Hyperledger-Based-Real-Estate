package v1

import (
	bc "application/blockchain"
	"application/pkg/app"
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AccountIdBody struct {
	AccountId string `json:"accountId"`
}

type AccountRequestBody struct {
	Args []AccountIdBody `json:"args"`
}

func QueryAccountList(c *gin.Context) {
	appG := app.Gin{C: c}
	body := new(AccountRequestBody)
	//Analyze the body parameter
	if err := c.ShouldBind(body); err != nil {
		print(err.Error(), "1")
		appG.Response(http.StatusBadRequest, "fail", fmt.Sprintf("Parameter error%s", err.Error()))
		return
	}
	var bodyBytes [][]byte
	for _, val := range body.Args {
		bodyBytes = append(bodyBytes, []byte(val.AccountId))
	}
	//Call smart contract
	resp, err := bc.ChannelQuery("queryAccountList", bodyBytes)
	if err != nil {
		print(err.Error(), "-2")
		appG.Response(http.StatusInternalServerError, "fail", err.Error())
		return
	}
	// Reverse serialization JSON
	var data []map[string]interface{}
	if err = json.Unmarshal(bytes.NewBuffer(resp.Payload).Bytes(), &data); err != nil {
		print(err.Error(), "3")
		appG.Response(http.StatusInternalServerError, "fail", err.Error())
		return
	}
	appG.Response(http.StatusOK, "success", data)
}
