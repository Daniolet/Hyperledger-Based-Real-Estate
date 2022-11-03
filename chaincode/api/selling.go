package api

import (
	"chaincode/model"
	"chaincode/pkg/utils"
	"encoding/json"
	"fmt"
	"strconv"
	"time"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
)

// CreateSelling Start sales
func CreateSelling(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	// Verification parameter
	if len(args) != 4 {
		return shim.Error("The number of parameters is not satisfied")
	}
	objectOfSale := args[0]
	seller := args[1]
	price := args[2]
	salePeriod := args[3]
	if objectOfSale == "" || seller == "" || price == "" || salePeriod == "" {
		return shim.Error("The parameter exists empty value")
	}
	// Parameter data format conversion
	var formattedPrice float64
	if val, err := strconv.ParseFloat(price, 64); err != nil {
		return shim.Error(fmt.Sprintf("Price parameter format converts errors: %s", err))
	} else {
		formattedPrice = val
	}
	var formattedSalePeriod int
	if val, err := strconv.Atoi(salePeriod); err != nil {
		return shim.Error(fmt.Sprintf("Saleperiod parameter format converts errors: %s", err))
	} else {
		formattedSalePeriod = val
	}
	//Determine whether Objectofsale belongs to Seller
	resultsRealEstate, err := utils.GetStateByPartialCompositeKeys2(stub, model.RealEstateKey, []string{seller, objectOfSale})
	if err != nil || len(resultsRealEstate) != 1 {
		return shim.Error(fmt.Sprintf("verify%sbelong%sfail: %s", objectOfSale, seller, err))
	}
	var realEstate model.RealEstate
	if err = json.Unmarshal(resultsRealEstate[0], &realEstate); err != nil {
		return shim.Error(fmt.Sprintf("CreateSelling-Deep -sequential errors: %s", err))
	}
	//Determine whether the record has existed, and the sales cannot be launched repeatedly
	//If Encumbrance is Tyue, it means that the property is already guaranteed
	if realEstate.Encumbrance {
		return shim.Error("This real estate has been used as a guarantee state and cannot be launched repeatedly")
	}
	createTime, _ := stub.GetTxTimestamp()
	selling := &model.Selling{
		ObjectOfSale:  objectOfSale,
		Seller:        seller,
		Buyer:         "",
		Price:         formattedPrice,
		CreateTime:    time.Unix(int64(createTime.GetSeconds()), int64(createTime.GetNanos())).Local().Format("2006-01-02 15:04:05"),
		SalePeriod:    formattedSalePeriod,
		SellingStatus: model.SellingStatusConstant()["saleStart"],
	}
	// Write account
	if err := utils.WriteLedger(selling, stub, model.SellingKey, []string{selling.Seller, selling.ObjectOfSale}); err != nil {
		return shim.Error(fmt.Sprintf("%s", err))
	}
	//Set the status of the house to the state of guarantee
	realEstate.Encumbrance = true
	if err := utils.WriteLedger(realEstate, stub, model.RealEstateKey, []string{realEstate.Proprietor, realEstate.RealEstateID}); err != nil {
		return shim.Error(fmt.Sprintf("%s", err))
	}
	//Return to the information created successfully
	sellingByte, err := json.Marshal(selling)
	if err != nil {
		return shim.Error(fmt.Sprintf("The information created by serialization is wrong: %s", err))
	}
	// Successfully return
	return shim.Success(sellingByte)
}

