package integration

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/jmoiron/sqlx"
)

type ConvertKitIntegration struct {
	APIKey    string
	APISecret string
	APIURL    string
	TagsPath  string
}

type ConvertKitTags struct {
	Tags []struct {
		ID   int32  `json:"id"`
		Name string `json:"name"`
	} `json:"tags"`
}

func (integration ConvertKitIntegration) URL() string {
	return integration.APIURL
}

func (integration ConvertKitIntegration) Authenticate(request *http.Request) {
	query := request.URL.Query()
	query.Set("api_key", integration.APIKey)
	request.URL.RawQuery = query.Encode()
}

func (integration ConvertKitIntegration) GetTagsPath() string {
	return integration.TagsPath
}

func (integration ConvertKitIntegration) UnmarshalTags(data []byte) ([]Tag, error) {
	var convertKitTags ConvertKitTags
	err := json.Unmarshal(data, &convertKitTags)
	if err != nil {
		return []Tag{}, err
	}
	var tags []Tag
	for _, tag := range convertKitTags.Tags {
		// Redacted to record video, uncomment line
		// tags = append(tags, Tag{ID: fmt.Sprintf("%d", tag.ID), Name: tag.Name})
		tags = append(tags, Tag{ID: fmt.Sprintf("%d", tag.ID), Name: "REDACTED"})
	}
	return tags, nil
}

func NewConvertKitIntegration(id int, conn *sqlx.DB) (*ConvertKitIntegration, error) {
	var apiKey string
	var apiSecret string

	err := GetFieldValue(conn, id, "api_key", &apiKey)
	if err != nil {
		return nil, err
	}

	err = GetFieldValue(conn, id, "api_secret", &apiSecret)
	if err != nil {
		return nil, err
	}

	return &ConvertKitIntegration{
		APIKey:    apiKey,
		APISecret: apiSecret,
		APIURL:    "https://api.convertkit.com",
		TagsPath:  "/v3/tags",
	}, nil
}
