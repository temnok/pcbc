#!/bin/bash
set -e
cd $(dirname $0)

grep --files-without-match \
  --recursive \
  --include=\*.go \
  --regexp='Copyright'
