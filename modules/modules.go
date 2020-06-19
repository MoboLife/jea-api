package modules

import (
	"jea-api/controller"
	"reflect"

	"github.com/gin-gonic/gin"
)

// Build build struct for create ginController
func Build(group *gin.RouterGroup, modules interface{}) {
	var modulesType = reflect.TypeOf(modules)
	if modulesType.Kind() == reflect.Ptr {
		modulesType = modulesType.Elem()
	}
	for i := 0; i < modulesType.NumField(); i++ {
		var field = modulesType.Field(i)
		if field.Type.Kind() != reflect.Struct {
			continue
		}
		var secure = true
		if secureTag, ok := field.Tag.Lookup("secure"); ok && secureTag == "false" {
			secure = false
		}
		if router, ok := field.Tag.Lookup("router"); ok {
			controller.NewGinControllerWrapper(group.Group(router), controller.NewGinController(reflect.New(field.Type).Interface()), secure)
		}
	}

}
