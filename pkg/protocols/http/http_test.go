package http

import (
	"testing"

	"github.com/stretchr/testify/require"

<<<<<<< HEAD:v2/pkg/protocols/http/http_test.go
	"github.com/Explorer1092/nuclei/v2/pkg/model"
	"github.com/Explorer1092/nuclei/v2/pkg/model/types/severity"
	"github.com/Explorer1092/nuclei/v2/pkg/protocols/common/generators"
	"github.com/Explorer1092/nuclei/v2/pkg/testutils"
=======
	"github.com/projectdiscovery/nuclei/v3/pkg/model"
	"github.com/projectdiscovery/nuclei/v3/pkg/model/types/severity"
	"github.com/projectdiscovery/nuclei/v3/pkg/protocols/common/generators"
	"github.com/projectdiscovery/nuclei/v3/pkg/testutils"
>>>>>>> 7f556f8e33080a9eb2e52a9c14fa2e16f32f62c3:pkg/protocols/http/http_test.go
)

func TestHTTPCompile(t *testing.T) {
	options := testutils.DefaultOptions
	options.CustomHeaders = []string{"User-Agent: test", "Hello: World"}

	testutils.Init(options)
	templateID := "testing-http"
	request := &Request{
		Name: "testing",
		Payloads: map[string]interface{}{
			"username": []string{"admin"},
			"password": []string{"admin", "guest", "password", "test", "12345", "123456"},
		},
		AttackType: generators.AttackTypeHolder{Value: generators.ClusterBombAttack},
		Raw: []string{`GET /manager/html HTTP/1.1
Host: {{Hostname}}
User-Agent: Nuclei - Open-source project (github.com/Explorer1092/nuclei)
Connection: close
Authorization: Basic {{username + ':' + password}}
Accept-Encoding: gzip`},
	}
	executerOpts := testutils.NewMockExecuterOptions(options, &testutils.TemplateInfo{
		ID:   templateID,
		Info: model.Info{SeverityHolder: severity.Holder{Severity: severity.Low}, Name: "test"},
	})
	err := request.Compile(executerOpts)
	require.Nil(t, err, "could not compile http request")
	require.Equal(t, 6, request.Requests(), "could not get correct number of requests")
	require.Equal(t, map[string]string{"User-Agent": "test", "Hello": "World"}, request.customHeaders, "could not get correct custom headers")
}
