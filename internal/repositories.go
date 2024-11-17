package internal

import "context"

type DocumentValidator interface {
	Validate(ctx context.Context, doc Document) error
}
