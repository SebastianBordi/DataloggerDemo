package socket

import (
	"fmt"
	"log"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
)

type webSocketStructure struct {
	mac         string
	channel     *chan string
	broadcast   *websocket.Conn
	subscribers map[string]*websocket.Conn
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  2048,
	WriteBufferSize: 2048,
}

var socketsChannels = make(map[string]*webSocketStructure)

var BasicResponse func(w *http.ResponseWriter, statusCode int, messge string)

func SocketEndpoint(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	mac := params["mac"]
	isSensor := params["isBroadcaster"] == "broadcast"
	socket, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		BasicResponse(&w, 500, "internal server error")
		return
	}

	sc := socketsChannels[mac]
	if sc == nil && !isSensor {
		BasicResponse(&w, 404, "don't exist sensor broadcast")
		return
	}
	if isSensor {
		if sc == nil {
			sc = &webSocketStructure{}
		}
		sc.broadcast = socket
		sc.subscribers = make(map[string]*websocket.Conn)
		socketsChannels[mac] = sc
		go broadcast(sc)
	} else {
		uuidString, err := uuid.NewUUID()
		if err != nil {
			log.Println(err)
			BasicResponse(&w, 500, "internal server error")
			return
		}
		sc.subscribers[uuidString.String()] = socket
		socket.WriteMessage(websocket.TextMessage, []byte(fmt.Sprintf("%s subscribed to %s sensor", uuidString, mac)))
	}
}

func broadcast(wss *webSocketStructure) {
	for {
		messageType, p, err := wss.broadcast.ReadMessage()
		if err != nil {
			log.Print(err)
			log.Print("Broadcast Closed")
			return
		}
		for _, item := range wss.subscribers {
			item.WriteMessage(messageType, p)
		}
	}
}
