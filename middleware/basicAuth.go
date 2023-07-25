package middleware

import "github.com/gin-gonic/gin"

func AuthDetails() gin.Accounts {
	account := gin.Accounts{
		"root": "root",
	}
	return account
}
