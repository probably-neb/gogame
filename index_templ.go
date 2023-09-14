// Code generated by templ@v0.2.316 DO NOT EDIT.

package main

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import "context"
import "io"
import "bytes"

import "fmt"

func Layout(title string) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) (err error) {
		templBuffer, templIsBuffer := w.(*bytes.Buffer)
		if !templIsBuffer {
			templBuffer = templ.GetBuffer()
			defer templ.ReleaseBuffer(templBuffer)
		}
		ctx = templ.InitializeContext(ctx)
		var_1 := templ.GetChildren(ctx)
		if var_1 == nil {
			var_1 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		_, err = templBuffer.WriteString("<!doctype html><html lang=\"en\"><head><meta charset=\"UTF-8\"><title>")
		if err != nil {
			return err
		}
		var var_2 string = title
		_, err = templBuffer.WriteString(templ.EscapeString(var_2))
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString("</title><link rel=\"icon\" type=\"image/svg+xml\" href=\"/dist/favicon.svg\"><meta name=\"viewport\" content=\"width=device-width, initial-scale=1\"><meta name=\"robots\" content=\"index, follow\"><meta name=\"revisit-after\" content=\"7 days\"><meta name=\"language\" content=\"English\"><script src=\"https://unpkg.com/htmx.org@1.9.5/dist/htmx.js\">")
		if err != nil {
			return err
		}
		var_3 := ``
		_, err = templBuffer.WriteString(var_3)
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString("</script><script src=\"https://cdn.tailwindcss.com\">")
		if err != nil {
			return err
		}
		var_4 := ``
		_, err = templBuffer.WriteString(var_4)
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString("</script><script src=\"https://unpkg.com/htmx.org@1.9.5/dist/ext/ws.js\">")
		if err != nil {
			return err
		}
		var_5 := ``
		_, err = templBuffer.WriteString(var_5)
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString("</script><script>")
		if err != nil {
			return err
		}
		var_6 := `
                htmx.on("htmx:wsConfigSend", function(e) {
                    console.dir(e, {depth: null})
                    let ty = e.detail.elt.dataset.type
                    let grp = e.detail.elt.dataset.group
                    e.detail.parameters = {type: ty, group: grp, data: e.detail.parameters}
                })
            `
		_, err = templBuffer.WriteString(var_6)
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString("</script></head><body id=\"body\">")
		if err != nil {
			return err
		}
		err = var_1.Render(ctx, templBuffer)
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString("</body></html>")
		if err != nil {
			return err
		}
		if !templIsBuffer {
			_, err = io.Copy(w, templBuffer)
		}
		return err
	})
}

func RoomURL(id string, rest string) string {
	if len(rest) > 0 && rest[0] != '/' {
		rest = "/" + rest
	}
	return "/rooms/" + id + rest
}

func LandingPage() templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) (err error) {
		templBuffer, templIsBuffer := w.(*bytes.Buffer)
		if !templIsBuffer {
			templBuffer = templ.GetBuffer()
			defer templ.ReleaseBuffer(templBuffer)
		}
		ctx = templ.InitializeContext(ctx)
		var_7 := templ.GetChildren(ctx)
		if var_7 == nil {
			var_7 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		var_8 := templ.ComponentFunc(func(ctx context.Context, w io.Writer) (err error) {
			templBuffer, templIsBuffer := w.(*bytes.Buffer)
			if !templIsBuffer {
				templBuffer = templ.GetBuffer()
				defer templ.ReleaseBuffer(templBuffer)
			}
			_, err = templBuffer.WriteString("<div hx-boost=\"true\"><button id=\"create-room\" class=\"rounded-md border-2 border-black\"><a href=\"/rooms/create\">")
			if err != nil {
				return err
			}
			var_9 := `Create Room`
			_, err = templBuffer.WriteString(var_9)
			if err != nil {
				return err
			}
			_, err = templBuffer.WriteString("</a></button><button id=\"join-room\" class=\"rounded-md border-2 border-black\"><a href=\"/rooms\" hx-boost=\"true\">")
			if err != nil {
				return err
			}
			var_10 := `Join Room`
			_, err = templBuffer.WriteString(var_10)
			if err != nil {
				return err
			}
			_, err = templBuffer.WriteString("</a></button></div>")
			if err != nil {
				return err
			}
			if !templIsBuffer {
				_, err = io.Copy(w, templBuffer)
			}
			return err
		})
		err = Layout("GOGAME!").Render(templ.WithChildren(ctx, var_8), templBuffer)
		if err != nil {
			return err
		}
		if !templIsBuffer {
			_, err = io.Copy(w, templBuffer)
		}
		return err
	})
}

