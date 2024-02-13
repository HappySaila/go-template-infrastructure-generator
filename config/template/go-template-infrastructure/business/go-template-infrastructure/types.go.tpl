package [[.Snake]]

import (
	"context"
	"dev/[[.Dashed]].git/business/i"
)

// Constants
const ()

// Service encapsulates core [[.Dashed]] functionality
type Service struct {
	Log   i.Logger
	Store Store
}

// Store encapsulates third-party dependencies
type Store interface {
	Create(ctx context.Context) error
}
