package api

import (
	"chaincode/model"
	"chaincode/pkg/utils"
	"encoding/json"
	"fmt"
	"time"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
)

// CreateDonating Initiate a donation
func CreateDonating(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	// Verification parameter
	if len(args) != 3 {
		return shim.Error("The number of parameters is not satisfied")
	}
	objectOfDonating := args[0]
	donor := args[1]
	grantee := args[2]
	if objectOfDonating == "" || donor == "" || grantee == "" {
		return shim.Error("The parameter exists empty value")
	}
	if donor == grantee {
		return shim.Error("The donor and the receiver cannot be the same person")
	}
	//Determine whether ObjectOfdonating belongs to DONOR
	resultsRealEstate, err := utils.GetStateByPartialCompositeKeys2(stub, model.RealEstateKey, []string{donor, objectOfDonating})
	if err != nil || len(resultsRealEstate) != 1 {
		return shim.Error(fmt.Sprintf("verify%S belongs to%S failed: %s", objectOfDonating, donor, err))
	}
	var realEstate model.RealEstate
	if err = json.Unmarshal(resultsRealEstate[0], &realEstate); err != nil {
		return shim.Error(fmt.Sprintf("CreateDonating-Deep -sequential errors: %s", err))
	}
	//Get information from Grantee
	resultsAccount, err := utils.GetStateByPartialCompositeKeys(stub, model.AccountKey, []string{grantee})
	if err != nil || len(resultsAccount) != 1 {
		return shim.Error(fmt.Sprintf("Grantee's gift information verification failed%s", err))
	}
	var accountGrantee model.Account
	if err = json.Unmarshal(resultsAccount[0], &accountGrantee); err != nil {
		return shim.Error(fmt.Sprintf("Query operator information-Deep -sequential errors: %s", err))
	}
	if accountGrantee.UserName == "administrator" {
		return shim.Error(fmt.Sprintf("Can't donate to administrator%s", err))
	}
	//Determine whether the record has existed, and the donation cannot be launched repeatedly
	//If Encumbrance is Tyue, it means that the property is already guaranteed
	if realEstate.Encumbrance {
		return shim.Error("This real estate has been used as a guarantee state and can no longer initiate a donation")
	}
	createTime, _ := stub.GetTxTimestamp()
	donating := &model.Donating{
		ObjectOfDonating: objectOfDonating,
		Donor:            donor,
		Grantee:          grantee,
		CreateTime:       time.Unix(int64(createTime.GetSeconds()), int64(createTime.GetNanos())).Local().Format("2006-01-02 15:04:05"),
		DonatingStatus:   model.DonatingStatusConstant()["donatingStart"],
	}
	// Write account
	if err := utils.WriteLedger(donating, stub, model.DonatingKey, []string{donating.Donor, donating.ObjectOfDonating, donating.Grantee}); err != nil {
		return shim.Error(fmt.Sprintf("%s", err))
	}
	//Set the status of the house to the state of guarantee
	realEstate.Encumbrance = true
	if err := utils.WriteLedger(realEstate, stub, model.RealEstateKey, []string{realEstate.Proprietor, realEstate.RealEstateID}); err != nil {
		return shim.Error(fmt.Sprintf("%s", err))
	}
	//Write the purchase transaction to the ledger,Can be used to inquire
	donatingGrantee := &model.DonatingGrantee{
		Grantee:    grantee,
		CreateTime: time.Unix(int64(createTime.GetSeconds()), int64(createTime.GetNanos())).Local().Format("2006-01-02 15:04:05"),
		Donating:   *donating,
	}
	if err := utils.WriteLedger(donatingGrantee, stub, model.DonatingGranteeKey, []string{donatingGrantee.Grantee, donatingGrantee.CreateTime}); err != nil {
		return shim.Error(fmt.Sprintf("Writing this donation transaction to the ledger failed%s", err))
	}
	donatingGranteeByte, err := json.Marshal(donatingGrantee)
	if err != nil {
		return shim.Error(fmt.Sprintf("The information created by serialization is wrong: %s", err))
	}
	// Successfully return
	return shim.Success(donatingGranteeByte)
}

