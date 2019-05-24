#!/bin/sh
set -e

if [[ "$1" = "server" ]] || [[ "$1" = "migrate" ]]; then
    exec "$@"
else
    exec server "$@"
fi
