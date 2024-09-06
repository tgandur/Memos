FROM neosmemo/memos:stable

RUN apk update && apk add --no-cache tini

EXPOSE 5230

VOLUME ["/var/opt/memos"]

ENTRYPOINT ["/sbin/tini", "--"]

CMD ["memos"]
