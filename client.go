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
	"runtime"

	client "github.com/sacloud/api-client-go"
	"github.com/sacloud/saclient-go"
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

func NewQueueClient(client saclient.ClientAPI) (*queue.Client, error) {
	return NewQueueClientWithApiUrl(DefaultQueueAPIRootURL, client)
}

func NewQueueClientWithApiUrl(apiUrl string, client saclient.ClientAPI) (*queue.Client, error) {
	var ss DummySecuritySource

	if dupable, ok := client.(saclient.ClientOptionAPI); !ok {
		return nil, NewError("client does not implement saclient.ClientOptionAPI", nil)
	} else if augmented, err := dupable.DupWith(
		saclient.WithUserAgent(UserAgent),
		// これはなにか:
		// DummySecuritySource.ApiKeyAuth()がBasic認証を生成
		// しかし実際の通信で必ずしもBasic認証が使われると限らない
		//　そのあたりをsaclient-go側で吸収させる設定が下記↓
		saclient.WithForceAutomaticAuthentication(),
	); err != nil {
		return nil, err
	} else {
		return queue.NewClient(apiUrl, ss, queue.WithClient(augmented))
	}
}

// DefaultMessageAPIRootURL デフォルトのMessage APIルートURL
const DefaultMessageAPIRootURL = "https://simplemq.tk1b.api.sacloud.jp"

type ApiKeySecuritySource struct {
	Token string
}

func (ss ApiKeySecuritySource) ApiKeyAuth(ctx context.Context, operationName message.OperationName) (message.ApiKeyAuth, error) {
	return message.ApiKeyAuth{Token: ss.Token}, nil
}

func NewMessageClient(apiKey string, client saclient.ClientAPI) (*message.Client, error) {
	return NewMessageClientWithApiUrl(DefaultMessageAPIRootURL, apiKey, client)
}

func NewMessageClientWithApiUrl(apiUrl, apiKey string, client saclient.ClientAPI) (*message.Client, error) {
	var ss ApiKeySecuritySource

	if dupable, ok := client.(saclient.ClientOptionAPI); !ok {
		return nil, NewError("client does not implement saclient.ClientOptionAPI", nil)
	} else if augmented, err := dupable.DupWith(
		saclient.WithUserAgent(UserAgent),
		// これはなにか:
		// ApiKeySecuritySource.ApiKeyAuth()がBasic認証を生成
		// しかし実際の通信で必要なのはBearer認証
		//　そのあたりをsaclient-go側で吸収させる設定が下記↓
		saclient.WithBearerToken(apiKey),
	); err != nil {
		return nil, err
	} else {
		return message.NewClient(apiUrl, ss, message.WithClient(augmented))
	}
}
