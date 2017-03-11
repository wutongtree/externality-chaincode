package main

import (
	"errors"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/hyperledger/fabric/core/crypto/primitives"
	"github.com/op/go-logging"
)

var myLogger = logging.MustGetLogger("externality")

// ExternalityChaincode ExternalityChaincode
type ExternalityChaincode struct {
	stub shim.ChaincodeStubInterface
	args []string
}

// Init is called during Deploy transaction after the container has been
// established, allowing the chaincode to initialize its internal data
func (e *ExternalityChaincode) Init(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	myLogger.Debug("Init Chaincode...")

	e.stub = stub
	e.args = args

	myLogger.Debug("Init Chaincode...done")

	return nil, nil
}

// Invoke is called for every Invoke transactions. The chaincode may change
// its state variables
func (e *ExternalityChaincode) Invoke(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	myLogger.Debug("Invoke Chaincode...")

	e.stub = stub
	e.args = args

	if function == "" {

	}
	myLogger.Debug("Invoke Chaincode...done")

	return nil, errors.New("Received unknown function invocation")
}

// Query is called for Query transactions. The chaincode may only read
// (but not modify) its state variables and return the result
func (e *ExternalityChaincode) Query(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	myLogger.Debug("Query Chaincode...")

	e.stub = stub
	e.args = args
	if function == "" {

	}
	myLogger.Debug("Query Chaincode...done")

	return nil, errors.New("Received unknown function query")
}

func main() {
	primitives.SetSecurityLevel("SHA3", 256)
	err := shim.Start(new(ExternalityChaincode))
	if err != nil {
		myLogger.Errorf("Error starting exchange chaincode: %s", err)
	}
}
