package cmdji

/*
	Make .kanji_today file for current day's kanji notes, so you don't have to check for new kanji every time you open the term
*/
import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	//"strings"
	"fmt"
)

func Kanji() KanjiADay {
	var k KanjiADay
	return k
}

func OpenKanji(b []byte) KanjiADay {
	var k KanjiADay
	json_blob = b
	return k
}

var json_blob []byte

type kanjiStruct struct {
	Kanji              string        `json:"kanji"`
	Meanings           []string      `json:"meanings"`
	Onyomis            []string      `json:"onyomis"`
	Kunyomis           []string      `json:"kunyomis`
	Nanoris            []string      `json:"nanoris"`
	Joyo               bool          `json:"joyo"`
	Jlpt               int           `json:"jlpt"`
	Newspaper_rank     int           `json:"newspaper_rank"`
	On_compounds       []interface{} `json:"on_compounds"`
	Kun_compounds      []interface{} `json:"kun_compounds"`
	Max_newspaper_rank int           `json:"max_newspaper_rank"`
	Published_at       string        `json:"published_at"`
	Image              []string      `json:"image"`
	Source_url         string        `json:"source_url"`
}

type KanjiADay struct {
	Kanji    kanjiStruct `json:"kanji"`
	Home_url string      `json:"home_url"`
}

func (k *KanjiADay) Print() {
	// sanity test
	fmt.Println("Kanji:", k.Kanji.Kanji)
	fmt.Println("Meanings:", k.Kanji.Meanings)
	fmt.Println("Onyomis:", k.Kanji.Onyomis)
	fmt.Println("Kunyomis:", k.Kanji.Kunyomis)
	fmt.Println("Nanoris:", k.Kanji.Nanoris)
	fmt.Println("Joyo:", k.Kanji.Joyo)
	fmt.Println("Jlpt:", k.Kanji.Jlpt)
	fmt.Println("Newspaper_ran:", k.Kanji.Newspaper_rank)
	fmt.Println("On_compounds:", k.Kanji.On_compounds)
	fmt.Println("Kun_compounds:", k.Kanji.Kun_compounds)
	fmt.Println("Max_newspaper_ran:", k.Kanji.Max_newspaper_rank)
	fmt.Println("Published_at:", k.Kanji.Published_at)
	fmt.Println("Image:", k.Kanji.Image)
	fmt.Println("Source_url:", k.Kanji.Source_url)
	fmt.Println("Home_url:", k.Home_url)

}

func (k *KanjiADay) Update() {
	const url = "http://one.kanji.website/api/v1/today.json"
	res, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
	}
	defer res.Body.Close()
	json_blob, err = ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
	}

}
func (k *KanjiADay) RawJson() []byte {
	return json_blob
}
func (k *KanjiADay) UnJson() {

	err := json.Unmarshal(json_blob, &k)
	if err != nil {
		fmt.Println("[Fail]", err)
	}

}
