
# dotnet-service

## How to build Docker image

 ```bash
 docker build . -t dotnet-service
 ```

## How to run

Go to docker-files folder and run

```bash
docker-compose up
```

Open localhost:8080/swagger to see swagger endpoints.
Localhost:8080/metrics to see exported metrics

localhost:9090 to go to prometheus UI
