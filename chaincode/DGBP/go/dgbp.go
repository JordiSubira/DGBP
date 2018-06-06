/*
Copyright IBM Corp. 2016 All Rights Reserved.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

		 http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package main

import (
	"bytes"
	"fmt"
	"encoding/json"

	"github.com/hyperledger/fabric/core/chaincode/lib/cid"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
)

// SimpleChaincode example simple Chaincode implementation
type SimpleChaincode struct {
}

type User struct{
	ObjectType string `json:"docType"` //docType is used to distinguish the various types of objects in state database
	PKI string `json:"pki"`
	EID string `json:"eid"`
	MSP string `json:"msp"`
	Dep Dep `json:"dep"`
}

type Dep struct{	
	ObjectType string `json:"docType"` //docType is used to distinguish the various types of objects in state database
	MSP string `json:"msp"`
	Name string `json:"name"`
}

type Resource struct{
	ObjectType string `json:"docType"` //docType is used to distinguish the various types of objects in state database
	MSP string `json:"msp"`
	EID string `json:"eid"`
}

type Policy struct{
	ObjectType string `json:"docType"` //docType is used to distinguish the various types of objects in state database
	//srcMSP string `json:"srcMSP"`
	FromUserEid string `json:"fromUser"`
	//dstMSP string `json:"srcMSP"`
	To Resource `json:"toRes"`
}

type PolicyDep struct{
	ObjectType string `json:"docType"` //docType is used to distinguish the various types of objects in state database
	//srcMSP string `json:"srcMSP"`
	From Dep `json:"fromDep"`
	//dstMSP string `json:"srcMSP"`
	To Resource `json:"toRes"`
}
// Init initializes the chaincode
func (t *SimpleChaincode) Init(stub shim.ChaincodeStubInterface) pb.Response {

	fmt.Println("abac Init")

	//
	// Demonstrate the use of Attribute-Based Access Control (ABAC) by checking
	// to see if the caller has the "abac.init" attribute with a value of true;
	// if not, return an error.
	//

	MSPid , err := cid.GetMSPID(stub)
	if err != nil {
		return shim.Error("In CREATEUSER: Error: Failed to get MSPID")
	}
	fmt.Printf("MSPID= %s", MSPid)

	/*_, args := stub.GetFunctionAndParameters()
	var AdminOrg1, AdminOrg2 string    // Entities
	var AO1PubK, AO2PubKey int // Asset holdings

	if len(args) != 4 {
		return shim.Error("Incorrect number of arguments. Expecting 4")
	}*/

	//	Initialize the chaincode
	// 	Deps
	objectType := "dep"

	Dep1Org1 := Dep{objectType, "Org1MSP","Dep1"}
	Dep1Org1JSONasBytes , err := json.Marshal(Dep1Org1)
	if err != nil {
		return shim.Error(err.Error())
	}

	indexName := "msp~name"
	depMSPIndexKey, err := stub.CreateCompositeKey(indexName, []string{Dep1Org1.MSP, Dep1Org1.Name})
	if err != nil {
		return shim.Error(err.Error())
	}

	err = stub.PutState(depMSPIndexKey, Dep1Org1JSONasBytes)
	if err != nil {
		return shim.Error(err.Error())
	}	

	Dep1Org2 := Dep{objectType, "Org2MSP","Dep1"}
	Dep1Org2JSONasBytes , err := json.Marshal(Dep1Org2)
	if err != nil {
		return shim.Error(err.Error())
	}

	depMSPIndexKey, err = stub.CreateCompositeKey(indexName, []string{Dep1Org2.MSP, Dep1Org2.Name})
	if err != nil {
		return shim.Error(err.Error())
	}

	err = stub.PutState(depMSPIndexKey, Dep1Org2JSONasBytes)
	if err != nil {
		return shim.Error(err.Error())
	}

	//	Users

	objectType = "user"

	User1Org1 := &User{		ObjectType:objectType,
							PKI: "-----BEGIN PUBLIC KEY-----\nMFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAE7LVwkDctZRz/pQwed7ZH/TDod852\nBdeKu0SAOQnLfmGfbYIzbIyfhy83sjpFcqeUoLCQyJaPe07hWqk7Fqf5Lg==\n-----END PUBLIC KEY-----",
							EID:"192.0.2.1",
							MSP:"Org1MSP",
							Dep: Dep{
								ObjectType: "dep",
								MSP: "Org1MSP",
								Name: "Dep1",
								},
						}
	User1Org1JSONasBytes, err := json.Marshal(User1Org1)
	if err != nil {
		return shim.Error(err.Error())
	}

	User1Org2 := &User{objectType,"-----BEGIN PUBLIC KEY-----\nMFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAEMRqb3/wG9jir88KUfg4OlHigwN/b\nc+4BeMVaiFIknn6af6Vd5X+oVA2qViDGVG3O30RkcR2DnLrJm0bWWwoUDA==\n-----END PUBLIC KEY-----","192.0.2.2","Org2MSP",Dep1Org2}
	User1Org2JSONasBytes, err := json.Marshal(User1Org2)
	if err != nil {
		return shim.Error(err.Error())
	}

	err = stub.PutState(User1Org1.EID, User1Org1JSONasBytes)
	if err != nil {
		return shim.Error(err.Error())
	}
	/*indexName := "msp~pki"
	mspPkiIndexKey, err := stub.CreateCompositeKey(indexName, []string{User1Org1.MSP, User1Org1.PKI})
	if err != nil {
		return shim.Error(err.Error())
	}
	value := []byte{0x00}
	stub.PutState(mspPkiIndexKey, value)*/


	err = stub.PutState(User1Org2.EID, User1Org2JSONasBytes)
	if err != nil {
		return shim.Error(err.Error())
	}
	/*mspPkiIndexKey2, err := stub.CreateCompositeKey(indexName, []string{User1Org2.MSP, User1Org2.PKI})
	if err != nil {
		return shim.Error(err.Error())
	}
	stub.PutState(mspPkiIndexKey2, value)
	}*/

	fmt.Printf("User1Org1 EID= %s, User1Org2 EID= %s\n", User1Org1.EID, User1Org2.EID)


	return shim.Success(nil)
}

