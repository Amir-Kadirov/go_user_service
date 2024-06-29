package service

import (
	"context"
	"user_service/config"
	"user_service/genproto/genproto/user_service"
	"user_service/grpc/client"
	"user_service/storage"

	"github.com/saidamir98/udevs_pkg/logger"
)

type TeacherService struct {
	user_service.UnimplementedTeacherServiceServer
	cfg      config.Config
	log      logger.LoggerI
	strg     storage.StorageI
	services client.ServiceManagerI
}

func NewTeacherService(cfg config.Config, log logger.LoggerI, strg storage.StorageI, srvs client.ServiceManagerI) *TeacherService {
	return &TeacherService{
		cfg:      cfg,
		log:      log,
		strg:     strg,
		services: srvs,
	}
}
func (c *TeacherService) Create(ctx context.Context, req *user_service.CreateTeacher) (*user_service.TeacherPrimaryKey, error) {

	c.log.Info("---CreateTeacher--->>>", logger.Any("req", req))

	resp, err := c.strg.Teacher().Create(ctx, req)
	if err != nil {
		c.log.Error("---CreateTeacher--->>>", logger.Error(err))
		return nil, err
	}

	return resp, nil
}

func (c *TeacherService) GetByID(ctx context.Context, req *user_service.TeacherPrimaryKey) (*user_service.Teacher, error) {
	c.log.Info("---GetByIdTeacher--->>>", logger.Any("req", req))

	resp, err := c.strg.Teacher().GetByID(ctx, req)
	if err != nil {
		c.log.Error("---GetByIdTeacher--->>>", logger.Error(err))
		return nil, err
	}

	return resp, nil
}

func (c *TeacherService) GetList(ctx context.Context, req *user_service.GetListTeacherRequest) (*user_service.GetListTeacherResponse, error) {
	c.log.Info("---GetAllTeacher--->>>", logger.Any("req", req))

	resp, err := c.strg.Teacher().GetList(ctx, req)
	if err != nil {
		c.log.Error("---GetAllTeacher--->>>", logger.Error(err))
		return nil, err
	}

	return resp, nil
}

func (c *TeacherService) Update(ctx context.Context, req *user_service.UpdateTeacherRequest) (*user_service.Message, error) {
	c.log.Info("---UpdateTeacher--->>>", logger.Any("req", req))

	resp, err := c.strg.Teacher().Update(ctx, req)
	if err != nil {
		c.log.Error("---UpdateTeacher--->>>", logger.Error(err))
		return nil, err
	}

	return resp, nil
}

func (c *TeacherService) Delete(ctx context.Context, req *user_service.TeacherPrimaryKey) (*user_service.Message, error) {
	c.log.Info("---DeleteTeacher--->>>", logger.Any("req", req))

	resp, err := c.strg.Teacher().Delete(ctx, req)
	if err != nil {
		c.log.Error("---DeleteTeacher--->>>", logger.Error(err))
		return nil, err
	}

	return resp, nil
}

func (c *TeacherService) GetByGmail(ctx context.Context, req *user_service.TeacherGmail) (*user_service.TeacherPrimaryKey, error) {
	c.log.Info("---GetByGmailTeacher--->>>", logger.Any("req", req))

	resp, err := c.strg.Teacher().GetByGmail(ctx, req)
	if err != nil {
		c.log.Error("---GetByGmailTeacher--->>>", logger.Error(err))
		return nil, err
	}

	return resp, nil
}


func (c *TeacherService) TeacherReport(ctx context.Context, req *user_service.GetListTeacherRequest) (*user_service.GetRepTeacherResponse, error) {
	c.log.Info("---TeacherReportTeacher--->>>", logger.Any("req", req))

	resp, err := c.strg.Teacher().TeacherReport(ctx, req)
	if err != nil {
		c.log.Error("---TeacherReportTeacher--->>>", logger.Error(err))
		return nil, err
	}

	return resp, nil
}
