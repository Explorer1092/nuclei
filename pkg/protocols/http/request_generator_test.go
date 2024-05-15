package http

import (
	"testing"

	"github.com/stretchr/testify/require"

<<<<<<< HEAD:v2/pkg/protocols/http/request_generator_test.go
	"github.com/Explorer1092/nuclei/v2/pkg/catalog/disk"
	"github.com/Explorer1092/nuclei/v2/pkg/protocols/common/generators"
=======
	"github.com/projectdiscovery/nuclei/v3/pkg/catalog/disk"
	"github.com/projectdiscovery/nuclei/v3/pkg/protocols/common/generators"
	"github.com/projectdiscovery/nuclei/v3/pkg/types"
>>>>>>> 7f556f8e33080a9eb2e52a9c14fa2e16f32f62c3:pkg/protocols/http/request_generator_test.go
)

func TestRequestGeneratorPaths(t *testing.T) {
	req := &Request{
		Path: []string{"{{BaseURL}}/test", "{{BaseURL}}/test.php"},
	}
	generator := req.newGenerator(false)
	var payloads []string
	for {
		raw, _, ok := generator.nextValue()
		if !ok {
			break
		}
		payloads = append(payloads, raw)
	}
	require.Equal(t, req.Path, payloads, "Could not get correct paths")
}

func TestRequestGeneratorClusterBombSingle(t *testing.T) {
	var err error

	req := &Request{
		Payloads:   map[string]interface{}{"username": []string{"admin", "tomcat", "manager"}, "password": []string{"password", "test", "secret"}},
		AttackType: generators.AttackTypeHolder{Value: generators.ClusterBombAttack},
		Raw:        []string{`GET /{{username}}:{{password}} HTTP/1.1`},
	}
	catalogInstance := disk.NewCatalog("")
	req.generator, err = generators.New(req.Payloads, req.AttackType.Value, "", catalogInstance, "", types.DefaultOptions())
	require.Nil(t, err, "could not create generator")

	generator := req.newGenerator(false)
	var payloads []map[string]interface{}
	for {
		_, data, ok := generator.nextValue()
		if !ok {
			break
		}
		payloads = append(payloads, data)
	}
	require.Equal(t, 9, len(payloads), "Could not get correct number of payloads")
}

func TestRequestGeneratorClusterBombMultipleRaw(t *testing.T) {
	var err error

	req := &Request{
		Payloads:   map[string]interface{}{"username": []string{"admin", "tomcat", "manager"}, "password": []string{"password", "test", "secret"}},
		AttackType: generators.AttackTypeHolder{Value: generators.ClusterBombAttack},
		Raw:        []string{`GET /{{username}}:{{password}} HTTP/1.1`, `GET /{{username}}@{{password}} HTTP/1.1`},
	}
	catalogInstance := disk.NewCatalog("")
	req.generator, err = generators.New(req.Payloads, req.AttackType.Value, "", catalogInstance, "", types.DefaultOptions())
	require.Nil(t, err, "could not create generator")

	generator := req.newGenerator(false)
	var payloads []map[string]interface{}
	for {
		_, data, ok := generator.nextValue()
		if !ok {
			break
		}
		payloads = append(payloads, data)
	}
	require.Equal(t, 18, len(payloads), "Could not get correct number of payloads")
}
