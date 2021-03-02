package models

import (
	"time"

	_ "github.com/jinzhu/gorm"
)

type Planet struct {
	ID       uint      `json:"id" gorm:"primary_key"`
	Name     string    `json:"name"`
	Climate  string    `json:"climate"`
	Terrain  string    `json:"terrain"`
	Films    int       `json:"films"`
	CreateAt time.Time `json:"create_at"`
}
