#! /usr/bin/env bash
CURDIR=$(cd $(dirname $0); pwd)
echo "$CURDIR/bin/hellotest"
exec "$CURDIR/bin/hellotest"