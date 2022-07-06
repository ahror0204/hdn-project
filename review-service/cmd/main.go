package main

import (
	"net"
    "google.golang.org/grpc/reflection"

	"github.com/hdn-project/review-service/config"
	pb "github.com/hdn-project/review-service/genproto"
	"github.com/hdn-project/review-service/pkg/db"
	"github.com/hdn-project/review-service/pkg/logger"
	"github.com/hdn-project/review-service/service"
	"google.golang.org/grpc"
    grpcClient "github.com/hdn-project/review-service/service/grpc_client"
)

func main() {
    cfg := config.Load()

    log := logger.New(cfg.LogLevel, "review-service")
    defer logger.Cleanup(log)

    log.Info("main: sqlxConfig",
        logger.String("host", cfg.PostgresHost),
        logger.Int("port", cfg.PostgresPort),
        logger.String("database", cfg.PostgresDatabase))

    connDB, err := db.ConnectToDB(cfg)
    if err != nil {
        log.Fatal("sqlx connection to postgres error", logger.Error(err))
    }
    grpcC, err := grpcClient.New(cfg)
	if err != nil {
		log.Fatal("grpc client error", logger.Error(err))	
		return
	}
    reviewService := service.NewReviewService(connDB, log,grpcC)

    lis, err := net.Listen("tcp", cfg.RPCPort)
    if err != nil {
        log.Fatal("Error while listening: %v", logger.Error(err))
    }

    s := grpc.NewServer()
    pb.RegisterReviewServiceServer(s, reviewService)
    log.Info("main: server running",
        logger.String("port", cfg.RPCPort))
    reflection.Register(s)
    if err := s.Serve(lis); err != nil {
        log.Fatal("Error while listening: %v", logger.Error(err))
    }
}
