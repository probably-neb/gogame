package main

import (
	"github.com/a-h/templ"
	. "gogame/player"
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
	RoomId string
	Player *Player
	IsHost bool
}

type Room struct {
	name   string
	host   *Player
	guests map[*Player]bool
	join   chan JoinRequest
	recv   chan PlayerMsg
}

func makeRoom(name string, hostName string) Room {
	return Room{
		name:   name,
		host:   &Player{DisplayName: hostName},
		guests: make(map[*Player]bool),
		join:   make(chan JoinRequest),
		recv:   make(chan PlayerMsg),
	}
}

func (r *Room) run() {
	for {
		select {
		case jrq := <-r.join:
			go r.HandleJoinRequest(jrq)
		case msg := <-r.recv:
            // TODO: disconnection logic
			switch msg.Message.Type {
			case "room":
				log.Println("info: room received message:", string(msg.Message.Data))
			case "game":
				log.Println("unimplemented: handling of game messages")
			}
		}
	}
}

func (r *Room) HandleJoinRequest(jrq JoinRequest) {
	if jrq.IsHost {
		jrq.Player.DisplayName = r.host.DisplayName
		r.host = jrq.Player
		go r.host.ListenForMessages(r.recv)
		go r.host.WriteMessages()
		return
	}

	guest := jrq.Player
	// go guest.ListenForMessages( /* TODO: make channel here */ )
	go guest.WriteMessages()
	// close join modal
	guest.Send <- CloseJoinModal()
	// update url
	guest.Send <- JoinRoomRedirect(r.name)
	// update guest list for host and other guests
	shouldAppend := true
	updatedGuestList := []string{guest.DisplayName}
	r.Broadcast(GuestList(updatedGuestList, shouldAppend))
	// add guest to guests
	r.guests[guest] = true
	// finally update entire room body for new guest
	guest.Send <- RoomPageBody(r.toHTMLRoom(), "guest")
}

func (r *Room) Broadcast(c templ.Component) {
	r.host.Send <- c
	for g, ok := range r.guests {
		if !ok {
			continue
		}
		g.Send <- c
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

type RoomRegistry struct {
	rooms      map[string]*Room
	Register   chan NewRoom
	Unregister chan string
	Join       chan JoinRequest
}

func newRoomRegistry() RoomRegistry {
	return RoomRegistry{
		rooms:      make(map[string]*Room),
		Register:   make(chan NewRoom),
		Unregister: make(chan string),
		Join:       make(chan JoinRequest),
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
			room, ok := rr.rooms[joinRequest.RoomId]
			if !ok {
				log.Println("error: attempt to join non-existant room:", joinRequest.RoomId)
				continue
			}
			room.join <- joinRequest
		}
	}
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

func (rr *RoomRegistry) Room(roomId string) *Room {
	return rr.rooms[roomId]
}
