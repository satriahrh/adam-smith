package catalogue

import (
	"context"

	"github.com/satriahrh/adam-smith/build/proto"
	"go.uber.org/zap"
)

func (c *catalog) AddBrand(ctx context.Context, brand *proto.Brand) error {
	createdBrand, err := c.ent.Brand.Create().
		SetCode(brand.Code).
		SetName(brand.Name).
		Save(ctx)
	if err != nil {
		c.logger.Error("Error when creating brand", zap.Any("action_id", ctx.Value("ActionID")), zap.Error(err))
		return err
	}

	brand.Id = createdBrand.ID
	return nil
}
