package main

import (
	"flag"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"path"
	"sort"
	"strings"
	"text/tabwriter"
)

var (
	verbose = flag.Bool("verbose", false, "be verbose")
	web     = flag.String("http", "", "open web browser instead")
)

const (
	statusOK            = "OK"
	statusInformational = "INFORMATIONAL"
	statusUnimplemented = "UNIMPLEMENTED"
	statusNonStrict     = "NON-STRICT"
	statusUnclean       = "UNCLEAN"
	statusFailed        = "FAILED"
)

//go:norace
func failing(behavior string) bool {
	_ = "STUB: not implemented"

	// case statusUnclean, statusFailed, statusNonStrict: // we should probably fix the nonstrict as well at some point
	return false
}

type statusCounter struct {
	Total         int
	OK            int
	Informational int
	Unimplemented int
	NonStrict     int
	Unclean       int
	Failed        int
}

//go:norace
func (c *statusCounter) Inc(s string) { _ = "STUB: not implemented"; return }

//go:norace
func main() {
	log.SetFlags(0)
	flag.Parse()

	if flag.NArg() < 1 {
		log.Fatalf("Usage: %s [options] <report-path>", os.Args[0])
	}

	base := path.Dir(flag.Arg(0))

	if addr := *web; addr != "" {
		http.HandleFunc("/", handlerIndex())
		http.Handle("/report/", http.StripPrefix("/report/",
			http.FileServer(http.Dir(base)),
		))
		log.Fatal(http.ListenAndServe(addr, nil))
		return
	}

	var report report
	if err := decodeFile(os.Args[1], &report); err != nil {
		log.Fatal(err)
	}

	var servers []string
	for s := range report {
		servers = append(servers, s)
	}
	sort.Strings(servers)

	var (
		failed bool
	)
	tw := tabwriter.NewWriter(os.Stderr, 0, 4, 1, ' ', 0)
	for _, server := range servers {
		var (
			srvFailed  bool
			hdrWritten bool
			counter    statusCounter
		)

		var cases []string
		for id := range report[server] {
			cases = append(cases, id)
		}
		sortBySegment(cases)
		for _, id := range cases {
			c := report[server][id]

			var r entryReport
			err := decodeFile(path.Join(base, c.ReportFile), &r)
			if err != nil {
				log.Fatal(err)
			}
			counter.Inc(c.Behavior)
			bad := failing(c.Behavior)
			if bad {
				srvFailed = true
				failed = true
			}
			if *verbose || bad {
				if !hdrWritten {
					hdrWritten = true
					n, _ := fmt.Fprintf(os.Stderr, "AGENT %q\n", server)
					_, _ = fmt.Fprintf(tw, "%s\n", strings.Repeat("=", n-1))
				}
				_, _ = fmt.Fprintf(tw, "%s\t%s\t%s\n", server, id, c.Behavior)
			}
			if bad {
				_, _ = fmt.Fprintf(tw, "\tdesc:\t%s\n", r.Description)
				_, _ = fmt.Fprintf(tw, "\texp: \t%s\n", r.Expectation)
				_, _ = fmt.Fprintf(tw, "\tact: \t%s\n", r.Result)
			}
		}
		if hdrWritten {
			_, _ = fmt.Fprint(tw, "\n")
		}
		var status string
		if srvFailed {
			status = statusFailed
		} else {
			status = statusOK
		}
		n, _ := fmt.Fprintf(tw, "AGENT %q SUMMARY (%s)\n", server, status)
		_, _ = fmt.Fprintf(tw, "%s\n", strings.Repeat("=", n-1))

		_, _ = fmt.Fprintf(tw, "TOTAL:\t%d\n", counter.Total)
		_, _ = fmt.Fprintf(tw, "%s:\t%d\n", statusOK, counter.OK)
		_, _ = fmt.Fprintf(tw, "%s:\t%d\n", statusInformational, counter.Informational)
		_, _ = fmt.Fprintf(tw, "%s:\t%d\n", statusUnimplemented, counter.Unimplemented)
		_, _ = fmt.Fprintf(tw, "%s:\t%d\n", statusNonStrict, counter.NonStrict)
		_, _ = fmt.Fprintf(tw, "%s:\t%d\n", statusUnclean, counter.Unclean)
		_, _ = fmt.Fprintf(tw, "%s:\t%d\n", statusFailed, counter.Failed)
		_, _ = fmt.Fprint(tw, "\n")
		_ = tw.Flush()
	}
	var rc int
	if failed {
		rc = 1
		_, _ = fmt.Fprintf(tw, "\n\nTEST %s\n\n", statusFailed)
	} else {
		_, _ = fmt.Fprintf(tw, "\n\nTEST %s\n\n", statusOK)
	}

	_ = tw.Flush()
	os.Exit(rc)
}

type report map[string]server

type server map[string]entry

type entry struct {
	Behavior        string `json:"behavior"`
	BehaviorClose   string `json:"behaviorClose"`
	Duration        int    `json:"duration"`
	RemoveCloseCode int    `json:"removeCloseCode"`
	ReportFile      string `json:"reportFile"`
}

type entryReport struct {
	Description string `json:"description"`
	Expectation string `json:"expectation"`
	Result      string `json:"result"`
	Duration    int    `json:"duration"`
}

//go:norace
func decodeFile(path string, x interface{}) error { _ = "STUB: not implemented"; return nil }

//go:norace
func compareBySegment(a, b string) int { _ = "STUB: not implemented"; return 0 }

//go:norace
func mustInt(s string) int64 { _ = "STUB: not implemented"; return 0 }

//go:norace
func min(a, b int) int { _ = "STUB: not implemented"; return 0 }

//go:norace
func handlerIndex() func(w http.ResponseWriter, r *http.Request) {
	_ = "STUB: not implemented"
	return nil
}

var index = template.Must(template.New("").Parse(`
<html>
<body>
<h1>Welcome to WebSocket test server!</h1>
<h4>Ready to Autobahn!</h4>
<a href="/report">Reports</a>
</body>
</html>
`))

//go:norace
func sortBySegment(s []string) { _ = "STUB: not implemented"; return }
