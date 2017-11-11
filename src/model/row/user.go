package row

import (
	"time"
)

type WantedlyUser struct {
	Id        int       `json:"id" xorm:"not null pk autoincr INT(10)"`
	Name      string    `json:"name" xorm:"not null VARCHAR(255)"`
	Email     string    `json:"email" xorm:"not null VARCHAR(255)"`
	CreatedAt time.Time `json:"created_at" xorm:"<-"`
	UpdatedAt time.Time `json:"updated_at" xorm:"<-"`
}

func (WantedlyUser) TableName() string {
	return "wantedlyusers"
}
