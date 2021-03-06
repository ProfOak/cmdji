package main

import (
	"fmt"
	"github.com/ProfOak/cmdji/go-bindings"
	"io/ioutil"
	"os/user"
	"strconv"
	"strings"
	"time"
)

func main() {
	// get home dir
	usr, err := user.Current()
	if err != nil {
		fmt.Println(err)
		// no user, no root, no service
		return
	}

	// .kanji_today in home dir
	kanji_today_filename := usr.HomeDir + "/.kanji_today"

	// file to store today's kanji
	// prevents multiple lookups on the server
	contents, err := ioutil.ReadFile(kanji_today_filename)

	var k cmdji.KanjiADay

	u2d, err := upToDate(contents)
	// first time running, or not up to date
	if err != nil || !u2d {
		// file not found/make file
		fmt.Println("Attempting to update from server...")
		k = cmdji.Kanji()
		err = k.Update()

		if err != nil {
			fmt.Println("Failed to retrieve new information from server")
			return
		} else {
			fmt.Println("Update successful")
		}

		ioutil.WriteFile(kanji_today_filename, k.RawJson(), 0744)

	} else {
		// already up to date contents in file
		k = cmdji.OpenKanji(contents)
	}

	// ready to use kanjiaday variable
	err = k.UnJson()

	// nicer printing
	str := k.KanjiCharacter() + "\n"
	str += strings.Join(k.Meanings(), ", ") + " "
	str += "JLPT: " + strconv.Itoa(k.Jlpt()) + "\n"

	if len(k.Kunyomis()) > 0 {
		str += "Kun'yomi: "
		str += strings.Join(k.Kunyomis(), ",")
	}
	if len(k.Kunyomis()) > 0 && len(k.Onyomis()) > 0 {
		str += " | "
	}
	if len(k.Onyomis()) > 0 {
		str += "On'yomi: "
		str += strings.Join(k.Onyomis(), ",")
	}

	fmt.Println(str)
	//k.Print()
}

func upToDate(contents []byte) (bool, error) {

	// check json blob from .kanji_today file
	past_kanji := cmdji.OpenKanji(contents)

	err := past_kanji.UnJson()
	if err != nil {
		// malformed json blob in file
		return false, err
	}

	date_raw := past_kanji.Date()

	// date_raw    = []string{date, time}
	// date_raw[0] = []string{year, month, day}
	old_date := strings.Split(strings.Split(date_raw, "T")[0], "-")

	old_year, _ := strconv.Atoi(old_date[0])
	old_month, _ := strconv.Atoi(old_date[1])
	old_day, _ := strconv.Atoi(old_date[2])

	// kanji a day uses UTC time
	// updates at midnight
	year, month, day := time.Now().UTC().Date()

	if year != old_year || int(month) != old_month || day != old_day {
		return false, err
	}
	return true, nil
}
