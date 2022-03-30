package dto

import "github.com/avatardev/ipos-mblb-backend/internal/admin/location/entity"

type LocationResponse struct {
	Id   int64  `json:"id"`
	Name string `json:"location_name"`
}

type LocationsResponse struct {
	Location []*LocationResponse `json:"location"`
	Offset   uint64              `json:"offset"`
	Limit    uint64              `json:"limit"`
	Total    uint64              `json:"total"`
}

func NewLocationResponse(location *entity.Location) *LocationResponse {
	return &LocationResponse{
		Id:   location.Id,
		Name: location.Name,
	}
}

func NewLocationsResponse(locs entity.Locations, limit uint64, offset uint64, total uint64) *LocationsResponse {
	res := &LocationsResponse{
		Offset: offset,
		Limit:  limit,
		Total:  total,
	}

	for _, loc := range locs {
		data := NewLocationResponse(loc)
		res.Location = append(res.Location, data)
	}

	return res
}
