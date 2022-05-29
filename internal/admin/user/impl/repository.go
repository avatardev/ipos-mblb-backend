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
	"github.com/avatardev/ipos-mblb-backend/pkg/util/privutil"
)

type UserRepositoryImpl struct {
	DB *sql.DB
}

func NewUserRepository(db *database.DatabaseClient) UserRepositoryImpl {
	return UserRepositoryImpl{DB: db.DB}
}

var (
	COUNT_USER         = sq.Select("COUNT(*)").From("users")
	SELECT_USER        = sq.Select("u.id", "u.username").From("users u")
	INSERT_USER        = sq.Insert("users").Columns("username", "password", "id_role", "created_at", "updated_at", "status_login")
	INSERT_USER_SELLER = sq.Insert("users").Columns("username", "password", "id_role", "id_seller", "created_at", "updated_at", "status_login")
	INSERT_USER_BUYER  = sq.Insert("users").Columns("username", "password", "id_role", "plat_truk", "created_at", "updated_at", "status_login")
	UPDATE_USER        = sq.Update("users")
)

func (ur UserRepositoryImpl) Count(ctx context.Context, keyword string, role int64) (uint64, error) {
	stmt, params, err := COUNT_USER.Where(sq.And{sq.Eq{"id_role": role}, sq.Like{"username": fmt.Sprintf("%%%s%%", keyword)}, sq.Eq{"deleted_at": nil}}).ToSql()
	if err != nil {
		log.Printf("[User.Count] role: %v, err: %v\n", role, err)
		return 0, err
	}

	var userCount uint64
	queryErr := ur.DB.QueryRowContext(ctx, stmt, params...).Scan(&userCount)
	if queryErr != nil {
		log.Printf("[User.Count] role: %v, err: %v\n", role, queryErr)
		return 0, queryErr
	}

	return userCount, nil
}

func (ur UserRepositoryImpl) CountSeller(ctx context.Context, seller int64, role int64) (uint64, error) {
	stmt, params, err := COUNT_USER.Where(sq.And{sq.Eq{"id_role": role}, sq.Eq{"id_seller": seller}, sq.Eq{"deleted_at": nil}}).ToSql()
	if err != nil {
		log.Printf("[User.CountSeller] role: %v, err: %v\n", role, err)
		return 0, err
	}

	var userCount uint64
	queryErr := ur.DB.QueryRowContext(ctx, stmt, params...).Scan(&userCount)
	if queryErr != nil {
		log.Printf("[User.CountSeller] role: %v, err: %v\n", role, queryErr)
		return 0, queryErr
	}

	return userCount, nil
}

func (ur UserRepositoryImpl) CountBuyer(ctx context.Context, v_plate string, role int64) (uint64, error) {
	stmt, params, err := COUNT_USER.Where(sq.And{sq.Eq{"id_role": role}, sq.Eq{"plat_truk": v_plate}, sq.Eq{"deleted_at": nil}}).ToSql()
	if err != nil {
		log.Printf("[User.CountBuyer] role: %v, err: %v\n", role, err)
		return 0, err
	}

	var userCount uint64
	queryErr := ur.DB.QueryRowContext(ctx, stmt, params...).Scan(&userCount)
	if queryErr != nil {
		log.Printf("[User.CountBuyer] role: %v, err: %v\n", role, queryErr)
		return 0, queryErr
	}

	return userCount, nil
}

func (ur UserRepositoryImpl) CountUserByUsername(ctx context.Context, uname string) (exist bool, err error) {
	stmt, params, err := COUNT_USER.Where(sq.Like{"username": fmt.Sprintf("%%%s%%", uname)}, sq.Eq{"deleted_at": nil}).ToSql()
	if err != nil {
		log.Printf("[User.CountUserByUsername] err: %v\n", err)
		return
	}

	var userCount uint64
	queryErr := ur.DB.QueryRowContext(ctx, stmt, params...).Scan(&userCount)
	if queryErr != nil {
		log.Printf("[User.CountUserByUsername] err: %v\n", queryErr)
		return
	}

	if userCount == 1 {
		exist = true
	}

	return
}

func (ur UserRepositoryImpl) GetAll(ctx context.Context, role int64, keyword string, limit uint64, offset uint64) (entity.Users, error) {
	stmt, params, err := SELECT_USER.Where(sq.And{sq.Eq{"u.id_role": role}, sq.Like{"u.username": fmt.Sprintf("%%%s%%", keyword)}, sq.Eq{"deleted_at": nil}}).Limit(limit).Offset(offset).ToSql()
	if err != nil {
		log.Printf("[User.GetAll] role: %v, err: %v\n", role, err)
		return nil, err
	}

	rows, err := ur.DB.QueryContext(ctx, stmt, params...)
	if err != nil {
		log.Printf("[User.GetAll] role: %v, err: %v\n", role, err)
		return nil, err
	}

	users := entity.Users{}

	for rows.Next() {
		var user entity.User
		err := rows.Scan(&user.Id, &user.Username)
		if err != nil {
			log.Printf("[User.GetAll] role: %v, err: %v\n", role, err)
			return nil, err
		}
		users = append(users, &user)
	}

	return users, nil
}

