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
	"errors"
	"strconv"

	"github.com/sacloud/simplemq-api-go/apis/v1/queue"
)

type QueueAPI interface {
	List(context.Context) ([]queue.CommonServiceItem, error)
	Get(_ context.Context, id string) (queue.CommonServiceItem, error)
	Create(context.Context, CreateQueueRequest) (queue.CommonServiceItem, error)
	Config(_ context.Context, id string, req queue.ConfigQueueRequest) (queue.CommonServiceItem, error)
	Delete(_ context.Context, id string) (queue.CommonServiceItem, error)

	CountMessages(_ context.Context, id string) (int, error)
	RotateAPIKey(_ context.Context, id string) (string, error)
	ClearMessages(_ context.Context, id string) error
}

var _ QueueAPI = (*queueOp)(nil)

type queueOp struct {
	client *queue.Client
}

func NewQueueOp(client *queue.Client) QueueAPI {
	return &queueOp{client: client}
}

func GetQueueID(q queue.CommonServiceItem) string {
	if q.ID.IsString() {
		return q.ID.String
	}
	return strconv.Itoa(q.ID.Int)
}

type CreateQueueRequest struct {
	QueueName   string
	Description string
}

func (op *queueOp) Create(ctx context.Context, req CreateQueueRequest) (queue.CommonServiceItem, error) {
	var empty queue.CommonServiceItem
	res, err := op.client.CreateQueue(ctx, &queue.CreateQueueRequest{
		CommonServiceItem: queue.CreateQueueRequestCommonServiceItem{
			Name:        queue.QueueName(req.QueueName),
			Description: queue.NewOptString(req.Description),
			Provider: queue.CreateQueueRequestCommonServiceItemProvider{
				Class: queue.NewOptCreateQueueRequestCommonServiceItemProviderClass(queue.CreateQueueRequestCommonServiceItemProviderClassSimplemq),
			},
		},
	})
	if err != nil {
		return empty, err
	}

	switch r := res.(type) {
	case *queue.CreateQueueCreated:
		return r.CommonServiceItem, nil
	case *queue.CreateQueueUnauthorized:
		return empty, errors.New(r.ErrorMsg.Value)
	case *queue.CreateQueueBadRequest:
		return empty, errors.New(r.ErrorMsg.Value)
	case *queue.CreateQueueConflict:
		return empty, errors.New(r.ErrorMsg.Value)
	case *queue.CreateQueueInternalServerError:
		return empty, errors.New(r.ErrorMsg.Value)
	default:
		return empty, errors.New("unknown error")
	}
}

func (op *queueOp) List(ctx context.Context) ([]queue.CommonServiceItem, error) {
	res, err := op.client.GetQueues(ctx)
	if err != nil {
		return nil, err
	}

	switch r := res.(type) {
	case *queue.GetQueuesOK:
		return r.CommonServiceItems, nil
	case *queue.GetQueuesUnauthorized:
		return nil, errors.New(r.ErrorMsg.Value)
	case *queue.GetQueuesBadRequest:
		return nil, errors.New(r.ErrorMsg.Value)
	case *queue.GetQueuesInternalServerError:
		return nil, errors.New(r.ErrorMsg.Value)
	default:
		return nil, errors.New("unknown error")
	}
}

func (op *queueOp) Get(ctx context.Context, id string) (queue.CommonServiceItem, error) {
	var empty queue.CommonServiceItem
	res, err := op.client.GetQueue(ctx, queue.GetQueueParams{ID: id})
	if err != nil {
		return empty, err
	}

	switch r := res.(type) {
	case *queue.GetQueueOK:
		return r.CommonServiceItem, nil
	case *queue.GetQueueUnauthorized:
		return empty, errors.New(r.ErrorMsg.Value)
	case *queue.GetQueueBadRequest:
		return empty, errors.New(r.ErrorMsg.Value)
	case *queue.GetQueueNotFound:
		return empty, errors.New(r.ErrorMsg.Value)
	case *queue.GetQueueInternalServerError:
		return empty, errors.New(r.ErrorMsg.Value)
	default:
		return empty, errors.New("unknown error")
	}
}