func (t *SimpleChaincode) Invoke(stub shim.ChaincodeStubInterface) pb.Response {
	fmt.Println("abac Invoke")
	function, args := stub.GetFunctionAndParameters()
	if function == "createUser" {
		// Make payment of X units from A to B
		return t.createUser(stub, args)
	} else if function == "deleteUser" {
		// Deletes an entity from its state
		return t.deleteUser(stub, args)
	} else if function == "queryUserByEID" {
		// the old "Query" is now implemtned in invoke
		return t.queryUserByEID(stub, args)
	} else if function == "queryUserByMSP" {
		// the old "Query" is now implemtned in invoke
		return t.queryUserByMSP(stub, args)
	} else if function == "createPolicy" {
		// the old "Query" is now implemtned in invoke
		return t.createPolicy(stub, args)
	} else if function == "deletePolicy" {
		// the old "Query" is now implemtned in invoke
		return t.deletePolicy(stub, args)
	} else if function == "queryPolicy" {
		// the old "Query" is now implemtned in invoke
		return t.queryPolicy(stub, args)
	}  else if function == "createPolicyDep" {
		// the old "Query" is now implemtned in invoke
		return t.createPolicyDep(stub, args)
	} else if function == "deletePolicyDep" {
		// the old "Query" is now implemtned in invoke
		return t.deletePolicyDep(stub, args)
	} else if function == "queryPolicyDep" {
		// the old "Query" is now implemtned in invoke
		return t.queryPolicyDep(stub, args)
	}else if function == "createDep" {
		// the old "Query" is now implemtned in invoke
		return t.createDep(stub, args)
	}else if function == "deleteDep" {
		// the old "Query" is now implemtned in invoke
		return t.deleteDep(stub, args)
	}else if function == "queryDep" {
		// the old "Query" is now implemtned in invoke
		return t.queryDep(stub, args)
	}else if function == "createResource" {
		// the old "Query" is now implemtned in invoke
		return t.createResource(stub, args)
	}else if function == "deleteResource" {
		// the old "Query" is now implemtned in invoke
		return t.deleteResource(stub, args)
	}else if function == "queryResource" {
		// the old "Query" is now implemtned in invoke
		return t.queryResource(stub, args)
	}/*else if function == "queryDepByMSP" {
		// the old "Query" is now implemtned in invoke
		return t.query(stub, args)
	}else if function == "queryTrustFrom" {
		// the old "Query" is now implemtned in invoke
		return t.query(stub, args)
	}*/



	return shim.Error("Invalid invoke function name. Expecting \"invoke\" \"delete\" \"query\"")
}

