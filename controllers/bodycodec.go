package controllers

import (
	"net/http"
	"net/url"
	"strconv"
	"test/testgin/models"

	"github.com/gin-gonic/gin"
)

/*
POST /testgin/urlencode?id=1234&page=1 HTTP/1.1
Content-Type: application/x-www-form-urlencoded

username=alice&password=1234
*/
func URLEncode(ctx *gin.Context) {
	//原生方式，querystring
	// req := ctx.Request
	// rw := ctx.Writer
	// req.ParseForm()
	// id := req.Form.Get("id")
	// page := req.Form.Get("page")
	// username := req.PostForm.Get("username")
	// password, _ := strconv.Atoi(req.PostForm.Get("password"))
	// fmt.Println(id, page, username, password)
	// data := make(url.Values) //url编码
	// data.Set("uid", "1001")
	// data.Add("result", "success")
	// rw.Write([]byte(data.Encode()))

	//gin方式
	// id := ctx.Query("id")
	// page := ctx.Query("page")
	// username := ctx.PostForm("username")
	// password := ctx.PostForm("password")
	// fmt.Println(id, page, username, password)
	var u models.User
	ctx.ShouldBind(&u)
	data := make(url.Values) //url编码
	data.Set("id", u.Id)
	data.Set("page", u.Page)
	data.Set("username", u.Username)
	data.Set("password", strconv.Itoa(u.Password))
	ctx.String(http.StatusOK, data.Encode())
}

func Json(ctx *gin.Context) {
	//原生方式
	// req := ctx.Request
	// rw := ctx.Writer
	// var u user
	// // json.Unmarshal(req.Body, &u)
	// json.NewDecoder(req.Body).Decode(&u)
	// fmt.Println(u.Username, u.Passwd)
	// r := &resultData{"1001", "success"} //映射数据库结构体
	// data, _ := json.Marshal(r)          //序列化
	// rw.Write(data)

	//gin方式
	var u models.User
	ctx.BindJSON(&u)
	// ctx.JSON(http.StatusOK, gin.H{
	// 	"uid":    "1001",
	// 	"result": "success",
	// 	"s":      []string{"lena", "austin", "foo"},
	// })
	ctx.JSON(http.StatusOK, &u)
	// ctx.SecureJSON(http.StatusOK, names) //是什么？
}

//可以理解为返回一个js文件
func Jsonp(ctx *gin.Context) {
	ctx.JSONP(http.StatusOK, gin.H{
		"price":   "200",
		"tickets": 9,
	})
}
