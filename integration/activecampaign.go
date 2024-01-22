package integration

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/jmoiron/sqlx"
)

type ActiveCampaignIntegration struct {
	APIKey   string
	APIURL   string
	TagsPath string
}

type ActiveCampaignTags struct {
	Tags []struct {
		ID  string `json:"id"`
		Tag string `json:"tag"`
	} `json:"tags"`
}

func (integration ActiveCampaignIntegration) URL() (string, error) {
	return integration.APIURL, nil
}

func (integration ActiveCampaignIntegration) Authenticate(request *http.Request) error {
	request.Header.Add("Api-Token", integration.APIKey)
	return nil
}

func (integration ActiveCampaignIntegration) GetTagsPath(path string) string {
	return fmt.Sprintf("%s/%s", path, integration.TagsPath)
}

func (integration ActiveCampaignIntegration) UnmarshalTags(data []byte) ([]Tag, error) {
	var activeCampaignTags ActiveCampaignTags
	err := json.Unmarshal(data, &activeCampaignTags)
	if err != nil {
		return []Tag{}, err
	}
	var tags []Tag
	for _, tag := range activeCampaignTags.Tags {
		// Redacted to record video, uncomment line
		// tags = append(tags, Tag{ID: tag.ID, Name: tag.Tag})
		tags = append(tags, Tag{ID: tag.ID, Name: "REDACTED"})
	}
	return tags, nil
}

func NewActiveCampaignIntegration(id int, conn *sqlx.DB) (*ActiveCampaignIntegration, error) {
	var apiKey string
	var apiURL string

	err := GetFieldValue(conn, id, "api_key", &apiKey)
	if err != nil {
		return nil, err
	}

	err = GetFieldValue(conn, id, "api_url", &apiURL)
	if err != nil {
		return nil, err
	}

	return &ActiveCampaignIntegration{
		APIURL:   apiURL,
		APIKey:   apiKey,
		TagsPath: "/api/3/tags",
	}, nil
}
