package main

import (
	"math/rand"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

const (
	screenWidth  = 1280
	screenHeight = 720
	playerSpeed  = 50
)

const (
	objTypeRabbit   = 0
	objTypeMooncake = 1
	objTypeChange   = 2
	objTypeBomb     = 3
)

var (
	moonImg     *ebiten.Image
	mooncakeImg *ebiten.Image
	rabbitImg   *ebiten.Image
	changeImg   *ebiten.Image
	bombImg     *ebiten.Image
)

var (
	mooncakeImgScale = 0.3
	rabbitImgScale   = 0.3
	changeImgScale   = 0.5
	bombImgScale     = 0.3
)

// 游戏对象类型
type GameObject struct {
	x, y          float64
	width, height float64
	speed         float64
	isActive      bool
	objType       int // 物体类型，0:玩家，1:月饼，2:嫦娥，3:炸弹
}

// 新建游戏对象
func NewGameObject(x, y float64, width, height float64, speed float64, isActive bool, objType int) GameObject {
	return GameObject{
		x:        x,
		y:        y,
		width:    width,
		height:   height,
		speed:    speed,
		isActive: isActive,
		objType:  objType,
	}
}

var (
	rd = rand.New(rand.NewSource(time.Now().UnixNano()))
)

func init() {
	var err error
	moonImg, _, err = ebitenutil.NewImageFromFile("moon.png")
	if err != nil {
		panic(err)
	}
	mooncakeImg, _, err = ebitenutil.NewImageFromFile("mooncake.png")
	if err != nil {
		panic(err)
	}
	rabbitImg, _, err = ebitenutil.NewImageFromFile("rabbit.png")
	if err != nil {
		panic(err)
	}
	changeImg, _, err = ebitenutil.NewImageFromFile("change.png")
	if err != nil {
		panic(err)
	}
	bombImg, _, err = ebitenutil.NewImageFromFile("bomb.png")
	if err != nil {
		panic(err)
	}
}

// 检测碰撞
func checkCollision(a, b GameObject) bool {
	return a.x < b.x+b.width &&
		a.x+a.width > b.x &&
		a.y < b.y+b.height &&
		a.y+a.height > b.y
}

func drawMooncake(screen *ebiten.Image, x, y float64) {
	opts := &ebiten.DrawImageOptions{}
	// 设置图片绘制的位置（x, y）
	opts.GeoM.Scale(mooncakeImgScale, mooncakeImgScale)
	opts.GeoM.Translate(x, y)
	// 绘制图片到屏幕
	screen.DrawImage(mooncakeImg, opts)
}

func drawRabbit(screen *ebiten.Image, x, y float64) {
	opts := &ebiten.DrawImageOptions{}
	// 设置图片绘制的位置（x, y）
	opts.GeoM.Scale(rabbitImgScale, rabbitImgScale)
	opts.GeoM.Translate(x, y)
	// 绘制图片到屏幕
	screen.DrawImage(rabbitImg, opts)
}
func drawMoon(screen *ebiten.Image, x, y float64) {
	opts := &ebiten.DrawImageOptions{}
	// 设置图片绘制的位置（x, y）
	opts.GeoM.Scale(0.8, 0.8)
	opts.GeoM.Translate(x, y)
	// 绘制图片到屏幕
	screen.DrawImage(moonImg, opts)
}

func drawChange(screen *ebiten.Image, x, y float64) {
	opts := &ebiten.DrawImageOptions{}
	// 设置图片绘制的位置（x, y）
	opts.GeoM.Scale(changeImgScale, changeImgScale)
	opts.GeoM.Translate(x, y)

	// 绘制图片到屏幕
	screen.DrawImage(changeImg, opts)
}

func drawBomb(screen *ebiten.Image, x, y float64) {
	opts := &ebiten.DrawImageOptions{}
	// 设置图片绘制的位置（x, y）
	opts.GeoM.Scale(bombImgScale, bombImgScale)
	opts.GeoM.Translate(x, y)

	// 绘制图片到屏幕
	screen.DrawImage(bombImg, opts)
}
