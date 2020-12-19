package catalogue

import (
	"context"

	"github.com/satriahrh/adam-smith/build/ent"
	"github.com/satriahrh/adam-smith/build/proto"
)

type Catalogue interface {
	AddBrand(ctx context.Context, newBrand *proto.Brand) (err error)
}

type catalog struct {
	ent *ent.Client
}

func New(ent *ent.Client) Catalogue {
	return &catalog{ent}
}
