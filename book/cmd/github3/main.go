package main

import (
	"flag"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"strings"

	"book/ch04"
)

var (
	host = flag.String("host", "", "listening host")
	port = flag.Int("port", 8013, "listening port")
)

var issueList = template.Must(template.New("issueList").Parse(`
<h1>{{.TotalCount}} issues</h1>
<table>
<tr style='text-align: left'>
  <th>#</th>
  <th>State</th>
  <th>User</th>
  <th>Title</th>
</tr>
{{range .Items}}
<tr>
  <td><a href='{{.HTMLURL}}'>{{.Number}}</a></td>
  <td>{{.State}}</td>
  <td><a href='{{.User.HTMLURL}}'>{{.User.Login}}</a></td>
  <td><a href='{{.HTMLURL}}'>{{.Title}}</a></td>
</tr>
{{end}}
</table>
`))

func main() {
	flag.Parse()
	if *host == "" {
		if err := cmdHandler(); err != nil {
			log.Fatal(err)
		}
	}
	addr := fmt.Sprintf("%s:%d", *host, *port)
	mux := http.NewServeMux()
	mux.HandleFunc("/", httpHandler)
	log.Fatal(http.ListenAndServe(addr, mux))
}

func httpHandler(w http.ResponseWriter, req *http.Request) {
	var terms []string
	for key, values := range req.URL.Query() {
		if len(values[0]) == 0 {
			terms = append(terms, key)
			continue
		}
		value := strings.Join(values, ",")
		terms = append(terms, fmt.Sprintf("%s=%v", key, value))
	}
	result, err := ch04.SearchIssues(terms)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Bad Request\n"))
		return
	}
	issueList.Execute(w, result)
}

func cmdHandler() error {
	result, err := ch04.SearchIssues(os.Args[1:])
	if err != nil {
		return err
	}
	return issueList.Execute(os.Stdout, result)
}
