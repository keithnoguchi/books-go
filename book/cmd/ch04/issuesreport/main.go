// issuesreport reports the github issues with text/template package
package main

import (
	"log"
	"os"
	"text/template"
	"time"

	"book/ch04"
)

const templ = `{{.TotalCount}} issues:
{{range .Items}}-------------------------------------
Number: {{.Number}}
User:   {{.User.Login}}
Title:  {{.Title | printf "%.64s"}}
Age:    {{.CreatedAt | daysAgo}} days
{{end}}`

func main() {
	result, err := ch04.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
	report := template.Must(template.New("issuelist").
		Funcs(template.FuncMap{"daysAgo": daysAgo}).
		Parse(templ))
	if err := report.Execute(os.Stdout, result); err != nil {
		log.Fatal(err)
	}
}

func daysAgo(t time.Time) int {
	return int(time.Since(t).Hours() / 24)
}
