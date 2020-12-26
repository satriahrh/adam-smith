package sheets

import (
	"context"
	"google.golang.org/api/option"
	"google.golang.org/api/sheets/v4"
)

type Sheets struct {
}

func Get(spreadsheetId, range_ string) (*sheets.ValueRange, error) {
	service, err := sheets.NewService(
		context.TODO(),
		option.WithAPIKey("AIzaSyCWpP7r5cNeqdx8OTq_A6-qOBYCuRBmHbg"),
	)
	if err != nil {
		return nil, err
	}
	return service.Spreadsheets.Values.Get(spreadsheetId, range_).Do()
}
