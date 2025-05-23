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

	"github.com/sacloud/simplemq-api-go/apis/v1/sacloud"
)

type SacloudAPI interface {
	ListQueues(context.Context) ([]sacloud.CommonServiceItem, error)
	GetQueue(_ context.Context, id string) (sacloud.CommonServiceItem, error)
	CreateQueue(context.Context, CreateQueueRequest) (sacloud.CommonServiceItem, error)
	ConfigQueue(_ context.Context, id string, req sacloud.ConfigQueueRequest) (sacloud.CommonServiceItem, error)
	DeleteQueue(_ context.Context, id string) (sacloud.CommonServiceItem, error)

	GetMessageCount(_ context.Context, id string) (int, error)
	RotateAPIKey(_ context.Context, id string) (string, error)
	ClearQueue(_ context.Context, id string) error
}

var _ SacloudAPI = (*sacloudOp)(nil)

type sacloudOp struct {
	client *sacloud.Client
}

func NewSacloudOp(client *sacloud.Client) SacloudAPI {
	return &sacloudOp{client: client}
}

type CreateQueueRequest struct {
	QueueName   string
	Description string
}

func (op *sacloudOp) CreateQueue(ctx context.Context, req CreateQueueRequest) (sacloud.CommonServiceItem, error) {
	var empty sacloud.CommonServiceItem
	res, err := op.client.CreateQueue(ctx, &sacloud.CreateQueueRequest{
		CommonServiceItem: sacloud.CreateQueueRequestCommonServiceItem{
			Name:        sacloud.QueueName(req.QueueName),
			Description: sacloud.NewOptString(req.Description),
			Provider: sacloud.CreateQueueRequestCommonServiceItemProvider{
				Class: sacloud.NewOptCreateQueueRequestCommonServiceItemProviderClass(sacloud.CreateQueueRequestCommonServiceItemProviderClassSimplemq),
			},
		},
	})
	if err != nil {
		return empty, err
	}

	switch r := res.(type) {
	case *sacloud.CreateQueueCreated:
		return r.CommonServiceItem, nil
	case *sacloud.CreateQueueUnauthorized:
		return empty, errors.New(r.ErrorMsg.Value)
	case *sacloud.CreateQueueBadRequest:
		return empty, errors.New(r.ErrorMsg.Value)
	case *sacloud.CreateQueueConflict:
		return empty, errors.New(r.ErrorMsg.Value)
	case *sacloud.CreateQueueInternalServerError:
		return empty, errors.New(r.ErrorMsg.Value)
	default:
		return empty, errors.New("unknown error")
	}
}

func (op *sacloudOp) ListQueues(ctx context.Context) ([]sacloud.CommonServiceItem, error) {
	res, err := op.client.GetQueues(ctx)
	if err != nil {
		return nil, err
	}

	switch r := res.(type) {
	case *sacloud.GetQueuesOK:
		return r.CommonServiceItems, nil
	case *sacloud.GetQueuesUnauthorized:
		return nil, errors.New(r.ErrorMsg.Value)
	case *sacloud.GetQueuesBadRequest:
		return nil, errors.New(r.ErrorMsg.Value)
	case *sacloud.GetQueuesInternalServerError:
		return nil, errors.New(r.ErrorMsg.Value)
	default:
		return nil, errors.New("unknown error")
	}
}

func (op *sacloudOp) GetQueue(ctx context.Context, id string) (sacloud.CommonServiceItem, error) {
	var empty sacloud.CommonServiceItem
	res, err := op.client.GetQueue(ctx, sacloud.GetQueueParams{ID: id})
	if err != nil {
		return empty, err
	}

	switch r := res.(type) {
	case *sacloud.GetQueueOK:
		return r.CommonServiceItem, nil
	case *sacloud.GetQueueUnauthorized:
		return empty, errors.New(r.ErrorMsg.Value)
	case *sacloud.GetQueueBadRequest:
		return empty, errors.New(r.ErrorMsg.Value)
	case *sacloud.GetQueueNotFound:
		return empty, errors.New(r.ErrorMsg.Value)
	case *sacloud.GetQueueInternalServerError:
		return empty, errors.New(r.ErrorMsg.Value)
	default:
		return empty, errors.New("unknown error")
	}
}

