package domain

import (
	"errors"
)

type BaseRequest struct {
	Username string `json:"username"`
}

type BaseItem struct {
	ID     int16  `json:"id"`
	Nama   string `json:"nama"`
	Link   string `json:"link"`
	IDMain int16  `json:"id_main"`
	Urut   int16  `json:"urut"`
}

type NewBaseItem struct {
	ID       int16          `json:"id"`
	IDMain   int16          `json:"id_main"`
	Urut     int16          `json:"urut"`
	Label    string         `json:"label"`
	Href     string         `json:"href,omitempty"`
	Children []*NewBaseItem `json:"children,omitempty"`
}

type BaseRepository interface {
	GetMenu(role string) ([]*BaseItem, error)
}

var ErrInvalidCredentials = errors.New("invalid credentials")
