package main

import (
	. "gogame/player"
    "gogame/htmx"
    "log"
)

type NewRoom struct {
	Name     string
	HostName string
}

func (nr *NewRoom) toHTMLRoom() HRoom {
	return HRoom{Name: nr.Name, Host: nr.HostName, Guests: []string{}}
}

type JoinRequest struct {
	roomId string
    isHost bool
	player *Player
}


type Room struct {
	name   string
	host   *Player
	guests map[*Player]bool
	join   chan JoinRequest
}

func makeRoom(name string, hostName string) Room {
	return Room{
        name: name,
        host: &Player{DisplayName: hostName},
        guests: make(map[*Player]bool),
        join: make(chan JoinRequest),
    }
}

func (r *Room) run() {
	for {
		select {
            // case jrq := <-r.join:
        }
	}
}

func (r *Room) toHTMLRoom() HRoom {
	players := make([]string, len(r.guests))
    i := 0
	for guest, ok := range r.guests {
        if !ok || guest == nil {
            continue
        }
		players[i] = guest.DisplayName
        i++
	}
	return HRoom{Name: r.name, Host: r.host.DisplayName, Guests: players}
}

func (r *Room) ListenForHostMessages() {
    // TODO: 
}

func (r *Room) StartGame(game string) {
    // TODO:
}

func (r *Room) DisconnectHost(err error) {
	// TODO: redirect players to lobby
}

func (r *Room) DisconnectGuest(g Player, err error) {
	// TODO: remove guest from guests
}

func (r *Room) ConnectGuest() {
	type GuestJoinMsg struct {
		Headers     htmx.Headers `json:"HEADERS"`
		DisplayName string       `json:"display-name"`
	}
	// guest := Player{conn: conn, DisplayName: msg.DisplayName}
	// // close join modal
	// conn.WriteMessage(websocket.TextMessage, []byte(`<div id="join-room-modal"></div>`))
	// // update url
	// rhtml := r.toHTMLRoom()
	// err = utils.ConnWriteTemplate(conn, JoinRoomRedirect(rhtml))
	// if err != nil {
	// 	log.Println(err)
	// 	return
	// }
	// // update room body
	// err = utils.ConnWriteTemplate(conn, RoomPageBody(rhtml, "guest"))
	// if err != nil {
	// 	log.Println(err)
	// 	return
	// }
	//
	// // TODO: update all connections with new guest
	// newGuestNameList := []string{guest.DisplayName}
	// appendGuestList := true
	// newGuestList := GuestList(newGuestNameList, appendGuestList)
}

type RoomRegistry struct {
	rooms      map[string]*Room
	Register   chan NewRoom
	Unregister chan string
	Join       chan JoinRequest
}

func newRoomRegistry() RoomRegistry {
    return RoomRegistry {
        rooms: make(map[string]*Room),
        Register: make(chan NewRoom),
        Unregister: make(chan string),
        Join: make(chan JoinRequest),
    }
}

func (rr *RoomRegistry) run() {
	for {
		select {
		case newRoom := <-rr.Register:
			room := makeRoom(newRoom.Name, newRoom.HostName)
			rr.rooms[room.name] = &room
			go room.run()
		case roomId := <-rr.Unregister:
			delete(rr.rooms, roomId)
		case joinRequest := <-rr.Join:
			room, ok := rr.rooms[joinRequest.roomId]
			if !ok {
				log.Println("error: attempt to join non-existant room:", joinRequest.roomId)
                continue
			}
			room.join <- joinRequest
		}
	}
}

func (rr *RoomRegistry) Room(id string) *Room {
    return rr.rooms[id]
}

func (rr *RoomRegistry) GetHTMLRooms() []HRoom {
	var hrooms = make([]HRoom, len(rr.rooms))
	i := 0
	for id, room := range rr.rooms {
		if id == "" || room == nil {
			continue
		}
		// TODO: add player names once they are collected
		hrooms[i] = room.toHTMLRoom()
		i++
	}
    return hrooms
}

func (rr *RoomRegistry) RoomExists(roomId string) bool {
    return roomId != "" && rr.rooms[roomId] != nil
}
