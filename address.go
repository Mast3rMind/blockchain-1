package blockchain

import (
    "encoding/hex"
    "crypto/rand"
)

type Address struct {
    Identifier string
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
