// Code generated by templ@v0.2.316 DO NOT EDIT.

package tictactoe

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import "context"
import "io"
import "bytes"

func Box(id string, symbol *rune) templ.Component {
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
		var var_2 = []any{"bg-muted", "flex", "items-center", "justify-center", "text-8xl", "font-bold", templ.KV("cursor-pointer", symbol == nil), templ.KV("cursor-default", symbol != nil), "w-full", "aspect-square", "text-primary"}
		err = templ.RenderCSSItems(ctx, templBuffer, var_2...)
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString("<div")
		if err != nil {
			return err
		}
		if symbol == nil {
			_, err = templBuffer.WriteString(" ws-send")
			if err != nil {
				return err
			}
		}
		_, err = templBuffer.WriteString(" name=\"cell\" id=\"")
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString(templ.EscapeString(id))
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString("\" class=\"")
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString(templ.EscapeString(templ.CSSClasses(var_2).String()))
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString("\">")
		if err != nil {
			return err
		}
		if symbol != nil {
			var var_3 string = string(*symbol)
			_, err = templBuffer.WriteString(templ.EscapeString(var_3))
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
}

func Board() templ.Component {
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
		_, err = templBuffer.WriteString("<div id=\"board\" class=\"grid grid-cols-3 gap-2 aspect-square w-3/5\">")
		if err != nil {
			return err
		}
		err = Box("c0", nil).Render(ctx, templBuffer)
		if err != nil {
			return err
		}
		err = Box("c1", nil).Render(ctx, templBuffer)
		if err != nil {
			return err
		}
		err = Box("c2", nil).Render(ctx, templBuffer)
		if err != nil {
			return err
		}
		err = Box("c3", nil).Render(ctx, templBuffer)
		if err != nil {
			return err
		}
		err = Box("c4", nil).Render(ctx, templBuffer)
		if err != nil {
			return err
		}
		err = Box("c5", nil).Render(ctx, templBuffer)
		if err != nil {
			return err
		}
		err = Box("c6", nil).Render(ctx, templBuffer)
		if err != nil {
			return err
		}
		err = Box("c7", nil).Render(ctx, templBuffer)
		if err != nil {
			return err
		}
		err = Box("c8", nil).Render(ctx, templBuffer)
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

func Game() templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) (err error) {
		templBuffer, templIsBuffer := w.(*bytes.Buffer)
		if !templIsBuffer {
			templBuffer = templ.GetBuffer()
			defer templ.ReleaseBuffer(templBuffer)
		}
		ctx = templ.InitializeContext(ctx)
		var_5 := templ.GetChildren(ctx)
		if var_5 == nil {
			var_5 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		_, err = templBuffer.WriteString("<div id=\"ws-connection\" hx-swap-oob=\"innerHTML\"><h1 class=\"text-2xl font-semibold mb-4\">")
		if err != nil {
			return err
		}
		var_6 := `Tic-Tac-Toe`
		_, err = templBuffer.WriteString(var_6)
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString("</h1><div id=\"board-container\" class=\"flex items-center justify-center\">")
		if err != nil {
			return err
		}
		err = Board().Render(ctx, templBuffer)
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
