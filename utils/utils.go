package utils

import (
	"context"
	"github.com/a-h/templ"
	"github.com/gorilla/websocket"
	"log"
)

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
