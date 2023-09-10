package main

import (
	"context"
	"encoding/json"
	"github.com/a-h/templ"
	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
	"gogame/html"
	"gogame/htmx"
	"log"
	"net/http"
	"time"
    "errors"
)

var rooms = make(map[string]*Room)

const MAX_GUESTS = 5

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	// TODO:actually check origin
	CheckOrigin: func(r *http.Request) bool { return true },
}

func ConnWriteTemplate(conn *websocket.Conn, t templ.Component) error {
	if conn == nil {
		log.Print("connection is nil, trying to send component")
		t.Render(context.TODO(), log.Writer())
		log.Writer().Write([]byte("\n"))
		return nil
	}
	w, err := conn.NextWriter(websocket.TextMessage)
	if err != nil {
		return err
	}
	t.Render(context.TODO(), w)
	err = w.Close()
	if err != nil {
		return err
	}
	return nil
}

type Player struct {
	Conn        *websocket.Conn
	DisplayName string
}

type Room struct {
	name   string
	host   Player
	guests []Player
}

func (r *Room) toHTMLRoom() html.Room {
	players := make([]string, len(r.guests))
	for i, guest := range r.guests {
		players[i] = guest.DisplayName
	}
	return html.Room{Name: r.name, Host: r.host.DisplayName, Guests: players}
}

func (r *Room) Start() {
	go r.ListenForHostMessages()
	for _, g := range r.guests {
		go r.ListenForGuestMessages(g)
	}
}

func (r *Room) ListenForHostMessages() {
	for {
		_, p, err := r.host.Conn.ReadMessage()
		if err != nil {
			r.DisconnectHost(err)
			return
		}
		log.Println(string(p))
	}
}

func (r *Room) ListenForGuestMessages(g Player) {
	for {
		_, p, err := g.Conn.ReadMessage()
		if err != nil {
			r.DisconnectGuest(g, err)
			return
		}
		log.Println(string(p))
	}
}

func (r *Room) DisconnectHost(err error) {
	r.host.Conn.Close()
	if err != nil {
		log.Println(err, "disconnecting host", r.host.DisplayName)
	} else {
		log.Println("disconnecting host", r.host.DisplayName)
	}
	// TODO: redirect players to lobby
}

func (r *Room) DisconnectGuest(g Player, err error) {
	g.Conn.Close()
	if err != nil {
		log.Println(err, "disconnecting", g.DisplayName)
	} else {
		log.Println("disconnecting", g.DisplayName)
	}
	// TODO: remove guest from guests
}

func (r *Room) ConnectGuest(conn *websocket.Conn) error {
	type GuestJoinMsg struct {
		Headers     htmx.Headers `json:"HEADERS"`
		DisplayName string       `json:"display-name"`
	}
	var err error = nil
	defer func() {
		// TODO: remove new guest from list if already added
		if err != nil {
			conn.Close()
		}
	}()
	var p []byte
	for {
		_, p, err = conn.ReadMessage()
		if err != nil {
			return err
		}
		break
	}
	var msg GuestJoinMsg
	if err = json.Unmarshal(p, &msg); err != nil {
		return err
	}
	log.Println(msg)
	guest := Player{Conn: conn, DisplayName: msg.DisplayName}
	r.guests = append(r.guests, guest)
	conn.WriteMessage(websocket.TextMessage, []byte(`<div id="join-room-modal"></div>`))

	newGuestNameList := []string{guest.DisplayName}
	newGuestList := html.GuestList(newGuestNameList, true)
	err = ConnWriteTemplate(r.host.Conn, newGuestList)
	if err != nil {
		return err
	}
	for _, g := range r.guests {
		err = ConnWriteTemplate(g.Conn, newGuestList)
		if err != nil {
			return err
		}
	}
	go r.ListenForGuestMessages(guest)
	return nil
}

func createRoomHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		html.CreateRoomPage().Render(r.Context(), w)
		break
	case http.MethodPost:
		roomName := r.FormValue("room-name")
		hostName := r.FormValue("display-name")
		// TODO: form validation
		if rooms[roomName] != nil {
			http.Error(w, "Room Already Created!", http.StatusForbidden)
		}
		room := Room{name: roomName, host: Player{DisplayName: hostName}}
		rooms[roomName] = &room
		w.Header().Set("HX-Replace-Url", "/rooms/"+roomName)
		html.HostRoomPage(room.toHTMLRoom()).Render(r.Context(), w)
		// TODO: MethodDELETE
	}
}

func joinRoomHandler(w http.ResponseWriter, r *http.Request) {
	roomid := mux.Vars(r)["roomid"]
	room := rooms[roomid]
    if room == nil {
        log.Println("access to room",roomid, "which does not exist")
        http.Error(w, "room does not exist", http.StatusBadRequest)
        return;
    }
	html.JoinRoomPage(room.toHTMLRoom()).Render(r.Context(), w)
}

func guestWSHandler(room *Room, conn *websocket.Conn, w http.ResponseWriter, r *http.Request) {
	if len(room.guests) == MAX_GUESTS {
		// TODO: redirect to lobby and show error
		log.Println("room is full")
        http.Error(w, "room is full", http.StatusForbidden)
	}
    log.Println("websocket connection established to new guest in room:", room.name)
    go room.ConnectGuest(conn)
}

func hostWSHandler(room *Room, conn *websocket.Conn, w http.ResponseWriter, r *http.Request) {
    if room.host.Conn != nil {
        http.Error(w, "host already connected", http.StatusForbidden)
    }
    room.host.Conn = conn
    go room.Start()
}

func wsHandler(w http.ResponseWriter, r *http.Request) {
	room_name := mux.Vars(r)["roomid"]

    // TODO: redirect to home in case of error
	if room_name == "" {
		log.Println("no room name specified for websocket connection")
        http.Error(w, "no room name specified for websocket connection", http.StatusBadRequest)
		return
	}
    room := rooms[room_name]
    if room == nil {
        log.Println("room does not exist")
        http.Error(w, "room does not exist", http.StatusBadRequest)
        return
    }
    // no errors: ok to upgrade connection
	conn, err := upgrader.Upgrade(w, r, nil)
	defer func() {
		if err == nil {
			return
		}
        log.Println(err)
		if conn != nil {
			conn.Close()
		}
	}()
    if err != nil {
        return
    }
	kind := mux.Vars(r)["kind"]
    switch kind {
        case "host":
            hostWSHandler(room, conn, w, r)
            break;
        case "guest":
            guestWSHandler(room, conn, w, r)
            break;
        default:
            err = errors.New("invalid websocket kind")
            http.Error(w, "invalid websocket kind", http.StatusBadRequest)
    }
}

func main() {
	// TODO: store this in db
	rt := mux.NewRouter()

	rt.Handle("/", templ.Handler(html.LandingPage()))

	rt.HandleFunc("/rooms", func(w http.ResponseWriter, r *http.Request) {
		var roomInfos = make([]html.Room, len(rooms))
		i := 0
		for _, room := range rooms {
			if room.name == "" {
				continue
			}
			// TODO: add player names once they are collected
			roomInfos[i] = room.toHTMLRoom()
			i++
		}
		// in case of hx-boost don't render (and send!) the whole page
		if r.Header.Get("HX-Request") == "true" {
			html.RoomList(roomInfos).Render(r.Context(), w)
		} else {
			html.RoomListPage(roomInfos).Render(r.Context(), w)
		}
	})
	rmrt := rt.PathPrefix("/rooms/{roomid}").Subrouter()
	rmrt.HandleFunc("/{kind}/ws", wsHandler)

	rt.HandleFunc("/rooms/create", createRoomHandler)
	rmrt.HandleFunc("/join", joinRoomHandler)

	rt.HandleFunc("/ok", func(w http.ResponseWriter, r *http.Request) {
		// NOTE: cannot be http.StatusNoContent because this is used with hx-swap="delete"
		// and htmx will not delete the element with a StatusNoContent response
		w.WriteHeader(http.StatusOK)
	})

	http.Handle("/", rt)
	s := &http.Server{
		Handler:      rt,
		Addr:         "127.0.0.1:8080",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(s.ListenAndServe())
}
