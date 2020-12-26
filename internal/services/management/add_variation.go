package management

import (
	"context"
	"fmt"
	"github.com/satriahrh/adam-smith/internal/instrumentations"

	"github.com/satriahrh/adam-smith/build/proto"
	"github.com/satriahrh/adam-smith/internal/helper"
	"go.uber.org/zap"
)

func (m *management) AddVariation(ctx context.Context, variation *proto.Variation) (err error) {
	errorParseFromDictionary := func(err error) error {
		return helper.ErrorParseFromDictionary(err, []helper.ErrorDictionaryItem{
			{From: 59, Value: "Error 1062", Error: fmt.Errorf("variation with given values is already existed")},
		})
	}

	createdVariation, err := m.ent.Variation.Create().
		SetType(variation.GetType()).
		SetValue(variation.GetValue()).
		Save(ctx)
	if err != nil {
		instrumentations.Logger.Error("Error when creating variation", zap.Any("action_id", ctx.Value("ActionID")), zap.Error(err))
		return errorParseFromDictionary(err)
	}
	variation.Id = createdVariation.ID
	return nil
}