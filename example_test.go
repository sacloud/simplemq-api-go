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
	"strconv"
	"time"

	"github.com/sacloud/simplemq-api-go"
	"github.com/sacloud/simplemq-api-go/apis/v1/sacloud"
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

func getQueueID(q sacloud.CommonServiceItem) string {
	if q.ID.IsString() {
		return q.ID.String
	}
	return strconv.Itoa(q.ID.Int)
}

func ExampleSacloudAPI() {
	checkEnvs()
	ctx := context.Background()

	client, err := simplemq.NewSacloudClient()
	if err != nil {
		panic(err)
	}
	sacloudOp := simplemq.NewSacloudOp(client)

	// CreateQueue
	resCreate, err := sacloudOp.CreateQueue(ctx, simplemq.CreateQueueRequest{
		QueueName:   "SDK-Test-Queue",
		Description: "SDK-Test-Queueの概要",
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(resCreate.Status.QueueName)
	queueID := getQueueID(resCreate)

	// ListQueues
	resList, err := sacloudOp.ListQueues(ctx)
	if err != nil {
		panic(err)
	}
	for _, q := range resList {
		if queueID == getQueueID(q) {
			fmt.Println(resCreate.Status.QueueName)
		}
	}

	// ConfigQueue
	resConfig, err := sacloudOp.ConfigQueue(ctx, queueID, sacloud.ConfigQueueRequest{
		CommonServiceItem: sacloud.ConfigQueueRequestCommonServiceItem{
			Description: sacloud.NewOptString("SDK-Test-Queueの概要を変更"),
			Settings: sacloud.ConfigQueueRequestCommonServiceItemSettings{
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
	resGet, err := sacloudOp.GetQueue(ctx, queueID)
	if err != nil {
		panic(err)
	}
	fmt.Println(resGet.Description.Value.String)

	// GetMessageCount
	resMessageCount, err := sacloudOp.GetMessageCount(ctx, queueID)
	if err != nil {
		panic(err)
	}
	fmt.Println(resMessageCount)

	// RotateAPIKey
	if _, err := sacloudOp.RotateAPIKey(ctx, queueID); err != nil {
		panic(err)
	}

	// ClearQueue
	if err := sacloudOp.ClearQueue(ctx, queueID); err != nil {
		panic(err)
	}

	// DeleteQueue
	resDelete, err := sacloudOp.DeleteQueue(ctx, queueID)
	if err != nil {
		panic(err)
	}
	fmt.Println(resDelete.Availability)

	// Output:
	// SDK-Test-Queue
	// SDK-Test-Queue
	// 99
	// SDK-Test-Queueの概要を変更
	// 0
	// discontinued
}

func ExampleSimpleMQAPI() {
	checkEnvs()
	ctx := context.Background()

	client, err := simplemq.NewSacloudClient()
	if err != nil {
		panic(err)
	}
	sacloudOp := simplemq.NewSacloudOp(client)

	resCreate, err := sacloudOp.CreateQueue(ctx, simplemq.CreateQueueRequest{
		QueueName: "SDK-Test-Queue",
	})
	if err != nil {
		panic(err)
	}
	// teardown
	defer func() {
		if _, err := sacloudOp.DeleteQueue(ctx, getQueueID(resCreate)); err != nil {
			panic(err)
		}
	}()
	queueName := resCreate.Status.QueueName
	apiKey, err := sacloudOp.RotateAPIKey(ctx, getQueueID(resCreate))
	if err != nil {
		panic(err)
	}

	qClient, err := simplemq.NewSimpleMQClient(queueName, apiKey)
	if err != nil {
		panic(err)
	}

	// SendMessage
	resSend, err := qClient.SendMessage(ctx, "HelloSimpleMQ")
	if err != nil {
		panic(err)
	}
	fmt.Println(resSend.Content)

	// ReceiveMessage
	resReceive, err := qClient.ReceiveMessage(ctx)
	if err != nil {
		panic(err)
	}
	fmt.Println(resReceive[0].Content)
	messageID := string(resReceive[0].ID)

	// ExtendMessageTimeout
	resExtend, err := qClient.ExtendMessageTimeout(ctx, messageID)
	if err != nil {
		panic(err)
	}
	timeBefore := time.UnixMilli(resReceive[0].VisibilityTimeoutAt)
	timeExtended := time.UnixMilli(resExtend.VisibilityTimeoutAt)
	fmt.Println(timeExtended.After(timeBefore))

	// DeleteMessage
	if err := qClient.DeleteMessage(ctx, messageID); err != nil {
		panic(err)
	}

	// Output:
	// HelloSimpleMQ
	// HelloSimpleMQ
	// true
}
