# Local Prometheus and Grafana

(Remember to build [go-service](../../go-service) and
[dotnet-service](../../dotnet-service))

To start the services

```
docker-compose up
```

By doing so, it is possible to access the services through a simple webbrowser.

| Service | URL |
| --- | --- |
| Go-Service | [Metrics Url](http://localhost:8080/metrics) |
| Promtheus | [Prometheus](http://localhost:9090) |
| Grafana | [Grafana](http://localhost:3000) |

When navigating to `Grafana` for the first time, please use admin/admin for
username and password, remember when creating a new password, that the instance
is only running locally, so choose an easy password.

From the `Grafana` startpage, choose `Add data source`, search for `Prometheus`
and set the `URL: http://prometheus:9090` thereby press `Save & Test` and
hopefully a green bar with the text `Data source is working` shows up.
