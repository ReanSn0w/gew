package css

type (
	Style interface {
		Style() []byte
	}

	Pair struct {
		Key   string
		Value string
	}
)

func Rules(selector string, pairs ...Pair) Style {
	r := make(rules)
	for _, pair := range pairs {
		r[pair.Key] = pair.Value
	}

	return r
}

func Group(selector string, blocks ...Style) Style {
	return &group{
		selector: selector,
		block:    blocks,
	}
}
