FROM alpine:latest as tini-stage
RUN apk update && apk add --no-cache tini bash

FROM neosmemo/memos:stable

COPY --from=tini-stage /sbin/tini /sbin/tini
COPY --from=tini-stage /bin/bash /bin/bash

EXPOSE 5230
VOLUME ["/var/opt/memos"]

COPY run-memos.sh /run-memos.sh
RUN chmod +x /run-memos.sh

ENTRYPOINT ["/sbin/tini", "--", "/bin/bash", "/run-memos.sh"]
