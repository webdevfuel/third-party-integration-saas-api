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

func (integration ConvertKitIntegration) URL() (string, error) {
	return integration.APIURL, nil
}

func (integration ConvertKitIntegration) Authenticate(request *http.Request) error {
	query := request.URL.Query()
	query.Set("api_key", integration.APIKey)
	request.URL.RawQuery = query.Encode()
	return nil
}

func (integration ConvertKitIntegration) GetTagsPath(path string) string {
	return fmt.Sprintf("%s/%s", path, integration.TagsPath)
}

func (integration ConvertKitIntegration) UnmarshalTags(data []byte) ([]Tag, error) {
	var convertKitTags ConvertKitTags
	err := json.Unmarshal(data, &convertKitTags)
	if err != nil {
		return []Tag{}, err
	}
	var tags []Tag
	for _, tag := range convertKitTags.Tags {
		tags = append(tags, Tag{ID: fmt.Sprintf("%d", tag.ID), Name: tag.Name})
	}
	return tags, nil
}

func NewConvertKitIntegration(id int, conn *sqlx.DB) (*ConvertKitIntegration, error) {
	var apiKey string
	var apiSecret string

	err := getFieldValue(conn, id, "api_key", &apiKey)
	if err != nil {
		return nil, err
	}

	err = getFieldValue(conn, id, "api_secret", &apiSecret)
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
