package main

import (
	"fmt"
	"github.com/savaukr/restApiGolang/pkg/config"
	"uber-go.org/zap"
)

func main() {
	prdLogger, _ := zap.NewProduction()
	defer prdLogger.Sync()
	logger := prdLogger.Sugar()

	cfg, err := config.newConfig()
}
