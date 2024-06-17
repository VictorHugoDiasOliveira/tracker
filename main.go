package main

import (
	"fmt"
	"io"
	"net"
	"net/http"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type user struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func main() {
	router := gin.Default()

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{"POST", "GET", "PUT", "DELETE"},
		AllowHeaders:     []string{"Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Lenght"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	router.GET("/", GetIpAddress)
	router.POST("/login", GetEmailAndPassword)

	router.Run("localhost:8080")
}

func GetIpAddress(c *gin.Context) {
	ip := "204.152.191.37"
	// ip := c.Request.Header.Get("X-Forwarded-For")
	if ip == "" {
		ip, _, _ = net.SplitHostPort(c.Request.RemoteAddr)
	}
	fmt.Println("IP: ", ip)

	url := "http://ip-api.com/json/" + ip

	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Erro ao fazer requisicao:", err)
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Erro ao lancar corpo da resposta:", err)
		return
	}

	fmt.Println("Infos: ", string(body))
}

func GetEmailAndPassword(c *gin.Context) {
	var newUser user

	if err := c.BindJSON(&newUser); err != nil {
		return
	}

	fmt.Println("Email: ", newUser.Email)
	fmt.Println("Password: ", newUser.Password)
}
