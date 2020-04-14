package api

import (
	"encoding/json"
	"net/http"
	"pigLatin/util"
)

// handlePigLatinTranslation handles the api /piglatin POST route.
func (s *API) handlePigLatinTranslation() http.HandlerFunc {

	type RequestBody struct {
		Text string `json:"text"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		var reqBody RequestBody

		if err := json.NewDecoder(r.Body).Decode(&reqBody); err != nil {
			util.ErrorResp(w, err)
			return
		}

		result, err := s.piglatinService.FormatText(reqBody.Text)

		if err != nil {
			util.ErrorResp(w, err)
			return
		}

		util.SuccessResp(w, result)
	}
}
