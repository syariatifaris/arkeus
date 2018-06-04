package txcontext

import "context"

//ContextOperation contract
type ContextOperation interface {
	//AcceptContext accept context from caller
	AcceptContext(ctx context.Context)
}

//BaseContextOperation structure
type BaseContextOperation struct {
	Context context.Context
}
