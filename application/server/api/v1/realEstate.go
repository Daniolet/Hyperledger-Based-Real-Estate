package v1

import (
	bc "application/blockchain"
	"application/pkg/app"
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type RealEstateRequestBody struct {
	AccountId   string  `json:"accountId"`   //Operator ID
	Proprietor  string  `json:"proprietor"`  //Owner (owner) (owner accountidid)
	TotalArea   float64 `json:"totalArea"`   //The total area
	LivingSpace float64 `json:"livingSpace"` //Living space
}

type RealEstateQueryRequestBody struct {
	Proprietor string `json:"proprietor"` //Owner (owner) (owner accountidid)
}

func CreateRealEstate(c *gin.Context) {
	appG := app.Gin{C: c}
	body := new(RealEstateRequestBody)
	//Analyze the body parameter
	if err := c.ShouldBind(body); err != nil {
		fmt.Println(body)
		appG.Response(http.StatusBadRequest, "fail", fmt.Sprintf("parameterError%s", err.Error()))
		return
	}
	NrawBody, _ := c.GetRawData()
	s := string(NrawBody)
	newdata := RealEstateRequestBody{}
	json.Unmarshal([]byte(s), &newdata)

	fmt.Println("second", (newdata))

	if newdata.TotalArea <= 0 || newdata.LivingSpace <= 0 || newdata.LivingSpace > newdata.TotalArea {
		fmt.Print(newdata)
		appG.Response(http.StatusBadRequest, "fail", "The total area of Total Area and Living Space must be greater than 0, and the living space is less than equal to the total"+string(newdata.AccountId))
		return
	}
	var bodyBytes [][]byte
	bodyBytes = append(bodyBytes, []byte(newdata.AccountId))
	bodyBytes = append(bodyBytes, []byte(newdata.Proprietor))
	bodyBytes = append(bodyBytes, []byte(strconv.FormatFloat(newdata.TotalArea, 'E', -1, 64)))
	bodyBytes = append(bodyBytes, []byte(strconv.FormatFloat(newdata.LivingSpace, 'E', -1, 64)))
	//Call smart contract
	resp, err := bc.ChannelExecute("createRealEstate", bodyBytes)
	if err != nil {
		fmt.Print("are we here")
		appG.Response(http.StatusInternalServerError, "fail", err.Error())
		return
	}
	var data map[string]interface{}
	if err = json.Unmarshal(bytes.NewBuffer(resp.Payload).Bytes(), &data); err != nil {
		fmt.Print("are we here 2")
		appG.Response(http.StatusInternalServerError, "fail", err.Error())
		return
	}
	appG.Response(http.StatusOK, "success", data)
}

func QueryRealEstateList(c *gin.Context) {
	appG := app.Gin{C: c}
	body := new(RealEstateQueryRequestBody)
	//Analyze the body parameter
	if err := c.ShouldBind(body); err != nil {
		appG.Response(http.StatusBadRequest, "fail", fmt.Sprintf("Parameter error%s", err.Error()))
		return
	}
	var bodyBytes [][]byte
	if body.Proprietor != "" {
		bodyBytes = append(bodyBytes, []byte(body.Proprietor))
	}
	//Call smart contract
	resp, err := bc.ChannelQuery("queryRealEstateList", bodyBytes)
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
