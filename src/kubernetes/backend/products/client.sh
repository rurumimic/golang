#!/usr/bin/env bash

PORT=${PORT:-3002}

for i in {1..10}
do
  curl http://localhost:$PORT
done

