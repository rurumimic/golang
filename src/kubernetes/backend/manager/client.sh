#!/usr/bin/env bash

PORT=${PORT:-3000}

for i in {1..1000}
do
  curl http://localhost:$PORT
done
