// Code generated by templ@v0.2.316 DO NOT EDIT.

package main

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import "context"
import "io"
import "bytes"

import (
	"fmt"
	"os"
)

var isProd = os.Getenv("ENV") == "PRODUCTION"

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
		_, err = templBuffer.WriteString("<!doctype html><html lang=\"en\" class=\"dark\" style=\"color-scheme: dark\"><head><meta charset=\"UTF-8\"><title>")
		if err != nil {
			return err
		}
		var var_2 string = title
		_, err = templBuffer.WriteString(templ.EscapeString(var_2))
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString("</title><link rel=\"icon\" type=\"image/svg+xml\" href=\"/dist/favicon.svg\"><meta name=\"viewport\" content=\"width=device-width, initial-scale=1\"><meta name=\"robots\" content=\"index, follow\"><meta name=\"revisit-after\" content=\"7 days\"><meta name=\"language\" content=\"English\"><script")
		if err != nil {
			return err
		}
		if isProd {
			_, err = templBuffer.WriteString(" src=\"https://unpkg.com/htmx.org@1.9.5/dist/htmx.min.js\"")
			if err != nil {
				return err
			}
		} else {
			_, err = templBuffer.WriteString(" src=\"https://unpkg.com/htmx.org@1.9.5/dist/htmx.js\"")
			if err != nil {
				return err
			}
		}
		_, err = templBuffer.WriteString(">")
		if err != nil {
			return err
		}
		var_3 := ``
		_, err = templBuffer.WriteString(var_3)
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString("</script><script src=\"https://unpkg.com/htmx.org@1.9.5/dist/ext/ws.js\">")
		if err != nil {
			return err
		}
		var_4 := ``
		_, err = templBuffer.WriteString(var_4)
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString("</script><link rel=\"stylesheet\" href=\"/assets/tailwind.css\"><script>")
		if err != nil {
			return err
		}
		var_5 := `
                htmx.on("htmx:wsConfigSend", function(e) {
                    console.dir(e, {depth: null})
                    let ty = e.detail.elt.dataset.type
                    let grp = e.detail.elt.dataset.group
                    e.detail.parameters = {type: ty, group: grp, data: e.detail.parameters}
                })
                htmx.on("htmx:configRequest", function(e) {
                    let sessionInput = htmx.find('#session-id')
                    if (!sessionInput) {
                        return
                    }
                    let sessionId = sessionInput.value
                    e.detail.headers["Session-Id"] = sessionId
                })
            `
		_, err = templBuffer.WriteString(var_5)
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString("</script></head><body id=\"body\" class=\"bg-background\">")
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

func Session(sessionId string) templ.Component {
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
		_, err = templBuffer.WriteString("<input type=\"hidden\" name=\"session-id\" id=\"session-id\" value=\"")
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString(templ.EscapeString(sessionId))
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString("\">")
		if err != nil {
			return err
		}
		if !templIsBuffer {
			_, err = io.Copy(w, templBuffer)
		}
		return err
	})
}

// CREDIT: shadcn-ui

