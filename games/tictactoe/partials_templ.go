// Code generated by templ@v0.2.316 DO NOT EDIT.

package tictactoe

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import "context"
import "io"
import "bytes"

func Box(id string, symbol rune) templ.Component {
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
		_, err = templBuffer.WriteString("<div ws-send name=\"cell\" id=\"")
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString(templ.EscapeString(id))
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString("\" class=\"bg-gray-300 h-16 flex items-center justify-center text-4xl font-bold cursor-pointer\">")
		if err != nil {
			return err
		}
		var var_2 string = string(symbol)
		_, err = templBuffer.WriteString(templ.EscapeString(var_2))
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

func Board() templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) (err error) {
		templBuffer, templIsBuffer := w.(*bytes.Buffer)
		if !templIsBuffer {
			templBuffer = templ.GetBuffer()
			defer templ.ReleaseBuffer(templBuffer)
		}
		ctx = templ.InitializeContext(ctx)
		var_3 := templ.GetChildren(ctx)
		if var_3 == nil {
			var_3 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		_, err = templBuffer.WriteString("<div class=\"grid grid-cols-3 gap-2\">")
		if err != nil {
			return err
		}
		err = Box("cell-1", 0).Render(ctx, templBuffer)
		if err != nil {
			return err
		}
		err = Box("cell-2", 0).Render(ctx, templBuffer)
		if err != nil {
			return err
		}
		err = Box("cell-3", 0).Render(ctx, templBuffer)
		if err != nil {
			return err
		}
		err = Box("cell-4", 0).Render(ctx, templBuffer)
		if err != nil {
			return err
		}
		err = Box("cell-5", 0).Render(ctx, templBuffer)
		if err != nil {
			return err
		}
		err = Box("cell-6", 0).Render(ctx, templBuffer)
		if err != nil {
			return err
		}
		err = Box("cell-7", 0).Render(ctx, templBuffer)
		if err != nil {
			return err
		}
		err = Box("cell-8", 0).Render(ctx, templBuffer)
		if err != nil {
			return err
		}
		err = Box("cell-9", 0).Render(ctx, templBuffer)
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

func Page() templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) (err error) {
		templBuffer, templIsBuffer := w.(*bytes.Buffer)
		if !templIsBuffer {
			templBuffer = templ.GetBuffer()
			defer templ.ReleaseBuffer(templBuffer)
		}
		ctx = templ.InitializeContext(ctx)
		var_4 := templ.GetChildren(ctx)
		if var_4 == nil {
			var_4 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		_, err = templBuffer.WriteString("<html lang=\"en\"><head><meta charset=\"UTF-8\"><title>")
		if err != nil {
			return err
		}
		var_5 := `GO GAMES!`
		_, err = templBuffer.WriteString(var_5)
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString("</title><link rel=\"icon\" type=\"image/svg+xml\" href=\"/dist/favicon.svg\"><meta name=\"viewport\" content=\"width=device-width, initial-scale=1\"><meta name=\"robots\" content=\"index, follow\"><meta name=\"revisit-after\" content=\"7 days\"><meta name=\"language\" content=\"English\"><script src=\"https://unpkg.com/htmx.org@1.9.5\" integrity=\"sha384-xcuj3WpfgjlKF+FXhSQFQ0ZNr39ln+hwjN3npfM9VBnUskLolQAcN80McRIVOPuO\" crossorigin=\"anonymous\">")
		if err != nil {
			return err
		}
		var_6 := ``
		_, err = templBuffer.WriteString(var_6)
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString("</script><script src=\"https://cdn.tailwindcss.com\">")
		if err != nil {
			return err
		}
		var_7 := ``
		_, err = templBuffer.WriteString(var_7)
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString("</script><script src=\"https://unpkg.com/htmx.org/dist/ext/ws.js\">")
		if err != nil {
			return err
		}
		var_8 := ``
		_, err = templBuffer.WriteString(var_8)
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString("</script></head><body><div class=\"bg-white shadow-lg rounded-lg p-4\" hx-ext=\"ws\" ws-connect=\"/games/tictactoe/ws\"><h1 class=\"text-2xl font-semibold mb-4\">")
		if err != nil {
			return err
		}
		var_9 := `Tic-Tac-Toe`
		_, err = templBuffer.WriteString(var_9)
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString("</h1>")
		if err != nil {
			return err
		}
		err = Board().Render(ctx, templBuffer)
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString("<button class=\"mt-4 bg-blue-500 hover:bg-blue-600 text-white px-4 py-2 rounded-md\">")
		if err != nil {
			return err
		}
		var_10 := `Reset`
		_, err = templBuffer.WriteString(var_10)
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString("</button></div></body></html>")
		if err != nil {
			return err
		}
		if !templIsBuffer {
			_, err = io.Copy(w, templBuffer)
		}
		return err
	})
}
