package goconsole

import (
	lib_goconsole "github.com/Explorer1092/nuclei/v3/pkg/js/libs/goconsole"

	"github.com/Explorer1092/nuclei/v3/pkg/js/gojs"
	"github.com/dop251/goja"
)

var (
	module = gojs.NewGojaModule("nuclei/goconsole")
)

func init() {
	module.Set(
		gojs.Objects{
			// Functions
			"NewGoConsolePrinter": lib_goconsole.NewGoConsolePrinter,

			// Var and consts

			// Objects / Classes
			"GoConsolePrinter": gojs.GetClassConstructor[lib_goconsole.GoConsolePrinter](&lib_goconsole.GoConsolePrinter{}),
		},
	).Register()
}

func Enable(runtime *goja.Runtime) {
	module.Enable(runtime)
}
