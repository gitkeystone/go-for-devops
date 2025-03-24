// lissajous 产生随机利萨茹图形的GIF动画
package main

import (
	"image"
	"image/color"
	"image/gif"
	"io"
	"log"
	"math"
	"math/rand"
	"net/http"
	"os"
	"strconv"
	"strings"
)

// var palette = []color.Color{color.White, color.Black}
// var palette = []color.Color{color.Black, color.RGBA{0x00, 0xFF, 0x00, 0xFF}}
var red = color.RGBA{0xFF, 0x00, 0x00, 0xFF}
var green = color.RGBA{0x00, 0xFF, 0x00, 0xFF}
var blue = color.RGBA{0x00, 0x00, 0xFF, 0xFF}

var palette = []color.Color{color.Black, red, green, blue}

const (
	whiteIndex = 0 // 画板中的第一种颜色
	blackIndex = 1 // 画板中的下一种颜色
)

var colorIndex byte
var cycles = 5 // 完整的X振荡器变化的个数

func main() {
	if len(os.Args) > 1 && os.Args[1] == "web" {
		handler := func(w http.ResponseWriter, r *http.Request) {
			colorIndex = byte(rand.Intn(3) + 1)
			if err := r.ParseForm(); err != nil {
				log.Print(err)
			}

			for k, v := range r.Form {
				if k == "cycles" {
					var err error
					cycles, err = strconv.Atoi(strings.Join(v, ""))
					if err != nil {
						cycles = 5
						log.Print(err)
					}
				}
			}

			lissajous(w)
		}

		http.HandleFunc("/", handler)
		log.Fatal(http.ListenAndServe(":8000", nil))
	}
	lissajous(os.Stdout)
}

func lissajous(out io.Writer) {
	const (
		res     = 0.001 // 角度分辨率
		size    = 100   // 图像画布包含[-size..+size]
		nframes = 64    // 动画中的帧数
		delay   = 8     // 以10ms 为单位的帧间延迟
	)

	freq := rand.Float64() * 3.0 // y振荡器的相对频率
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0 // phrase difference
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < float64(cycles) * 2 * math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5),
				colorIndex)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	if err := gif.EncodeAll(out, &anim); err != nil {
		log.Fatal(err)
	}
}
