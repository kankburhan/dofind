package dofind

import (
	"fmt"
	"github.com/kankburhan/dofind/pkg/errors"
	"github.com/logrusorgru/aurora/v3"
	"github.com/reiver/go-telnet"
	"net/http"
	"net/url"
	"os"
	"strings"
	"sync"
)

func Find(target string, domains []string, concurrency int, silent, verbose bool, output *os.File) {
	jobs := make(chan string)
	var wg sync.WaitGroup

	for i := 0; i < concurrency; i++ {
		wg.Add(1)
		go func() {
			for domain := range jobs {
				domain := fmt.Sprintf("%s.%s", target, strings.ToLower(domain))
				res, e := GetHttp(domain)
				if e != nil {
					if !silent {
						if verbose {
							errors.Show(e.Error())
						} else {
							errors.Show(domain)
						}
					}

					err := GetTelnet(domain)
					if err != nil {
						if !silent {
							if verbose {
								errors.Show(err.Error())
							} else {
								errors.Show(domain)
							}
						}
					} else {
						if silent {
							fmt.Println(domain)
						} else {
							fmt.Printf("[%s] %s\n", aurora.Green("TELNET OK").String(), aurora.Green(domain).String())
						}

						if output != nil {
							_, f := output.WriteString(fmt.Sprintf("%s\n", domain))
							if f != nil {
								errors.Show(f.Error())
							}
						}
					}
				}

				if res != nil {
					if silent {
						fmt.Println(domain)
					} else {
						fmt.Printf("[%s] %s\n", aurora.Green("HTTP OK").String(), aurora.Green(domain).String())
					}

					if output != nil {
						_, f := output.WriteString(fmt.Sprintf("%s\n", domain))
						if f != nil {
							errors.Show(f.Error())
						}
					}
				}
			}
		}()
	}

	for _, t := range domains[1:] {
		jobs <- t
	}

	close(jobs)
	wg.Wait()
}

func GetHttp(s string) (*http.Response, error) {
	//default https
	schema := "https"
	res, _ := http.Get(getUrlString(s, schema))
	if res != nil && res.StatusCode == http.StatusOK {
		return res, nil
	}

	schema = "http"
	res, _ = http.Get(getUrlString(s, schema))
	if res != nil && res.StatusCode == http.StatusOK {
		return res, nil
	}
	return nil, fmt.Errorf("domain not available")
}

func GetTelnet(s string) error {
	var caller = telnet.StandardCaller

	// dial port 80
	err := telnet.DialToAndCall(fmt.Sprintf("%s:%s", s, "80"), caller)
	if err != nil {
		return err
	}

	// dial port 443
	err = telnet.DialToAndCall(fmt.Sprintf("%s:%s", s, "443"), caller)
	if err != nil {
		return err
	}
	return nil
}

func getUrlString(h, schema string) string {
	u := url.URL{
		Scheme: schema,
		Host:   h,
	}
	return u.String()
}
