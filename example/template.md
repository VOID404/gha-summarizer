# Sample template

Generated from `{{ env "$GITHUB_REPOSITORY" }}`

{{ if filep "./example/code.go" -}}
```go
{{ file "./example/code.go"}}
```
{{- end }}
{{if filep "./example/other_code.go" }}
```go
{{ file "./example/other_code.go" }}
```
{{- end -}}
