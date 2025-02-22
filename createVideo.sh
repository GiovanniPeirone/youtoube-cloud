ffmpeg -r 30 -i frames/frame_%03d.png -vf "format=yuv420p" output.mp4
