/* RxMed - Chaincode Base
* @author Ananthapadmanabhan (ananthan.vr@netobjex.com)
* Copyright netObjex, Inc. 2018 All Rights Reserved.
**/

package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strconv"
	"time"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
)

var logger = shim.NewLogger("example_cc0")

type rxMedChaincode struct {
}

type Doctor struct {
	DoctorID       string `json:"doctorid"`
	Name           string `json:"name"`
	RegisterNumber string `json:"registernumber"`
	Hospital       string `json:"hospital"`
}

type PatientPrivate struct {
	PatientID  string `json:"patientid"`
	Name       string `json:"name"`
	Dob        string `json:"dob"`
	Bloodgroup string `json:"bloodgroup"`
	Address    string `json:"address"`
}

type Medication struct {
	MedName  string `json:"medname"`
	Compound string `json:"compound"`
	Dosage   string `json:"dosage"`
	Quantity string `json:"quantity"`
}

type Patient struct {
	PatientID   string `json:"patientid`
	Medications []Medication
	Pin         int `json:"pin"`
}

type Pharmacy struct {
	PharmacyID string `json:"pharmacyid"`
	Name       string `json:"name"`
	Pin        string `json:"pin"`
	Owner      string `json:"owner"`
}

func (t *rxMedChaincode) Init(stub shim.ChaincodeStubInterface) pb.Response {
	logger.Info("########### RxMed Init ###########")

	// Get the function and arguments from the request
	function, _ := stub.GetFunctionAndParameters()
	//	var err error

	// Check if the request is the init function
	if function != "init" {
		return shim.Error("Unknown function call")
	}

	// Adding to the ledger.
	doctors := []Doctor{
		Doctor{DoctorID: "DOC1", Name: "Tittu Varghese", RegisterNumber: "MSD123", Hospital: "HOSP1"},
		Doctor{DoctorID: "DOC2", Name: "ABC CDF", RegisterNumber: "ABC143", Hospital: "HOSP2"},
	}

	i := 0
	for i < len(doctors) {
		fmt.Println("i is ", i)
		doctorAsBytes, _ := json.Marshal(doctors[i])
		stub.PutPrivateData("doctorcollection", "DOC"+strconv.Itoa(i), doctorAsBytes)
		fmt.Println("Added", doctors[i])
		i = i + 1
	}

	patients := []PatientPrivate{
		PatientPrivate{PatientID: "PAT1", Name: "Mahesh", Dob: "11/2/1990", Bloodgroup: "O+", Address: "bgfgdggdgdgf"},
		PatientPrivate{PatientID: "PAT2", Name: "Maheshwe", Dob: "11/2/1996", Bloodgroup: "B+", Address: "bgfgdggdgdgf"},
	}

	j := 0
	for j < len(patients) {
		fmt.Println("j is ", j)
		patientAsBytes, _ := json.Marshal(patients[j])
		stub.PutPrivateData("collectionPatientPrivate", "PAT"+strconv.Itoa(j), patientAsBytes)
		fmt.Println("Added", patients[j])
		j = j + 1
	}

	patients := []Patient{
		Patient{PatientID: "PAT1",[]Medication{ Medication{MedName: "ccc", Compound: "xxxxx", Dosage: "vvvv", Quantity: "bbbbb"}},Pin: "686101"}
		
	}

	j := 0
	for j < len(patients) {
		fmt.Println("j is ", j)
		patientAsBytes, _ := json.Marshal(patients[j])
		stub.PutPrivateData("collectionPatient", "PAT"+strconv.Itoa(j), patientAsBytes)
		fmt.Println("Added", patients[j])
		j = j + 1
	}

	pharmacies := []Pharmacy{
		Pharmacy{PharmacyID: "PHARM1", Name: "Shahnaz Pharmaceuticals", Pin: "691683", Owner: "Shah"},
	}

	k := 0
	for k < len(pharmacies) {
		fmt.Println("k is ", k)
		pharmacyAsBytes, _ := json.Marshal(pharmacies[k])
		stub.PutPrivateData("pharmacycollection", "PHARM"+strconv.Itoa(k), pharmacyAsBytes)
		fmt.Println("Added", pharmacies[k])
		k = k + 1
	}

	return shim.Success(nil)
}

