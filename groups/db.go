package groups

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
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

func GetGroups(name string) ([]*Group, error) {
	groups := make([]*Group, 0)

	err := db.
		Table("groups as g").
		Select("g.id", "g.name", "g.description", "g.created_at", "u.id as username",
			"u.display_name as display").
		Joins("LEFT JOIN users as u ON g.username = u.id").
		Where("g.name LIKE ?", name+"%").
		Order("g.name ASC").
		Find(&groups).Error

	return groups, err
}

func CreateGroup(group *Group, username string) error {
	user := User{}
	err := db.Where("username = ?", username).First(&user).Error

	if err != nil {
		return err
	}

	group.Username = user.Id

	err = db.Select("name", "description", "username").Create(group).Error

	return err
}

func RemoveGroup(group *Group) error {
	err := db.Where("property_id = ?", group.Id).Delete(&PropGroup{}).Error

	if err != nil {
		return err
	}

	err = db.Delete(group).Error

	return err

}

func AddToGroup(props []*PropGroup) error {
	return db.Select("property_id", "group_id").Create(&props).Error
}

func RemoveFromGroup(props []*PropGroup) error {
	for _, prop := range props {
		err := db.
			Where("group_id = ?", prop.GroupId).
			Where("property_id = ?", prop.PropertyId).
			Delete(&PropGroup{}).
			Error
		if err != nil {
			return err
		}
	}
	return nil
}

func IsInGroup(prop *PropGroup) (*PropGroup, error) {
	err := db.Where("property_id = ?", prop.PropertyId).
		Where("group_id = ?", prop.GroupId).First(prop).Error
	return prop, err
}
