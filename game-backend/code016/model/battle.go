package model

// Skill 定义技能结构体
type Skill struct {
	Name         string `json:"name"`
	Damage       int    `json:"damage"`
	ManaCost     int    `json:"manaCost"`
	CoolDown     int    `json:"coolDown"`
	IsOnCooldown bool   `json:"isOnCooldown"`
}

// Position 定义位置结构体
type Position struct {
	X int `json:"x"`
	Y int `json:"y"`
}

// BattleInfo 定义战斗信息结构体
type BattleInfo struct {
	PlayerID       string   `json:"playerID"`
	Health         int      `json:"health"`
	Mana           int      `json:"mana"`
	Action         string   `json:"action"`
	Skills         []Skill  `json:"skills"`
	Position       Position `json:"position"`
	StatusEffects  []string `json:"statusEffects"`
	TargetPlayerID string   `json:"targetPlayerID"`
}