// CREDIT: shadcn-ui dialog

func Modal(id string) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) (err error) {
		templBuffer, templIsBuffer := w.(*bytes.Buffer)
		if !templIsBuffer {
			templBuffer = templ.GetBuffer()
			defer templ.ReleaseBuffer(templBuffer)
		}
		ctx = templ.InitializeContext(ctx)
		var_11 := templ.GetChildren(ctx)
		if var_11 == nil {
			var_11 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		_, err = templBuffer.WriteString("<div id=\"")
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString(templ.EscapeString(id))
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString("\"><div data-state=\"open\" class=\"fixed inset-0 z-50 bg-background/80 backdrop-blur-sm data-[state=open]:animate-in data-[state=closed]:animate-out data-[state=closed]:fade-out-0 data-[state=open]:fade-in-0\" style=\"pointer-events: auto;\" data-aria-hidden=\"true\" aria-hidden=\"true\"></div><div role=\"dialog\" data-state=\"open\" class=\"fixed left-[50%] top-[50%] z-50 grid w-full max-w-lg translate-x-[-50%] translate-y-[-50%] gap-4 border bg-background p-6 shadow-lg duration-200 data-[state=open]:animate-in data-[state=closed]:animate-out data-[state=closed]:fade-out-0 data-[state=open]:fade-in-0 data-[state=closed]:zoom-out-95 data-[state=open]:zoom-in-95 data-[state=closed]:slide-out-to-left-1/2 data-[state=closed]:slide-out-to-top-[48%] data-[state=open]:slide-in-from-left-1/2 data-[state=open]:slide-in-from-top-[48%] sm:rounded-lg md:w-full sm:max-w-[425px]\" tabindex=\"-1\" style=\"pointer-events: auto;\">")
		if err != nil {
			return err
		}
		err = var_11.Render(ctx, templBuffer)
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString("</div></div>")
		if err != nil {
			return err
		}
		if !templIsBuffer {
			_, err = io.Copy(w, templBuffer)
		}
		return err
	})
}

func CreateRoomModal() templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) (err error) {
		templBuffer, templIsBuffer := w.(*bytes.Buffer)
		if !templIsBuffer {
			templBuffer = templ.GetBuffer()
			defer templ.ReleaseBuffer(templBuffer)
		}
		ctx = templ.InitializeContext(ctx)
		var_12 := templ.GetChildren(ctx)
		if var_12 == nil {
			var_12 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		var_13 := templ.ComponentFunc(func(ctx context.Context, w io.Writer) (err error) {
			templBuffer, templIsBuffer := w.(*bytes.Buffer)
			if !templIsBuffer {
				templBuffer = templ.GetBuffer()
				defer templ.ReleaseBuffer(templBuffer)
			}
			_, err = templBuffer.WriteString("<form class=\"text-small text-gray-800\" hx-post=\"/rooms/create\" hx-boost=\"true\" hx-target=\"body\"><label for=\"room-name\">")
			if err != nil {
				return err
			}
			var_14 := `Room Name:`
			_, err = templBuffer.WriteString(var_14)
			if err != nil {
				return err
			}
			_, err = templBuffer.WriteString("</label><input type=\"text\" name=\"room-name\" id=\"room-name\" placeholder=\"Room Name\" class=\"rounded-md border-2 border-black\"><label for=\"display-name\">")
			if err != nil {
				return err
			}
			var_15 := `Display Name:`
			_, err = templBuffer.WriteString(var_15)
			if err != nil {
				return err
			}
			_, err = templBuffer.WriteString("</label><input type=\"text\" name=\"display-name\" id=\"display-name\" placeholder=\"Display Name\" class=\"rounded-md border-2 border-black\"><button id=\"create\" type=\"submit\" class=\"rounded-md border-2 border-black\">")
			if err != nil {
				return err
			}
			var_16 := `Create`
			_, err = templBuffer.WriteString(var_16)
			if err != nil {
				return err
			}
			_, err = templBuffer.WriteString("</button><button type=\"button\" class=\"mt-3 inline-flex w-full justify-center rounded-md bg-white px-3 py-2 text-sm font-semibold text-gray-900 shadow-sm ring-1 ring-inset ring-gray-300 hover:bg-gray-50 sm:mt-0 sm:w-auto\"><a href=\"/\" hx-boost=\"true\">")
			if err != nil {
				return err
			}
			var_17 := `Cancel`
			_, err = templBuffer.WriteString(var_17)
			if err != nil {
				return err
			}
			_, err = templBuffer.WriteString("</a></button></form>")
			if err != nil {
				return err
			}
			if !templIsBuffer {
				_, err = io.Copy(w, templBuffer)
			}
			return err
		})
		err = Modal("create-room-modal").Render(templ.WithChildren(ctx, var_13), templBuffer)
		if err != nil {
			return err
		}
		if !templIsBuffer {
			_, err = io.Copy(w, templBuffer)
		}
		return err
	})
}

