package telnet

import (
	lib_telnet "github.com/Explorer1092/nuclei/v3/pkg/js/libs/telnet"

	"github.com/Explorer1092/nuclei/v3/pkg/js/gojs"
	"github.com/dop251/goja"
)

var (
	module = gojs.NewGojaModule("nuclei/telnet")
)

func init() {
	module.Set(
		gojs.Objects{
			// Functions
			"IsTelnet": lib_telnet.IsTelnet,

			// Var and consts

			// Objects / Classes
			"IsTelnetResponse": gojs.GetClassConstructor[lib_telnet.IsTelnetResponse](&lib_telnet.IsTelnetResponse{}),
		},
	).Register()
}

func Enable(runtime *goja.Runtime) {
	module.Enable(runtime)
}
