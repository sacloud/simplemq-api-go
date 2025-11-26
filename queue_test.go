package simplemq_test

import (
	"context"
	"testing"

	"github.com/sacloud/packages-go/testutil"
	"github.com/sacloud/simplemq-api-go"
	"github.com/sacloud/simplemq-api-go/apis/v1/queue"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestQueueAPI(t *testing.T) {
	testutil.PreCheckEnvsFunc("SAKURACLOUD_ACCESS_TOKEN", "SAKURACLOUD_ACCESS_TOKEN_SECRET")(t)

	ctx := context.Background()

	client, err := simplemq.NewQueueClient()
	require.NoError(t, err)
	queueOp := simplemq.NewQueueOp(client)

	// CreateQueue
	resCreate, err := queueOp.Create(ctx, queue.CreateQueueRequest{
		CommonServiceItem: queue.CreateQueueRequestCommonServiceItem{
			Name:        "SDK-Test-Queue",
			Description: queue.NewOptString("SDK-Test-Queueの概要"),
		},
	})
	require.NoError(t, err)
	assert.Equal(t, "SDK-Test-Queue", resCreate.Status.QueueName)

	queueID := simplemq.GetQueueID(resCreate)

	// ListQueues
	resList, err := queueOp.List(ctx)
	assert.NoError(t, err)
	found := false
	for _, q := range resList {
		if queueID == simplemq.GetQueueID(&q) {
			found = true
			assert.Equal(t, "SDK-Test-Queue", resCreate.Status.QueueName)
		}
	}
	assert.True(t, found, "Created queue not found in list")

	// ConfigQueue
	resConfig, err := queueOp.Config(ctx, queueID, queue.ConfigQueueRequest{
		CommonServiceItem: queue.ConfigQueueRequestCommonServiceItem{
			Description: queue.NewOptString("SDK-Test-Queueの概要を変更"),
			Settings: queue.Settings{
				VisibilityTimeoutSeconds: 99,
				ExpireSeconds:            resCreate.Settings.ExpireSeconds,
			},
			Tags: []string{"tag1", "tag2"},
			Icon: queue.NewOptNilConfigQueueRequestCommonServiceItemIcon(queue.ConfigQueueRequestCommonServiceItemIcon{
				ID: queue.NewOptConfigQueueRequestCommonServiceItemIconID(queue.NewStringConfigQueueRequestCommonServiceItemIconID("112901627751")),
			}),
		},
	})
	assert.NoError(t, err)
	assert.Equal(t, queue.VisibilityTimeoutSeconds(99), resConfig.Settings.VisibilityTimeoutSeconds)

	// ReadQueue
	resRead, err := queueOp.Read(ctx, queueID)
	assert.NoError(t, err)
	assert.Equal(t, "SDK-Test-Queueの概要を変更", resRead.Description.Value)

	// GetMessageCount
	resMessageCount, err := queueOp.CountMessages(ctx, queueID)
	assert.NoError(t, err)
	assert.Equal(t, 0, resMessageCount)

	// RotateAPIKey
	_, err = queueOp.RotateAPIKey(ctx, queueID)
	assert.NoError(t, err)

	// ClearQueue
	err = queueOp.ClearMessages(ctx, queueID)
	assert.NoError(t, err)

	// DeleteQueue
	err = queueOp.Delete(ctx, queueID)
	require.NoError(t, err)
}
