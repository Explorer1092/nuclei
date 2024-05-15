package engine

import (
<<<<<<< HEAD:v2/pkg/protocols/headless/engine/util.go
	"github.com/Explorer1092/nuclei/v2/pkg/protocols/common/marker"
=======
	"github.com/projectdiscovery/nuclei/v3/pkg/protocols/common/marker"
>>>>>>> 7f556f8e33080a9eb2e52a9c14fa2e16f32f62c3:pkg/protocols/headless/engine/util.go
	"github.com/valyala/fasttemplate"
)

func replaceWithValues(data string, values map[string]interface{}) string {
	return fasttemplate.ExecuteStringStd(data, marker.ParenthesisOpen, marker.ParenthesisClose, values)
}
