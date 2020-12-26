package instrumentations

import "go.uber.org/zap"

var Logger = func() *zap.Logger {
	l, _ := zap.NewDevelopment()
	return l
}()
