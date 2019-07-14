FROM scratch

ARG CREATED
ARG REVISION
ARG VERSION

LABEL org.opencontainers.image.created=$CREATED
LABEL org.opencontainers.image.url="https://hub.docker.com/r/karljohangrahn/microtest"
LABEL org.opencontainers.image.source="https://github.com/karl-johan-grahn/microtest"
LABEL org.opencontainers.image.revision=$REVISION
LABEL org.opencontainers.image.version=$VERSION

COPY microtest /
COPY handlers/openapi.yaml /
CMD ["/microtest"]
