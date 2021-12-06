FROM python:3.10
LABEL maintainer="Mark Medum Bundgaard <mmedum@gmail.com>"

WORKDIR api

COPY requirements.txt ./

RUN pip install -r requirements.txt

COPY main.py main.py

EXPOSE 80
CMD ["uvicorn", "main:app", "--host", "0.0.0.0", "--port", "80"]
