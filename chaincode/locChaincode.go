package main

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

// LetterOfCredit defines the structure for the Letter of Credit
type LetterOfCredit struct {
	LOCID            string   `json:"locId"`
	Buyer            string   `json:"buyer"`
	Seller           string   `json:"seller"`
	IssuingBank      string   `json:"issuingBank"`
	AdvisingBank     string   `json:"advisingBank"`
	Amount           string   `json:"amount"`
	Currency         string   `json:"currency"`
	ExpiryDate       string   `json:"expiryDate"`
	GoodsDescription string   `json:"goodsDescription"`
	Status           string   `json:"status"`
	DocumentHashes   []string `json:"documentHashes"`
	History          []string `json:"history"`
}

type LOCStatus struct {
	Status string `json:"status"`
	Owner  string `json:"owner"`
}

// SmartContract provides functions for managing the Letter of Credit
type SmartContract struct {
	contractapi.Contract
}

// InitLedger initializes the chaincode (optional)
func (s *SmartContract) InitLedger(ctx contractapi.TransactionContextInterface) error {
	// TODO: Initialization code if needed
	return nil
}

// RequestLOC creates a new LoC request
func (s *SmartContract) RequestLOC(ctx contractapi.TransactionContextInterface, locID string, buyer string, seller string, issuingBank string, advisingBank string, amount string, currency string, expiryDate string, goodsDescription string) error {
	// TODO: Implement RequestLOC function
	// Verify caller is TataMotors
	// Create new LoC with status "Requested"
	// Add to history

	clientMSP, err := ctx.GetClientIdentity().GetMSPID()
	if err != nil {
		return fmt.Errorf("failed to get client msp id")
	}

	if clientMSP != "TataMotorsMSP" {
		return fmt.Errorf("unauthorized to request LOC only tata motors can request it")
	}

	if buyer != "TataMotors" {
		return fmt.Errorf("unverified buyer")
	}
	loc := &LetterOfCredit{
		LOCID:            locID,
		Buyer:            buyer,
		Seller:           seller,
		IssuingBank:      issuingBank,
		AdvisingBank:     advisingBank,
		Amount:           amount,
		Currency:         currency,
		ExpiryDate:       expiryDate,
		GoodsDescription: goodsDescription,
		Status:           "Requested",
		History:          []string{"Tata motors requested LOC"},
		DocumentHashes:   []string{},
	}

	locbytes, err := json.Marshal(loc)
	if err != nil {
		return err
	}

	err = ctx.GetStub().PutState(locID, locbytes)
	if err != nil {
		return err
	}

	err = ctx.GetStub().SetEvent("LOC_ISSUED", locbytes)
	if err != nil {
		return err
	}

	return nil
}

// TODO: Implement other functions here (IssueLOC, AcceptLOC, etc.)
func (s *SmartContract) IssueLOC(ctx contractapi.TransactionContextInterface, locID string) error {

	clientMSP, err := ctx.GetClientIdentity().GetMSPID()
	if err != nil {
		return fmt.Errorf("failed to get client msp id")
	}

	if clientMSP != "IciciMSP" {
		return fmt.Errorf("unauthorized to issue LOC only icici bank can issue it")
	}

	var loc LetterOfCredit
	locbytes, err := ctx.GetStub().GetState(locID)
	if err != nil {
		return err
	}

	err = json.Unmarshal(locbytes, &loc)
	if err != nil {
		return err
	}

	loc.Status = "Issued"
	loc.History = append(loc.History, "Letter od credit issued")

	locbytes, err = json.Marshal(loc)
	if err != nil {
		return err
	}

	err = ctx.GetStub().PutState(locID, locbytes)
	if err != nil {
		return err
	}

	return nil
}
func (s *SmartContract) AcceptLOC(ctx contractapi.TransactionContextInterface, locID string) error {

	clientMSP, err := ctx.GetClientIdentity().GetMSPID()
	if err != nil {
		return fmt.Errorf("failed to get client msp id")
	}

	if clientMSP != "TeslaMSP" {
		return fmt.Errorf("unauthorized to accept LOC only tesla can accept it")
	}

	var loc LetterOfCredit
	locbytes, err := ctx.GetStub().GetState(locID)
	if err != nil {
		return err
	}

	err = json.Unmarshal(locbytes, &loc)
	if err != nil {
		return err
	}

	loc.Status = "Accepted"
	loc.History = append(loc.History, "Letter of credit accepted")

	locbytes, err = json.Marshal(loc)
	if err != nil {
		return err
	}

	err = ctx.GetStub().PutState(locID, locbytes)
	if err != nil {
		return err
	}

	return nil
}

