package staff

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"strings"
	"time"
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

func GetStaff(params *SearchStaff) ([]*Employee, error) {
	staff := make([]*Employee, 0)

	query := db.Table("staff")

	query = ParseSearchParams(query, params)

	err := query.Find(&staff).Error

	return staff, err
}

func GetEmployeeById(id uint32) (*Employee, error) {
	emp := Employee{}

	err := db.Where("id = ?", id).Last(&emp).Error

	return &emp, err
}

func GetCount(params *SearchStaff) (uint32, error) {
	count := int64(0)

	query := db.Table("staff")

	query = ParseSearchParams(query, params)

	err := query.Count(&count).Error

	return uint32(count), err
}

func CreateEmployee(employee Employee) (Employee, error) {
	employee.CreatedAt = time.Now()
	err := db.Select("table", "name", "manager", "department", "job").
		Create(&employee).Error
	return employee, err
}

func GetEmployeesProp(emp Employee) ([]*Property, error) {
	props := make([]*Property, 0)
	query := db.Table("fixed_property AS fp").
		Select("fp.id", "fp.inventory", "fp.serial", "fp.name", "fp.created_at", "fp.warehouse", "fp.state",
			"u.username", "sp.created_at AS given", "sp.id as record").
		Joins("LEFT JOIN staff_property AS sp ON sp.property_id = fp.id").
		Joins("LEFT JOIN users AS u ON u.id = sp.user_id").
		Where("sp.staff_id = ?", emp.Id).
		Find(&props)
	err := query.Error

	return props, err
}

func GiveToEmployee(empId, usrId uint32, props []*Property) error {

	for _, prop := range props {
		err := db.Select("property_id", "staff_id", "user_id").
			Create(&PropRecord{
				PropertyId: prop.Id,
				StaffId:    empId,
				UserId:     usrId,
			}).Error

		if err != nil {
			return err
		}

	}

	return nil
}

func TakeFromEmployee(empId uint32, props []*Property) error {

	for _, prop := range props {
		err := db.Where("staff_id = ?", empId).
			Where("property_id = ?", prop.Id).
			Delete(&PropRecord{}).Error

		if err != nil {
			return err
		}

	}

	return nil
}

func GetUser(username string) (User, error) {
	user := User{}

	err := db.Where("username = ?", username).First(&user).Error

	return user, err
}

func ParseSearchParams(query *gorm.DB, params *SearchStaff) *gorm.DB {
	if params.Search != "" {
		query = query.Where("UPPER(name) LIKE ?", "%"+strings.ToUpper(params.Search)+"%")
	}
	if params.TableNumber != "" {
		query = query.Where("UPPER(table) LIKE ?", "%"+strings.ToUpper(params.TableNumber)+"%")
	}
	if params.Department != "" {
		query = query.Where("UPPER(department) LIKE ?", "%"+strings.ToUpper(params.Department)+"%")
	}
	if params.Manager != "" {
		query = query.Where("UPPER(manager) LIKE ?", "%"+strings.ToUpper(params.Manager)+"%")
	}
	if params.Job != "" {
		query = query.Where("UPPER(job) LIKE ?", "%"+strings.ToUpper(params.Job)+"%")
	}

	//Такая тупая проверка для защиты от SQL инъекций
	switch params.Order {
	case "name0":
		query = query.Order("name asc")
	case "name1":
		query = query.Order("name desc")
	case "table0":
		query = query.Order("table asc")
	case "table1":
		query = query.Order("table desc")
	case "department0":
		query = query.Order("department asc")
	case "department1":
		query = query.Order("department desc")
	case "manager0":
		query = query.Order("manager asc")
	case "manager1":
		query = query.Order("manager desc")
	case "job0":
		query = query.Order("job asc")
	case "job1":
		query = query.Order("job desc")
	case "created_at0":
		query = query.Order("created_at asc")
	case "created_at1":
		query = query.Order("created_at desc")
	default:
		query = query.Order("name asc")
	}
	query = query.Offset(int(params.Offset))
	query = query.Limit(int(params.Limit))
	return query

}

func GetRecord(record PropRecord) (PropRecord, error) {
	err := db.Where("property_id = ?", record.PropertyId).Last(&record).Error
	return record, err
}

func GetRecordById(recordId uint32) (*PropRecord, error) {
	record := PropRecord{}

	err := db.Where("id = ?", recordId).Last(&record).Error

	return &record, err
}

func GetProperty(propId uint32) (*Property, error) {
	prop := Property{}

	err := db.Where("id = ?", propId).Last(&prop).Error

	return &prop, err
}
