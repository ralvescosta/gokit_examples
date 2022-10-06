package services

import "context"

type (
	BookRepository interface {
		Create(ctx context.Context)
		Get(ctx context.Context)
		List(ctx context.Context)
		Update(ctx context.Context)
		Delete(ctx context.Context)
	}
)
