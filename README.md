### ts 视频合成命令

```code (mac)
cat *.ts > x.ts

brew install ffmpeg

ffmpeg -y -i x.ts -c:v libx264 -c:a copy -bsf:a aac_adtstoasc x.mp4
```