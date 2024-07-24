package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/getsentry/sentry-go"
	"github.com/tahacodes/boilerplate-go/configs"
	"github.com/tahacodes/boilerplate-go/internal/platform/application"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func init() {
	// Initialize Sentry
	err := sentry.Init(sentry.ClientOptions{
		Dsn: configs.C.Log.SentryDSN,
	})
	if err != nil {
		log.Fatal(fmt.Errorf("failed to initiate sentry"))
	}

	// Initialize zap logger
	level := zap.NewAtomicLevel()
	err = level.UnmarshalText([]byte(configs.C.Log.LogLevel))
	if err != nil {
		log.Fatal(fmt.Errorf("failed to set log level"))
	}

	encoder := zap.NewProductionEncoderConfig()
	encoder.EncodeTime = zapcore.ISO8601TimeEncoder
	encoder.TimeKey = "time"
	encoder.MessageKey = "message"

	zapConfig := zap.Config{
		Level:             level,
		Development:       false,
		Encoding:          "json",
		EncoderConfig:     encoder,
		OutputPaths:       []string{"stdout"},
		ErrorOutputPaths:  []string{"stderr"},
		DisableCaller:     true,
		DisableStacktrace: true,
	}

	logger, err := zapConfig.Build(zap.Hooks(func(entry zapcore.Entry) error {
		// Send errors to a tracking system like sentry
		return nil
	}))
	if err != nil {
		log.Fatal(fmt.Errorf("failed to set zap config"))
	}

	zap.ReplaceGlobals(logger)
}

func main() {
	var err error

	ctx, cancel := context.WithCancel(context.Background())
	wg := &sync.WaitGroup{}

	// Register application
	a, err := application.NewApplication(ctx)
	if err != nil {
		zap.L().Fatal("failed to register application", zap.Error(err))
	}

	// Start async services
	a.RunAwesome(ctx, wg)

	// Watch for termination or close signals
	closeSignal := make(chan os.Signal, 1)
	signal.Notify(closeSignal, syscall.SIGTERM, syscall.SIGINT, os.Interrupt)
	select {
	case <-closeSignal:
		zap.L().Info("terminating by os signal")
	case <-ctx.Done():
		zap.L().Info("terminating by context cancellation")
	}

	cancel()
	wg.Wait()

	zap.L().Debug("shutting down gracefully")
	if err := a.Close(); err != nil {
		log.Fatal(err)
	}
}
