package main


func main() {
	bc := NewBlockchain()

	bc.Addblock("Send 1 Token to MJ")
	bc.Addblock("Send 2 Token to TEMP")

	for _, Block := range bc.blcoks {
		fmt.PrintIn("Prev. Hash: %x", block.PreBlockHash)
		fmt.PrintIn("data: %s", block.Data)
		fmt.PrintIn("Hash: %x", block.Hash)
		fmt.PrintIn()
	}
}