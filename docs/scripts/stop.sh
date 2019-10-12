#!/bin/sh
set +e

PIDS=`pgrep ^ease-gateway$`
if [ $? -ne 0 ]; then
  echo "INFO: The service Ease Gateway did not started!"
  exit 0
fi

echo -e "Stopping the service Ease Gateway ...\c"
for PID in $PIDS ; do
  kill $PID > /dev/null 2>&1
done

while [ true ]; do
  echo -e ".\c"

  IDS=`pgrep ^ease-gateway$`
  if [ $? -ne 0 ]; then
    echo
    echo "PID: $PIDS"
    echo "Service Ease Gateway stopped."
    exit 0
  fi

  sleep 1
done
