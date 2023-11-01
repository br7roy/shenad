package opts

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"log"
)

var DB *gorm.DB

func InitDB() {
	var err error
	DB, err = gorm.Open("sqlite3", Cfg.App.Db.DbFilePath)
	if err != nil {
		log.Panicln("err:", err.Error())
	}
	DB.SingularTable(true)
}

type User struct {
	ID           int    `json:"id",gorm:"primary_key"`
	Email        string `json:"email" binding:"email"`
	Password     string `json:"password"`
	Username     string `json:"username"`
	Token        string `json:"token"`
	Avatar       string `json:"avatar"`
	Introduction string `json:"introduction"`
	Name         string `json:"name"`
	Roleid       string `json:"roleid"`
}

type FoodTrucks struct {
	Name      string `json:"name"`
	Status    string `json:"status"`
	Latitude  string `json:"latitude"`
	Longitude string `json:"longitude"`
}

func (user *User) QueryByEntry(userName string, password string) (User, error) {
	e := DB.Find(user, "username=? and password=?", userName, password).Error
	return *user, e
}

func (user *User) QueryByToken(token string) (User, error) {
	e := DB.Find(user, "token=?", token).Error
	return *user, e
}

func (user *User) UpdateTokenByUser() {
	update := DB.Model(&user).Where("id = ?", user.ID).Update("token", user.Token)
	println(update)
}

func (user *User) ClearTokenByUser() {
	DB.Model(&user).Update("token", "")
}

func (foodTrucks *FoodTrucks) GetFoodTrucks() ([]FoodTrucks, error) {
	var foodTruckList []FoodTrucks
	e := DB.Raw("select Applicant as name, Status as status ,Latitude as latitude," +
		"Longitude as longitude  from food_trucks").Scan(&foodTruckList)
	return foodTruckList, e.Error

}
