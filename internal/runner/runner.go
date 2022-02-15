package runner

import (
	"os"

	"github.com/kankburhan/dofind/pkg/config"
	"github.com/kankburhan/dofind/pkg/dofind"
	"github.com/kankburhan/dofind/pkg/domains"
	"github.com/projectdiscovery/gologger"
)

func New(options *Options) {
	cfg := config.UserHomeDir() + pathfile + configfile
	if options.UpdateTld {
		gologger.Print().Msgf("Starting update from iana.org")
		err := domains.Download(cfg, "https://data.iana.org/TLD/tlds-alpha-by-domain.txt")
		if err != nil {
			gologger.Error().Msgf("Failed update domains")
			os.Exit(1)
		}
	}

	domain, err := config.Get(cfg)
	if err != nil {
		gologger.Error().Msgf("No domain list found, please update with -u, -update-domain")
		os.Exit(1)
	}
	options.Domains = domain

	gologger.Info().Msgf("%d domains loaded.", len(domain))
	dofind.Find(
		options.Target,
		options.Domains,
		options.Concurrency,
		options.Silent,
		options.Verbose,
		options.Output,
	)
}
