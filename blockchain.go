package main

import "time"

type Blockchain struct {
	chain []Block
}

type Block struct {
	index        int64
	timestamp    time.Time
	transactions []Transaction
	proof        int64
	previousHash Hash
}

type Transaction struct {
	sender    Hash
	recipient Hash
	amount    int64
}

type Hash string

func NewBlockchain() *Blockchain {
	emp := &Blockchain{}
}

func NewBlock(proof int64, previous Hash) *Block {
	// TODO
}

func (b *Block) NewTransaction(sender, recipient Hash, amount int64) {
	b.transactions = append(b.transactions, Transaction{
		sender:    sender,
		recipient: recipient,
		amount:    amount,
	})
}
