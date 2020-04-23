FROM scratch
ADD docker/ca-certificates.crt /etc/ssl/certs/
ADD docker/docker-entrypoint.sh /
ADD oea-go /

ENTRYPOINT ["docker-entrypoint.sh"]

CMD ["/oea-go"]
