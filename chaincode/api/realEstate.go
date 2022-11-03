package api

import (
	"chaincode/model"
	"chaincode/pkg/utils"
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
)

// CreateRealEstate New real estate (administrator)
func CreateRealEstate(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	// Verification parameter
	if len(args) != 4 {
		return shim.Error("The number of parameters is not satisfied")
	}
	accountId := args[0] //accountid is used to verify whether it is an administrator
	proprietor := args[1]
	totalArea := args[2]
	livingSpace := args[3]
	if accountId == "" || proprietor == "" || totalArea == "" || livingSpace == "" {
		return shim.Error("The parameter exists empty value")
	}
	if accountId == proprietor {
		return shim.Error("The operator should be an administrator and cannot be the same as everyone")
	}
	// Parameter data format conversion
	var formattedTotalArea float64
	if val, err := strconv.ParseFloat(totalArea, 64); err != nil {
		return shim.Error(fmt.Sprintf("Totalarea parameter format converts errors: %s", err))
	} else {
		formattedTotalArea = val
	}
	var formattedLivingSpace float64
	if val, err := strconv.ParseFloat(livingSpace, 64); err != nil {
		return shim.Error(fmt.Sprintf("Livingspace parameter format converts errors: %s", err))
	} else {
		formattedLivingSpace = val
	}
	//Determine whether the administrator's operation
	resultsAccount, err := utils.GetStateByPartialCompositeKeys(stub, model.AccountKey, []string{accountId})
	if err != nil || len(resultsAccount) != 1 {
		return shim.Error(fmt.Sprintf("Operation human rights limit verification failure%s", err))
	}
	var account model.Account
	if err = json.Unmarshal(resultsAccount[0], &account); err != nil {
		return shim.Error(fmt.Sprintf("Query operator information-Deep -sequential errors: %s", err))
	}
	if account.UserName != "administrator" {
		return shim.Error(fmt.Sprintf("Insufficient human rights%s", err))
	}
	//Determine whether the owner exists
	resultsProprietor, err := utils.GetStateByPartialCompositeKeys(stub, model.AccountKey, []string{proprietor})
	if err != nil || len(resultsProprietor) != 1 {
		return shim.Error(fmt.Sprintf("Owner Proprietor information verification failed%s", err))
	}
	realEstate := &model.RealEstate{
		RealEstateID: stub.GetTxID()[:16],
		Proprietor:   proprietor,
		Encumbrance:  false,
		TotalArea:    formattedTotalArea,
		LivingSpace:  formattedLivingSpace,
	}
	// Write account
	if err := utils.WriteLedger(realEstate, stub, model.RealEstateKey, []string{realEstate.Proprietor, realEstate.RealEstateID}); err != nil {
		return shim.Error(fmt.Sprintf("%s", err))
	}
	//Return to the information created successfully
	realEstateByte, err := json.Marshal(realEstate)
	if err != nil {
		return shim.Error(fmt.Sprintf("The information created by serialization is wrong: %s", err))
	}
	// Successfully return
	return shim.Success(realEstateByte)
}

// QueryRealEstateList Query real estate (you can query everything, or you can query the real estate according to everyone)
func QueryRealEstateList(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	var realEstateList []model.RealEstate
	results, err := utils.GetStateByPartialCompositeKeys2(stub, model.RealEstateKey, args)
	if err != nil {
		return shim.Error(fmt.Sprintf("%s", err))
	}
	for _, v := range results {
		if v != nil {
			var realEstate model.RealEstate
			err := json.Unmarshal(v, &realEstate)
			if err != nil {
				return shim.Error(fmt.Sprintf("QueryRealEstateList-Deep -sequential errors: %s", err))
			}
			realEstateList = append(realEstateList, realEstate)
		}
	}
	realEstateListByte, err := json.Marshal(realEstateList)
	if err != nil {
		return shim.Error(fmt.Sprintf("QueryRealEstateList-Serial errors: %s", err))
	}
	return shim.Success(realEstateListByte)
}
