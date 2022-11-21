package middlewares

import (
	"jwt-golang/auth"

	"github.com/gin-gonic/gin"
)

func Auth() gin.HandlerFunc {
	return func(context *gin.Context) {
		// Extracts the Authorization header from the HTTP context. Ideally, we expect the token to be sent as a header
		// by the client. If there are no tokens found at the header, the application throws a 401 error with the
		// appropriate error message.
		tokenString := context.GetHeader("Authorization")
		if tokenString == "" {
			context.JSON(401, gin.H{"error": "request does not contain an access token"})
			context.Abort()
			return
		}

		// we validate the token using the earlier created helper function. If the token is found to be invalid or
		// expired, the application would throw a 401 Unauthorized exception. If the token is valid, the middleware
		// allows the flow and the request reaches the required controller’s endpoint. As simple as that
		err := auth.ValidateToken(tokenString)
		if err != nil {
			context.JSON(401, gin.H{"error": err.Error()})
			context.Abort()
			return
		}
		context.Next()
	}
}
