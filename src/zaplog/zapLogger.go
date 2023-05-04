package zaplog

import (
	"fmt"
	"os"

	"go.uber.org/zap"
)

var Logger *zap.Logger

func init() {
	var err error
	Logger, err = zap.NewDevelopment()
	if err != nil {
		fmt.Printf("ZapLogger init failed. \n")
		os.Exit(1)
	}
	defer Logger.Sync()
}
