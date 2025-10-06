package middleware

import (
	"accountantapp/go-service/internal/utils"
	"net/http"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
)

type client struct {
	lastSeen time.Time
	requests int
}

var (
	clients = make(map[string]*client)
	mutex   sync.Mutex
)

func RateLimit(maxRequests int, window time.Duration) gin.HandlerFunc {
	go cleanupClients()

	return func(c *gin.Context) {
		ip := c.ClientIP()

		mutex.Lock()
		defer mutex.Unlock()

		if _, exists := clients[ip]; !exists {
			clients[ip] = &client{}
		}

		cl := clients[ip]

		if time.Since(cl.lastSeen) > window {
			cl.requests = 0
			cl.lastSeen = time.Now()
		}

		if cl.requests >= maxRequests {
			utils.Error(c, http.StatusTooManyRequests, "Rate limit exceeded", nil)
			c.Abort()
			return
		}

		cl.requests++
		cl.lastSeen = time.Now()
		c.Next()
	}
}

func cleanupClients() {
	for {
		time.Sleep(time.Minute)
		mutex.Lock()
		for ip, client := range clients {
			if time.Since(client.lastSeen) > 10*time.Minute {
				delete(clients, ip)
			}
		}
		mutex.Unlock()
	}
}
