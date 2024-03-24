package domain

import (
	"context"
	"wishlist/internal/gateway"
)

type App struct {
	Queries *gateway.Queries
	Ctx     context.Context
}
