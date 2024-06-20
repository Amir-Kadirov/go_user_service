package grpc

import (
	"user_service/config"
	"user_service/genproto/genproto/user_service"
	"user_service/grpc/client"
	"user_service/grpc/service"
	"user_service/storage"

	"github.com/saidamir98/udevs_pkg/logger"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func SetUpServer(cfg config.Config, log logger.LoggerI, strg storage.StorageI, srvc client.ServiceManagerI) (grpcServer *grpc.Server) {

	grpcServer = grpc.NewServer()

	user_service.RegisterCustomerServiceServer(grpcServer, service.NewCustomerService(cfg, log, strg, srvc))

	user_service.RegisterShopServiceServer(grpcServer, service.NewShopService(cfg, log, strg, srvc))

	user_service.RegisterSellerServiceServer(grpcServer, service.NewSellerService(cfg, log, strg, srvc))

	user_service.RegisterBranchServiceServer(grpcServer, service.NewBranchService(cfg, log, strg, srvc))

	user_service.RegisterSystemUserServiceServer(grpcServer, service.NewSystemUserService(cfg, log, strg, srvc))

	reflection.Register(grpcServer)
	return
}
