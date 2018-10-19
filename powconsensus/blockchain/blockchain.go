package blockchain

import (
	"fmt"
	"go-blackchain/powconsensus/block"
)

type Node struct {
	NextNode *Node
	Block *block.Block
}

func CreateHeaderNode(block *block.Block) *Node{
	var headerNode *Node = new(Node)
	headerNode.NextNode = nil
	headerNode.Block = block
	return headerNode
}

func AddNode(block *block.Block, node *Node) *Node {
	var newNode *Node = new(Node)
	newNode.NextNode = nil
	newNode.Block = block
	node.NextNode = newNode
	return newNode
}

func ShowNodes(node *Node) {
	n := node
	for {
		if n.NextNode == nil {
			fmt.Println(n.Block)
			break
		} else {
			fmt.Println(n.Block)
			n = n.NextNode
		}
	}
}