// QueryDonatingList Query donation list(You can query everything, or you can inquire according to the initiative of the donor)(Initiative)(For donor inquiries)
func QueryDonatingList(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	var donatingList []model.Donating
	results, err := utils.GetStateByPartialCompositeKeys2(stub, model.DonatingKey, args)
	if err != nil {
		return shim.Error(fmt.Sprintf("%s", err))
	}
	for _, v := range results {
		if v != nil {
			var donating model.Donating
			err := json.Unmarshal(v, &donating)
			if err != nil {
				return shim.Error(fmt.Sprintf("QueryDonatingList-Deep -sequential errors: %s", err))
			}
			donatingList = append(donatingList, donating)
		}
	}
	donatingListByte, err := json.Marshal(donatingList)
	if err != nil {
		return shim.Error(fmt.Sprintf("QueryDonatingList-Serial errors: %s", err))
	}
	return shim.Success(donatingListByte)
}

// QueryDonatingListByGrantee According to the gift(ACCOUNTID)Query donation(Given)(Inquiry from the donor)
func QueryDonatingListByGrantee(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	if len(args) != 1 {
		return shim.Error(fmt.Sprintf("Must be specified by the gift of Acountid inquiries"))
	}
	var donatingGranteeList []model.DonatingGrantee
	results, err := utils.GetStateByPartialCompositeKeys2(stub, model.DonatingGranteeKey, args)
	if err != nil {
		return shim.Error(fmt.Sprintf("%s", err))
	}
	for _, v := range results {
		if v != nil {
			var donatingGrantee model.DonatingGrantee
			err := json.Unmarshal(v, &donatingGrantee)
			if err != nil {
				return shim.Error(fmt.Sprintf("QueryDonatingListByGrantee-Deep -sequential errors: %s", err))
			}
			donatingGranteeList = append(donatingGranteeList, donatingGrantee)
		}
	}
	donatingGranteeListByte, err := json.Marshal(donatingGranteeList)
	if err != nil {
		return shim.Error(fmt.Sprintf("QueryDonatingListByGrantee-Serial errors: %s", err))
	}
	return shim.Success(donatingGranteeListByte)
}

