package models

import (
	"database/sql"
	"time"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct{
	Id         string   `gorm:"type:uuid;primaryKey"`
	Name      *sql.NullString	
	Email     string
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
	CreatedRooms   []Rooms      `gorm:"foreignKey:createdBy"`
	JoinedRooms    []Rooms      `gorm:"many2many:user_joined_room;"`
}

type VerificationToken struct{
	Id  string        `gorm:"type:uuid;primaryKey"`
	EmailIdentifier string  
	Token           string
	CreatedAt time.Time
	ExpiresAt time.Time
}

type Session struct{
	Id              string      `gorm:"primaryKey"`
	AccountProvider string
	SignInAt    	  time.Time
	SignedOutAt 	  time.Time
}

type Maps struct{
	Id 			string
	Image          string
	Tiles	        string
	Info		    string
	ExistedFrom    time.Time
}

type Avatars struct{
	Id             string
	Name           string
	Image          string
	ExistedFrom    time.Time
	MetaUserId     string
}

type Rooms struct{
	Id           string
	Name         string
	CreatedBy    string
	MapId        string
	Maps         Maps        `gorm:"references:Id"`
	UsersJoined  []User      `gorm:"many2many:user_joined_room;"`
	MetaUsers    []MetaUsers `gorm:"foreignKey:roomId"`

}

type MetaUsers struct{
	Id          string
	Name        string
	UserAvatar  Avatars   `gorm:"references:meta_user_id"`
	UserId		 string   
	RoomId      string     
	Position    PlayerPosition `gorm:"references:meta_users_id"`
}

type PlayerPosition struct{
	Id          string
	X_cordinate float64
	Y_cordinate float64
	MetaUsersId string
}


func (u *User) BeforeCreate(tx *gorm.DB) (err error){
	u.Id = uuid.New().String()
	return 
} 

func (u *VerificationToken) BeforeCreate(tx *gorm.DB) (err error){
	u.Id = uuid.New().String()
	return 
} 
