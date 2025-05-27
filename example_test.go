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

package simplemq_test

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/sacloud/simplemq-api-go"
	"github.com/sacloud/simplemq-api-go/apis/v1/queue"
)

var requriedEnvs = []string{
	"SAKURACLOUD_ACCESS_TOKEN",
	"SAKURACLOUD_ACCESS_TOKEN_SECRET",
}

func checkEnvs() {
	for _, env := range requriedEnvs {
		if os.Getenv(env) == "" {
			panic(env + " is not set")
		}
	}
}

func ExampleQueueAPI() {
	checkEnvs()
	ctx := context.Background()

	client, err := simplemq.NewQueueClient()
	if err != nil {
		panic(err)
	}
	queueOp := simplemq.NewQueueOp(client)

	// CreateQueue
	resCreate, err := queueOp.Create(ctx, simplemq.CreateQueueRequest{
		QueueName:   "SDK-Test-Queue",
		Description: "SDK-Test-Queueの概要",
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(resCreate.Status.QueueName)
	queueID := simplemq.GetQueueID(*resCreate)

	// ListQueues
	resList, err := queueOp.List(ctx)
	if err != nil {
		panic(err)
	}
	for _, q := range resList {
		if queueID == simplemq.GetQueueID(q) {
			fmt.Println(resCreate.Status.QueueName)
		}
	}

	// ConfigQueue
	resConfig, err := queueOp.Config(ctx, queueID, queue.ConfigQueueRequest{
		CommonServiceItem: queue.ConfigQueueRequestCommonServiceItem{
			Description: queue.NewOptString("SDK-Test-Queueの概要を変更"),
			Settings: queue.ConfigQueueRequestCommonServiceItemSettings{
				VisibilityTimeoutSeconds: 99,
				ExpireSeconds:            resCreate.Settings.ExpireSeconds, // NOTE: VisibilityTimeoutSecondsのみを変更
			},
		},
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(resConfig.Settings.VisibilityTimeoutSeconds)

	// GetQueue
	resGet, err := queueOp.Get(ctx, queueID)
	if err != nil {
		panic(err)
	}
	fmt.Println(resGet.Description.Value.String)

	// GetMessageCount
	resMessageCount, err := queueOp.CountMessages(ctx, queueID)
	if err != nil {
		panic(err)
	}
	fmt.Println(resMessageCount)

	// RotateAPIKey
	if _, err := queueOp.RotateAPIKey(ctx, queueID); err != nil {
		panic(err)
	}

	// ClearQueue
	if err := queueOp.ClearMessages(ctx, queueID); err != nil {
		panic(err)
	}

	// DeleteQueue
	if err := queueOp.Delete(ctx, queueID); err != nil {
		panic(err)
	}

	// Output:
	// SDK-Test-Queue
	// SDK-Test-Queue
	// 99
	// SDK-Test-Queueの概要を変更
	// 0
}

func ExampleMessageAPI() {
	checkEnvs()
	ctx := context.Background()

	client, err := simplemq.NewQueueClient()
	if err != nil {
		panic(err)
	}
	queueOp := simplemq.NewQueueOp(client)

	resCreate, err := queueOp.Create(ctx, simplemq.CreateQueueRequest{
		QueueName: "SDK-Test-Queue",
	})
	if err != nil {
		panic(err)
	}
	// teardown
	defer func() {
		if err := queueOp.Delete(ctx, simplemq.GetQueueID(*resCreate)); err != nil {
			panic(err)
		}
	}()
	queueName := simplemq.GetQueueName(*resCreate)
	apiKey, err := queueOp.RotateAPIKey(ctx, simplemq.GetQueueID(*resCreate))
	if err != nil {
		panic(err)
	}

	messageClient, err := simplemq.NewMessageClient(queueName, apiKey)
	if err != nil {
		panic(err)
	}

	// SendMessage
	resSend, err := messageClient.Send(ctx, "HelloSimpleMQ")
	if err != nil {
		panic(err)
	}
	fmt.Println(resSend.Content)

	// ReceiveMessage
	resReceive, err := messageClient.Receive(ctx)
	if err != nil {
		panic(err)
	}
	fmt.Println(resReceive[0].Content)
	messageID := string(resReceive[0].ID)

	// ExtendMessageTimeout
	resExtend, err := messageClient.ExtendTimeout(ctx, messageID)
	if err != nil {
		panic(err)
	}
	timeBefore := time.UnixMilli(resReceive[0].VisibilityTimeoutAt)
	timeExtended := time.UnixMilli(resExtend.VisibilityTimeoutAt)
	fmt.Println(timeExtended.After(timeBefore))

	// DeleteMessage
	if err := messageClient.Delete(ctx, messageID); err != nil {
		panic(err)
	}

	// Output:
	// HelloSimpleMQ
	// HelloSimpleMQ
	// true
}
