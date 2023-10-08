package player

import "gorm.io/gorm"

type (
	CommonRequest struct {
		Id uint `uri:"player_id"`
	}
)

type Images struct {
	gorm.Model
	PlayerID   uint   `json:"PlayerID"  gorm:"unique;not null;column:player_id"`
	Avatar     string `json:"Avatar"`
	Background string `json:"Background"`
}

func (Images) TableName() string {
	return "player_images"
}
