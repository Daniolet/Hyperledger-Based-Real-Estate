package cron

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"time"

	bc "application/blockchain"
	"application/model"

	"github.com/robfig/cron/v3"
)

const spec = "0 0 0 * * ?" // Execute at 0 o'clock every day
//const spec = "*/10 * * * * ?" //Execute once every 10 seconds for testing

func Init() {
	c := cron.New(cron.WithSeconds()) //Support to second level
	_, err := c.AddFunc(spec, GoRun)
	if err != nil {
		log.Printf("Timing task opens failure %s", err)
	}
	c.Start()
	log.Printf("Timing task has been opened")
	select {}
}

func GoRun() {
	log.Printf("Timing task has been started")
	//Inquiry all sales first
	resp, err := bc.ChannelQuery("querySellingList", [][]byte{}) //Call smart contract
	if err != nil {
		log.Printf("Timing task-querySellinglist failed%s", err.Error())
		return
	}
	// Reverse serialization JSON
	var data []model.Selling
	if err = json.Unmarshal(bytes.NewBuffer(resp.Payload).Bytes(), &data); err != nil {
		log.Printf("Timing task-Reverse serialization JSON failed%s", err.Error())
		return
	}
	for _, v := range data {
		//Filter the status as a sales and delivery
		if v.SellingStatus == model.SellingStatusConstant()["saleStart"] ||
			v.SellingStatus == model.SellingStatusConstant()["delivery"] {
			//Validity period
			day, _ := time.ParseDuration(fmt.Sprintf("%dh", v.SalePeriod*24))
			local, _ := time.LoadLocation("Local")
			t, _ := time.ParseInLocation("2006-01-02 15:04:05", v.CreateTime, local)
			vTime := t.Add(day)
			//If time.now () is greater than vtime explanation explanation
			if time.Now().Local().After(vTime) {
				//Change the status to expired
				var bodyBytes [][]byte
				bodyBytes = append(bodyBytes, []byte(v.ObjectOfSale))
				bodyBytes = append(bodyBytes, []byte(v.Seller))
				bodyBytes = append(bodyBytes, []byte(v.Buyer))
				bodyBytes = append(bodyBytes, []byte("expired"))
				//Call smart contract
				resp, err := bc.ChannelExecute("updateSelling", bodyBytes)
				if err != nil {
					return
				}
				var data map[string]interface{}
				if err = json.Unmarshal(bytes.NewBuffer(resp.Payload).Bytes(), &data); err != nil {
					return
				}
				fmt.Println(data)
			}
		}
	}
}
