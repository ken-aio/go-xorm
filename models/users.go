package models

import (
	"time"
)

type Users struct {
	Id        int64     `xorm:"pk autoincr BIGINT(20)"`
	Name      string    `xorm:"not null VARCHAR(255)"`
	Birthdate time.Time `xorm:"DATE"`
	Gender    string    `xorm:"not null VARCHAR(10)"`
	CreatedAt time.Time `xorm:"not null DATETIME"`
	CreatedBy string    `xorm:"not null VARCHAR(64)"`
	UpdatedAt time.Time `xorm:"not null DATETIME"`
	UpdatedBy string    `xorm:"not null VARCHAR(64)"`
}

func (u *Users) BeforeInsert() {
	u.CreatedAt = time.Now()
	u.UpdatedAt = time.Now()
	u.CreatedBy = "dummy"
	u.UpdatedBy = "dummy"
}

func (u *Users) BeforeUpdate() {
	u.UpdatedAt = time.Now()
	u.UpdatedBy = "dummy"
}
