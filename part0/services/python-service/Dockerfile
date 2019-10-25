FROM python:3.7
LABEL maintainer="Mark Medum Bundgaard <mmedum@gmail.com>"

WORKDIR api

RUN pip install fastapi uvicorn pipenv

COPY Pipfile* ./
RUN pipenv lock --requirements > requirements.txt

RUN pip install -r requirements.txt

COPY main.py main.py

EXPOSE 80
CMD ["uvicorn", "main:app", "--host", "0.0.0.0", "--port", "80"]
