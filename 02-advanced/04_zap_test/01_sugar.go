package main

import "go.uber.org/zap"

func main() {
	// 生产环境
	logger, _ := zap.NewProduction()
	// 开发环境
	//logger, _ := zap.NewDevelopment()
	defer logger.Sync() // flushes buffer, if any
	url := "https://notes.einscat.com"
	sugar := logger.Sugar()
	sugar.Infow("failed to fetch URL",
		// Structured context as loosely typed key-value pairs.
		"url", url,
		"attempt", 3,
		//"backoff", time.Second,
	)
	sugar.Infof("Failed to fetch URL: %s", url)
}

/**
{"level":"info","ts":1642007697.784308,"caller":"01_sugar/orm_test.go:13","msg":"failed to fetch URL","url":"https://notes.einscat.com","attempt":3}
{"level":"info","ts":1642007697.784385,"caller":"01_sugar/orm_test.go:19","msg":"Failed to fetch URL: https://notes.einscat.com"}
*/
