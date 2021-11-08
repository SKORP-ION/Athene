package property

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"strings"
)

var db *gorm.DB

func connectDB(user string, password string, host string, port string, database string) error {
	dbURi := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		host, user, password, database, port)

	//dialect := postgres.Open(dbURi)
	dialect := postgres.New(postgres.Config{
		DSN: dbURi,
	})

	conn, err := gorm.Open(dialect, &gorm.Config{})

	if err != nil {
		return err
	}
	db = conn
	return nil
}

func ParseSearchParams(query *gorm.DB, params *SearchParams) *gorm.DB {
	if params.Serial != "" {
		query = query.Where("fp.serial LIKE ?", "%"+strings.ToUpper(params.Serial)+"%")
	}
	if params.Inventory != "" {
		query = query.Where("fp.inventory LIKE ?", "%"+strings.ToUpper(params.Inventory)+"%")
	}
	if params.Name != "" {
		query = query.Where("UPPER(fp.name) LIKE ?", "%"+strings.ToUpper(params.Name)+"%")
	}
	if params.Action != 0 {
		query = query.Where("fp.state = ?", params.Action)
	}
	if params.Warehouse != 0 {
		query = query.Where("fp.warehouse = ?", params.Warehouse)
	}

	if len(params.Groups) > 0 {
		newParams := make([]string, 0)
		for _, param := range params.Groups {
			if param != "" {
				newParams = append(newParams, strings.ToLower(param))
			}
		}
		if len(newParams) > 0 {
			query = query.Where("LOWER(g.name) in ?", newParams)
		}
	}

	//Такая тупая проверка для защиты от SQL инъекций
	switch params.Order {
	case "serial0":
		query = query.Order("serial asc")
	case "serial1":
		query = query.Order("serial desc")
	case "inventory0":
		query = query.Order("inventory asc")
	case "inventory1":
		query = query.Order("inventory desc")
	case "name0":
		query = query.Order("name asc")
	case "name1":
		query = query.Order("name desc")
	case "warehouse0":
		query = query.Order("warehouse asc")
	case "warehouse1":
		query = query.Order("warehouse desc")
	case "created_at0":
		query = query.Order("created_at asc")
	case "created_at1":
		query = query.Order("created_at desc")
	case "updated_at0":
		query = query.Order("updated_at asc")
	case "updated_at1":
		query = query.Order("updated_at desc")
	default:
		query = query.Order("updated_at desc")
	}
	query = query.Offset(int(params.Offset))
	query = query.Limit(int(params.Limit))
	return query

}

func GetProperty(params *SearchParams) ([]*Property, error) {
	properties := make([]*Property, 0)
	query := db.Table("fixed_property as fp").
		Distinct("fp.id").
		Select("fp.id", "fp.inventory", "fp.serial", "fp.name",
			"fp.created_at", "fp.warehouse",
			"ph.date as updated_at", "fp.state as action").
		Joins("LEFT JOIN (SELECT DISTINCT ON (property_id) date, " +
			"action, property_id FROM property_history ORDER BY property_id, date DESC) AS ph ON fp.id = ph.property_id").
		Joins("LEFT JOIN property_groups AS pg On pg.property_id = fp.id").
		Joins("LEFT JOIN groups AS g ON g.id = pg.group_id")
	query = ParseSearchParams(query, params)
	query = query.Find(&properties)
	err := query.Error

	if err != nil {
		return nil, err
	}

	return properties, err
}

func GetWarehouses() ([]Warehouse, error) {
	warehouses := make([]Warehouse, 0)

	err := db.Find(&warehouses).Error

	if err != nil {
		return nil, err
	}

	return warehouses, nil
}

func GetCount(params *SearchParams) (int64, error) {

	/*
		query := db.Table("fixed_property as fp").
			Distinct("fp.id").
			Select("fp.id", "fp.inventory", "fp.serial", "fp.name",
				"fp.created_at","fp.warehouse",
				"ph.date as updated_at", "fp.state as action").
			Joins("LEFT JOIN (SELECT DISTINCT ON (property_id) date, " +
				"action, property_id FROM property_history ORDER BY property_id, date DESC) AS ph ON fp.id = ph.property_id").
			Joins("LEFT JOIN property_groups AS pg On pg.property_id = fp.id").
			Joins("LEFT JOIN groups AS g ON g.id = pg.group_id")
	*/

	query := db.Table("fixed_property as fp").
		Select("COUNT(DISTINCT(fp.id))").
		Joins("LEFT JOIN (SELECT DISTINCT ON (property_id) date, " +
			"action, property_id FROM property_history ORDER BY property_id, date DESC) AS ph ON fp.id = ph.property_id").
		Joins("LEFT JOIN property_groups AS pg On pg.property_id = fp.id").
		Joins("LEFT JOIN groups AS g ON g.id = pg.group_id").
		Order("fp.id")

	query = ParseSearchParams(query, params)
	count := int64(0)

	err := query.Model(&Property{}).Count(&count).Error

	return count, err
}

func GetActions() ([]*Action, error) {
	actions := make([]*Action, 0)

	err := db.Order("id asc").Find(&actions).Error

	if err != nil {
		return nil, err
	}

	return actions, err
}

func GetOneProperty(id uint32) (*Property, error) {
	prop := &Property{Id: id}

	query := db.Table("fixed_property as fp").
		Select("fp.id", "fp.inventory", "fp.serial", "fp.name",
			"fp.created_at", "fp.warehouse",
			"ph.date as updated_at", "fp.state as action").
		Joins("LEFT JOIN (SELECT DISTINCT ON (property_id) date, " +
			"action, property_id FROM property_history ORDER BY property_id, date DESC) AS ph ON fp.id = ph.property_id")

	err := query.Last(&prop).Error

	if err != nil {
		return nil, err
	}

	return prop, nil
}

func IsOnWarehouse(prop *Property) (*Property, error) {
	err := db.
		Where("id = ?", prop.Id).
		First(prop).Error

	return prop, err
}

func ChangeWarehouse(props []*Property) error {
	for _, prop := range props {
		err := db.Model(prop).Select("warehouse").Updates(prop).Error
		if err != nil {
			return err
		}
	}
	return nil
}
