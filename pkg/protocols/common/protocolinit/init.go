package protocolinit

import (
	"github.com/Explorer1092/nuclei/v3/pkg/js/compiler"
	"github.com/Explorer1092/nuclei/v3/pkg/protocols/common/protocolstate"
	"github.com/Explorer1092/nuclei/v3/pkg/protocols/dns/dnsclientpool"
	"github.com/Explorer1092/nuclei/v3/pkg/protocols/http/httpclientpool"
	"github.com/Explorer1092/nuclei/v3/pkg/protocols/http/signerpool"
	"github.com/Explorer1092/nuclei/v3/pkg/protocols/network/networkclientpool"
	"github.com/Explorer1092/nuclei/v3/pkg/protocols/whois/rdapclientpool"
	"github.com/Explorer1092/nuclei/v3/pkg/types"
)

// Init initializes the client pools for the protocols
func Init(options *types.Options) error {

	if err := protocolstate.Init(options); err != nil {
		return err
	}
	if err := dnsclientpool.Init(options); err != nil {
		return err
	}
	if err := httpclientpool.Init(options); err != nil {
		return err
	}
	if err := signerpool.Init(options); err != nil {
		return err
	}
	if err := networkclientpool.Init(options); err != nil {
		return err
	}
	if err := rdapclientpool.Init(options); err != nil {
		return err
	}
	if err := compiler.Init(options); err != nil {
		return err
	}
	return nil
}

func Close() {
	protocolstate.Dialer.Close()
}
