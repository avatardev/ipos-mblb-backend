package dto

import (
	"encoding/json"
	"io"

	"github.com/avatardev/ipos-mblb-backend/internal/admin/location/entity"
)

type LocationRequest struct {
	Name string `json:"location_name" validate:"required"`
}

func (l *LocationRequest) FromJSON(r io.Reader) error {
	return json.NewDecoder(r).Decode(l)
}

func (l *LocationRequest) ToEntity() entity.Location {
	return entity.Location{
		Name: l.Name,
	}
}
