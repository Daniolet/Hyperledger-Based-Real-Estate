package api

import (
	"chaincode/model"
	"chaincode/pkg/utils"
	"encoding/json"
	"fmt"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
)

// QueryAccountList Query account list
func QueryAccountList(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	var accountList []model.Account
	results, err := utils.GetStateByPartialCompositeKeys(stub, model.AccountKey, args)
	if err != nil {
		return shim.Error(fmt.Sprintf("%s", err))
	}
	for _, v := range results {
		if v != nil {
			var account model.Account
			err := json.Unmarshal(v, &account)
			if err != nil {
				return shim.Error(fmt.Sprintf("QueryAccountList-Deep -sequential errors: %s", err))
			}
			accountList = append(accountList, account)
		}
	}
	accountListByte, err := json.Marshal(accountList)
	if err != nil {
		return shim.Error(fmt.Sprintf("QueryAccountList-Serial errors: %s", err))
	}
	return shim.Success(accountListByte)
}
