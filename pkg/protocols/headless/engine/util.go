package engine

import (
<<<<<<< HEAD
	"github.com/Explorer1092/nuclei/v3/pkg/protocols/common/marker"
=======
<<<<<<< HEAD:v2/pkg/protocols/headless/engine/util.go
	"github.com/Explorer1092/nuclei/v2/pkg/protocols/common/marker"
=======
	"github.com/projectdiscovery/nuclei/v3/pkg/protocols/common/marker"
>>>>>>> 419f08f61ce5ca2d3f0eae9fe36dc7c44c1f532a:pkg/protocols/headless/engine/util.go
>>>>>>> projectdiscovery-main
	"github.com/valyala/fasttemplate"
)

func replaceWithValues(data string, values map[string]interface{}) string {
	return fasttemplate.ExecuteStringStd(data, marker.ParenthesisOpen, marker.ParenthesisClose, values)
}
