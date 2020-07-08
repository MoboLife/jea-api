package database

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type Connection struct {
	ConnectionInfo
	EID				string		`json:"eid"`
}

type Configuration struct {
	Connections		[]Connection
}

type Farm interface {
	Open() error
	GetConnections() map[string]*gorm.DB
	GetConnection(eid string) *gorm.DB
}

type Context struct {
	Configuration Configuration
	Connections   map[string]*gorm.DB

}

func (c *Context) GetConnection(eid string) *gorm.DB {
	return c.Connections[eid]
}

func (c *Context) Open() error {
	for _, connection := range c.Configuration.Connections {
		db, err := NewDatabase(connection.ConnectionInfo, false)
		if err != nil {
			return err
		}
		c.Connections[connection.EID] = db
	}
	return nil
}

func (c *Context) GetConnections() map[string]*gorm.DB {
	return c.Connections
}

func NewFarm(configuration Configuration) Farm {
	return &Context{Configuration: configuration}
}

func UseFarm(farm Farm) func(c *gin.Context) {
	return func(c *gin.Context) {
		c.Set("farm", farm)
	}
}

func GetFarm(c *gin.Context) Farm {
	value, exists := c.Get("farm")
	if exists {
		return value.(Farm)
	}
	return nil
}

func GetConnection(eid string, c *gin.Context) *gorm.DB {
	var farm = GetFarm(c)
	if farm == nil {
		return nil
	}
	return farm.GetConnection(eid)
}
