package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

var balance = 1000

func main() {
	router := gin.Default()
	router.GET("/balance/", getBalance)
	router.GET("/deposit/:input", deposit)
	
	// router.GET("/json3", returnJson3)
	// router.GET("/para1", para1)
	// router.GET("/para2/:input", para2)
	// router.POST("/post", post)
	router.Run(":80")
}

// func returnJson3(c *gin.Context) {
// 	type Result struct {
// 		Status  string `json:"status"`
// 		Message string `json:"message"`
// 	}

// 	var result = Result{
// 		Status:  "OK",
// 		Message: "This is Json",
// 	}

// 	c.JSON(http.StatusOK, result)
// }

// func para1(c *gin.Context){
// 	input := c.Query("input")
// 	msg := []byte("您輸入的文字為: \n" + input)
// 	c.Data(http.StatusOK, "text/plain; charset=utf-8;", msg)
// }

// func para2(c *gin.Context) {
// 	msg := c.Param("input")
// 	c.String(http.StatusOK, "您輸入的文字為: \n%s", msg)
// }

// func post(c *gin.Context){
// 	msg := c.DefaultPostForm("input", "表單沒有input。")
// 	c.String(http.StatusOK, "您輸入的文字為: \n%s", msg)
// }

func getBalance(context *gin.Context) {
	var msg = "您的帳戶內有:" + strconv.Itoa(balance) + "元"
	context.JSON(http.StatusOK, gin.H{
		"amount":  balance,
		"status":  "ok",
		"message": msg,
	})
}

func deposit(context *gin.Context) {
	var status string
	var msg string

	input := context.Param("input")
	amount, err := strconv.Atoi(input)

	if err == nil {
		if amount <= 0 {
			amount = 0
			status = "failed"
			msg = "操作失敗，存款金額需大於0元！"
		} else {
			balance += amount
			status = "ok"
			msg = "已成功存款" + strconv.Itoa(amount) + "元"
		}
	} else {
		amount = 0
		status = "failed"
		msg = "操作失敗，輸入有誤！"
	}
	context.JSON(http.StatusOK, gin.H{
		"amount":  amount,
		"status":  status,
		"message": msg,
	})
}

func wrapResponse(context *gin.Context, amount int, err error) {
	var r = struct {
		Amount  int    `json:"amount"`
		Status  string `json:"status"`
		Message string `json:"message"`
	}{
		Amount:  amount,
		Status:  "ok", // 預設狀態為ok
		Message: "",
	}

	if err != nil {
		r.Amount = 0
		r.Status = "failed"     // 若出現任何err，狀態改為failed
		r.Message = err.Error() // Message回傳錯誤訊息
	}

	context.JSON(http.StatusOK, r)
}
