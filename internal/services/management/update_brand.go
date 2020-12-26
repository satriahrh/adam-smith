package management

import (
	"context"
	"fmt"
	"github.com/satriahrh/adam-smith/internal/helper"
	"github.com/satriahrh/adam-smith/internal/instrumentations"

	"github.com/satriahrh/adam-smith/build/proto"
	"go.uber.org/zap"
)

func (m *management) UpdateBrand(ctx context.Context, protoBrand *proto.Brand) (err error) {
	errorParseFromDictionary := func(err error) error {
		return helper.ErrorParseFromDictionary(err, []helper.ErrorDictionaryItem{
			{From: 0, Value: "ent: brand not found", Error: fmt.Errorf("brand with given id is not found")},
		})
	}

	defer func() {
		if err != nil {
			err = errorParseFromDictionary(err)
		}
	}()

	err = m.ent.Brand.UpdateOneID(protoBrand.GetId()).
		SetName(protoBrand.GetName()).
		Exec(ctx)
	if err != nil {
		instrumentations.Logger.Error("Error when updating brand", zap.Any("action_id", ctx.Value("ActionID")), zap.Error(err))
	}
	return
}
