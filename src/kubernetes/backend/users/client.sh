#!/usr/bin/env bash

PORT=${PORT:-3001}

for i in {1..10}
do
  curl http://localhost:$PORT
done
