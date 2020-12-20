package catalogue

import (
	"context"

	"github.com/satriahrh/adam-smith/build/ent"
	"github.com/satriahrh/adam-smith/build/proto"
	"go.uber.org/zap"
)

type Catalogue interface {
	AddBrand(ctx context.Context, brand *proto.Brand) (err error)
	AddProduct(ctx context.Context, brandId uint64, product *proto.Product) (err error)
	AddVariation(ctx context.Context, variation *proto.Variation) (err error)
}

type catalog struct {
	ent    *ent.Client
	logger *zap.Logger
}

func New(ent *ent.Client, logger *zap.Logger) Catalogue {
	return &catalog{ent, logger}
}
