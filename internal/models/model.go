package models

import (
	"database/sql"
	"time"
	"github.com/google/uuid"
	"gorm.io/gorm"
)



type User struct{
	Id             string   	`gorm:"unique;primaryKey"`
	Name      	   *sql.NullString	
	Email      	   string     	`gorm:"unique"`
	Password   	   string
	CreatedAt      time.Time
	UpdatedAt      time.Time
	CreatedRooms   []Rooms      `gorm:"foreignKey:CreatedBy; constraint:OnDelete:CASCADE;"`
	JoinedRooms    []Rooms      `gorm:"many2many:user_joined_room;"`
}

type VerificationToken struct{
	Id  			string        `gorm:"unique;primaryKey"`
	EmailIdentifier string  
	Token           string
	CreatedAt 		time.Time
	ExpiresAt 		time.Time
}

type Session struct{
	Id              	string      `gorm:"unique;primaryKey"`
	AccountProvider 	string
	SignInAt    	  time.Time
	SignedOutAt 	  time.Time
}

type Maps struct{
	Id 				string    `gorm:"unique;primaryKey"`
	Image          	string
	Tiles	        string
	Info		    string
	ExistedFrom    	time.Time
	Rooms          	[]Rooms   `gorm:"foreignKey:map_id; constraint:OnDelete:CASCADE;"`  
}

type Avatars struct{
	Id             string    `gorm:"unique;primaryKey"`
	Name           string
	Image          string
	ExistedFrom    time.Time
	MetaUsers      []MetaUsers `gorm:"foreignKey:UserAvatarId; constraint:OnDelete:CASCADE;"`
}

type Rooms struct{
	Id           string      `gorm:"unique;primaryKey"`
	Name         string
	CreatedBy    string
	CreatedUser  User        `gorm:"references:Id; foreignKey:CreatedBy;"`
	MapId        string
	Map          Maps        `gorm:"references:Id; foreignKey:MapId;"`
	UsersJoined  []User      `gorm:"many2many:user_joined_room;"`
	MetaUsers    []MetaUsers `gorm:"foreignKey:RoomId; constraint:OnDelete:CASCADE;"`

}

type MetaUsers struct  {
	Id           string         `gorm:"unique;primaryKey"`
	Name         string
	UserAvatarId string
	UserAvatar   Avatars   		`gorm:"foreignKey:UserAvatarId; references:Id; "`
	UserId		 string   
	RoomId       string 
	Room         Rooms          `gorm:"references:Id; foreignKey:RoomId"`    
	Position     PlayerPosition `gorm:"foreignKey:meta_users_id;constraint:OnDelete:CASCADE;"`
}

type PlayerPosition struct{
	Id          string       `gorm:"unique"`
	X_cordinate string
	Y_cordinate string
	MetaUsersId string       `gorm:"unique;primaryKey"`

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