// Code generated by ogen, DO NOT EDIT.

package message

import (
	"context"
	"net/url"
	"strings"

	"github.com/go-faster/errors"

	"github.com/ogen-go/ogen/conv"
	ht "github.com/ogen-go/ogen/http"
	"github.com/ogen-go/ogen/ogenerrors"
	"github.com/ogen-go/ogen/uri"
)

func trimTrailingSlashes(u *url.URL) {
	u.Path = strings.TrimRight(u.Path, "/")
	u.RawPath = strings.TrimRight(u.RawPath, "/")
}

// Invoker invokes operations described by OpenAPI v3 specification.
type Invoker interface {
	// DeleteMessage invokes deleteMessage operation.
	//
	// 読み取り済みのメッセージを削除 (ack).
	//
	// DELETE /v1/queues/{queueName}/messages/{messageId}
	DeleteMessage(ctx context.Context, params DeleteMessageParams) (DeleteMessageRes, error)
	// ExtendMessageTimeout invokes extendMessageTimeout operation.
	//
	// メッセージのタイムアウトをキューの設定値で延長.
	//
	// PUT /v1/queues/{queueName}/messages/{messageId}
	ExtendMessageTimeout(ctx context.Context, params ExtendMessageTimeoutParams) (ExtendMessageTimeoutRes, error)
	// ReceiveMessage invokes receiveMessage operation.
	//
	// キューに対するメッセージのreceive (dequeue).
	//
	// GET /v1/queues/{queueName}/messages
	ReceiveMessage(ctx context.Context, params ReceiveMessageParams) (ReceiveMessageRes, error)
	// SendMessage invokes sendMessage operation.
	//
	// キューに対するメッセージのsend (enqueue).
	//
	// POST /v1/queues/{queueName}/messages
	SendMessage(ctx context.Context, request *SendRequest, params SendMessageParams) (SendMessageRes, error)
}

// Client implements OAS client.
type Client struct {
	serverURL *url.URL
	sec       SecuritySource
	baseClient
}

// NewClient initializes new Client defined by OAS.
func NewClient(serverURL string, sec SecuritySource, opts ...ClientOption) (*Client, error) {
	u, err := url.Parse(serverURL)
	if err != nil {
		return nil, err
	}
	trimTrailingSlashes(u)

	c, err := newClientConfig(opts...).baseClient()
	if err != nil {
		return nil, err
	}
	return &Client{
		serverURL:  u,
		sec:        sec,
		baseClient: c,
	}, nil
}

type serverURLKey struct{}

// WithServerURL sets context key to override server URL.
func WithServerURL(ctx context.Context, u *url.URL) context.Context {
	return context.WithValue(ctx, serverURLKey{}, u)
}

func (c *Client) requestURL(ctx context.Context) *url.URL {
	u, ok := ctx.Value(serverURLKey{}).(*url.URL)
	if !ok {
		return c.serverURL
	}
	return u
}

// DeleteMessage invokes deleteMessage operation.
//
// 読み取り済みのメッセージを削除 (ack).
//
// DELETE /v1/queues/{queueName}/messages/{messageId}
func (c *Client) DeleteMessage(ctx context.Context, params DeleteMessageParams) (DeleteMessageRes, error) {
	res, err := c.sendDeleteMessage(ctx, params)
	return res, err
}

func (c *Client) sendDeleteMessage(ctx context.Context, params DeleteMessageParams) (res DeleteMessageRes, err error) {

	u := uri.Clone(c.requestURL(ctx))
	var pathParts [4]string
	pathParts[0] = "/v1/queues/"
	{
		// Encode "queueName" parameter.
		e := uri.NewPathEncoder(uri.PathEncoderConfig{
			Param:   "queueName",
			Style:   uri.PathStyleSimple,
			Explode: false,
		})
		if err := func() error {
			if unwrapped := string(params.QueueName); true {
				return e.EncodeValue(conv.StringToString(unwrapped))
			}
			return nil
		}(); err != nil {
			return res, errors.Wrap(err, "encode path")
		}
		encoded, err := e.Result()
		if err != nil {
			return res, errors.Wrap(err, "encode path")
		}
		pathParts[1] = encoded
	}
	pathParts[2] = "/messages/"
	{
		// Encode "messageId" parameter.
		e := uri.NewPathEncoder(uri.PathEncoderConfig{
			Param:   "messageId",
			Style:   uri.PathStyleSimple,
			Explode: false,
		})
		if err := func() error {
			if unwrapped := string(params.MessageId); true {
				return e.EncodeValue(conv.StringToString(unwrapped))
			}
			return nil
		}(); err != nil {
			return res, errors.Wrap(err, "encode path")
		}
		encoded, err := e.Result()
		if err != nil {
			return res, errors.Wrap(err, "encode path")
		}
		pathParts[3] = encoded
	}
	uri.AddPathParts(u, pathParts[:]...)

	r, err := ht.NewRequest(ctx, "DELETE", u)
	if err != nil {
		return res, errors.Wrap(err, "create request")
	}

	{
		type bitset = [1]uint8
		var satisfied bitset
		{

			switch err := c.securityApiKeyAuth(ctx, DeleteMessageOperation, r); {
			case err == nil: // if NO error
				satisfied[0] |= 1 << 0
			case errors.Is(err, ogenerrors.ErrSkipClientSecurity):
				// Skip this security.
			default:
				return res, errors.Wrap(err, "security \"ApiKeyAuth\"")
			}
		}

		if ok := func() bool {
		nextRequirement:
			for _, requirement := range []bitset{
				{0b00000001},
			} {
				for i, mask := range requirement {
					if satisfied[i]&mask != mask {
						continue nextRequirement
					}
				}
				return true
			}
			return false
		}(); !ok {
			return res, ogenerrors.ErrSecurityRequirementIsNotSatisfied
		}
	}

	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	defer resp.Body.Close()

	result, err := decodeDeleteMessageResponse(resp)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}

