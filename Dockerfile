FROM alpine:latest as tini-stage
RUN apk update && apk add --no-cache tini

FROM neosmemo/memos:stable

COPY --from=tini-stage /sbin/tini /sbin/tini

EXPOSE 5230
VOLUME ["/var/opt/memos"]

COPY run-memos.sh /run-memos.sh
ENTRYPOINT ["/sbin/tini", "--", "/run-memos.sh"]
