package middleware

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func GetUserFromContext(ctx *gin.Context) {
	sub := ctx.GetString("sub")
	if sub == "" {
		fmt.Println("Cannot retrieve user id, user not authenticated/authorized")
	} else {
		fmt.Println(sub)
	}
}
