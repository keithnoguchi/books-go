package main

import (
	"flag"
	"html/template"
	"log"
	"os"

	"book/ch04"
)

var (
	host = flag.String("host", "", "listening host")
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
}

func cmdHandler() error {
	result, err := ch04.SearchIssues(os.Args[1:])
	if err != nil {
		return err
	}
	return issueList.Execute(os.Stdout, result)
}
