// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.680
package views

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import "context"
import "io"
import "bytes"

import "github.com/gustavomtborges/orcamento-auto/views/layouts"

func Login() templ.Component {
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
		templ_7745c5c3_Var2 := templ.ComponentFunc(func(ctx context.Context, templ_7745c5c3_W io.Writer) (templ_7745c5c3_Err error) {
			templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templ_7745c5c3_W.(*bytes.Buffer)
			if !templ_7745c5c3_IsBuffer {
				templ_7745c5c3_Buffer = templ.GetBuffer()
				defer templ.ReleaseBuffer(templ_7745c5c3_Buffer)
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<div class=\"flex min-h-full flex-col justify-center px-6 lg:px-8\"><div class=\"sm:mx-auto sm:w-full sm:max-w-sm\"><img class=\"mx-auto h-36 w-auto\" src=\"/static/logo-full.svg\" alt=\"Orçamento Auto logo\"><h2 class=\"mt-2 text-center text-2xl font-bold leading-9 tracking-tight text-gray-900\">Entre em sua conta</h2></div><div class=\"mt-2 sm:mx-auto sm:w-full sm:max-w-sm\"><form class=\"space-y-6\" action=\"/login\" method=\"POST\"><div class=\"w-full max-w-xs mx-auto\"><input name=\"email\" type=\"email\" placeholder=\"E-mail\" class=\"flex w-full h-10 px-3 py-2 text-sm bg-white border rounded-md border-neutral-300 ring-offset-background placeholder:text-neutral-500 focus:border-neutral-300 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-neutral-400 disabled:cursor-not-allowed disabled:opacity-50\"></div><div class=\"w-full max-w-xs mx-auto flex flex-col\"><input name=\"password\" type=\"password\" placeholder=\"Senha\" class=\"flex w-full h-10 px-3 py-2 text-sm bg-white border rounded-md border-neutral-300 ring-offset-background placeholder:text-neutral-500 focus:border-neutral-300 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-neutral-400 disabled:cursor-not-allowed disabled:opacity-50\"><div class=\"text-sm self-end mt-1\"><a href=\"#\" class=\"font-semibold text-cyan-600 hover:text-cyan-500\">Esqueceu a senha?</a></div></div><div class=\"w-full max-w-xs mx-auto\"><button type=\"submit\" class=\"inline-flex items-center justify-center px-4 py-2 text-sm font-medium tracking-wide text-white transition-colors duration-200 bg-cyan-600 rounded-md hover:bg-cyan-700 focus:ring-2 focus:ring-offset-2 focus:ring-blue-700 focus:shadow-outline focus:outline-none w-full\">Entrar</button></div></form><p class=\"mt-10 text-center text-sm text-gray-500\">Voltar <a href=\"/\" class=\"font-semibold leading-6 text-cyan-600 hover:text-cyan-500\">para início</a></p></div></div>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			if !templ_7745c5c3_IsBuffer {
				_, templ_7745c5c3_Err = io.Copy(templ_7745c5c3_W, templ_7745c5c3_Buffer)
			}
			return templ_7745c5c3_Err
		})
		templ_7745c5c3_Err = layouts.Base().Render(templ.WithChildren(ctx, templ_7745c5c3_Var2), templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}
