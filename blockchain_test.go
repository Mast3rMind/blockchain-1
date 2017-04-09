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
)

func TestBlockchain_AddBlock(t *testing.T) {
	pow := ProofOfWork{}

	blockchain := Blockchain{POW: pow, Difficulty: 1, DifficultyChanges: 10, Reward: 50}
	blockchain.AddGenesisBlock("The Times 03/Jan/2009 Chancellor on brink of second bailout for banks")

	miner := Miner{
		Name:        "Ricardo V.",
		POW:         pow,
		RewardTotal: 0,
		Challenge:   blockchain.Challenge,
		Difficulty:  blockchain.Difficulty,
	}

	for len(blockchain.Blocks) != 45 {
		t.Logf("Challenge is now: %s", miner.Challenge)
		t.Logf("Difficulty is now: %d", miner.Difficulty)
		t.Logf("Total Reward Received: %f", miner.RewardTotal)

		proof := miner.Mine()
		blockchain.AddBlock("xx", proof, &miner)

		miner.Challenge = blockchain.Challenge
		miner.Difficulty = blockchain.Difficulty
	}

	if blockchain.Difficulty != 5 {
		t.Errorf("The final difficulty after all the blocks have been added should be 5 and it's %d", blockchain.Difficulty)
	}

	if len(blockchain.Blocks) != 41 {
		t.Error("34 blocks should have been added to the blockchain after the tests are over")
	}

	if blockchain.IsValid() != true {
		t.Error("Blockchain should be valid")
	}
}