func JoinRoomModal() templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) (err error) {
		templBuffer, templIsBuffer := w.(*bytes.Buffer)
		if !templIsBuffer {
			templBuffer = templ.GetBuffer()
			defer templ.ReleaseBuffer(templBuffer)
		}
		ctx = templ.InitializeContext(ctx)
		var_18 := templ.GetChildren(ctx)
		if var_18 == nil {
			var_18 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		var_19 := templ.ComponentFunc(func(ctx context.Context, w io.Writer) (err error) {
			templBuffer, templIsBuffer := w.(*bytes.Buffer)
			if !templIsBuffer {
				templBuffer = templ.GetBuffer()
				defer templ.ReleaseBuffer(templBuffer)
			}
			_, err = templBuffer.WriteString("<form class=\"text-small text-gray-800\" ws-send data-group=\"room\" data-type=\"display-name\"><label for=\"display-name\">")
			if err != nil {
				return err
			}
			var_20 := `Display Name:`
			_, err = templBuffer.WriteString(var_20)
			if err != nil {
				return err
			}
			_, err = templBuffer.WriteString("</label><input type=\"text\" name=\"display-name\" id=\"display-name\" placeholder=\"Display Name\" class=\"rounded-md border-2 border-black\"><button id=\"join\" type=\"submit\" class=\"rounded-md border-2 border-black\">")
			if err != nil {
				return err
			}
			var_21 := `Join`
			_, err = templBuffer.WriteString(var_21)
			if err != nil {
				return err
			}
			_, err = templBuffer.WriteString("</button><button type=\"button\" class=\"mt-3 inline-flex w-full justify-center rounded-md bg-white px-3 py-2 text-sm font-semibold text-gray-900 shadow-sm ring-1 ring-inset ring-gray-300 hover:bg-gray-50 sm:mt-0 sm:w-auto\"><a href=\"/rooms\" hx-boost=\"true\">")
			if err != nil {
				return err
			}
			var_22 := `Cancel`
			_, err = templBuffer.WriteString(var_22)
			if err != nil {
				return err
			}
			_, err = templBuffer.WriteString("</a></button></form>")
			if err != nil {
				return err
			}
			if !templIsBuffer {
				_, err = io.Copy(w, templBuffer)
			}
			return err
		})
		err = Modal("join-room-modal").Render(templ.WithChildren(ctx, var_19), templBuffer)
		if err != nil {
			return err
		}
		if !templIsBuffer {
			_, err = io.Copy(w, templBuffer)
		}
		return err
	})
}

// TODO: remove redeclaration of Room, make Room have public fields instead
type HRoom struct {
	Name   string
	Host   string
	Guests []string
}

func LobbyGuest(name string) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) (err error) {
		templBuffer, templIsBuffer := w.(*bytes.Buffer)
		if !templIsBuffer {
			templBuffer = templ.GetBuffer()
			defer templ.ReleaseBuffer(templBuffer)
		}
		ctx = templ.InitializeContext(ctx)
		var_23 := templ.GetChildren(ctx)
		if var_23 == nil {
			var_23 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		_, err = templBuffer.WriteString("<li id=\"")
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString(templ.EscapeString(name))
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString("\">")
		if err != nil {
			return err
		}
		var var_24 string = name
		_, err = templBuffer.WriteString(templ.EscapeString(var_24))
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString("</li>")
		if err != nil {
			return err
		}
		if !templIsBuffer {
			_, err = io.Copy(w, templBuffer)
		}
		return err
	})
}

