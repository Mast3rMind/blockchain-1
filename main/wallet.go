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
	"flag"
	"github.com/rvelhote/blockchain/cli"
	"github.com/rvelhote/blockchain/core"
	"log"
	"os"
	"path/filepath"
)

func main() {
	if len(os.Args) < 2 {
		log.Println("You must specify 'send' or 'balance' commands otherwise there is nothing I can do.")
		os.Exit(1)
	}

	path, _ := filepath.Abs(".wallet")
	err := os.MkdirAll(path, 0700)

	if err != nil {
		log.Fatal(err)
	}

	createCommand := flag.NewFlagSet("create", flag.ExitOnError)
	memorypoolCommand := flag.NewFlagSet("memorypool", flag.ExitOnError)

	sendCommand := flag.NewFlagSet("send", flag.ExitOnError)
	balanceCommand := flag.NewFlagSet("balance", flag.ExitOnError)

	addressTo := sendCommand.String("address", "", "The address to send money to")
	amountTo := sendCommand.Float64("amount", 0, "The amount of money to send to the address")

	addressCheck := balanceCommand.String("address", "", "The address to check the balance of. If empty it will use the wallet's ")

	switch os.Args[1] {
	case "send":
		log.Println(os.Args)
		sendCommand.Parse(os.Args[2:])

	case "balance":
		balanceCommand.Parse(os.Args[2:])

	case "create":
		createCommand.Parse(os.Args[2:])

	case "memorypool":
		memorypoolCommand.Parse(os.Args[2:])

	default:
		os.Exit(1)
	}

	if createCommand.Parsed() {
		addr, err := cli.Create()

		if err != nil {
			log.Fatal(err)
		}

		log.Println(addr.Identifier)
	}

	if sendCommand.Parsed() {
		cli.Send(core.NewTransaction(core.NewAddress("abc"), core.NewAddress(*addressTo), *amountTo))
	}

	if balanceCommand.Parsed() {
		cli.Balance(core.NewAddress(*addressCheck))
	}

	if memorypoolCommand.Parsed() {
		cli.MemoryPool()
	}
}
