# Build-time variables
ARG RELEASE=prod
ARG ALPINE_VERSION=3.15
ARG GOLANG_VERSION=1.15
ARG NODE_VERSION=16

# backend build (api server)
FROM golang:${GOLANG_VERSION}-alpine AS api-build
RUN apk add --no-cache bash dep make git curl g++

ARG RELEASE
COPY ./api /go/src/commento/api/
WORKDIR /go/src/commento/api
RUN make ${RELEASE} -j$(($(nproc) + 1))


# frontend build (html, js, css, images)
FROM node:${NODE_VERSION}-alpine${ALPINE_VERSION} AS frontend-build
RUN apk add --no-cache bash make python2 g++

ARG RELEASE
COPY ./frontend /commento/frontend
WORKDIR /commento/frontend/
RUN make ${RELEASE} -j$(($(nproc) + 1))


# templates and db build
FROM alpine:${ALPINE_VERSION} AS templates-db-build
RUN apk add --no-cache bash make

ARG RELEASE
COPY ./templates /commento/templates
WORKDIR /commento/templates
RUN make ${RELEASE} -j$(($(nproc) + 1))

COPY ./db /commento/db
WORKDIR /commento/db
RUN make ${RELEASE} -j$(($(nproc) + 1))


# final image
FROM gcr.io/distroless/static-debian11

ARG RELEASE

COPY --from=api-build /go/src/commento/api/build/${RELEASE}/commento /commento/commento
COPY --from=frontend-build /commento/frontend/build/${RELEASE}/js /commento/js
COPY --from=frontend-build /commento/frontend/build/${RELEASE}/css /commento/css
COPY --from=frontend-build /commento/frontend/build/${RELEASE}/images /commento/images
COPY --from=frontend-build /commento/frontend/build/${RELEASE}/fonts /commento/fonts
COPY --from=frontend-build /commento/frontend/build/${RELEASE}/i18n /commento/i18n
COPY --from=frontend-build /commento/frontend/build/${RELEASE}/*.html /commento/
COPY --from=templates-db-build /commento/templates/build/${RELEASE}/templates /commento/templates/
COPY --from=templates-db-build /commento/db/build/${RELEASE}/db /commento/db/

EXPOSE 8080
WORKDIR /commento/
ENV COMMENTO_BIND_ADDRESS="0.0.0.0"
USER nobody
ENTRYPOINT ["/commento/commento"]
