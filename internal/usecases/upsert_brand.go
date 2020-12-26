package usecases

import (
	"context"
	"github.com/satriahrh/adam-smith/build/proto"
	"github.com/satriahrh/adam-smith/internal/services/management"
	"go.uber.org/zap"
)

func UpsertBrand(ctx context.Context, managementClient management.Management, protoBrands []*proto.Brand, logger *zap.Logger) {
	for _, protoBrand := range protoBrands {
		if retrievedBrand, err := managementClient.GetBrand(ctx, 0, protoBrand.GetCode()); err != nil {
			if err != management.ErrorBrandNotFound {
				logger.Error("Something wrong when getting brand by code", zap.Any("action_id", ctx.Value("ActionID")), zap.Error(err))
			} else if err = managementClient.AddBrand(ctx, protoBrand); err != nil {
				logger.Error("Something wrong when add brand", zap.Any("action_id", ctx.Value("ActionID")), zap.Error(err))
			}
		} else {
			protoBrand.Id = retrievedBrand.Id
			if err = managementClient.UpdateBrand(ctx, protoBrand); err != nil {
				logger.Error("Something wrong when update brand", zap.Any("action_id", ctx.Value("ActionID")), zap.Error(err))
			}
		}
	}
}