func (t *rxMedChaincode) Invoke(stub shim.ChaincodeStubInterface) pb.Response {
	logger.Info("########### example_cc0 Invoke ###########")

	function, args := stub.GetFunctionAndParameters()

	if function == "query" {
		// queries an entity state
		return t.query(stub, args)
	}

	if function == "delete" {
		// Deletes an entity from its state
		return t.delete(stub, args)
	}

	if function == "queryAll" {
		// query all records
		return t.queryAll(stub, args)
	}

	if function == "queryHistory" {
		return t.queryHistory(stub, args)
	}

	if function == "createDoctor" {
		return t.createDoctor(stub, args)
	}

	if function == "createPatient" {
		return t.createDoctor(stub, args)
	}

	if function == "createPharmacy" {
		return t.createDoctor(stub, args)
	}

	if function == "updateDoctor" {
		return t.updateDoctor(stub, args)
	}

	if function == "updatePatient" {
		return t.updateDoctor(stub, args)
	}

	if function == "updatePharmacy" {
		return t.updateDoctor(stub, args)
	}

	logger.Errorf("Unknown action, check the first argument, must be one of 'delete', 'query', 'createDoctor' or 'updateDoctor'. But got: %v", args[0])
	return shim.Error(fmt.Sprintf("Unknown action, check the first argument, must be one of 'delete', 'query', 'createDoctor' or 'updateDoctor'. But got: %v", args[0]))
}

func (t *rxMedChaincode) query(stub shim.ChaincodeStubInterface, args []string) pb.Response {

	var A string // Entities
	var err error

	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting name of the person to query")
	}

	A = args[0]

	// Get the state from the ledger
	Avalbytes, err := stub.GetPrivateData("doctorcollections",A)
	if err != nil {
		jsonResp := "{\"Error\":\"Failed to get state for " + A + "\"}"
		return shim.Error(jsonResp)
	}

	if Avalbytes == nil {
		jsonResp := "{\"Error\":\"Nil amount for " + A + "\"}"
		return shim.Error(jsonResp)
	}

	jsonResp := "{\"Name\":\"" + A + "\",\"Amount\":\"" + string(Avalbytes) + "\"}"
	logger.Infof("Query Response:%s\n", jsonResp)
	return shim.Success(Avalbytes)
}

// Deletes an entity from state
func (t *rxMedChaincode) delete(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}

	searchKey := args[0]

	// Delete the key from the state in ledger
	err := stub.DelState(searchKey)
	if err != nil {
		return shim.Error("Failed to delete state")
	}

	return shim.Success(nil)
}

