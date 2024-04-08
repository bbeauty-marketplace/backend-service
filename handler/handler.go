package handler

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

func HelloWorldHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Hello world",
	})
}

var (
	googleOauthConfig *oauth2.Config
)

func init() {
	googleOauthConfig = &oauth2.Config{
		RedirectURL:  "http://localhost:8080/auth/google/callback",
		ClientID:     "YOUR_CLIENT_ID",
		ClientSecret: "YOUR_CLIENT_SECRET",
		Scopes: []string{
			"https://www.googleapis.com/auth/userinfo.email",
			"https://www.googleapis.com/auth/userinfo.profile",
		},
		Endpoint: google.Endpoint,
	}
}

func GoogleLoginHandler(c *gin.Context) {
	url := googleOauthConfig.AuthCodeURL("state")
	c.Redirect(http.StatusTemporaryRedirect, url)
}

func GoogleCallbackHandler(c *gin.Context) {
	code := c.Query("code")
	token, err := googleOauthConfig.Exchange(context.Background(), code)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to exchange token",
		})
		return
	}
	if !token.Valid() {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Invalid token",
		})
		return
	}

	// Use the token to make API requests
	// ...

	c.JSON(http.StatusOK, gin.H{
		"message": "Logged in via Google",
	})
}