// Transaction makes payment of X units from A to B
func (t *SimpleChaincode) createUser(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	var eid,pki,msp,name_dep string    // Entities
	var err error

	//Check department exists

	objectType := "user"

	msp , err = cid.GetMSPID(stub)
	if err != nil {
		return shim.Error("In CREATEUSER: Error: Failed to get MSPID")
	}
	fmt.Printf("MSPID = %s", msp)

	if len(args) != 3 {
		return shim.Error("Incorrect number of arguments. Expecting 3")
	}

	pki = args[0]
	eid = args[1]
	name_dep = args[2]

	//Check department
	ret, err := checkDepartmentExists(stub,msp,name_dep)
	if err != nil{
		fmt.Printf("Error checking if Department exists")
		return shim.Error(err.Error())
	}else if ret == false{
		fmt.Println("Error Department %s does not exists, create dep before", name_dep)
		return shim.Error("Error Department does not exists: " + name_dep)
	}

	dep := Dep{"dep",msp,name_dep}
	

	userAsBytes, err := stub.GetState(eid)
	if err != nil {
		return shim.Error("Failed to get user: " + err.Error())
	} else if userAsBytes != nil {
		fmt.Println("User with EID = %s already exists: ", eid)
		return shim.Error("User with PKI already exists: " + eid)
	}

	user := &User{objectType,pki,eid,msp,dep}
	userAsBytes, err = json.Marshal(user)
	if err != nil {
		return shim.Error(err.Error())
	}

	err = stub.PutState(user.EID, userAsBytes)
	if err != nil {
		return shim.Error(err.Error())
	}
	/*indexName := "msp~pki"
	mspPkiIndexKey, err := stub.CreateCompositeKey(indexName, []string{user.MSP, user.PKI})
	if err != nil {
		return shim.Error(err.Error())
	}
	value := []byte{0x00}
	stub.PutState(mspPkiIndexKey, value)*/

	fmt.Printf("Success creating User EID= %s, from MSP = %s\n", user.EID, user.MSP)

	return shim.Success(nil)
}

// Deletes an entity from state
func (t *SimpleChaincode) deleteUser(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	var jsonResp string
	var userJSON User
	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}

	msp , err := cid.GetMSPID(stub)
	if err != nil {
		return shim.Error("In DELETEUSER: Error: Failed to get MSPID")
	}
	fmt.Printf("MSPID = %s", msp)

	eid := args[0]

	//Check if invoker MSP == user.MSP
	valAsbytes, err := stub.GetState(eid) //get the marble from chaincode state
	if err != nil {
		jsonResp = "{\"Error\":\"Failed to get state for user with EID = " + eid + "\"}"
		return shim.Error(jsonResp)
	} else if valAsbytes == nil {
		jsonResp = "{\"Error\":\"User with this EID does not exist: " + eid + "\"}"
		return shim.Error(jsonResp)
	}

	err = json.Unmarshal([]byte(valAsbytes), &userJSON)
	if err != nil {
		jsonResp = "{\"Error\":\"Failed to decode JSON of user: " + eid + "\"}"
		return shim.Error(jsonResp)
	}

	if userJSON.MSP != msp {
		fmt.Printf("User MSP = %s is different from invoker MSP = %s", userJSON.MSP, msp)
		jsonResp = "{\"Error\":\"User MSP:"+ userJSON.MSP +" != invoker MSP: " + msp + "\"}"
		return shim.Error(jsonResp)
	}

	// Delete the key from the state in ledger
	err = stub.DelState(eid)
	if err != nil {
		return shim.Error("Failed to delete state")
	}
	/*indexName := "msp~pki"
	mspPkiIndexKey, err := stub.CreateCompositeKey(indexName, []string{user.MSP, pki})
	if err != nil {
		return shim.Error(err.Error())
	}
	stub.DelState(mspPkiIndexKey)*/

	return shim.Success(nil)
}

// query callback representing the query of a chaincode
func (t *SimpleChaincode) queryUserByEID(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	var eid, jsonResp string
	var err error

	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting eid of the user to query")
	}

	eid = args[0]

	// Get the state from the ledger
	valAsbytes, err := stub.GetState(eid)
	if err != nil {
		jsonResp = "{\"Error\":\"Failed to get state for user with EID = " + eid + "\"}"
		return shim.Error(jsonResp)
	}else if valAsbytes == nil {
		jsonResp = "{\"Error\":\"User does not exist: " + eid + "\"}"
		return shim.Error(jsonResp)
	}

	jsonResp = "{\"User\":\"" + string(valAsbytes) + "\"}"
	fmt.Printf("Query Response:%s\n", jsonResp)
	return shim.Success(valAsbytes)
}

