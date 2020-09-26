package main

// Champion describes a LOL champion
type Champion struct {
	Version   string   `json:"version"`
	ID        string   `json:"id"`
	Key       string   `json:"key"`
	Name      string   `json:"name"`
	Title     string   `json:"title"`
	Blurb     string   `json:"blurb"`
	Info      Info     `json:"info"`
	Tags      []string `json:"tags"`
	Partype   string   `json:"partype"`
	Stats     Stats    `json:"stats"`
	AllyTips  []string `json:"allytips"`
	EnemyTips []string `json:"enemytips"`
	Lore      string   `json:"lore"`
	Spells    []Spell  `json:"spells"`
}

// Info describes a champions playstyle
type Info struct {
	Attack     int `json:"attack"`
	Defense    int `json:"defense"`
	Magic      int `json:"magic"`
	Difficulty int `json:"difficulty"`
}

// Stats describes a champions detailed stats
type Stats struct {
	Hp                   int     `json:"hp"`
	Hpperlevel           int     `json:"hpperlevel"`
	Mp                   int     `json:"mp"`
	Mpperlevel           int     `json:"mpperlevel"`
	Movespeed            int     `json:"movespeed"`
	Armor                int     `json:"armor"`
	Armorperlevel        int     `json:"armorperlevel"`
	Spellblock           int     `json:"spellblock"`
	Spellblockperlevel   float64 `json:"spellblockperlevel"`
	Attackrange          int     `json:"attackrange"`
	Hpregen              float64 `json:"hpregen"`
	Hpregenperlevel      float64 `json:"hpregenperlevel"`
	Mpregen              int     `json:"mpregen"`
	Mpregenperlevel      float64 `json:"mpregenperlevel"`
	Crit                 int     `json:"crit"`
	Critperlevel         int     `json:"critperlevel"`
	Attackdamage         float64 `json:"attackdamage"`
	Attackdamageperlevel float64 `json:"attackdamageperlevel"`
	Attackspeedperlevel  float64 `json:"attackspeedperlevel"`
	Attackspeed          float64 `json:"attackspeed"`
}

// Spell describes a champion spell
type Spell struct {
	ID           string    `json:"id"`
	Name         string    `json:"name"`
	Description  string    `json:"description"`
	Tooltip      string    `json:"tooltip"`
	Maxrank      int       `json:"maxrank"`
	Cooldown     []float64 `json:"cooldown"`
	CooldownBurn string    `json:"cooldownBurn"`
	Cost         []int     `json:"cost"`
	CostBurn     string    `json:"costBurn"`
	Effect       []int     `json:"effect"`
	EffectBurn   []int     `json:"effectBurn"`
	Vars         []Vars    `json:"vars"`
	CostType     string    `json:"costType"`
	Maxammo      string    `json:"maxammo"`
	Range        []int     `json:"range"`
	RangeBurn    string    `json:"rangeBurn"`
	Resource     string    `json:"resource"`
}

// Vars describes spell values
type Vars struct {
	Key   string  `json:"vars"`
	Link  string  `json:"link"`
	Value float64 `json:"Value"`
}
