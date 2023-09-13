package main

import (
	"encoding/json"
	"github.com/a-h/templ"
	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
	"gogame/htmx"
	"gogame/player"
	"log"
	"net/http"
	"time"
)

const MAX_GUESTS = 5

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	// TODO:actually check origin
	CheckOrigin: func(r *http.Request) bool { return true },
}

func createRoomHandler(rr *RoomRegistry, w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		CreateRoomPage().Render(r.Context(), w)
	case http.MethodPost:
		hostName := r.FormValue("display-name")
		roomId := r.FormValue("room-name")
		if rr.RoomExists(roomId) {
			log.Println("error: attempt to create room that already exists:", roomId)
			http.Error(w, "Room already exists", http.StatusBadRequest)
			return
		}
		room := NewRoom{Name: roomId, HostName: hostName}
		rr.Register <- room
		w.Header().Set("HX-Replace-Url", "/rooms/"+roomId)
		HostRoomPage(room.toHTMLRoom()).Render(r.Context(), w)
		log.Println("created room:", roomId)
		// TODO: MethodDELETE
	}
}

func joinRoomHandler(rr *RoomRegistry, w http.ResponseWriter, r *http.Request) {
    // TODO: use switch on method like createRoomHandler instead of joining websocket and waiting for name
	roomId := mux.Vars(r)["roomid"]
	if !rr.RoomExists(roomId) {
		errmsg := "error: tried to join non-existant room: " + roomId
		log.Println(errmsg)
		http.Error(w, errmsg, http.StatusBadRequest)
	}
	JoinRoomPage(rr.Room(roomId).toHTMLRoom()).Render(r.Context(), w)
}

func wsHandler(rr *RoomRegistry, w http.ResponseWriter, r *http.Request) {
	roomId := mux.Vars(r)["roomid"]
	if !rr.RoomExists(roomId) {
		log.Println("error: attempt to connect to non-existant room:", roomId)
		http.Error(w, "error: attempt to connect to non-existant room:", http.StatusBadRequest)
	}
	kind := mux.Vars(r)["kind"]
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("failed to upgrade connection:", err)
		http.Error(w, "failed to upgrade connection", http.StatusInternalServerError)
	}
	switch kind {
	case "host":
		p := player.NewPlayer(conn, "")
		rr.Join <- JoinRequest{RoomId: roomId, IsHost: true, Player: &p}
	case "guest":
		// NOTE: must be in go routine because this handler must respond
		// for the websocket connection to be established
		go waitForGuestName(rr, roomId, conn)
	}
}

func waitForGuestName(rr *RoomRegistry, roomId string, conn *websocket.Conn) {
	// TODO: redirect to home in case of error
	type SetNameMsg struct {
		Headers htmx.Headers `json:"HEADERS"`
        Group string `json:"group"`
        Type string `json:"type"`
		Data    struct {
			DisplayName string `json:"display-name"`
		} `json:"data"`
	}
	// block until name message is recieved
	_, msgBytes, err := conn.ReadMessage()
	if err != nil {
		log.Println("error: failed to read name message while joining room:", roomId)
		// TODO: conn.WriteCloseMessage
	}
	var msg SetNameMsg
	// TODO: error if name is ""
	if err = json.Unmarshal(msgBytes, &msg); err != nil {
		errmsg := "error: failed to decode set-name message while joining room: " + roomId
		log.Printf("%s: %v\n", errmsg, err)
		// TODO: conn.WriteCloseMessage
	}
    if msg.Type != "display-name" {
        log.Println("error: recieved message of type", msg.Type,"but expected display-name")
        return
    }
	p := player.NewPlayer(conn, msg.Data.DisplayName)
	rr.Join <- JoinRequest{RoomId: roomId, IsHost: false, Player: &p}
}

func roomBrowserHandler(rr *RoomRegistry, w http.ResponseWriter, r *http.Request) {
	roomInfos := rr.GetHTMLRooms()
	// in case of hx-boost don't render (and send!) the whole page
	if r.Header.Get("HX-Request") == "true" {
		RoomList(roomInfos).Render(r.Context(), w)
	} else {
		RoomListPage(roomInfos).Render(r.Context(), w)
	}
}

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.Method, r.RequestURI)
		next.ServeHTTP(w, r)
	})
}

func main() {
	rr := newRoomRegistry()
	go rr.run()
	rrwrap := func(handler func(reg *RoomRegistry, w http.ResponseWriter, r *http.Request)) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			handler(&rr, w, r)
		}
	}
	rt := mux.NewRouter()

	rt.Handle("/", templ.Handler(LandingPage()))
	rt.HandleFunc("/rooms", rrwrap(roomBrowserHandler))
	rt.HandleFunc("/rooms/create", rrwrap(createRoomHandler))

	rmrt := rt.PathPrefix("/rooms/{roomid}").Subrouter()
	rmrt.HandleFunc("/{kind}/ws", rrwrap(wsHandler))
	rmrt.HandleFunc("/join", rrwrap(joinRoomHandler))

	// helper function to do hmtx client side actions (like hx-push-url, or hx-swap="delete")
	// especially when using a websocket connection that cannot send responses in the same way
	rt.HandleFunc("/ok", func(w http.ResponseWriter, r *http.Request) {
		// NOTE: cannot be http.StatusNoContent because this is used with hx-swap="delete"
		// and htmx will not delete the element with a StatusNoContent response
		w.WriteHeader(http.StatusOK)
	})

	http.Handle("/", rt)
	rt.Use(loggingMiddleware)
	s := &http.Server{
		Handler:      rt,
		Addr:         "127.0.0.1:8080",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	log.Fatal(s.ListenAndServe())
}
