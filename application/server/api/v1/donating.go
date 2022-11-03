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

type DonatingRequestBody struct {
	ObjectOfDonating string `json:"objectOfDonating"` //Donation
	Donor            string `json:"donor"`            //Donors
	Grantee          string `json:"grantee"`          //Recipient
}

type DonatingListQueryRequestBody struct {
	Donor string `json:"donor"`
}

type DonatingListQueryByGranteeRequestBody struct {
	Grantee string `json:"grantee"`
}

type UpdateDonatingRequestBody struct {
	ObjectOfDonating string `json:"objectOfDonating"` //Donation
	Donor            string `json:"donor"`            //Donors
	Grantee          string `json:"grantee"`          //Recipient
	Status           string `json:"status"`           //The state that needs to be changed
}

func CreateDonating(c *gin.Context) {
	appG := app.Gin{C: c}
	body := new(DonatingRequestBody)
	//Analyze the body parameter
	if err := c.ShouldBind(body); err != nil {
		appG.Response(http.StatusBadRequest, "fail", fmt.Sprintf("Parameter error%s", err.Error()))
		return
	}
	if body.ObjectOfDonating == "" || body.Donor == "" || body.Grantee == "" {
		appG.Response(http.StatusBadRequest, "fail", "Objectofdonating donation objects and Donor donors and Grantee gifts cannot be empty")
		return
	}
	var bodyBytes [][]byte
	bodyBytes = append(bodyBytes, []byte(body.ObjectOfDonating))
	bodyBytes = append(bodyBytes, []byte(body.Donor))
	bodyBytes = append(bodyBytes, []byte(body.Grantee))
	//Call smart contract
	resp, err := bc.ChannelExecute("createDonating", bodyBytes)
	if err != nil {
		appG.Response(http.StatusInternalServerError, "fail", err.Error())
		return
	}
	var data map[string]interface{}
	if err = json.Unmarshal(bytes.NewBuffer(resp.Payload).Bytes(), &data); err != nil {
		appG.Response(http.StatusInternalServerError, "fail", err.Error())
		return
	}
	appG.Response(http.StatusOK, "success", data)
}

func QueryDonatingList(c *gin.Context) {
	appG := app.Gin{C: c}
	body := new(DonatingListQueryRequestBody)
	//Analyze the body parameter
	if err := c.ShouldBind(body); err != nil {
		appG.Response(http.StatusBadRequest, "fail", fmt.Sprintf("Parameter error%s", err.Error()))
		return
	}
	var bodyBytes [][]byte
	if body.Donor != "" {
		bodyBytes = append(bodyBytes, []byte(body.Donor))
	}
	//Call smart contract
	resp, err := bc.ChannelQuery("queryDonatingList", bodyBytes)
	if err != nil {
		appG.Response(http.StatusInternalServerError, "fail", err.Error())
		return
	}
	// Reverse serialization JSON
	var data []map[string]interface{}
	if err = json.Unmarshal(bytes.NewBuffer(resp.Payload).Bytes(), &data); err != nil {
		appG.Response(http.StatusInternalServerError, "fail", err.Error())
		return
	}
	appG.Response(http.StatusOK, "success", data)
}

func QueryDonatingListByGrantee(c *gin.Context) {
	appG := app.Gin{C: c}
	body := new(DonatingListQueryByGranteeRequestBody)
	//Analyze the body parameter
	if err := c.ShouldBind(body); err != nil {
		appG.Response(http.StatusBadRequest, "fail", fmt.Sprintf("Parameter error%s", err.Error()))
		return
	}
	if body.Grantee == "" {
		appG.Response(http.StatusBadRequest, "fail", "You must specify an accountid query")
		return
	}
	var bodyBytes [][]byte
	bodyBytes = append(bodyBytes, []byte(body.Grantee))
	//Call smart contract
	resp, err := bc.ChannelQuery("queryDonatingListByGrantee", bodyBytes)
	if err != nil {
		appG.Response(http.StatusInternalServerError, "fail", err.Error())
		return
	}
	// Reverse serialization JSON
	var data []map[string]interface{}
	if err = json.Unmarshal(bytes.NewBuffer(resp.Payload).Bytes(), &data); err != nil {
		appG.Response(http.StatusInternalServerError, "fail", err.Error())
		return
	}
	appG.Response(http.StatusOK, "success", data)
}

func UpdateDonating(c *gin.Context) {
	appG := app.Gin{C: c}
	body := new(UpdateDonatingRequestBody)
	//Analyze the body parameter
	if err := c.ShouldBind(body); err != nil {
		appG.Response(http.StatusBadRequest, "fail", fmt.Sprintf("Parameter error%s", err.Error()))
		return
	}
	if body.ObjectOfDonating == "" || body.Donor == "" || body.Grantee == "" || body.Status == "" {
		appG.Response(http.StatusBadRequest, "fail", "Parameters cannot be empty")
		return
	}
	var bodyBytes [][]byte
	bodyBytes = append(bodyBytes, []byte(body.ObjectOfDonating))
	bodyBytes = append(bodyBytes, []byte(body.Donor))
	bodyBytes = append(bodyBytes, []byte(body.Grantee))
	bodyBytes = append(bodyBytes, []byte(body.Status))
	//Call smart contract
	resp, err := bc.ChannelExecute("updateDonating", bodyBytes)
	if err != nil {
		appG.Response(http.StatusInternalServerError, "fail", err.Error())
		return
	}
	var data map[string]interface{}
	if err = json.Unmarshal(bytes.NewBuffer(resp.Payload).Bytes(), &data); err != nil {
		appG.Response(http.StatusInternalServerError, "fail", err.Error())
		return
	}
	appG.Response(http.StatusOK, "success", data)
}
