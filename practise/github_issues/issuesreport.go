package github_issues


const Templ = `{{ .TotalCount }} issues:
{{range .Items }}-------------------------------------
Number: {{ .Number }}
User: {{ .User.Login }}
Title: {{ .Title | printf "%.64s" }}
Age: {{ .CreatedAt | daysAgo }} days
{{ end }}`

// | 操作符表示将前一个表达式的结果作为后一个函数的输入
