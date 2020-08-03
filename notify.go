package main

import (
	"fmt"
	"os"
	"runtime"
	"strconv"
	"strings"
	"time"

	"github.com/0xAX/notificator"
)

var notify *notificator.Notificator

func main() {
	goos := runtime.GOOS

	for {
		time.Sleep(30 * time.Minute)
		now := time.Now().Format("15:05")
		nowPrev, _ := strconv.Atoi(strings.Split(now, ":")[0])
		nowNex, _ := strconv.Atoi(strings.Split(now, ":")[1])
		fmt.Println(goos)
		if nowPrev == 17 && nowNex >= 30 {
			continue
		} else if nowPrev > 17 {
			continue
		} else if nowPrev == 8 && nowNex <= 30 {
			continue
		} else if nowPrev < 8 {
			continue
		}

		notify = notificator.New(notificator.Options{
			DefaultIcon: "icon/default.png",
			AppName:     "my test app",
		})
		pngPath, _ := os.Getwd()
		pngPath = pngPath + "/appointment-soon.png"
		notify.Push("现在时间"+now+"，休息一下吧", "", pngPath, notificator.UR_CRITICAL)
	}
}
