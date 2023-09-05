// Code generated by templ@v0.2.316 DO NOT EDIT.

package partials

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import "context"
import "io"
import "bytes"

func Box() templ.Component {
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
		_, err = templBuffer.WriteString("<div class=\"relative border-1 h-[10px] w-[10px]\"><button name=\"top\" class=\"bg-blue-600 absolute top-0 inset-x-0 w-[10px] h-[2px]\"></button><button name=\"left\" class=\"bg-blue-600 absolute left-0 inset-y-0 h-[10px] w-[2px]\"></button><button name=\"right\" class=\"bg-blue-600 absolute right-0 inset-y-0 h-[10px] w-[2px]\"></button><button name=\"bottom\" class=\"bg-blue-600 absolute bottom-0 inset-x-0 w-[10px] h-[2px]\"></button></div>")
		if err != nil {
			return err
		}
		if !templIsBuffer {
			_, err = io.Copy(w, templBuffer)
		}
		return err
	})
}

func Board(rows int, cols int) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) (err error) {
		templBuffer, templIsBuffer := w.(*bytes.Buffer)
		if !templIsBuffer {
			templBuffer = templ.GetBuffer()
			defer templ.ReleaseBuffer(templBuffer)
		}
		ctx = templ.InitializeContext(ctx)
		var_2 := templ.GetChildren(ctx)
		if var_2 == nil {
			var_2 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		_, err = templBuffer.WriteString("<div id=\"board\" class=\"grid grid-cols-10 grid-cols-10 w-[100px] h-[100px] gap-x-0\">")
		if err != nil {
			return err
		}
		for row := 0; row < 10; row++ {
			for col := 0; col < 10; col++ {
				_, err = templBuffer.WriteString("<div>")
				if err != nil {
					return err
				}
				err = Box().Render(ctx, templBuffer)
				if err != nil {
					return err
				}
				_, err = templBuffer.WriteString("</div>")
				if err != nil {
					return err
				}
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
