package staff

import "time"

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

type Property struct {
	Id        uint32
	Inventory string
	Serial    string
	Name      string
	CreatedAt time.Time `gorm:"column:created_at"`
	Warehouse uint32
	State     uint32
	Username  string
	GivenAt   time.Time `gorm:"column:given"`
	Record    uint32
}

func (Property) TableName() string {
	return "fixed_property"
}

type PropRecord struct {
	Id         uint32
	PropertyId uint32    `gorm:"column:property_id"`
	StaffId    uint32    `gorm:"column:staff_id"`
	UserId     uint32    `gorm:"column:user_id"`
	CreatedAt  time.Time `gorm:"column:created_at"`
}

func (PropRecord) TableName() string {
	return "staff_property"
}

type User struct {
	Id       uint32
	Username string
}

func (User) TableName() string {
	return "users"
}

type SearchStaff struct {
	Search      string
	TableNumber string
	Department  string
	Manager     string
	Job         string
	Offset      uint32
	Limit       uint32
	Order       string
}
