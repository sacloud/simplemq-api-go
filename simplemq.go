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

	"github.com/sacloud/simplemq-api-go/apis/v1/simplemq"
)

type SimpleMQAPI interface {
	SendMessage(_ context.Context, message string) (simplemq.NewMessage, error)
	ReceiveMessage(_ context.Context) ([]simplemq.Message, error)
	ExtendMessageTimeout(_ context.Context, messageID string) (simplemq.Message, error)
	DeleteMessage(_ context.Context, messageID string) error
}

var _ SimpleMQAPI = (*simplemqOp)(nil)

type simplemqOp struct {
	queueName simplemq.QueueName
	client    *simplemq.Client
}

func newSimpleMQOp(client *simplemq.Client, queueName string) SimpleMQAPI {
	return &simplemqOp{
		queueName: simplemq.QueueName(queueName),
		client:    client,
	}
}

type MQError simplemq.Error

func (e MQError) Error() string {
	return e.Message.Value
}

func (op *simplemqOp) SendMessage(ctx context.Context, message string) (simplemq.NewMessage, error) {
	var empty simplemq.NewMessage
	res, err := op.client.SendMessage(ctx,
		&simplemq.SendRequest{
			Content: simplemq.MessageContent(message),
		},
		simplemq.SendMessageParams{
			QueueName: op.queueName,
		})
	if err != nil {
		return empty, err
	}

	switch r := res.(type) {
	case *simplemq.SendMessageOK:
		return r.Message, nil
	case *simplemq.SendMessageUnauthorized:
		return empty, MQError(*r)
	case *simplemq.SendMessageBadRequest:
		return empty, MQError(*r)
	case *simplemq.SendMessageInternalServerError:
		return empty, MQError(*r)
	default:
		return empty, MQError{
			Message: simplemq.NewOptString("unknown error"),
		}
	}
}

func (op *simplemqOp) ReceiveMessage(ctx context.Context) ([]simplemq.Message, error) {
	res, err := op.client.ReceiveMessage(ctx,
		simplemq.ReceiveMessageParams{
			QueueName: op.queueName,
		})
	if err != nil {
		return nil, err
	}

	switch r := res.(type) {
	case *simplemq.ReceiveMessageOK:
		return r.Messages, nil
	case *simplemq.ReceiveMessageUnauthorized:
		return nil, MQError(*r)
	case *simplemq.ReceiveMessageBadRequest:
		return nil, MQError(*r)
	case *simplemq.ReceiveMessageInternalServerError:
		return nil, MQError(*r)
	default:
		return nil, MQError{
			Message: simplemq.NewOptString("unknown error"),
		}
	}
}

func (op *simplemqOp) ExtendMessageTimeout(ctx context.Context, messageID string) (simplemq.Message, error) {
	var empty simplemq.Message
	res, err := op.client.ExtendMessageTimeout(ctx,
		simplemq.ExtendMessageTimeoutParams{
			QueueName: op.queueName,
			MessageId: simplemq.MessageId(messageID),
		})
	if err != nil {
		return empty, err
	}

	switch r := res.(type) {
	case *simplemq.ExtendMessageTimeoutOK:
		return r.Message, nil
	case *simplemq.ExtendMessageTimeoutUnauthorized:
		return empty, MQError(*r)
	case *simplemq.ExtendMessageTimeoutBadRequest:
		return empty, MQError(*r)
	case *simplemq.ExtendMessageTimeoutNotFound:
		return empty, MQError(*r)
	case *simplemq.ExtendMessageTimeoutInternalServerError:
		return empty, MQError(*r)
	default:
		return empty, MQError{
			Message: simplemq.NewOptString("unknown error"),
		}
	}
}

func (op *simplemqOp) DeleteMessage(ctx context.Context, messageID string) error {
	res, err := op.client.DeleteMessage(ctx,
		simplemq.DeleteMessageParams{
			QueueName: op.queueName,
			MessageId: simplemq.MessageId(messageID),
		})
	if err != nil {
		return err
	}

	switch r := res.(type) {
	case *simplemq.DeleteMessageOK:
		return nil
	case *simplemq.DeleteMessageUnauthorized:
		return MQError(*r)
	case *simplemq.DeleteMessageBadRequest:
		return MQError(*r)
	case *simplemq.DeleteMessageNotFound:
		return MQError(*r)
	case *simplemq.DeleteMessageInternalServerError:
		return MQError(*r)
	default:
		return MQError{
			Message: simplemq.NewOptString("unknown error"),
		}
	}
}
