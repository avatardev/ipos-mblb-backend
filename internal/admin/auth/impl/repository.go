package impl

import (
	"context"
	"database/sql"
	"log"

	sq "github.com/Masterminds/squirrel"
	"github.com/avatardev/ipos-mblb-backend/internal/admin/auth/entity"
	"github.com/avatardev/ipos-mblb-backend/internal/global/database"
)

type AuthRepositoryImpl struct {
	DB *sql.DB
}

var SELECT_USER = sq.Select("u.id", "u.username", "u.password", "u.id_role", "d.nama_role", "u.id_seller", "u.plat_truk").From("users u").LeftJoin("role_users d ON d.id = u.id_role")

func NewAuthRepository(db *database.DatabaseClient) AuthRepositoryImpl {
	return AuthRepositoryImpl{DB: db.DB}
}

func (a *AuthRepositoryImpl) GetById(ctx context.Context, userId int64) (*entity.UserData, error) {
	stmt, params, err := SELECT_USER.Where(sq.And{sq.Eq{"u.id": userId}, sq.Eq{"u.deleted_at": nil}}).ToSql()
	if err != nil {
		log.Printf("[Auth.GetById] userId: %v, error:%v\n", userId, err)
		return nil, err
	}

	row := a.DB.QueryRowContext(ctx, stmt, params...)
	user := &entity.UserData{}
	queryErr := user.FromSql(row)
	if queryErr != nil && queryErr != sql.ErrNoRows {
		log.Printf("[Auth.GetById] userId: %v, error: %v\n", userId, queryErr)
		return nil, queryErr
	} else if queryErr == sql.ErrNoRows {
		log.Printf("[Auth.GetById] userId: %v, error: %v\n", userId, queryErr)
		return nil, nil
	}

	return user, nil
}

func (a *AuthRepositoryImpl) GetByUsername(ctx context.Context, username string) (*entity.UserData, error) {
	stmt, params, err := SELECT_USER.Where(sq.And{sq.Eq{"username": username}, sq.Eq{"deleted_at": nil}}).ToSql()
	if err != nil {
		log.Printf("[Auth.GetByUsername] username: %v, error: %v\n", username, err)
		return nil, err
	}

	row := a.DB.QueryRowContext(ctx, stmt, params...)
	user := &entity.UserData{}
	queryErr := user.FromSql(row)
	if queryErr != nil && queryErr != sql.ErrNoRows {
		log.Printf("[Auth.GetByUsername] username: %v, error: %v\n", username, queryErr)
		return nil, err
	} else if queryErr == sql.ErrNoRows {
		log.Printf("[Auth.GetByUsername] username: %v, error: %v\n", username, queryErr)
		return nil, nil
	}

	return user, nil
}
