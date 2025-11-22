package main

import (
	"image"
	"image/color"
	"strconv"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

// 游戏状态
type Game struct {
	player             GameObject
	mooncakes          []GameObject
	score              int
	award              bool
	time               int
	speedMultiplier    float64
	maxSpeedMultiplier float64 //最高速度倍速，防止速度过快
}

// 初始化游戏
func NewGame() *Game {
	rd.Seed(time.Now().UnixNano())
	bounds := rabbitImg.Bounds()
	width := float64(bounds.Dx()) * rabbitImgScale
	height := float64(bounds.Dy()) * rabbitImgScale
	return &Game{
		player:             NewGameObject(screenWidth/2, screenHeight-height, width, height, 0, true, 0),
		mooncakes:          make([]GameObject, 0),
		score:              0,
		award:              false,
		time:               0,
		speedMultiplier:    1.0,
		maxSpeedMultiplier: 1.2,
	}
}

// 更新游戏状态
func (g *Game) Update() error {
	// 移动玩家
	if ebiten.IsKeyPressed(ebiten.KeyLeft) && g.player.x > 0 {
		g.player.x -= playerSpeed
	}
	if ebiten.IsKeyPressed(ebiten.KeyRight) && g.player.x < screenWidth-g.player.width {
		g.player.x += playerSpeed
	}

	// 随时间生成月饼或炸弹
	g.time++
	if g.time%90 == 0 { // 每90帧生成一个物品
		objType := objTypeMooncake // 默认生成月饼
		speed := 2 + rd.Float64()*2

		// 10%概率生成炸弹
		if rd.Float64() < 0.1 {
			objType = objTypeBomb
		}
		if g.score%20 == 0 && g.score != 0 && !g.award {
			//每得20分奖励一个嫦娥
			objType = objTypeChange
			g.award = true
			speed = 0.815
		} else if g.score%20 != 0 {
			g.award = false
		}

		var bounds image.Rectangle
		var width, height float64
		if objType == objTypeRabbit {
			bounds = rabbitImg.Bounds()
			width = float64(bounds.Dx()) * rabbitImgScale
			height = float64(bounds.Dy()) * rabbitImgScale
		} else if objType == objTypeMooncake {
			bounds = mooncakeImg.Bounds()
			width = float64(bounds.Dx()) * mooncakeImgScale
			height = float64(bounds.Dy()) * mooncakeImgScale
		} else if objType == objTypeChange {
			bounds = changeImg.Bounds()
			width = float64(bounds.Dx()) * changeImgScale
			height = float64(bounds.Dy()) * changeImgScale
		} else if objType == objTypeBomb {
			bounds = bombImg.Bounds()
			width = float64(bounds.Dx()) * bombImgScale
			height = float64(bounds.Dy()) * bombImgScale
		}

		mooncake := NewGameObject(rd.Float64()*(screenWidth-width),
			-height, width, height, speed, true, objType)
		g.mooncakes = append(g.mooncakes, mooncake)
	}

	// 增加游戏难度
	if g.time%600 == 0 && g.speedMultiplier < g.maxSpeedMultiplier { // 每10秒增加一次速度，最高为1.2
		g.speedMultiplier += 0.1
	}

	// 更新月饼位置并检测碰撞
	for i := range g.mooncakes {
		if g.mooncakes[i].isActive {
			g.mooncakes[i].y += g.mooncakes[i].speed * g.speedMultiplier

			// 检测碰撞
			if checkCollision(g.player, g.mooncakes[i]) {
				g.mooncakes[i].isActive = false
				if g.mooncakes[i].objType == objTypeChange {
					// 接到嫦娥加分
					g.score += 10
				}
				if g.mooncakes[i].objType == objTypeBomb {
					// 接到炸弹减分
					g.score -= 5
					if g.score < 0 {
						g.score = 0
					}
				} else {
					// 接到月饼加分
					g.score += 1
				}
			}

			// 移除超出屏幕的月饼
			if g.mooncakes[i].y > screenHeight {
				g.mooncakes[i].isActive = false
			}
		}
	}

	// 清理不活跃的月饼
	activeMooncakes := make([]GameObject, 0)
	for _, m := range g.mooncakes {
		if m.isActive {
			activeMooncakes = append(activeMooncakes, m)
		}
	}
	g.mooncakes = activeMooncakes

	return nil
}

// 绘制游戏画面
func (g *Game) Draw(screen *ebiten.Image) {
	// 绘制背景 (夜空)
	screen.Fill(color.RGBA{0x0a, 0x0a, 0x23, 0xf0})

	// 绘制月亮
	drawMoon(screen, 815, -20)

	// 绘制玩家 (玉兔)
	drawRabbit(screen, g.player.x, g.player.y)

	// 绘制月饼和炸弹
	for _, m := range g.mooncakes {
		if m.isActive {
			if m.objType == objTypeChange {
				drawChange(screen, m.x, m.y)
			} else if m.objType == objTypeBomb {
				drawBomb(screen, m.x, m.y)
			} else {
				drawMooncake(screen, m.x, m.y)
			}
		}
	}

	// 绘制分数
	ebitenutil.DebugPrint(screen, "Scores: "+strconv.Itoa(g.score))
}

// 游戏窗口尺寸
func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func main() {
	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("中秋节 - 玉兔接月饼")
	ebiten.SetTPS(60)

	game := NewGame()
	opts := &ebiten.RunGameOptions{
		ScreenTransparent: true,
	}
	if err := ebiten.RunGameWithOptions(game, opts); err != nil {
		panic(err)
	}
}
