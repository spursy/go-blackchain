package block

import (
	"strconv"
	"crypto/sha256"
	"encoding/hex"
	"time"
	"strings"
	"fmt"
)

type Block struct {
	Index int
	TimeStamp string
	Diff int
	PreHash string
	Nonce int
	Data string
	HashCode string
}

func GenerateBlockHashValue(block Block) string {
	var hashData = strconv.Itoa(block.Index) + block.TimeStamp + block.Data + strconv.Itoa(block.Diff) + strconv.Itoa(block.Nonce)
	var hash = sha256.New()
	hash.Write([]byte(hashData))
	return hex.EncodeToString(hash.Sum(nil))
}

func GenerateFirstBlock(data string) Block {
	var firstBlock Block
	firstBlock.Index = 1
	firstBlock.TimeStamp = time.Now().String()
	firstBlock.Diff = 4
	firstBlock.Nonce = 0
	firstBlock.Data = data
	firstBlock.HashCode = GenerateBlockHashValue(firstBlock)

	return firstBlock
}

func GenerateNextBlock(data string, oldBlock Block) Block {
	var newBlock Block
	newBlock.Data = data
	newBlock.TimeStamp = time.Now().String()

	newBlock.Nonce = 0
	newBlock.PreHash = oldBlock.HashCode

	newBlock.Index = oldBlock.Index + 1
	newBlock.Diff = oldBlock.Diff

	Mine(newBlock.Diff, &newBlock)
	return newBlock
}


func Mine(diff int, block *Block) {
	for {
		var hash = GenerateBlockHashValue(*block)
		if strings.HasPrefix(hash, strings.Repeat("0", diff)) {
			fmt.Println(hash)
			fmt.Println(block.Nonce)
			fmt.Println("挖矿成功")
			block.HashCode = hash
			return
		} else {
			fmt.Println(hash)
			block.Nonce ++
		}
	}
}