// ExtendMessageTimeout invokes extendMessageTimeout operation.
//
// メッセージのタイムアウトをキューの設定値で延長.
//
// PUT /v1/queues/{queueName}/messages/{messageId}
func (c *Client) ExtendMessageTimeout(ctx context.Context, params ExtendMessageTimeoutParams) (ExtendMessageTimeoutRes, error) {
	res, err := c.sendExtendMessageTimeout(ctx, params)
	return res, err
}

func (c *Client) sendExtendMessageTimeout(ctx context.Context, params ExtendMessageTimeoutParams) (res ExtendMessageTimeoutRes, err error) {

	u := uri.Clone(c.requestURL(ctx))
	var pathParts [4]string
	pathParts[0] = "/v1/queues/"
	{
		// Encode "queueName" parameter.
		e := uri.NewPathEncoder(uri.PathEncoderConfig{
			Param:   "queueName",
			Style:   uri.PathStyleSimple,
			Explode: false,
		})
		if err := func() error {
			if unwrapped := string(params.QueueName); true {
				return e.EncodeValue(conv.StringToString(unwrapped))
			}
			return nil
		}(); err != nil {
			return res, errors.Wrap(err, "encode path")
		}
		encoded, err := e.Result()
		if err != nil {
			return res, errors.Wrap(err, "encode path")
		}
		pathParts[1] = encoded
	}
	pathParts[2] = "/messages/"
	{
		// Encode "messageId" parameter.
		e := uri.NewPathEncoder(uri.PathEncoderConfig{
			Param:   "messageId",
			Style:   uri.PathStyleSimple,
			Explode: false,
		})
		if err := func() error {
			if unwrapped := string(params.MessageId); true {
				return e.EncodeValue(conv.StringToString(unwrapped))
			}
			return nil
		}(); err != nil {
			return res, errors.Wrap(err, "encode path")
		}
		encoded, err := e.Result()
		if err != nil {
			return res, errors.Wrap(err, "encode path")
		}
		pathParts[3] = encoded
	}
	uri.AddPathParts(u, pathParts[:]...)

	r, err := ht.NewRequest(ctx, "PUT", u)
	if err != nil {
		return res, errors.Wrap(err, "create request")
	}

	{
		type bitset = [1]uint8
		var satisfied bitset
		{

			switch err := c.securityApiKeyAuth(ctx, ExtendMessageTimeoutOperation, r); {
			case err == nil: // if NO error
				satisfied[0] |= 1 << 0
			case errors.Is(err, ogenerrors.ErrSkipClientSecurity):
				// Skip this security.
			default:
				return res, errors.Wrap(err, "security \"ApiKeyAuth\"")
			}
		}

		if ok := func() bool {
		nextRequirement:
			for _, requirement := range []bitset{
				{0b00000001},
			} {
				for i, mask := range requirement {
					if satisfied[i]&mask != mask {
						continue nextRequirement
					}
				}
				return true
			}
			return false
		}(); !ok {
			return res, ogenerrors.ErrSecurityRequirementIsNotSatisfied
		}
	}

	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	defer resp.Body.Close()

	result, err := decodeExtendMessageTimeoutResponse(resp)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}

// ReceiveMessage invokes receiveMessage operation.
//
// キューに対するメッセージのreceive (dequeue).
//
// GET /v1/queues/{queueName}/messages
func (c *Client) ReceiveMessage(ctx context.Context, params ReceiveMessageParams) (ReceiveMessageRes, error) {
	res, err := c.sendReceiveMessage(ctx, params)
	return res, err
}

