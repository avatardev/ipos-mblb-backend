package impl

import (
	"context"
	"log"

	"github.com/avatardev/ipos-mblb-backend/internal/dto"
	"github.com/avatardev/ipos-mblb-backend/pkg/errors"
	"golang.org/x/crypto/bcrypt"
)

type UserServiceImpl struct {
	Ur UserRepositoryImpl
}

func (u *UserServiceImpl) GetUser(ctx context.Context, role int64, keyword string, limit uint64, offset uint64) (*dto.UsersResponse, error) {
	userCount, err := u.Ur.Count(ctx, keyword, role)
	if err != nil {
		return nil, err
	}

	if userCount == 0 {
		return nil, errors.ErrInvalidResources
	}

	users, err := u.Ur.GetAll(ctx, role, keyword, limit, offset)
	if err != nil {
		return nil, err
	}

	if len(users) == 0 {
		return nil, errors.ErrInvalidResources
	}

	return dto.NewUsersResponse(users, limit, offset, userCount), nil
}

func (u *UserServiceImpl) GetUserSeller(ctx context.Context, role int64, seller int64, keyword string, limit uint64, offset uint64) (*dto.UsersResponse, error) {
	userCount, err := u.Ur.CountSeller(ctx, seller, role)
	if err != nil {
		return nil, err
	}

	if userCount == 0 {
		return nil, errors.ErrInvalidResources
	}

	users, err := u.Ur.GetAllSeller(ctx, role, seller, keyword, limit, offset)
	if err != nil {
		return nil, err
	}

	if len(users) == 0 {
		return nil, errors.ErrInvalidResources
	}

	return dto.NewUsersResponse(users, limit, offset, userCount), nil
}

func (u *UserServiceImpl) GetUserBuyer(ctx context.Context, role int64, v_plate string, limit uint64, offset uint64) (*dto.UsersResponse, error) {
	userCount, err := u.Ur.CountBuyer(ctx, v_plate, role)
	if err != nil {
		return nil, err
	}

	if userCount == 0 {
		return nil, errors.ErrInvalidResources
	}

	users, err := u.Ur.GetAllBuyer(ctx, role, v_plate, limit, offset)
	if err != nil {
		return nil, err
	}

	if len(users) == 0 {
		return nil, errors.ErrInvalidResources
	}

	return dto.NewUsersResponse(users, limit, offset, userCount), nil
}

func (u *UserServiceImpl) GetUserById(ctx context.Context, role int64, id int64) (*dto.UserResponse, error) {
	user, err := u.Ur.GetById(ctx, role, id)
	if err != nil {
		return nil, err
	}

	if user == nil {
		return nil, errors.ErrNotFound
	}

	return dto.NewUserResponse(user), nil
}

func (u *UserServiceImpl) StoreUser(ctx context.Context, role int64, req *dto.UserPostRequest) (*dto.UserResponse, error) {
	user := req.ToEntity()

	hashed, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Printf("[Service.StoreUser] role: %v, error: %v\n", role, err)
		return nil, err
	} else {
		user.Password = string(hashed)
	}

	data, err := u.Ur.Store(ctx, role, user)
	if err != nil {
		return nil, err

	}

	if data == nil {
		return nil, errors.ErrUnknown
	}

	return dto.NewUserResponse(data), nil
}

func (u *UserServiceImpl) UpdateUser(ctx context.Context, role int64, id int64, req *dto.UserPutRequest) (*dto.UserResponse, error) {
	user := req.ToEntity()
	user.Id = id

	exists, err := u.Ur.GetById(ctx, role, id)
	if err != nil {
		return nil, err
	}

	if exists == nil {
		return nil, errors.ErrNotFound
	}

	if user.Password != "" {
		hashed, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
		if err != nil {
			log.Printf("[Service.StoreUser] role: %v, error: %v\n", role, err)
			return nil, err
		} else {
			user.Password = string(hashed)
		}
	}

	data, err := u.Ur.Update(ctx, role, user)
	if err != nil {
		return nil, err
	}

	if data == nil {
		return nil, errors.ErrNotFound
	}

	return dto.NewUserResponse(data), nil
}

func (u *UserServiceImpl) DeleteUser(ctx context.Context, role int64, id int64) error {
	exists, err := u.Ur.GetById(ctx, role, id)
	if err != nil {
		return err
	}

	if exists == nil {
		return errors.ErrNotFound
	}

	if err := u.Ur.Delete(ctx, role, id); err != nil {
		return err
	}

	return nil
}
