package postgres

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	ct "user_service/genproto/genproto/user_service"
	"user_service/pkg/hash"
	"user_service/pkg/helper"
	"user_service/storage"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v4/pgxpool"
)

type SupportTeacherRepo struct {
	db *pgxpool.Pool
}

func NewSupportTeacherRepo(db *pgxpool.Pool) storage.SupportTeacherRepoI {
	return &SupportTeacherRepo{
		db: db,
	}
}

func (c *SupportTeacherRepo) Create(ctx context.Context, req *ct.CreateSupportTeacher) (*ct.SupportTeacherPrimaryKey, error) {
	id := uuid.NewString()
	resp := &ct.SupportTeacherPrimaryKey{Id: id}

	login,err:=helper.GenerateLoginID(c.db,"SupportTeacher")
	if err != nil {
		return nil, err
	}

	query := `INSERT INTO "SupportTeacher" (
			"ID",
			"FullName",
			"Phone",
			"Password",
			"Email",
			"Salary",
			"IeltsScore",
			"IeltsAttemptsCount",
			"BranchID",
			"created_at",
			"LoginID") VALUES (
				$1,
				$2,
				$3,
				$4,
				$5,
				$6,
				$7,
				$8,
				$9,
				NOW(),
				$10
			)`
		hashedPassword,err:=hash.HashPassword(req.Password)
		if err != nil {
			return nil, err
		}

	_, err = c.db.Exec(ctx, query, id,req.Fullname, req.Phone, hashedPassword, req.Email, 
		req.Salary, req.Ieltsscore,req.Ieltsattemptscount,req.Branchid,login)
	if err != nil {
		log.Println("error while creating supportteacher")
		return nil, err
	}

	return resp, err
}

func (c *SupportTeacherRepo) GetByID(ctx context.Context, req *ct.SupportTeacherPrimaryKey) (*ct.SupportTeacher, error) {
	resp := &ct.SupportTeacher{}
	query := `SELECT "ID",
					"FullName",
					"Phone",
					"Email",
					"Salary",
					"IeltsScore",
					"IeltsAttemptsCount",
					"BranchID",
					"created_at",
					"updated_at"
			FROM "SupportTeacher"
			WHERE "ID"=$1 AND "deleted_at" is null`

	row := c.db.QueryRow(ctx, query, req.Id)

	var updatedAt, createdAt sql.NullTime
	if err := row.Scan(
		&resp.Id,
		&resp.Fullname,
		&resp.Phone,
		&resp.Email,
		&resp.Salary,
		&resp.Ieltsscore,
		&resp.Ieltsattemptscount,
		&resp.Branchid,
		&createdAt,
		&updatedAt); err != nil {
		return nil, err
	}

	resp.CreatedAt = helper.NullTimeStampToString(createdAt)
	resp.UpdatedAt = helper.NullTimeStampToString(updatedAt)

	return resp, nil
}

