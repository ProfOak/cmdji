package cmdji

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

func unpack(cmpd [][]interface{}) []Compound {
	// used for on and kun compounds
	comps := make([]Compound, 0)
	for _, i := range cmpd {

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

func (k KanjiADay) On_compounds() []Compound {
	return unpack(k.Kanji.On_compounds)
}

func (k KanjiADay) Kun_compounds() []Compound {
	return unpack(k.Kanji.Kun_compounds)
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
