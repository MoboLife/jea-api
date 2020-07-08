package repository

import (
	"fmt"
	"jea-api/common"
	"strconv"

	"github.com/gin-gonic/gin"
)

// Filter base filter
type Filter interface {
	Apply(ctx *gin.Context) []Options
}

// APIFilter api Filter builder
type APIFilter struct {
	ApplyFunc		func(ctx *gin.Context) []Options
}

// Apply apply filter in gin context
func (a *APIFilter) Apply(ctx *gin.Context) []Options {
	return a.ApplyFunc(ctx)
}

// LimitFilter filter for limit data for API
func LimitFilter() Filter {
	return &APIFilter{ApplyFunc: func(ctx *gin.Context) []Options {
		var options []Options
		if limitStr := ctx.Query("limit"); limitStr != "" {
			limit, err := strconv.ParseInt(limitStr, 10, 64)
			if err != nil {
				common.SendError(ctx, err, 400)
				return options
			}
			if limit > 20 {
				limit = 20
			}
			options = append(options, WithLimit(int(limit)))
		}
		return options
	}}
}

// OrderingFilter filter for order data in API
func OrderingFilter() Filter{
	return &APIFilter{ApplyFunc: func(ctx *gin.Context) []Options {
		orderStr := ctx.Query("order")
		var orderValue = "id"
		var orderType = "asc"
		var customType = false
		if len(orderStr) == 0{
			return []Options{WithOrder(fmt.Sprintf("%s %s", orderValue, orderType))}
		}
		if orderStr[0] == '+'{
			orderType = "asc"
			customType = true
		}
		if orderStr[0] == '-'{
			orderType = "desc"
			customType = true
		}
		orderValue = orderStr
		if customType {
			orderValue = orderStr[1:]
		}
		return []Options{WithOrder(fmt.Sprintf("%s %s", orderValue, orderType))}
	}}
}

// LimitAndPageFilter page and limit in API
func LimitAndPageFilter() Filter {
	return &APIFilter{ApplyFunc: func(ctx *gin.Context) []Options {
		var options []Options
		var err error
		var limit int64
		if limitStr := ctx.Query("limit"); limitStr != "" {
			limit, err = strconv.ParseInt(limitStr, 10, 64)
			if err != nil {
				return options
			}
			if limit > 20 {
				limit = 20
			}
			options = append(options, WithLimit(int(limit)))
		}
		if pageStr := ctx.Query("page"); pageStr != "" && limit != 0 {
			page, err := strconv.ParseInt(pageStr, 10, 64)
			if err != nil {
				return options
			}
			var offset = (limit * page) - limit
			if offset >= 0 {
				options = append(options, WithOffset(int(offset)))
			}
		}
		return options
	}}
}