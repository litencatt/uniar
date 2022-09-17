package entity

import (
	"math"
)

type Scene struct {
	Photograph string
	Member     string
	Color      string
	Expect     float32
	Total      int32
	All35      int32
	All35Rank  int32
	VoDa50     int32
	VoDa50Rank int32
	DaPe50     int32
	DaPe50Rank int32
	VoPe50     int32
	VoPe50Rank int32
	Vo85       int32
	Vo85Rank   int32
	Da85       int32
	Da85Rank   int32
	Pe85       int32
	Pe85Rank   int32
	Vo         int32
	Da         int32
	Pe         int32
	SsrPlus    bool
}

type SceneTotal struct {
	Total  int32
	All35  int32
	All40  int32
	VoDa50 int32
	DaPe50 int32
	VoPe50 int32
	Vo85   int32
	Da85   int32
	Pe85   int32
}

type OfficeBonus struct {
	Vocal       int32
	Dance       int32
	Performance int32
}

var (
	all35  = 0.35
	voda50 = 0.50
	dape50 = 0.50
	vope50 = 0.50
	vo85   = 0.85
	da85   = 0.85
	pe85   = 0.85

	// フロントスキン
	frontSkinAll5  = 0.05
	frontSkinVoda8 = 0.08
	frontSkinDape8 = 0.08
	frontSkinVope8 = 0.08
)

var (
	bonds = []int32{
		0,
		0, 0, 10, 20, 30, 45, 45, 60, 75, 75, // 1-10
		95, 95, 115, 135, 155, 175, 175, 195, 220, 245, // 11-20
		270, 270, 295, 320, 345, 370, 370, 395, 420, 445,
		475, 475, 505, 535, 565, 595, 595, 625, 655, 685,
		715, 715, 745, 775, 805, 835, 835, 865, 895, 925,
		955, 985, 1020, 1055, 1090, 1125, 1125, 1160, 1195, 1230,
		1265, 1300, 1335, 1370, 1405, 1440, 1475, 1510, 1545, 1580,
		1615, 1650, 1685, 1720, 1755, 1790, 1825, 1860, 1895, 1930,
		1965, 2000, 2035, 2070, 2105, 2140, 2175, 2210, 2245, 2280,
		2315, 2350, 2385, 2420, 2455, 2490, 2525, 2560, 2595, 2630,
		2665, 2700, 2735, 2770, 2805, 2840, 2875, 2910, 2945, 2980,
		3015, 3050, 3085, 3120, 3155, 3190, 3225, 3260, 3295, 3330,
		3365, 3400, 3435, 3470, 3505, 3540, 3575, 3610, 3645, 3680,
		3715, 3750, 3785, 3820, 3855, 3890, 3925, 3960, 3995, 4030,
		4065, 4100, 4135, 4170, 4205, 4240, 4275, 4310, 4345, 4380,
		4415, 4450, 4485, 4520, 4555, 4590, 4625, 4660, 4695, 4730,
		4765, 4800, 4835, 4870, 4905, 4940, 4975, 5010, 5045, 5080,
		5115, 5150, 5185, 5220, 5255, 5290, 5325, 5360, 5395, 5430,
		5465, 5500, 5535, 5570, 5605, 5640, 5675, 5710, 5745, 5780,
		5815, 5850, 5885, 5920, 5955, 5990, 6025, 6060, 6095, 6130,
		6165, 6200, 6235, 6270, 6305, 6340, 6375, 6410, 6445, 6480,
		6515, 6550, 6585, 6620, 6655, 6690, 6725, 6760, 6795, 6830,
		6865, 6900, 6935, 6970, 7005, 7040, 7075, 7110, 7145, 7180,
		7215, 7250, 7285, 7320, 7355, 7390, 7425, 7460, 7495, 7530,
		7565, 7600, 7635, 7670, 7705, 7740, 7775, 7810, 7845, 7880, // 241-250
	}
)

type BonusRate struct {
	Vo float32
	Da float32
	Pe float32
}

func (s *Scene) CalcTotal(bondLevel int32, discography int32) {
	centerSkills := []string{"All35", "VoDa50", "DaPe50", "VoPe50", "Vo85", "Da85", "Pe85", "All40"}
	vo := s.Vo
	da := s.Da
	pe := s.Pe
	e := s.Expect
	t := s.Total

	//typeBonus := 0.3
	//musicBonus := 0.3

	// Max level
	voOb1 := 0.12
	daOb1 := 0.12
	peOb1 := 0.12

	// Max level
	voOb2 := 0.05
	daOb2 := 0.05
	peOb2 := 0.05

	// Max level
	sceneSkillLvMaxBonus := int32(430)

	// Fixed. because this is tiny score for rank
	costumeBonus := int32(300)

	cs := map[string]BonusRate{
		"All35":  {Vo: 0.35, Da: 0.35, Pe: 0.35},
		"VoDa50": {Vo: 0.5, Da: 0.5, Pe: 0},
		"DaPe50": {Vo: 0, Da: 0.5, Pe: 0.5},
		"VoPe50": {Vo: 0.5, Da: 0, Pe: 0.5},
		"Vo85":   {Vo: 0.85, Da: 0, Pe: 0},
		"Da85":   {Vo: 0, Da: 0.85, Pe: 0},
		"Pe85":   {Vo: 0, Da: 0, Pe: 0.85},
	}
	fs := map[string]BonusRate{
		"All5":  {Vo: 0.05, Da: 0.05, Pe: 0.05},
		"VoDa8": {Vo: 0.08, Da: 0.08, Pe: 0},
		"DaPe8": {Vo: 0, Da: 0.08, Pe: 0.08},
		"VoPe8": {Vo: 0.08, Da: 0, Pe: 0.08},
	}
	// 対象シーンカードの総合力計算
	totals := map[string]int32{}
	for _, skill := range centerSkills {
		csr := cs[skill]
		voCsb := int32(math.Ceil(float64(vo) * float64(csr.Vo)))
		daCsb := int32(math.Ceil(float64(da) * float64(csr.Da)))
		peCsb := int32(math.Ceil(float64(pe) * float64(csr.Pe)))
		csb := voCsb + daCsb + peCsb

		fsr := fs[skill]
		voFsb := int32(math.Ceil(float64(vo) * float64(fsr.Vo)))
		daFsb := int32(math.Ceil(float64(da) * float64(fsr.Da)))
		peFsb := int32(math.Ceil(float64(pe) * float64(fsr.Pe)))
		fsb := voFsb + daFsb + peFsb

		voOb := int32(math.Ceil(float64(vo) * (voOb1 + voOb2)))
		daOb := int32(math.Ceil(float64(da) * (daOb1 + daOb2)))
		peOb := int32(math.Ceil(float64(pe) * (peOb1 + peOb2)))
		ob := voOb + daOb + peOb

		// メンバーステータス
		member := bonds[bondLevel] + discography

		totals[skill] = int32(e * float32(t+csb+fsb+ob+member+sceneSkillLvMaxBonus+costumeBonus))
	}

	s.All35 = totals["All35"]
	s.VoDa50 = totals["VoDa50"]
	s.DaPe50 = totals["DaPe50"]
	s.VoPe50 = totals["VoPe50"]
	s.Vo85 = totals["Vo85"]
	s.Da85 = totals["Da85"]
	s.Pe85 = totals["Pe85"]
}
