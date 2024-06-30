package service

import (
	"context"
	"user_service/config"
	"user_service/genproto/genproto/user_service"
	"user_service/grpc/client"
	"user_service/storage"

	"github.com/saidamir98/udevs_pkg/logger"
)

type SupportTeacherService struct {
	user_service.UnimplementedSupportTeacherServiceServer
	cfg      config.Config
	log      logger.LoggerI
	strg     storage.StorageI
	services client.ServiceManagerI
}

func NewSupportTeacherService(cfg config.Config, log logger.LoggerI, strg storage.StorageI, srvs client.ServiceManagerI) *SupportTeacherService {
	return &SupportTeacherService{
		cfg:      cfg,
		log:      log,
		strg:     strg,
		services: srvs,
	}
}
func (c *SupportTeacherService) Create(ctx context.Context, req *user_service.CreateSupportTeacher) (*user_service.SupportTeacherPrimaryKey, error) {

	c.log.Info("---CreateSupportTeacher--->>>", logger.Any("req", req))

	resp, err := c.strg.SupportTeacher().Create(ctx, req)
	if err != nil {
		c.log.Error("---CreateSupportTeacher--->>>", logger.Error(err))
		return nil, err
	}

	return resp, nil
}

func (c *SupportTeacherService) GetByID(ctx context.Context, req *user_service.SupportTeacherPrimaryKey) (*user_service.SupportTeacher, error) {
	c.log.Info("---GetByIdSupportTeacher--->>>", logger.Any("req", req))

	resp, err := c.strg.SupportTeacher().GetByID(ctx, req)
	if err != nil {
		c.log.Error("---GetByIdSupportTeacher--->>>", logger.Error(err))
		return nil, err
	}

	return resp, nil
}

func (c *SupportTeacherService) GetList(ctx context.Context, req *user_service.GetListSupportTeacherRequest) (*user_service.GetListSupportTeacherResponse, error) {
	c.log.Info("---GetAllSupportTeacher--->>>", logger.Any("req", req))

	resp, err := c.strg.SupportTeacher().GetList(ctx, req)
	if err != nil {
		c.log.Error("---GetAllSupportTeacher--->>>", logger.Error(err))
		return nil, err
	}

	return resp, nil
}

func (c *SupportTeacherService) Update(ctx context.Context, req *user_service.UpdateSupportTeacherRequest) (*user_service.STMessage, error) {
	c.log.Info("---UpdateSupportTeacher--->>>", logger.Any("req", req))

	resp, err := c.strg.SupportTeacher().Update(ctx, req)
	if err != nil {
		c.log.Error("---UpdateSupportTeacher--->>>", logger.Error(err))
		return nil, err
	}

	return resp, nil
}

func (c *SupportTeacherService) Delete(ctx context.Context, req *user_service.SupportTeacherPrimaryKey) (*user_service.STMessage, error) {
	c.log.Info("---DeleteSupportTeacher--->>>", logger.Any("req", req))

	resp, err := c.strg.SupportTeacher().Delete(ctx, req)
	if err != nil {
		c.log.Error("---DeleteSupportTeacher--->>>", logger.Error(err))
		return nil, err
	}

	return resp, nil
}

func (c *SupportTeacherService) GetByGmail(ctx context.Context, req *user_service.SupportTeacherGmail) (*user_service.SupportTeacherGmailRes, error) {
	c.log.Info("---GetByGmailSupportTeacher--->>>", logger.Any("req", req))

	resp, err := c.strg.SupportTeacher().GetByGmail(ctx, req)
	if err != nil {
		c.log.Error("---GetByGmailSupportTeacher--->>>", logger.Error(err))
		return nil, err
	}

	return resp, nil
}

func (c *SupportTeacherService) SupportTeacherReport(ctx context.Context, req *user_service.GetListSupportTeacherRequest) (*user_service.GetRepSupportTeacherResponse, error) {
	c.log.Info("---SupportTeacherReport--->>>", logger.Any("req", req))

	resp, err := c.strg.SupportTeacher().SupportTeacherReport(ctx, req)
	if err != nil {
		c.log.Error("---SupportTeacherReport--->>>", logger.Error(err))
		return nil, err
	}

	return resp, nil
}