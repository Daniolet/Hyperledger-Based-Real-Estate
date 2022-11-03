package utils

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/hyperledger/fabric/core/chaincode/shim"
)

// WriteLedger Write account
func WriteLedger(obj interface{}, stub shim.ChaincodeStubInterface, objectType string, keys []string) error {
	//Create a composite main key
	var key string
	if val, err := stub.CreateCompositeKey(objectType, keys); err != nil {
		return errors.New(fmt.Sprintf("%s-Create a composite main key to make an error %s", objectType, err))
	} else {
		key = val
	}
	bytes, err := json.Marshal(obj)
	if err != nil {
		return errors.New(fmt.Sprintf("%s-Serialized JSON data failed errors: %s", objectType, err))
	}
	//Write to the blockchain account book
	if err := stub.PutState(key, bytes); err != nil {
		return errors.New(fmt.Sprintf("%s-Write into the blockchain account book error: %s", objectType, err))
	}
	return nil
}

// DelLedger Delete the ledger
func DelLedger(stub shim.ChaincodeStubInterface, objectType string, keys []string) error {
	//Create a composite main key
	var key string
	if val, err := stub.CreateCompositeKey(objectType, keys); err != nil {
		return errors.New(fmt.Sprintf("%s-Create a composite main key to make an error %s", objectType, err))
	} else {
		key = val
	}
	//Write to the blockchain account book
	if err := stub.DelState(key); err != nil {
		return errors.New(fmt.Sprintf("%s-Delete the blockchain account error: %s", objectType, err))
	}
	return nil
}

// GetStateByPartialCompositeKeys Inquiry data based on the composite main key (suitable for obtaining all, multiple, single data)
// Demonstrate KEYS split inquiry
func GetStateByPartialCompositeKeys(stub shim.ChaincodeStubInterface, objectType string, keys []string) (results [][]byte, err error) {
	if len(keys) == 0 {
		// The length of the passing key is 0, then find and return all the data
		// Related data from the blockchain through the main key, which is equivalent to the vague query of the primary key
		resultIterator, err := stub.GetStateByPartialCompositeKey(objectType, keys)
		if err != nil {
			return nil, errors.New(fmt.Sprintf("%s-Get all data errors: %s", objectType, err))
		}
		defer resultIterator.Close()

		//Check whether the returned data is empty.
		for resultIterator.HasNext() {
			val, err := resultIterator.Next()
			if err != nil {
				return nil, errors.New(fmt.Sprintf("%s-Return data error: %s", objectType, err))
			}
			results = append(results, val.GetValue())
		}
	} else {
		// The length of the passing key is not 0, find the corresponding data and return
		for _, v := range keys {
			// Create a combination key
			key, err := stub.CreateCompositeKey(objectType, []string{v})
			if err != nil {
				return nil, errors.New(fmt.Sprintf("%s-Create a combination key error: %s", objectType, err))
			}
			// Get data from the ledger
			bytes, err := stub.GetState(key)
			if err != nil {
				return nil, errors.New(fmt.Sprintf("%s-Obtain data errors: %s", objectType, err))
			}

			if bytes != nil {
				results = append(results, bytes)
			}
		}
	}

	return results, nil
}

// GetStateByPartialCompositeKeys2 Inquiry data based on the composite main key (suitable for obtaining all or specified data)
func GetStateByPartialCompositeKeys2(stub shim.ChaincodeStubInterface, objectType string, keys []string) (results [][]byte, err error) {
	// Related data from the blockchain through the main key, which is equivalent to the vague query of the primary key
	resultIterator, err := stub.GetStateByPartialCompositeKey(objectType, keys)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("%s-Get all data errors: %s", objectType, err))
	}
	defer resultIterator.Close()

	//Check whether the returned data is empty.
	for resultIterator.HasNext() {
		val, err := resultIterator.Next()
		if err != nil {
			return nil, errors.New(fmt.Sprintf("%s-Return data error: %s", objectType, err))
		}

		results = append(results, val.GetValue())
	}
	return results, nil
}
