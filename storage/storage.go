package storage

import (
	"context"
	ct "user_service/genproto/genproto/user_service"
)

type StorageI interface {
	CloseDB()
	Teacher() TeacherRepoI
	SupportTeacher() SupportTeacherRepoI
	Branch() BranchRepoI
	Admin() AdminRepoI
}

type TeacherRepoI interface {
	Create(ctx context.Context, req *ct.CreateTeacher) (resp *ct.TeacherPrimaryKey, err error)
	GetByID(ctx context.Context, req *ct.TeacherPrimaryKey) (resp *ct.Teacher, err error)
	GetList(ctx context.Context, req *ct.GetListTeacherRequest) (resp *ct.GetListTeacherResponse, err error)
	Update(ctx context.Context, req *ct.UpdateTeacherRequest) (resp *ct.Message, err error)
	Delete(ctx context.Context, req *ct.TeacherPrimaryKey) (resp *ct.Message, err error)
	GetByGmail(ctx context.Context, req *ct.TeacherGmail) (*ct.TeacherPrimaryKey, error)
}

type SupportTeacherRepoI interface {
	Create(ctx context.Context, req *ct.CreateSupportTeacher) (resp *ct.SupportTeacherPrimaryKey, err error)
	GetByID(ctx context.Context, req *ct.SupportTeacherPrimaryKey) (resp *ct.SupportTeacher, err error)
	GetList(ctx context.Context, req *ct.GetListSupportTeacherRequest) (resp *ct.GetListSupportTeacherResponse, err error)
	Update(ctx context.Context, req *ct.UpdateSupportTeacherRequest) (resp *ct.STMessage, err error)
	Delete(ctx context.Context, req *ct.SupportTeacherPrimaryKey) (resp *ct.STMessage, err error)
	GetByGmail(ctx context.Context, req *ct.SupportTeacherGmail) (*ct.SupportTeacherPrimaryKey, error)
}

type BranchRepoI interface {
	Create(ctx context.Context, req *ct.CreateBranch) (resp *ct.BranchPrimaryKey, err error)
	GetByID(ctx context.Context, req *ct.BranchPrimaryKey) (resp *ct.Branch, err error)
	GetList(ctx context.Context, req *ct.GetListBranchRequest) (resp *ct.GetListBranchResponse, err error)
	Update(ctx context.Context, req *ct.UpdateBranchRequest) (resp *ct.BMessage, err error)
	Delete(ctx context.Context, req *ct.BranchPrimaryKey) (resp *ct.BMessage, err error)
}
	
type AdminRepoI interface {
	Create(ctx context.Context, req *ct.CreateAdmin) (resp *ct.AdminPrimaryKey, err error)
	GetByID(ctx context.Context, req *ct.AdminPrimaryKey) (resp *ct.Admin, err error)
	GetList(ctx context.Context, req *ct.GetListAdminRequest) (resp *ct.GetListAdminResponse, err error)
	Update(ctx context.Context, req *ct.UpdateAdminRequest) (resp *ct.ADMessage, err error)
	Delete(ctx context.Context, req *ct.AdminPrimaryKey) (resp *ct.ADMessage, err error)
	GetByGmail(ctx context.Context, req *ct.AdminGmail) (*ct.AdminPrimaryKey, error)
}