package main

import (
	"go.uber.org/zap"
)

func main() {
	// 2. Logger 性能高
	logger, _ := zap.NewProduction()
	defer logger.Sync() // flushes buffer, if any
	url := "https://notes.einscat.com"
	logger.Info("failed to fetch URL",
		zap.String("url", url),
		zap.Int("nums", 3),
	)
}

/**
{"level":"info","ts":1642007755.473988,"caller":"02_logger/orm_test.go:12","msg":"failed to fetch URL","url":"https://notes.einscat.com","nums":3}
*/