func (c *Client) sendReceiveMessage(ctx context.Context, params ReceiveMessageParams) (res ReceiveMessageRes, err error) {

	u := uri.Clone(c.requestURL(ctx))
	var pathParts [3]string
	pathParts[0] = "/v1/queues/"
	{
		// Encode "queueName" parameter.
		e := uri.NewPathEncoder(uri.PathEncoderConfig{
			Param:   "queueName",
			Style:   uri.PathStyleSimple,
			Explode: false,
		})
		if err := func() error {
			if unwrapped := string(params.QueueName); true {
				return e.EncodeValue(conv.StringToString(unwrapped))
			}
			return nil
		}(); err != nil {
			return res, errors.Wrap(err, "encode path")
		}
		encoded, err := e.Result()
		if err != nil {
			return res, errors.Wrap(err, "encode path")
		}
		pathParts[1] = encoded
	}
	pathParts[2] = "/messages"
	uri.AddPathParts(u, pathParts[:]...)

	r, err := ht.NewRequest(ctx, "GET", u)
	if err != nil {
		return res, errors.Wrap(err, "create request")
	}

	{
		type bitset = [1]uint8
		var satisfied bitset
		{

			switch err := c.securityApiKeyAuth(ctx, ReceiveMessageOperation, r); {
			case err == nil: // if NO error
				satisfied[0] |= 1 << 0
			case errors.Is(err, ogenerrors.ErrSkipClientSecurity):
				// Skip this security.
			default:
				return res, errors.Wrap(err, "security \"ApiKeyAuth\"")
			}
		}

		if ok := func() bool {
		nextRequirement:
			for _, requirement := range []bitset{
				{0b00000001},
			} {
				for i, mask := range requirement {
					if satisfied[i]&mask != mask {
						continue nextRequirement
					}
				}
				return true
			}
			return false
		}(); !ok {
			return res, ogenerrors.ErrSecurityRequirementIsNotSatisfied
		}
	}

	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	defer resp.Body.Close()

	result, err := decodeReceiveMessageResponse(resp)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}

// SendMessage invokes sendMessage operation.
//
// キューに対するメッセージのsend (enqueue).
//
// POST /v1/queues/{queueName}/messages
func (c *Client) SendMessage(ctx context.Context, request *SendRequest, params SendMessageParams) (SendMessageRes, error) {
	res, err := c.sendSendMessage(ctx, request, params)
	return res, err
}

func (c *Client) sendSendMessage(ctx context.Context, request *SendRequest, params SendMessageParams) (res SendMessageRes, err error) {
	// Validate request before sending.
	if err := func() error {
		if err := request.Validate(); err != nil {
			return err
		}
		return nil
	}(); err != nil {
		return res, errors.Wrap(err, "validate")
	}

	u := uri.Clone(c.requestURL(ctx))
	var pathParts [3]string
	pathParts[0] = "/v1/queues/"
	{
		// Encode "queueName" parameter.
		e := uri.NewPathEncoder(uri.PathEncoderConfig{
			Param:   "queueName",
			Style:   uri.PathStyleSimple,
			Explode: false,
		})
		if err := func() error {
			if unwrapped := string(params.QueueName); true {
				return e.EncodeValue(conv.StringToString(unwrapped))
			}
			return nil
		}(); err != nil {
			return res, errors.Wrap(err, "encode path")
		}
		encoded, err := e.Result()
		if err != nil {
			return res, errors.Wrap(err, "encode path")
		}
		pathParts[1] = encoded
	}
	pathParts[2] = "/messages"
	uri.AddPathParts(u, pathParts[:]...)

	r, err := ht.NewRequest(ctx, "POST", u)
	if err != nil {
		return res, errors.Wrap(err, "create request")
	}
	if err := encodeSendMessageRequest(request, r); err != nil {
		return res, errors.Wrap(err, "encode request")
	}

	{
		type bitset = [1]uint8
		var satisfied bitset
		{

			switch err := c.securityApiKeyAuth(ctx, SendMessageOperation, r); {
			case err == nil: // if NO error
				satisfied[0] |= 1 << 0
			case errors.Is(err, ogenerrors.ErrSkipClientSecurity):
				// Skip this security.
			default:
				return res, errors.Wrap(err, "security \"ApiKeyAuth\"")
			}
		}

		if ok := func() bool {
		nextRequirement:
			for _, requirement := range []bitset{
				{0b00000001},
			} {
				for i, mask := range requirement {
					if satisfied[i]&mask != mask {
						continue nextRequirement
					}
				}
				return true
			}
			return false
		}(); !ok {
			return res, ogenerrors.ErrSecurityRequirementIsNotSatisfied
		}
	}

	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	defer resp.Body.Close()

	result, err := decodeSendMessageResponse(resp)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}
