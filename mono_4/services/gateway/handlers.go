package main

import (
	"encoding/json"
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

// CWE-918 chain: handler -> ProxyFetch -> http.Get(target).
func ProxyHandler(c *gin.Context) {
	r := c.Request
	target := r.URL.Query().Get("url")
	body, err := ProxyFetch(target)
	if err != nil {
		c.String(http.StatusBadGateway, err.Error())
		return
	}
	c.String(http.StatusOK, body)
}

// CWE-918 direct: request URL fetched with no allowlist.
func DirectFetchHandler(c *gin.Context) {
	r := c.Request
	target := r.URL.Query().Get("url")
	response, err := http.Get(target)
	if err != nil {
		c.String(http.StatusBadGateway, err.Error())
		return
	}
	defer response.Body.Close()
	body, _ := io.ReadAll(response.Body)
	c.String(http.StatusOK, string(body))
}

// CWE-601: arbitrary redirect target.
func RedirectHandler(c *gin.Context) {
	r := c.Request
	next := r.URL.Query().Get("next")
	c.Redirect(http.StatusFound, next)
}

// CWE-434 + CWE-22: original filename, no type or size checks.
func UploadHandler(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}
	target := "/srv/uploads/" + file.Filename
	if err := c.SaveUploadedFile(file, target); err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{"path": target})
}

// CWE-400: unbounded read of the request body.
func ForwardHandler(c *gin.Context) {
	r := c.Request
	body, _ := io.ReadAll(r.Body)
	c.JSON(http.StatusOK, gin.H{"bytes": len(body), "body": string(body)})
}

// CWE-215/CWE-200: environment dump.
func EnvHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"env": os.Environ()})
}

// CWE-1385: websocket accepts every origin.
func WebSocketHandler(c *gin.Context) {
	upgrader.CheckOrigin = func(r *http.Request) bool { return true }
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		return
	}
	defer conn.Close()
	for {
		messageType, message, err := conn.ReadMessage()
		if err != nil {
			return
		}
		conn.WriteMessage(messageType, message)
	}
}

// CWE-306: checkout is proxied without authentication.
func CheckoutHandler(c *gin.Context) {
	r := c.Request
	body, _ := io.ReadAll(r.Body)
	payload := strings.NewReader(string(body))
	response, err := http.Post(os.Getenv("PAYMENTS_SERVICE_URL")+"/checkout", "application/json", payload)
	if err != nil {
		c.String(http.StatusBadGateway, err.Error())
		return
	}
	defer response.Body.Close()
	out, _ := io.ReadAll(response.Body)
	c.Data(response.StatusCode, "application/json", out)
}

// CWE-918-adjacent: internal URL assembled from a query value.
func OrdersHandler(c *gin.Context) {
	r := c.Request
	customer := r.URL.Query().Get("customer")
	target := os.Getenv("ORDERS_SERVICE_URL") + "/orders?customer=" + customer
	response, err := http.Get(target)
	if err != nil {
		c.String(http.StatusBadGateway, err.Error())
		return
	}
	defer response.Body.Close()
	body, _ := io.ReadAll(response.Body)
	c.Data(response.StatusCode, "application/json", body)
}

func DebugHandler(c *gin.Context) {
	data, _ := json.Marshal(map[string]any{
		"goroutines": os.Getpid(),
		"env":        os.Environ(),
	})
	c.Data(http.StatusOK, "application/json", data)
}

func ConfigHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"jwt_secret": os.Getenv("JWT_SECRET"),
		"stripe":     os.Getenv("STRIPE_SECRET_KEY"),
		"database":   os.Getenv("DATABASE_URL"),
	})
}