// TODO: add kind var for kick functionality

func GuestList(guests []string, append bool) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) (err error) {
		templBuffer, templIsBuffer := w.(*bytes.Buffer)
		if !templIsBuffer {
			templBuffer = templ.GetBuffer()
			defer templ.ReleaseBuffer(templBuffer)
		}
		ctx = templ.InitializeContext(ctx)
		var_25 := templ.GetChildren(ctx)
		if var_25 == nil {
			var_25 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		_, err = templBuffer.WriteString("<ul id=\"players\"")
		if err != nil {
			return err
		}
		if append {
			_, err = templBuffer.WriteString(" hx-swap-oob=\"beforeend\"")
			if err != nil {
				return err
			}
		}
		_, err = templBuffer.WriteString(">")
		if err != nil {
			return err
		}
		for _, guest := range guests {
			err = LobbyGuest(guest).Render(ctx, templBuffer)
			if err != nil {
				return err
			}
		}
		_, err = templBuffer.WriteString("</ul>")
		if err != nil {
			return err
		}
		if !templIsBuffer {
			_, err = io.Copy(w, templBuffer)
		}
		return err
	})
}

func GamesList(kind string) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) (err error) {
		templBuffer, templIsBuffer := w.(*bytes.Buffer)
		if !templIsBuffer {
			templBuffer = templ.GetBuffer()
			defer templ.ReleaseBuffer(templBuffer)
		}
		ctx = templ.InitializeContext(ctx)
		var_26 := templ.GetChildren(ctx)
		if var_26 == nil {
			var_26 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		_, err = templBuffer.WriteString("<div><h3>")
		if err != nil {
			return err
		}
		var_27 := `Games`
		_, err = templBuffer.WriteString(var_27)
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString("</h3><ul><li>")
		if err != nil {
			return err
		}
		var_28 := `Tic-Tac-Toe`
		_, err = templBuffer.WriteString(var_28)
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString(" ")
		if err != nil {
			return err
		}
		if kind == "host" {
			_, err = templBuffer.WriteString("<button id=\"start-tic-tac-toe\" name=\"play\" value=\"tictactoe\" ws-send data-group=\"room\" data-type=\"play\">")
			if err != nil {
				return err
			}
			var_29 := `Play!`
			_, err = templBuffer.WriteString(var_29)
			if err != nil {
				return err
			}
			_, err = templBuffer.WriteString("</button>")
			if err != nil {
				return err
			}
		}
		_, err = templBuffer.WriteString("</li></ul></div>")
		if err != nil {
			return err
		}
		if !templIsBuffer {
			_, err = io.Copy(w, templBuffer)
		}
		return err
	})
}

func RoomWS(roomid string, kind string) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) (err error) {
		templBuffer, templIsBuffer := w.(*bytes.Buffer)
		if !templIsBuffer {
			templBuffer = templ.GetBuffer()
			defer templ.ReleaseBuffer(templBuffer)
		}
		ctx = templ.InitializeContext(ctx)
		var_30 := templ.GetChildren(ctx)
		if var_30 == nil {
			var_30 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		_, err = templBuffer.WriteString("<div id=\"ws-connection\" hx-preserve=\"true\" hx-ext=\"ws\" ws-connect=\"")
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString(templ.EscapeString(RoomURL(roomid, kind+"/ws")))
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString("\">")
		if err != nil {
			return err
		}
		err = var_30.Render(ctx, templBuffer)
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString("</div>")
		if err != nil {
			return err
		}
		if !templIsBuffer {
			_, err = io.Copy(w, templBuffer)
		}
		return err
	})
}

func RoomPageBody(room HRoom, kind string) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) (err error) {
		templBuffer, templIsBuffer := w.(*bytes.Buffer)
		if !templIsBuffer {
			templBuffer = templ.GetBuffer()
			defer templ.ReleaseBuffer(templBuffer)
		}
		ctx = templ.InitializeContext(ctx)
		var_31 := templ.GetChildren(ctx)
		if var_31 == nil {
			var_31 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		_, err = templBuffer.WriteString("<div id=\"room-page-body\"><h3>")
		if err != nil {
			return err
		}
		var var_32 string = room.Name
		_, err = templBuffer.WriteString(templ.EscapeString(var_32))
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString("</h3><h2>")
		if err != nil {
			return err
		}
		var_33 := `Host:`
		_, err = templBuffer.WriteString(var_33)
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString("</h2><p id=\"host\">")
		if err != nil {
			return err
		}
		var var_34 string = room.Host
		_, err = templBuffer.WriteString(templ.EscapeString(var_34))
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString("</p><h2>")
		if err != nil {
			return err
		}
		var_35 := `Guests:`
		_, err = templBuffer.WriteString(var_35)
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString("</h2>")
		if err != nil {
			return err
		}
		err = GuestList(room.Guests, false).Render(ctx, templBuffer)
		if err != nil {
			return err
		}
		err = GamesList(kind).Render(ctx, templBuffer)
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString("</div>")
		if err != nil {
			return err
		}
		if !templIsBuffer {
			_, err = io.Copy(w, templBuffer)
		}
		return err
	})
}

