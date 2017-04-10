package blockchain

type Transaction struct {
    From Address
    To Address
    Amount float64
    Fee float64
    Hash string
}
