package models

import (
	"database/sql"
	"time"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct{
	Id         string   `gorm:"unique;primaryKey"`
	Name       *sql.NullString	
	Email      string     `gorm:"unique"`
	Password   string
	CreatedAt      time.Time
	UpdatedAt      time.Time
	CreatedRooms   []Rooms      `gorm:"foreignKey:CreatedBy"`
	JoinedRooms    []Rooms      `gorm:"many2many:user_joined_room;"`
}

type VerificationToken struct{
	Id  			string        `gorm:"type:uuid;primaryKey"`
	EmailIdentifier string  
	Token           string
	CreatedAt 		time.Time
	ExpiresAt 		time.Time
}

type Session struct{
	Id              	string      `gorm:"primaryKey"`
	AccountProvider 	string
	SignInAt    	  time.Time
	SignedOutAt 	  time.Time
}

type Maps struct{
	Id 				string
	Image          	string
	Tiles	        string
	Info		    string
	ExistedFrom    	time.Time
	Rooms          	[]Rooms   `gorm:"foreignKey:map_id"`  
}

type Avatars struct{
	Id             string
	Name           string
	Image          string
	ExistedFrom    time.Time
	MetaUsers      []MetaUsers `gorm:"foreignKey:UserAvatarId"`
}

type Rooms struct{
	Id           string
	Name         string
	CreatedBy    string
	MapId        string
	Map          Maps        `gorm:"references:Id; foreignKey:MapId"`
	UsersJoined  []User      `gorm:"many2many:user_joined_room;"`
	MetaUsers    []MetaUsers `gorm:"foreignKey:room_id"`

}

type MetaUsers struct  {
	Id           string
	Name         string
	UserAvatarId string
	UserAvatar   Avatars   		`gorm:"foreignKey:UserAvatarId; references:Id; "`
	UserId		 string   
	RoomId       string 
	Room         Rooms          `gorm:"references:Id; foreignKey:RoomId"`    
	Position     PlayerPosition `gorm:"foreignKey:meta_users_id"`
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

func (u *Avatars) BeforeCreate(tx *gorm.DB) (err error){
	u.Id = uuid.New().String()
	return 
} 
func (u *Rooms) BeforeCreate(tx *gorm.DB) (err error){
	u.Id = uuid.New().String()
	return 
} 
func (u *MetaUsers) BeforeCreate(tx *gorm.DB) (err error){
	u.Id = uuid.New().String()
	return 
} 
func (u *Maps) BeforeCreate(tx *gorm.DB) (err error){
	u.Id = uuid.New().String()
	return 
} 

func (u *PlayerPosition) BeforeCreate(tx *gorm.DB) (err error){
	u.Id = uuid.New().String()
	return 
} 