func CreateRoomPage() templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) (err error) {
		templBuffer, templIsBuffer := w.(*bytes.Buffer)
		if !templIsBuffer {
			templBuffer = templ.GetBuffer()
			defer templ.ReleaseBuffer(templBuffer)
		}
		ctx = templ.InitializeContext(ctx)
		var_36 := templ.GetChildren(ctx)
		if var_36 == nil {
			var_36 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		var_37 := templ.ComponentFunc(func(ctx context.Context, w io.Writer) (err error) {
			templBuffer, templIsBuffer := w.(*bytes.Buffer)
			if !templIsBuffer {
				templBuffer = templ.GetBuffer()
				defer templ.ReleaseBuffer(templBuffer)
			}
			err = CreateRoomModal().Render(ctx, templBuffer)
			if err != nil {
				return err
			}
			_, err = templBuffer.WriteString(" ")
			if err != nil {
				return err
			}
			err = RoomPageBody(HRoom{}, "host").Render(ctx, templBuffer)
			if err != nil {
				return err
			}
			if !templIsBuffer {
				_, err = io.Copy(w, templBuffer)
			}
			return err
		})
		err = Layout("Create Room").Render(templ.WithChildren(ctx, var_37), templBuffer)
		if err != nil {
			return err
		}
		if !templIsBuffer {
			_, err = io.Copy(w, templBuffer)
		}
		return err
	})
}

func GuestRoomPage(room HRoom) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) (err error) {
		templBuffer, templIsBuffer := w.(*bytes.Buffer)
		if !templIsBuffer {
			templBuffer = templ.GetBuffer()
			defer templ.ReleaseBuffer(templBuffer)
		}
		ctx = templ.InitializeContext(ctx)
		var_38 := templ.GetChildren(ctx)
		if var_38 == nil {
			var_38 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		var_39 := templ.ComponentFunc(func(ctx context.Context, w io.Writer) (err error) {
			templBuffer, templIsBuffer := w.(*bytes.Buffer)
			if !templIsBuffer {
				templBuffer = templ.GetBuffer()
				defer templ.ReleaseBuffer(templBuffer)
			}
			var_40 := templ.ComponentFunc(func(ctx context.Context, w io.Writer) (err error) {
				templBuffer, templIsBuffer := w.(*bytes.Buffer)
				if !templIsBuffer {
					templBuffer = templ.GetBuffer()
					defer templ.ReleaseBuffer(templBuffer)
				}
				err = var_38.Render(ctx, templBuffer)
				if err != nil {
					return err
				}
				_, err = templBuffer.WriteString(" ")
				if err != nil {
					return err
				}
				err = RoomPageBody(room, "guest").Render(ctx, templBuffer)
				if err != nil {
					return err
				}
				if !templIsBuffer {
					_, err = io.Copy(w, templBuffer)
				}
				return err
			})
			err = RoomWS(room.Name, "guest").Render(templ.WithChildren(ctx, var_40), templBuffer)
			if err != nil {
				return err
			}
			if !templIsBuffer {
				_, err = io.Copy(w, templBuffer)
			}
			return err
		})
		err = Layout(room.Name).Render(templ.WithChildren(ctx, var_39), templBuffer)
		if err != nil {
			return err
		}
		if !templIsBuffer {
			_, err = io.Copy(w, templBuffer)
		}
		return err
	})
}

