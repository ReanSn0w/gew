package view

import "context"

// Text, Empty, Group
type View interface {
	Body(context.Context) View
}