func (t *SimpleChaincode) queryUserByMSP(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	var msp string
	var err error

	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting MSP to query")
	}

	mspInvoker , err := cid.GetMSPID(stub)
	if err != nil {
		return shim.Error("In DELETEUSER: Error: Failed to get MSPID")
	}
	fmt.Printf("MSPID = %s", mspInvoker)

	msp = args[0]

	queryString := fmt.Sprintf("{\"selector\":{\"docType\":\"user\",\"msp\":\"%s\"}}", msp)

	queryResults, err := getQueryResultForQueryString(stub, queryString)
	if err != nil {
		return shim.Error(err.Error())
	}
	return shim.Success(queryResults)
}

func (t *SimpleChaincode) createPolicy(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	var dstMSP,fromUserEid,toResEid string    // Entities
	var err error

	objectType := "policy"


	dstMSP, err = cid.GetMSPID(stub)
	if err != nil {
		return shim.Error("In createPolicy: Error: Failed to get MSPID")
	}
	fmt.Printf("MSPID = %s", dstMSP)

	if len(args) != 2 {
		return shim.Error("Incorrect number of arguments. Expecting toResEid, fromUserEid")
	}

	toResEid = args[0]
	fromUserEid = args[1]

	//Check destination EID exists
	ret, err := checkResourceExists(stub,dstMSP,toResEid)
	if err != nil{
		fmt.Printf("Error checking if dest Resource exists")
		return shim.Error(err.Error())
	}else if ret == false{
		fmt.Println("Error Resource %s does not exists for MSP %s, create res before", toResEid,dstMSP)
		return shim.Error("Error Resource does not exists: " + toResEid)
	}

	toRes := Resource{"resource",dstMSP,toResEid}

	//Check source user exists
	ret, err = checkUserExists(stub,fromUserEid)
	if err != nil{
		fmt.Printf("Error checking if user exists")
		return shim.Error(err.Error())
	}else if ret == false{
		fmt.Println("Error User %s not found", fromUserEid)
		return shim.Error("Error User either does not exists: " + fromUserEid)
	}
	
	//...TO BE CONTINUED: include only composite key

	indexName := "dstres~srcuser"
	destResSrcUsrIndexKey, err := stub.CreateCompositeKey(indexName, []string{toResEid,fromUserEid})
	if err != nil {
		return shim.Error(err.Error())
	}

	AsBytes, err := stub.GetState(destResSrcUsrIndexKey)
	if err != nil {
		return shim.Error("Failed to get trust relationship: " + err.Error())
	} else if AsBytes != nil {
		fmt.Println("Trust relationship = %s already exists: ", destResSrcUsrIndexKey)
		return shim.Error("Trust relationship already exists: " + destResSrcUsrIndexKey)
	}

	policy := &Policy{objectType,fromUserEid,toRes}
	policyAsBytes, err := json.Marshal(policy)
	if err != nil {
		return shim.Error(err.Error())
	}

	err = stub.PutState(destResSrcUsrIndexKey, policyAsBytes)
	if err != nil {
		return shim.Error(err.Error())
	}

	fmt.Printf("Success creating relation from user = %s; to res = %s MSP = %s\n", fromUserEid, toResEid, dstMSP)

	return shim.Success(nil)
}

func (t *SimpleChaincode) deletePolicy(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	var err error

	dstMSP, err := cid.GetMSPID(stub)
	if err != nil {
		return shim.Error("In DELTRUST: Error: Failed to get dst MSPID")
	}
	fmt.Printf("MSPID = %s", dstMSP)

	if len(args) != 2 {
		return shim.Error("Incorrect number of arguments. Expecting toResEid, fromUserEid")
	}

	toResEid := args[0]
	fromUserEid := args[1]

	indexName := "dstres~srcuser"
	destResSrcUsrIndexKey, err := stub.CreateCompositeKey(indexName, []string{toResEid,fromUserEid})
	if err != nil {
		return shim.Error(err.Error())
	}

	//Check destination EID exists
	ret, err := checkResourceExists(stub,dstMSP,toResEid)
	if err != nil{
		fmt.Printf("Error checking if dest Resource exists")
		return shim.Error(err.Error())
	}else if ret == false{
		fmt.Println("Error Resource %s does not exists for MSP %s, create res before", toResEid,dstMSP)
		return shim.Error("Error Resource does not exists: " + toResEid)
	}

	// Delete the key from the state in ledger
	err = stub.DelState(destResSrcUsrIndexKey)
	if err != nil {
		return shim.Error("Failed to delete state")
	}
	/*indexName := "msp~pki"
	mspPkiIndexKey, err := stub.CreateCompositeKey(indexName, []string{user.MSP, pki})
	if err != nil {
		return shim.Error(err.Error())
	}
	stub.DelState(mspPkiIndexKey)*/

	return shim.Success(nil)
}

