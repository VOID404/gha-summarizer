# GHA Summarizer [![GHA Summarizer](https://github.com/VOID404/gha-summarizer/actions/workflows/gha-summarizer.yaml/badge.svg)](https://github.com/VOID404/gha-summarizer/actions/workflows/gha-summarizer.yaml)

Simple action for Go template based summarization of GitHub Actions workflow runs.
It supports usual Go templating features, check [text/template](https://golang.org/pkg/text/template/) for more details.

## Usage
```yaml
- name: Render template
  uses: VOID404/gha-summarizer@1.0.0
  with:
    template: hack/summary_template.md
    out-file: summary.md
- name: Set step output
  run: cat summary.md >> $GITHUB_STEP_SUMMARY
```

## Functions
Additionally some helper functions are provided:

### `file`

Embedd file content in the template.
````md
# Sample template
```go
{{ file "./example/code.go"}}
```
````
---

### `env`

Get environment variable value.
This has access to github context, so you can use it to get values from the workflow run.
````md
```md
# Sample template
Generated from `{{ env "$GITHUB_REPOSITORY" }}`
```
````
---
