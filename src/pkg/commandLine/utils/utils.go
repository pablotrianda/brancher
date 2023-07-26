package utils

import (
	"os"
	"os/signal"
	"syscall"
)

func GracefullShutdown() {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	os.Exit(0)
}
