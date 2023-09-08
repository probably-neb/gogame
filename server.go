package main

import (
	"fmt"
	"github.com/a-h/templ"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
    "gogame/html"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     func(r *http.Request) bool { return true },
}

type Room struct {
    name string
    host *websocket.Conn
    guests []*websocket.Conn
}

// TODO: store this in db
var rooms = make(map[string]Room)

func main() {
	http.Handle("/", templ.Handler(html.LandingPage()))
    http.HandleFunc("/rooms", func(w http.ResponseWriter, r *http.Request) {
        var room = r.URL.Query().Get("q")
        log.Println("access to room: [", room, "]")
        html.RoomPage(room).Render(r.Context(), w)
    })
    http.HandleFunc("/rooms/create", func(w http.ResponseWriter, r *http.Request) {
        if r.Method == http.MethodGet {
            html.CreateRoomModal().Render(r.Context(), w)
        } else if r.Method == http.MethodPost {
            name := r.FormValue("name")
            // TODO: error handling (form validation!)
            log.Println("create room", name)
            location := "/rooms?q="+name
            w.Header().Set("HX-Redirect", location)
            w.WriteHeader(http.StatusSeeOther)
            rooms[name] = Room{name: name}
        }
    })
    http.HandleFunc("/rooms/create/cancel", func(w http.ResponseWriter, r *http.Request) {
        // NOTE: cannot be http.StatusNoContent because htmx will do nothing when it should
        // remove the modal according to the hx- attributes
        w.WriteHeader(http.StatusOK)
    })
    http.HandleFunc("/rooms/join", func(w http.ResponseWriter, r *http.Request) {
        var room_infos = make([]html.Room, len(rooms));
        i := 0
        for _, room := range rooms {
            if room.name == "" {
                continue;
            }
            room_infos[i] = html.Room{Name: room.name, Players: make([]string, 0)}
            i++
        }
        html.JoinRoomPage(room_infos).Render(r.Context(), w)
    })

	fmt.Println("Server started at 8080 port")
	//Use the default DefaultServeMux.
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
