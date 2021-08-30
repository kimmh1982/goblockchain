package main

<<<<<<< Updated upstream

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
=======
func main() {
	// bc := NewBlockchain()

	// bc.AddBlock("Send 1 Token to MJ")
	// bc.AddBlock("Send 2 Token to TEMP")

	// for _, block := range bc.blocks {
	// 	fmt.Printf("Prev. Hash: %x\n", block.PrevBlockHash)
	// 	fmt.Printf("Data: %s\n", block.Data)
	// 	fmt.Printf("Hash: %x\n", block.Hash)
	// 	pow := NewProofOfWork(block)
	// 	fmt.Printf("PoW: %s\n", strconv.FormatBool(pow.Validate()))
	// 	fmt.Println()
	// }

	cli := CLI{}
	cli.Run()
>>>>>>> Stashed changes
}