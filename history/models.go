package history

import (
	"time"
)

type Record struct {
	Action      uint32
	Note        string
	Date        time.Time
	User        uint32 `gorm:"column:username"`
	StringUser  string `gorm:"column:stringuser"`
	Property_id uint32
}

func (Record) TableName() string {
	return "property_history"
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

type Property struct {
	Id        uint32
	Inventory string
	Serial    string
	Name      string
	State     uint32
}

func (Property) TableName() string {
	return "fixed_property"
}

type Group struct {
	Id          uint32
	Name        string
	Description string
	Username    uint32
	CreatedAt   time.Time
}

func (Group) TableName() string {
	return "groups"
}

type Warehouse struct {
	Id   uint32
	Name string
}

func (Warehouse) TableName() string {
	return "warehouses"
}

type Employee struct {
	Id         uint32
	Table      string
	Name       string
	Manager    string
	Department string
	Job        string
	CreatedAt  time.Time `gorm:"column:created_at"`
}

func (Employee) TableName() string {
	return "staff"
}
