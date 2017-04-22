// Package main contains the main applications of the blockchain. A CLI interface and the blockchain server.
package main

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
	"encoding/json"
	"github.com/rvelhote/blockchain/core"
	"go.uber.org/zap"
	"net/http"
)

var pow core.ProofOfWork = core.ProofOfWork{}
var blockchain core.Blockchain = core.Blockchain{POW: pow, Difficulty: 1, DifficultyChanges: 10, Reward: 50}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/send", sendHandler)
	mux.HandleFunc("/balance", balanceHandler)
	mux.HandleFunc("/memorypool", memorypoolHandler)

	http.ListenAndServe(":8080", mux)
}
func memorypoolHandler(writer http.ResponseWriter, request *http.Request) {
	logger, _ := zap.NewDevelopment()
	sugar := logger.Sugar()

	for _, transaction := range blockchain.MemoryPool {
		sugar.Debugf("In Memory Pool -- From: %s To: %s Amount: %f", transaction.From.Identifier, transaction.To.Identifier, transaction.Amount)
	}

	writer.Header().Add("Content-Type", "application/json")
	writer.Write([]byte("true"))
}

func balanceHandler(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Add("Content-Type", "text/html")
	writer.Write([]byte("250"))
}

func sendHandler(writer http.ResponseWriter, request *http.Request) {
	logger, _ := zap.NewDevelopment()
	sugar := logger.Sugar()

	var transaction core.Transaction
	err := json.NewDecoder(request.Body).Decode(&transaction)

	sugar.Debug("Decoded JSON Request")

	if err != nil {
		sugar.Error(err)
		http.Error(writer, err.Error(), 400)
	}

	blockchain.AddTransaction(transaction)

	sugar.Debugf("Added Transaction From: %s To: %s Amount: %f", transaction.From.Identifier, transaction.To.Identifier, transaction.Amount)

	writer.Header().Add("Content-Type", "application/json")
	writer.Write([]byte("true"))
}
