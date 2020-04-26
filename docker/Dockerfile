FROM scratch

ADD ca-certificates.crt /etc/ssl/certs/
ADD passwd /etc/passwd

ENTRYPOINT ["/oea-go"]

ADD oea-go /

USER appuser
