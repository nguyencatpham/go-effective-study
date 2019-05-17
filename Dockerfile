FROM alpine:latest

ARG APP_NAME
ENV APP_NAME ${APP_NAME}

RUN apk --no-cache add curl  \
	ca-certificates \
	;
RUN apk add --no-cache tzdata && \
	 rm -rf /var/cache/apk/* /tmp/*
# ENV GOPATH /onsky/golang/
# ENV PATH $GOPATH/bin:$PATH

WORKDIR /onsky/apps/
# RUN mkdir plugins
COPY ${APP_NAME} .
# COPY config ./config
# COPY plugins/alerts ./plugins/alerts
# COPY plugins ./plugins
# COPY assets/ ./assets
COPY ./cmd/api/conf.local.yaml ./cmd/api/conf.local.yaml

EXPOSE 8080 10001

ENTRYPOINT [ "./$APP_NAME" ]

# CMD ["sh", "-c","./$APP_NAME --server_name=$APP_NAME --server_address=0.0.0.0:8080 --broker_address=0.0.0.0:10001"]