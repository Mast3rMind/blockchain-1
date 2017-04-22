// Package core contains code related to this amazing custom blockchain including miners and a PoW algorithm
package core

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
)

type Address struct {
	Identifier string `json:"Identifier"`
}

func NewAddress(address string) Address {
	return Address{
		Identifier: address,
	}
}

// GenerateAddress generates a new wallet address to be used for transactions
func GenerateAddress() string {
	size := 64

	bytes := make([]byte, size)
	rand.Read(bytes)

	return hex.EncodeToString(bytes)
}

func (a *Address) GetBalance(blockchain *Blockchain) float64 {
	var total float64 = 0.0

	for _, block := range blockchain.Blocks {
		for _, transaction := range block.Transactions {
			if transaction.From.Identifier == a.Identifier {
				total -= transaction.Amount
			} else if transaction.To.Identifier == a.Identifier {
				total += transaction.Amount
			}
		}
	}

	return total
}
