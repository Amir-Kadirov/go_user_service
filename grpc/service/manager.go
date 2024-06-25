package service

import (
	"context"
	"user_service/config"
	"user_service/genproto/genproto/user_service"
	"user_service/grpc/client"
	"user_service/storage"

	"github.com/saidamir98/udevs_pkg/logger"
)

type ManagerService struct {
	user_service.UnimplementedManagerServiceServer
	cfg      config.Config
	log      logger.LoggerI
	strg     storage.StorageI
	services client.ServiceManagerI
}

func NewManagerService(cfg config.Config, log logger.LoggerI, strg storage.StorageI, srvs client.ServiceManagerI) *ManagerService {
	return &ManagerService{
		cfg:      cfg,
		log:      log,
		strg:     strg,
		services: srvs,
	}
}
func (c *ManagerService) Create(ctx context.Context, req *user_service.CreateManager) (*user_service.ManagerPrimaryKey, error) {

	c.log.Info("---CreateManager--->>>", logger.Any("req", req))

	resp, err := c.strg.Manager().Create(ctx, req)
	if err != nil {
		c.log.Error("---CreateManager--->>>", logger.Error(err))
		return nil, err
	}

	return resp, nil
}

func (c *ManagerService) GetByID(ctx context.Context, req *user_service.ManagerPrimaryKey) (*user_service.Manager, error) {
	c.log.Info("---GetByIdManager--->>>", logger.Any("req", req))

	resp, err := c.strg.Manager().GetByID(ctx, req)
	if err != nil {
		c.log.Error("---GetByIdManager--->>>", logger.Error(err))
		return nil, err
	}

	return resp, nil
}

func (c *ManagerService) GetList(ctx context.Context, req *user_service.GetListManagerRequest) (*user_service.GetListManagerResponse, error) {
	c.log.Info("---GetAllManager--->>>", logger.Any("req", req))

	resp, err := c.strg.Manager().GetList(ctx, req)
	if err != nil {
		c.log.Error("---GetAllManager--->>>", logger.Error(err))
		return nil, err
	}

	return resp, nil
}

func (c *ManagerService) Update(ctx context.Context, req *user_service.UpdateManagerRequest) (*user_service.MGMessage, error) {
	c.log.Info("---UpdateManager--->>>", logger.Any("req", req))

	resp, err := c.strg.Manager().Update(ctx, req)
	if err != nil {
		c.log.Error("---UpdateManager--->>>", logger.Error(err))
		return nil, err
	}

	return resp, nil
}

func (c *ManagerService) Delete(ctx context.Context, req *user_service.ManagerPrimaryKey) (*user_service.MGMessage, error) {
	c.log.Info("---DeleteManager--->>>", logger.Any("req", req))

	resp, err := c.strg.Manager().Delete(ctx, req)
	if err != nil {
		c.log.Error("---DeleteManager--->>>", logger.Error(err))
		return nil, err
	}

	return resp, nil
}

func (c *ManagerService) GetByGmail(ctx context.Context, req *user_service.ManagerGmail) (*user_service.ManagerPrimaryKey, error) {
	c.log.Info("---GetByGmailManager--->>>", logger.Any("req", req))

	resp, err := c.strg.Manager().GetByGmail(ctx, req)
	if err != nil {
		c.log.Error("---GetByGmailManager--->>>", logger.Error(err))
		return nil, err
	}

	return resp, nil
}