func (c *SupportTeacherRepo) GetList(ctx context.Context, req *ct.GetListSupportTeacherRequest) (*ct.GetListSupportTeacherResponse, error) {
	resp := &ct.GetListSupportTeacherResponse{}
	if req.Offset==0 {
		req.Offset=1
	}

	filter := ""
	offset := (req.Offset - 1) * req.Limit

	if req.Search != "" {
		filter = ` AND "IeltsScore" ILIKE '%` + req.Search + `%' `
	}

	query := `SELECT 
					"ID",
					"FullName",
					"Phone",
					"Email",
					"Salary",
					"IeltsScore",
					"IeltsAttemptsCount",
					"BranchID",
					"created_at",
					"updated_at"
			FROM "SupportTeacher"
        	WHERE "deleted_at" is null AND TRUE ` + filter + `
           OFFSET $1 LIMIT $2
    `

	rows, err := c.db.Query(ctx, query, offset, req.Limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		SupportTeacher := &ct.SupportTeacher{}
		var createdAt, updatedAt sql.NullTime
		if err := rows.Scan(
			&SupportTeacher.Id,
			&SupportTeacher.Fullname,
			&SupportTeacher.Phone,
			&SupportTeacher.Email,
			&SupportTeacher.Salary,
			&SupportTeacher.Ieltsscore,
			&SupportTeacher.Ieltsattemptscount,
			&SupportTeacher.Branchid,
			&createdAt,
			&updatedAt); err != nil {
			return nil, err
		}

		SupportTeacher.CreatedAt = helper.NullTimeStampToString(createdAt)
		SupportTeacher.UpdatedAt = helper.NullTimeStampToString(updatedAt)
		resp.Supportteacher = append(resp.Supportteacher, SupportTeacher)
	}

	queryCount := `SELECT COUNT(*) FROM "SupportTeacher" WHERE "deleted_at" is null AND TRUE ` + filter
	err = c.db.QueryRow(ctx, queryCount).Scan(&resp.Count)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (c *SupportTeacherRepo) Update(ctx context.Context, req *ct.UpdateSupportTeacherRequest) (*ct.STMessage, error) {
	resp := &ct.STMessage{Message: "SupportTeacher updated successfully"}
	query := `UPDATE "SupportTeacher" SET
								 "FullName"=$1,
								 "Phone"=$2,
								 "Email"=$3,
								 "Salary"=$4,
								 "IeltsScore"=$5,
								 "IeltsAttemptsCount"=$6,
								 "BranchID"=$7,
								 "updated_at"=NOW()
								 WHERE "ID"=$8 AND "deleted_at" is null`
	_, err := c.db.Exec(ctx, query, req.Fullname, req.Phone, req.Email, req.Salary, req.Ieltsscore,
									req.Ieltsattemptscount,req.Branchid,req.Id)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (c *SupportTeacherRepo) Delete(ctx context.Context, req *ct.SupportTeacherPrimaryKey) (*ct.STMessage, error) {
	resp := &ct.STMessage{Message: "Support teacher deleted successfully"}
	query := `UPDATE "SupportTeacher" SET
							 "deleted_at"=NOW()
							 WHERE "ID"=$1 AND "deleted_at" is null RETURNING "created_at"`

	var createdAt sql.NullTime
	err := c.db.QueryRow(ctx, query, req.Id).Scan(&createdAt)
	if err != nil {
		return nil, err
	}

	if err = helper.DeleteChecker(createdAt); err != nil {
		return resp, nil
	}

	return resp, nil
}

func (c *SupportTeacherRepo) GetByGmail(ctx context.Context, req *ct.SupportTeacherGmail) (*ct.SupportTeacherGmailRes, error) {
	resp:=&ct.SupportTeacherGmailRes{}
	query := `SELECT "ID","Password" FROM "SupportTeacher" WHERE "Email"=$1`
	err := c.db.QueryRow(ctx, query, req.Gmail).Scan(&resp.Gmail,&resp.Password)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (c *SupportTeacherRepo) SupportTeacherReport(ctx context.Context, req *ct.GetListSupportTeacherRequest) (*ct.GetRepSupportTeacherResponse, error) {
	resp := &ct.GetRepSupportTeacherResponse{}

	if req.Offset==0 {
		req.Offset=1
	}

	filter := ""
	offset := (req.Offset - 1) * req.Limit

	if req.Search != "" {
		filter = ` AND "BranchID" ILIKE '%` + req.Search + `%' `
	}

	query := `SELECT
					"ID",
					"FullName",
					"Phone",
					"Salary",
					"IeltsScore",
					"BranchID",
					"created_at",
					"updated_at"
			FROM "SupportTeacher"
        	WHERE "deleted_at" is null AND TRUE ` + filter + `
           OFFSET $1 LIMIT $2
    `	

	rows, err := c.db.Query(ctx, query, offset, req.Limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		SupportTeacher := &ct.SupportTeacherRep{}
		var createdAt, updatedAt sql.NullTime
		if err := rows.Scan(
			&SupportTeacher.Id,
			&SupportTeacher.Fullname,
			&SupportTeacher.Phone,
			&SupportTeacher.Salary,
			&SupportTeacher.Ieltsscore,
			&SupportTeacher.BranchId,
			&createdAt,
			&updatedAt); err != nil {
			return nil, err
		}

		SupportTeacher.MonhtWorked=int32(helper.DateSince(createdAt))
		fmt.Println("Month worked",SupportTeacher.MonhtWorked)
		SupportTeacher.TotalSum=SupportTeacher.MonhtWorked*SupportTeacher.Salary
		fmt.Println("Total sum",SupportTeacher.TotalSum)
		SupportTeacher.CreatedAt = helper.NullTimeStampToString(createdAt)
		SupportTeacher.UpdatedAt = helper.NullTimeStampToString(updatedAt)
		resp.SupportTeacher= append(resp.SupportTeacher, SupportTeacher)
	}

	queryCount := `SELECT COUNT(*) FROM "SupportTeacher" WHERE "deleted_at" is null AND TRUE ` + filter
	err = c.db.QueryRow(ctx, queryCount).Scan(&resp.Count)
	if err != nil {
		return nil, err
	}

	return resp, nil
}