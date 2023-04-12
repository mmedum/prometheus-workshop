FROM python:3.12
LABEL maintainer="Mark Medum Bundgaard <mmedum@gmail.com>"

WORKDIR api

COPY requirements.txt ./

RUN pip install --no-cache-dir -r requirements.txt

COPY main.py main.py

EXPOSE 80
CMD ["uvicorn", "main:app", "--host", "0.0.0.0", "--port", "80"]
