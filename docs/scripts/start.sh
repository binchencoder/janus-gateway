#!/bin/sh
# resolve links - $0 may be a softlink
PRG="$0"

while [ -h "$PRG" ] ; do
  ls=`ls -ld "$PRG"`
  link=`expr "$ls" : '.*-> \(.*\)$'`
  if expr "$link" : '/.*' > /dev/null; then
    PRG="$link"
  else
    PRG=`dirname "$PRG"`/"$link"
  fi
done

PRGDIR=`dirname "$PRG"`
cd $PRGDIR/..
HOMEDIR=`pwd`
chmod +x $HOMEDIR/bin/*

if [ -z $SERVICE_LOG_DIR ]; then
  SERVICE_LOG_DIR=$HOMEDIR/logs
fi

if [ -z $SKYLB_ENDPOINTS ]; then
  SKYLB_ENDPOINTS=skylbserver:1900
fi


STDOUT_FILE=$HOMEDIR/std.out

set +e
PIDS=`pgrep ^janus-gateway$`
if [ $? -eq 0 ]; then
  echo "ERROR: The service Janus Gateway already started!"
  echo "PID: $PIDS"
  exit 1
fi
set -e

echo "Service Janus Gateway start..."
nohup $HOMEDIR/bin/janus-gateway \
-debug-svc-endpoint=vexillary-service=192.168.10.41:4100 \
-debug-svc-endpoint=pay-grpc-service=192.168.32.18:10008 \
-skylb-endpoints=$SKYLB_ENDPOINTS \
-log_dir=$SERVICE_LOG_DIR \
-g_log_dir=$SERVICE_LOG_DIR \
-roll_type=date \
-config-profiles=gateway \
> $STDOUT_FILE 2>&1 &

sleep 2

PIDS_COUNT=`pgrep ^janus-gateway$|wc -l`
if [ $PIDS_COUNT -eq 0 ]; then
  echo "ERROR: The service Janus Gateway does not started!"
  exit 1
fi

echo "Service Janus Gateway start success!"
