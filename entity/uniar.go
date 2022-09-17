package entity

import (
	"math"
)

type Scene struct {
	Photograph string
	Member     string
	Color      string
	Expect     string
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

type BonusRate struct {
	Vo float32
	Da float32
	Pe float32
}

func (s *Scene) CalcTotal() {
	centerSkills := []string{"All35", "VoDa50", "DaPe50", "VoPe50", "Vo85", "Da85", "Pe85", "All40"}
	vo := s.Vo
	da := s.Da
	pe := s.Pe
	t := s.Total

	// フロント共通
	// groupTotal := 600000
	//typeBonus := 0.3
	//musicBonus := 0.3

	voOb1 := 0.12
	daOb1 := 0.12
	peOb1 := 0.12

	voOb2 := 0.05
	daOb2 := 0.05
	peOb2 := 0.05

	bondsBonus := int32(1000)
	discography := int32(1000)

	sceneSkillLvMaxBonus := int32(430)
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
		member := bondsBonus + discography

		totals[skill] = t + csb + fsb + ob + member + sceneSkillLvMaxBonus + costumeBonus
	}

	s.All35 = totals["All35"]
	s.VoDa50 = totals["VoDa50"]
	s.DaPe50 = totals["DaPe50"]
	s.VoPe50 = totals["VoPe50"]
	s.Vo85 = totals["Vo85"]
	s.Da85 = totals["Da85"]
	s.Pe85 = totals["Pe85"]
}
