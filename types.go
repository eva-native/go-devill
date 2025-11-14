package godevill

type SuperPet struct {
	Class   string `json:"pet_class"`
	Level   string `json:"pet_level"`
	Name    string `json:"pet_name"`
	Wounded bool   `json:"wounded"`
}

type Superhero struct {
	GodName  string `json:"godname"`
	GodPower int    `json:"godpower"`

	Name        string `json:"name"`
	Gender      string `json:"gender"`
	Level       int    `json:"level"`
	Health      int    `json:"health"`
	MaxHealth   int    `json:"max_health"`
	Motto       string `json:"motto"`
	Alignment   string `json:"alignment"`
	Aura        string `json:"string"`
	ExpProgress int    `json:"exp_progress"`
	Distance    int    `json:"distance"`
	Gold        string `json:"gold_approx"`
	Town        string `json:"town_name"`

	Quest         string `json:"quest"`
	QuestProgress int    `json:"quest_progress"`

	SiedJobProgress int    `json:"side_job_progress"`
	SiedJob         string `json:"side_job"`

	InventoryMax int   `json:"inventory_max_num"`
	Inventory    int   `json:"inventory_num"`
	Activa       []any `json:"activatables"`

	RelicsPercent int `json:"relics_percent"`

	FightType string `json:"fight_type"`

	Clan         string `json:"clan"`
	ClanPosition string `json:"clan_position"`

	Pet SuperPet `json:"pet"`

	DiaryLast string `json:"diary_last"`
	EyeLast   string `json:"eye_last"`

	BricksCounr int `json:"bricks_cnt"`

	Savings            int    `json:"savings"`
	SavingsCompletedAt string `json:"savings_completed_at"`
	TempleCompletedAt  string `json:"temple_completed_at"`
	WoodCounr          int    `json:"wood_cnt"`

	ShopName   string `json:"shop_name"`
	TradeLevel int    `json:"t_level"`

	ArcFemale      int    `json:"ark_f"`
	ArcMale        int    `json:"ark_m"`
	ArcCompletedAt string `json:"ark_completed_at"`
	ArcName        string `json:"arc_name"`
	PairsAt        string `json:"pairs_at"`

	ArenaWon     int  `json:"arena_won"`
	ArenaLost    int  `json:"arena_lost"`
	ArenaInFight bool `json:"arena_fight"`

	Words  int    `json:"words"`
	BookAt string `json:"book_at"`

	SoulsAt      string `json:"souls_at"`
	SoulsPercent int    `json:"souls_percent"`

	BossName  string `json:"boss_name"`
	BossPower int    `json:"boss_power"`

	Expired bool `json:"expired"`
}
