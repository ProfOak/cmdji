package cmdji

type Compound struct {
	KanjiCompound    string
	KanaCompound     string
	CompoundMeanings []string
}

type kanjiStruct struct {
	Kanji              string          `json:"kanji"`
	Meanings           []string        `json:"meanings"`
	Onyomis            []string        `json:"onyomis"`
	Kunyomis           []string        `json:"kunyomis`
	Nanoris            []string        `json:"nanoris"`
	Joyo               bool            `json:"joyo"`
	Jlpt               int             `json:"jlpt"`
	Newspaper_rank     int             `json:"newspaper_rank"`
	On_compounds       [][]interface{} `json:"on_compounds"`
	Kun_compounds      [][]interface{} `json:"kun_compounds"`
	Max_newspaper_rank int             `json:"max_newspaper_rank"`
	Published_at       string          `json:"published_at"`
	Image              []string        `json:"image"`
	Source_url         string          `json:"source_url"`
}

type KanjiADay struct {
	Kanji    kanjiStruct `json:"kanji"`
	Home_url string      `json:"home_url"`
}
