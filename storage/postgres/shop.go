package postgres

import (
	"context"
	"database/sql"
	"log"
	ct "user_service/genproto/user_service"
	"user_service/pkg/helper"
	"user_service/storage"

	"github.com/google/uuid"
	"github.com/gosimple/slug"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/lib/pq"
)

type shopRepo struct {
	db *pgxpool.Pool
}

func NewShopRepo(db *pgxpool.Pool) storage.ShopRepoI {
	return &shopRepo{
		db: db,
	}
}

func (c *shopRepo) 	Create(ctx context.Context,req *ct.CreateShop) (resp *ct.ShopPrimaryKey,err error){
	id := uuid.NewString()
	resp = &ct.ShopPrimaryKey{Id: id}
	slug:=slug.Make(req.NameEn)

	query := `INSERT INTO shop (
			slug,
			phone,
			name_uz,
			name_ru,
			name_en,
			description_uz,
			description_ru,
			description_en,
			location,
			currency,
			payment_types,
			id,
			created_at) VALUES (
				$1,
				$2,
				$3,
				$4,
				$5,
				$6,
				$7,
				$8,
				$9,
				$10,
				$11, 
				$12,
				NOW()
			)`
	_, err = c.db.Exec(ctx, query, slug, req.Phone, req.NameUz, req.NameRu, req.NameEn, req.DescriptionUz,
					   req.DescriptionRu,req.DescriptionEn,req.Location,req.Currency,pq.Array(req.PaymentTypes),id)
	if err != nil { 
		log.Println("error while creating customer")
		return nil, err
	}

	return resp, err
}

func (c *shopRepo) GetById(ctx context.Context, req *ct.ShopPrimaryKey) (resp *ct.GetByID,err error) {
	resp=&ct.GetByID{}
	
	query:=`SELECT 
			slug,
			phone,
			name_uz,
			name_ru,
			name_en,	
			description_uz,
			description_ru,
			description_en,
			currency,
			id,
			COALESCE(payment_types, '{}'),
			created_at,
			updated_at
			FROM shop
			WHERE id=$1 AND deleted_at is null`

	row:=c.db.QueryRow(ctx,query,req.Id)
	var createdAt,updatedAt sql.NullTime
	if err=row.Scan(
		&resp.Slug,
		&resp.Phone,
		&resp.NameUz,
		&resp.NameRu,
		&resp.NameEn,
		&resp.DescriptionUz,
		&resp.DescriptionRu,
		&resp.DescriptionEn,
		&resp.Currency,
		&resp.Id,
		&resp.PaymentTypes,
		&createdAt,
		&updatedAt);err!=nil {
		return nil,err
	}
	resp.CreatedAt=helper.NullTimeStampToString(createdAt)
	resp.UpdatedAt=helper.NullTimeStampToString(updatedAt)

	return resp,nil
}


func (c *shopRepo) Update(ctx context.Context, req *ct.UpdateShopRequest) (resp *ct.ShopEmpty, err error) {
	resp = &ct.ShopEmpty{}
	query := `UPDATE shop SET 	phone=$1,
								name_uz=$2,
								name_ru=$3,
								name_en=$4,	
								description_uz=$5,
								description_ru=$6,
								description_en=$7,
								location=$8,
								currency=$9,
								payment_types=$10,
								 updated_at=NOW()
								 WHERE id=$11 AND deleted_at is null`
	_, err = c.db.Exec(ctx, query, req.Phone, req.NameUz, req.NameRu, req.NameEn, req.DescriptionUz,
									req.DescriptionRu,req.DescriptionEn,req.Location,req.Currency,pq.Array(req.PaymentTypes),req.Id)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (c *shopRepo) Delete(ctx context.Context,req *ct.ShopPrimaryKey) (resp *ct.ShopEmpty,err error) {
	query := `UPDATE shop SET
							 deleted_at=NOW()
							 WHERE id=$1 RETURNING created_at`

	var createdAt sql.NullTime
	err = c.db.QueryRow(ctx, query, req.Id).Scan(&createdAt)
	if err != nil {
		return nil, err
	}

	if err=helper.DeleteChecker(createdAt);err!=nil {
		return resp,nil
	}

	return resp, nil
}

func (c *shopRepo) GetList(ctx context.Context,req *ct.GetListShopRequest) (resp *ct.GetListShopResponse,err error) {
	resp = &ct.GetListShopResponse{}
	shop := &ct.Shop{}

	filter := ""
    offset := (req.Offset - 1) * req.Limit

    if req.Search != "" {
        filter = ` AND description_uz ILIKE '%` + req.Search + `%' `
    }

	query := `SELECT 
				slug,
				phone,
				name_uz,
				name_ru,
				name_en,	
				description_uz,
				description_ru,
				description_en,
				currency,
				id,
				COALESCE(payment_types, '{}'),
				created_at,
				updated_at
			FROM shop
			WHERE deleted_at is null AND TRUE ` + filter + `
			OFFSET $1 LIMIT $2
    `

	rows, err := c.db.Query(ctx, query,offset,req.Limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var createdAt, updatedAt sql.NullTime
		if err := rows.Scan(
			&shop.Slug,
			&shop.Phone,
			&shop.NameUz,
			&shop.NameRu,
			&shop.NameEn,
			&shop.DescriptionUz,
			&shop.DescriptionRu,
			&shop.DescriptionEn,
			&shop.Currency,
			&shop.Id,
			&shop.PaymentTypes,
			&createdAt,
			&updatedAt); err != nil {
			return nil, err
		}

		shop.CreatedAt = helper.NullTimeStampToString(createdAt)
		shop.UpdatedAt = helper.NullTimeStampToString(updatedAt)
		resp.Shop = append(resp.Shop, shop)
	}

	queryCount := `SELECT COUNT(*) FROM shop WHERE deleted_at is null AND TRUE ` + filter +``
	err = c.db.QueryRow(ctx, queryCount).Scan(&resp.Count)
	if err != nil {
		return nil, err
	}

	return resp, nil
}