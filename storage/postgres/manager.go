package postgres

import (
	"context"
	"database/sql"
	"log"
	ct "user_service/genproto/genproto/user_service"
	"user_service/pkg/hash"
	"user_service/pkg/helper"
	"user_service/storage"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v4/pgxpool"
)

type ManagerRepo struct {
	db *pgxpool.Pool
}

func NewManagerRepo(db *pgxpool.Pool) storage.ManagerRepoI {
	return &ManagerRepo{
		db: db,
	}
}

func (c *ManagerRepo) Create(ctx context.Context, req *ct.CreateManager) (*ct.ManagerPrimaryKey, error) {
	id := uuid.NewString()
	resp := &ct.ManagerPrimaryKey{Id: id}

	login,err:=helper.GenerateLoginID(c.db,"Manager")
	if err != nil {
		return nil, err
	}

	query := `INSERT INTO "Manager" (
			"ID",
			"FullName",
			"Phone",
			"Password",
			"Email",
			"Salary",
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
				NOW(),
				$8
			)`
		hashedPassword,err:=hash.HashPassword(req.Password)
		if err != nil {
			return nil, err
		}

	_, err = c.db.Exec(ctx, query, id,req.Fullname, req.Phone, hashedPassword, req.Email, 
		req.Salary,req.Branchid,login)
	if err != nil {
		log.Println("error while creating manager")
		return nil, err
	}

	return resp, err
}

func (c *ManagerRepo) GetByID(ctx context.Context, req *ct.ManagerPrimaryKey) (*ct.Manager, error) {
	resp := &ct.Manager{}
	query := `SELECT "ID",
					 "FullName",
					 "Phone",
					 "Email",
					 "Salary",
					 "BranchID",
					 "created_at",
					"updated_at"
			FROM "Manager"
			WHERE "ID"=$1 AND "deleted_at" is null`

	row := c.db.QueryRow(ctx, query, req.Id)

	var updatedAt, createdAt sql.NullTime
	if err := row.Scan(
		&resp.Id,
		&resp.Fullname,
		&resp.Phone,
		&resp.Email,
		&resp.Salary,
		&resp.Branchid,
		&createdAt,
		&updatedAt); err != nil {
		return nil, err
	}

	resp.CreatedAt = helper.NullTimeStampToString(createdAt)
	resp.UpdatedAt = helper.NullTimeStampToString(updatedAt)

	return resp, nil
}

func (c *ManagerRepo) GetList(ctx context.Context, req *ct.GetListManagerRequest) (*ct.GetListManagerResponse, error) {
	resp := &ct.GetListManagerResponse{}
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
					 "BranchID",
					 "created_at",
					 "updated_at"
			FROM "Manager"
        	WHERE "deleted_at" is null AND TRUE ` + filter + `
           OFFSET $1 LIMIT $2
    `

	rows, err := c.db.Query(ctx, query, offset, req.Limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		Manager := &ct.Manager{}
		var createdAt, updatedAt sql.NullTime
		if err := rows.Scan(
			&Manager.Id,
			&Manager.Fullname,
			&Manager.Phone,
			&Manager.Email,
			&Manager.Salary,
			&Manager.Branchid,
			&createdAt,
			&updatedAt); err != nil {
			return nil, err
		}

		Manager.CreatedAt = helper.NullTimeStampToString(createdAt)
		Manager.UpdatedAt = helper.NullTimeStampToString(updatedAt)
		resp.Manager = append(resp.Manager, Manager)
	}

	queryCount := `SELECT COUNT(*) FROM "Manager" WHERE "deleted_at" is null AND TRUE ` + filter
	err = c.db.QueryRow(ctx, queryCount).Scan(&resp.Count)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (c *ManagerRepo) Update(ctx context.Context, req *ct.UpdateManagerRequest) (*ct.MGMessage, error) {
	resp := &ct.MGMessage{Message: "Manager updated successfully"}
	query := `UPDATE "Manager" SET
								 "FullName"=$1,
								 "Phone"=$2,
								 "Email"=$3,
								 "Salary"=$4,
								 "BranchID"=$5,
								 "updated_at"=NOW()
								 WHERE "ID"=$6 AND "deleted_at" is null`
	_, err := c.db.Exec(ctx, query, req.Fullname, req.Phone, req.Email, req.Salary, req.Branchid,req.Id)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (c *ManagerRepo) Delete(ctx context.Context, req *ct.ManagerPrimaryKey) (*ct.MGMessage, error) {
	resp := &ct.MGMessage{Message: "Manager deleted successfully"}
	query := `UPDATE "Manager" SET
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

func (c *ManagerRepo) GetByGmail(ctx context.Context, req *ct.ManagerGmail) (*ct.ManagerGmailRes, error) {
	resp:=&ct.ManagerGmailRes{}
	query := `SELECT "ID","Password" FROM "Manager" WHERE "Email"=$1`
	err := c.db.QueryRow(ctx, query, req.Gmail).Scan(&resp.Gmail,&resp.Password)
	if err != nil {
		return nil, err
	}

	return resp, nil
}