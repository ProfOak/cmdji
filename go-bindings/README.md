# Go language bindings for Kanji A Day (incomplete)

## Example usage

```
k := cmpdji.Kanji()
k.Update()
k.UnJson()
// Now you have a usable KanjiADay struct
// See Exported methods for data retrieval
```

## Exported Types

**KanjiADay** - struct

* **Kanji** - type kanjiStruct 
  * kanjiStruct is not exported. There is no reason to have a kanjiStruct structure because of the json blob layout.
* **HomeUrl** - string
  * url of the website (http://one.kanji.website)

## Exported methods

**Kanji**

```
func Kanji() KanjiADay {
    initialize and return KanjiADay object

}
```

**OpenKanji**

```
func OpenKanji(raw_json []bytes) KanjiADay {
    store json blob in unexported field
    initialize and return KanjiADay object

}
```


**Update**

```
func (k &KanjiADay) Update() {
    GET requst for json blob
    store json blob in unexported field
}

```

**UnJson**

```
func (k &KanjiADay) Unjson() {
    Take json blob, parse and store in k
    For use after Update() or OpenKanji()
}
```

**RawJson** 
 
```
func (k *KanjiADay) RawJson() []byte {
    return raw json blob
}
```

 **KanjiCharacter**

```
func (k KanjiADay) KanjiCharacter() string {
    return kanji character of the day
}
```


**Meanings**
```
func (k KanjiADay) Meanings() []string {
    return array of meanings for kanji character
}
```

**Onyomis**

```
func (k KanjiADay) Onyomis() []string {
    return list of on'yomi readings
}
```

**Kunyomis**

```
func (k KanjiADay) Kunyomis() []string {
    return list of kun'yomi readings
}
```

**Nanoris**

```
func (k KanjiADay) Nanoris() []string {
    return list of nanori readings
}
```

**Joyo**

```
func (k KanjiADay) Joyo() bool {
    return true if it's a joyo kanji
```

**Jlpt**

```
func (k KanjiADay) Jlpt() int {
    return jlpt level
}
```

**Newspaper_rank**

```
func (k KanjiADay) Newspaper_rank() int {
    return newspaper ranking
}
````

****
```
func (k KanjiADay) Max_newspaper_rank() int {
    return max newspaper rank
}
```

**Date**

```
func (k KanjiADay) Date() string {
    return the date the kanji of the day was from
}
```

**Image**

```
func (k KanjiADay) Image() []string {
    return the two URLs related to the kanji
    [
        direct image link,
        link to image's flicker page
    ]
}
```

**Source_url**

```
func (k KanjiADay) Source_url() string {
    return the date's source url for the kanji
}
```
