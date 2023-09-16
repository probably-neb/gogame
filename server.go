package main

import (
	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
	"gogame/player"
    "gogame/sessions"
	"log"
	"net/http"
	"os"
	"time"
    "net/url"
    "strings"
)

const MAX_GUESTS = 5

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	// TODO:actually check origin
	CheckOrigin: func(r *http.Request) bool { return true },
}

func createRoomHandler(rr *RoomRegistry, w http.ResponseWriter, r *http.Request) {
    sessionId := r.Header.Get("Session-Id")
    log.Println("create room with sessionId=", sessionId)
	switch r.Method {
	case http.MethodGet:
		CreateRoomPage(sessionId).Render(r.Context(), w)
	case http.MethodPost:
		roomId := r.FormValue("room-name")
		if rr.RoomExists(roomId) {
			log.Println("error: attempt to create room that already exists:", roomId)
			http.Error(w, "Room already exists", http.StatusBadRequest)
			return
		}
		room := NewRoom{Name: roomId, HostSessionId: sessionId}
		rr.Register <- room
        hostSession, ok := rr.SessionManager.Get(sessionId)
        if !ok {
            http.Error(w, "Host session does not exist", http.StatusBadRequest)
            return
        }
        hostName := hostSession.Name
        hroom := HRoom {Name: roomId, Host: hostName, Guests: []string{}}
		w.Header().Set("HX-Replace-Url", "/rooms/"+roomId)
		HostRoomPage(hroom, sessionId).Render(r.Context(), w)
		log.Println("created room:", roomId)
		// TODO: MethodDELETE
	}
}

func joinRoomHandler(rr *RoomRegistry, w http.ResponseWriter, r *http.Request) {
	// TODO: use switch on method like createRoomHandler instead of joining websocket and waiting for name
	// maybe create playerId and include it in ws url for connecting name to player
	// question is how to handle accepting/rejecting players w/o keeping the request open forever
	roomId := mux.Vars(r)["roomid"]
	if !rr.RoomExists(roomId) {
		errmsg := "error: tried to join non-existant room: " + roomId
		log.Println(errmsg)
		http.Error(w, errmsg, http.StatusBadRequest)
	}
    sessionId := r.Header.Get("Session-Id")
    // FIXME: joining room without session doesn't prompt for name
    if sessionId == "" {
        // TODO: handle by asking player for name in modal
        log.Println("player tried to join room without session id")
        sessionId = rr.SessionManager.NewSession()
        JoinRoomPage(rr.Room(roomId).toHTMLRoom(), sessionId).Render(r.Context(), w)
        return
    }
    GuestRoomPage(rr.Room(roomId).toHTMLRoom(), sessionId).Render(r.Context(), w)
}

func wsHandler(rr *RoomRegistry, w http.ResponseWriter, r *http.Request) {
	roomId := mux.Vars(r)["roomid"]
    sessionId := mux.Vars(r)["sessionid"]
	if !rr.RoomExists(roomId) {
		log.Println("error: attempt to connect to non-existant room:", roomId)
		http.Error(w, "error: attempt to connect to non-existant room:", http.StatusBadRequest)
        return
	}
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("failed to upgrade connection:", err)
		http.Error(w, "failed to upgrade connection", http.StatusInternalServerError)
        return
	}
    playerSession, ok := rr.SessionManager.Get(sessionId)
    if !ok {
        http.Error(w, "player session does not exist", http.StatusBadRequest)
        conn.Close()
        return
    }
    p := player.NewPlayer(conn, playerSession.Name, sessionId)
    rr.Join <- JoinRequest{RoomId: roomId,Player: &p}
}

