package row

import (
	"time"
)

type WantedlyUser struct {
	Id        int       `json:"id" xorm:"'id'"`
	Name      string    `json:"name" xorm:"'name'"`
	Email     string    `json:"email" xorm:"'email'"`
	CreatedAt time.Time `json:"created_at" xorm:"'created_at'"`
	UpdatedAr time.Time `json:"updated_ar" xorm:"'updated_at'"`
}

func (WantedlyUser) TableName() string {
	return "wantedlyusers"
}

//type WantedlyUsers []WantedlyUser
