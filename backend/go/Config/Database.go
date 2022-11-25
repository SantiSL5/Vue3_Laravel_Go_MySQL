package Config

import (
	"fmt"
	"os"
	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
	"strconv"
)

var DB *gorm.DB

// DBConfig represents db configuration
type DBConfig struct {
	Host     string
	Port     int
	User     string
	DBName   string
	Password string
}

func BuildDBConfig() *DBConfig {
	
	godotenv.Load();

	fmt.Println("asdad")


	s := os.Getenv("DB_PORT")

    v, err := strconv.Atoi(s)
	if err != nil{
		fmt.Println(err)
	}
    
	// if err != nil {
    //     return 0, err
	// }
	
	dbConfig := DBConfig{
		Host:     os.Getenv("DB_HOST"),
		Port:     v,
		User:     os.Getenv("DB_USERNAME"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   os.Getenv("DB_DATABASE"),
	}
	return &dbConfig
}

func DbURL(dbConfig *DBConfig) string {
	return fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
		dbConfig.User,
		dbConfig.Password,
		dbConfig.Host,
		dbConfig.Port,
		dbConfig.DBName,
	)
}