func roomBrowserHandler(rr *RoomRegistry, w http.ResponseWriter, r *http.Request) {
	roomInfos := rr.GetHTMLRooms()
	// in case of hx-boost don't render (and send!) the whole page
	if r.Header.Get("HX-Request") == "true" {
        sessionId := r.Header.Get("Session-Id")
        if sessionId == "" {
            errmsg := "error: hx-request to /rooms but no sessionId included in request"
            log.Printf(errmsg)
            http.Error(w, errmsg, http.StatusBadRequest)
        }
		RoomList(roomInfos, sessionId).Render(r.Context(), w)
	} else {
        // NOTE: session will be created once player tries to join a room
        sessionId := ""
		RoomListPage(roomInfos, sessionId).Render(r.Context(), w)
	}
}

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.Method, r.RequestURI)
		next.ServeHTTP(w, r)
	})
}

func landingHandler(m *sessions.Manager, w http.ResponseWriter, r *http.Request) {
    sessionId := m.NewSession()
    if sessionId == "" {
        log.Println("failed to generate session id")
        http.Error(w, "Failed to generate session id", http.StatusInternalServerError)
        return
    }
    LandingPage(sessionId).Render(r.Context(), w)
}

func handleSessionUpdate(m *sessions.Manager, w http.ResponseWriter, r *http.Request) {
    sessionId := r.Header.Get("Session-Id")
    name := r.FormValue("display-name")
    currentUrl, err := url.Parse(r.Header.Get("HX-Current-Url"))
    if err != nil {
        log.Println("error: attempted to access session data not through htmx")
        http.Error(w, "sessions must be mutated through the website", http.StatusBadRequest)
        return
    }
    log.Println("current url:", currentUrl.Path)
    err = m.Set(sessionId, "Name", name)
    log.Printf("set player with sessionid=%s name to %s\n", sessionId, name)
    if strings.HasPrefix(currentUrl.Path, "/rooms/") {
        // NOTE: use HX-Location instead of normal redirect
        // so htmx event listener that includes session id in request to room is ran
        w.Header().Set("HX-Location", currentUrl.Path)
        return
    }
    if err != nil {
        log.Println(r)
        log.Println(err)
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    SessionInput(name).Render(r.Context(), w)
}

func main() {
    m := sessions.NewManager()
	rr := newRoomRegistry(m)
	go rr.run()
	rrwrap := func(handler func(reg *RoomRegistry, w http.ResponseWriter, r *http.Request)) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			handler(&rr, w, r)
		}
	}
	rt := mux.NewRouter()

	rt.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        landingHandler(m, w, r)
    })
    rt.HandleFunc("/sessions", func(w http.ResponseWriter, r *http.Request) {
        handleSessionUpdate(m, w, r)
    })
	rt.HandleFunc("/rooms", rrwrap(roomBrowserHandler))
	rt.HandleFunc("/rooms/create", rrwrap(createRoomHandler))

	rmrt := rt.PathPrefix("/rooms/{roomid}").Subrouter()
	rmrt.HandleFunc("", rrwrap(joinRoomHandler))
	rmrt.HandleFunc("/{sessionid}/ws", rrwrap(wsHandler))

	// helper function to do hmtx client side actions (like hx-push-url, or hx-swap="delete")
	// especially when using a websocket connection that cannot send responses in the same way
	rt.HandleFunc("/ok", func(w http.ResponseWriter, r *http.Request) {
		// NOTE: cannot be http.StatusNoContent because this is used with hx-swap="delete"
		// and htmx will not delete the element with a StatusNoContent response
		w.WriteHeader(http.StatusOK)
	})

	rt.HandleFunc("/assets/{asset}", func(w http.ResponseWriter, r *http.Request) {
		asset := mux.Vars(r)["asset"]
		path := "/assets/" + asset
		file, err := os.Open("." + path)
		if err != nil {
			log.Printf("Error opening file: %v\n", err)
			return
		}
		defer file.Close()
		http.ServeContent(w, r, path, time.Now(), file)
	})

	rt.Use(loggingMiddleware)

	addr := "127.0.0.1"
	port := ":8080"

	isProd := os.Getenv("ENV") == "PRODUCTION"
	if isProd {
		addr = "0.0.0.0"
	}

	s := &http.Server{
		Handler:      rt,
		Addr:         addr + port,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	log.Fatal(s.ListenAndServe())
}
