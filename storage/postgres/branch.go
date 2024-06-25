package postgres

import (
	"context"
	ct "user_service/genproto/genproto/user_service"
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

func (b *BranchRepo) Create(ctx context.Context, req *ct.CreateBranch) (*ct.BranchPrimaryKey, error){
	id:=uuid.NewString()
	query:=`INSERT INTO "Branch" (
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
	_,err:=b.db.Exec(ctx,query,id,req.Location.Longitude,req.Location.Latitude,req.Address)
	if err != nil {
		return nil, err
	}

	return &ct.BranchPrimaryKey{Id: id},nil
}