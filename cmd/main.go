package main

import (
	"context"
	"errors"
	"github.com/ThalesMonteir0/backend-test/internal/application/part/create"
	"github.com/ThalesMonteir0/backend-test/internal/application/part/delete"
	"github.com/ThalesMonteir0/backend-test/internal/application/part/get"
	"github.com/ThalesMonteir0/backend-test/internal/application/part/update"
	"github.com/ThalesMonteir0/backend-test/internal/application/restock/priorieties"
	"github.com/ThalesMonteir0/backend-test/internal/config"
	"github.com/ThalesMonteir0/backend-test/internal/controller/parts"
	"github.com/ThalesMonteir0/backend-test/internal/controller/restock"
	"github.com/ThalesMonteir0/backend-test/internal/infra/database/postgres"
	"github.com/ThalesMonteir0/backend-test/internal/repository"
	"github.com/ThalesMonteir0/backend-test/internal/service/part"
	"go.uber.org/zap"
	"net/http"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

func main() {
	ctx, stop := signal.NotifyContext(
		context.Background(),
		syscall.SIGINT,
		syscall.SIGTERM,
	)
	defer stop()

	logger, _ := zap.NewProduction()
	defer logger.Sync()

	cfg, err := config.Load()
	if err != nil {
		logger.Error("Error loading config", zap.Error(err))
		return
	}

	app := http.NewServeMux()

	srv := &http.Server{
		Addr:    ":8080",
		Handler: app,
	}

	dbConn, err := postgres.NewConnection(cfg.DatabaseURL())
	if err != nil {
		logger.Error("Error connecting to database", zap.Error(err))
		return
	}

	partRepository := repository.New(dbConn)
	partService := part.New(partRepository)
	partCreateProcessor := create.New(partService, logger)
	partDeleteProcessor := delete.New(partService, logger)
	partUpdateProcessor := update.New(partService, logger)
	partGetProcessor := get.New(partService, logger)
	partCrtlrs := parts.New(
		partCreateProcessor,
		partGetProcessor,
		partDeleteProcessor,
		partUpdateProcessor,
	)

	priorietiesProcessor := priorieties.NewProcessor(partService, logger)
	restockController := restock.New(priorietiesProcessor)

	//PARTS ROUTES
	app.HandleFunc("GET /part", partCrtlrs.GetParts)
	app.HandleFunc("POST /part", partCrtlrs.CreateParts)
	app.HandleFunc("PUT /part/{id}", partCrtlrs.UpdateParts)
	app.HandleFunc("DELETE /part/{id}", partCrtlrs.DeleteParts)

	//RESTOCK
	app.HandleFunc("GET /restock/priorities", restockController.Priorities)

	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()

		if err = srv.ListenAndServe(); !errors.Is(err, http.ErrServerClosed) {
			//todo: logger error
			stop()
		}
	}()
	<-ctx.Done()

	logger.Info("Shutting down server...")

	shutdownCtx, shutdownCancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer shutdownCancel()

	if err = srv.Shutdown(shutdownCtx); err != nil {
		logger.Error("Failed shutting down server", zap.Error(err))
	}

	wg.Wait()

	logger.Info("Server shutdown complete")
}
