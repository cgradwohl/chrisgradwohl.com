// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.598
package component

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import "context"
import "io"
import "bytes"

func MainLayout() templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, templ_7745c5c3_W io.Writer) (templ_7745c5c3_Err error) {
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templ_7745c5c3_W.(*bytes.Buffer)
		if !templ_7745c5c3_IsBuffer {
			templ_7745c5c3_Buffer = templ.GetBuffer()
			defer templ.ReleaseBuffer(templ_7745c5c3_Buffer)
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var1 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var1 == nil {
			templ_7745c5c3_Var1 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<!doctype html><html lang=\"en\"><head><meta charset=\"UTF-8\"><meta name=\"viewport\" content=\"width=device-width, initial-scale=1.0\"><meta name=\"color-scheme\" content=\"light dark\"><link rel=\"icon\" type=\"image/png\" href=\"/img/favicon.ico\"><link rel=\"stylesheet\" href=\"/css/pico.jade.min.css\"><title>Chris Gradwohl</title></head><body><header><nav class=\"container\"><ul></ul><ul><li><a href=\"/\">Home</a></li><li><a href=\"/interests\">Interests</a></li><li><a href=\"/blog\">Blog</a></li></ul></nav></header><main class=\"container\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = templ_7745c5c3_Var1.Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</main><footer class=\"container\"></footer></body></html><style>\n\t\theader {\n\t\t\tpadding-block: 0 !important;\n\t\t\tbox-shadow: 0 .125rem .25rem rgba(255,255,255,.11);\n\t\t}\n\t\tmain {\n\t\t\tmargin: 2% 0;\n\t\t}\n\n\t\t.container {\n\t\t\tmax-width: 1200px;\n\t\t}\n\t</style>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}