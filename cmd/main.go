package main

import (
	"context"
	"errors"
	"github.com/ThalesMonteir0/backend-test/internal/application/part/create"
	delete "github.com/ThalesMonteir0/backend-test/internal/application/part/delete"
	"github.com/ThalesMonteir0/backend-test/internal/application/part/get"
	"github.com/ThalesMonteir0/backend-test/internal/application/part/update"
	"github.com/ThalesMonteir0/backend-test/internal/controller/parts"
	"github.com/ThalesMonteir0/backend-test/internal/infra/database/postgres"
	"github.com/ThalesMonteir0/backend-test/internal/repository"
	"github.com/ThalesMonteir0/backend-test/internal/service/part"
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

	//todo: config env
	//todo: config logger

	app := http.NewServeMux()

	srv := &http.Server{
		Addr:    ":8080",
		Handler: app,
	}

	dbConn, err := postgres.NewConnection("")
	if err != nil {
		//todo logs
		return
	}

	partRepository := repository.New(dbConn)
	partService := part.New(partRepository)
	partCreateProcessor := create.New(partService)
	partDeleteProcessor := delete.New(partService)
	partUpdateProcessor := update.New(partService)
	partGetProcessor := get.New(partService)
	partCrtlrs := parts.New(
		partCreateProcessor,
		partGetProcessor,
		partDeleteProcessor,
		partUpdateProcessor,
	)

	//PARTS ROUTES
	app.HandleFunc("GET /part", partCrtlrs.GetParts)
	app.HandleFunc("POST /part", partCrtlrs.CreateParts)
	app.HandleFunc("UPDATE /part", partCrtlrs.UpdateParts)
	app.HandleFunc("DELETE /part", partCrtlrs.DeleteParts)

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

	//todo: logger

	shutdownCtx, shutdownCancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer shutdownCancel()

	if err := srv.Shutdown(shutdownCtx); err != nil {
		//todo: logger graceshutdown failed
	}

	wg.Wait()

	//todo: logger shutdown completed
}
