#!/bin/sh
xk6 build v1.5.0\
    --with xk6-k6-cache=.

./k6 run k6caches.js