func Button() templ.Component {
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
		_, err = templBuffer.WriteString("<div class=\"inline-flex items-center justify-center rounded-md text-sm font-medium ring-offset-background transition-colors focus-visible:outline-none focus-visible:ring-2 focus-visible:ring-ring focus-visible:ring-offset-2 disabled:pointer-events-none disabled:opacity-50 border-ring border-4 bg-primary-foreground text-foreground hover:bg-ring h-10 px-4 py-2\">")
		if err != nil {
			return err
		}
		err = var_7.Render(ctx, templBuffer)
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

var colors = []string{
	"bg-foreground",
	"bg-border",
	"bg-input",
	"bg-ring",
	"bg-background",
	"bg-primary",
	"bg-primary-foreground",
	"bg-secondary",
	"bg-secondary-foreground",
	"bg-destructive",
	"bg-muted",
	"bg-muted-foreground",
	"bg-accent",
	"bg-accent-foreground",
	"bg-popover",
	"bg-popover-foreground",
	"bg-card",
	"bg-card-foreground",
}

func SessionInput(displayName string) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) (err error) {
		templBuffer, templIsBuffer := w.(*bytes.Buffer)
		if !templIsBuffer {
			templBuffer = templ.GetBuffer()
			defer templ.ReleaseBuffer(templBuffer)
		}
		ctx = templ.InitializeContext(ctx)
		var_8 := templ.GetChildren(ctx)
		if var_8 == nil {
			var_8 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		_, err = templBuffer.WriteString("<form id=\"session-form\" class=\"text-small text-gray-800\" hx-post=\"/sessions\" hx-sync=\"this:abort\"><div class=\"grid gap-y-1.5 gap-x-1.5 grid-cols-3 grid-rows-1 items-center text-foreground\"><label for=\"display-name\" class=\"text-right\">")
		if err != nil {
			return err
		}
		var_9 := `Display Name:`
		_, err = templBuffer.WriteString(var_9)
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString("</label><input type=\"text\" name=\"display-name\" id=\"display-name\"")
		if err != nil {
			return err
		}
		if displayName != "" {
			_, err = templBuffer.WriteString(" value=\"")
			if err != nil {
				return err
			}
			_, err = templBuffer.WriteString(templ.EscapeString(displayName))
			if err != nil {
				return err
			}
			_, err = templBuffer.WriteString("\"")
			if err != nil {
				return err
			}
		}
		_, err = templBuffer.WriteString(" data-default=\"")
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString(templ.EscapeString(displayName))
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString("\" hx-target=\"#session-form\" hx-swap=\"outerHTML\" hx-post=\"/sessions\" required=\"true\" class=\"rounded-sm p-2 col-span-2\" placeholder=\"Display Name\" hx-on:focus=\"if(this.value === this.dataset.default) {this.value = &#39;&#39;}\" hx-on:blur=\"if(this.value === &#39;&#39;) {this.value = this.dataset.default}\"></div></form>")
		if err != nil {
			return err
		}
		if !templIsBuffer {
			_, err = io.Copy(w, templBuffer)
		}
		return err
	})
}

