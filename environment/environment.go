package environment

import (
	"fmt"
	"github.com/jinzhu/gorm"
)

func UseEnvironment(eid string, db *gorm.DB) *gorm.DB {
	return db.Set("environment", eid)
}

func TableNameHandler(db *gorm.DB, defaultTableName string) string {
	if envConfig, ok := db.Get("environment"); ok{
		schema, ok := envConfig.(string)
		if ok && schema != ""{
			return fmt.Sprintf("%s.%s", schema, defaultTableName)
		}
	}
	return fmt.Sprintf("public.%s", defaultTableName)
}