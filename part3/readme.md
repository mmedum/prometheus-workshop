# Grafana Dashboard

Navigate to running Grafana instance and create a new dashboard

## Request rate

Add a new panel with a graph, click on `general` and rename the panel to
`Request Rate`. Navigate to the `Queries` and add

```bash
sum(rate(promhttp_metric_handler_requests_total[5m]))
```

This should output a graph, with the current request rate for `/metrics` on one
of the `go-services`. It is possible to change to query to match a specific app

```bash
sum(rate(promhttp_metric_handler_requests_total{app="go-service"}[5m]))
```

This means that it should be a specific `go-service`, remember the name of
the `app` is defined in the `prometheus.yml`.

## Error rate

So the query language of Prometheus is quite good, so it is possible to create
more complex queries. Start by creating a new row and add a new panel with a
`singlestat`. Add the query

```bas
sum(rate(ping_total_number_of_requests{status="500"}[5m])) / sum(rate(ping_total_number_of_requests[5m]))
```

This should give an indication of how many responses have a statuscode of `500`.

Navigate to Visualization and change the unit to percent, this will translate
to a more direct percentage. On top of that we should see the value as the most
current value, which means changing Stat to Current.

## Exercise

- Be creative and create even more metrics based on the data returned from
  services
