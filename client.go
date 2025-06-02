// Copyright 2025- The sacloud/simplemq-api-go authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package simplemq

import (
	"context"
	"fmt"
	"net/http"
	"runtime"

	client "github.com/sacloud/api-client-go"
	sacloudhttp "github.com/sacloud/go-http"
	"github.com/sacloud/simplemq-api-go/apis/v1/message"
	"github.com/sacloud/simplemq-api-go/apis/v1/queue"
)

// DefaultQueueAPIRootURL デフォルトのQueue APIルートURL
const DefaultQueueAPIRootURL = "https://secure.sakura.ad.jp/cloud/zone/tk1b/api/cloud/1.1"

// UserAgent APIリクエスト時のユーザーエージェント
var UserAgent = fmt.Sprintf(
	"simplemq-api-go/%s (%s/%s; +https://github.com/sacloud/simplemq-api-go) %s",
	Version,
	runtime.GOOS,
	runtime.GOARCH,
	client.DefaultUserAgent,
)

// SecuritySourceはOpenAPI定義で使用されている認証のための仕組み。api-client-goが処理するので、ogen用はダミーで誤魔化す
type DummySecuritySource struct {
	Token string
}

func (ss DummySecuritySource) ApiKeyAuth(ctx context.Context, operationName queue.OperationName) (queue.ApiKeyAuth, error) {
	return queue.ApiKeyAuth{Username: ss.Token}, nil
}

func NewQueueClient(params ...client.ClientParam) (*queue.Client, error) {
	return NewQueueClientWithApiUrl(DefaultQueueAPIRootURL, params...)
}

func NewQueueClientWithApiUrl(apiUrl string, params ...client.ClientParam) (*queue.Client, error) {
	c, err := client.NewClient(apiUrl, append([]client.ClientParam{
		client.WithUserAgent(UserAgent),
	}, params...)...)
	if err != nil {
		return nil, err
	}

	v1Client, err := queue.NewClient(c.ServerURL(), DummySecuritySource{Token: "simplemq-client"}, queue.WithClient(c.NewHttpRequestDoer()))
	if err != nil {
		return nil, fmt.Errorf("create client: %w", err)
	}

	return v1Client, nil
}

// DefaultMessageAPIRootURL デフォルトのMessage APIルートURL
const DefaultMessageAPIRootURL = "https://simplemq.tk1b.api.sacloud.jp"

type ApiKeySecuritySource struct {
	Token string
}

func (ss ApiKeySecuritySource) ApiKeyAuth(ctx context.Context, operationName message.OperationName) (message.ApiKeyAuth, error) {
	return message.ApiKeyAuth{Token: ss.Token}, nil
}

func NewMessageClient(apiKey string, params ...client.ClientParam) (*message.Client, error) {
	return NewMessageClientWithApiUrl(DefaultMessageAPIRootURL, apiKey, params...)
}

func NewMessageClientWithApiUrl(apiUrl, apiKey string, params ...client.ClientParam) (*message.Client, error) {
	c, err := client.NewClient(apiUrl, append([]client.ClientParam{
		client.WithUserAgent(UserAgent),
		client.WithOptions(&client.Options{
			RequestCustomizers: []sacloudhttp.RequestCustomizer{
				func(req *http.Request) error {
					req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", apiKey))
					return nil
				},
			}}),
	},
		params...)...)
	if err != nil {
		return nil, err
	}

	v1Client, err := message.NewClient(apiUrl, ApiKeySecuritySource{Token: apiKey}, message.WithClient(c.NewHttpRequestDoer()))
	if err != nil {
		return nil, fmt.Errorf("create client: %w", err)
	}

	return v1Client, nil
}
