package management

import (
	"context"

	"github.com/satriahrh/adam-smith/build/proto"
	"go.uber.org/zap"
)

func (m *management) AddVariation(ctx context.Context, variation *proto.Variation) (err error) {
	createdVariation, err := m.ent.Variation.Create().
		SetType(variation.GetType()).
		SetValue(variation.GetValue()).
		Save(ctx)
	if err != nil {
		m.logger.Error("Error when creating variation", zap.Any("action_id", ctx.Value("ActionID")), zap.Error(err))
		return err
	}
	variation.Id = createdVariation.ID
	return nil
}
