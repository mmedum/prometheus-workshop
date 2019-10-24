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

## Quick docker tutorial

Open the terminal and navigate to the `part0` folder, in this folder there
should be a `docker-compose.yml` located, which is the `yml` file that defines
how the services should be started, and in our case connected.

### Commands

The three following commands should be enough for setting up the environment.
When running `docker-compose up` the first time, docker will `build`, `tag`,
`download` and `run` each service and place them in the same network called
`public`, which will make it easy for referring services later.

#### Start

For `starting` the services

```
docker-compose up
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
service on these urls.

| Service | URL |
| --- | --- |
| Go-Service | [Metrics Url](http://localhost:8080/metrics) |
| Promtheus | [Prometheus](http://localhost:9090) |
| Grafana | [Grafana](http://localhost:3000) |

When navigating to `Grafana` for the first time, please use admin/admin for
username and password, remember when creating a new password, that the instance
is only running locally, so choose an easy password like for instance `admin`.

From the `Grafana` startpage, choose `Add data source`, search for `Prometheus`
and set the `URL` to `http://prometheus:9090`, press `Save & Test` and
hopefully a green bar with the text `Data source is working` shows up and with
that the local deployment is done.
