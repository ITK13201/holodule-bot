#!/bin/bash

BASE_DIR=..

cd `dirname $0`

MIGRATION_DIR=$BASE_DIR/migrations
DRIVER=mysql
if [ -z "$DATABASE_URL" ]; then
    DATABASE_URL="user:password@tcp(localhost:3306)/holodule_bot?parseTime=true"
fi
GOOSE=$BASE_DIR/bin/goose

ONE_ARGS_COMMANDS=("up" "up-by-one" "down" "redo" "reset" "status" "version")
TWO_ARGS_COMMANDS=("up-to" "down-to")


if [ $# = 1 ]; then
    flag=0
    for i in "${!ONE_ARGS_COMMANDS[@]}"; do
        if [ "${ONE_ARGS_COMMANDS[$i]}" = "$1" ]; then
            $GOOSE -dir $MIGRATION_DIR $DRIVER $DATABASE_URL $1
            flag=1
        fi
    done
    if [ $flag = 0 ]; then
        echo "invalid command."
        exit 2
    fi
elif [ $# = 2 ]; then
    flag=0
    for i in "${!TWO_ARGS_COMMANDS[@]}"; do
        if [ "${TWO_ARGS_COMMANDS[$i]}" = "$1" ]; then
            $GOOSE -dir $MIGRATION_DIR $DRIVER $DATABASE_URL $1 $2
            flag=1
        fi
    done
    if [ $flag = 0 ]; then
        echo "invalid command."
        exit 2
    fi
else
    echo "invalid command."
    exit 1
fi

exit 0