func JoinRoomPage(room HRoom) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) (err error) {
		templBuffer, templIsBuffer := w.(*bytes.Buffer)
		if !templIsBuffer {
			templBuffer = templ.GetBuffer()
			defer templ.ReleaseBuffer(templBuffer)
		}
		ctx = templ.InitializeContext(ctx)
		var_41 := templ.GetChildren(ctx)
		if var_41 == nil {
			var_41 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		var_42 := templ.ComponentFunc(func(ctx context.Context, w io.Writer) (err error) {
			templBuffer, templIsBuffer := w.(*bytes.Buffer)
			if !templIsBuffer {
				templBuffer = templ.GetBuffer()
				defer templ.ReleaseBuffer(templBuffer)
			}
			err = JoinRoomModal().Render(ctx, templBuffer)
			if err != nil {
				return err
			}
			if !templIsBuffer {
				_, err = io.Copy(w, templBuffer)
			}
			return err
		})
		err = GuestRoomPage(room).Render(templ.WithChildren(ctx, var_42), templBuffer)
		if err != nil {
			return err
		}
		if !templIsBuffer {
			_, err = io.Copy(w, templBuffer)
		}
		return err
	})
}

func HostRoomPage(room HRoom) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) (err error) {
		templBuffer, templIsBuffer := w.(*bytes.Buffer)
		if !templIsBuffer {
			templBuffer = templ.GetBuffer()
			defer templ.ReleaseBuffer(templBuffer)
		}
		ctx = templ.InitializeContext(ctx)
		var_43 := templ.GetChildren(ctx)
		if var_43 == nil {
			var_43 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		var_44 := templ.ComponentFunc(func(ctx context.Context, w io.Writer) (err error) {
			templBuffer, templIsBuffer := w.(*bytes.Buffer)
			if !templIsBuffer {
				templBuffer = templ.GetBuffer()
				defer templ.ReleaseBuffer(templBuffer)
			}
			var_45 := templ.ComponentFunc(func(ctx context.Context, w io.Writer) (err error) {
				templBuffer, templIsBuffer := w.(*bytes.Buffer)
				if !templIsBuffer {
					templBuffer = templ.GetBuffer()
					defer templ.ReleaseBuffer(templBuffer)
				}
				err = RoomPageBody(room, "host").Render(ctx, templBuffer)
				if err != nil {
					return err
				}
				if !templIsBuffer {
					_, err = io.Copy(w, templBuffer)
				}
				return err
			})
			err = RoomWS(room.Name, "host").Render(templ.WithChildren(ctx, var_45), templBuffer)
			if err != nil {
				return err
			}
			if !templIsBuffer {
				_, err = io.Copy(w, templBuffer)
			}
			return err
		})
		err = Layout(room.Name).Render(templ.WithChildren(ctx, var_44), templBuffer)
		if err != nil {
			return err
		}
		if !templIsBuffer {
			_, err = io.Copy(w, templBuffer)
		}
		return err
	})
}

func JoinRoomRedirect(roomId string) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) (err error) {
		templBuffer, templIsBuffer := w.(*bytes.Buffer)
		if !templIsBuffer {
			templBuffer = templ.GetBuffer()
			defer templ.ReleaseBuffer(templBuffer)
		}
		ctx = templ.InitializeContext(ctx)
		var_46 := templ.GetChildren(ctx)
		if var_46 == nil {
			var_46 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		_, err = templBuffer.WriteString("<div id=\"body\" hx-swap-oob=\"beforeend\"><div hx-get=\"/ok\" hx-replace-url=\"")
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString(templ.EscapeString(RoomURL(roomId, "")))
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString("\" hx-swap=\"none\" hx-trigger=\"load\"></div></div>")
		if err != nil {
			return err
		}
		if !templIsBuffer {
			_, err = io.Copy(w, templBuffer)
		}
		return err
	})
}

func CloseJoinModal() templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) (err error) {
		templBuffer, templIsBuffer := w.(*bytes.Buffer)
		if !templIsBuffer {
			templBuffer = templ.GetBuffer()
			defer templ.ReleaseBuffer(templBuffer)
		}
		ctx = templ.InitializeContext(ctx)
		var_47 := templ.GetChildren(ctx)
		if var_47 == nil {
			var_47 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		_, err = templBuffer.WriteString("<div id=\"join-room-modal\"></div>")
		if err != nil {
			return err
		}
		if !templIsBuffer {
			_, err = io.Copy(w, templBuffer)
		}
		return err
	})
}

