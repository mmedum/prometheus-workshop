# Instrumenting code

Now that both Prometheus and Grafana is running, it is time to instrument code
to output some interesting custom metrics. By using the standard Prometheus
client for both go and dotnet, both services are already outputting quite a few
metrics.

The `go-service` has defined a `/v1/ping` endpoint, when called with a `GET`
request, returns json and a random status code. The idea is to create a counter
metric, which just counts the number of calls to the `/v1/ping` endpoint.

In the `go-service` folder, locate the file `main.go` and open it. The code for
responding to a `GET` call on `/v1/ping` is as follows.

```go
...
func pong(w http.ResponseWriter, r *http.Request) {
	response := make(map[string]string)
	response["message"] = "pong"

	rand.Seed(time.Now().Unix())

	responseCode := responseCodes[rand.Intn(len(responseCodes))]

	render.Status(r, responseCode)

	render.JSON(w, r, response)
}
...

```

First it will be needed to define the Prometheus variable, which should store
the total number of requests.

```go
var (
	pongCount = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "ping_total_number_of_requests",
			Help: "Number of ping requests.",
		},
		[]string{"status"},
	)
)
```

The `pongCount` variable is defined with a parameter `status`, such that every
single response code can be isolated and summarized also.

And a `init` function should be created, so the variable is registered to
prometheus on startup.

```go
func initPrometheusMetric() {
	prometheus.MustRegister(pongCount)
}
```

This function should be called in the `main` function, such that prometheus
tracks the new variable.

```go
func main() {
	initResponseCodes()

    initPrometheusMetric()

	router := Routes()
	port := 80
}
```

With that, it is needed to instrument the `pong` function, so that every single
response increments the counter with the return statuscode as key.

```go
func pong(w http.ResponseWriter, r *http.Request) {
...

	responseCode := responseCodes[rand.Intn(len(responseCodes))]

	pongCount.WithLabelValues(strconv.Itoa(responseCode)).Inc()

	render.Status(r, responseCode)
...
```

With that code added, build a new image of the `go-service` and run
`docker-compose up` again, this should give extra metrics paths, but before
checking it would be a good idea to generate some traffic, by doing `GET`
requests on `localhost:8081/v1/ping`.

After a couple of hits, look at the metrics of the go service by navigating to
[localhost:8080/metrics](localhost:8080/metrics), and search for the new
variable `ping_total_number_of_requests`, this should output something like
this.

```
# HELP ping_total_number_of_requests Number of ping requests.
# TYPE ping_total_number_of_requests counter
ping_total_number_of_requests{status="500"} 6
ping_total_number_of_requests{status="503"} 1
```

## Exercise

- Prometheus supports many different metrics types, look at the `go-service` and
  `dotnet-service` and create a couple of new metrics. The different types can
  be found [here](https://prometheus.io/docs/concepts/metric_types/). Be very
  creative! 
