package impl

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"

	sq "github.com/Masterminds/squirrel"
	"github.com/avatardev/ipos-mblb-backend/internal/admin/user/entity"
	"github.com/avatardev/ipos-mblb-backend/internal/global/database"
)

type UserRepositoryImpl struct {
	DB *sql.DB
}

func NewUserRepository(db *database.DatabaseClient) UserRepositoryImpl {
	return UserRepositoryImpl{DB: db.DB}
}

var (
	COUNT_USER  = sq.Select("COUNT(*)").From("users")
	SELECT_USER = sq.Select("u.id", "u.username").From("users u")
	INSERT_USER = sq.Insert("users").Columns("username", "password", "id_role", "created_at", "updated_at", "status_login")
	UPDATE_USER = sq.Update("users")
	DELETE_USER = sq.Delete("users")
)

func (ur UserRepositoryImpl) Count(ctx context.Context, keyword string, role int64) (uint64, error) {
	stmt, params, err := COUNT_USER.Where(sq.And{sq.Eq{"id_role": role}, sq.Like{"username": fmt.Sprintf("%%%s%%", keyword)}}).ToSql()
	if err != nil {
		log.Printf("[User.Count] role: %v, err: %v\n", role, err)
		return 0, err
	}

	prpd, err := ur.DB.PrepareContext(ctx, stmt)
	if err != nil {
		log.Printf("[User.Count] role: %v, err: %v\n", role, err)
		return 0, err
	}

	var userCount uint64
	queryErr := prpd.QueryRowContext(ctx, params...).Scan(&userCount)
	if queryErr != nil {
		log.Printf("[User.Count] role: %v, err: %v\n", role, queryErr)
		return 0, queryErr
	}

	return userCount, nil
}

func (ur UserRepositoryImpl) GetAll(ctx context.Context, role int64, keyword string, limit uint64, offset uint64) (entity.Users, error) {
	stmt, params, err := SELECT_USER.Where(sq.And{sq.Eq{"u.id_role": role}, sq.Like{"username": fmt.Sprintf("%%%s%%", keyword)}}).Limit(limit).Offset(offset).ToSql()
	if err != nil {
		log.Printf("[User.GetAll] role: %v, err: %v\n", role, err)
		return nil, err
	}

	prpd, err := ur.DB.PrepareContext(ctx, stmt)
	if err != nil {
		log.Printf("[User.GetAll] role: %v, err: %v\n", role, err)
		return nil, err
	}

	rows, err := prpd.QueryContext(ctx, params...)
	if err != nil {
		log.Printf("[User.GetAll] role: %v, err: %v\n", role, err)
		return nil, err
	}

	users := entity.Users{}

	for rows.Next() {
		var user entity.User
		err := rows.Scan(&user.Id, &user.Username)
		if err != nil {
			log.Printf("[Category.GetAll] role: %v, err: %v\n", role, err)
			return nil, err
		}
		users = append(users, &user)
	}

	return users, nil
}

func (ur UserRepositoryImpl) GetById(ctx context.Context, role int64, id int64) (*entity.User, error) {
	stmt, params, err := SELECT_USER.Where(sq.And{sq.Eq{"u.id": id}, sq.Eq{"u.id_role": role}}).ToSql()
	if err != nil {
		log.Printf("[User.GetById] id: %v, err: %v\n", id, err)
		return nil, err
	}

	prpd, err := ur.DB.PrepareContext(ctx, stmt)
	if err != nil {
		log.Printf("[User.GetById] id: %v, err: %v\n", id, err)
		return nil, err
	}

	user := &entity.User{}

	rows := prpd.QueryRowContext(ctx, params...)
	queryErr := rows.Scan(&user.Id, &user.Username)
	if queryErr != nil && queryErr != sql.ErrNoRows {
		log.Printf("[User.GetById] id: %v, err: %v\n", id, queryErr)
		return nil, err
	} else if queryErr == sql.ErrNoRows {
		log.Printf("[User.GetById] id: %v, err: %v\n", id, queryErr)
		return nil, nil
	}

	return user, nil
}

func (ur UserRepositoryImpl) Store(ctx context.Context, role int64, user entity.User) (*entity.User, error) {
	currTime := time.Now()
	stmt, params, err := INSERT_USER.Values(user.Username, user.Password, role, currTime, currTime, 0).ToSql()
	if err != nil {
		log.Printf("[User.Store] role: %v, err: %v\n", role, err)
		return nil, err
	}

	prpd, err := ur.DB.PrepareContext(ctx, stmt)
	if err != nil {
		log.Printf("[User.Store] role: %v, err: %v\n", role, err)
		return nil, err
	}

	res, err := prpd.ExecContext(ctx, params...)
	if err != nil {
		log.Printf("[User.Store] role: %v, err: %v\n", role, err)
		return nil, err
	}

	lid, err := res.LastInsertId()
	if err != nil {
		log.Printf("[User.Store] role: %v, err: %v\n", role, err)
		return nil, err
	}

	return ur.GetById(ctx, role, lid)
}

func (ur UserRepositoryImpl) Update(ctx context.Context, role int64, user entity.User) (*entity.User, error) {
	updateMap := map[string]interface{}{
		"username": user.Username,
	}

	stmt, params, err := UPDATE_USER.SetMap(updateMap).Where(sq.Eq{"id": user.Id}).ToSql()
	if err != nil {
		log.Printf("[User.Update] id: %v, err: %v\n", user.Id, err)
		return nil, err
	}

	prpd, err := ur.DB.PrepareContext(ctx, stmt)
	if err != nil {
		log.Printf("[User.Update] id: %v, err: %v\n", user.Id, err)
		return nil, err
	}

	if _, err := prpd.ExecContext(ctx, params...); err != nil {
		log.Printf("[User.Update] id: %v, err: %v\n", user.Id, err)
		return nil, err
	}

	return ur.GetById(ctx, role, user.Id)
}

func (ur UserRepositoryImpl) Delete(ctx context.Context, role int64, id int64) error {
	stmt, params, err := DELETE_USER.Where(sq.And{sq.Eq{"id": id}, sq.Eq{"id_role": role}}).ToSql()
	if err != nil {
		log.Printf("[User.Update] id: %v, err: %v\n", id, err)
		return err
	}

	prpd, err := ur.DB.PrepareContext(ctx, stmt)
	if err != nil {
		log.Printf("[User.Update] id: %v, err: %v\n", id, err)
		return err
	}

	if _, err := prpd.ExecContext(ctx, params...); err != nil {
		log.Printf("[User.Update] id: %v, err: %v\n", id, err)
		return err
	}

	return nil
}