// CreateSellingByBuy Participate in sales (buyer purchased)
func CreateSellingByBuy(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	// Verification parameter
	if len(args) != 3 {
		return shim.Error("The number of parameters is not satisfied")
	}
	objectOfSale := args[0]
	seller := args[1]
	buyer := args[2]
	if objectOfSale == "" || seller == "" || buyer == "" {
		return shim.Error("The parameter exists empty value")
	}
	if seller == buyer {
		return shim.Error("Buyers and sellers cannot the same person")
	}
	//According to Objectofsale and Seller to obtain the real estate information you want to buy, it is confirmed that the property exists
	resultsRealEstate, err := utils.GetStateByPartialCompositeKeys2(stub, model.RealEstateKey, []string{seller, objectOfSale})
	if err != nil || len(resultsRealEstate) != 1 {
		return shim.Error(fmt.Sprintf("according to%sand%sObtaining the real estate information you want to buy failed: %s", objectOfSale, seller, err))
	}
	//Get sales information based on Objectofsale and Seller
	resultsSelling, err := utils.GetStateByPartialCompositeKeys2(stub, model.SellingKey, []string{seller, objectOfSale})
	if err != nil || len(resultsSelling) != 1 {
		return shim.Error(fmt.Sprintf("according to%sand%sFailure to get sales information: %s", objectOfSale, seller, err))
	}
	var selling model.Selling
	if err = json.Unmarshal(resultsSelling[0], &selling); err != nil {
		return shim.Error(fmt.Sprintf("CreateSellingBuy-Deep -sequential errors: %s", err))
	}
	//Determine whether Selling's status is in sales
	if selling.SellingStatus != model.SellingStatusConstant()["saleStart"] {
		return shim.Error("This transaction does not belong to the state of sales, and it can no longer be purchased")
	}
	//Get buyer information based on Buyer
	resultsAccount, err := utils.GetStateByPartialCompositeKeys(stub, model.AccountKey, []string{buyer})
	if err != nil || len(resultsAccount) != 1 {
		return shim.Error(fmt.Sprintf("Buyer buyer information verification failed%s", err))
	}
	var buyerAccount model.Account
	if err = json.Unmarshal(resultsAccount[0], &buyerAccount); err != nil {
		return shim.Error(fmt.Sprintf("Query Buyer buyer information-Deep -sequential errors: %s", err))
	}
	if buyerAccount.UserName == "administrator" {
		return shim.Error(fmt.Sprintf("Administrator cannot buy%s", err))
	}
	//Determine whether the balance is sufficient
	if buyerAccount.Balance < selling.Price {
		return shim.Error(fmt.Sprintf("Price of real estate%f,Your current balance is%f,Failed purchase", selling.Price, buyerAccount.Balance))
	}
	//Write buyer into the transaction selling and modify the transaction status
	selling.Buyer = buyer
	selling.SellingStatus = model.SellingStatusConstant()["delivery"]
	if err := utils.WriteLedger(selling, stub, model.SellingKey, []string{selling.Seller, selling.ObjectOfSale}); err != nil {
		return shim.Error(fmt.Sprintf("Write buyer into the transaction selling, modify the transaction status failed%s", err))
	}
	createTime, _ := stub.GetTxTimestamp()
	//Write the purchase transaction to the ledger,Available for buyers
	sellingBuy := &model.SellingBuy{
		Buyer:      buyer,
		CreateTime: time.Unix(int64(createTime.GetSeconds()), int64(createTime.GetNanos())).Local().Format("2006-01-02 15:04:05"),
		Selling:    selling,
	}
	if err := utils.WriteLedger(sellingBuy, stub, model.SellingBuyKey, []string{sellingBuy.Buyer, sellingBuy.CreateTime}); err != nil {
		return shim.Error(fmt.Sprintf("Writing the purchase transaction to the ledger failed%s", err))
	}
	sellingBuyByte, err := json.Marshal(sellingBuy)
	if err != nil {
		return shim.Error(fmt.Sprintf("The information created by serialization is wrong: %s", err))
	}
	//The purchase is successful, the balance is deducted, and the balance of the ledger is updated. Note that at this time, the seller is required to confirm the payment, and the payment will be transferred to the seller's account.
	buyerAccount.Balance -= selling.Price
	if err := utils.WriteLedger(buyerAccount, stub, model.AccountKey, []string{buyerAccount.AccountId}); err != nil {
		return shim.Error(fmt.Sprintf("Failure to deduct the balance of buyers%s", err))
	}
	// Successfully return
	return shim.Success(sellingBuyByte)
}

// QuerySellingList Query sales (which can be queried, or according to the initiative of the initiative) (initiated) (for the seller inquiry)
func QuerySellingList(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	var sellingList []model.Selling
	results, err := utils.GetStateByPartialCompositeKeys2(stub, model.SellingKey, args)
	if err != nil {
		return shim.Error(fmt.Sprintf("%s", err))
	}
	for _, v := range results {
		if v != nil {
			var selling model.Selling
			err := json.Unmarshal(v, &selling)
			if err != nil {
				return shim.Error(fmt.Sprintf("QuerySellingList-Retice-series errors: %s", err))
			}
			sellingList = append(sellingList, selling)
		}
	}
	sellingListByte, err := json.Marshal(sellingList)
	if err != nil {
		return shim.Error(fmt.Sprintf("QuerySellingList-serialization error: %s", err))
	}
	return shim.Success(sellingListByte)
}

// QuerySellingListByBuyer inquires sales based
func QuerySellingListByBuyer(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	if len(args) != 1 {
		return shim.Error(fmt.Sprintf("You must specify the buyer Acountid to query"))
	}
	var sellingBuyList []model.SellingBuy
	results, err := utils.GetStateByPartialCompositeKeys2(stub, model.SellingBuyKey, args)
	if err != nil {
		return shim.Error(fmt.Sprintf("%s", err))
	}
	for _, v := range results {
		if v != nil {
			var sellingBuy model.SellingBuy
			err := json.Unmarshal(v, &sellingBuy)
			if err != nil {
				return shim.Error(fmt.Sprintf("QuerySellingListBYBUYER-Reverse sequence error: %s", err))
			}
			sellingBuyList = append(sellingBuyList, sellingBuy)
		}
	}
	sellingBuyListByte, err := json.Marshal(sellingBuyList)
	if err != nil {
		return shim.Error(fmt.Sprintf("QuerySellingListBYBUYER-serialization error: %s", err))
	}
	return shim.Success(sellingBuyListByte)
}

