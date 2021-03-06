/*
Copyright 2020 IBM All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package gateway

import "github.com/hyperledger/fabric-sdk-go/pkg/client/channel"

// A Contract object represents a smart contract instance in a network.
// Applications should get a Contract instance from a Network using the GetContract method
type Contract struct {
	chaincodeID string
	name        string
	network     *Network
	client      *channel.Client
}

func newContract(network *Network, chaincodeID string, name string) *Contract {
	return &Contract{network: network, client: network.client, chaincodeID: chaincodeID, name: name}
}

// Name returns the name of the smart contract
func (c *Contract) Name() string {
	return c.chaincodeID
}

// EvaluateTransaction will evaluate a transaction function and return its results.
// The transaction function 'name'
// will be evaluated on the endorsing peers but the responses will not be sent to
// the ordering service and hence will not be committed to the ledger.
// This can be used for querying the world state.
func (c *Contract) EvaluateTransaction(name string, args ...string) ([]byte, error) {
	txn, err := c.CreateTransaction(name)

	if err != nil {
		return nil, err
	}

	return txn.Evaluate(args...)
}

// SubmitTransaction will submit a transaction to the ledger. The transaction function 'name'
// will be evaluated on the endorsing peers and then submitted to the ordering service
// for committing to the ledger.
func (c *Contract) SubmitTransaction(name string, args ...string) ([]byte, error) {
	txn, err := c.CreateTransaction(name)

	if err != nil {
		return nil, err
	}

	return txn.Submit(args...)
}

// CreateTransaction creates an object representing a specific invocation of a transaction
// function implemented by this contract, and provides more control over
// the transaction invocation using the optional arguments. A new transaction object must
// be created for each transaction invocation.
func (c *Contract) CreateTransaction(name string, args ...TransactionOption) (*Transaction, error) {
	return newTransaction(name, c, args...)
}
