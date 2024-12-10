---
page_title: "{{.Name}} {{.Type}} - {{.ProviderName}}"
subcategory: "Data Warehouse"
description: |-
{{ .Description | plainmarkdown | trimspace | prefixlines "  " }}
---

# {{.Name}} ({{.Type}})

{{ .Description | trimspace }}

{{ if .HasExample -}}
## Example Usage

{{ tffile .ExampleFile }}
{{- end }}

{{ .SchemaMarkdown | trimspace }}

{{- if .HasImport }}
## Import

Import is supported using the following syntax:

{{codefile "shell" .ImportFile }}
{{- end }}
