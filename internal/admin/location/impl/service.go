package impl

import (
	"context"

	"github.com/avatardev/ipos-mblb-backend/internal/dto"
	"github.com/avatardev/ipos-mblb-backend/pkg/errors"
)

type LocationServiceImpl struct {
	Lr LocationRepositoryImpl
}

func (ls *LocationServiceImpl) GetLocation(ctx context.Context, keyword string, limit uint64, offset uint64) (*dto.LocationsResponse, error) {
	count, err := ls.Lr.Count(ctx, keyword)
	if err != nil {
		return nil, err
	}

	if count == 0 {
		return nil, errors.ErrInvalidResources
	}

	locations, err := ls.Lr.GetAll(ctx, keyword, limit, offset)
	if err != nil {
		return nil, err
	}

	if len(locations) == 0 {
		return nil, errors.ErrInvalidResources
	}

	return dto.NewLocationsResponse(locations, limit, offset, count), nil
}

func (ls *LocationServiceImpl) GetLocationById(ctx context.Context, id int64) (*dto.LocationResponse, error) {
	loc, err := ls.Lr.GetById(ctx, id)
	if err != nil {
		return nil, err
	}

	if loc == nil {
		return nil, errors.ErrNotFound
	}

	return dto.NewLocationResponse(loc), nil
}

func (ls *LocationServiceImpl) StoreLocation(ctx context.Context, req *dto.LocationRequest) (*dto.LocationResponse, error) {
	loc := req.ToEntity()

	data, err := ls.Lr.Store(ctx, loc)
	if err != nil {
		return nil, err
	}

	if data == nil {
		return nil, errors.ErrUnknown
	}

	return dto.NewLocationResponse(data), nil
}

func (ls *LocationServiceImpl) UpdateLocation(ctx context.Context, id int64, req *dto.LocationRequest) (*dto.LocationResponse, error) {
	loc := req.ToEntity()
	loc.Id = id

	exists, err := ls.Lr.GetById(ctx, loc.Id)
	if err != nil {
		return nil, err
	}

	if exists == nil {
		return nil, errors.ErrNotFound
	}

	data, err := ls.Lr.Update(ctx, loc)
	if err != nil {
		return nil, err
	}

	if data == nil {
		return nil, errors.ErrNotFound
	}

	return dto.NewLocationResponse(data), nil
}

func (ls *LocationServiceImpl) DeleteLocation(ctx context.Context, id int64) error {
	exists, err := ls.Lr.GetById(ctx, id)
	if err != nil {
		return err
	}

	if exists == nil {
		return errors.ErrNotFound
	}

	if err := ls.Lr.Delete(ctx, id); err != nil {
		return err
	}

	return nil
}
