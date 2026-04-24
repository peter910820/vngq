package vndb

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"strings"
)

type UlistRequest struct {
	User    string `json:"user"`
	Fields  string `json:"fields"`
	Filters []any  `json:"filters,omitempty"`
	Sort    string `json:"sort,omitempty"`
	Reverse bool   `json:"reverse,omitempty"`
	Results int    `json:"results,omitempty"`
}

type UlistResponse struct {
	More    bool              `json:"more"`
	Results []UlistResultItem `json:"results"`
}

type UlistResultItem struct {
	Added int64         `json:"added"`
	ID    string        `json:"id"`
	VN    UlistResultVN `json:"vn"`
	Vote  *int          `json:"vote"`
}

type UlistResultVN struct {
	Title    string                  `json:"title"`
	Released string                  `json:"released"`
	Rating   *float64                `json:"rating"`
	Image    UlistResultVNImageBlock `json:"image"`
}

type UlistResultVNImageBlock struct {
	URL string `json:"url"`
}

func ulistRequestFactory(userID, fields string) *UlistRequest {
	return &UlistRequest{
		User:    userID,
		Fields:  fields,
		Reverse: true,
	}
}

func PostUlist(userID string) (*UlistResponse, error) {
	if !strings.HasPrefix(userID, "u") {
		return nil, fmt.Errorf("user id prefix error: prefix must be u")
	}

	jsonReq, err := json.Marshal(ulistRequestFactory(userID, makeUlistFields()))
	if err != nil {
		return nil, err
	}

	jsonResp, err := sendPostRequest("ulist", jsonReq)
	if err != nil {
		slog.Error(err.Error())
		return nil, err
	}

	var parsed UlistResponse
	if err := json.Unmarshal(jsonResp, &parsed); err != nil {
		return nil, err
	}

	slog.Info("ulist response", "results", len(parsed.Results), "more", parsed.More)
	return &parsed, nil
}

func makeUlistFields() string {
	// Keep payload small: only fetch fields required by current game room UI.
	return strings.Join([]string{
		"id",
		"added",
		"vote",
		"vn.id",
		"vn.title",
		"vn.released",
		"vn.rating",
		"vn.image.url",
	}, ",")
}
