package management

import (
	"context"

	"github.com/satriahrh/adam-smith/build/ent"
	"github.com/satriahrh/adam-smith/build/ent/brand"
	"github.com/satriahrh/adam-smith/build/proto"
	"github.com/satriahrh/adam-smith/internal/helper"
	"go.uber.org/zap"
)

// GetBrand get single brand by ID or by Code. Use one of them but not both.
func (m *management) GetBrand(ctx context.Context, brandID uint64, brandCode string) (protoBrand *proto.Brand, err error) {
	defer func() {
		if err != nil {
			err = func(err error) error {
				return helper.ErrorParseFromDictionary(err, []helper.ErrorDictionaryItem{
					{From: 0, Value: "ent: brand not found", Error: ErrorBrandNotFound},
				})
			}(err)
		}
	}()

	var retrievedBrand *ent.Brand
	if brandID != 0 {
		if retrievedBrand, err = m.ent.Brand.Get(ctx, brandID); err != nil {
			m.logger.Error("Error getting brand by ID", zap.Any("action_id", ctx.Value("ActionID")), zap.Error(err))
		}
	} else {
		if retrievedBrand, err = m.ent.Brand.Query().Where(brand.CodeEQ(brandCode)).First(ctx); err != nil {
			m.logger.Error("Error getting brand by Code", zap.Any("action_id", ctx.Value("ActionID")), zap.Error(err))
		}
	}

	if err == nil {
		protoBrand = &proto.Brand{
			Id:   retrievedBrand.ID,
			Code: retrievedBrand.Code,
			Name: retrievedBrand.Name,
		}
	}

	return
}