// UpdateSelling Update sales status (buyer confirmation, buyer cancel)
func UpdateSelling(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	// Verification parameter
	if len(args) != 4 {
		return shim.Error("The number of parameters is not satisfied")
	}
	objectOfSale := args[0]
	seller := args[1]
	buyer := args[2]
	status := args[3]
	if objectOfSale == "" || seller == "" || status == "" {
		return shim.Error("The parameter exists empty value")
	}
	if buyer == seller {
		return shim.Error("Buyers and sellers cannot the same person")
	}
	//According to Objectofsale and Seller to obtain the real estate information you want to buy, it is confirmed that the property exists
	resultsRealEstate, err := utils.GetStateByPartialCompositeKeys2(stub, model.RealEstateKey, []string{seller, objectOfSale})
	if err != nil || len(resultsRealEstate) != 1 {
		return shim.Error(fmt.Sprintf("according to%sand%S failed to get the real estate information you want to buy: %s", objectOfSale, seller, err))
	}
	var realEstate model.RealEstate
	if err = json.Unmarshal(resultsRealEstate[0], &realEstate); err != nil {
		return shim.Error(fmt.Sprintf("UpdateSellingBy Seller-Reverse sequence error: %s", err))
	}
	//Get sales information based on Objectofsale and Seller
	resultsSelling, err := utils.GetStateByPartialCompositeKeys2(stub, model.SellingKey, []string{seller, objectOfSale})
	if err != nil || len(resultsSelling) != 1 {
		return shim.Error(fmt.Sprintf("according to%sand%sFailure to get sales information: %s", objectOfSale, seller, err))
	}
	var selling model.Selling
	if err = json.Unmarshal(resultsSelling[0], &selling); err != nil {
		return shim.Error(fmt.Sprintf("UpdateSellingBy Seller-Reverse sequence error: %s", err))
	}
	//Get the buyer to buy information for buyers based on Buyer
	var sellingBuy model.SellingBuy
	//If the current state is in salestart sales, there is no buyer
	if selling.SellingStatus != model.SellingStatusConstant()["saleStart"] {
		resultsSellingByBuyer, err := utils.GetStateByPartialCompositeKeys2(stub, model.SellingBuyKey, []string{buyer})
		if err != nil || len(resultsSellingByBuyer) == 0 {
			return shim.Error(fmt.Sprintf("according to%S failed to get buyers' purchase information: %s", buyer, err))
		}
		for _, v := range resultsSellingByBuyer {
			if v != nil {
				var s model.SellingBuy
				err := json.Unmarshal(v, &s)
				if err != nil {
					return shim.Error(fmt.Sprintf("UpdateSellingBySeller-Deep -sequential errors: %s", err))
				}
				if s.Selling.ObjectOfSale == objectOfSale && s.Selling.Seller == seller && s.Buyer == buyer {
					//It must also be judged that the status must be delivered,To prevent the house from trading, it is just canceled
					if s.Selling.SellingStatus == model.SellingStatusConstant()["delivery"] {
						sellingBuy = s
						break
					}
				}
			}
		}
	}
	var data []byte
	//Judgment sales status
	switch status {
	case "done":
		//If the buyer confirms the receipt operation,Must ensure that the sales are in delivery
		if selling.SellingStatus != model.SellingStatusConstant()["delivery"] {
			return shim.Error("This transaction is not in the delivery, confirming the failure of the payment")
		}
		//Get seller information based on Seller
		resultsSellerAccount, err := utils.GetStateByPartialCompositeKeys(stub, model.AccountKey, []string{seller})
		if err != nil || len(resultsSellerAccount) != 1 {
			return shim.Error(fmt.Sprintf("Seller seller information verification failed%s", err))
		}
		var accountSeller model.Account
		if err = json.Unmarshal(resultsSellerAccount[0], &accountSeller); err != nil {
			return shim.Error(fmt.Sprintf("Query Seller seller information-Deep -sequential errors: %s", err))
		}
		//confirmed paid,Add the money to the seller account
		accountSeller.Balance += selling.Price
		if err := utils.WriteLedger(accountSeller, stub, model.AccountKey, []string{accountSeller.AccountId}); err != nil {
			return shim.Error(fmt.Sprintf("The seller confirmed that the receiving funds failed%s", err))
		}
		//Transfer real estate information to the buyer and reset the guarantee status
		realEstate.Proprietor = buyer
		realEstate.Encumbrance = false
		//realEstate.RealEstateID = stub.GetTxID() //Re -update real estate ID
		if err := utils.WriteLedger(realEstate, stub, model.RealEstateKey, []string{realEstate.Proprietor, realEstate.RealEstateID}); err != nil {
			return shim.Error(fmt.Sprintf("%s", err))
		}
		//Clear the original real estate information
		if err := utils.DelLedger(stub, model.RealEstateKey, []string{seller, objectOfSale}); err != nil {
			return shim.Error(fmt.Sprintf("%s", err))
		}
		//The order status is set to complete, write the ledger
		selling.SellingStatus = model.SellingStatusConstant()["done"]
		selling.ObjectOfSale = realEstate.RealEstateID //Re -update real estate ID
		if err := utils.WriteLedger(selling, stub, model.SellingKey, []string{selling.Seller, objectOfSale}); err != nil {
			return shim.Error(fmt.Sprintf("%s", err))
		}
		sellingBuy.Selling = selling
		if err := utils.WriteLedger(sellingBuy, stub, model.SellingBuyKey, []string{sellingBuy.Buyer, sellingBuy.CreateTime}); err != nil {
			return shim.Error(fmt.Sprintf("Writing the purchase transaction to the ledger failed%s", err))
		}
		data, err = json.Marshal(sellingBuy)
		if err != nil {
			return shim.Error(fmt.Sprintf("The information of the serialization purchase transaction is wrong: %s", err))
		}
		break
	case "cancelled":
		data, err = closeSelling("cancelled", selling, realEstate, sellingBuy, buyer, stub)
		if err != nil {
			return shim.Error(fmt.Sprintf("%s", err))
		}
		break
	case "expired":
		data, err = closeSelling("expired", selling, realEstate, sellingBuy, buyer, stub)
		if err != nil {
			return shim.Error(fmt.Sprintf("%s", err))
		}
		break
	default:
		return shim.Error(fmt.Sprintf("%S status does not support", status))
	}
	return shim.Success(data)
}

