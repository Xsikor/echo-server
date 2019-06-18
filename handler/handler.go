package handler

import (
	"fmt"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"sync"
)

const (
	authHeader = "secretKey"
	secret     = "54686973206973206d7920626f6f6d737469636b"
)

type Handler struct {
	up     websocket.Upgrader
	secret string

	clients  []*websocket.Conn
	clientMu sync.RWMutex
}

func NewHandler() *Handler {
	return &Handler{
		up:       websocket.Upgrader{},
		clientMu: sync.RWMutex{},
		clients:  make([]*websocket.Conn, 0),
		secret:   secret,
	}
}

func (h *Handler) Connect(w http.ResponseWriter, r *http.Request) {
	secretKey := r.Header.Get(authHeader)
	if secretKey == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if secretKey != secret {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	log.Println("new client connected", r.Header)

	c, err := h.up.Upgrade(w, r, nil)
	if err != nil {
		log.Print("upgrade error:", err)
		return
	}

	h.clientMu.Lock()
	clientIndex := len(h.clients)
	h.clients = append(h.clients, c)
	h.clientMu.Unlock()

	go h.clientWorker(clientIndex, c)
}

func (h *Handler) clientWorker(i int, c *websocket.Conn) {
	defer func() {
		c.Close()
		h.clientMu.Lock()
		h.clients = append(h.clients[:i], h.clients[i+1:]...)
		h.clientMu.Unlock()

		fmt.Println("Removed client", i, "left", len(h.clients))
	}()

	for {
		mt, message, err := c.ReadMessage()
		if err != nil {
			log.Println("client", i, "read error:", err)
			break
		}

		log.Printf("client %d recv: %s", i, message)
		err = c.WriteMessage(mt, message)
		if err != nil {
			log.Println("client", i, "write error:", err)
			break
		}
	}
}
