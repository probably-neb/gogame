package utils

import ("github.com/gorilla/websocket"
"github.com/a-h/templ"
"log"
"context")



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

