FROM ubuntu:22.04

WORKDIR /app

COPY ./bin ./

COPY ./web/dist ./wwwroot

ENV ADMIN_PASSWORD ""
ENV DATABASE_URL ""

ENTRYPOINT ["./canarails"]
