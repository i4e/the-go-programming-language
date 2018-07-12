#!/bin/bash

curl "http://localhost:8000/?cycles=64&size=200&nframes=64&res=0.001&delay=1" > out0.gif
curl "http://localhost:8000/?cycles=100&res=0.001&nframes=64&delay=0.1" > out1.gif
curl "http://localhost:8000/?cycles=128&res=0.0001&nframes=64&delay=1" > out2.gif