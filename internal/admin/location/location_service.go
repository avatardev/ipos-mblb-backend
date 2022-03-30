package location

import (
	"context"

	"github.com/avatardev/ipos-mblb-backend/internal/admin/location/impl"
	"github.com/avatardev/ipos-mblb-backend/internal/dto"
)

type LocationService interface {
	GetLocation(ctx context.Context, keyword string, limit uint64, offset uint64) (*dto.LocationsResponse, error)
	GetLocationById(ctx context.Context, id int64) (*dto.LocationResponse, error)
	StoreLocation(ctx context.Context, req *dto.LocationRequest) (*dto.LocationResponse, error)
	UpdateLocation(ctx context.Context, id int64, req *dto.LocationRequest) (*dto.LocationResponse, error)
	DeleteLocation(ctx context.Context, id int64) error
}

func NewLocationService(Lr impl.LocationRepositoryImpl) LocationService {
	return &impl.LocationServiceImpl{Lr: Lr}
}
