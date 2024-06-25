package postgres

import (
	"context"
	"log"
	ct "user_service/genproto/genproto/user_service"
	"user_service/pkg/hash"
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
			"created_at") VALUES (
				$1,
				$2,
				$3,
				$4,
				$5,
				$6,
				$7,
				$8,
				$9,
				NOW()
			)`
		hashedPassword,err:=hash.HashPassword(req.Password)
		if err != nil {
			return nil, err
		}

	_, err = c.db.Exec(ctx, query, id,req.Fullname, req.Phone, hashedPassword, req.Email, 
		req.Salary, req.Ieltsscore,req.Ieltsattemptscount,req.Branchid)
	if err != nil {
		log.Println("error while creating supportteacher")
		return nil, err
	}

	return resp, err
}

// func (c *SupportTeacherRepo) GetByID(ctx context.Context, req *ct.SupportTeacherPrimaryKey) (*ct.SupportTeacher, error) {
// 	resp := &ct.SupportTeacher{}
// 	query := `SELECT "ID",
// 				   phone,
// 				   gmail,
// 				   language,
// 				   date_of_birth,
// 				   gender,
// 				   created_at,
// 				   updated_at
// 			FROM SupportTeacher
// 			WHERE id=$1 AND deleted_at is null`

// 	row := c.db.QueryRow(ctx, query, req.Id)

// 	var updatedAt, dateOfBirth, createdAt sql.NullTime
// 	if err := row.Scan(
// 		&resp.Id,
// 		&resp.Phone,
// 		&resp.Gmail,
// 		&resp.Language,
// 		&dateOfBirth,
// 		&resp.Gender,
// 		&createdAt,
// 		&updatedAt); err != nil {
// 		return nil, err
// 	}

// 	resp.DateOfBirth = helper.NullDateToString(dateOfBirth)
// 	resp.CreatedAt = helper.NullTimeStampToString(createdAt)
// 	resp.UpdatedAt = helper.NullTimeStampToString(updatedAt)

// 	return resp, nil
// }

// func (c *SupportTeacherRepo) GetList(ctx context.Context, req *ct.GetListSupportTeacherRequest) (*ct.GetListSupportTeacherResponse, error) {
// 	resp := &ct.GetListSupportTeacherResponse{}

// 	filter := ""
// 	offset := (req.Offset - 1) * req.Limit

// 	if req.Search != "" {
// 		filter = ` AND gender ILIKE '%` + req.Search + `%' `
// 	}

// 	query := `SELECT 
// 				id,
// 				phone,
// 				gmail,
// 				language,
// 				date_of_birth,
// 				gender,
// 				created_at,
// 				updated_at
// 			FROM SupportTeacher
//         	WHERE deleted_at is null AND TRUE ` + filter + `
//            OFFSET $1 LIMIT $2
//     `

// 	rows, err := c.db.Query(ctx, query, offset, req.Limit)
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer rows.Close()

// 	for rows.Next() {
// 		SupportTeacher := &ct.SupportTeacher{}
// 		var createdAt, updatedAt, dateOfBirth sql.NullTime
// 		if err := rows.Scan(
// 			&SupportTeacher.Id,
// 			&SupportTeacher.Phone,
// 			&SupportTeacher.Gmail,
// 			&SupportTeacher.Language,
// 			&dateOfBirth,
// 			&SupportTeacher.Gender,
// 			&createdAt,
// 			&updatedAt); err != nil {
// 			return nil, err
// 		}
// 		SupportTeacher.DateOfBirth = helper.NullDateToString(dateOfBirth)
// 		SupportTeacher.CreatedAt = helper.NullTimeStampToString(createdAt)
// 		SupportTeacher.UpdatedAt = helper.NullTimeStampToString(updatedAt)
// 		resp.SupportTeacher = append(resp.SupportTeacher, SupportTeacher)
// 	}

// 	queryCount := `SELECT COUNT(*) FROM SupportTeacher WHERE deleted_at is null AND TRUE ` + filter
// 	err = c.db.QueryRow(ctx, queryCount).Scan(&resp.Count)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return resp, nil
// }

// func (c *SupportTeacherRepo) Update(ctx context.Context, req *ct.UpdateSupportTeacherRequest) (*ct.UpdateSupportTeacherResponse, error) {
// 	resp := &ct.UpdateSupportTeacherResponse{Message: "SupportTeacher updated successfully"}
// 	query := `UPDATE SupportTeacher SET phone=$1,
// 								 gmail=$2,
// 								 language=$3,
// 								 date_of_birth=$4,
// 								 gender=$5,
// 								 updated_at=NOW()
// 								 WHERE id=$6 AND deleted_at is null`
// 	_, err := c.db.Exec(ctx, query, req.Phone, req.Gmail, req.Language, req.DateOfBirth, req.Gender, req.Id)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return resp, nil
// }

// func (c *SupportTeacherRepo) Delete(ctx context.Context, req *ct.SupportTeacherPrimaryKey) (*ct.Empty, error) {
// 	resp := &ct.Empty{}
// 	query := `UPDATE SupportTeacher SET
// 							 deleted_at=NOW()
// 							 WHERE id=$1 AND deleted_at is null RETURNING created_at`

// 	var createdAt sql.NullTime
// 	err := c.db.QueryRow(ctx, query, req.Id).Scan(&createdAt)
// 	if err != nil {
// 		return nil, err
// 	}

// 	if err = helper.DeleteChecker(createdAt); err != nil {
// 		return resp, nil
// 	}

// 	return resp, nil
// }

// func (c *SupportTeacherRepo) GetByGmail(ctx context.Context, req *ct.SupportTeacherGmail) (*ct.SupportTeacherPrimaryKey, error) {
// 	query := `SELECT id FROM SupportTeacher WHERE gmail=$1`
// 	var id string
// 	err := c.db.QueryRow(ctx, query, req.Gmail).Scan(&id)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return &ct.SupportTeacherPrimaryKey{Id: id}, nil
// }
