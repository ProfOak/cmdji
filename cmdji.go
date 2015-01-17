package main

import (
	"fmt"
	"github.com/ProfOak/cmdji/go-bindings"
	"io/ioutil"
	"os/user"
	"time"
)

func main() {
	// open file .kanji_today
	// throw in json.Unmarshal
	// check date - make date function
	// if date > .kanji_today date,

	usr, err := user.Current()
	if err != nil {
		fmt.Println(err)
	}

	kanji_today_filename := usr.HomeDir + "/.kanji_today"

	// file to store today's kanji
	// prevents multiple lookups on the server
	contents, err := ioutil.ReadFile(kanji_today_filename)

	var k cmdji.KanjiADay

	// first time running, or not up to date
	if err != nil || !upToDate(contents) {
		// file not found/make file
		fmt.Println("UPDATING")
		k = cmdji.Kanji()
		k.Update()
		ioutil.WriteFile(kanji_today_filename, k.RawJson(), 0744)
	} else {
		// already up to date contents in file
		k = cmdji.OpenKanji(contents)
	}

	// ready to use kanjiaday variable
	k.UnJson()
	//k.Print()

	// check date to match today (match time zone of server)

	fmt.Println(k.Kanji.Kanji)
}

func upToDate(contents []byte) bool {

	past_kanji := cmdji.OpenKanji(contents)
	past_kanji.UnJson()
	fmt.Println("date", past_kanji.Kanji.Published_at)
	fmt.Println(time.Now())
	return true
}
