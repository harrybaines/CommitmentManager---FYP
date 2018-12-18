package blockchain

import (
	"fmt"
	"github.com/hyperledger/fabric-sdk-go/pkg/client/channel"
)

// QueryHello query the chaincode to get the state of hello
func (setup *FabricSetup) QueryHello() (string, error) {

	// Prepare arguments
	var args []string
	args = append(args, "query")
	args = append(args, "hello")

	response, err := setup.client.Query(channel.Request{ChaincodeID: setup.ChainCodeID, Fcn: args[0], Args: [][]byte{[]byte(args[1])}})
	if err != nil {
		return "", fmt.Errorf("failed to query: %v", err)
	}

	return string(response.Payload), nil
}

// QueryCommitment query the chaincode to get the state of a commitment
func (setup *FabricSetup) QueryCommitment(name string) (string, error) {

  // Prepare arguments
  var args []string
  args = append(args, "readCommitment")
  args = append(args, name)

  response, err := setup.client.Query(channel.Request{ChaincodeID: setup.ChainCodeID, Fcn: args[0], Args: [][]byte{[]byte(args[1])}})
  if err != nil {
    return "", fmt.Errorf("failed to query: %v", err)
  }

  return string(response.Payload), nil
}

// RichQuery query the chaincode to perform an ad hoc rich query based on input
func (setup *FabricSetup) RichQuery(query string) (string, error) {

  // Prepare arguments
  var args []string
  args = append(args, "richQuery")
  args = append(args, query)

  response, err := setup.client.Query(channel.Request{ChaincodeID: setup.ChainCodeID, Fcn: args[0], Args: [][]byte{[]byte(args[1])}})
  if err != nil {
    return "", fmt.Errorf("failed to perform rich query: %v", err)
  }

  return string(response.Payload), nil
}

// ================================================================================
// QueryCreated obtains all the created commitments from the database
// ================================================================================
func (setup *FabricSetup) QueryCreated(query string) (string, error) {
  queryStr := fmt.Sprintf("{\"selector\":{\"docType\":\"commitment\"}}")
  return setup.RichQuery(queryStr)
}
