#!/bin/sh
# wait_for.sh

# Ожидание доступности указанного сервиса
set -e

host="$1"
shift
cmd="$@"

until nc -z "$host"; do
  >&2 echo "Service is unavailable - waiting..."
  sleep 1
done

>&2 echo "Service is up - executing command"
exec $cmd
