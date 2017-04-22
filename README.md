# Blockchain
A very basic blockchain that I am implementing just for fun and learning (no profit). It's a centralized in-memory 
blockchain (the blocks are stored in an array).

The proof of work is simply hashing a string until X amount of zeroes (given by current difficulty) is found. There 
is no secret to it - just looping until found.

## Features
This package features what you'd expect from a blockchain:
- The star of the show... the blockchain
- Miners will solve challenges and win rewards (only a single miner for now)
- A Proof of Work algorithm that miners need to solve to be able to add messages to the blockchain

## Future
- Support multiple miners that send blocks to the blockchain (the blockchain will always be centralized)
- Port the [Cuckoo Proof of Work](https://github.com/tromp/cuckoo) algorithm to Golang so we can use it
- Cancelation of mining work if a block is added which changes the challenge and the difficulty
- Transaction and Wallet Addresses
- Blockchain Explorer
- Improve the cryptography to use more "realistic" blockchain cryptography (e.g. merkle trees)