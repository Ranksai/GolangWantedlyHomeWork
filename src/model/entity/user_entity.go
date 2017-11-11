package entity

import "GolangWantedlyHomeWork/src/model/row"

type WantedlyUser struct {
	row.WantedlyUser `xorm:"extends"`
}

type WantedlyUsers []WantedlyUser
