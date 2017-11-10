package row

import (
	"time"
)

type wantedlyUser struct {
	Id        int       `json:"id" xorm:"'id'"`
	Name      string    `json:"name" xorm:"'name'"`
	Email     string    `json:"email" xorm:"'email'"`
	CreatedAt time.Time `json:"created_at" xorm:"'created_at'"`
	UpdatedAr time.Time `json:"updated_ar" xorm:"'updated_at'"`
}

type Users []wantedlyUser
