package main

import (
    "fmt"
    "github.com/a-h/templ"
    "github.com/gorilla/websocket"
    "log"
    "net/http"
    "gogame/html"
    "gogame/htmx"
    "encoding/json"
    "context"
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
            Headers htmx.Headers `json:"HEADERS"`
            DisplayName string `json:"display-name"`
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
            break;
        }
        var msg GuestJoinMsg
        if err = json.Unmarshal(p, &msg); err != nil {
            return err
        }
        log.Println(msg)
        guest := Player{Conn: conn, DisplayName: msg.DisplayName};
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

func main() {
    // TODO: store this in db
    var rooms = make(map[string]*Room)

    http.Handle("/", templ.Handler(html.LandingPage()))
    http.HandleFunc("/rooms", func(w http.ResponseWriter, r *http.Request) {
        room_name := r.URL.Query().Get("q")
        isNewGuest := r.URL.Query().Get("host") != "true"
        if room_name == "" {
            log.Println("tried to join room without query")
            return
        }
        room := rooms[room_name]
        log.Println("access to room: [", room.name, "]", "by")
        html.RoomPage(room.toHTMLRoom(), isNewGuest).Render(r.Context(), w)
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
        room_name := r.URL.Query().Get("q")
        is_host := r.URL.Query().Get("host") == "true"
        if room_name == "" {
            log.Println("no room name specified for websocket connection")
            return
        }
        if rooms[room_name] == nil {
            log.Println("room does not exist")
            // TODO: redirect to home (before upgrading connection!)
            return
        }
        if is_host && rooms[room_name].host.Conn == nil {
            rooms[room_name].host.Conn = conn
            log.Println("websocket connection established to host in room:",room_name)
            go rooms[room_name].Start();
            // TODO: display error when trying to join room with host already
        } else if len(rooms[room_name].guests) == MAX_GUESTS {
            // TODO: error in client
            log.Println("room is full")
        } else {
            log.Println("websocket connection established to new guest in room:", room_name)
            go rooms[room_name].ConnectGuest(conn)
        }
    })
    http.HandleFunc("/rooms/create", func(w http.ResponseWriter, r *http.Request) {
        if r.Method == http.MethodGet {
            html.CreateRoomModal().Render(r.Context(), w)
        } else if r.Method == http.MethodPost {
            // TODO: error handling (form validation!)
            room_name := r.FormValue("room-name")
            display_name := r.FormValue("display-name")
            log.Println("create room", room_name, "host:", display_name)
            location := "/rooms?q="+room_name+"&host=true"
            w.Header().Set("HX-Redirect", location)
            w.WriteHeader(http.StatusSeeOther)
            rooms[room_name] = &Room{name: room_name, host: Player{DisplayName: display_name}}
        }
    })
    http.HandleFunc("ok", func(w http.ResponseWriter, r *http.Request) {
        // NOTE: cannot be http.StatusNoContent because this is used with hx-swap="delete"
        // and htmx will not delete the element with a StatusNoContent response
        w.WriteHeader(http.StatusOK)
    })
    http.HandleFunc("/rooms/join", func(w http.ResponseWriter, r *http.Request) {
        var roomInfos = make([]html.Room, len(rooms));
        i := 0
        for _, room := range rooms {
            if room.name == "" {
                continue;
            }
            // TODO: add player names once they are collected
            roomInfos[i] = room.toHTMLRoom();
            i++;
        }
        html.JoinRoomPage(roomInfos).Render(r.Context(), w)
    })

    fmt.Println("Server started at 8080 port")
    //Use the default DefaultServeMux.
    err := http.ListenAndServe(":8080", nil)
    if err != nil {
        log.Fatal(err)
    }
}
