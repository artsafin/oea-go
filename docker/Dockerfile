FROM scratch

ADD ca-certificates.crt /etc/ssl/certs/
ADD passwd /etc/passwd

RUN mkdir -p /certs && \
    chown appuser:appuser /certs

ENTRYPOINT ["/oea-go"]

ADD oea-go /

USER appuser
