package database

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
	log "github.com/sirupsen/logrus"
	"jea-api/environment"
)

type ConnectionInfo struct {
	Host		string
	Port		string
	User		string
	Database	string
	Password	string
}

func NewDatabase(info ConnectionInfo) *gorm.DB {
	gorm.DefaultTableNameHandler = environment.TableNameHandler
	db, err := gorm.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable", info.Host, info.Port, info.User, info.Database, info.Password))
	if err != nil {
		log.Fatal("Error in connect with database. Error", err.Error())
	}
	return db
}

func UseDatabase(db *gorm.DB) func (ctx *gin.Context){
	return func(ctx *gin.Context) {
		ctx.Set("db", db)
		ctx.Next()
	}
}

func GetDatabase(ctx *gin.Context) *gorm.DB {
	db, exists := ctx.Get("db")
	database, ok := db.(*gorm.DB)
	if ok && exists{
		return database
	}
	return nil
}
