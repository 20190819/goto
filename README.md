### ts 视频合成命令

```code (mac)
cat *.ts > x.ts

brew install ffmpeg

ffmpeg -y -i x.ts -c:v libx264 -c:a copy -bsf:a aac_adtstoasc x.mp4
```

```code 
# 直接下载
ffmpeg -i "http://xxxx.com/file_name.m3u8"  "save_video.mp4" 
```