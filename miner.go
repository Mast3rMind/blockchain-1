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

// Miner is an instance of each miner that is registered to the Blockchain. It can be seens as the equivalent to a node
// in the Bitcoin/Cryptocoin ecosystem. The miner will register in the Blockchain and produce the work required to be
// able to register blocks in the system and receive rewards when they add/mine a block.
type Miner struct {
	Name        string
	POW         ProofOfWork
	Address		Address
}

func (m *Miner) Mine(challenge string, difficulty int) Proof {
	return m.POW.Solve(challenge, difficulty)
}

func (m *Miner) SelectTransactions(blockchain *Blockchain, n int) []Transaction {
	var transactions []Transaction

	if len(blockchain.MemoryPool) == 0 {
		return transactions
	}

	// We must also check the memory pool for the transaction values and make sure that, if multiple transactions are
	// made in the same block, the balance does not go below zero. This will make sure that such transactions are not
	// confirmed to the blockchain.
	totals := make(map[string]float64)

	for i := 0; i < n; i++ {
		from := blockchain.MemoryPool[i].From
		txamount := blockchain.MemoryPool[i].Amount

		if (from.GetBalance(blockchain) - totals[from.Identifier]) - txamount >= 0 {
			transactions = append(transactions, blockchain.MemoryPool[i])
			totals[from.Identifier] += txamount
		}
	}

	// TODO For sure there is a better way to do this
	blockchain.MemoryPool = blockchain.MemoryPool[n:]
	return transactions
}
