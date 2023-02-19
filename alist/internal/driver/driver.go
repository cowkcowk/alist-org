package driver

import (
	"context"

	"github.com/alist-org/alist/v3/internal/model"
)

type Driver interface {
	Meta
}

type Meta interface {
	Config() Config
	GetStorage() *model.Storage
	SetStorage(model.Storage)
	GetAddition() Additional
	Init(ctx context.Context) error
	Drop(ctx context.Context) error
}