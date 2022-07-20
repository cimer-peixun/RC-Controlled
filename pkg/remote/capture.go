package remote

import (
	"fmt"
	"image"
	"image/png"
	"os"

	"github.com/kbinani/screenshot"
)

func save(img *image.RGBA, filePath string) {
	file, err := os.Create(filePath)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	png.Encode(file, img)
}

func CaptureFullScreen() {
	img, err := screenshot.CaptureDisplay(0)
	if err != nil {
		panic(err)
	}
	save(img, "自定义.png")
}

func main() {

	//自定义截图
	img, err := screenshot.Capture(0, 0, 500, 500)
	if err != nil {
		panic(err)
	}
	save(img, "自定义.png")

	//获取所有活动屏幕
	n := screenshot.NumActiveDisplays()
	if n <= 0 {
		panic("没有发现活动的显示器")
	}

	//全屏截取所有活动屏幕
	if n > 0 {
		for i := 0; i < n; i++ {
			img, err = screenshot.CaptureDisplay(i)
			if err != nil {
				panic(err)
			}
			fileName := fmt.Sprintf("第%d屏幕截图.png", i)
			save(img, fileName)
		}
	}

	//使用 Rectangle 自定义截图
	// 获取第一个屏幕显示范围
	var new image.Rectangle = image.Rect(0, 0, 600, 700)
	img, err = screenshot.CaptureRect(new)
	if err != nil {
		panic(err)
	}
	save(img, "new.png")

	//使用 GetDisplayBounds获取指定屏幕显示范围，全屏截图
	bounds := screenshot.GetDisplayBounds(0)
	img, err = screenshot.CaptureRect(bounds)
	if err != nil {
		panic(err)
	}
	save(img, "all.png")

	//使用Union获取指定屏幕 矩形最小点和最大点
	var all image.Rectangle = image.Rect(0, 0, 0, 0)
	bounds1 := screenshot.GetDisplayBounds(0)
	all = bounds1.Union(all)
	fmt.Println(all.Min.X, all.Min.Y, all.Dx(), all.Dy())

}
