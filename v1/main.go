package main

import "fmt"

const data  = "go语言太牛逼了"

func main() {
	blockChain := NewBlockChain()
	blockChain.AddBlock("hello")
	blockChain.AddBlock("world")

	for index, block := range blockChain.Blocks {
		fmt.Println(" ============== current block index :", index,"===============")
		fmt.Printf("Version : %d\n", block.Version)
		fmt.Printf("PrevBlockHash : %x\n", block.PrevHash)
		fmt.Printf("Hash : %x\n", block.Hash)
		fmt.Printf("MerkleRoot : %x\n", block.MerKleRoot)
		fmt.Printf("TimeStamp : %d\n", block.TimeStamp)
		fmt.Printf("Difficuty : %d\n", block.Difficulty)
		fmt.Printf("Nonce : %d\n", block.Nonce)
		fmt.Printf("Data : %s\n", block.Data)
	}
}