func (t *SimpleChaincode) queryPolicy(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	var jsonResp string
	var err error

	if len(args) != 2 {
		return shim.Error("Incorrect number of arguments. Expecting toResEid, fromUserEid")
	}

	toResEid := args[0]
	fromUserEid := args[1]

	indexName := "dstres~srcuser"
	destResSrcUsrIndexKey, err := stub.CreateCompositeKey(indexName, []string{toResEid,fromUserEid})
	if err != nil {
		return shim.Error(err.Error())
	}

	// Get the state from the ledger
	valAsbytes, err := stub.GetState(destResSrcUsrIndexKey)
	if err != nil {
		jsonResp = "{\"Error\":\"Failed to get state for policy = " + destResSrcUsrIndexKey + "\"}"
		return shim.Error(jsonResp)
	}else if valAsbytes == nil {
		jsonResp = "{\"Error\":\"Policy does not exist: " + destResSrcUsrIndexKey + "\"}"
		return shim.Error(jsonResp)
	}

	jsonResp = "{\"Policy\":\"" + string(valAsbytes) + "\"}"
	fmt.Printf("Query Response:%s\n", jsonResp)
	return shim.Success(valAsbytes)
}

func (t *SimpleChaincode) createPolicyDep(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	var srcMSP,dstMSP,fromDepName,toResEid string    // Entities
	var err error

	objectType := "policydep"


	dstMSP, err = cid.GetMSPID(stub)
	if err != nil {
		return shim.Error("In createPolicy: Error: Failed to get MSPID")
	}
	fmt.Printf("MSPID = %s", dstMSP)

	if len(args) != 3 {
		return shim.Error("Incorrect number of arguments. Expecting toResEid, srcMSP and srcDep")
	}

	toResEid = args[0]
	srcMSP = args[1]
	fromDepName = args[2]

	//Check destination department exists
	ret, err := checkDepartmentExists(stub,dstMSP,toResEid)
	if err != nil{
		fmt.Printf("Error checking if dest Resource exists")
		return shim.Error(err.Error())
	}else if ret == false{
		fmt.Println("Error Resource %s does not exists, create res before", toResEid)
		return shim.Error("Error Resource does not exists: " + toResEid)
	}

	toRes := Resource{"resource",dstMSP,toResEid}

	//Check source department exists
	ret, err = checkDepartmentExists(stub,srcMSP,fromDepName)
	if err != nil{
		fmt.Printf("Error checking if Department exists")
		return shim.Error(err.Error())
	}else if ret == false{
		fmt.Println("Error Department %s does not exists, create dep before", fromDepName)
		return shim.Error("Error Department does not exists: " + fromDepName)
	}

	fromDep := Dep{"dep",srcMSP,fromDepName}	
	
	//...TO BE CONTINUED: include only composite key

	indexName := "dstmsp~dstres~srcmsp~srcdep"
	destResSrcDepIndexKey, err := stub.CreateCompositeKey(indexName, []string{dstMSP,toResEid,srcMSP,fromDepName})
	if err != nil {
		return shim.Error(err.Error())
	}


	AsBytes, err := stub.GetState(destResSrcDepIndexKey)
	if err != nil {
		return shim.Error("Failed to get trust relationship: " + err.Error())
	} else if AsBytes != nil {
		fmt.Println("Trust relationship = %s already exists: ", destResSrcDepIndexKey)
		return shim.Error("Trust relationship already exists: " + destResSrcDepIndexKey)
	}

	policy := &PolicyDep{objectType,fromDep,toRes}
	policyAsBytes, err := json.Marshal(policy)
	if err != nil {
		return shim.Error(err.Error())
	}

	err = stub.PutState(destResSrcDepIndexKey, policyAsBytes)
	if err != nil {
		return shim.Error(err.Error())
	}

	fmt.Printf("Success creating relation from dep = %s MSP = %s; to res = %s MSP = %s\n", fromDepName, srcMSP, toResEid, dstMSP)

	return shim.Success(nil)
}

