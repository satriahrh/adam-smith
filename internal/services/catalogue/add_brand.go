package catalogue

import (
	"context"

	"github.com/satriahrh/adam-smith/build/proto"
)

func (c *catalog) AddBrand(ctx context.Context, brand *proto.Brand) error {
	createdBrand, err := c.ent.Brand.Create().
		SetCode(brand.Code).
		SetName(brand.Name).
		Save(ctx)
	if err != nil {
		return err
	}

	brand.Id = createdBrand.ID
	return nil
}
