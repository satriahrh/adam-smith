package catalogue

import (
	"context"

	"github.com/satriahrh/adam-smith/build/proto"
)

func (c *catalog) AddBrand(ctx context.Context, protoBrand *proto.Brand) error {
	brand, err := c.ent.Brand.Create().
		SetCode(protoBrand.Code).
		SetName(protoBrand.Name).
		Save(ctx)
	if err != nil {
		return err
	}

	protoBrand.Id = brand.ID
	return nil
}
