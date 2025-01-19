package main

import (
	"ApiGateway/internal/clients/grpc"
	"ApiGateway/internal/clients/redis"
	"ApiGateway/internal/clients/users"
	"ApiGateway/internal/httphandler"
	"ApiGateway/internal/service"
	"ApiGateway/pkg/config"
	"ApiGateway/pkg/lib/logger/handler/slogpretty"
	"context"
	"errors"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"
)

func main() {
	cfg := config.InitConfig()

	logg := setupPrettySlogLocal()

	logg.Info("starting api gateway", slog.String("env", cfg.Env))

	redisClient := redis.NewRedisClient(cfg.RedisConfig.Host+":"+strconv.Itoa(cfg.RedisConfig.Port), logg)
	services := service.NewService(logg, redisClient)
	ssoClient, err := grpc.New(
		context.Background(),
		logg,
		cfg.Clients.SSO.Addr,
		cfg.Clients.SSO.Timeout,
		cfg.Clients.SSO.RetriesCount,
	)
	logg.Info("sso client initialized",
		slog.String("addr", cfg.Clients.SSO.Addr),
		slog.String("timeout", cfg.Clients.SSO.Timeout.String()),
		slog.String("retriesCount",
			strconv.Itoa(cfg.Clients.SSO.RetriesCount)),
	)
	if err != nil {
		logg.Error("failed to init sso client", err)
		return
	}

	userSRV := users.New(cfg.Clients.Users.Address+":"+strconv.Itoa(cfg.Clients.Users.Port), logg)
	logg.Info("user client initialized", slog.String("addr", cfg.Clients.Users.Address+strconv.Itoa(cfg.Clients.Users.Port)))

	handlers := httphandler.NewHandler(services, logg, ssoClient, userSRV)

	router := handlers.InitRoutes()

	srv := http.Server{
		Addr:         cfg.Address + ":" + strconv.Itoa(cfg.HTTPServer.Port),
		Handler:      router,
		ReadTimeout:  cfg.HTTPServer.ReadTimeout,
		WriteTimeout: cfg.HTTPServer.WriteTimeout,
		IdleTimeout:  cfg.HTTPServer.IdleTimeout,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			logg.Error("failed to start server", err)
		}
	}()
	logg.Info("server started", slog.String("address", cfg.Address+":"+strconv.Itoa(cfg.HTTPServer.Port)))

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		logg.Error("failed to stop server", err)
	}

	logg.Info("server stopped by graceful shutdown")
}

func setupPrettySlogLocal() *slog.Logger {
	opts := slogpretty.PrettyHandlersOptions{
		SlogOpts: &slog.HandlerOptions{
			Level: slog.LevelDebug,
		},
	}

	handler := opts.NewPrettyHandler(os.Stdout)

	return slog.New(handler)
}

//TODO: fix user_id
//TODO: add method by teacher's id
//TODO: fix kafka in lessons
//TODO: add consumer in Notification service
//TODO: add otros methods in courses service
//TODO: refactor redis for another bussiness logic
