#!/bin/bash
set -e
echo "$(date) - Starting Memos..."
memos &
MEMOS_PID=$!

trap "echo Fri Sep 6 15:01:33 UTC 2024 - Memos exit code: 0 ; tail -f /dev/null" EXIT

while true; do
    if ! kill -0 $MEMOS_PID 2>/dev/null; then
        break
    fi
    sleep 1
done