func LandingPage(sessionId string, displayName string) templ.Component {
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
		var_11 := templ.ComponentFunc(func(ctx context.Context, w io.Writer) (err error) {
			templBuffer, templIsBuffer := w.(*bytes.Buffer)
			if !templIsBuffer {
				templBuffer = templ.GetBuffer()
				defer templ.ReleaseBuffer(templBuffer)
			}
			_, err = templBuffer.WriteString("<div class=\"w-screen h-screen flex flex-col items-center justify-around\">")
			if err != nil {
				return err
			}
			err = Session(sessionId).Render(ctx, templBuffer)
			if err != nil {
				return err
			}
			err = SessionInput(displayName).Render(ctx, templBuffer)
			if err != nil {
				return err
			}
			_, err = templBuffer.WriteString("<div hx-boost=\"true\" class=\"flex flex-row items-center justfiy-around\">")
			if err != nil {
				return err
			}
			var_12 := templ.ComponentFunc(func(ctx context.Context, w io.Writer) (err error) {
				templBuffer, templIsBuffer := w.(*bytes.Buffer)
				if !templIsBuffer {
					templBuffer = templ.GetBuffer()
					defer templ.ReleaseBuffer(templBuffer)
				}
				_, err = templBuffer.WriteString("<button id=\"create-room\"><a href=\"/rooms/create\">")
				if err != nil {
					return err
				}
				var_13 := `Create Room`
				_, err = templBuffer.WriteString(var_13)
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
			err = Button().Render(templ.WithChildren(ctx, var_12), templBuffer)
			if err != nil {
				return err
			}
			var_14 := templ.ComponentFunc(func(ctx context.Context, w io.Writer) (err error) {
				templBuffer, templIsBuffer := w.(*bytes.Buffer)
				if !templIsBuffer {
					templBuffer = templ.GetBuffer()
					defer templ.ReleaseBuffer(templBuffer)
				}
				_, err = templBuffer.WriteString("<button id=\"join-room\"><a href=\"/rooms\" hx-push-url=\"true\">")
				if err != nil {
					return err
				}
				var_15 := `Join Room`
				_, err = templBuffer.WriteString(var_15)
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
			err = Button().Render(templ.WithChildren(ctx, var_14), templBuffer)
			if err != nil {
				return err
			}
			_, err = templBuffer.WriteString("</div>")
			if err != nil {
				return err
			}
			if !isProd {
				_, err = templBuffer.WriteString("<table><tbody>")
				if err != nil {
					return err
				}
				for _, color := range colors {
					_, err = templBuffer.WriteString("<tr><td>")
					if err != nil {
						return err
					}
					var var_16 string = color
					_, err = templBuffer.WriteString(templ.EscapeString(var_16))
					if err != nil {
						return err
					}
					_, err = templBuffer.WriteString("</td><td>")
					if err != nil {
						return err
					}
					var var_17 = []any{"w-8 h-5 " + color}
					err = templ.RenderCSSItems(ctx, templBuffer, var_17...)
					if err != nil {
						return err
					}
					_, err = templBuffer.WriteString("<div class=\"")
					if err != nil {
						return err
					}
					_, err = templBuffer.WriteString(templ.EscapeString(templ.CSSClasses(var_17).String()))
					if err != nil {
						return err
					}
					_, err = templBuffer.WriteString("\"></div></td></tr>")
					if err != nil {
						return err
					}
				}
				_, err = templBuffer.WriteString("</tbody></table>")
				if err != nil {
					return err
				}
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
		err = Layout("GOGAME!").Render(templ.WithChildren(ctx, var_11), templBuffer)
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
		var_18 := templ.GetChildren(ctx)
		if var_18 == nil {
			var_18 = templ.NopComponent
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
		err = var_18.Render(ctx, templBuffer)
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

func CreateRoomModal(sessionId string) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) (err error) {
		templBuffer, templIsBuffer := w.(*bytes.Buffer)
		if !templIsBuffer {
			templBuffer = templ.GetBuffer()
			defer templ.ReleaseBuffer(templBuffer)
		}
		ctx = templ.InitializeContext(ctx)
		var_19 := templ.GetChildren(ctx)
		if var_19 == nil {
			var_19 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		var_20 := templ.ComponentFunc(func(ctx context.Context, w io.Writer) (err error) {
			templBuffer, templIsBuffer := w.(*bytes.Buffer)
			if !templIsBuffer {
				templBuffer = templ.GetBuffer()
				defer templ.ReleaseBuffer(templBuffer)
			}
			_, err = templBuffer.WriteString("<form class=\"text-small text-gray-800\" hx-post=\"/rooms/create\" hx-boost=\"true\" hx-target=\"body\"><div class=\"grid gap-y-1.5 gap-x-1.5 grid-cols-3 grid-rows-1 items-center text-foreground\"><label for=\"room-name\" class=\"text-right\">")
			if err != nil {
				return err
			}
			var_21 := `Room Name`
			_, err = templBuffer.WriteString(var_21)
			if err != nil {
				return err
			}
			_, err = templBuffer.WriteString("</label><input type=\"text\" name=\"room-name\" id=\"room-name\" placeholder=\"Room Name\" class=\"rounded-sm p-2 col-span-2\">")
			if err != nil {
				return err
			}
			err = Session(sessionId).Render(ctx, templBuffer)
			if err != nil {
				return err
			}
			_, err = templBuffer.WriteString("</div><div class=\"h-4\"></div><div class=\"flex justify-around\">")
			if err != nil {
				return err
			}
			var_22 := templ.ComponentFunc(func(ctx context.Context, w io.Writer) (err error) {
				templBuffer, templIsBuffer := w.(*bytes.Buffer)
				if !templIsBuffer {
					templBuffer = templ.GetBuffer()
					defer templ.ReleaseBuffer(templBuffer)
				}
				_, err = templBuffer.WriteString("<button type=\"button\"><a href=\"/\" hx-boost=\"true\">")
				if err != nil {
					return err
				}
				var_23 := `Cancel`
				_, err = templBuffer.WriteString(var_23)
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
			err = Button().Render(templ.WithChildren(ctx, var_22), templBuffer)
			if err != nil {
				return err
			}
			var_24 := templ.ComponentFunc(func(ctx context.Context, w io.Writer) (err error) {
				templBuffer, templIsBuffer := w.(*bytes.Buffer)
				if !templIsBuffer {
					templBuffer = templ.GetBuffer()
					defer templ.ReleaseBuffer(templBuffer)
				}
				_, err = templBuffer.WriteString("<button id=\"create\" type=\"submit\">")
				if err != nil {
					return err
				}
				var_25 := `Create`
				_, err = templBuffer.WriteString(var_25)
				if err != nil {
					return err
				}
				_, err = templBuffer.WriteString("</button>")
				if err != nil {
					return err
				}
				if !templIsBuffer {
					_, err = io.Copy(w, templBuffer)
				}
				return err
			})
			err = Button().Render(templ.WithChildren(ctx, var_24), templBuffer)
			if err != nil {
				return err
			}
			_, err = templBuffer.WriteString("</div></form>")
			if err != nil {
				return err
			}
			if !templIsBuffer {
				_, err = io.Copy(w, templBuffer)
			}
			return err
		})
		err = Modal("create-room-modal").Render(templ.WithChildren(ctx, var_20), templBuffer)
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
		var_26 := templ.GetChildren(ctx)
		if var_26 == nil {
			var_26 = templ.NopComponent
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
		var var_27 string = name
		_, err = templBuffer.WriteString(templ.EscapeString(var_27))
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
		var_28 := templ.GetChildren(ctx)
		if var_28 == nil {
			var_28 = templ.NopComponent
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
		var_29 := templ.GetChildren(ctx)
		if var_29 == nil {
			var_29 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		_, err = templBuffer.WriteString("<div><h3>")
		if err != nil {
			return err
		}
		var_30 := `Games`
		_, err = templBuffer.WriteString(var_30)
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString("</h3><ul><li>")
		if err != nil {
			return err
		}
		var_31 := `Tic-Tac-Toe`
		_, err = templBuffer.WriteString(var_31)
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString(" ")
		if err != nil {
			return err
		}
		if kind == "host" {
			var_32 := templ.ComponentFunc(func(ctx context.Context, w io.Writer) (err error) {
				templBuffer, templIsBuffer := w.(*bytes.Buffer)
				if !templIsBuffer {
					templBuffer = templ.GetBuffer()
					defer templ.ReleaseBuffer(templBuffer)
				}
				_, err = templBuffer.WriteString("<button id=\"start-tic-tac-toe\" name=\"play\" value=\"tictactoe\" ws-send data-group=\"room\" data-type=\"play\">")
				if err != nil {
					return err
				}
				var_33 := `Play!`
				_, err = templBuffer.WriteString(var_33)
				if err != nil {
					return err
				}
				_, err = templBuffer.WriteString("</button>")
				if err != nil {
					return err
				}
				if !templIsBuffer {
					_, err = io.Copy(w, templBuffer)
				}
				return err
			})
			err = Button().Render(templ.WithChildren(ctx, var_32), templBuffer)
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

func RoomWS(roomid string, sessionId string) templ.Component {
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
		_, err = templBuffer.WriteString("<div id=\"ws-connection\" hx-preserve=\"true\" hx-ext=\"ws\" ws-connect=\"")
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString(templ.EscapeString(RoomURL(roomid, sessionId+"/ws")))
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString("\">")
		if err != nil {
			return err
		}
		err = var_34.Render(ctx, templBuffer)
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
		var_35 := templ.GetChildren(ctx)
		if var_35 == nil {
			var_35 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		_, err = templBuffer.WriteString("<div id=\"room-page-body\"><h3>")
		if err != nil {
			return err
		}
		var var_36 string = room.Name
		_, err = templBuffer.WriteString(templ.EscapeString(var_36))
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString("</h3><h2>")
		if err != nil {
			return err
		}
		var_37 := `Host:`
		_, err = templBuffer.WriteString(var_37)
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString("</h2><p id=\"host\">")
		if err != nil {
			return err
		}
		var var_38 string = room.Host
		_, err = templBuffer.WriteString(templ.EscapeString(var_38))
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString("</p><h2>")
		if err != nil {
			return err
		}
		var_39 := `Guests:`
		_, err = templBuffer.WriteString(var_39)
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

func CreateRoomPage(sessionId string) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) (err error) {
		templBuffer, templIsBuffer := w.(*bytes.Buffer)
		if !templIsBuffer {
			templBuffer = templ.GetBuffer()
			defer templ.ReleaseBuffer(templBuffer)
		}
		ctx = templ.InitializeContext(ctx)
		var_40 := templ.GetChildren(ctx)
		if var_40 == nil {
			var_40 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		var_41 := templ.ComponentFunc(func(ctx context.Context, w io.Writer) (err error) {
			templBuffer, templIsBuffer := w.(*bytes.Buffer)
			if !templIsBuffer {
				templBuffer = templ.GetBuffer()
				defer templ.ReleaseBuffer(templBuffer)
			}
			err = CreateRoomModal(sessionId).Render(ctx, templBuffer)
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
		err = Layout("Create Room").Render(templ.WithChildren(ctx, var_41), templBuffer)
		if err != nil {
			return err
		}
		if !templIsBuffer {
			_, err = io.Copy(w, templBuffer)
		}
		return err
	})
}

func GuestRoomPage(room HRoom, sessionId string) templ.Component {
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
		var_43 := templ.ComponentFunc(func(ctx context.Context, w io.Writer) (err error) {
			templBuffer, templIsBuffer := w.(*bytes.Buffer)
			if !templIsBuffer {
				templBuffer = templ.GetBuffer()
				defer templ.ReleaseBuffer(templBuffer)
			}
			err = Session(sessionId).Render(ctx, templBuffer)
			if err != nil {
				return err
			}
			_, err = templBuffer.WriteString(" ")
			if err != nil {
				return err
			}
			var_44 := templ.ComponentFunc(func(ctx context.Context, w io.Writer) (err error) {
				templBuffer, templIsBuffer := w.(*bytes.Buffer)
				if !templIsBuffer {
					templBuffer = templ.GetBuffer()
					defer templ.ReleaseBuffer(templBuffer)
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
			err = RoomWS(room.Name, sessionId).Render(templ.WithChildren(ctx, var_44), templBuffer)
			if err != nil {
				return err
			}
			if !templIsBuffer {
				_, err = io.Copy(w, templBuffer)
			}
			return err
		})
		err = Layout(room.Name).Render(templ.WithChildren(ctx, var_43), templBuffer)
		if err != nil {
			return err
		}
		if !templIsBuffer {
			_, err = io.Copy(w, templBuffer)
		}
		return err
	})
}

// TODO: further divide up session input and have join button here

func JoinRoomModal(displayName string) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) (err error) {
		templBuffer, templIsBuffer := w.(*bytes.Buffer)
		if !templIsBuffer {
			templBuffer = templ.GetBuffer()
			defer templ.ReleaseBuffer(templBuffer)
		}
		ctx = templ.InitializeContext(ctx)
		var_45 := templ.GetChildren(ctx)
		if var_45 == nil {
			var_45 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		var_46 := templ.ComponentFunc(func(ctx context.Context, w io.Writer) (err error) {
			templBuffer, templIsBuffer := w.(*bytes.Buffer)
			if !templIsBuffer {
				templBuffer = templ.GetBuffer()
				defer templ.ReleaseBuffer(templBuffer)
			}
			err = SessionInput(displayName).Render(ctx, templBuffer)
			if err != nil {
				return err
			}
			_, err = templBuffer.WriteString(" ")
			if err != nil {
				return err
			}
			var_47 := templ.ComponentFunc(func(ctx context.Context, w io.Writer) (err error) {
				templBuffer, templIsBuffer := w.(*bytes.Buffer)
				if !templIsBuffer {
					templBuffer = templ.GetBuffer()
					defer templ.ReleaseBuffer(templBuffer)
				}
				_, err = templBuffer.WriteString("<button hx-post=\"/sessions\" hx-include=\"#session-form\">")
				if err != nil {
					return err
				}
				var_48 := `Join`
				_, err = templBuffer.WriteString(var_48)
				if err != nil {
					return err
				}
				_, err = templBuffer.WriteString("</button>")
				if err != nil {
					return err
				}
				if !templIsBuffer {
					_, err = io.Copy(w, templBuffer)
				}
				return err
			})
			err = Button().Render(templ.WithChildren(ctx, var_47), templBuffer)
			if err != nil {
				return err
			}
			if !templIsBuffer {
				_, err = io.Copy(w, templBuffer)
			}
			return err
		})
		err = Modal("join-room-modal").Render(templ.WithChildren(ctx, var_46), templBuffer)
		if err != nil {
			return err
		}
		if !templIsBuffer {
			_, err = io.Copy(w, templBuffer)
		}
		return err
	})
}

func JoinRoomPage(room HRoom, sessionId string, displayName string) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) (err error) {
		templBuffer, templIsBuffer := w.(*bytes.Buffer)
		if !templIsBuffer {
			templBuffer = templ.GetBuffer()
			defer templ.ReleaseBuffer(templBuffer)
		}
		ctx = templ.InitializeContext(ctx)
		var_49 := templ.GetChildren(ctx)
		if var_49 == nil {
			var_49 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		var_50 := templ.ComponentFunc(func(ctx context.Context, w io.Writer) (err error) {
			templBuffer, templIsBuffer := w.(*bytes.Buffer)
			if !templIsBuffer {
				templBuffer = templ.GetBuffer()
				defer templ.ReleaseBuffer(templBuffer)
			}
			err = Session(sessionId).Render(ctx, templBuffer)
			if err != nil {
				return err
			}
			_, err = templBuffer.WriteString(" ")
			if err != nil {
				return err
			}
			err = JoinRoomModal(displayName).Render(ctx, templBuffer)
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
		err = Layout(room.Name).Render(templ.WithChildren(ctx, var_50), templBuffer)
		if err != nil {
			return err
		}
		if !templIsBuffer {
			_, err = io.Copy(w, templBuffer)
		}
		return err
	})
}

func HostRoomPage(room HRoom, sessionId string) templ.Component {
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
		var_52 := templ.ComponentFunc(func(ctx context.Context, w io.Writer) (err error) {
			templBuffer, templIsBuffer := w.(*bytes.Buffer)
			if !templIsBuffer {
				templBuffer = templ.GetBuffer()
				defer templ.ReleaseBuffer(templBuffer)
			}
			err = Session(sessionId).Render(ctx, templBuffer)
			if err != nil {
				return err
			}
			_, err = templBuffer.WriteString(" ")
			if err != nil {
				return err
			}
			var_53 := templ.ComponentFunc(func(ctx context.Context, w io.Writer) (err error) {
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
			err = RoomWS(room.Name, sessionId).Render(templ.WithChildren(ctx, var_53), templBuffer)
			if err != nil {
				return err
			}
			if !templIsBuffer {
				_, err = io.Copy(w, templBuffer)
			}
			return err
		})
		err = Layout(room.Name).Render(templ.WithChildren(ctx, var_52), templBuffer)
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
		var_54 := templ.GetChildren(ctx)
		if var_54 == nil {
			var_54 = templ.NopComponent
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
		var_55 := templ.GetChildren(ctx)
		if var_55 == nil {
			var_55 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		_, err = templBuffer.WriteString("<h3>")
		if err != nil {
			return err
		}
		var var_56 string = room.Name
		_, err = templBuffer.WriteString(templ.EscapeString(var_56))
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString("</h3><h2>")
		if err != nil {
			return err
		}
		var_57 := `Host:`
		_, err = templBuffer.WriteString(var_57)
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString("</h2><p>")
		if err != nil {
			return err
		}
		var var_58 string = room.Host
		_, err = templBuffer.WriteString(templ.EscapeString(var_58))
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString("</p><h2>")
		if err != nil {
			return err
		}
		var_59 := `Guests: `
		_, err = templBuffer.WriteString(var_59)
		if err != nil {
			return err
		}
		var var_60 string = fmt.Sprint(len(room.Guests))
		_, err = templBuffer.WriteString(templ.EscapeString(var_60))
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString("</h2>")
		if err != nil {
			return err
		}
		var_61 := templ.ComponentFunc(func(ctx context.Context, w io.Writer) (err error) {
			templBuffer, templIsBuffer := w.(*bytes.Buffer)
			if !templIsBuffer {
				templBuffer = templ.GetBuffer()
				defer templ.ReleaseBuffer(templBuffer)
			}
			_, err = templBuffer.WriteString("<button id=\"join-room\" class=\"rounded-md border-2 border-black\"><a href=\"")
			if err != nil {
				return err
			}
			var var_62 templ.SafeURL = templ.SafeURL("/rooms/" + room.Name)
			_, err = templBuffer.WriteString(templ.EscapeString(string(var_62)))
			if err != nil {
				return err
			}
			_, err = templBuffer.WriteString("\" hx-boost=\"true\">")
			if err != nil {
				return err
			}
			var_63 := `Join Room`
			_, err = templBuffer.WriteString(var_63)
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
		err = Button().Render(templ.WithChildren(ctx, var_61), templBuffer)
		if err != nil {
			return err
		}
		if !templIsBuffer {
			_, err = io.Copy(w, templBuffer)
		}
		return err
	})
}

func RoomList(rooms []HRoom, sessionId string) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) (err error) {
		templBuffer, templIsBuffer := w.(*bytes.Buffer)
		if !templIsBuffer {
			templBuffer = templ.GetBuffer()
			defer templ.ReleaseBuffer(templBuffer)
		}
		ctx = templ.InitializeContext(ctx)
		var_64 := templ.GetChildren(ctx)
		if var_64 == nil {
			var_64 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		err = Session(sessionId).Render(ctx, templBuffer)
		if err != nil {
			return err
		}
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

func RoomListPage(rooms []HRoom, sessionId string) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) (err error) {
		templBuffer, templIsBuffer := w.(*bytes.Buffer)
		if !templIsBuffer {
			templBuffer = templ.GetBuffer()
			defer templ.ReleaseBuffer(templBuffer)
		}
		ctx = templ.InitializeContext(ctx)
		var_65 := templ.GetChildren(ctx)
		if var_65 == nil {
			var_65 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		var_66 := templ.ComponentFunc(func(ctx context.Context, w io.Writer) (err error) {
			templBuffer, templIsBuffer := w.(*bytes.Buffer)
			if !templIsBuffer {
				templBuffer = templ.GetBuffer()
				defer templ.ReleaseBuffer(templBuffer)
			}
			err = RoomList(rooms, sessionId).Render(ctx, templBuffer)
			if err != nil {
				return err
			}
			if !templIsBuffer {
				_, err = io.Copy(w, templBuffer)
			}
			return err
		})
		err = Layout("Join Room").Render(templ.WithChildren(ctx, var_66), templBuffer)
		if err != nil {
			return err
		}
		if !templIsBuffer {
			_, err = io.Copy(w, templBuffer)
		}
		return err
	})
}
