package catalogue

import (
	"context"
	"github.com/satriahrh/adam-smith/build/proto"
	"go.uber.org/zap"
)

func (c *catalog) AddVariant(ctx context.Context, productID uint64, variationID uint64, variant *proto.Variant) (err error) {
	createdVariant, err := c.ent.Variant.Create().
		SetImages(variant.GetImages()).
		SetStock(variant.GetStock()).
		SetPrice(variant.GetPrice()).
		SetProductID(productID).
		SetVariationID(variationID).
		Save(ctx)
	if err != nil {
		c.logger.Error("Error when creating variant", zap.Any("action_id", ctx.Value("ActionID")), zap.Error(err))
		return err
	}
	variant.Id = createdVariant.ID
	return nil
}
