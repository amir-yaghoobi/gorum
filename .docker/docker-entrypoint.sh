#!/bin/sh
set -e

if [[ "$1" = "server" ]]; then
    exec "$@"
else
    exec server "$@"
fi
