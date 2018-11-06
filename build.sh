#!/bin/bash

if [ ! -e "cucumber" ]; then
  mkdir cucumber
else
    if [ -e cucumber/report.json ]; then
        rm cucumber/report.json
    fi
fi

if [ $# -lt 1 ]; then
   godog -f cucumber > cucumber/report.json $*
   exit 0
fi

if [ $1 = "-h" ]; then
    godog -h
    exit 0
fi

godog -f cucumber > cucumber/report.json $*
