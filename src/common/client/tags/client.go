package tags

import (
	"context"

	tagapisdk "github.com/phoenixnap/go-sdk-bmc/tagapi/v2"
	"golang.org/x/oauth2/clientcredentials"
	"phoenixnap.com/pnapctl/commands/version"
	"phoenixnap.com/pnapctl/common/client"
	configuration "phoenixnap.com/pnapctl/configs"
)

var Client TagSdkClient

type TagSdkClient interface {
	TagPost(tagCreate tagapisdk.TagCreate) (*tagapisdk.Tag, error)
	TagsGet(name string) ([]tagapisdk.Tag, error)
	TagGetById(tagId string) (*tagapisdk.Tag, error)
	TagDelete(tagId string) (*tagapisdk.DeleteResult, error)
	TagPatch(tagId string, tagUpdate tagapisdk.TagUpdate) (*tagapisdk.Tag, error)
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
	tagConfiguration.XPoweredBy = configuration.XPoweredBy + version.AppVersion.Version

	api_client := tagapisdk.NewAPIClient(tagConfiguration)

	return MainClient{
		TagSdkClient: api_client.TagsApi,
	}
}

func (m MainClient) TagPost(tagCreate tagapisdk.TagCreate) (*tagapisdk.Tag, error) {
	return client.HandleResponse(m.TagSdkClient.TagsPost(context.Background()).TagCreate(tagCreate).Execute())
}

func (m MainClient) TagsGet(name string) ([]tagapisdk.Tag, error) {
	request := m.TagSdkClient.TagsGet(context.Background())

	if name != "" {
		request = request.Name(name)
	}

	return client.HandleResponse(request.Execute())
}

func (m MainClient) TagGetById(tagId string) (*tagapisdk.Tag, error) {
	return client.HandleResponse(m.TagSdkClient.TagsTagIdGet(context.Background(), tagId).Execute())
}

func (m MainClient) TagDelete(tagId string) (*tagapisdk.DeleteResult, error) {
	return client.HandleResponse(m.TagSdkClient.TagsTagIdDelete(context.Background(), tagId).Execute())
}

func (m MainClient) TagPatch(tagId string, tagUpdate tagapisdk.TagUpdate) (*tagapisdk.Tag, error) {
	return client.HandleResponse(m.TagSdkClient.TagsTagIdPatch(context.Background(), tagId).TagUpdate(tagUpdate).Execute())
}
