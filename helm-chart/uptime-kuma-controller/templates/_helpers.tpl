{{/*
Return the name of the service account to use
*/}}
{{- define "uptime-kuma-controller.serviceAccountName" -}}
{{- if .Values.serviceAccount.name }}
{{- .Values.serviceAccount.name }}
{{- else }}
{{- include "uptime-kuma-controller.fullname" . }}
{{- end }}
{{- end }}

{{/*
Return the fullname of the release
*/}}
{{- define "uptime-kuma-controller.fullname" -}}
{{- printf "%s-%s" .Release.Name "uptime-kuma-controller" | trunc 63 | trimSuffix "-" }}
{{- end }}

{{/*
Return the chart name
*/}}
{{- define "uptime-kuma-controller.name" -}}
uptime-kuma-controller
{{- end }}