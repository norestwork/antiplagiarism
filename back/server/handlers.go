package server

import (
	"antiplagiarism/core"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"antiplagiarism/server/dto"
)

type Handlers struct {
	core core.Core
}

func NewHandlers(core core.Core) Handlers {
	return Handlers{
		core: core,
	}
}

func (h *Handlers) HandleAnalyse(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	
	var input dto.TextsDTO

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		h.writeError(w, err, http.StatusBadRequest)
		return
	}

	mode := core.AnalyseMode(input.Mode)

	if mode == "" || mode == "general" {
		h.writeGeneralAnalyseResult(w, input.Text1, input.Text2)
		return
	}

	h.writeAnalyseResult(w, input.Text1, input.Text2, mode)
}

func (h *Handlers) writeGeneralAnalyseResult(w http.ResponseWriter, text1 string, text2 string) {
		generalAnalyseResult := h.core.GeneralAnalyse(text1, text2)
		generalAnalyseResultDTO := dto.NewGeneralAnalyseResultDTO(generalAnalyseResult)

		bin, err := json.Marshal(generalAnalyseResultDTO)
		if err != nil {
			h.writeError(w, err, http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		if _, err := w.Write(bin); err != nil {
			fmt.Println("Failed to write HTTP response:", err)
			return
		}
}

func (h *Handlers) writeAnalyseResult(w http.ResponseWriter, text1 string, text2 string, mode core.AnalyseMode) {
	similarity, err := h.core.Analyse(text1, text2, mode)
	if err != nil {
		if errors.Is(err, core.ErrModeDoesNotExist) || errors.Is(err, core.ErrTextIsEmpty) {
			h.writeError(w, err, http.StatusBadRequest)
		} else {
			h.writeError(w, err, http.StatusInternalServerError)
		}

		return
	}

	similarityDTO := dto.NewSimilarityDTO(similarity)

	bin, err := json.Marshal(similarityDTO)
	if err != nil {
		h.writeError(w, err, http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if _, err := w.Write(bin); err != nil {
		fmt.Println("Failed to write HTTP response:", err)
		return
	}
}

func (h *Handlers) writeError(w http.ResponseWriter, err error, status int) {
	errDTO := dto.NewErrDTO(err)

	bin, err := json.Marshal(errDTO)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Println("JSON line feed error:", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	
	if _, err = w.Write(bin); err != nil {
		fmt.Println("Failed to write HTTP response:", err)
		return
	}
}