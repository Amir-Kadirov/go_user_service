package postgres

import (
	"context"
	"database/sql"
	ct "user_service/genproto/genproto/user_service"
	"user_service/pkg/helper"
	"user_service/storage"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v4/pgxpool"
)

type BranchRepo struct {
	db *pgxpool.Pool
}

func NewBranchRepo(db *pgxpool.Pool) storage.BranchRepoI {
	return &BranchRepo{
		db: db,
	}
}

func (b *BranchRepo) Create(ctx context.Context, req *ct.CreateBranch) (*ct.BranchPrimaryKey, error) {
	id := uuid.NewString()
	query := `INSERT INTO "Branch" (
					"ID",
					"Location",
					"Addres",
					"created_at"
					)
					VALUES (
					$1,
					ST_SetSRID(ST_MakePoint($2, $3), 4326),
					$4,
					NOW()
					);
	`
	_, err := b.db.Exec(ctx, query, id, req.Location.Longitude, req.Location.Latitude, req.Address)
	if err != nil {
		return nil, err
	}

	return &ct.BranchPrimaryKey{Id: id}, nil
}

func (c *BranchRepo) GetByID(ctx context.Context, req *ct.BranchPrimaryKey) (*ct.Branch, error) {
	resp := &ct.Branch{}
	query := `SELECT "ID",
					 "Addres",
					 ST_Y("Location") AS latitude, 
      				 ST_X("Location") AS longitude,
					 "created_at",
					 "updated_at"
			FROM "Branch"
			WHERE "ID"=$1 AND "deleted_at" is null`

	row := c.db.QueryRow(ctx, query, req.Id)

	var (createdAt, updatedAt sql.NullTime
		longitude, latitude      float64)
	err := row.Scan(
		&resp.Id,
		&resp.Address,
		&latitude,
		&longitude,
		&createdAt,
		&updatedAt)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil,err
		}
		return nil, err
	}

	resp.Location=&ct.Location{Longitude: longitude,Latitude: latitude}
	resp.CreatedAt = helper.NullTimeStampToString(createdAt)
	resp.UpdatedAt = helper.NullTimeStampToString(updatedAt)

	return resp, nil
}


func (c *BranchRepo) GetList(ctx context.Context, req *ct.GetListBranchRequest) (*ct.GetListBranchResponse, error) {
	resp := &ct.GetListBranchResponse{}
	if req.Offset == 0 {
		req.Offset = 1
	}

	filter := ""
	offset := (req.Offset - 1) * req.Limit

	if req.Search != "" {
		filter = ` AND "Addres" ILIKE '%` + req.Search + `%' `
	}

	query := `SELECT 
					 "ID",
					 "Addres",
					 ST_Y(location) AS latitude, 
      				 ST_X(location) AS longitude,
					 "created_at",
					 "updated_at"
			FROM "Branch"
        	WHERE "deleted_at" is null AND TRUE ` + filter + `
           OFFSET $1 LIMIT $2
    `

	rows, err := c.db.Query(ctx, query, offset, req.Limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		Branch := &ct.Branch{}
		var createdAt, updatedAt sql.NullTime
		if err := rows.Scan(
			&Branch.Id,
			&Branch.Address,
			&Branch.Location.Latitude,
			&Branch.Location.Longitude,
			&createdAt,
			&updatedAt); err != nil {
			return nil, err
		}

		Branch.CreatedAt = helper.NullTimeStampToString(createdAt)
		Branch.UpdatedAt = helper.NullTimeStampToString(updatedAt)
		resp.Branches = append(resp.Branches, Branch)
	}

	queryCount := `SELECT COUNT(*) FROM "Branch" WHERE "deleted_at" is null AND TRUE ` + filter
	err = c.db.QueryRow(ctx, queryCount).Scan(&resp.Count)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (c *BranchRepo) Update(ctx context.Context, req *ct.UpdateBranchRequest) (*ct.BMessage, error) {
	resp := &ct.BMessage{Message: "Branch updated successfully"}
	query := `UPDATE "Branch" SET
    						  "Addres" = $1,
    						  "Location" = ST_SetSRID(ST_MakePoint($2, $3), 4326),
    						  "updated_at" = NOW()
				WHERE "ID" = $4 AND "deleted_at" is null;
`
	_, err := c.db.Exec(ctx, query, req.Address, req.Location.Longitude,req.Location.Latitude,req.Id)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (c *BranchRepo) Delete(ctx context.Context, req *ct.BranchPrimaryKey) (*ct.BMessage, error) {
	resp := &ct.BMessage{Message: "Branch deleted successfully"}
	query := `UPDATE "Branch" SET
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