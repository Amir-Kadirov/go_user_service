package service

import (
	"context"
	"user_service/config"
	"user_service/genproto/genproto/user_service"
	"user_service/grpc/client"
	"user_service/storage"

	"github.com/saidamir98/udevs_pkg/logger"
)

type AdminService struct {
	user_service.UnimplementedAdminServiceServer
	cfg      config.Config
	log      logger.LoggerI
	strg     storage.StorageI
	services client.ServiceManagerI
}

func NewAdminService(cfg config.Config, log logger.LoggerI, strg storage.StorageI, srvs client.ServiceManagerI) *AdminService {
	return &AdminService{
		cfg:      cfg,
		log:      log,
		strg:     strg,
		services: srvs,
	}
}
func (c *AdminService) Create(ctx context.Context, req *user_service.CreateAdmin) (*user_service.AdminPrimaryKey, error) {

	c.log.Info("---CreateAdmin--->>>", logger.Any("req", req))

	resp, err := c.strg.Admin().Create(ctx, req)
	if err != nil {
		c.log.Error("---CreateAdmin--->>>", logger.Error(err))
		return nil, err
	}

	return resp, nil
}

func (c *AdminService) GetByID(ctx context.Context, req *user_service.AdminPrimaryKey) (*user_service.Admin, error) {
	c.log.Info("---GetByIdAdmin--->>>", logger.Any("req", req))

	resp, err := c.strg.Admin().GetByID(ctx, req)
	if err != nil {
		c.log.Error("---GetByIdAdmin--->>>", logger.Error(err))
		return nil, err
	}

	return resp, nil
}

func (c *AdminService) GetList(ctx context.Context, req *user_service.GetListAdminRequest) (*user_service.GetListAdminResponse, error) {
	c.log.Info("---GetAllAdmin--->>>", logger.Any("req", req))

	resp, err := c.strg.Admin().GetList(ctx, req)
	if err != nil {
		c.log.Error("---GetAllAdmin--->>>", logger.Error(err))
		return nil, err
	}

	return resp, nil
}

func (c *AdminService) Update(ctx context.Context, req *user_service.UpdateAdminRequest) (*user_service.ADMessage, error) {
	c.log.Info("---UpdateAdmin--->>>", logger.Any("req", req))

	resp, err := c.strg.Admin().Update(ctx, req)
	if err != nil {
		c.log.Error("---UpdateAdmin--->>>", logger.Error(err))
		return nil, err
	}

	return resp, nil
}

func (c *AdminService) Delete(ctx context.Context, req *user_service.AdminPrimaryKey) (*user_service.ADMessage, error) {
	c.log.Info("---DeleteAdmin--->>>", logger.Any("req", req))

	resp, err := c.strg.Admin().Delete(ctx, req)
	if err != nil {
		c.log.Error("---DeleteAdmin--->>>", logger.Error(err))
		return nil, err
	}

	return resp, nil
}

func (c *AdminService) GetByGmail(ctx context.Context, req *user_service.AdminGmail) (*user_service.AdminPrimaryKey, error) {
	c.log.Info("---GetByGmailAdmin--->>>", logger.Any("req", req))

	resp, err := c.strg.Admin().GetByGmail(ctx, req)
	if err != nil {
		c.log.Error("---GetByGmailAdmin--->>>", logger.Error(err))
		return nil, err
	}

	return resp, nil
}
