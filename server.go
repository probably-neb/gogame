package main

import (
	"github.com/a-h/templ"
	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
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

func joinRoomHandler(w http.ResponseWriter, r *http.Request) {
    // TODO:
}

func guestWSHandler(room *Room, conn *websocket.Conn, w http.ResponseWriter, r *http.Request) {
    // TODO:
}

func hostWSHandler(room *Room, conn *websocket.Conn, w http.ResponseWriter, r *http.Request) {
    // TODO:
}

func createRoomHandler(rr RoomRegistry, w http.ResponseWriter, r *http.Request) {
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
        rr.Register<- room
		w.Header().Set("HX-Replace-Url", "/rooms/"+roomId)
        HostRoomPage(room.toHTMLRoom()).Render(r.Context(), w)
        log.Println("created room:", roomId)
    // TODO: MethodDELETE
	}
}


func wsHandler(w http.ResponseWriter, r *http.Request) {
    // TODO:

	// TODO: redirect to home in case of error
}

func roomBrowserHandler(rr RoomRegistry, w http.ResponseWriter, r *http.Request) {
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
	// TODO: store this in db
    rr := newRoomRegistry();
    go rr.run()
    rrwrap := func(handler func(reg RoomRegistry, w http.ResponseWriter, r *http.Request)) http.HandlerFunc {
        return func(w http.ResponseWriter, r *http.Request) {
            handler(rr, w, r)
        }
    }
	rt := mux.NewRouter()

	rt.Handle("/", templ.Handler(LandingPage()))
	rt.HandleFunc("/rooms", rrwrap(roomBrowserHandler))
	rt.HandleFunc("/rooms/create", rrwrap(createRoomHandler))

	rmrt := rt.PathPrefix("/rooms/{roomid}").Subrouter()
	rmrt.HandleFunc("/{kind}/ws", wsHandler)
	rmrt.HandleFunc("/join", joinRoomHandler)

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
