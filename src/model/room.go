package model

// RoomMgr 模型
type RoomMgr struct {
	RoomID      int64      `gorm:"primary_key"`
	Users       []UserMgr  `gorm:"many2many:room_users;"`
	Comets      []CometMgr `gorm:"many2many:room_comets;"`
	OnlineCount int        `gorm:"not null"`
}

// // RoomComet 模型
// type RoomComet struct {
// 	RoomID    int64  `gorm:"primary_key"`
// 	CometAddr string `gorm:"primary_key;type:varchar(255)"`
// }

// // RoomUserid 模型
// type RoomUser struct {
// 	RoomID int64 `gorm:"primary_key"`
// 	UserID int64 `gorm:"primary_key"`
// }
