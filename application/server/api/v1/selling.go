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

type SellingRequestBody struct {
	ObjectOfSale string  `json:"objectOfSale"` //Sales target (Realestateid, real estate is being sold)
	Seller       string  `json:"seller"`       //Initiating seller and seller (seller accountidid)
	Price        float64 `json:"price"`        //price
	SalePeriod   int     `json:"salePeriod"`   //The validity period of smart contracts (the unit is heaven)
}

type SellingByBuyRequestBody struct {
	ObjectOfSale string `json:"objectOfSale"` //Sales target (Realestateid, real estate is being sold)
	Seller       string `json:"seller"`       //Initiating seller and seller (seller accountidid)
	Buyer        string `json:"buyer"`        //Buyer (buyer accountidid)
}

type SellingListQueryRequestBody struct {
	Seller string `json:"seller"` //Initiating seller and seller (seller accountidid)
}

type SellingListQueryByBuyRequestBody struct {
	Buyer string `json:"buyer"` //Buyer (buyer accountidid)
}

type UpdateSellingRequestBody struct {
	ObjectOfSale string `json:"objectOfSale"` //Sales target (Realestateid, real estate is being sold)
	Seller       string `json:"seller"`       //Initiating seller and seller (seller accountidid)
	Buyer        string `json:"buyer"`        //Buyer (buyer accountidid)
	Status       string `json:"status"`       //The state that needs to be changed
}

func CreateSelling(c *gin.Context) {
	appG := app.Gin{C: c}
	body := new(SellingRequestBody)
	//解析Body参数
	if err := c.ShouldBind(body); err != nil {
		appG.Response(http.StatusBadRequest, "fail", fmt.Sprintf("Parameter error%s", err.Error()))
		return
	}
	NrawBody, _ := c.GetRawData()
	s := string(NrawBody)
	newdata := SellingRequestBody{}
	json.Unmarshal([]byte(s), &newdata)
	fmt.Println(newdata, body)
	if newdata.ObjectOfSale == "" || newdata.Seller == "" {
		appG.Response(http.StatusBadRequest, "fail", "Objectofsale sales target and Seller initiated seller cannot be empty")
		return
	}
	if newdata.Price <= 0 || newdata.SalePeriod <= 0 {
		appG.Response(http.StatusBadRequest, "fail", "The validity period of Price price and saleperiod smart contract(The unit is the sky)Must be greater than 0")
		return
	}
	var bodyBytes [][]byte
	bodyBytes = append(bodyBytes, []byte(newdata.ObjectOfSale))
	bodyBytes = append(bodyBytes, []byte(newdata.Seller))
	bodyBytes = append(bodyBytes, []byte(strconv.FormatFloat(newdata.Price, 'E', -1, 64)))
	bodyBytes = append(bodyBytes, []byte(strconv.Itoa(newdata.SalePeriod)))
	//Call smart contract
	resp, err := bc.ChannelExecute("createSelling", bodyBytes)
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

func CreateSellingByBuy(c *gin.Context) {
	appG := app.Gin{C: c}
	body := new(SellingByBuyRequestBody)
	//Analyze the body parameter
	if err := c.ShouldBind(body); err != nil {
		appG.Response(http.StatusBadRequest, "fail", fmt.Sprintf("Parameter error%s", err.Error()))
		return
	}
	NrawBody, _ := c.GetRawData()
	s := string(NrawBody)
	newdata := SellingByBuyRequestBody{}
	json.Unmarshal([]byte(s), &newdata)
	if newdata.ObjectOfSale == "" || newdata.Seller == "" || newdata.Buyer == "" {
		appG.Response(http.StatusBadRequest, "fail", "Parameters cannot be empty")
		return
	}
	var bodyBytes [][]byte
	bodyBytes = append(bodyBytes, []byte(newdata.ObjectOfSale))
	bodyBytes = append(bodyBytes, []byte(newdata.Seller))
	bodyBytes = append(bodyBytes, []byte(newdata.Buyer))
	//Call smart contract
	resp, err := bc.ChannelExecute("createSellingByBuy", bodyBytes)
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

func QuerySellingList(c *gin.Context) {
	appG := app.Gin{C: c}
	body := new(SellingListQueryRequestBody)
	//Analyze the body parameter
	if err := c.ShouldBind(body); err != nil {
		appG.Response(http.StatusBadRequest, "fail", fmt.Sprintf("Parameter error%s", err.Error()))
		return
	}
	var bodyBytes [][]byte
	if body.Seller != "" {
		bodyBytes = append(bodyBytes, []byte(body.Seller))
	}
	//Call smart contract
	resp, err := bc.ChannelQuery("querySellingList", bodyBytes)
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

func QuerySellingListByBuyer(c *gin.Context) {
	appG := app.Gin{C: c}
	body := new(SellingListQueryByBuyRequestBody)
	//Analyze the body parameter
	if err := c.ShouldBind(body); err != nil {
		appG.Response(http.StatusBadRequest, "fail", fmt.Sprintf("Parameter error%s", err.Error()))
		return
	}
	if body.Buyer == "" {
		appG.Response(http.StatusBadRequest, "fail", "You must specify the buyer Acountid to query")
		return
	}
	var bodyBytes [][]byte
	bodyBytes = append(bodyBytes, []byte(body.Buyer))
	//Call smart contract
	resp, err := bc.ChannelQuery("querySellingListByBuyer", bodyBytes)
	if err != nil {
		appG.Response(http.StatusInternalServerError, "fail", err.Error())
		return
	}
	// 反序列化json
	var data []map[string]interface{}
	if err = json.Unmarshal(bytes.NewBuffer(resp.Payload).Bytes(), &data); err != nil {
		appG.Response(http.StatusInternalServerError, "fail", err.Error())
		return
	}
	appG.Response(http.StatusOK, "success", data)
}

func UpdateSelling(c *gin.Context) {
	appG := app.Gin{C: c}
	body := new(UpdateSellingRequestBody)
	//Analyze the body parameter
	if err := c.ShouldBind(body); err != nil {
		appG.Response(http.StatusBadRequest, "fail", fmt.Sprintf("Parameter error%s", err.Error()))
		return
	}
	NrawBody, _ := c.GetRawData()
	s := string(NrawBody)
	newdata := UpdateSellingRequestBody{}
	json.Unmarshal([]byte(s), &newdata)
	if newdata.ObjectOfSale == "" || newdata.Seller == "" || newdata.Status == "" {
		appG.Response(http.StatusBadRequest, "fail", "Parameters cannot be empty")
		return
	}
	var bodyBytes [][]byte
	bodyBytes = append(bodyBytes, []byte(newdata.ObjectOfSale))
	bodyBytes = append(bodyBytes, []byte(newdata.Seller))
	bodyBytes = append(bodyBytes, []byte(newdata.Buyer))
	bodyBytes = append(bodyBytes, []byte(newdata.Status))
	//Call smart contract
	resp, err := bc.ChannelExecute("updateSelling", bodyBytes)
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
