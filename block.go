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
	"crypto/sha512"
	"encoding/hex"
	"strconv"
	"time"
)

type Block struct {
	Index        int
	PreviousHash string
	Timestamp    time.Time
	Hash         string
	MinedBy string
	Transactions []Transaction
}

func (b *Block) compute() string {
	hasher := sha512.New()
	hasher.Write([]byte(strconv.Itoa(b.Index) + b.PreviousHash + b.Timestamp.String() + b.MinedBy))
	return hex.EncodeToString(hasher.Sum(nil))
}

func Genesis(minedBy string, transactions []Transaction) Block {
	genesis := Block{
		Index:        0,
		PreviousHash: "0",
		Timestamp:    time.Now(),
		MinedBy: minedBy,
		Transactions: transactions,
	}

	genesis.Hash = genesis.compute()
	return genesis
}

func NewBlock(previous Block, minedBy string, transactions []Transaction) Block {
	block := Block{
		Index:        previous.Index + 1,
		PreviousHash: previous.Hash,
		Timestamp:    time.Now(),
		Transactions: transactions,
		MinedBy: minedBy,
	}

	block.Hash = block.compute()
	return block
}
