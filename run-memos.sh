#!/bin/sh
echo "Starting Memos..."
memos
EXIT_CODE=$?
echo "Memos exited with code ${EXIT_CODE}"
sleep 1000
exit ${EXIT_CODE}
