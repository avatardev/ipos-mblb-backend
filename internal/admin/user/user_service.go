package user

import (
	"context"

	"github.com/avatardev/ipos-mblb-backend/internal/admin/user/impl"
	"github.com/avatardev/ipos-mblb-backend/internal/dto"
)

type UserService interface {
	GetUser(ctx context.Context, role int64, keyword string, limit uint64, offset uint64) (*dto.UsersResponse, error)
	GetUserBuyer(ctx context.Context, role int64, v_plate string, limit uint64, offset uint64) (*dto.UsersResponse, error)
	GetUserSeller(ctx context.Context, role int64, seller int64, keyword string, limit uint64, offset uint64) (*dto.UsersResponse, error)
	GetUserById(ctx context.Context, role int64, id int64) (*dto.UserResponse, error)
	StoreUser(ctx context.Context, role int64, req *dto.UserPostRequest) (*dto.UserResponse, error)
	UpdateUser(ctx context.Context, role int64, id int64, req *dto.UserPutRequest) (*dto.UserResponse, error)
	DeleteUser(ctx context.Context, role int64, id int64) error
}

func NewUserService(Ur impl.UserRepositoryImpl) UserService {
	return &impl.UserServiceImpl{Ur: Ur}
}
