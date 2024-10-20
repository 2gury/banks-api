FROM ubuntu:22.04

ADD ./bin/app /app
ADD ./config.yaml /config.yaml
ADD ./translate-key.json /translate-key.json

RUN apt-get update && apt-get install -y ca-certificates

CMD ["/bin/bash", "-c", "/app 2>&1 >> /data/logs.txt"]