// Get History of a transaction by passing Key
func (t *rxMedChaincode) queryHistory(stub shim.ChaincodeStubInterface, args []string) pb.Response {

	if len(args) < 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}

	searchKey := args[0]
	fmt.Printf("##### start History of Record: %s\n", searchKey)

	resultsIterator, err := stub.GetHistoryForKey(searchKey)
	if err != nil {
		return shim.Error(err.Error())
	}
	defer resultsIterator.Close()

	// buffer is a JSON array containing historic values for the marble
	var buffer bytes.Buffer
	buffer.WriteString("[")

	bArrayMemberAlreadyWritten := false
	for resultsIterator.HasNext() {
		response, err := resultsIterator.Next()
		if err != nil {
			return shim.Error(err.Error())
		}
		// Add a comma before array members, suppress it for the first array member
		if bArrayMemberAlreadyWritten == true {
			buffer.WriteString(",")
		}
		buffer.WriteString("{\"TxId\":")
		buffer.WriteString("\"")
		buffer.WriteString(response.TxId)
		buffer.WriteString("\"")

		buffer.WriteString(", \"Value\":")
		// if it was a delete operation on given key, then we need to set the
		//corresponding value null. Else, we will write the response.Value
		//as-is (as the Value itself a JSON marble)
		if response.IsDelete {
			buffer.WriteString("null")
		} else {
			buffer.WriteString(string(response.Value))
		}

		buffer.WriteString(", \"Timestamp\":")
		buffer.WriteString("\"")
		buffer.WriteString(time.Unix(response.Timestamp.Seconds, int64(response.Timestamp.Nanos)).String())
		buffer.WriteString("\"")

		buffer.WriteString(", \"IsDelete\":")
		buffer.WriteString("\"")
		buffer.WriteString(strconv.FormatBool(response.IsDelete))
		buffer.WriteString("\"")

		buffer.WriteString("}")
		bArrayMemberAlreadyWritten = true
	}
	buffer.WriteString("]")

	fmt.Printf("- History of Doctor Record with key returning:\n%s\n", buffer.String())

	return shim.Success(buffer.Bytes())
}