func JoinRoomEntry(room HRoom) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) (err error) {
		templBuffer, templIsBuffer := w.(*bytes.Buffer)
		if !templIsBuffer {
			templBuffer = templ.GetBuffer()
			defer templ.ReleaseBuffer(templBuffer)
		}
		ctx = templ.InitializeContext(ctx)
		var_48 := templ.GetChildren(ctx)
		if var_48 == nil {
			var_48 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		_, err = templBuffer.WriteString("<h3>")
		if err != nil {
			return err
		}
		var var_49 string = room.Name
		_, err = templBuffer.WriteString(templ.EscapeString(var_49))
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString("</h3><h2>")
		if err != nil {
			return err
		}
		var_50 := `Host:`
		_, err = templBuffer.WriteString(var_50)
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString("</h2><p>")
		if err != nil {
			return err
		}
		var var_51 string = room.Host
		_, err = templBuffer.WriteString(templ.EscapeString(var_51))
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString("</p><h2>")
		if err != nil {
			return err
		}
		var_52 := `Guests: `
		_, err = templBuffer.WriteString(var_52)
		if err != nil {
			return err
		}
		var var_53 string = fmt.Sprint(len(room.Guests))
		_, err = templBuffer.WriteString(templ.EscapeString(var_53))
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString("</h2><button id=\"join-room\" class=\"rounded-md border-2 border-black\"><a href=\"")
		if err != nil {
			return err
		}
		var var_54 templ.SafeURL = templ.SafeURL(RoomURL(room.Name, "/join"))
		_, err = templBuffer.WriteString(templ.EscapeString(string(var_54)))
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString("\" hx-boost=\"true\">")
		if err != nil {
			return err
		}
		var_55 := `Join Room`
		_, err = templBuffer.WriteString(var_55)
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString("</a></button>")
		if err != nil {
			return err
		}
		if !templIsBuffer {
			_, err = io.Copy(w, templBuffer)
		}
		return err
	})
}

func RoomList(rooms []HRoom) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) (err error) {
		templBuffer, templIsBuffer := w.(*bytes.Buffer)
		if !templIsBuffer {
			templBuffer = templ.GetBuffer()
			defer templ.ReleaseBuffer(templBuffer)
		}
		ctx = templ.InitializeContext(ctx)
		var_56 := templ.GetChildren(ctx)
		if var_56 == nil {
			var_56 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		_, err = templBuffer.WriteString("<ul>")
		if err != nil {
			return err
		}
		for _, room := range rooms {
			_, err = templBuffer.WriteString("<li id=\"")
			if err != nil {
				return err
			}
			_, err = templBuffer.WriteString(templ.EscapeString(room.Name))
			if err != nil {
				return err
			}
			_, err = templBuffer.WriteString("\">")
			if err != nil {
				return err
			}
			err = JoinRoomEntry(room).Render(ctx, templBuffer)
			if err != nil {
				return err
			}
			_, err = templBuffer.WriteString("</li>")
			if err != nil {
				return err
			}
		}
		_, err = templBuffer.WriteString("</ul>")
		if err != nil {
			return err
		}
		if !templIsBuffer {
			_, err = io.Copy(w, templBuffer)
		}
		return err
	})
}

// TODO: use HX headers to differentiate between hx request vs normal fetch
// to decide whether or not to render entire page or just roomslist

func RoomListPage(rooms []HRoom) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) (err error) {
		templBuffer, templIsBuffer := w.(*bytes.Buffer)
		if !templIsBuffer {
			templBuffer = templ.GetBuffer()
			defer templ.ReleaseBuffer(templBuffer)
		}
		ctx = templ.InitializeContext(ctx)
		var_57 := templ.GetChildren(ctx)
		if var_57 == nil {
			var_57 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		var_58 := templ.ComponentFunc(func(ctx context.Context, w io.Writer) (err error) {
			templBuffer, templIsBuffer := w.(*bytes.Buffer)
			if !templIsBuffer {
				templBuffer = templ.GetBuffer()
				defer templ.ReleaseBuffer(templBuffer)
			}
			err = RoomList(rooms).Render(ctx, templBuffer)
			if err != nil {
				return err
			}
			if !templIsBuffer {
				_, err = io.Copy(w, templBuffer)
			}
			return err
		})
		err = Layout("Join Room").Render(templ.WithChildren(ctx, var_58), templBuffer)
		if err != nil {
			return err
		}
		if !templIsBuffer {
			_, err = io.Copy(w, templBuffer)
		}
		return err
	})
}
