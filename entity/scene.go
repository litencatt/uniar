package entity

import (
	"math"
)

type Scene struct {
	Color       string
	Photograph  string
	PhotographID int64
	SsrPlus     bool
	Member      string
	MemberID    int64
	Expect      float32
	Total       int64
	All35Score  int64
	All35       int64
	VoDa50Score int64
	VoDa50      int64
	DaPe50Score int64
	DaPe50      int64
	VoPe50Score int64
	VoPe50      int64
	Vo85Score   int64
	Vo85        int64
	Da85Score   int64
	Da85        int64
	Pe85Score   int64
	Pe85        int64
	Vo          int64
	Da          int64
	Pe          int64
}

type SceneTotal struct {
	Total  int64
	All35  int64
	All40  int64
	VoDa50 int64
	DaPe50 int64
	VoPe50 int64
	Vo85   int64
	Da85   int64
	Pe85   int64
}

type OfficeBonus struct {
	Vocal       int64
	Dance       int64
	Performance int64
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
	bonds = []int64{
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

func (s *Scene) CalcTotal(bondLevel int64, discography int64) {
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
	sceneSkillLvMaxBonus := int64(430)

	// Fixed. because this is tiny score for rank
	costumeBonus := int64(300)

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
	totals := map[string]int64{}
	for _, skill := range centerSkills {
		csr := cs[skill]
		voCsb := int64(math.Ceil(float64(vo) * float64(csr.Vo)))
		daCsb := int64(math.Ceil(float64(da) * float64(csr.Da)))
		peCsb := int64(math.Ceil(float64(pe) * float64(csr.Pe)))
		csb := voCsb + daCsb + peCsb

		fsr := fs[skill]
		voFsb := int64(math.Ceil(float64(vo) * float64(fsr.Vo)))
		daFsb := int64(math.Ceil(float64(da) * float64(fsr.Da)))
		peFsb := int64(math.Ceil(float64(pe) * float64(fsr.Pe)))
		fsb := voFsb + daFsb + peFsb

		voOb := int64(math.Ceil(float64(vo) * (voOb1 + voOb2)))
		daOb := int64(math.Ceil(float64(da) * (daOb1 + daOb2)))
		peOb := int64(math.Ceil(float64(pe) * (peOb1 + peOb2)))
		ob := voOb + daOb + peOb

		// メンバーステータス
		member := bonds[bondLevel] + discography

		totals[skill] = int64(e * float32(t+csb+fsb+ob+member+sceneSkillLvMaxBonus+costumeBonus))
	}

	s.All35Score = totals["All35"]
	s.VoDa50Score = totals["VoDa50"]
	s.DaPe50Score = totals["DaPe50"]
	s.VoPe50Score = totals["VoPe50"]
	s.Vo85Score = totals["Vo85"]
	s.Da85Score = totals["Da85"]
	s.Pe85Score = totals["Pe85"]
}