func (t *SimpleChaincode) deletePolicyDep(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	var err error

	dstMSP, err := cid.GetMSPID(stub)
	if err != nil {
		return shim.Error("In DELTRUST: Error: Failed to get src MSPID")
	}
	fmt.Printf("MSPID = %s", dstMSP)

	if len(args) != 3 {
		return shim.Error("Incorrect number of arguments. Expecting toResEid, srcMSP and srcDep")
	}

	toResEid := args[0]
	srcMSP := args[1]
	fromDepName := args[2]

	indexName := "dstmsp~dstres~srcmsp~srcdep"
	destResSrcDepIndexKey, err := stub.CreateCompositeKey(indexName, []string{dstMSP,toResEid,srcMSP,fromDepName})
	if err != nil {
		return shim.Error(err.Error())
	}

	// Delete the key from the state in ledger
	err = stub.DelState(destResSrcDepIndexKey)
	if err != nil {
		return shim.Error("Failed to delete state")
	}
	/*indexName := "msp~pki"
	mspPkiIndexKey, err := stub.CreateCompositeKey(indexName, []string{user.MSP, pki})
	if err != nil {
		return shim.Error(err.Error())
	}
	stub.DelState(mspPkiIndexKey)*/

	return shim.Success(nil)
}

func (t *SimpleChaincode) queryPolicyDep(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	var jsonResp string
	var err error

	if len(args) != 4 {
		return shim.Error("Incorrect number of arguments. Expecting dst MSP, toRes, srcMSP and fromDep")
	}

	toResEid := args[1]
	dstMSP := args[0]
	fromDepName := args[3]
	srcMSP := args[2]

	indexName := "dstmsp~dstres~srcmsp~srcdep"
	destResSrcDepIndexKey, err := stub.CreateCompositeKey(indexName, []string{dstMSP,toResEid,srcMSP,fromDepName})
	if err != nil {
		return shim.Error(err.Error())
	}

	// Get the state from the ledger
	valAsbytes, err := stub.GetState(destResSrcDepIndexKey)
	if err != nil {
		jsonResp = "{\"Error\":\"Failed to get state for policy = " + destResSrcDepIndexKey + "\"}"
		return shim.Error(jsonResp)
	}else if valAsbytes == nil {
		jsonResp = "{\"Error\":\"Policy does not exist: " + destResSrcDepIndexKey + "\"}"
		return shim.Error(jsonResp)
	}

	jsonResp = "{\"Policy\":\"" + string(valAsbytes) + "\"}"
	fmt.Printf("Query Response:%s\n", jsonResp)
	return shim.Success(valAsbytes)
}

func (t *SimpleChaincode) createDep(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	var name string    // Entities
	var err error

	objectType := "dep"


	msp, err := cid.GetMSPID(stub)
	if err != nil {
		return shim.Error("In CREATEDEP: Error: Failed to get MSPID")
	}
	fmt.Printf("MSPID = %s", msp)

	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting department name")
	}

	name = args[0]
	
	//...TO BE CONTINUED: include only composite key

	indexName := "msp~name"
	depMSPIndexKey, err := stub.CreateCompositeKey(indexName, []string{msp, name})
	if err != nil {
		return shim.Error(err.Error())
	}


	AsBytes, err := stub.GetState(depMSPIndexKey)
	if err != nil {
		return shim.Error("Failed to get department: " + err.Error())
	} else if AsBytes != nil {
		fmt.Println("Department = %s already exists: ", depMSPIndexKey)
		return shim.Error("Department already exists: " + depMSPIndexKey)
	}

	dep := &Dep{objectType,msp,name}
	depAsBytes, err := json.Marshal(dep)
	if err != nil {
		return shim.Error(err.Error())
	}

	err = stub.PutState(depMSPIndexKey, depAsBytes)
	if err != nil {
		return shim.Error(err.Error())
	}

	fmt.Printf("Success creating Department = %s; to MSP = %s\n", name, msp)

	return shim.Success(nil)
}

