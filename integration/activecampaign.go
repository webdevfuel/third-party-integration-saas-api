package integration

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
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
		tags = append(tags, Tag{ID: tag.ID, Name: tag.Tag})
	}
	return tags, nil
}

func NewActiveCampaignIntegration() *ActiveCampaignIntegration {
	return &ActiveCampaignIntegration{
		APIURL:   os.Getenv("ACTIVECAMPAIGN_API_URL"),
		APIKey:   os.Getenv("ACTIVECAMPAIGN_API_KEY"),
		TagsPath: "/api/3/tags",
	}
}
