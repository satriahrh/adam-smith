package catalogue

import (
	"context"

	"github.com/satriahrh/adam-smith/build/proto"
	"github.com/satriahrh/adam-smith/ent/schema"
	"go.uber.org/zap"
)

func (c *catalog) AddProduct(ctx context.Context, brandId uint64, product *proto.Product) (err error) {
	createdProduct, err := c.ent.Product.Create().
		SetBrandID(brandId).
		SetName(product.GetName()).
		SetSku(product.GetSku()).
		SetDescriptions(
			normalizeProductDescriptions(product.GetDescriptions()),
		).
		SetImages(
			normalizeProductImages(product.GetImages()),
		).
		SetMarketplaces(
			normalizeProductMarketplace(product.GetMarketplaces()),
		).
		Save(ctx)
	if err != nil {
		c.logger.Error("Error when creating brand", zap.Any("action_id", ctx.Value("ActionID")), zap.Error(err))
		return err
	}
	product.Id = createdProduct.ID

	return err
}

func normalizeProductDescriptions(protoDescriptions []*proto.ProductDescription) []schema.ProductDescription {
	entDescriptions := make([]schema.ProductDescription, 0)
	for _, protoDescription := range protoDescriptions {
		entDescriptions = append(entDescriptions,
			schema.ProductDescription{
				Heading: protoDescription.GetHeading(),
				Body:    protoDescription.GetBody(),
			},
		)
	}
	return entDescriptions
}

func normalizeProductImages(protoImages *proto.ProductImages) schema.ProductImages {
	return schema.ProductImages{
		Thumbnail: protoImages.GetThumbnail(),
		Displays:  protoImages.GetDisplays(),
	}
}

func normalizeProductMarketplace(protoMarketplaces *proto.ProductMarketplaces) schema.ProductMarketplaces {
	return schema.ProductMarketplaces{
		Tokopedia: protoMarketplaces.GetTokopedia(),
		Shopee:    protoMarketplaces.GetShopee(),
		Bukalapak: protoMarketplaces.GetBukalapak(),
	}
}
