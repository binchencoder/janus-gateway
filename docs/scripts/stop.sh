#!/bin/sh
set +e

PIDS=`pgrep ^janus-gateway$`
if [ $? -ne 0 ]; then
  echo "INFO: The service Janus Gateway did not started!"
  exit 0
fi

echo -e "Stopping the service Janus Gateway ...\c"
for PID in $PIDS ; do
  kill $PID > /dev/null 2>&1
done

while [ true ]; do
  echo -e ".\c"

  IDS=`pgrep ^janus-gateway$`
  if [ $? -ne 0 ]; then
    echo
    echo "PID: $PIDS"
    echo "Service Janus Gateway stopped."
    exit 0
  fi

  sleep 1
done