// Query All callback representing the query of a chaincode
func (t *rxMedChaincode) queryAll(stub shim.ChaincodeStubInterface, args []string) pb.Response {

	resultsIterator, err := stub.GetStateByRange("", "")
	if err != nil {
		return shim.Error(err.Error())
	}
	defer resultsIterator.Close()
	// buffer is a JSON array containing QueryResults
	var buffer bytes.Buffer
	buffer.WriteString("[")

	bArrayMemberAlreadyWritten := false
	for resultsIterator.HasNext() {
		queryResponse, err := resultsIterator.Next()
		if err != nil {
			return shim.Error(err.Error())
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

	logger.Infof("Query All Doctor Info:%s\n", buffer.String())

	return shim.Success(buffer.Bytes())

}

func (t *rxMedChaincode) createDoctor(stub shim.ChaincodeStubInterface, args []string) pb.Response {

	if len(args) != 5 {
		return shim.Error("Incorrect number of arguments. Expecting 5 arguments for the invoke")
	}

	var doctor = Doctor{DoctorID: args[1], Name: args[2], RegisterNumber: args[3], Hospital: args[4]}
	docAsBytes, _ := json.Marshal(doctor)
	stub.PutPrivateData("doctorcollection", args[0], docAsBytes)

	logger.Infof("Create Doctor Response:%s\n", string(docAsBytes))

	// Transaction Response
	return shim.Success(docAsBytes)
}

func (t *rxMedChaincode) createPatientPrivate(stub shim.ChaincodeStubInterface, args []string) pb.Response {

	if len(args) != 5 {
		return shim.Error("Incorrect number of arguments. Expecting 5 arguments for the invoke")
	}

	var patient = Patient{PatientID: args[1], Name: args[2], Dob: args[3], Bloodgroup: args[4], Address: args[5]}
	patAsBytes, _ := json.Marshal(patient)
	stub.PutPrivateData("collectionPatientPrivate", args[0], patAsBytes)

	logger.Infof("Create Patient Response:%s\n", string(patAsBytes))

	// Transaction Response
	return shim.Success(patAsBytes)
}

// func (t *rxMedChaincode) createPatient(stub shim.ChaincodeStubInterface, args []string) pb.Response {

// 	if len(args) != 5 {
// 		return shim.Error("Incorrect number of arguments. Expecting 5 arguments for the invoke")
// 	}

// 	var patient = Patient{PatientID: args[1], []Medication{arg[2]}, Dob: args[3], Bloodgroup: args[4], Address: args[5]}
// 	patAsBytes, _ := json.Marshal(patient)
// 	stub.PutPrivateData("collectionPatientPrivate", args[0], patAsBytes)

// 	logger.Infof("Create Patient Response:%s\n", string(patAsBytes))

// 	// Transaction Response
// 	return shim.Success(patAsBytes)
// }

func (t *rxMedChaincode) createPharmacy(stub shim.ChaincodeStubInterface, args []string) pb.Response {

	if len(args) != 5 {
		return shim.Error("Incorrect number of arguments. Expecting 5 arguments for the invoke")
	}

	var pharmacy = Pharmacy{PharmacyID: args[1], Name: args[2], Pin: args[3], Owner: args[4]}
	pharmAsBytes, _ := json.Marshal(pharmacy)
	stub.PutPrivateData("pharmacycollection", args[0], pharmAsBytes)

	logger.Infof("Create Pharmacy Response:%s\n", string(pharmAsBytes))

	// Transaction Response
	return shim.Success(pharmAsBytes)
}

// Update a Doctor callback representing the transactionID
func (t *rxMedChaincode) updateDoctor(stub shim.ChaincodeStubInterface, args []string) pb.Response {

	if len(args) != 5 {
		return shim.Error("Incorrect number of arguments. Expecting 5 arguments for the invoke")
	}

	docAsBytes, _ := stub.GetState(args[0])
	/* If want to update all fields */
	doctor := Doctor{DoctorID: args[1], Name: args[2], RegisterNumber: args[3], Hospital: args[4]}
	/* if want to update a single field */
	/*
		docAsBytes, _ := stub.GetState(args[0])
		doctor := Doctor{}

		json.Unmarshal(docAsBytes, &doctor)
		doctor.Hospital = args[4]
	*/

	docAsBytes, _ = json.Marshal(doctor)
	stub.PutPrivateData("doctorcollection", args[0], docAsBytes)

	logger.Infof("Create Doctor Response:%s\n", string(docAsBytes))

	// Transaction Response
	return shim.Success(nil)
}

// Update a Patient callback representing the transactionID
func (t *rxMedChaincode) updatePatient(stub shim.ChaincodeStubInterface, args []string) pb.Response {

	if len(args) != 5 {
		return shim.Error("Incorrect number of arguments. Expecting 5 arguments for the invoke")
	}

	patAsBytes, _ := stub.GetState(args[0])
	/* If want to update all fields */
	patient := Patient{PatientID: args[1], Name: args[2], Dob: args[3], Bloodgroup: args[4]}
	/* if want to update a single field */
	/*
		patAsBytes, _ := stub.GetState(args[0])
		patient := Patient{}

		json.Unmarshal(patAsBytes, &patient)
		patient.Hospital = args[4]
	*/

	patAsBytes, _ = json.Marshal(patient)
	stub.PutPrivateData(args[0], patAsBytes)

	logger.Infof("Create Patient Response:%s\n", string(patAsBytes))

	// Transaction Response
	return shim.Success(nil)
}

// Update a Pharmacy callback representing the transactionID
func (t *rxMedChaincode) updatePharmacy(stub shim.ChaincodeStubInterface, args []string) pb.Response {

	if len(args) != 5 {
		return shim.Error("Incorrect number of arguments. Expecting 5 arguments for the invoke")
	}

	pharmAsBytes, _ := stub.GetState(args[0])
	/* If want to update all fields */
	pharmacy := Pharmacy{PharmacyID: args[1], Name: args[2], Pin: args[3], Owner: args[4]}
	/* if want to update a single field */
	/*
		pharmAsBytes, _ := stub.GetState(args[0])
		pharmacy := Pharmacy{}

		json.Unmarshal(pharmAsBytes, &pharmacy)
		pharmacy.Hospital = args[4]
	*/

	pharmAsBytes, _ = json.Marshal(pharmacy)
	stub.PutPrivateData(args[0], pharmAsBytes)

	logger.Infof("Create Pharmacy Response:%s\n", string(pharmAsBytes))

	// Transaction Response
	return shim.Success(nil)
}

func main() {
	err := shim.Start(new(rxMedChaincode))
	if err != nil {
		logger.Errorf("Error starting rxMed Chaincode: %s", err)
	}
}

