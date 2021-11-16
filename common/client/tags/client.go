package tags

import (
	"context"
	"net/http"

	tagapisdk "gitlab.com/phoenixnap/bare-metal-cloud/go-sdk.git/tagapi"
	"golang.org/x/oauth2/clientcredentials"
	configuration "phoenixnap.com/pnap-cli/configs"
)

var Client TagSdkClient

type TagSdkClient interface {
	TagPost(tagCreate tagapisdk.TagCreate) (tagapisdk.Tag, *http.Response, error)
	TagsGet() ([]tagapisdk.Tag, *http.Response, error)
	TagGetById(tagId string) (tagapisdk.Tag, *http.Response, error)
	TagDelete(tagId string) (tagapisdk.DeleteResult, *http.Response, error)
	TagPatch(tagId string, tagUpdate tagapisdk.TagUpdate) (tagapisdk.Tag, *http.Response, error)
}

type MainClient struct {
	TagSdkClient tagapisdk.TagsApi
}

func NewMainClient(clientId string, clientSecret string) TagSdkClient {
	rancherConfiguration := tagapisdk.NewConfiguration()

	config := clientcredentials.Config{
		ClientID:     clientId,
		ClientSecret: clientSecret,
		TokenURL:     configuration.TokenURL,
		Scopes:       []string{"bmc", "bmc.read"},
	}

	rancherConfiguration.HTTPClient = config.Client(context.Background())

	api_client := tagapisdk.NewAPIClient(rancherConfiguration)

	return MainClient{
		TagSdkClient: api_client.TagsApi,
	}
}

func (m MainClient) TagPost(tagCreate tagapisdk.TagCreate) (tagapisdk.Tag, *http.Response, error) {
	return m.TagSdkClient.TagsPost(context.Background()).TagCreate(tagCreate).Execute()
}

func (m MainClient) TagsGet() ([]tagapisdk.Tag, *http.Response, error) {
	return m.TagSdkClient.TagsGet(context.Background()).Execute()
}

func (m MainClient) TagGetById(tagId string) (tagapisdk.Tag, *http.Response, error) {
	return m.TagSdkClient.TagsTagIdGet(context.Background(), tagId).Execute()
}

func (m MainClient) TagDelete(tagId string) (tagapisdk.DeleteResult, *http.Response, error) {
	return m.TagSdkClient.TagsTagIdDelete(context.Background(), tagId).Execute()
}

func (m MainClient) TagPatch(tagId string, tagUpdate tagapisdk.TagUpdate) (tagapisdk.Tag, *http.Response, error) {
	return m.TagSdkClient.TagsTagIdPatch(context.Background(), tagId).TagUpdate(tagUpdate).Execute()
}
