package main


//引入区块链
type BlockChain struct {
	Blocks []*Block//区块数组
}

//创建区块链
func NewBlockChain() *BlockChain {
	genesisBlock := NewGenesisBlock(data,[]byte{})//构建创世区块
	blockChain := BlockChain{//添加创世区块
		[]*Block{genesisBlock},
	}
	return &blockChain
}

//添加区块
func (bc *BlockChain)AddBlock(data string)  {
	//获取最后一个区块
	lastBlock := bc.Blocks[len(bc.Blocks)-1]
	//获取最后一个区块的哈希，作为新区块的prevhash
	prevHash := lastBlock.Hash
	block := NewBlock(data,prevHash)
	bc.Blocks = append(bc.Blocks,block)
}
