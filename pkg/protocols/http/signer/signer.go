package signer

import (
	"context"
	"errors"
	"net/http"

<<<<<<< HEAD:v2/pkg/protocols/http/signer/signer.go
	"github.com/Explorer1092/nuclei/v2/pkg/types"
=======
	"github.com/projectdiscovery/nuclei/v3/pkg/types"
>>>>>>> 419f08f61ce5ca2d3f0eae9fe36dc7c44c1f532a:pkg/protocols/http/signer/signer.go
)

// An Argument that can be passed to Signer
type SignerArg string

type Signer interface {
	SignHTTP(ctx context.Context, request *http.Request) error
}

type SignerArgs interface {
	Validate() error
}

func NewSigner(args SignerArgs) (signer Signer, err error) {
	switch signerArgs := args.(type) {
	case *AWSOptions:
		awsSigner, err := NewAwsSigner(signerArgs)
		if err != nil {
			awsSigner, err = NewAwsSignerFromConfig(signerArgs)
			if err != nil {
				return nil, err
			}
		}
		return awsSigner, err
	default:
		return nil, errors.New("unknown signature arguments type")
	}
}

// GetCtxWithArgs creates and returns context with signature args
func GetCtxWithArgs(maps ...map[string]interface{}) context.Context {
	var region, service string
	for _, v := range maps {
		for key, val := range v {
			if key == "region" && region == "" {
				region = types.ToString(val)
			}
			if key == "service" && service == "" {
				service = types.ToString(val)
			}
		}
	}
	// type ctxkey string
	ctx := context.WithValue(context.Background(), SignerArg("service"), service)
	return context.WithValue(ctx, SignerArg("region"), region)
}
