package blockchain

type BlockChain struct {
	Blocks        map[string]*Block
	LastBlockHash []byte
}

func (bc *BlockChain) AddBlock(data string) {
	prevBlock := bc.Blocks[string(bc.LastBlockHash)]
	newBlock := NewBlock(data, prevBlock.Hash)
	bc.Blocks[string(newBlock.Hash)] = newBlock
	bc.LastBlockHash = newBlock.Hash
}

func NewGenesisBlock() *Block {
	return NewBlock("Genesis Block", []byte{})
}

func NewBlockChain() *BlockChain {
	genesis := NewGenesisBlock()
	blocks := make(map[string]*Block)
	blocks[string(genesis.Hash)] = genesis
	return &BlockChain{Blocks: blocks, LastBlockHash: genesis.Hash}
}