func (t *SimpleChaincode) deleteDep(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	var err error

	msp, err := cid.GetMSPID(stub)
	if err != nil {
		return shim.Error("In DELDEPARTMENT: Error: Failed to get src MSPID")
	}
	fmt.Printf("MSPID = %s", msp)

	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting Department name")
	}

	name := args[0]

	indexName := "msp~name"
	depMSPIndexKey, err := stub.CreateCompositeKey(indexName, []string{msp, name})
	if err != nil {
		return shim.Error(err.Error())
	}

	// Delete the key from the state in ledger
	err = stub.DelState(depMSPIndexKey)
	if err != nil {
		return shim.Error("Failed to delete state")
	}

	return shim.Success(nil)
}

func (t *SimpleChaincode) queryDep(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	var jsonResp string
	var err error

	if len(args) != 2 {
		return shim.Error("Incorrect number of arguments. Expecting MSP and name")
	}

	msp := args[0]
	name := args[1]

	indexName := "msp~name"
	depMSPIndexKey, err := stub.CreateCompositeKey(indexName, []string{msp, name})
	if err != nil {
		return shim.Error(err.Error())
	}

	// Get the state from the ledger
	valAsbytes, err := stub.GetState(depMSPIndexKey)
	if err != nil {
		jsonResp = "{\"Error\":\"Failed to get state for Department = " + depMSPIndexKey + "\"}"
		return shim.Error(jsonResp)
	}else if valAsbytes == nil {
		jsonResp = "{\"Error\":\"Dep does not exist: " + depMSPIndexKey + "\"}"
		return shim.Error(jsonResp)
	}

	jsonResp = "{\"Department\":\"" + string(valAsbytes) + "\"}"
	fmt.Printf("Query Response:%s\n", jsonResp)
	return shim.Success(valAsbytes)
}

func (t *SimpleChaincode) createResource(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	var eid string    // Entities
	var err error

	objectType := "resource"


	msp, err := cid.GetMSPID(stub)
	if err != nil {
		return shim.Error("In CREATERESOURCE: Error: Failed to get MSPID")
	}
	fmt.Printf("MSPID = %s", msp)

	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting Resource eid")
	}

	eid = args[0]
	
	//...TO BE CONTINUED: include only composite key

	indexName := "msp~eid"
	eidMSPIndexKey, err := stub.CreateCompositeKey(indexName, []string{msp, eid})
	if err != nil {
		return shim.Error(err.Error())
	}


	AsBytes, err := stub.GetState(eidMSPIndexKey)
	if err != nil {
		return shim.Error("Failed to get department: " + err.Error())
	} else if AsBytes != nil {
		fmt.Println("Resource = %s already exists: ", eidMSPIndexKey)
		return shim.Error("Resource already exists: " + eidMSPIndexKey)
	}

	res := &Resource{objectType,msp,eid}
	resAsBytes, err := json.Marshal(res)
	if err != nil {
		return shim.Error(err.Error())
	}

	err = stub.PutState(eidMSPIndexKey, resAsBytes)
	if err != nil {
		return shim.Error(err.Error())
	}

	fmt.Printf("Success creating Resource = %s; to MSP = %s\n", eid, msp)

	return shim.Success(nil)
}

func (t *SimpleChaincode) deleteResource(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	var err error

	msp, err := cid.GetMSPID(stub)
	if err != nil {
		return shim.Error("In DELRESOURCE: Error: Failed to get src MSPID")
	}
	fmt.Printf("MSPID = %s", msp)

	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting Resource eid")
	}

	eid := args[0]

	indexName := "msp~eid"
	eidMSPIndexKey, err := stub.CreateCompositeKey(indexName, []string{msp, eid})
	if err != nil {
		return shim.Error(err.Error())
	}

	// Delete the key from the state in ledger
	err = stub.DelState(eidMSPIndexKey)
	if err != nil {
		return shim.Error("Failed to delete state")
	}

	return shim.Success(nil)
}

func (t *SimpleChaincode) queryResource(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	var jsonResp string
	var err error

	if len(args) != 2 {
		return shim.Error("Incorrect number of arguments. Expecting MSP and eid")
	}

	msp := args[0]
	eid := args[1]

	indexName := "msp~eid"
	eidMSPIndexKey, err := stub.CreateCompositeKey(indexName, []string{msp, eid})
	if err != nil {
		return shim.Error(err.Error())
	}

	// Get the state from the ledger
	valAsbytes, err := stub.GetState(eidMSPIndexKey)
	if err != nil {
		jsonResp = "{\"Error\":\"Failed to get state for Resource = " + eidMSPIndexKey + "\"}"
		return shim.Error(jsonResp)
	}else if valAsbytes == nil {
		jsonResp = "{\"Error\":\"Resource does not exist: " + eidMSPIndexKey + "\"}"
		return shim.Error(jsonResp)
	}

	jsonResp = "{\"Resource\":\"" + string(valAsbytes) + "\"}"
	fmt.Printf("Query Response:%s\n", jsonResp)
	return shim.Success(valAsbytes)
}

