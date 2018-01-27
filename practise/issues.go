package main

import (
	"fmt"
	"log"
	"os"
	"html/template"
	"plum/practise/github_issues"
	"time"
)

//var report = template.Must(template.New("issuslist").Funcs(template.FuncMap{"daysAgo": daysAgo}).Parse(github_issues.Templ)) // 普通文本模板
// html 模版
var report = template.Must(template.New("issuelist").Parse(`
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

func daysAgo(t time.Time) int {
	return int(time.Since(t).Hours() / 24)
}

func printTest1(result *github_issues.IssusesSearchResult) {
	// go run issues.go python  > test.txt
	fmt.Printf("%d issues: \n", result.TotalCount)
	for _, item := range result.Items {
		fmt.Printf("#%-5d %9.9s %.55s\n", item.Number, item.User.Login, item.Title)
	}
}

func printTest2(result *github_issues.IssusesSearchResult) {
	// go run issues.go python  > test.html
	if err := report.Execute(os.Stdout, result); err != nil {
		log.Fatal(err)
	}
}
func main() {
	result, err := github_issues.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
	//printTest1(result) // 普通文本输出
	printTest2(result) // html格式输出

	const templ = `<p>A: {{.A}}</p><p>B: {{.B}}</p>`
	t := template.Must(template.New("escape").Parse(templ))
	var data struct {
		A string        // untrusted plain text
		B template.HTML // trusted HTML
	}
	data.A = "<b>Hello!</b>"
	data.B = "<b>Hello!</b>"
	if err := t.Execute(os.Stdout, data); err != nil {
		log.Fatal(err)
	}
}
