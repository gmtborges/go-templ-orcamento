// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.747
package views

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import templruntime "github.com/a-h/templ/runtime"

import (
	"encoding/json"
	"fmt"
	"github.com/gmtborges/orcamento-auto/types"
	"github.com/gmtborges/orcamento-auto/views/layouts"
)

func orcamentoCreateMetaTags() templ.Component {
	return templruntime.GeneratedTemplate(func(templ_7745c5c3_Input templruntime.GeneratedComponentInput) (templ_7745c5c3_Err error) {
		templ_7745c5c3_W, ctx := templ_7745c5c3_Input.Writer, templ_7745c5c3_Input.Context
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templruntime.GetBuffer(templ_7745c5c3_W)
		if !templ_7745c5c3_IsBuffer {
			defer func() {
				templ_7745c5c3_BufErr := templruntime.ReleaseBuffer(templ_7745c5c3_Buffer)
				if templ_7745c5c3_Err == nil {
					templ_7745c5c3_Err = templ_7745c5c3_BufErr
				}
			}()
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var1 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var1 == nil {
			templ_7745c5c3_Var1 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<meta name=\"description\" content=\"Lista de orçamentos\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return templ_7745c5c3_Err
	})
}

func vmToAlpine(vm types.OrcamentoCreateViewModel) string {
	produtoJson, _ := json.Marshal(vm.AutoCategorias["acProduto"])
	servicoJson, _ := json.Marshal(vm.AutoCategorias["acServico"])
	return fmt.Sprintf(`
  {
    itens: [],
    acTipo: "",
    acID: "",
    acDesc: "",
    acObs: "",
    acProduto: %s,  
    acServico: %s,
    addItem() {
      this.itens.push({
        acTipo: this.acTipo,
        acID: this.acID,
        acDesc: this.acDesc,
        acObs: this.acObs
        });
      this.acTipo = "";
      this.acID = "";
      this.acDesc = "";
      this.acObs = "";
    }
  }`,
		produtoJson,
		servicoJson,
	)
}

func selectVeiculoCor() templ.Component {
	return templruntime.GeneratedTemplate(func(templ_7745c5c3_Input templruntime.GeneratedComponentInput) (templ_7745c5c3_Err error) {
		templ_7745c5c3_W, ctx := templ_7745c5c3_Input.Writer, templ_7745c5c3_Input.Context
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templruntime.GetBuffer(templ_7745c5c3_W)
		if !templ_7745c5c3_IsBuffer {
			defer func() {
				templ_7745c5c3_BufErr := templruntime.ReleaseBuffer(templ_7745c5c3_Buffer)
				if templ_7745c5c3_Err == nil {
					templ_7745c5c3_Err = templ_7745c5c3_BufErr
				}
			}()
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var2 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var2 == nil {
			templ_7745c5c3_Var2 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<select id=\"veiculoCor\" name=\"veiculoCor\" class=\"select select-bordered mb-5\" required><option value=\"branco\">Branco</option> <option value=\"preto\">Preto</option> <option value=\"cinza\">Cinza</option> <option value=\"prata\">Prata</option> <option value=\"vermelho\">Vermelho</option> <option value=\"azul\">Azul</option> <option value=\"outro\">Outro (informe na observação)</option></select>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return templ_7745c5c3_Err
	})
}

func OrcamentoCreate(vm types.OrcamentoCreateViewModel) templ.Component {
	return templruntime.GeneratedTemplate(func(templ_7745c5c3_Input templruntime.GeneratedComponentInput) (templ_7745c5c3_Err error) {
		templ_7745c5c3_W, ctx := templ_7745c5c3_Input.Writer, templ_7745c5c3_Input.Context
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templruntime.GetBuffer(templ_7745c5c3_W)
		if !templ_7745c5c3_IsBuffer {
			defer func() {
				templ_7745c5c3_BufErr := templruntime.ReleaseBuffer(templ_7745c5c3_Buffer)
				if templ_7745c5c3_Err == nil {
					templ_7745c5c3_Err = templ_7745c5c3_BufErr
				}
			}()
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var3 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var3 == nil {
			templ_7745c5c3_Var3 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		templ_7745c5c3_Var4 := templruntime.GeneratedTemplate(func(templ_7745c5c3_Input templruntime.GeneratedComponentInput) (templ_7745c5c3_Err error) {
			templ_7745c5c3_W, ctx := templ_7745c5c3_Input.Writer, templ_7745c5c3_Input.Context
			templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templruntime.GetBuffer(templ_7745c5c3_W)
			if !templ_7745c5c3_IsBuffer {
				defer func() {
					templ_7745c5c3_BufErr := templruntime.ReleaseBuffer(templ_7745c5c3_Buffer)
					if templ_7745c5c3_Err == nil {
						templ_7745c5c3_Err = templ_7745c5c3_BufErr
					}
				}()
			}
			ctx = templ.InitializeContext(ctx)
			templ_7745c5c3_Var5 := templruntime.GeneratedTemplate(func(templ_7745c5c3_Input templruntime.GeneratedComponentInput) (templ_7745c5c3_Err error) {
				templ_7745c5c3_W, ctx := templ_7745c5c3_Input.Writer, templ_7745c5c3_Input.Context
				templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templruntime.GetBuffer(templ_7745c5c3_W)
				if !templ_7745c5c3_IsBuffer {
					defer func() {
						templ_7745c5c3_BufErr := templruntime.ReleaseBuffer(templ_7745c5c3_Buffer)
						if templ_7745c5c3_Err == nil {
							templ_7745c5c3_Err = templ_7745c5c3_BufErr
						}
					}()
				}
				ctx = templ.InitializeContext(ctx)
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<div class=\"flex flex-col items-center mx-20 mt-6\"><h1 class=\"font-bold text-4xl mb-10\">Novo orçamento</h1><main x-data=\"")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				var templ_7745c5c3_Var6 string
				templ_7745c5c3_Var6, templ_7745c5c3_Err = templ.JoinStringErrs(vmToAlpine(vm))
				if templ_7745c5c3_Err != nil {
					return templ.Error{Err: templ_7745c5c3_Err, FileName: `views/orcamento/create.templ`, Line: 62, Col: 28}
				}
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var6))
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\" class=\"mt-5 flex flex-col items-center w-full max-w-5xl\"><form hx-boost=\"false\" action=\"/orcamentos/salvar\" method=\"POST\" class=\"flex flex-col bg-gray-50 dark:bg-gray-950 p-5 rounded-lg w-full mb-6\"><div class=\"flex flex-col lg:flex-row gap-10\"><div class=\"form-control w-full\"><label for=\"associadoNome\" class=\"label\">Nome do associado *</label> <input id=\"associadoNome\" name=\"associadoNome\" type=\"text\" class=\"input input-bordered mb-5\" placeholder=\"João da Silva\" value=\"")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				var templ_7745c5c3_Var7 string
				templ_7745c5c3_Var7, templ_7745c5c3_Err = templ.JoinStringErrs(vm.AssociadoNome)
				if templ_7745c5c3_Err != nil {
					return templ.Error{Err: templ_7745c5c3_Err, FileName: `views/orcamento/create.templ`, Line: 80, Col: 33}
				}
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var7))
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\" required></div><div class=\"form-control w-full\"><label for=\"veiculoMarca\" class=\"label\">Marca do veículo *</label> <input id=\"veiculoMarca\" name=\"veiculoMarca\" class=\"input input-bordered mb-5\" placeholder=\"Ford, Hyunday, Volkswagen ...\" value=\"")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				var templ_7745c5c3_Var8 string
				templ_7745c5c3_Var8, templ_7745c5c3_Err = templ.JoinStringErrs(vm.VeiculoMarca)
				if templ_7745c5c3_Err != nil {
					return templ.Error{Err: templ_7745c5c3_Err, FileName: `views/orcamento/create.templ`, Line: 91, Col: 32}
				}
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var8))
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\" required></div></div><div class=\"flex flex-col lg:flex-row gap-10\"><div class=\"form-control w-full\"><label for=\"veiculoNome\" class=\"label\">Nome do veículo *</label> <input id=\"veiculoNome\" name=\"veiculoNome\" type=\"text\" class=\"input input-bordered mb-5\" placeholder=\"Gol, Creta, Hilux ...\" value=\"")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				var templ_7745c5c3_Var9 string
				templ_7745c5c3_Var9, templ_7745c5c3_Err = templ.JoinStringErrs(vm.VeiculoNome)
				if templ_7745c5c3_Err != nil {
					return templ.Error{Err: templ_7745c5c3_Err, FileName: `views/orcamento/create.templ`, Line: 105, Col: 31}
				}
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var9))
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\" required></div><div class=\"form-control w-full\"><label for=\"veiculoAno\" class=\"label\">Ano do veículo *</label> <input id=\"veiculoAno\" name=\"veiculoAno\" class=\"input input-bordered mb-5\" type=\"number\" min=\"1950\" max=\"2024\" value=\"")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				var templ_7745c5c3_Var10 string
				templ_7745c5c3_Var10, templ_7745c5c3_Err = templ.JoinStringErrs(fmt.Sprintf("%d", vm.VeiculoAno))
				if templ_7745c5c3_Err != nil {
					return templ.Error{Err: templ_7745c5c3_Err, FileName: `views/orcamento/create.templ`, Line: 118, Col: 49}
				}
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var10))
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\" required></div></div><div class=\"flex flex-col lg:flex-row gap-10\"><div class=\"form-control w-full\"><label for=\"veiculoCor\" class=\"label\">Cor do veículo *</label>")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				templ_7745c5c3_Err = selectVeiculoCor().Render(ctx, templ_7745c5c3_Buffer)
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</div><div class=\"form-control w-full\"><label for=\"observacao\" class=\"label\">Descrição</label> <textarea id=\"observacao\" name=\"observacao\" class=\"textarea textarea-bordered mb-5\" rows=\"3\" placeholder=\"Alguma observação sobre o associado ou veículo.\"></textarea></div></div><h2 class=\"font-bold text-2xl\">Items</h2><p class=\"text-error font-bold my-2\">")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				var templ_7745c5c3_Var11 string
				templ_7745c5c3_Var11, templ_7745c5c3_Err = templ.JoinStringErrs(vm.Errors["itens"])
				if templ_7745c5c3_Err != nil {
					return templ.Error{Err: templ_7745c5c3_Err, FileName: `views/orcamento/create.templ`, Line: 140, Col: 63}
				}
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var11))
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</p><div class=\"flex flex-col mt-5 overflow-y-auto max-h-64\"><table class=\"table table-zebra\"><thead class=\"text-base\"><tr><th>#</th><th>Tipo</th><th>Descrição</th><th>Observação</th><th>Remover</th></tr></thead> <tbody><template x-for=\"(item, i) in itens.reverse()\"><tr><td x-text=\"itens.length - i\"></td><td x-text=\"item.acTipo === &#39;servico&#39; ? &#39;Serviço&#39; : &#39;Produto&#39;\"></td><td x-text=\"item.acDesc\"></td><td x-text=\"item.acObs\"></td><td><button class=\"text-error p-1 hover:bg-base-300 rounded\" @click=\"itens = [...itens.slice(i + 1), ...itens.slice(0, i)]\"><svg xmlns=\"http://www.w3.org/2000/svg\" fill=\"none\" viewBox=\"0 0 24 24\" stroke-width=\"1.5\" stroke=\"currentColor\" class=\"w-6\"><path stroke-linecap=\"round\" stroke-linejoin=\"round\" d=\"M6 18L18 6M6 6l12 12\"></path></svg></button></td></tr></template><input type=\"hidden\" name=\"itens\" :value=\"JSON.stringify(itens.map(i =&gt; ({ autoCategoriaID: parseInt(i.acID), observacao: i.acObs })))\"></tbody></table></div><div class=\"flex items-center mt-10 justify-end gap-4\"><p class=\"text-error font-bold\">")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				var templ_7745c5c3_Var12 string
				templ_7745c5c3_Var12, templ_7745c5c3_Err = templ.JoinStringErrs(vm.Errors["db"])
				if templ_7745c5c3_Err != nil {
					return templ.Error{Err: templ_7745c5c3_Err, FileName: `views/orcamento/create.templ`, Line: 180, Col: 56}
				}
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var12))
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</p><button class=\"text-lg self-end btn btn-wide btn-primary\" type=\"submit\" :disabled=\"!itens.length\">Salvar</button></div></form><div class=\"flex flex-col w-full rounded-lg bg-gray-50 dark:bg-gray-950 p-5\"><p class=\"text-error font-bold mb-2\">")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				var templ_7745c5c3_Var13 string
				templ_7745c5c3_Var13, templ_7745c5c3_Err = templ.JoinStringErrs(vm.Errors["ac"])
				if templ_7745c5c3_Err != nil {
					return templ.Error{Err: templ_7745c5c3_Err, FileName: `views/orcamento/create.templ`, Line: 189, Col: 60}
				}
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var13))
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</p><div class=\"flex flex-col gap-5 lg:flex-row w-full\"><select name=\"acTipo\" class=\"select select-bordered mb-5 w-full lg:w-1/2\" x-model=\"acTipo\" @change=\"acID = &#39;&#39;\"><option value=\"\" disabled selected>Selecione o tipo</option> <option value=\"servico\">Serviço</option> <option value=\"produto\">Peça</option></select> <select name=\"descricaoItem\" class=\"select select-bordered mb-5 w-full\" x-model=\"acID\" @change=\"acDesc = $el.options[$el.selectedIndex].innerHTML\"><option value=\"\" disabled selected>Selecione o serviço ou produto</option><template x-if=\"acTipo === &#39;servico&#39;\"><template x-for=\"ac in acServico\"><option :value=\"ac.ID\" x-text=\"ac.Descricao\"></option></template></template><template x-if=\"acTipo === &#39;produto&#39;\"><template x-for=\"ac in acProduto\"><option :value=\"ac.ID\" x-text=\"ac.Descricao\"></option></template></template></select> <textarea name=\"acObs\" x-model=\"acObs\" class=\"textarea textarea-bordered mb-5 w-full\" rows=\"1\" placeholder=\"Observação sobre o serviço ou peça\"></textarea></div><button :disabled=\"acTipo == &#39;&#39; || acID == &#39;&#39;\" class=\"mt-2 btn btn-primary text-lg font-semibold\" @click=\"addItem()\"><svg xmlns=\"http://www.w3.org/2000/svg\" fill=\"none\" viewBox=\"0 0 24 24\" stroke-width=\"1.5\" stroke=\"currentColor\" class=\"w-6\"><path stroke-linecap=\"round\" stroke-linejoin=\"round\" d=\"M12 4.5v15m7.5-7.5h-15\"></path></svg> Adicionar item no orçamento</button></div></main></div>")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				return templ_7745c5c3_Err
			})
			templ_7745c5c3_Err = layouts.SideBarOrg().Render(templ.WithChildren(ctx, templ_7745c5c3_Var5), templ_7745c5c3_Buffer)
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			return templ_7745c5c3_Err
		})
		templ_7745c5c3_Err = layouts.Base("Orçamento Auto - Novo orçamento", orcamentoCreateMetaTags()).Render(templ.WithChildren(ctx, templ_7745c5c3_Var4), templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return templ_7745c5c3_Err
	})
}
