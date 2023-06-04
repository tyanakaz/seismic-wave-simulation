#!/bin/bash

echo 'start'
go run main.go
echo 'go'
python3 main.py
echo 'python'
cd images; ffmpeg -framerate 30 -i frame_%04d.png -c:v libx264 -pix_fmt yuv420p output.mp4
echo 'finish'