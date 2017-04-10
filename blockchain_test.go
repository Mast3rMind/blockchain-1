// Package blockchain contains code related to this amazing custom blockchain including miners and a PoW algorithm
package blockchain

/**
 * The MIT License (MIT)
 *
 * Permission is hereby granted, free of charge, to any person obtaining a copy
 * of this software and associated documentation files (the "Software"), to deal
 * in the Software without restriction, including without limitation the rights
 * to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
 * copies of the Software, and to permit persons to whom the Software is
 * furnished to do so, subject to the following conditions:
 *
 * The above copyright notice and this permission notice shall be included in all
 * copies or substantial portions of the Software.
 *
 * THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
 * IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
 * FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
 * AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
 * LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
 * OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
 * SOFTWARE.
 */
import (
	"testing"
	"math/rand"
	"time"
)

func initializeBlockchain() (*Blockchain, Miner, []Address) {
	pow := ProofOfWork{}

	blockchain := &Blockchain{POW: pow, Difficulty: 1, DifficultyChanges: 10, Reward: 50}
	var initialTx []Transaction

	// Miners
	miner := Miner{Name: "Velhote Pool", POW: pow }

	seed := rand.NewSource(time.Now().UnixNano())
	rnd := rand.New(seed)

	genesisAddress := GenerateAddress()

	// Wallets
	var addresses []Address
	for i := 0; i < 15; i++ {
		address := Address{ Identifier: GenerateAddress() }
		addresses = append(addresses, address)
		initialTx = append(initialTx, Transaction{Hash: genesisAddress, From: Address{ Identifier: GenerateAddress() }, To: address, Amount: 1})
	}

	// Generate Transactions
	for i := 0; i < 1000; i++ {
		from := addresses[rnd.Intn(len(addresses))]
		to :=  addresses[rnd.Intn(len(addresses))]

		for from == to {
			to =  addresses[rnd.Intn(len(addresses))]
		}

		amount := rnd.Float64()

		blockchain.MemoryPool = append(blockchain.MemoryPool, Transaction{ From: from, To: to, Amount: amount, Fee: 0 })
	}

	blockchain.AddGenesisBlock("The Times 03/Jan/2009 Chancellor on brink of second bailout for banks", initialTx)
	return blockchain, miner, addresses
}

func TestBlockchain_AddBlock(t *testing.T) {
	blockchain, miner, addresses := initializeBlockchain()

	for len(blockchain.Blocks) != 45 {
		proof := miner.Mine(blockchain.Challenge, blockchain.Difficulty)
		blockchain.AddBlock(proof, &miner, miner.SelectTransactions(blockchain, 10))
	}

	for _, block := range blockchain.Blocks {
		t.Logf("%s, %d Transactions", block.MinedBy, len(block.Transactions))

		for _, transaction := range block.Transactions {
			t.Logf("From %s To %s Amount: %f", transaction.From.Identifier, transaction.To.Identifier, transaction.Amount)
		}
	}

	for _, address := range addresses {
		t.Logf("%s Balance is %f", address.Identifier, address.GetBalance(blockchain))
	}

	if blockchain.Difficulty != 5 {
		t.Errorf("The final difficulty after all the blocks have been added should be 5 and it's %d", blockchain.Difficulty)
	}

	if len(blockchain.Blocks) != 45 {
		t.Error("45 blocks should have been added to the blockchain after the tests are over")
	}

	if blockchain.IsValid() != true {
		t.Error("Blockchain should be valid")
	}
}
