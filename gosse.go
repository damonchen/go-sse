package sse

import (
	"fmt"
	"net/http"
)

type SSEServer struct {
	// type string
	messageChan chan []byte
	clients map[chan []byte]struct{}
}

func NewSSEServer() *SSEServer {
	return &SSEServer{
		messageChan: make(chan []byte),
	}
}

func (sse *SSEServer) Publish(message []byte) {
	sse.messageChan <- message
}

func (sse *SSEServer) Handler() http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {

		// Make sure that the writer supports flushing.
		//
		flusher, ok := rw.(http.Flusher)

		if !ok {
			http.Error(rw, "Streaming unsupported!", http.StatusInternalServerError)
			return
		}

		rw.Header().Set("Content-Type", "text/event-stream")
		rw.Header().Set("Cache-Control", "no-cache")
		rw.Header().Set("Connection", "keep-alive")
		rw.Header().Set("Access-Control-Allow-Origin", "*")

		// Listen to connection close and un-register messageChan
		notify := rw.(http.d).CloseNotify()

		go func() {
			<-notify
		}()

		for {

			select {
			case msg :<- sse.messageChan:
						// Write to the ResponseWriter
			// Server Sent Events compatible
			fmt.Fprintf(rw, "data: %s\n\n", msg)

			// Flush the data immediatly instead of buffering it for later.
			flusher.Flush()
		case <-closingClients:
			delete(broker.clients, s)
			}

	
		}

	}
}
