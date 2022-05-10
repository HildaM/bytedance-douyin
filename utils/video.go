package utils

import (
	"bytes"
	"fmt"
	ffmpeg "github.com/u2takey/ffmpeg-go"
	"io"
	"os"
	"strconv"
	"time"
)

// use ffmpeg read first frame of video as a jpeg

func ReadFrameAsJpeg(inFileName string) io.Reader {
	frameNum := 1
	buf := bytes.NewBuffer(nil)
	err := ffmpeg.Input(inFileName).
		Filter("select", ffmpeg.Args{fmt.Sprintf("gte(n,%d)", frameNum)}).
		Output("pipe:", ffmpeg.KwArgs{"vframes": 1, "format": "image2", "vcodec": "mjpeg"}).
		WithOutput(buf, os.Stdout).
		Run()
	if err != nil {
		panic(err)
	}
	return buf
}

// depend on some algorithm to generate a name for uploaded videos

func GenerateFilename(username string, userId int64) string {
	timestamp := strconv.FormatInt(time.Now().Unix(), 10)
	return timestamp + "_" + username + "_" + strconv.FormatInt(userId, 10)
}
