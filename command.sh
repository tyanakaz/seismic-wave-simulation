#!/bin/bash

echo 'start'
echo 'seismic wave simulation'
go run simple_wave_and_absorption_boundary_simulation.go
# go run sh_wave_simulation.go
# go run simple_wave_simulation.go
# go run sv_wave_simulation.go
echo 'visualize'
python3 main.py
echo 'creat mp4'
cd images; ffmpeg -framerate 30 -i frame_%04d.png -c:v libx264 -pix_fmt yuv420p output.mp4
echo 'finish'