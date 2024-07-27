package css

import "bytes"

type (
	group struct {
		selector string
		block    []Style
	}

	rules map[string]string
)

func (g *group) Style() []byte {
	buffer := new(bytes.Buffer)
	buffer.WriteString(g.selector + "{")
	for _, block := range g.block {
		buffer.Write(block.Style())
	}
	buffer.WriteString("}")
	return buffer.Bytes()
}

func (r rules) Style() []byte {
	buffer := new(bytes.Buffer)
	for key, value := range r {
		buffer.WriteString(key + ": " + value + "; ")
	}
	return buffer.Bytes()
}
