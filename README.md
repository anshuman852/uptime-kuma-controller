# Uptime Kuma Controller

Uptime Kuma Controller is a Kubernetes operator that manages Uptime Kuma monitors and notification channels using Custom Resources. It enables you to declaratively define and automate uptime monitoring for your services within your Kubernetes cluster.

## Building the Go Controller

1. Ensure you have [Go](https://golang.org/dl/) installed (version 1.19+ recommended).
2. Navigate to the `controller` directory:
   ```sh
   cd controller
   ```
3. Build the controller binary:
   ```sh
   go build -o uptime-kuma-controller main.go
   ```

## Deploying with Helm

The recommended way to deploy the controller is using the provided Helm chart.

1. Ensure you have [Helm](https://helm.sh/) installed.
2. Package or reference the chart in `helm-chart/uptime-kuma-controller`.
3. Install the chart into your Kubernetes cluster:
   ```sh
   helm install uptime-kuma-controller ./helm-chart/uptime-kuma-controller
   ```
   You can customize values using the `--set` flag or by editing `values.yaml`.

---

## Uptime Kuma Custom Resources

### UptimeKumaInstance

Represents a connection to an Uptime Kuma server. Other resources reference this to associate monitors and notification channels with a specific Uptime Kuma instance.

**Example:**
```yaml
apiVersion: uptimekuma.example.com/v1alpha1
kind: UptimeKumaInstance
metadata:
  name: kuma-main
spec:
  url: https://uptime-kuma.example.com
  username: admin
  passwordSecretRef: kuma-admin-secret
```

### Monitor

Defines a monitor in Uptime Kuma. Each Monitor references an UptimeKumaInstance.

**Example:**
```yaml
apiVersion: uptimekuma.example.com/v1alpha1
kind: Monitor
metadata:
  name: example-monitor
spec:
  url: https://example.com
  type: http
  interval: 60
  uptimeKumaInstanceRef: kuma-main
```

### NotificationChannel

Defines a notification channel in Uptime Kuma. Each NotificationChannel references an UptimeKumaInstance.

**Example:**
```yaml
apiVersion: uptimekuma.example.com/v1alpha1
kind: NotificationChannel
metadata:
  name: slack-channel
spec:
  type: slack
  name: Slack Alerts
  data:
    webhookUrl: https://hooks.slack.com/services/XXX/YYY/ZZZ
  uptimeKumaInstanceRef: kuma-main
```

---

Apply your CRD manifests using:
```sh
kubectl apply -f your-resource.yaml
