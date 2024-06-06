package service

import (
	"context"
	"fmt"
	"user_service/config"
	"user_service/genproto/user_service"
	"user_service/grpc/client"
	"user_service/storage"

	"github.com/saidamir98/udevs_pkg/logger"
)

type ShopService struct {
	cfg      config.Config
	log      logger.LoggerI
	strg     storage.StorageI
	services client.ServiceManagerI
	*user_service.UnimplementedShopServiceServer
}

func NewShopService(cfg config.Config, log logger.LoggerI, strg storage.StorageI, srvs client.ServiceManagerI) *ShopService {
	return &ShopService{
		cfg:      cfg,
		log:      log,
		strg:     strg,
		services: srvs,
	}
}
func (c *ShopService) Create(ctx context.Context, req *user_service.CreateShop) (resp *user_service.ShopPrimaryKey, err error) {

	c.log.Info("---CreateShop--->>>", logger.Any("req", req))

	resp, err = c.strg.Shop().Create(ctx, req)
	if err != nil {
		c.log.Error("---CreateShop--->>>", logger.Error(err))
		return nil, err
	}

	return resp, nil
}

func (c *ShopService) GetById(ctx context.Context, req *user_service.ShopPrimaryKey) (resp *user_service.GetByID,err error) {
	fmt.Println("herreee")
	c.log.Info("---GetByIdShop--->>>", logger.Any("req", req))

	resp, err = c.strg.Shop().GetById(ctx, req)
	if err != nil {
		c.log.Error("---GetByIdShop--->>>", logger.Error(err))
		return nil, err
	}

	return resp, nil
}

func (c *ShopService) GetList(ctx context.Context, req *user_service.GetListShopRequest) (resp *user_service.GetListShopResponse, err error) {
	c.log.Info("---GetAllShop--->>>", logger.Any("req", req))

	resp, err = c.strg.Shop().GetList(ctx, req)
	if err != nil {
		c.log.Error("---GetAllShop--->>>", logger.Error(err))
		return nil, err
	}

	return resp, nil
}

func (c *ShopService) Update(ctx context.Context,req *user_service.UpdateShopRequest) (resp *user_service.ShopEmpty,err error) {
	c.log.Info("---UpdateShop--->>>", logger.Any("req", req))

	resp, err = c.strg.Shop().Update(ctx, req)
	if err != nil {
		c.log.Error("---UpdateShop--->>>", logger.Error(err))
		return nil, err
	}

	return resp, nil
}

func (c *ShopService) Delete(ctx context.Context,req *user_service.ShopPrimaryKey) (resp *user_service.ShopEmpty,err error) {
	c.log.Info("---DeleteShop--->>>", logger.Any("req", req))

	resp, err = c.strg.Shop().Delete(ctx, req)
	if err != nil {
		c.log.Error("---DeleteShop--->>>", logger.Error(err))
		return nil, err
	}

	return resp, nil
}