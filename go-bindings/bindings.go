package cmdji

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
)

var json_blob []byte

func Kanji() KanjiADay {
	var k KanjiADay
	return k
}

func OpenKanji(b []byte) KanjiADay {
	var k KanjiADay
	json_blob = b
	return k
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
