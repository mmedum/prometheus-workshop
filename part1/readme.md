# Prometheus

The core of Prometheus is a timeseries DB and query language, which can be
extended through a modular approach (Exporters, Dashboard, Alertmanager)

Prometheus uses a white box monitoring approach, where applications is aware of
or rather instrumented to be scraped by Prometheus.

## Note

Before really starting any of the work, remember to complete [Part 0](../part0)

## Prometheus UI

The running `Prometheus` is located at `localhost:9090`

When the local instance is running, navigate to the Prometheus UI where you
should be presented with the following tabs:

- `Alerts`, which list all alerting rules and the current status for each alert
- `Graph`, where it is possible to query the timeseries data
- `Status`, which represents the current status of scraping targets.
- `Help`, which links to the Prometheus documentation

Navigate to `Status -> Targets` and look at the four targets, which hopefully
have a status of `UP`. Every target which Prometheus should scrape is defined in
`prometheus.yml`, which for this exercise is created manually.

```yaml
global:
  scrape_interval:     15s # Set the scrape interval to every 15 seconds.
  evaluation_interval: 15s # Evaluate rules every 15 seconds.
  # scrape_timeout is set to the global default (10s).

  # Attach these labels to any time series or alerts when communicating with
  # external systems (federation, remote storage, Alertmanager).
  external_labels:
      monitor: 'Alertmanager'

rule_files:

scrape_configs:
  - job_name: 'prometheus'

    static_configs:
      - targets: ['localhost:9090']

  - job_name: 'docker'
         # metrics_path defaults to '/metrics'
         # scheme defaults to 'http'.

    # docker.for.mac.localhost:9090 -> MAC
    # docker.for.win.localhost:9090 -> WIN
    # localhost -> Linux

    static_configs:
      - targets: ['docker.for.mac.localhost:9090']

  - job_name: 'go-service'

    static_configs:
      - targets: ['go-service']

  - job_name: 'dotnet-service'

    static_configs:
      - targets: ['dotnet-service']
```

Prometheus assumes that every target has an URL with a `/metrics` path defined,
so Prometheus can pull/scrape the metrics.

By navigating to `Graph` and query `up`, it should be possible to see the
targets, which is being scraped currently.

### Prometheus Exercise

- Add the `python-service` as a new scrape target in the `prometheus.yml` file
  and start all the services again.
- Add another `dotnet-service` and `go-service` to `docker-compose.yml` and
  extend `prometheus.yml` to scrape the new targets. Remember that since every
  service is running in docker, it is possible to use the name of the service,
  which is defined in `docker-compose.yml`, to point Prometheus to the new scrape
  targets.
- Navigate to the Prometheus UI and locate the two new targets, check whether it
  is possible to see the `UP` state, when querying through `Graph`.

## Grafana

This workshop uses Grafana as the visualization tool for outputting and querying
Prometheus metrics. For getting some quick information, navigate to `Dashboards`
on the left side, press `Manage` and then `Import`. It is possible to import
third party dashboard, where most of them can be found at [Grafana
Labs](https://grafana.com/dashboards), in this workshop it is recommended using
`6671` for getting all `go` metrics.

### Grafana Exercise

- Add another panel to the current dashboard, this should visualize the `up`
  parameter from Prometheus.
