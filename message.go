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

	"github.com/sacloud/simplemq-api-go/apis/v1/message"
)

type MessageAPI interface {
	Send(_ context.Context, content string) (*message.NewMessage, error)
	Receive(_ context.Context) ([]message.Message, error)
	ExtendTimeout(_ context.Context, messageID string) (*message.Message, error)
	Delete(_ context.Context, messageID string) error
}

var _ MessageAPI = (*messageOp)(nil)

type messageOp struct {
	queueName message.QueueName
	client    *message.Client
}

func NewMessageOp(client *message.Client, queueName string) MessageAPI {
	return &messageOp{
		queueName: message.QueueName(queueName),
		client:    client,
	}
}

type MessageError message.Error

func (e MessageError) Error() string {
	return e.Message.Value
}

func (op *messageOp) Send(ctx context.Context, content string) (*message.NewMessage, error) {
	res, err := op.client.SendMessage(ctx,
		&message.SendRequest{
			Content: message.MessageContent(content),
		},
		message.SendMessageParams{
			QueueName: op.queueName,
		})
	if err != nil {
		return nil, err
	}

	switch r := res.(type) {
	case *message.SendMessageOK:
		return &r.Message, nil
	case *message.SendMessageUnauthorized:
		return nil, MessageError(*r)
	case *message.SendMessageBadRequest:
		return nil, MessageError(*r)
	case *message.SendMessageInternalServerError:
		return nil, MessageError(*r)
	default:
		return nil, MessageError{
			Message: message.NewOptString("unknown error"),
		}
	}
}

func (op *messageOp) Receive(ctx context.Context) ([]message.Message, error) {
	res, err := op.client.ReceiveMessage(ctx,
		message.ReceiveMessageParams{
			QueueName: op.queueName,
		})
	if err != nil {
		return nil, err
	}

	switch r := res.(type) {
	case *message.ReceiveMessageOK:
		return r.Messages, nil
	case *message.ReceiveMessageUnauthorized:
		return nil, MessageError(*r)
	case *message.ReceiveMessageBadRequest:
		return nil, MessageError(*r)
	case *message.ReceiveMessageInternalServerError:
		return nil, MessageError(*r)
	default:
		return nil, MessageError{
			Message: message.NewOptString("unknown error"),
		}
	}
}

func (op *messageOp) ExtendTimeout(ctx context.Context, messageID string) (*message.Message, error) {
	res, err := op.client.ExtendMessageTimeout(ctx,
		message.ExtendMessageTimeoutParams{
			QueueName: op.queueName,
			MessageId: message.MessageId(messageID),
		})
	if err != nil {
		return nil, err
	}

	switch r := res.(type) {
	case *message.ExtendMessageTimeoutOK:
		return &r.Message, nil
	case *message.ExtendMessageTimeoutUnauthorized:
		return nil, MessageError(*r)
	case *message.ExtendMessageTimeoutBadRequest:
		return nil, MessageError(*r)
	case *message.ExtendMessageTimeoutNotFound:
		return nil, MessageError(*r)
	case *message.ExtendMessageTimeoutInternalServerError:
		return nil, MessageError(*r)
	default:
		return nil, MessageError{
			Message: message.NewOptString("unknown error"),
		}
	}
}

func (op *messageOp) Delete(ctx context.Context, messageID string) error {
	res, err := op.client.DeleteMessage(ctx,
		message.DeleteMessageParams{
			QueueName: op.queueName,
			MessageId: message.MessageId(messageID),
		})
	if err != nil {
		return err
	}

	switch r := res.(type) {
	case *message.DeleteMessageOK:
		return nil
	case *message.DeleteMessageUnauthorized:
		return MessageError(*r)
	case *message.DeleteMessageBadRequest:
		return MessageError(*r)
	case *message.DeleteMessageNotFound:
		return MessageError(*r)
	case *message.DeleteMessageInternalServerError:
		return MessageError(*r)
	default:
		return MessageError{
			Message: message.NewOptString("unknown error"),
		}
	}
}