// =========================================================================================
// getQueryResultForQueryString executes the passed in query string.
// Result set is built and returned as a byte array containing the JSON results.
// =========================================================================================
func getQueryResultForQueryString(stub shim.ChaincodeStubInterface, queryString string) ([]byte, error) {

	fmt.Printf("- getQueryResultForQueryString queryString:\n%s\n", queryString)

	resultsIterator, err := stub.GetQueryResult(queryString)
	if err != nil {
		return nil, err
	}
	defer resultsIterator.Close()

	// buffer is a JSON array containing QueryRecords
	var buffer bytes.Buffer
	buffer.WriteString("[")

	bArrayMemberAlreadyWritten := false
	for resultsIterator.HasNext() {
		queryResponse, err := resultsIterator.Next()
		if err != nil {
			return nil, err
		}
		// Add a comma before array members, suppress it for the first array member
		if bArrayMemberAlreadyWritten == true {
			buffer.WriteString(",")
		}
		buffer.WriteString("{\"Key\":")
		buffer.WriteString("\"")
		buffer.WriteString(queryResponse.Key)
		buffer.WriteString("\"")

		buffer.WriteString(", \"Record\":")
		// Record is a JSON object, so we write as-is
		buffer.WriteString(string(queryResponse.Value))
		buffer.WriteString("}")
		bArrayMemberAlreadyWritten = true
	}
	buffer.WriteString("]")

	fmt.Printf("- getQueryResultForQueryString queryResult:\n%s\n", buffer.String())

	return buffer.Bytes(), nil
}



func checkDepartmentExists(stub shim.ChaincodeStubInterface, msp string, name string) (bool, error){
	fmt.Println("checkDepartmentExists msp: %s, name: %s", msp, name)

	indexName := "msp~name"
	depMSPIndexKey, err := stub.CreateCompositeKey(indexName, []string{msp, name})
	if err != nil {
		return false, err
	}

	valAsbytes, err := stub.GetState(depMSPIndexKey)
	if err != nil {
		return false, err
	}else if valAsbytes == nil {
		return false, nil
	}else {
		return true, nil
	}
	
}

func checkUserExists(stub shim.ChaincodeStubInterface, eid string) (bool, error){
	//var userJSON User
	
	fmt.Println("checkUserExists msp: %s, eid: %s", eid)

	valAsbytes, err := stub.GetState(eid)
	if err != nil {
		fmt.Printf("Failed to get state for user with EID = %s", eid)
		return false, err
	} else if valAsbytes == nil {
		//err = "User with this EID does not exist" + eid
		return false, nil
	} else {
		return true, nil
	}

	/*err = json.Unmarshal([]byte(valAsbytes), &userJSON)
	if err != nil {
		return false, err
	}

	if userJSON.MSP != msp {
		fmt.Printf("User MSP = %s is different from invoker MSP = %s", userJSON.MSP, msp)
		//err = "User MSP is different from invoker MSP "
		return false, nil
	}
	return true, nil*/
	
}

func checkResourceExists(stub shim.ChaincodeStubInterface, msp string, eid string) (bool, error){
	var resJSON Resource
	
	fmt.Println("checkResourceExists msp: %s, eid: %s", msp, eid)

	valAsbytes, err := stub.GetState(eid)
	if err != nil {
		fmt.Printf("Failed to get state for resource with EID = %s", eid)
		return false, err
	} else if valAsbytes == nil {
		//err = "User with this EID does not exist" + eid
		return false, nil
	}

	err = json.Unmarshal([]byte(valAsbytes), &resJSON)
	if err != nil {
		return false, err
	}

	if resJSON.MSP != msp {
		fmt.Printf("User MSP = %s is different from invoker MSP = %s", resJSON.MSP, msp)
		//err = "User MSP is different from invoker MSP "
		return false, nil
	}
	return true, nil
	
}

func main() {
	err := shim.Start(new(SimpleChaincode))
	if err != nil {
		fmt.Printf("Error starting Simple chaincode: %s", err)
	}
}
