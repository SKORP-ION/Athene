package history

import (
	. "Athena/log"
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

func CreateCard(username, note string, properties []Property) ([]Property, error) {
	err := db.Create(&properties).Error

	if err != nil {
		return nil, err
	}

	user := User{}

	err = db.Select("Id").Where("username = ?", username).First(&user).Error

	if err != nil {
		return nil, err
	}

	records := make([]Record, 0)

	for _, prop := range properties {
		records = append(records, Record{
			Action: 1, //Приход
			Note: fmt.Sprintf("Создана карточка %s - %s - %s. \nЗаметка: %s",
				prop.Serial, prop.Inventory, prop.Name, note),
			User:        user.Id,
			Property_id: prop.Id,
		})
	}

	err = db.Select("Action", "Note", "User", "Property_id").Create(&records).Error

	if err != nil {
		return nil, err
	}

	return properties, nil
}

func Archive(username, note string, properties []Property) error {
	err := db.Select("archived").Save(&properties).Error

	if err != nil {
		return err
	}

	user := User{}

	err = db.Select("Id").Where("username = ?", username).First(&user).Error

	if err != nil {
		return err
	}

	records := make([]Record, 0)

	for _, prop := range properties {
		records = append(records, Record{
			Action:      11, //Отправлено в архив
			Note:        note,
			User:        user.Id,
			Property_id: prop.Id,
		})
	}

	err = db.Create(&records).Error

	if err != nil {
		return err
	}

	return nil
}

func DoAction(username, note string, properties []*Property, action uint32, needUpdate bool) error {
	user, err := getUser(username)

	if err != nil {
		return err
	}

	records := make([]*Record, 0)

	for _, prop := range properties {
		records = append(records, &Record{
			Action:      action,
			Note:        note,
			User:        user.Id,
			Property_id: prop.Id,
		})
	}

	err = db.Select("action", "note", "username", "property_id").Create(&records).Error

	if err != nil {
		return err
	}

	if needUpdate {
		for _, prop := range properties {
			err = db.Model(prop).Select("state").Updates(prop).Error

			if err != nil {
				return err
			}
		}
	}

	return nil
}

func ChangeField(prop *Property, field string) error {
	err := db.Model(prop).Select(field).Updates(prop).Error

	return err
}

//Проверки
func GetPropertyBySerial(serial string) (Property, error) {
	prop := Property{}
	err := db.Where("serial = ?", strings.ToUpper(serial)).Last(&prop).Error

	return prop, err
}

func GetLastStateBySerial(serial string) (Record, error) {
	record := Record{}

	err := db.Table("property_history as ph").
		Select("ph.action", "ph.note", "ph.date").
		Joins("LEFT JOIN fixed_property as fp ON ph.property_id = fp.id").
		Where("fp.serial = ?", strings.ToUpper(serial)).
		Order("ph.id DESC").
		First(&record).Error

	return record, err
}

func GetHistory(id uint32) ([]*Record, error) {
	records := make([]*Record, 0)

	err := db.Table("property_history as ph").
		Select("ph.action", "ph.note", "ph.date", "u.display_name as stringuser").
		Joins("LEFT JOIN users as u ON ph.username = u.id").
		Where("property_id = ?", id).
		Order("ph.id asc").
		Find(&records).Error

	if err != nil {
		Error.Println(err)
		return nil, err
	}

	return records, nil
}

//Асинхронные функции
func asyncGetUser(username string, userChan chan User, errorChan chan error) {
	user := User{}

	err := db.Where("username = ?", username).First(&user).Error

	if err != nil {
		errorChan <- err
		return
	}

	userChan <- user
	errorChan <- nil
}

func getUser(username string) (User, error) {
	user := User{}

	err := db.Where("LOWER(username) = ?", strings.ToLower(username)).First(&user).Error

	return user, err
}

//Вспомогательные функции

func GetGroup(id uint32) (Group, error) {
	group := Group{Id: id}

	err := db.First(&group).Error
	return group, err
}

func GetWarehouse(id uint32) (Warehouse, error) {
	warehouse := Warehouse{Id: id}

	err := db.First(&warehouse).Error

	return warehouse, err
}

func GetEmployee(id uint32) (Employee, error) {
	emp := Employee{Id: id}

	err := db.First(&emp).Error

	return emp, err
}
