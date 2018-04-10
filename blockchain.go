package blockchain

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"time"
)

type Blockchain struct {
	chain               []*Block
	currentTransactions []*Transaction
}

type Block struct {
	index        int
	timestamp    time.Time
	transactions []*Transaction
	proof        int
	previousHash Hash
}

type Transaction struct {
	sender    Hash
	recipient Hash
	amount    int
}

type Hash string

func NewBlockchain() *Blockchain {
	bc := &Blockchain{}
	// genesis block
	bc.AddBlock(100, Hash("1"))
	return bc
}

func (bc *Blockchain) AddBlock(proof int, previousHash Hash) {
	b := &Block{
		index:        len(bc.chain) + 1,
		timestamp:    time.Now(),
		transactions: bc.currentTransactions,
		proof:        proof,
		previousHash: previousHash,
	}

	// reset current transactions
	bc.currentTransactions = []*Transaction{}

	bc.chain = append(bc.chain, b)
}

// AddTransaction add new transaction to chain, and returns latest block address
func (bc *Blockchain) AddTransaction(sender, recipient Hash, amount int) int {
	bc.currentTransactions = append(bc.currentTransactions, &Transaction{
		sender:    sender,
		recipient: recipient,
		amount:    amount,
	})
	return len(bc.chain)
}

func (bc *Blockchain) LastBlock() *Block {
	return bc.chain[len(bc.chain)-1]
}

func GenHash(b *Block) Hash {
	block, _ := json.Marshal(b)
	hash := sha256.Sum256(block)
	return Hash(hex.EncodeToString(hash[:]))
}

// ProofOfWork シンプルなプルーフ・オブ・ワークのアルゴリズム:
//         - hash(pp') の最初の4つが0となるような p' を探す
//         - p は1つ前のブロックのプルーフ、 p' は新しいブロックのプルーフ
func (bc *Blockchain) ProofOfWork(p int) int {
	proof := 0
	for {
		if ValidProof(p, proof) {
			return proof
		}
		proof++
	}
}

// ValidProof returns is valid proof of params
func ValidProof(lastProof, proof int) bool {
	guess := fmt.Sprintf("%d%d", lastProof, proof)
	guessHash := hex.EncodeToString([]byte(guess))

	return guessHash[:4] == "0000"
}
