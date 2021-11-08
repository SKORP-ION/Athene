package property

import "time"

type Property struct {
	Id         uint32
	Inventory  string
	Serial     string
	Name       string
	Created_at time.Time
	Updated_at time.Time
	Warehouse  uint32
	Action     uint32
	State      uint32
}

func (Property) TableName() string {
	return "fixed_property"
}

type Warehouse struct {
	Id   uint32
	Name string
}

func (Warehouse) TableName() string {
	return "warehouses"
}

type Action struct {
	Id     uint32
	Name   string `gorm:"column:state"`
	Action string
}

func (Action) TableName() string {
	return "actions"
}

type SearchParams struct {
	Inventory string
	Serial    string
	Name      string
	Action    uint32
	Warehouse uint32
	Offset    uint32
	Limit     uint32
	Order     string
	Groups    []string
}
