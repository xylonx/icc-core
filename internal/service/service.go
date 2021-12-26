package service

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"time"

	"github.com/xylonx/icc-core/internal/config"
	"github.com/xylonx/icc-core/internal/router"
	"github.com/xylonx/zapx"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

var (
	httpServer *http.Server
	grpcServer *grpc.Server
)

func StartService() error {
	httpAddr := fmt.Sprintf("%s:%d", config.Config.Application.HttpHost, config.Config.Application.HttpPort)
	grpcAddr := fmt.Sprintf("%s:%d", config.Config.Application.GrpcHost, config.Config.Application.GrpcPort)

	grpcLis, err := net.Listen("tcp", grpcAddr)
	if err != nil {
		zapx.Error("listen grpc addr failed", zap.String("addr", grpcAddr))
		return err
	}

	httpServer = router.InitHttpServer(&router.HttpOption{
		Addr:         httpAddr,
		ReadTimeout:  time.Duration(config.Config.Application.HttpReadTimeoutSeconds) * time.Second,
		WriteTimeout: time.Duration(config.Config.Application.HttpWriteTimeoutSeconds) * time.Second,
		AllowOrigins: config.Config.Application.HttpAllowOrigins,
	})
	grpcServer = router.InitRPCServer()

	go func() {
		zapx.Info("start http server", zap.String("host", httpAddr))
		if err := httpServer.ListenAndServe(); err != nil {
			zapx.Error("http run error", zap.Error(err))
		}
	}()

	go func() {
		zapx.Info("start grpc server", zap.String("addr", grpcAddr))
		if err := grpcServer.Serve(grpcLis); err != nil {
			zapx.Error("serve grpc listen failed", zap.Error(err))
		}
	}()

	return nil
}

func StopService(ctx context.Context) {
	grpcServer.GracefulStop()
	httpServer.Shutdown(ctx)
}
