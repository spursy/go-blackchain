package main

import (
	"go-blackchain/powconsensus/block"
	"go-blackchain/powconsensus/blockchain"
)

func main()  {
	var first = block.GenerateFirstBlock("创世区块")
	var second = block.GenerateNextBlock("第二个区块", first)

	var header = blockchain.CreateHeaderNode(&first)
	blockchain.AddNode(&second, header)

	blockchain.ShowNodes(header)
}