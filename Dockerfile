FROM neosmemo/memos:stable

FROM alpine:latest as tini-bash
RUN apk add --no-cache tini bash

COPY --from=tini-bash /sbin/tini /sbin/tini
COPY --from=tini-bash /bin/bash /bin/bash

EXPOSE 5230
VOLUME ["/var/opt/memos"]

COPY run-memos.sh /run-memos.sh
RUN chmod +x /run-memos.sh

ENTRYPOINT ["/sbin/tini", "--", "/bin/bash", "/run-memos.sh"]
