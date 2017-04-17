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
	"crypto/rand"
	"encoding/hex"
	"errors"
	"log"
)

// Blockchain contains the information related to the current blockchain. The most important things are the blocks
// that make up the blockchain and the proof of work required to add blocks to the chain.
type Blockchain struct {
	Blocks            []Block
	POW               ProofOfWork
	Challenge         string
	DifficultyChanges int
	Difficulty        int
	Reward            float64
	MemoryPool        []Transaction
}

func (chain *Blockchain) AddTransaction(transaction Transaction) {
	chain.MemoryPool = append(chain.MemoryPool, transaction)
}

// AddGenesisBlock should be the first operation to perform when initializing a new blockchain. It adds the first
// block to the blockchain and does not require a PoW.
func (chain *Blockchain) AddGenesisBlock(minedBy string, transactions []Transaction) {
	chain.Blocks = append(chain.Blocks, Genesis(minedBy, transactions))
	chain.Challenge = chain.generateNewChallenge()
}

func (chain *Blockchain) getLastBlock() Block {
	if len(chain.Blocks) == 0 {
		log.Fatal("Genesis block does not exist. The genesis block must be generated first!")
	}
	return chain.Blocks[len(chain.Blocks)-1]
}

func (chain *Blockchain) generateNewChallenge() string {
	size := 64
	bytes := make([]byte, size)
	rand.Read(bytes)

	return hex.EncodeToString(bytes)
}

func (chain *Blockchain) DoWork() Proof {
	return chain.POW.Solve(chain.Challenge, chain.Difficulty)
}

// AddBlock adds a new block to the blockchain. It will require the data to be added to the blockchain as well as a
// proof of work that proves that the client has done the necessary work.
func (chain *Blockchain) AddBlock(proof Proof, miner *Miner, transactions []Transaction) error {
	if !chain.POW.Verify(chain.Challenge, chain.Difficulty, proof) {
		return errors.New("Invalid proof of work. Request to add block rejected")
	}

	if !chain.IsValid() {
		return errors.New("The chain is in an invalid state. Cannot add more blocks")
	}

	previous := chain.getLastBlock()
	block := NewBlock(previous, miner.Name, transactions)

	chain.Blocks = append(chain.Blocks, block)
	chain.Challenge = chain.generateNewChallenge()

	if len(chain.Blocks)%chain.DifficultyChanges == 0 {
		chain.Difficulty++
		chain.Reward /= 2
	}

	return nil
}

// IsValid confirms that the blockchain is in a valid state and all the blocks (except the genesis block) have a
// previous hash that points to the previous block.
func (chain *Blockchain) IsValid() bool {
	var previous Block

	for index, block := range chain.Blocks {
		// Ignores the genesis block from the previous block verification
		if index == 0 {
			previous = block
			continue
		}

		if block.PreviousHash != previous.Hash {
			return false
		}

		previous = block
	}

	return true
}
