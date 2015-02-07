package cmdji

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

var json_blob []byte

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

func Kanji() KanjiADay {
	var k KanjiADay
	return k
}

func OpenKanji(b []byte) KanjiADay {
	var k KanjiADay
	json_blob = b
	return k
}

func (k *KanjiADay) Print() {
	// sanity test
	fmt.Println("Kanji:", k.KanjiCharacter())
	fmt.Println("Meanings:", k.Meanings())
	fmt.Println("Onyomis:", k.Onyomis())
	fmt.Println("Kunyomis:", k.Kunyomis())
	fmt.Println("Nanoris:", k.Nanoris())
	fmt.Println("Joyo:", k.Joyo())
	fmt.Println("Jlpt:", k.Jlpt())
	fmt.Println("Newspaper_ran:", k.Newspaper_rank())
	fmt.Println("On_compounds:", k.On_compounds())
	fmt.Println("Kun_compounds:", k.Kun_compounds())
	fmt.Println("Max_newspaper_ran:", k.Newspaper_rank())
	fmt.Println("Published_at:", k.Date())
	fmt.Println("Image:", k.Image())
	fmt.Println("Source_url:", k.Source_url())
	fmt.Println("Home_url:", k.Home())

}

func (k *KanjiADay) Update() error {
	// retrieve from website
	const url = "http://one.kanji.website/api/v1/today.json"
	res, err := http.Get(url)
	if err != nil {
		return errors.New("Unable to connect to website")
	}
	defer res.Body.Close()
	json_blob, _ = ioutil.ReadAll(res.Body)

	return nil
}

func (k *KanjiADay) UnJson() error {
	// take json blob and unmarshal it
	// TODO: Check for len > 0
	// Update or OpenKanji must be run before this
	err := json.Unmarshal(json_blob, &k)
	if err != nil {
		return errors.New("Error: Malformed json blob")
	}
	return nil
}

func (k *KanjiADay) RawJson() []byte {
	return json_blob
}

func (k KanjiADay) KanjiCharacter() string {
	return k.Kanji.Kanji
}

func (k KanjiADay) Meanings() []string {
	return k.Kanji.Meanings
}

func (k KanjiADay) Onyomis() []string {
	return k.Kanji.Onyomis
}

func (k KanjiADay) Kunyomis() []string {
	return k.Kanji.Kunyomis
}

func (k KanjiADay) Nanoris() []string {
	return k.Kanji.Nanoris
}

func (k KanjiADay) Joyo() bool {
	return k.Kanji.Joyo
}

func (k KanjiADay) Jlpt() int {
	return k.Kanji.Jlpt
}

func (k KanjiADay) Newspaper_rank() int {
	return k.Kanji.Newspaper_rank
}

func (k KanjiADay) On_compounds() []Compound {
	comps := make([]Compound, 0)
	for _, i := range k.Kanji.On_compounds {

		x := i[0].(string)     // kanji reading
		y := i[1].(string)     // kana reading
		z := make([]string, 0) // meanings

		for _, j := range i[2].([]interface{}) {
			// collect all the meanings
			z = append(z, j.(string))
		}

		// build Compound struct
		var c Compound
		c.KanjiCompound = x
		c.KanaCompound = y
		c.CompoundMeanings = z

		// collect all Compound structs, for returning
		comps = append(comps, c)
	}

	return comps
}

func (k KanjiADay) Kun_compounds() []Compound {
	comps := make([]Compound, 0)
	for _, i := range k.Kanji.Kun_compounds {

		x := i[0].(string)     // kanji reading
		y := i[1].(string)     // kana reading
		z := make([]string, 0) // meanings

		for _, j := range i[2].([]interface{}) {
			// collect all the meanings
			z = append(z, j.(string))
		}

		// build Compound struct
		var c Compound
		c.KanjiCompound = x
		c.KanaCompound = y
		c.CompoundMeanings = z

		// collect all Compound structs, for returning
		comps = append(comps, c)
	}
	return comps
}

func (k KanjiADay) Max_newspaper_rank() int {
	return k.Kanji.Max_newspaper_rank
}

func (k KanjiADay) Date() string {
	return k.Kanji.Published_at
}

func (k KanjiADay) Image() []string {
	return k.Kanji.Image
}

func (k KanjiADay) Source_url() string {
	return k.Kanji.Source_url
}

func (k KanjiADay) Home() string {
	return k.Home_url
}
