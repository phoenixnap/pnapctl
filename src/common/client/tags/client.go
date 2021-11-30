package tags

import (
	"context"
	"net/http"

	tagapisdk "github.com/phoenixnap/go-sdk-bmc/tagapi"
	"golang.org/x/oauth2/clientcredentials"
	"phoenixnap.com/pnapctl/commands/version"
	configuration "phoenixnap.com/pnapctl/configs"
)

var Client TagSdkClient

type TagSdkClient interface {
	TagPost(tagCreate tagapisdk.TagCreate) (tagapisdk.Tag, *http.Response, error)
	TagsGet(name string) ([]tagapisdk.Tag, *http.Response, error)
	TagGetById(tagId string) (tagapisdk.Tag, *http.Response, error)
	TagDelete(tagId string) (tagapisdk.DeleteResult, *http.Response, error)
	TagPatch(tagId string, tagUpdate tagapisdk.TagUpdate) (tagapisdk.Tag, *http.Response, error)
}

type MainClient struct {
	TagSdkClient tagapisdk.TagsApi
}

func NewMainClient(clientId string, clientSecret string, customUrl string, customTokenURL string) TagSdkClient {
	tagConfiguration := tagapisdk.NewConfiguration()

	if customUrl != "" {
		tagConfiguration.Servers = tagapisdk.ServerConfigurations{
			{
				URL: customUrl,
			},
		}
	}

	tokenUrl := configuration.TokenURL
	if customTokenURL != "" {
		tokenUrl = customTokenURL
	}

	config := clientcredentials.Config{
		ClientID:     clientId,
		ClientSecret: clientSecret,
		TokenURL:     tokenUrl,
		Scopes:       []string{"tags", "tags.read"},
	}

	tagConfiguration.HTTPClient = config.Client(context.Background())
	tagConfiguration.UserAgent = configuration.UserAgent + version.AppVersion.Version

	api_client := tagapisdk.NewAPIClient(tagConfiguration)

	return MainClient{
		TagSdkClient: api_client.TagsApi,
	}
}

func (m MainClient) TagPost(tagCreate tagapisdk.TagCreate) (tagapisdk.Tag, *http.Response, error) {
	return m.TagSdkClient.TagsPost(context.Background()).TagCreate(tagCreate).Execute()
}

func (m MainClient) TagsGet(name string) ([]tagapisdk.Tag, *http.Response, error) {
	request := m.TagSdkClient.TagsGet(context.Background())

	if name != "" {
		request.Name(name)
	}

	return request.Execute()
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
