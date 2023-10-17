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

type CommonPlayerRequest struct {
	ID uint `json:"ID"`

	Name       string   `json:"Name"`                 // 选手名
	WcaID      string   `json:"WcaID,omitempty"`      // 选手WcaID，用于查询选手WCA的成绩
	ActualName string   `json:"ActualName,omitempty"` // 真实姓名
	TitlesVal  []string `json:"TitlesVal,omitempty"`  // 头衔
	LoginID    string   `json:"LoginID"`              // 登录自定义ID
	QQ         string   `json:"QQ"`                   // qq号
	WeChat     string   `json:"WeChat"`               // 微信号
	Phone      string   `json:"Phone"`                // 手机号
}
