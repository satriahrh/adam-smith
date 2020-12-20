package management

import (
	"context"
	"fmt"

	"github.com/satriahrh/adam-smith/build/proto"
	"github.com/satriahrh/adam-smith/internal/helper"
	"go.uber.org/zap"
)

func (m *management) AddVariation(ctx context.Context, variation *proto.Variation) (err error) {
	errorParseFromDictionary := func(err error) error {
		return helper.ErrorParseFromDictionary(err, []helper.ErrorDictionaryItem{
			{59, "Error 1062", fmt.Errorf("these value is already created")},
		})
	}

	createdVariation, err := m.ent.Variation.Create().
		SetType(variation.GetType()).
		SetValue(variation.GetValue()).
		Save(ctx)
	if err != nil {
		m.logger.Error("Error when creating variation", zap.Any("action_id", ctx.Value("ActionID")), zap.Error(err))
		return errorParseFromDictionary(err)
	}
	variation.Id = createdVariation.ID
	return nil
}
