package css

import "bytes"

type (
	selector struct {
		selector string
		blocks   []Style
	}

	group struct {
		block []Style
	}

	rules map[string]string
)

func (g *selector) Style() []byte {
	buffer := new(bytes.Buffer)
	buffer.WriteString(g.selector + "{")
	for _, block := range g.blocks {
		buffer.Write(block.Style())
	}
	buffer.WriteString("}")
	return buffer.Bytes()
}

func (g *group) Style() []byte {
	buffer := new(bytes.Buffer)
	for _, block := range g.block {
		buffer.Write(block.Style())
	}
	return buffer.Bytes()
}

func (r rules) Style() []byte {
	buffer := new(bytes.Buffer)
	for key, value := range r {
		buffer.WriteString(key + ": " + value + "; ")
	}
	return buffer.Bytes()
}
