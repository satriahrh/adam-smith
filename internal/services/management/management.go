package management

import (
	"context"
	"errors"

	"github.com/satriahrh/adam-smith/build/ent"
	"github.com/satriahrh/adam-smith/build/proto"
)

type Management interface {
	AddBrand(ctx context.Context, brand *proto.Brand) (err error)
	AddProduct(ctx context.Context, brandID uint64, product *proto.Product) (err error)
	AddVariation(ctx context.Context, variation *proto.Variation) (err error)
	GetBrand(ctx context.Context, brandID uint64, code string) (protoBrand *proto.Brand, err error)
	GetProductByID(ctx context.Context, productID uint64) (product *proto.Product, err error)
	UpdateBrand(ctx context.Context, protoBrand *proto.Brand) (err error)
	UpsertSelectedVariantsUnderProduct(ctx context.Context, productID uint64, variants []*proto.Variant) (err error)
}

type management struct {
	ent    *ent.Client
}

func New(ent *ent.Client) Management {
	return &management{ent}
}

var (
	ErrorBrandNotFound = errors.New("brand not found with given id/code")
)
