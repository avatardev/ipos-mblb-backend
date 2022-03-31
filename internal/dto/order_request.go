package dto

import (
	"encoding/json"
	"io"
)

type InsertNoteRequest struct {
	Note string `json:"note" validate:"required"`
}

func (i *InsertNoteRequest) FromJSON(r io.Reader) error {
	return json.NewDecoder(r).Decode(i)
}
