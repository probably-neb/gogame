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
	if rest[0] != '/' {
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
		var_6 := templ.GetChildren(ctx)
		if var_6 == nil {
			var_6 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		var_7 := templ.ComponentFunc(func(ctx context.Context, w io.Writer) (err error) {
			templBuffer, templIsBuffer := w.(*bytes.Buffer)
			if !templIsBuffer {
				templBuffer = templ.GetBuffer()
				defer templ.ReleaseBuffer(templBuffer)
			}
			_, err = templBuffer.WriteString("<div hx-boost=\"true\"><button id=\"create-room\" class=\"rounded-md border-2 border-black\"><a href=\"/rooms/create\">")
			if err != nil {
				return err
			}
			var_8 := `Create Room`
			_, err = templBuffer.WriteString(var_8)
			if err != nil {
				return err
			}
			_, err = templBuffer.WriteString("</a></button><button id=\"join-room\" class=\"rounded-md border-2 border-black\"><a href=\"/rooms\" hx-boost=\"true\">")
			if err != nil {
				return err
			}
			var_9 := `Join Room`
			_, err = templBuffer.WriteString(var_9)
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
		err = Layout("GOGAME!").Render(templ.WithChildren(ctx, var_7), templBuffer)
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
		var_10 := templ.GetChildren(ctx)
		if var_10 == nil {
			var_10 = templ.NopComponent
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
		err = var_10.Render(ctx, templBuffer)
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
		var_11 := templ.GetChildren(ctx)
		if var_11 == nil {
			var_11 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		var_12 := templ.ComponentFunc(func(ctx context.Context, w io.Writer) (err error) {
			templBuffer, templIsBuffer := w.(*bytes.Buffer)
			if !templIsBuffer {
				templBuffer = templ.GetBuffer()
				defer templ.ReleaseBuffer(templBuffer)
			}
			_, err = templBuffer.WriteString("<form class=\"text-small text-gray-800\" hx-post=\"/rooms/create\" hx-boost=\"true\" hx-target=\"body\"><label for=\"room-name\">")
			if err != nil {
				return err
			}
			var_13 := `Room Name:`
			_, err = templBuffer.WriteString(var_13)
			if err != nil {
				return err
			}
			_, err = templBuffer.WriteString("</label><input type=\"text\" name=\"room-name\" id=\"room-name\" placeholder=\"Room Name\" class=\"rounded-md border-2 border-black\"><label for=\"display-name\">")
			if err != nil {
				return err
			}
			var_14 := `Display Name:`
			_, err = templBuffer.WriteString(var_14)
			if err != nil {
				return err
			}
			_, err = templBuffer.WriteString("</label><input type=\"text\" name=\"display-name\" id=\"display-name\" placeholder=\"Display Name\" class=\"rounded-md border-2 border-black\"><button id=\"create\" type=\"submit\" class=\"rounded-md border-2 border-black\">")
			if err != nil {
				return err
			}
			var_15 := `Create`
			_, err = templBuffer.WriteString(var_15)
			if err != nil {
				return err
			}
			_, err = templBuffer.WriteString("</button><button type=\"button\" class=\"mt-3 inline-flex w-full justify-center rounded-md bg-white px-3 py-2 text-sm font-semibold text-gray-900 shadow-sm ring-1 ring-inset ring-gray-300 hover:bg-gray-50 sm:mt-0 sm:w-auto\" hx-get=\"/ok\" hx-swap=\"delete\" hx-target=\"#create-room-modal\">")
			if err != nil {
				return err
			}
			var_16 := `Cancel`
			_, err = templBuffer.WriteString(var_16)
			if err != nil {
				return err
			}
			_, err = templBuffer.WriteString("</button></form>")
			if err != nil {
				return err
			}
			if !templIsBuffer {
				_, err = io.Copy(w, templBuffer)
			}
			return err
		})
		err = Modal("create-room-modal").Render(templ.WithChildren(ctx, var_12), templBuffer)
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
		var_17 := templ.GetChildren(ctx)
		if var_17 == nil {
			var_17 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		var_18 := templ.ComponentFunc(func(ctx context.Context, w io.Writer) (err error) {
			templBuffer, templIsBuffer := w.(*bytes.Buffer)
			if !templIsBuffer {
				templBuffer = templ.GetBuffer()
				defer templ.ReleaseBuffer(templBuffer)
			}
			_, err = templBuffer.WriteString("<form class=\"text-small text-gray-800\" ws-send><label for=\"display-name\">")
			if err != nil {
				return err
			}
			var_19 := `Display Name:`
			_, err = templBuffer.WriteString(var_19)
			if err != nil {
				return err
			}
			_, err = templBuffer.WriteString("</label><input type=\"text\" name=\"display-name\" id=\"display-name\" placeholder=\"Display Name\" class=\"rounded-md border-2 border-black\"><button id=\"create\" type=\"submit\" class=\"rounded-md border-2 border-black\">")
			if err != nil {
				return err
			}
			var_20 := `Join`
			_, err = templBuffer.WriteString(var_20)
			if err != nil {
				return err
			}
			_, err = templBuffer.WriteString("</button><button type=\"button\" class=\"mt-3 inline-flex w-full justify-center rounded-md bg-white px-3 py-2 text-sm font-semibold text-gray-900 shadow-sm ring-1 ring-inset ring-gray-300 hover:bg-gray-50 sm:mt-0 sm:w-auto\"><a href=\"/rooms\" hx-boost=\"true\">")
			if err != nil {
				return err
			}
			var_21 := `Cancel`
			_, err = templBuffer.WriteString(var_21)
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
		err = Modal("join-room-modal").Render(templ.WithChildren(ctx, var_18), templBuffer)
		if err != nil {
			return err
		}
		if !templIsBuffer {
			_, err = io.Copy(w, templBuffer)
		}
		return err
	})
}

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
		var_22 := templ.GetChildren(ctx)
		if var_22 == nil {
			var_22 = templ.NopComponent
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
		var var_23 string = name
		_, err = templBuffer.WriteString(templ.EscapeString(var_23))
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

func GuestList(guests []string, append bool) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) (err error) {
		templBuffer, templIsBuffer := w.(*bytes.Buffer)
		if !templIsBuffer {
			templBuffer = templ.GetBuffer()
			defer templ.ReleaseBuffer(templBuffer)
		}
		ctx = templ.InitializeContext(ctx)
		var_24 := templ.GetChildren(ctx)
		if var_24 == nil {
			var_24 = templ.NopComponent
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

func RoomWS(roomid string, kind string) templ.Component {
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
		_, err = templBuffer.WriteString("<div id=\"ws-connection\" hx-ext=\"ws\" ws-connect=\"")
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
		err = var_25.Render(ctx, templBuffer)
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

func RoomPageBody(room HRoom) templ.Component {
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
		_, err = templBuffer.WriteString("<div id=\"room-page-body\"><h3>")
		if err != nil {
			return err
		}
		var var_27 string = room.Name
		_, err = templBuffer.WriteString(templ.EscapeString(var_27))
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString("</h3><h2>")
		if err != nil {
			return err
		}
		var_28 := `Host:`
		_, err = templBuffer.WriteString(var_28)
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString("</h2><p id=\"host\">")
		if err != nil {
			return err
		}
		var var_29 string = room.Host
		_, err = templBuffer.WriteString(templ.EscapeString(var_29))
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString("</p><h2>")
		if err != nil {
			return err
		}
		var_30 := `Guests:`
		_, err = templBuffer.WriteString(var_30)
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

func RoomPage(room HRoom, kind string) templ.Component {
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
		var_32 := templ.ComponentFunc(func(ctx context.Context, w io.Writer) (err error) {
			templBuffer, templIsBuffer := w.(*bytes.Buffer)
			if !templIsBuffer {
				templBuffer = templ.GetBuffer()
				defer templ.ReleaseBuffer(templBuffer)
			}
			var_33 := templ.ComponentFunc(func(ctx context.Context, w io.Writer) (err error) {
				templBuffer, templIsBuffer := w.(*bytes.Buffer)
				if !templIsBuffer {
					templBuffer = templ.GetBuffer()
					defer templ.ReleaseBuffer(templBuffer)
				}
				err = RoomPageBody(room).Render(ctx, templBuffer)
				if err != nil {
					return err
				}
				if !templIsBuffer {
					_, err = io.Copy(w, templBuffer)
				}
				return err
			})
			err = RoomWS(room.Name, kind).Render(templ.WithChildren(ctx, var_33), templBuffer)
			if err != nil {
				return err
			}
			if !templIsBuffer {
				_, err = io.Copy(w, templBuffer)
			}
			return err
		})
		err = Layout(room.Name).Render(templ.WithChildren(ctx, var_32), templBuffer)
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
		var_34 := templ.GetChildren(ctx)
		if var_34 == nil {
			var_34 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		var_35 := templ.ComponentFunc(func(ctx context.Context, w io.Writer) (err error) {
			templBuffer, templIsBuffer := w.(*bytes.Buffer)
			if !templIsBuffer {
				templBuffer = templ.GetBuffer()
				defer templ.ReleaseBuffer(templBuffer)
			}
			var_36 := templ.ComponentFunc(func(ctx context.Context, w io.Writer) (err error) {
				templBuffer, templIsBuffer := w.(*bytes.Buffer)
				if !templIsBuffer {
					templBuffer = templ.GetBuffer()
					defer templ.ReleaseBuffer(templBuffer)
				}
				err = JoinRoomModal().Render(ctx, templBuffer)
				if err != nil {
					return err
				}
				_, err = templBuffer.WriteString(" ")
				if err != nil {
					return err
				}
				err = RoomPageBody(room).Render(ctx, templBuffer)
				if err != nil {
					return err
				}
				if !templIsBuffer {
					_, err = io.Copy(w, templBuffer)
				}
				return err
			})
			err = RoomWS(room.Name, "guest").Render(templ.WithChildren(ctx, var_36), templBuffer)
			if err != nil {
				return err
			}
			if !templIsBuffer {
				_, err = io.Copy(w, templBuffer)
			}
			return err
		})
		err = Layout(room.Name).Render(templ.WithChildren(ctx, var_35), templBuffer)
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
		var_37 := templ.GetChildren(ctx)
		if var_37 == nil {
			var_37 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		var_38 := templ.ComponentFunc(func(ctx context.Context, w io.Writer) (err error) {
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
			err = RoomPageBody(HRoom{}).Render(ctx, templBuffer)
			if err != nil {
				return err
			}
			if !templIsBuffer {
				_, err = io.Copy(w, templBuffer)
			}
			return err
		})
		err = Layout("Create Room").Render(templ.WithChildren(ctx, var_38), templBuffer)
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
		var_39 := templ.GetChildren(ctx)
		if var_39 == nil {
			var_39 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		var_40 := templ.ComponentFunc(func(ctx context.Context, w io.Writer) (err error) {
			templBuffer, templIsBuffer := w.(*bytes.Buffer)
			if !templIsBuffer {
				templBuffer = templ.GetBuffer()
				defer templ.ReleaseBuffer(templBuffer)
			}
			var_41 := templ.ComponentFunc(func(ctx context.Context, w io.Writer) (err error) {
				templBuffer, templIsBuffer := w.(*bytes.Buffer)
				if !templIsBuffer {
					templBuffer = templ.GetBuffer()
					defer templ.ReleaseBuffer(templBuffer)
				}
				err = RoomPageBody(room).Render(ctx, templBuffer)
				if err != nil {
					return err
				}
				if !templIsBuffer {
					_, err = io.Copy(w, templBuffer)
				}
				return err
			})
			err = RoomWS(room.Name, "host").Render(templ.WithChildren(ctx, var_41), templBuffer)
			if err != nil {
				return err
			}
			if !templIsBuffer {
				_, err = io.Copy(w, templBuffer)
			}
			return err
		})
		err = Layout(room.Name).Render(templ.WithChildren(ctx, var_40), templBuffer)
		if err != nil {
			return err
		}
		if !templIsBuffer {
			_, err = io.Copy(w, templBuffer)
		}
		return err
	})
}

func JoinRoomRedirect(roomid string) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) (err error) {
		templBuffer, templIsBuffer := w.(*bytes.Buffer)
		if !templIsBuffer {
			templBuffer = templ.GetBuffer()
			defer templ.ReleaseBuffer(templBuffer)
		}
		ctx = templ.InitializeContext(ctx)
		var_42 := templ.GetChildren(ctx)
		if var_42 == nil {
			var_42 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		_, err = templBuffer.WriteString("<div id=\"url-push\" hx-get=\"/ok\" hx-push-url=\"")
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString(templ.EscapeString(RoomURL(roomid, "")))
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString("\" hx-trigger=\"load\"></div>")
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
		var_43 := templ.GetChildren(ctx)
		if var_43 == nil {
			var_43 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		_, err = templBuffer.WriteString("<h3>")
		if err != nil {
			return err
		}
		var var_44 string = room.Name
		_, err = templBuffer.WriteString(templ.EscapeString(var_44))
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString("</h3><h2>")
		if err != nil {
			return err
		}
		var_45 := `Host:`
		_, err = templBuffer.WriteString(var_45)
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString("</h2><p>")
		if err != nil {
			return err
		}
		var var_46 string = room.Host
		_, err = templBuffer.WriteString(templ.EscapeString(var_46))
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString("</p><h2>")
		if err != nil {
			return err
		}
		var_47 := `Guests: `
		_, err = templBuffer.WriteString(var_47)
		if err != nil {
			return err
		}
		var var_48 string = fmt.Sprint(len(room.Guests))
		_, err = templBuffer.WriteString(templ.EscapeString(var_48))
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString("</h2><button id=\"join-room\" class=\"rounded-md border-2 border-black\"><a href=\"")
		if err != nil {
			return err
		}
		var var_49 templ.SafeURL = templ.SafeURL(RoomURL(room.Name, "/join"))
		_, err = templBuffer.WriteString(templ.EscapeString(string(var_49)))
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString("\" hx-boost=\"true\">")
		if err != nil {
			return err
		}
		var_50 := `Join Room`
		_, err = templBuffer.WriteString(var_50)
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
		var_51 := templ.GetChildren(ctx)
		if var_51 == nil {
			var_51 = templ.NopComponent
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
		var_52 := templ.GetChildren(ctx)
		if var_52 == nil {
			var_52 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		var_53 := templ.ComponentFunc(func(ctx context.Context, w io.Writer) (err error) {
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
		err = Layout("Join Room").Render(templ.WithChildren(ctx, var_53), templBuffer)
		if err != nil {
			return err
		}
		if !templIsBuffer {
			_, err = io.Copy(w, templBuffer)
		}
		return err
	})
}
