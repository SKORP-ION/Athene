package groups

import "time"

type Group struct {
	Id          uint32
	Name        string
	Description string
	CreatedAt   time.Time
	Username    uint32
	Display     string
}

func (Group) TableName() string {
	return "groups"
}

type PropGroup struct {
	Id         uint32
	PropertyId uint32 `gorm:"column:property_id"`
	GroupId    uint32 `gorm:"column:group_id"`
}

func (PropGroup) TableName() string {
	return "property_groups"
}

type User struct {
	Id          uint32
	Username    string
	Password    string
	Description string
	Created_at  time.Time
	Active      bool
	Role        uint32
}

func (User) TableName() string {
	return "users"
}
