package main

import (
    "fmt"
    "github.com/a-h/templ"
    "github.com/gorilla/websocket"
    "log"
    "net/http"
    "gogame/html"
)

const MAX_GUESTS = 5 

var upgrader = websocket.Upgrader{
    ReadBufferSize:  1024,
    WriteBufferSize: 1024,
    CheckOrigin:     func(r *http.Request) bool { return true },
}

type Player struct {
    Conn *websocket.Conn
    DisplayName string
}

type Room struct {
    name string
    host Player
    num_guests int
    guests [MAX_GUESTS]Player
}

func (r *Room) toHTMLRoom() html.Room {
    players := make([]string, len(r.guests)+1)
    players[0] = r.host.DisplayName
    for i, guest := range r.guests {
        players[i+1] = guest.DisplayName
    }
    return html.Room{Name: r.name, Players: players}
}

// TODO: store this in db
var rooms = make(map[string]*Room)

func main() {
    http.Handle("/", templ.Handler(html.LandingPage()))
    http.HandleFunc("/rooms", func(w http.ResponseWriter, r *http.Request) {
        room_name := r.URL.Query().Get("q")
        room := rooms[room_name]
        log.Println("access to room: [", room.name, "]")
        html.RoomPage(room.toHTMLRoom()).Render(r.Context(), w)
    })
    http.HandleFunc("/rooms/ws", func(w http.ResponseWriter, r *http.Request) {
        conn, err := upgrader.Upgrade(w, r, nil)
        defer func() {
            if err == nil {
                return
            }
            if conn != nil {
                conn.Close()
            }
        }();
        if err != nil {
            log.Println(err)
            return
        }
        log.Println("websocket connection established")
        room_name := r.URL.Query().Get("q")
        is_host := r.URL.Query().Get("host") == "true"
        if room_name == "" {
            log.Println("no room name specified for websocket connection")
            return
        }
        if rooms[room_name] == nil {
            log.Println("room does not exist")
            return
        }
        if is_host && rooms[room_name].host.Conn == nil {
            rooms[room_name] = &Room{name: room_name, host: Player{Conn: conn}}
            // TODO: display error when trting to join room with host already
        } else if rooms[room_name].num_guests == MAX_GUESTS {
            // TODO: error in client
            log.Println("room is full")
        } else {
            num_guests := rooms[room_name].num_guests
            rooms[room_name].guests[num_guests] = Player{Conn: conn}
            rooms[room_name].num_guests++
        }
    })
    http.HandleFunc("/rooms/create", func(w http.ResponseWriter, r *http.Request) {
        if r.Method == http.MethodGet {
            html.CreateRoomModal().Render(r.Context(), w)
        } else if r.Method == http.MethodPost {
            name := r.FormValue("name")
            // TODO: error handling (form validation!)
            log.Println("create room", name)
            location := "/rooms?q="+name+"&host=true"
            w.Header().Set("HX-Redirect", location)
            w.WriteHeader(http.StatusSeeOther)
            rooms[name] = &Room{name: name}
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
            // TODO: add player names once they are collected
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
