package main

import (
	"context"
	"os"
	"strconv"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/satriahrh/adam-smith/build/ent"
	"github.com/satriahrh/adam-smith/build/proto"
	"github.com/satriahrh/adam-smith/internal/infrastructures/sheets"
	"github.com/satriahrh/adam-smith/internal/services/management"
	"github.com/satriahrh/adam-smith/internal/usecases"
	"go.uber.org/zap"
)

func main() {
	logger, _ := zap.NewDevelopment()
	ctx := context.WithValue(context.TODO(), "ActionID", strconv.Itoa(int(time.Now().Unix())))

	client, err := ent.Open("mysql", "root:rootpw@/adam_smith?parseTime=true")
	if err != nil {
		logger.Fatal("Error opening ent client.", zap.Error(err))
	}
	defer client.Close()
	usecases.UpsertBrand(ctx, managementClient(ctx, client, logger), protoBrands(logger), logger)
}

func managementClient(ctx context.Context, client *ent.Client, logger *zap.Logger) management.Management {
	//Run the auto migration tool.
	if err := client.Schema.Create(ctx); err != nil {
		logger.Fatal("Error running auto migration tool.", zap.Error(err))
	}

	return management.New(client, logger)
}

func protoBrands(logger *zap.Logger) []*proto.Brand {
	args := os.Args
	if len(args) != 2 {
		logger.Fatal("No spreadsheet ID given")
	}
	spreadsheetId := os.Args[1]
	valueRange, err := sheets.Get(spreadsheetId, "brands")
	if err != nil {
		logger.Error("Error on getting spreadsheet", zap.Error(err))
	}

	protoBrands := make([]*proto.Brand, 0)
	for _, value := range valueRange.Values[1:] {
		protoBrands = append(protoBrands, &proto.Brand{
			Code: value[0].(string),
			Name: value[1].(string),
		})
	}
	return protoBrands
}
