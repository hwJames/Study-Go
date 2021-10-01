package models

import (
	"time"
)

type User struct {
	NO uint `gorm:"primaryKey;comment:유저 번호;"`
	ID string `gorm:"unique;size:50;comment:유저 아이디(이메일);"`
	Nick string `gorm:"size:20;comment:유저 닉네임;"`
	Pwd string `gorm:"size:255;comment:유저 비밀번호;"`
	CreatedAt time.Time `gorm:"comment:유저 등록 날짜;"`
	UpdatedAt time.Time `gorm:"comment:유저 업데이트 날짜;"`
  }