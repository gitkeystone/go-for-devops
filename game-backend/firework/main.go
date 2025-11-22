package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"image/color"
	"math"
	"math/rand"
	"time"
)

const (
	// screenWidth  = 1000
	// screenHeight = 800
	gravity = 0.1
	// 阶段控制常量
	phase1Duration = 10 * time.Second // 第一阶段：烟花少且慢
	phase2Duration = 15 * time.Second // 第二阶段：烟花逐渐变多变快
	phase3Duration = 20 * time.Second // 第三阶段：烟花密集
)

var (
	rd           *rand.Rand
	screenWidth  = 1000
	screenHeight = 800
)

// 初始化字体
func init() {
	rd = rand.New(rand.NewSource(time.Now().Unix()))
	screenWidth, screenHeight = ebiten.Monitor().Size()
}

// 粒子结构体，表示烟花的粒子
type Particle struct {
	x, y     float64
	vx, vy   float64
	color    color.RGBA
	lifetime int
	maxLife  int
	targetX  float64 // 用于文字阶段的目标位置
	targetY  float64
}

// NewParticle 初始化新粒子
func NewParticle(x, y, vx, vy float64, color color.RGBA, maxLife int, targetX, targetY float64) *Particle {
	return &Particle{
		x:        x,
		y:        y,
		vx:       vx,
		vy:       vy,
		color:    color,
		lifetime: 0,
		maxLife:  maxLife,
		targetX:  targetX,
		targetY:  targetY,
	}
}

// Update 更新粒子状态
func (p *Particle) Update() {
	p.x += p.vx
	p.y += p.vy
	p.vy += gravity // 应用重力
	p.lifetime++
}

// IsDead 检查粒子是否已过期
func (p *Particle) IsDead() bool {
	return p.lifetime >= p.maxLife
}

// Firework 烟花结构体
type Firework struct {
	particles []*Particle
	exploded  bool
	x, y      float64
	vx, vy    float64
	color     color.RGBA
}

// NewFirework 创建新烟花
func NewFirework() *Firework {
	// 随机位置和初始速度
	x := rd.Float64() * float64(screenWidth)
	y := float64(screenHeight)
	vx := (rd.Float64() - 0.5) * 2
	vy := -(rd.Float64()*6 + 10)

	// 国庆节主题颜色：红色、黄色为主
	colors := []color.RGBA{
		{255, 0, 0, 255},   // 红色
		{255, 50, 50, 255}, // 浅红
		{255, 100, 0, 255}, // 橙色
		{255, 255, 0, 255}, // 黄色
		{255, 165, 0, 255}, // 橙黄色
		{255, 215, 0, 255}, // 金黄色
	}
	c := colors[rd.Intn(len(colors))]

	return &Firework{
		particles: []*Particle{},
		exploded:  false,
		x:         x,
		y:         y,
		vx:        vx,
		vy:        vy,
		color:     c,
	}
}

// 更新烟花状态
func (f *Firework) Update() {
	if !f.exploded {
		// 烟花上升阶段
		f.x += f.vx
		f.y += f.vy
		f.vy += gravity // 应用重力

		// 当烟花到达顶点或速度变为正时，爆炸
		if f.vy >= 0 {
			f.explode()
		}
	} else {
		// 更新所有粒子
		for i := len(f.particles) - 1; i >= 0; i-- {
			f.particles[i].Update()
			if f.particles[i].IsDead() {
				// 移除过期粒子
				f.particles = append(f.particles[:i], f.particles[i+1:]...)
			}
		}
	}
}

// 烟花爆炸效果
func (f *Firework) explode() {
	f.exploded = true
	particleCount := 80 + rd.Intn(40) // 80-120个粒子

	// 创建爆炸粒子
	for i := 0; i < particleCount; i++ {
		angle := rd.Float64() * 2 * math.Pi
		speed := rd.Float64()*3 + 1

		vx := math.Cos(angle) * speed
		vy := math.Sin(angle) * speed

		// 为每个粒子添加一些颜色变化，增加视觉效果
		r := f.color.R
		g := f.color.G
		b := f.color.B
		if r > 50 {
			r -= uint8(rd.Intn(50))
		}
		if g > 50 {
			g -= uint8(rd.Intn(50))
		}
		if b > 50 {
			b -= uint8(rd.Intn(50))
		}

		particleColor := color.RGBA{r, g, b, 255}
		maxLife := 50 + rd.Intn(40) // 粒子生命周期随机

		f.particles = append(f.particles, NewParticle(f.x, f.y, vx, vy, particleColor, maxLife, 0, 0))
	}
}

// 检查烟花是否已结束
func (f *Firework) IsFinished() bool {
	return f.exploded && len(f.particles) == 0
}

// 绘制烟花
func (f *Firework) Draw(screen *ebiten.Image) {
	if !f.exploded {
		// 绘制上升中的烟花
		vector.DrawFilledCircle(screen, float32(f.x), float32(f.y), 3, f.color, false)
	} else {
		// 绘制爆炸后的粒子
		for _, p := range f.particles {
			// 粒子随时间逐渐消失
			alpha := uint8(255 * (1 - float64(p.lifetime)/float64(p.maxLife)*0.8))
			c := color.RGBA{p.color.R, p.color.G, p.color.B, alpha}
			radius := 2.0
			if p.lifetime < p.maxLife/3 {
				radius = 3.0 // 刚爆炸时粒子更大
			}
			vector.DrawFilledCircle(screen, float32(p.x), float32(p.y), float32(radius), c, false)
		}
	}
}

// 游戏结构体
type Game struct {
	fireworks []*Firework
	startTime time.Time
	spawnRate float64
	phase     int // 0:准备阶段, 1:第一阶段, 2:第二阶段, 3:第三阶段
}

// 初始化游戏
func NewGame() *Game {
	return &Game{
		fireworks: []*Firework{},
		startTime: time.Now(),
		spawnRate: 0.02, // 初始生成概率
		phase:     1,
	}
}

// 更新游戏状态
func (g *Game) Update() error {
	elapsed := time.Since(g.startTime)

	// 阶段切换逻辑
	if elapsed > phase1Duration+phase2Duration {
		g.phase = 3 // 第三阶段：烟花密集
	} else if elapsed > phase1Duration {
		g.phase = 2 // 第二阶段：逐渐变快变多
		// 逐渐增加生成速率
		progress := (elapsed - phase1Duration).Seconds() / phase2Duration.Seconds()
		g.spawnRate = 0.02 + 0.10*progress // 从0.02增加到0.15
	}

	if rd.Float64() < g.spawnRate {
		g.fireworks = append(g.fireworks, NewFirework())
	}

	// 更新所有烟花
	for i := len(g.fireworks) - 1; i >= 0; i-- {
		g.fireworks[i].Update()
		if g.fireworks[i].IsFinished() {
			// 移除已结束的烟花
			g.fireworks = append(g.fireworks[:i], g.fireworks[i+1:]...)
		}
	}

	return nil
}

// 绘制游戏画面
func (g *Game) Draw(screen *ebiten.Image) {
	// 绘制所有烟花
	for _, f := range g.fireworks {
		f.Draw(screen)
	}
}

// 获取游戏窗口尺寸
func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return outsideWidth, outsideHeight
}

func main() {
	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("祝祖国母亲76周年快乐")
	ebiten.SetTPS(60)

	game := NewGame()

	if err := ebiten.RunGameWithOptions(game, &ebiten.RunGameOptions{
		ScreenTransparent: true,
	}); err != nil {
		panic(err)
	}
}
