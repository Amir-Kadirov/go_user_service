package service

import (
	"context"
	"user_service/config"
	"user_service/genproto/user_service"
	"user_service/grpc/client"
	"user_service/storage"

	"github.com/saidamir98/udevs_pkg/logger"
)

type SellerService struct {
	cfg      config.Config
	log      logger.LoggerI
	strg     storage.StorageI
	services client.ServiceManagerI
	*user_service.UnimplementedSellerServiceServer
}

func NewSellerService(cfg config.Config, log logger.LoggerI, strg storage.StorageI, srvs client.ServiceManagerI) *SellerService {
	return &SellerService{
		cfg:      cfg,	
		log:      log,
		strg:     strg,
		services: srvs,
	}
}
func (c *SellerService) Create(ctx context.Context, req *user_service.CreateSeller) (resp *user_service.SellerPrimaryKey, err error) {

	c.log.Info("---CreateSeller--->>>", logger.Any("req", req))

	resp, err = c.strg.Seller().Create(ctx, req)
	if err != nil {
		c.log.Error("---CreateSeller--->>>", logger.Error(err))
		return nil, err
	}

	return resp, nil
}	

func (c *SellerService) GetByID(ctx context.Context, req *user_service.SellerPrimaryKey) (resp *user_service.Seller, err error) {
	c.log.Info("---GetByIdSeller--->>>", logger.Any("req", req))

	resp, err = c.strg.Seller().GetByID(ctx, req)
	if err != nil {
		c.log.Error("---GetByIdSeller--->>>", logger.Error(err))
		return nil, err
	}

	return resp, nil
}

func (c *SellerService) GetList(ctx context.Context, req *user_service.GetListSellerRequest) (resp *user_service.GetListSellerResponse, err error) {
	c.log.Info("---GetAllSeller--->>>", logger.Any("req", req))

	resp, err = c.strg.Seller().GetList(ctx, req)
	if err != nil {
		c.log.Error("---GetAllSeller--->>>", logger.Error(err))
		return nil, err
	}

	return resp, nil
}

func (c *SellerService) Update(ctx context.Context,req *user_service.UpdateSellerRequest) (resp *user_service.UpdateSellerResponse,err error) {
	c.log.Info("---UpdateSeller--->>>", logger.Any("req", req))

	resp, err = c.strg.Seller().Update(ctx, req)
	if err != nil {
		c.log.Error("---UpdateSeller--->>>", logger.Error(err))
		return nil, err
	}

	return resp, nil
}

func (c *SellerService) Delete(ctx context.Context,req *user_service.SellerPrimaryKey) (resp *user_service.SellerEmpty,err error) {
	c.log.Info("---DeleteSeller--->>>", logger.Any("req", req))

	resp, err = c.strg.Seller().Delete(ctx, req)
	if err != nil {
		c.log.Error("---DeleteSeller--->>>", logger.Error(err))
		return nil, err
	}

	return resp, nil
}