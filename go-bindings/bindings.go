package cmdji

/*
	Make .kanji_today file for current day's kanji notes, so you don't have to check for new kanji every time you open the term
*/
import (
	"encoding/json"
	//"io/ioutil"
	//"net/http"
	//"strings"
	"fmt"
)

func Kanji() kanjiADay {
	/*
		// will uncomment this when I impliment the .kanji_today file

		url := "http://one.kanji.website/api/v1/today.json"
		res, err := http.Get(url)
		if err != nil {
		fmt.Println(err)
		}
		defer res.Body.Close()
		body, err := ioutil.ReadAll(res.Body)
		if err != nil {
		fmt.Println(err)
		}
	*/
	var k kanjiADay
	var body = []byte(`
	{"kanji":{"kanji":"示","meanings":["show","indicate","point out","express","display"],"onyomis":["ジ","シ"],"kunyomis":["しめ.す"],"joyo":true,"jlpt":3,"newspaper_rank":237,"max_newspaper_rank":2500,"on_compounds":[["示威運動","ジイウンドウ",["a demonstration"]],["示談","ジダン",["settlement out of court"]],["示威","ジイ",["demonstration","show of force"]],["示指","ジシ",["index finger","forefinger"]],["示寂","ジジャク",["death of a high-ranking priest"]],["示準化石","ジジュンカセキ",["index fossil (fossil which suggests a date of the formation of a rock stratum)","key fossil","leading fossil"]],["示現","ジゲン",["manifestation (of a celestial being)"]],["示達","ジタツ",["instructions","directions"]],["指示","シジ",["indication","instruction","designation","directions"]],["提示","テイジ",["presentation","exhibit","suggest","citation"]],["表示","ヒョウジ",["indication","expression","showing","manifestation","demonstration","display","displaying","representation"]],["掲示","ケイジ",["notice","bulletin","post","posting","placard"]],["示す","シメス",["to (take out and) show","to demonstrate","to tell","to exemplify","to make apparent","to point out (finger","clock hand","needle","etc.)","to indicate","to show","to represent","to signify","to display"]],[" 示唆","シサ",["suggestion","hint","implication"]],["示し","シメシ",["discipline","revelation"]],["示し合わせる","シメシアワセル",["to arrange beforehand","to make a sign to each other","to conspire"]],["示威","ジイ",["demonstration","show of force"]],["示差","シサ",["differential"]],["示相化石","シソウカセキ",["facies fossil"]],["示唆的","シサテキ",["suggestive","pregnant (e.g. pregnant pause)"]],["教示","キョウジ",["instruction","teaching"]],["公示","コウジ",["edict","public announcement"]],["内示","ナイジ",["unofficial announcement"]]],"kun_compounds":[["示す","しめす",["to (take out and) show","to demonstrate","to tell","to exemplify","to make apparent","to point out (finger","clock hand","needle","etc.)","to indicate","to show","to represent","to signify","to display"]],["示偏"," しめすへん",["kanji \"show\" radical at left (radical 113)"]]],"published_at":"2015-01-15T00:00:00.000Z","image":["https://farm4.staticflickr.com/3210/3047742878_944b1cc018_o.jpg","https://www.flickr.com/photos/74743437@N00/3047742878"],"source_url":"http://one.kanji.website/date/2015-01-15"},"home_url":"http://one.kanji.website"}
	`)

	err := json.Unmarshal(body, &k)
	if err != nil {
		fmt.Println("[Fail]", err)
	}

	return k

}

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

type kanjiADay struct {
	Kanji    kanjiStruct `json:"kanji"`
	Home_url string      `json:"home_url"`
}

func (k *kanjiADay) Print() {
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
