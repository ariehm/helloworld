FROM scratch
LABEL org.opencontainers.image.source https://github.com/ariehm/helloworld
ENTRYPOINT ["/sbin/hw"]
COPY hw /sbin/
