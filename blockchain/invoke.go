package blockchain

import (
  "fmt"
  "github.com/hyperledger/fabric-sdk-go/pkg/client/channel"
  "time"
)

// InvokeHello
func (setup *FabricSetup) InvokeHello(value string) (string, error) {

  // Prepare arguments
  var args []string
  args = append(args, "invoke")
  args = append(args, "hello")
  args = append(args, value)

  eventID := "eventInvoke"

  // Add data that will be visible in the proposal, like a description of the invoke request
  transientDataMap := make(map[string][]byte)
  transientDataMap["result"] = []byte("Transient data in hello invoke")

  reg, notifier, err := setup.event.RegisterChaincodeEvent(setup.ChainCodeID, eventID)
  if err != nil {
    return "", err
  }
  defer setup.event.Unregister(reg)

  // Create a request (proposal) and send it
  response, err := setup.client.Execute(channel.Request{ChaincodeID: setup.ChainCodeID, Fcn: args[0], Args: [][]byte{[]byte(args[1]), []byte(args[2])}, TransientMap: transientDataMap})
  if err != nil {
    return "", fmt.Errorf("failed to move funds: %v", err)
  }

  // Wait for the result of the submission
  select {
  case ccEvent := <-notifier:
    fmt.Printf("Received CC event: %v\n", ccEvent)
  case <-time.After(time.Second * 20):
    return "", fmt.Errorf("did NOT receive CC event for eventId(%s)", eventID)
  }

  return string(response.TransactionID), nil
}

// Initialise a new commitment spec on the blockchain
func (setup *FabricSetup) InvokeInitCommitment(inpargs []string) (string, error) {

  // Prepare arguments
  var args []string
  args = append(args, "initCommitment")
  args = append(args, inpargs[0])
  args = append(args, inpargs[1])
  args = append(args, inpargs[2])
  args = append(args, inpargs[3])
  args = append(args, inpargs[4])

  eventID := "eventInvoke"

  // Add data that will be visible in the proposal, like a description of the invoke request
  transientDataMap := make(map[string][]byte)
  transientDataMap["result"] = []byte("Transient data in init commitment invoke")

  reg, notifier, err := setup.event.RegisterChaincodeEvent(setup.ChainCodeID, eventID)
  if err != nil {
    return "", err
  }
  defer setup.event.Unregister(reg)

  // Create a request (proposal) and send it
  response, err := setup.client.Execute(channel.Request{ChaincodeID: setup.ChainCodeID, Fcn: args[0], Args: [][]byte{[]byte(args[1]), []byte(args[2]), []byte(args[3]), []byte(args[4]), []byte(args[5])}, TransientMap: transientDataMap})
  if err != nil {
    return "", fmt.Errorf("failed to move funds: %v", err)
  }

  // Wait for the result of the submission
  select {
  case ccEvent := <-notifier:
    fmt.Printf("Received CC event: %v\n", ccEvent)
  case <-time.After(time.Second * 20):
    return "", fmt.Errorf("did NOT receive CC event for eventId(%s) in init commitment", eventID)
  }

  return string(response.TransactionID), nil
}

// Add data to blockchain
func (setup *FabricSetup) InvokeInitCommitmentData(inpargs []string) (string, error) {

  // Prepare arguments
  var args []string
  args = append(args, "initCommitmentData")
  args = append(args, inpargs[0])
  args = append(args, inpargs[1])
  args = append(args, inpargs[2])
  args = append(args, inpargs[3])

  eventID := "eventInvoke"

  // Add data that will be visible in the proposal, like a description of the invoke request
  transientDataMap := make(map[string][]byte)
  transientDataMap["result"] = []byte("Transient data in init commitment data invoke")

  reg, notifier, err := setup.event.RegisterChaincodeEvent(setup.ChainCodeID, eventID)
  if err != nil {
    return "", err
  }
  defer setup.event.Unregister(reg)

  // Create a request (proposal) and send it
  response, err := setup.client.Execute(channel.Request{ChaincodeID: setup.ChainCodeID, Fcn: args[0], Args: [][]byte{[]byte(args[1]), []byte(args[2]), []byte(args[3]), []byte(args[4])}, TransientMap: transientDataMap})
  if err != nil {
    return "", fmt.Errorf("failed to move funds: %v", err)
  }

  // Wait for the result of the submission
  select {
  case ccEvent := <-notifier:
    fmt.Printf("Received CC event: %v\n", ccEvent)
  case <-time.After(time.Second * 20):
    return "", fmt.Errorf("did NOT receive CC event for eventId(%s) in init commitment data", eventID)
  }

  return string(response.TransactionID), nil
}
