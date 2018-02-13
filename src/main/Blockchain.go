package main

import (
	"fmt"
	"strconv"
	"bytes"
	"crypto/sha256"
	"time"
)

type Block struct {
	Index		  int
	Timestamp     int64
	Data          []byte
	PrevBlockHash []byte
	Hash          []byte
}

type Blockchain struct {
	blocks []*Block
}

func (b *Block) SetHash() {
	timestamp := []byte(strconv.FormatInt(b.Timestamp, 10))
	headers := bytes.Join([][]byte{b.PrevBlockHash, b.Data, timestamp}, []byte{})
	hash := sha256.Sum256(headers)
	b.Hash = hash[:]
}

var index = 0
func NewBlock(data string,  prevBlockHash []byte) *Block {
	index += 1
	block := &Block{index,time.Now().Unix(), []byte(data), prevBlockHash, []byte{}}
	block.SetHash()
	return block
}


func (bc *Blockchain) AddBlock(data string) {
	prevBlock := bc.blocks[len(bc.blocks)-1]
	newBlock := NewBlock(data, prevBlock.Hash)
	bc.blocks = append(bc.blocks, newBlock)
}
func NewGenesisBlock() *Block {
	return NewBlock("Genesis Block", []byte{})
}
func NewBlockchain() *Blockchain {
	return &Blockchain{[]*Block{NewGenesisBlock()}}
}

func main() {
	bc := NewBlockchain()

	bc.AddBlock("Làm Quen")
	bc.AddBlock("Về Đích")
	bc.AddBlock("Khởi Động")
	bc.AddBlock("Vượt Chướng Ngại Vật")
	bc.AddBlock("I Love You")

	for _, block := range bc.blocks {
		fmt.Printf("Index: %d\n", block.Index)
		fmt.Printf("Prev. hash: %x\n", block.PrevBlockHash)
		fmt.Printf("Data: %s\n", block.Data)
		fmt.Printf("Hash: %x\n", block.Hash)
		fmt.Println()
	}
}


