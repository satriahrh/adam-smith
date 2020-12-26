package management

import (
	"context"
	"errors"

	"github.com/satriahrh/adam-smith/build/ent"
	"github.com/satriahrh/adam-smith/build/proto"
	"go.uber.org/zap"
)

type Management interface {
	AddBrand(ctx context.Context, brand *proto.Brand) (err error)
	AddProduct(ctx context.Context, brandID uint64, product *proto.Product) (err error)
	AddVariant(ctx context.Context, productID uint64, variationID uint64, variant *proto.Variant) (err error)
	AddVariation(ctx context.Context, variation *proto.Variation) (err error)
	GetBrand(ctx context.Context, brandID uint64, code string) (protoBrand *proto.Brand, err error)
	UpdateBrand(ctx context.Context, protoBrand *proto.Brand) (err error)
}

type management struct {
	ent    *ent.Client
	logger *zap.Logger
}

func New(ent *ent.Client, logger *zap.Logger) Management {
	return &management{ent, logger}
}

var (
	ErrorBrandNotFound = errors.New("brand not found with given id/code")
)