func (ur UserRepositoryImpl) GetAllSeller(ctx context.Context, role int64, seller int64, keyword string, limit uint64, offset uint64) (entity.Users, error) {
	stmt, params, err := SELECT_USER.Where(sq.And{sq.Eq{"u.id_role": role}, sq.Like{"u.username": fmt.Sprintf("%%%s%%", keyword)}, sq.Eq{"u.id_seller": seller}, sq.Eq{"deleted_at": nil}}).Limit(limit).Offset(offset).ToSql()
	if err != nil {
		log.Printf("[User.GetAllSeller] role: %v, err: %v\n", role, err)
		return nil, err
	}

	rows, err := ur.DB.QueryContext(ctx, stmt, params...)
	if err != nil {
		log.Printf("[User.GetAllSeller] role: %v, err: %v\n", role, err)
		return nil, err
	}

	users := entity.Users{}

	for rows.Next() {
		var user entity.User
		err := rows.Scan(&user.Id, &user.Username)
		if err != nil {
			log.Printf("[User.GetAllSeller] role: %v, err: %v\n", role, err)
			return nil, err
		}
		users = append(users, &user)
	}

	return users, nil
}

func (ur UserRepositoryImpl) GetAllBuyer(ctx context.Context, role int64, v_plate string, limit uint64, offset uint64) (entity.Users, error) {
	stmt, params, err := SELECT_USER.Where(sq.And{sq.Eq{"u.id_role": role}, sq.Eq{"plat_truk": v_plate}, sq.Eq{"deleted_at": nil}}).Limit(limit).Offset(offset).ToSql()
	if err != nil {
		log.Printf("[User.GetAllBuyer] role: %v, err: %v\n", role, err)
		return nil, err
	}

	rows, err := ur.DB.QueryContext(ctx, stmt, params...)
	if err != nil {
		log.Printf("[User.GetAllBuyer] role: %v, err: %v\n", role, err)
		return nil, err
	}

	users := entity.Users{}

	for rows.Next() {
		var user entity.User
		err := rows.Scan(&user.Id, &user.Username)
		if err != nil {
			log.Printf("[User.GetAllBuyer] role: %v, err: %v\n", role, err)
			return nil, err
		}
		users = append(users, &user)
	}

	return users, nil
}

func (ur UserRepositoryImpl) GetById(ctx context.Context, role int64, id int64) (*entity.User, error) {
	stmt, params, err := SELECT_USER.Where(sq.And{sq.Eq{"u.id": id}, sq.Eq{"u.id_role": role}, sq.Eq{"deleted_at": nil}}).ToSql()
	if err != nil {
		log.Printf("[User.GetById] id: %v, err: %v\n", id, err)
		return nil, err
	}

	user := &entity.User{}

	rows, err := ur.DB.QueryContext(ctx, stmt, params...)
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
	query := INSERT_USER.Values(user.Username, user.Password, role, currTime, currTime, 0)

	if role == privutil.USER_SELLER {
		query = INSERT_USER_SELLER.Values(user.Username, user.Password, role, user.SellerId, currTime, currTime, 0)
	} else if role == privutil.USER_BUYER {
		query = INSERT_USER_BUYER.Values(user.Username, user.Password, role, user.VPlate, currTime, currTime, 0)
	}

	stmt, params, err := query.ToSql()
	if err != nil {
		log.Printf("[User.Store] role: %v, err: %v\n", role, err)
		return nil, err
	}

	res, err := ur.DB.ExecContext(ctx, stmt, params...)
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

	if user.Password != "" {
		updateMap = map[string]interface{}{
			"username": user.Username,
			"password": user.Password,
		}
	}

	stmt, params, err := UPDATE_USER.SetMap(updateMap).Where(sq.Eq{"id": user.Id}).ToSql()
	if err != nil {
		log.Printf("[User.Update] id: %v, err: %v\n", user.Id, err)
		return nil, err
	}

	if _, err := ur.DB.ExecContext(ctx, stmt, params...); err != nil {
		log.Printf("[User.Update] id: %v, err: %v\n", user.Id, err)
		return nil, err
	}

	return ur.GetById(ctx, role, user.Id)
}

func (ur UserRepositoryImpl) Delete(ctx context.Context, role int64, id int64) error {
	updateMap := map[string]interface{}{
		"deleted_at": time.Now(),
	}

	stmt, params, err := UPDATE_USER.SetMap(updateMap).Where(sq.And{sq.Eq{"id": id}, sq.Eq{"id_role": role}}).ToSql()
	if err != nil {
		log.Printf("[User.Update] id: %v, err: %v\n", id, err)
		return err
	}

	if _, err := ur.DB.ExecContext(ctx, stmt, params...); err != nil {
		log.Printf("[User.Update] id: %v, err: %v\n", id, err)
		return err
	}

	return nil
}
