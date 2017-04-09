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
	RewardTotal float64
	Challenge   string
	Difficulty  int
}

func (m *Miner) Mine() Proof {
	return m.POW.Solve(m.Challenge, m.Difficulty)
}
