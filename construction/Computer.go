package construction

import "fmt"


// Computer 产品类
type Computer struct {
	CPU     string
	Memory  string
	Storage string
	GPU     string
}

func (c *Computer) Show() string {
	return fmt.Sprintf("Computer: CPU=%s, Memory=%s, Storage=%s, GPU=%s",
		c.CPU, c.Memory, c.Storage, c.GPU)
}
