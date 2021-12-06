# Prometheus and Grafana

Before diving into how `Prometheus` is scraping targets, we will need to have a
couple of services running, such that we have `targets` to scrape and `grafana`
to visualize the scraped data.

This project uses `docker` to spin up the full `production` environment, so
start by opening `docker-compose.yml` and have a look at each of the services.

- `go-service` is a small service, which can receive requests
- `dotnet-service` is another small service, which can receive requets
- `http-simulator` is program for simulating `GET` request to the different
  services
- `prometheus` is the prometheus instance, which uses the config saved in the
  `config` folder
- `grafana` is the grafana instance, which will visualize the scraped data from
  `prometheus`

For those who have not used `docker` before, it is recommended to read the
`docker` [tutorial](https://docker-curriculum.com/), but note for this project
it should be enough to only use a couple of commands.

## Prometheus config

For making certain that `prometheus` can scrape the correct docker installation,
it is needed to set the `docker` scrape target in the
[prometheus.yml](configs/prometheus.yml)

```yaml
  - job_name: 'docker'
         # metrics_path defaults to '/metrics'
         # scheme defaults to 'http'.

    # docker.for.mac.localhost:9090 -> MAC
    # docker.for.win.localhost:9090 -> WIN
    # localhost -> Linux

    static_configs:
      - targets: ['HERE']
```

Such that `prometheus` is scraping the health of the running `docker engine`.

## Quick docker tutorial

Open the terminal and navigate to the `part0` folder, in this folder there
should be a `docker-compose.yml` located, which is the `yml` file that defines
how the services should be started, and in our case connected.

### Commands

The following three commands should be enough for setting up the environment.
When running `docker-compose up` the first time, docker will `build`, `tag`,
`download` and `run` each service and place them in the same network called
`public`, which will make it easy for referring to services later.

#### Start

For `starting` the services in the foreground

```
docker-compose up
```

For `starting` the services in the background

```
docker-compose up -d
```

#### Rebuild

For `rebuilding` the services after a code change

```
docker-compose up --build
```

#### Stop

For `stopping` the services fully.

```
docker-compose down
```

## Run it

After `docker-compose up` is done, it should be possible to navigate to each
service using these urls.

| Service | URL |
| --- | --- |
| Go-Service | [Metrics Url](http://localhost:8080/metrics) |
| Promtheus | [Prometheus](http://localhost:9090) |
| Grafana | [Grafana](http://localhost:3000) |

### Grafana

When navigating to `Grafana` for the first time, please use `admin` for both
username and password, remember when creating a new password, that the instance
is only running locally, so choose an easy password like for instance `admin`.

From the `Grafana` startpage, choose `Add data source`, pick `Prometheus`
and set the `URL` to `http://prometheus:9090`, press `Save & Test` and
hopefully a green bar with the text `Data source is working` shows up and with
that the local deployment is done.

## Conclusion

This should conclude the setup for the `production` environment. Remember to
keep the containers running and restart them with the `docker-compose` commands
when changing code or config files.