// UpdateDonating Update the donation status (confirm the gift, cancel)
func UpdateDonating(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	// Verification parameter
	if len(args) != 4 {
		return shim.Error("The number of parameters is not satisfied")
	}
	objectOfDonating := args[0]
	donor := args[1]
	grantee := args[2]
	status := args[3]
	if objectOfDonating == "" || donor == "" || grantee == "" || status == "" {
		return shim.Error("The parameter exists empty value")
	}
	if donor == grantee {
		return shim.Error("The donor and the receiver cannot be the same person")
	}
	//According to Objectofdonating and DONOR obtaining the real estate information you want to buy, it is confirmed that the property exists
	resultsRealEstate, err := utils.GetStateByPartialCompositeKeys2(stub, model.RealEstateKey, []string{donor, objectOfDonating})
	if err != nil || len(resultsRealEstate) != 1 {
		return shim.Error(fmt.Sprintf("according to%sand%S failed to get the real estate information you want to buy: %s", objectOfDonating, donor, err))
	}
	var realEstate model.RealEstate
	if err = json.Unmarshal(resultsRealEstate[0], &realEstate); err != nil {
		return shim.Error(fmt.Sprintf("UpdateDonating-Deep -sequential errors: %s", err))
	}
	//Obtained a gift based on Grantee
	resultsGranteeAccount, err := utils.GetStateByPartialCompositeKeys(stub, model.AccountKey, []string{grantee})
	if err != nil || len(resultsGranteeAccount) != 1 {
		return shim.Error(fmt.Sprintf("Grantee's gift information verification failed%s", err))
	}
	var accountGrantee model.Account
	if err = json.Unmarshal(resultsGranteeAccount[0], &accountGrantee); err != nil {
		return shim.Error(fmt.Sprintf("Query Grantee Gift Information: %s", err))
	}
	//Get donation information based on ObjectOfdonating and Donor and Grantee
	resultsDonating, err := utils.GetStateByPartialCompositeKeys2(stub, model.DonatingKey, []string{donor, objectOfDonating, grantee})
	if err != nil || len(resultsDonating) != 1 {
		return shim.Error(fmt.Sprintf("according to%sand%sand%sFailure to get sales information: %s", objectOfDonating, donor, grantee, err))
	}
	var donating model.Donating
	if err = json.Unmarshal(resultsDonating[0], &donating); err != nil {
		return shim.Error(fmt.Sprintf("UpdateDonating-Deep -sequential errors: %s", err))
	}
	//Regardless of whether it is completed or canceled, it is necessary to ensure that the donation is in the state of donation
	if donating.DonatingStatus != model.DonatingStatusConstant()["donatingStart"] {
		return shim.Error("This transaction is not in a donation, confirmed/Devising donation failed")
	}
	//Purchase information DonatingGraantee based on Grantee acquisition of buyers
	var donatingGrantee model.DonatingGrantee
	resultsDonatingGrantee, err := utils.GetStateByPartialCompositeKeys2(stub, model.DonatingGranteeKey, []string{grantee})
	if err != nil || len(resultsDonatingGrantee) == 0 {
		return shim.Error(fmt.Sprintf("according to%sObtaining the information of the gift failed: %s", grantee, err))
	}
	for _, v := range resultsDonatingGrantee {
		if v != nil {
			var s model.DonatingGrantee
			err := json.Unmarshal(v, &s)
			if err != nil {
				return shim.Error(fmt.Sprintf("UpdatedOnating: %s", err))
			}
			if s.Donating.ObjectOfDonating == objectOfDonating && s.Donating.Donor == donor && s.Grantee == grantee {
				//It must also be judged that the status must be delivered,To prevent the house from trading, it is just canceled
				if s.Donating.DonatingStatus == model.DonatingStatusConstant()["donatingStart"] {
					donatingGrantee = s
					break
				}
			}
		}
	}
	var data []byte
	//Determine donation status
	switch status {
	case "done":
		//Transfer real estate information to the gift and reset the guarantee status
		realEstate.Proprietor = grantee
		realEstate.Encumbrance = false
		//realEstate.RealEstateID = stub.GetTxID() //Re -update real estate ID
		if err := utils.WriteLedger(realEstate, stub, model.RealEstateKey, []string{realEstate.Proprietor, realEstate.RealEstateID}); err != nil {
			return shim.Error(fmt.Sprintf("%s", err))
		}
		//Clear the original real estate information
		if err := utils.DelLedger(stub, model.RealEstateKey, []string{donor, objectOfDonating}); err != nil {
			return shim.Error(fmt.Sprintf("%s", err))
		}
		//The donation status is set to complete, write the ledger
		donating.DonatingStatus = model.DonatingStatusConstant()["done"]
		donating.ObjectOfDonating = realEstate.RealEstateID //Re -update real estate ID
		if err := utils.WriteLedger(donating, stub, model.DonatingKey, []string{donating.Donor, objectOfDonating, grantee}); err != nil {
			return shim.Error(fmt.Sprintf("%s", err))
		}
		donatingGrantee.Donating = donating
		if err := utils.WriteLedger(donatingGrantee, stub, model.DonatingGranteeKey, []string{donatingGrantee.Grantee, donatingGrantee.CreateTime}); err != nil {
			return shim.Error(fmt.Sprintf("Writing this donation transaction to the ledger failed%s", err))
		}
		data, err = json.Marshal(donatingGrantee)
		if err != nil {
			return shim.Error(fmt.Sprintf("Information about serialized donation transactions is wrong: %s", err))
		}
		break
	case "cancelled":
		//Reset the status of real estate information guarantee status
		realEstate.Encumbrance = false
		if err := utils.WriteLedger(realEstate, stub, model.RealEstateKey, []string{realEstate.Proprietor, realEstate.RealEstateID}); err != nil {
			return shim.Error(fmt.Sprintf("%s", err))
		}
		//Update donation status
		donating.DonatingStatus = model.DonatingStatusConstant()["cancelled"]
		if err := utils.WriteLedger(donating, stub, model.DonatingKey, []string{donating.Donor, donating.ObjectOfDonating, donating.Grantee}); err != nil {
			return shim.Error(fmt.Sprintf("%s", err))
		}
		donatingGrantee.Donating = donating
		if err := utils.WriteLedger(donatingGrantee, stub, model.DonatingGranteeKey, []string{donatingGrantee.Grantee, donatingGrantee.CreateTime}); err != nil {
			return shim.Error(fmt.Sprintf("%s", err))
		}
		data, err = json.Marshal(donatingGrantee)
		if err != nil {
			return shim.Error(fmt.Sprintf("%s", err))
		}
		break
	default:
		return shim.Error(fmt.Sprintf("%sState is not supported", status))
	}
	return shim.Success(data)
}
