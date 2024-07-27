package css

type (
	Style interface {
		Style() []byte
	}

	pair struct {
		Key   string
		Value string
	}
)

func Pair(key string, value string) pair {
	return pair{Key: key, Value: value}
}

func Rules(pairs ...pair) Style {
	r := make(rules)
	for _, pair := range pairs {
		r[pair.Key] = pair.Value
	}

	return r
}

func Selector(s string, blocks ...Style) Style {
	return &selector{
		selector: s,
		blocks:   blocks,
	}
}

func Group(blocks ...Style) Style {
	return &group{
		block: blocks,
	}
}
