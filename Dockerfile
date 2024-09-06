FROM neosmemo/memos:stable

EXPOSE 5230

VOLUME ["/var/opt/memos"]

ENTRYPOINT ["/sbin/tini", "--", "memos"]
