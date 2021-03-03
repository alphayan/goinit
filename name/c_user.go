package main

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

func init() {
	ops = append(ops, addUser)
}
func addUser(e *gin.Engine) {
	apiV1 := e.Group("/api/v1")
	apiV1.GET("/users/:id", getUser)
	apiV1.GET("/users", getsUser)
	apiV1.PUT("/users/:id", putUser)
	apiV1.POST("/users", postUser)
	apiV1.DELETE("/users/:id", delUser)
}
func getUser(c *gin.Context) {
	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
	u := &User{ID: id}
	if err := u.get(); err != nil {
		c.JSON(200, NewRespFault("获取user失败："+err.Error()))
		return
	}
	c.JSON(200, NewRespWithData(200, "获取user成功", u))
}
func getsUser(c *gin.Context) {
	u := new(User)
	limit, _ := strconv.Atoi(c.Query("limit"))
	page, _ := strconv.Atoi(c.Query("page"))
	us, count := u.gets(limit, page)
	c.JSON(200, NewRespWithCount(200, "获取user列表成功", us, count))
}
func postUser(c *gin.Context) {
	u := &User{}
	if err := c.Bind(u); err != nil {
		c.JSON(200, NewRespBad("参数失败:"+err.Error()))
		return
	}
	if err := u.update(); err != nil {
		c.JSON(200, NewRespFault("创建user失败:"+err.Error()))
		return
	}
	c.JSON(200, NewRespOK("创建user成功"))
}
func putUser(c *gin.Context) {
	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
	u := &User{ID: id}
	if err := c.Bind(u); err != nil {
		c.JSON(200, NewRespBad("参数失败:"+err.Error()))
		return
	}
	if err := u.update(); err != nil {
		c.JSON(200, NewRespFault("更新user失败:"+err.Error()))
		return
	}
	c.JSON(200, NewRespOK("更新user成功"))
}
func delUser(c *gin.Context) {
	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
	u := &User{ID: id}
	if err := u.delete(); err != nil {
		c.JSON(200, NewRespFault("删除user失败:"+err.Error()))
		return
	}
	c.JSON(200, NewRespOK("删除user成功"))
}
