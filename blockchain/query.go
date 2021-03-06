package blockchain

import (
	"fmt"
  "errors"
  "encoding/json"
	"github.com/hyperledger/fabric-sdk-go/pkg/client/channel"
)

// Mapping of commitment state to commitment query functions
var comStateFunctions = map[string]interface{}{
  "created": "getCreatedCommitments",
  "detached": "getDetachedCommitments",
  "expired": "getExpiredCommitments",
  "discharged": "getDischargedCommitments",
  "violated": "getViolatedCommitments",
}

// Represents a single commitment with an ID and slice of states
type Commitment struct {
  ComID    string
  States []ComState
}

// Represents a single commitment state - each has a name and a map of data associated with that state
type ComState struct {
  Name  string
  Data  map[string]interface{}
}

// A commitment specification
type Spec struct {
  ObjectType  string `json:"docType"`  // docType - used to distinguish the various types of objects in state database
  Name        string `json:"name"`     // Spec name - the name of the specification
  Source      string `json:"source"`   // Source - string to store spec source code (.quark file)
}

// GetSpec - query the chaincode to get the state of a spec
func (setup *FabricSetup) GetSpec(name string) (res *Spec, err error) {

  // Prepare results
  com := &Spec{}

  // Prepare arguments
  var args []string
  args = append(args, "getSpec")
  args = append(args, name)

  response, er := setup.client.Query(channel.Request{ChaincodeID: setup.ChainCodeID, Fcn: args[0], Args: [][]byte{[]byte(args[1])}})
  if er != nil {
    return com, fmt.Errorf("failed to query: %v", er)
  }

  json.Unmarshal([]byte(response.Payload), &com)
  return com, nil
}

// GetCommitments - query the chaincode to obtain commitments for a particular state
// States: created, detached, expired, discharged, violated
func (setup *FabricSetup) GetCommitments(comName string, comState string) (coms[] Commitment, err error) {

  // Prepare results
  commitments := []Commitment{}
  var chaincodeFunc string

  // Sanity check
  if chaincodeFunc = comStateFunctions[comState].(string); chaincodeFunc == "" {
    return commitments, errors.New("Unsupported commitment state chosen")
  }

  // Prepare arguments
  var args []string
  args = append(args, chaincodeFunc)
  args = append(args, comName)

  // Calls getDetachedCommitments/getDischargedCommitments in the chaincode logic with extra arg
  // Prevents repetition of code by using a boolean flag
  if chaincodeFunc == "getExpiredCommitments" || chaincodeFunc == "getViolatedCommitments" {
    args = append(args, "true")
  } else {
    args = append(args, "false")
  }

  response, err := setup.client.Query(channel.Request{ChaincodeID: setup.ChainCodeID, Fcn: args[0], Args: [][]byte{[]byte(args[1]), []byte(args[2])}})
  if err != nil {
    return commitments, fmt.Errorf("failed to query: %v", err)
  }

  json.Unmarshal([]byte(response.Payload), &commitments)
  return commitments, nil
}

// RichQuery - query the chaincode to perform an ad hoc rich query based on input
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