func (s *SmartContract) RejectLOC(ctx contractapi.TransactionContextInterface, locID string) error {

	clientMSP, err := ctx.GetClientIdentity().GetMSPID()
	if err != nil {
		return fmt.Errorf("failed to get client msp id")
	}

	if clientMSP != "TeslaMSP" {
		return fmt.Errorf("unauthorized to reject LOC only tesla can reject it")
	}

	var loc LetterOfCredit
	locbytes, err := ctx.GetStub().GetState(locID)
	if err != nil {
		return err
	}

	err = json.Unmarshal(locbytes, &loc)
	if err != nil {
		return err
	}

	loc.Status = "Rejected"
	loc.History = append(loc.History, "Letter of credit is rejected")

	locbytes, err = json.Marshal(loc)
	if err != nil {
		return err
	}

	err = ctx.GetStub().PutState(locID, locbytes)
	if err != nil {
		return err
	}

	return nil
}

func (s *SmartContract) ShipGoods(ctx contractapi.TransactionContextInterface, locID string) error {

	clientMSP, err := ctx.GetClientIdentity().GetMSPID()
	if err != nil {
		return fmt.Errorf("failed to get client msp id")
	}

	if clientMSP != "TeslaMSP" {
		return fmt.Errorf("unauthorized to ship goods only tesla can ship goods")
	}

	var loc LetterOfCredit
	locbytes, err := ctx.GetStub().GetState(locID)
	if err != nil {
		return err
	}

	err = json.Unmarshal(locbytes, &loc)
	if err != nil {
		return err
	}

	loc.Status = "Shipped"
	loc.History = append(loc.History, "Shipped Goods")

	h := sha256.New()
	h.Write([]byte("docs of shipment from tesla"))
	h.Sum(nil)
	loc.DocumentHashes = append(loc.DocumentHashes, string(h.Sum(nil)))

	locbytes, err = json.Marshal(loc)
	if err != nil {
		return err
	}

	err = ctx.GetStub().PutState(locID, locbytes)
	if err != nil {
		return err
	}

	err = ctx.GetStub().SetEvent("GOODS_SHIPPED", locbytes)
	if err != nil {
		return err
	}

	return nil
}

func (s *SmartContract) VerifyDocuments(ctx contractapi.TransactionContextInterface, locID string) error {

	clientMSP, err := ctx.GetClientIdentity().GetMSPID()
	if err != nil {
		return fmt.Errorf("failed to get client msp id")
	}

	if clientMSP != "ChaseMSP" {
		return fmt.Errorf("unauthorized to verify only chase bank can ship goods")
	}

	var loc LetterOfCredit
	locbytes, err := ctx.GetStub().GetState(locID)
	if err != nil {
		return err
	}

	err = json.Unmarshal(locbytes, &loc)
	if err != nil {
		return err
	}

	loc.Status = "Verified"
	loc.History = append(loc.History, "verified documents")

	locbytes, err = json.Marshal(loc)
	if err != nil {
		return err
	}

	err = ctx.GetStub().PutState(locID, locbytes)
	if err != nil {
		return err
	}

	err = ctx.GetStub().SetEvent("DOCUMENTS_VERIFIED", locbytes)
	if err != nil {
		return err
	}

	return nil
}

func (s *SmartContract) ReleasePayment(ctx contractapi.TransactionContextInterface, locID string) error {

	clientMSP, err := ctx.GetClientIdentity().GetMSPID()
	if err != nil {
		return fmt.Errorf("failed to get client msp id")
	}

	if clientMSP != "IciciMSP" {
		return fmt.Errorf("unauthorized to release payment only icici bank can release it")
	}

	var loc LetterOfCredit
	locbytes, err := ctx.GetStub().GetState(locID)
	if err != nil {
		return err
	}

	err = json.Unmarshal(locbytes, &loc)
	if err != nil {
		return err
	}

	loc.Status = "Paid"
	loc.History = append(loc.History, "paid the amount")

	locbytes, err = json.Marshal(loc)
	if err != nil {
		return err
	}

	err = ctx.GetStub().PutState(locID, locbytes)
	if err != nil {
		return err
	}

	err = ctx.GetStub().SetEvent("PAYMENT_RELEASED", locbytes)
	if err != nil {
		return err
	}

	return nil
}

func (s *SmartContract) GetLOCHistory(ctx contractapi.TransactionContextInterface, locID string) ([]string, error) {
	var loc LetterOfCredit
	locbytes, err := ctx.GetStub().GetState(locID)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(locbytes, &loc)
	if err != nil {
		return nil, err
	}

	return loc.History, nil
}

func (s *SmartContract) GetLOCStatus(ctx contractapi.TransactionContextInterface, locID string) (*LOCStatus, error) {
	var loc LetterOfCredit
	locbytes, err := ctx.GetStub().GetState(locID)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(locbytes, &loc)
	if err != nil {
		return nil, err
	}

	status := &LOCStatus{
		Status: loc.Status,
		Owner:  loc.Buyer,
	}

	return status, nil
}

func main() {
	chaincode, err := contractapi.NewChaincode(&SmartContract{})
	if err != nil {
		fmt.Printf("Error creating loc chaincode: %s", err.Error())
		return
	}

	if err := chaincode.Start(); err != nil {
		fmt.Printf("Error starting loc chaincode: %s", err.Error())
	}
}
