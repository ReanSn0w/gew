package css

import (
	"bytes"
	"net/http"
)

var (
	defaultStyles = New()
)

func New() *CSS {
	return &CSS{}
}

func Default() *CSS {
	return defaultStyles
}

type CSS struct {
	registredNodes []Style
	compiledStyle  []byte
}

func (s *CSS) Register(nodes ...Style) {
	s.registredNodes = append(s.registredNodes, nodes...)
	s.compile()
}

func (c *CSS) Handler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/css")
		w.WriteHeader(http.StatusOK)
		w.Write(c.compiledStyle)
	})
}

func (c *CSS) Plain() []byte {
	return c.compiledStyle
}

func (c *CSS) compile() {
	buffer := new(bytes.Buffer)
	for _, node := range c.registredNodes {
		buffer.Write(node.Style())
	}
	c.compiledStyle = buffer.Bytes()
}
