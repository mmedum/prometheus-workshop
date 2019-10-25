from fastapi import FastAPI
from starlette.responses import PlainTextResponse
import prometheus_client

app = FastAPI(title="Python Service")

CONTENT_TYPE_LATEST = str('text/plain; version=0.0.4; charset=utf-8')


@app.get("/")
def read_root():
    return {"Hello": "World"}


@app.get("/metrics")
def read_metrics():
    return PlainTextResponse(prometheus_client.generate_latest())
