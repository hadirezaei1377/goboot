package zap

import (
	"go.uber.org/zap"
)

func main() {

	logger, _ := zap.NewProduction()
	defer logger.Sync()

	logger.Info("info:", zap.String("key", "value"))
	logger.Warn("")
	logger.Error("")

	logger.Info("مشکلی رخ داده است",
		zap.String("موقعیت", "Server1"),
		zap.Int("حالت", 200),
	)
}
