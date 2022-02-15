package domains

import (
	"fmt"
	"github.com/kankburhan/dofind/pkg/config"
	"io"
	"log"
	"net/http"
	"strings"

	"github.com/projectdiscovery/gologger"
)

func Download(path string, url string) error {
	resp, err := http.Get(url)
	if err != nil {
		gologger.Error().Msgf("Failed download list domains")
		return err
	}

	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		bodyBytes, err := io.ReadAll(resp.Body)
		if err != nil {
			log.Fatal(err)
		}
		bodyString := string(bodyBytes)
		er := config.Save(path, strings.Split(bodyString, "\n"))
		if er != nil {
			return fmt.Errorf("failed to save domain")
		}
		return nil
	}

	return fmt.Errorf("failed download domain")
}
