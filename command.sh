#!/bin/bash

echo 'start'
echo 'seismic wave simulation'
go run main.go
echo 'visualize'
python3 main.py
echo 'creat mp4'
cd images; ffmpeg -framerate 30 -i frame_%04d.png -c:v libx264 -pix_fmt yuv420p output.mp4
echo 'finish'