FROM ubuntu:latest
LABEL authors="reosh"

ENTRYPOINT ["top", "-b"]