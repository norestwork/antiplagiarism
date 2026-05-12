package core

import (
	"antiplagiarism/core/set"
	"math"
	"strings"
)

type (
	AnalyseMode string
)

const (
	LexicalAnalyse AnalyseMode = "lexical"
	StructuralAnalyse AnalyseMode = "structural"
)

type Core struct {}

func NewCore() Core {
	return Core{}
}

type GeneralAnalyseResult struct {
	General float64
	Structural float64
	Lexical float64
}

func (c *Core) GeneralAnalyse(text1 string, text2 string) GeneralAnalyseResult {
	text1 = c.clearText(text1)
	text2 = c.clearText(text2)

	wordsCountText1 := len(strings.Split(text1, " "))
	wordsCountText2 := len(strings.Split(text2, " "))

	averageCount := (wordsCountText1 + wordsCountText2) / 2
	imbalance := float64(min(wordsCountText1, wordsCountText2)) / float64(max(wordsCountText1, wordsCountText2))

	var structuralAnalyseWeight float64

	if averageCount < 20 {
		structuralAnalyseWeight = 0.2
	} else if averageCount < 100 {
		structuralAnalyseWeight = 0.5
	} else {
		structuralAnalyseWeight = 0.7
	}

	structuralAnalyseWeight = max(0.2, structuralAnalyseWeight * imbalance)
	lexicalAnalyseWeight := 1 - structuralAnalyseWeight

	structuralAnalyse, _ := c.Analyse(text1, text2, StructuralAnalyse)
	lexicalAnalyse, _ := c.Analyse(text1, text2, LexicalAnalyse)

	generalAnalyse := structuralAnalyse * structuralAnalyseWeight + lexicalAnalyse * lexicalAnalyseWeight

	return GeneralAnalyseResult{
		General: math.Round(generalAnalyse * 100) / 100,
		Structural: structuralAnalyse,
		Lexical: lexicalAnalyse,
	}
}

func (c *Core) Analyse(text1 string, text2 string, mode AnalyseMode) (float64, error) {
	var shingleSize int
	
	switch mode {
	case LexicalAnalyse:
		shingleSize = 1
	case StructuralAnalyse:
		shingleSize = 3
	default:
		return 0, ErrModeDoesNotExist
	}

	text1 = c.clearText(text1)
	text2 = c.clearText(text2)

	if text1 == "" || text2 == "" {
		return 0, ErrTextIsEmpty
	}

	sliceText1 := strings.Fields(text1)
	sliceText2 := strings.Fields(text2)

	shinglesText1 := c.getShingles(sliceText1, shingleSize)
	shinglesText2 := c.getShingles(sliceText2, shingleSize)

	setShinglesText1 := set.New(shinglesText1)
	setShinglesText2 := set.New(shinglesText2)

	intersection := setShinglesText1.Intersection(setShinglesText2)
	union := setShinglesText1.Union(setShinglesText2)

	if union.GetLen() == 0 {
		return 0, ErrTextIsEmpty
	}

	similarity := intersection.GetLen() / union.GetLen()

	return math.Round(similarity * 100) / 100, nil
}

func (c *Core) clearText(text string) string {
	text = strings.ToLower(text)

	text = strings.ReplaceAll(text, ".", "")
	text = strings.ReplaceAll(text, ",", "")
	text = strings.ReplaceAll(text, ";", "")
	text = strings.ReplaceAll(text, ":", "")
	text = strings.ReplaceAll(text, "!", "")
	text = strings.ReplaceAll(text, "?", "")
	text = strings.ReplaceAll(text, "—", "")

	text = strings.Join(strings.Fields(text), " ")

	return text
}

func (c *Core) stem(word string) string {
    suffixes := []string{
        "ающими", "ающего", "ающему",
        "ованию", "ования",
        "ующего", "ующему",
        "ающей", "ающий", "ающих",
        "ований", "ованием",
        "ками", "гами", "нами", "рами", "тами",
        "ами", "ями",
        "ого", "ому", "ему",
        "ах", "ях", "ий", "ых", "ой", "ую", "юю",
        "ом", "ем", "ие", "ия", "ии",
        "ть", "ся", "сь",
        "ам", "ям",
        "ей", "ны", "ан",
        "ов", "ев",
        "ы", "и", "е", "а", "я",
    }
    for _, s := range suffixes {
        r := []rune(word)
        rs := []rune(s)
        if strings.HasSuffix(word, s) && len(r)-len(rs) > 3 {
            return string(r[:len(r)-len(rs)])
        }
    }
    return word
}

func (c *Core) getShingles(words []string, shingleSize int) []string {
	if len(words) < shingleSize {
		return []string{}
	}

	for i := range words {
    words[i] = c.stem(words[i])
	}

	shingles := make([]string, 0, len(words)-shingleSize+1)

	for i := 0; i <= len(words)-shingleSize; i++ {
		shingle := strings.Join(words[i:i+shingleSize], " ")
		shingles = append(shingles, shingle)
	}

	return shingles
}