func (op *queueOp) Config(ctx context.Context, id string, req queue.ConfigQueueRequest) (queue.CommonServiceItem, error) {
	var empty queue.CommonServiceItem
	res, err := op.client.ConfigQueue(ctx, &req, queue.ConfigQueueParams{ID: id})
	if err != nil {
		return empty, err
	}

	switch r := res.(type) {
	case *queue.ConfigQueueOK:
		return r.CommonServiceItem, nil
	case *queue.ConfigQueueUnauthorized:
		return empty, errors.New(r.ErrorMsg.Value)
	case *queue.ConfigQueueBadRequest:
		return empty, errors.New(r.ErrorMsg.Value)
	case *queue.ConfigQueueNotFound:
		return empty, errors.New(r.ErrorMsg.Value)
	case *queue.ConfigQueueInternalServerError:
		return empty, errors.New(r.ErrorMsg.Value)
	default:
		return empty, errors.New("unknown error")
	}
}

func (op *queueOp) Delete(ctx context.Context, id string) (queue.CommonServiceItem, error) {
	var empty queue.CommonServiceItem
	res, err := op.client.DeleteQueue(ctx, queue.DeleteQueueParams{
		ID: id,
	})
	if err != nil {
		return empty, err
	}

	switch r := res.(type) {
	case *queue.DeleteQueueOK:
		return r.CommonServiceItem, nil
	case *queue.DeleteQueueUnauthorized:
		return empty, errors.New(r.ErrorMsg.Value)
	case *queue.DeleteQueueBadRequest:
		return empty, errors.New(r.ErrorMsg.Value)
	case *queue.DeleteQueueNotFound:
		return empty, errors.New(r.ErrorMsg.Value)
	case *queue.DeleteQueueInternalServerError:
		return empty, errors.New(r.ErrorMsg.Value)
	default:
		return empty, errors.New("unknown error")
	}
}

func (op *queueOp) CountMessages(ctx context.Context, id string) (int, error) {
	res, err := op.client.GetMessageCount(ctx, queue.GetMessageCountParams{ID: id})
	if err != nil {
		return 0, err
	}

	switch r := res.(type) {
	case *queue.GetMessageCountOK:
		return r.SimpleMQ.GetCount(), nil
	case *queue.GetMessageCountUnauthorized:
		return 0, errors.New(r.ErrorMsg.Value)
	case *queue.GetMessageCountBadRequest:
		return 0, errors.New(r.ErrorMsg.Value)
	case *queue.GetMessageCountNotFound:
		return 0, errors.New(r.ErrorMsg.Value)
	case *queue.GetMessageCountInternalServerError:
		return 0, errors.New(r.ErrorMsg.Value)
	default:
		return 0, errors.New("unknown error")
	}
}

func (op *queueOp) RotateAPIKey(ctx context.Context, id string) (string, error) {
	res, err := op.client.RotateAPIKey(ctx, queue.RotateAPIKeyParams{ID: id})
	if err != nil {
		return "", err
	}

	switch r := res.(type) {
	case *queue.RotateAPIKeyOK:
		return r.SimpleMQ.GetApikey(), nil
	case *queue.RotateAPIKeyUnauthorized:
		return "", errors.New(r.ErrorMsg.Value)
	case *queue.RotateAPIKeyBadRequest:
		return "", errors.New(r.ErrorMsg.Value)
	case *queue.RotateAPIKeyNotFound:
		return "", errors.New(r.ErrorMsg.Value)
	case *queue.RotateAPIKeyInternalServerError:
		return "", errors.New(r.ErrorMsg.Value)
	default:
		return "", errors.New("unknown error")
	}
}

func (op *queueOp) ClearMessages(ctx context.Context, id string) error {
	res, err := op.client.ClearQueue(ctx, queue.ClearQueueParams{ID: id})
	if err != nil {
		return err
	}

	switch r := res.(type) {
	case *queue.ClearQueueOK:
		return nil
	case *queue.ClearQueueUnauthorized:
		return errors.New(r.ErrorMsg.Value)
	case *queue.ClearQueueBadRequest:
		return errors.New(r.ErrorMsg.Value)
	case *queue.ClearQueueNotFound:
		return errors.New(r.ErrorMsg.Value)
	case *queue.ClearQueueInternalServerError:
		return errors.New(r.ErrorMsg.Value)
	default:
		return errors.New("unknown error")
	}
}
