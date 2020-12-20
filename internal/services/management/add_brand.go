package management

import (
	"context"
	"fmt"
	"github.com/satriahrh/adam-smith/internal/helper"

	"github.com/satriahrh/adam-smith/build/proto"
	"go.uber.org/zap"
)

func (m *management) AddBrand(ctx context.Context, brand *proto.Brand) error {
	errorParseFromDictionary := func(err error) error {
		return helper.ErrorParseFromDictionary(err, []helper.ErrorDictionaryItem{
			{From: 55, Value: "Error 1062", Error: fmt.Errorf("brand with given code is already existed")},
		})
	}

	createdBrand, err := m.ent.Brand.Create().
		SetCode(brand.Code).
		SetName(brand.Name).
		Save(ctx)
	if err != nil {
		m.logger.Error("Error when creating brand", zap.Any("action_id", ctx.Value("ActionID")), zap.Error(err))
		return errorParseFromDictionary(err)
	}

	brand.Id = createdBrand.ID
	return nil
}