// closeSelling Whether it is canceled or expired, there are two cases
// 1. It is currently in salestart sales status
// 2. It is currently in Delivery delivery status
func closeSelling(closeStart string, selling model.Selling, realEstate model.RealEstate, sellingBuy model.SellingBuy, buyer string, stub shim.ChaincodeStubInterface) ([]byte, error) {
	switch selling.SellingStatus {
	case model.SellingStatusConstant()["saleStart"]:
		selling.SellingStatus = model.SellingStatusConstant()[closeStart]
		//Reset the status of real estate information guarantee status
		realEstate.Encumbrance = false
		if err := utils.WriteLedger(realEstate, stub, model.RealEstateKey, []string{realEstate.Proprietor, realEstate.RealEstateID}); err != nil {
			return nil, err
		}
		if err := utils.WriteLedger(selling, stub, model.SellingKey, []string{selling.Seller, selling.ObjectOfSale}); err != nil {
			return nil, err
		}
		data, err := json.Marshal(selling)
		if err != nil {
			return nil, err
		}
		return data, nil
	case model.SellingStatusConstant()["delivery"]:
		//Get seller information based on Buyer
		resultsBuyerAccount, err := utils.GetStateByPartialCompositeKeys(stub, model.AccountKey, []string{buyer})
		if err != nil || len(resultsBuyerAccount) != 1 {
			return nil, err
		}
		var accountBuyer model.Account
		if err = json.Unmarshal(resultsBuyerAccount[0], &accountBuyer); err != nil {
			return nil, err
		}
		//Cancellation at this time, you need to return the funds to the buyer
		accountBuyer.Balance += selling.Price
		if err := utils.WriteLedger(accountBuyer, stub, model.AccountKey, []string{accountBuyer.AccountId}); err != nil {
			return nil, err
		}
		//Reset the status of real estate information guarantee status
		realEstate.Encumbrance = false
		if err := utils.WriteLedger(realEstate, stub, model.RealEstateKey, []string{realEstate.Proprietor, realEstate.RealEstateID}); err != nil {
			return nil, err
		}
		//Update sales status
		selling.SellingStatus = model.SellingStatusConstant()[closeStart]
		if err := utils.WriteLedger(selling, stub, model.SellingKey, []string{selling.Seller, selling.ObjectOfSale}); err != nil {
			return nil, err
		}
		sellingBuy.Selling = selling
		if err := utils.WriteLedger(sellingBuy, stub, model.SellingBuyKey, []string{sellingBuy.Buyer, sellingBuy.CreateTime}); err != nil {
			return nil, err
		}
		data, err := json.Marshal(sellingBuy)
		if err != nil {
			return nil, err
		}
		return data, nil
	default:
		return nil, nil
	}
}
