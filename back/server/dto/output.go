package dto

import (
	"antiplagiarism/core"
	"encoding/json"
	"time"
)

type SimilarityDTO struct {
	Similarity float64 `json:"similarity"`
}

func NewSimilarityDTO(similarity float64) SimilarityDTO {
	return SimilarityDTO{
		Similarity: similarity,
	}
}

type ErrDTO struct {
	Text string `json:"text"`
	Time time.Time `json:"time"`
}

func NewErrDTO(err error) ErrDTO {
	return ErrDTO{
		Text: err.Error(),
		Time: time.Now(),
	}
}

func (e *ErrDTO) ToString() string {
	bin, err := json.MarshalIndent(e, "", "	")
	if err != nil {
		panic(err)
	}
	
	return string(bin)
}

type GeneralAnalyseResultDTO struct {
	General float64 `json:"general"`
	Structural float64 `json:"structural"`
	Lexical float64 `json:"lexical"`
}

func NewGeneralAnalyseResultDTO(generalAnalyseResult core.GeneralAnalyseResult) GeneralAnalyseResultDTO {
	return GeneralAnalyseResultDTO{
		General: generalAnalyseResult.General,
		Structural: generalAnalyseResult.Structural,
		Lexical: generalAnalyseResult.Lexical,
	}
}