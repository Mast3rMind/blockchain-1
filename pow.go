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
	"strings"
)

type Proof struct {
	Variation string
	Solution  string
}

type ProofOfWork struct {
}

func (pow *ProofOfWork) Verify(work string, difficulty int, proof Proof) bool {
	zeroes := strings.Repeat("0", difficulty)
	hasher := sha512.New()

	hasher.Write([]byte(work + proof.Variation))
	solution := hex.EncodeToString(hasher.Sum(nil))

	return strings.HasPrefix(solution, zeroes) && solution == proof.Solution
}

func (pow *ProofOfWork) Solve(work string, difficulty int) Proof {
	solution := ""
	zeroes := strings.Repeat("0", difficulty)

	variation := 0

	for {
		hasher := sha512.New()

		hasher.Write([]byte(work + strconv.Itoa(variation)))
		solution = hex.EncodeToString(hasher.Sum(nil))

		if strings.HasPrefix(solution, zeroes) {
			break
		}

		variation++
	}

	return Proof{
		Variation: strconv.Itoa(variation),
		Solution:  solution,
	}
}