func (op *sacloudOp) ConfigQueue(ctx context.Context, id string, req sacloud.ConfigQueueRequest) (sacloud.CommonServiceItem, error) {
	var empty sacloud.CommonServiceItem
	res, err := op.client.ConfigQueue(ctx, &req, sacloud.ConfigQueueParams{ID: id})
	if err != nil {
		return empty, err
	}

	switch r := res.(type) {
	case *sacloud.ConfigQueueOK:
		return r.CommonServiceItem, nil
	case *sacloud.ConfigQueueUnauthorized:
		return empty, errors.New(r.ErrorMsg.Value)
	case *sacloud.ConfigQueueBadRequest:
		return empty, errors.New(r.ErrorMsg.Value)
	case *sacloud.ConfigQueueNotFound:
		return empty, errors.New(r.ErrorMsg.Value)
	case *sacloud.ConfigQueueInternalServerError:
		return empty, errors.New(r.ErrorMsg.Value)
	default:
		return empty, errors.New("unknown error")
	}
}

func (op *sacloudOp) DeleteQueue(ctx context.Context, id string) (sacloud.CommonServiceItem, error) {
	var empty sacloud.CommonServiceItem
	res, err := op.client.DeleteQueue(ctx, sacloud.DeleteQueueParams{
		ID: id,
	})
	if err != nil {
		return empty, err
	}

	switch r := res.(type) {
	case *sacloud.DeleteQueueOK:
		return r.CommonServiceItem, nil
	case *sacloud.DeleteQueueUnauthorized:
		return empty, errors.New(r.ErrorMsg.Value)
	case *sacloud.DeleteQueueBadRequest:
		return empty, errors.New(r.ErrorMsg.Value)
	case *sacloud.DeleteQueueNotFound:
		return empty, errors.New(r.ErrorMsg.Value)
	case *sacloud.DeleteQueueInternalServerError:
		return empty, errors.New(r.ErrorMsg.Value)
	default:
		return empty, errors.New("unknown error")
	}
}

func (op *sacloudOp) GetMessageCount(ctx context.Context, id string) (int, error) {
	res, err := op.client.GetMessageCount(ctx, sacloud.GetMessageCountParams{ID: id})
	if err != nil {
		return 0, err
	}

	switch r := res.(type) {
	case *sacloud.GetMessageCountOK:
		return r.SimpleMQ.GetCount(), nil
	case *sacloud.GetMessageCountUnauthorized:
		return 0, errors.New(r.ErrorMsg.Value)
	case *sacloud.GetMessageCountBadRequest:
		return 0, errors.New(r.ErrorMsg.Value)
	case *sacloud.GetMessageCountNotFound:
		return 0, errors.New(r.ErrorMsg.Value)
	case *sacloud.GetMessageCountInternalServerError:
		return 0, errors.New(r.ErrorMsg.Value)
	default:
		return 0, errors.New("unknown error")
	}
}

func (op *sacloudOp) RotateAPIKey(ctx context.Context, id string) (string, error) {
	res, err := op.client.RotateAPIKey(ctx, sacloud.RotateAPIKeyParams{ID: id})
	if err != nil {
		return "", err
	}

	switch r := res.(type) {
	case *sacloud.RotateAPIKeyOK:
		return r.SimpleMQ.GetApikey(), nil
	case *sacloud.RotateAPIKeyUnauthorized:
		return "", errors.New(r.ErrorMsg.Value)
	case *sacloud.RotateAPIKeyBadRequest:
		return "", errors.New(r.ErrorMsg.Value)
	case *sacloud.RotateAPIKeyNotFound:
		return "", errors.New(r.ErrorMsg.Value)
	case *sacloud.RotateAPIKeyInternalServerError:
		return "", errors.New(r.ErrorMsg.Value)
	default:
		return "", errors.New("unknown error")
	}
}

func (op *sacloudOp) ClearQueue(ctx context.Context, id string) error {
	res, err := op.client.ClearQueue(ctx, sacloud.ClearQueueParams{ID: id})
	if err != nil {
		return err
	}

	switch r := res.(type) {
	case *sacloud.ClearQueueOK:
		return nil
	case *sacloud.ClearQueueUnauthorized:
		return errors.New(r.ErrorMsg.Value)
	case *sacloud.ClearQueueBadRequest:
		return errors.New(r.ErrorMsg.Value)
	case *sacloud.ClearQueueNotFound:
		return errors.New(r.ErrorMsg.Value)
	case *sacloud.ClearQueueInternalServerError:
		return errors.New(r.ErrorMsg.Value)
	default:
		return errors.New("unknown error")
	}
}
