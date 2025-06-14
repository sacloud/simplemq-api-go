// Code generated by ogen, DO NOT EDIT.

package message

import (
	"math/bits"
	"strconv"

	"github.com/go-faster/errors"
	"github.com/go-faster/jx"

	"github.com/ogen-go/ogen/validate"
)

// Encode encodes DeleteMessageBadRequest as json.
func (s *DeleteMessageBadRequest) Encode(e *jx.Encoder) {
	unwrapped := (*Error)(s)

	unwrapped.Encode(e)
}

// Decode decodes DeleteMessageBadRequest from json.
func (s *DeleteMessageBadRequest) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode DeleteMessageBadRequest to nil")
	}
	var unwrapped Error
	if err := func() error {
		if err := unwrapped.Decode(d); err != nil {
			return err
		}
		return nil
	}(); err != nil {
		return errors.Wrap(err, "alias")
	}
	*s = DeleteMessageBadRequest(unwrapped)
	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s *DeleteMessageBadRequest) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *DeleteMessageBadRequest) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode encodes DeleteMessageInternalServerError as json.
func (s *DeleteMessageInternalServerError) Encode(e *jx.Encoder) {
	unwrapped := (*Error)(s)

	unwrapped.Encode(e)
}

// Decode decodes DeleteMessageInternalServerError from json.
func (s *DeleteMessageInternalServerError) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode DeleteMessageInternalServerError to nil")
	}
	var unwrapped Error
	if err := func() error {
		if err := unwrapped.Decode(d); err != nil {
			return err
		}
		return nil
	}(); err != nil {
		return errors.Wrap(err, "alias")
	}
	*s = DeleteMessageInternalServerError(unwrapped)
	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s *DeleteMessageInternalServerError) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *DeleteMessageInternalServerError) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode encodes DeleteMessageNotFound as json.
func (s *DeleteMessageNotFound) Encode(e *jx.Encoder) {
	unwrapped := (*Error)(s)

	unwrapped.Encode(e)
}

// Decode decodes DeleteMessageNotFound from json.
func (s *DeleteMessageNotFound) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode DeleteMessageNotFound to nil")
	}
	var unwrapped Error
	if err := func() error {
		if err := unwrapped.Decode(d); err != nil {
			return err
		}
		return nil
	}(); err != nil {
		return errors.Wrap(err, "alias")
	}
	*s = DeleteMessageNotFound(unwrapped)
	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s *DeleteMessageNotFound) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *DeleteMessageNotFound) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode implements json.Marshaler.
func (s *DeleteMessageOK) Encode(e *jx.Encoder) {
	e.ObjStart()
	s.encodeFields(e)
	e.ObjEnd()
}

// encodeFields encodes fields.
func (s *DeleteMessageOK) encodeFields(e *jx.Encoder) {
	{
		e.FieldStart("result")
		e.Str(s.Result)
	}
}

var jsonFieldsNameOfDeleteMessageOK = [1]string{
	0: "result",
}

// Decode decodes DeleteMessageOK from json.
func (s *DeleteMessageOK) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode DeleteMessageOK to nil")
	}
	var requiredBitSet [1]uint8

	if err := d.ObjBytes(func(d *jx.Decoder, k []byte) error {
		switch string(k) {
		case "result":
			requiredBitSet[0] |= 1 << 0
			if err := func() error {
				v, err := d.Str()
				s.Result = string(v)
				if err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"result\"")
			}
		default:
			return d.Skip()
		}
		return nil
	}); err != nil {
		return errors.Wrap(err, "decode DeleteMessageOK")
	}
	// Validate required fields.
	var failures []validate.FieldError
	for i, mask := range [1]uint8{
		0b00000001,
	} {
		if result := (requiredBitSet[i] & mask) ^ mask; result != 0 {
			// Mask only required fields and check equality to mask using XOR.
			//
			// If XOR result is not zero, result is not equal to expected, so some fields are missed.
			// Bits of fields which would be set are actually bits of missed fields.
			missed := bits.OnesCount8(result)
			for bitN := 0; bitN < missed; bitN++ {
				bitIdx := bits.TrailingZeros8(result)
				fieldIdx := i*8 + bitIdx
				var name string
				if fieldIdx < len(jsonFieldsNameOfDeleteMessageOK) {
					name = jsonFieldsNameOfDeleteMessageOK[fieldIdx]
				} else {
					name = strconv.Itoa(fieldIdx)
				}
				failures = append(failures, validate.FieldError{
					Name:  name,
					Error: validate.ErrFieldRequired,
				})
				// Reset bit.
				result &^= 1 << bitIdx
			}
		}
	}
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}

	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s *DeleteMessageOK) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *DeleteMessageOK) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode encodes DeleteMessageUnauthorized as json.
func (s *DeleteMessageUnauthorized) Encode(e *jx.Encoder) {
	unwrapped := (*Error)(s)

	unwrapped.Encode(e)
}

// Decode decodes DeleteMessageUnauthorized from json.
func (s *DeleteMessageUnauthorized) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode DeleteMessageUnauthorized to nil")
	}
	var unwrapped Error
	if err := func() error {
		if err := unwrapped.Decode(d); err != nil {
			return err
		}
		return nil
	}(); err != nil {
		return errors.Wrap(err, "alias")
	}
	*s = DeleteMessageUnauthorized(unwrapped)
	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s *DeleteMessageUnauthorized) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *DeleteMessageUnauthorized) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode implements json.Marshaler.
func (s *Error) Encode(e *jx.Encoder) {
	e.ObjStart()
	s.encodeFields(e)
	e.ObjEnd()
}

// encodeFields encodes fields.
func (s *Error) encodeFields(e *jx.Encoder) {
	{
		if s.Code.Set {
			e.FieldStart("code")
			s.Code.Encode(e)
		}
	}
	{
		if s.Message.Set {
			e.FieldStart("message")
			s.Message.Encode(e)
		}
	}
}

var jsonFieldsNameOfError = [2]string{
	0: "code",
	1: "message",
}

// Decode decodes Error from json.
func (s *Error) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode Error to nil")
	}

	if err := d.ObjBytes(func(d *jx.Decoder, k []byte) error {
		switch string(k) {
		case "code":
			if err := func() error {
				s.Code.Reset()
				if err := s.Code.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"code\"")
			}
		case "message":
			if err := func() error {
				s.Message.Reset()
				if err := s.Message.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"message\"")
			}
		default:
			return d.Skip()
		}
		return nil
	}); err != nil {
		return errors.Wrap(err, "decode Error")
	}

	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s *Error) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *Error) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode encodes ExtendMessageTimeoutBadRequest as json.
func (s *ExtendMessageTimeoutBadRequest) Encode(e *jx.Encoder) {
	unwrapped := (*Error)(s)

	unwrapped.Encode(e)
}

// Decode decodes ExtendMessageTimeoutBadRequest from json.
func (s *ExtendMessageTimeoutBadRequest) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode ExtendMessageTimeoutBadRequest to nil")
	}
	var unwrapped Error
	if err := func() error {
		if err := unwrapped.Decode(d); err != nil {
			return err
		}
		return nil
	}(); err != nil {
		return errors.Wrap(err, "alias")
	}
	*s = ExtendMessageTimeoutBadRequest(unwrapped)
	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s *ExtendMessageTimeoutBadRequest) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *ExtendMessageTimeoutBadRequest) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode encodes ExtendMessageTimeoutInternalServerError as json.
func (s *ExtendMessageTimeoutInternalServerError) Encode(e *jx.Encoder) {
	unwrapped := (*Error)(s)

	unwrapped.Encode(e)
}

// Decode decodes ExtendMessageTimeoutInternalServerError from json.
func (s *ExtendMessageTimeoutInternalServerError) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode ExtendMessageTimeoutInternalServerError to nil")
	}
	var unwrapped Error
	if err := func() error {
		if err := unwrapped.Decode(d); err != nil {
			return err
		}
		return nil
	}(); err != nil {
		return errors.Wrap(err, "alias")
	}
	*s = ExtendMessageTimeoutInternalServerError(unwrapped)
	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s *ExtendMessageTimeoutInternalServerError) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *ExtendMessageTimeoutInternalServerError) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode encodes ExtendMessageTimeoutNotFound as json.
func (s *ExtendMessageTimeoutNotFound) Encode(e *jx.Encoder) {
	unwrapped := (*Error)(s)

	unwrapped.Encode(e)
}

// Decode decodes ExtendMessageTimeoutNotFound from json.
func (s *ExtendMessageTimeoutNotFound) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode ExtendMessageTimeoutNotFound to nil")
	}
	var unwrapped Error
	if err := func() error {
		if err := unwrapped.Decode(d); err != nil {
			return err
		}
		return nil
	}(); err != nil {
		return errors.Wrap(err, "alias")
	}
	*s = ExtendMessageTimeoutNotFound(unwrapped)
	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s *ExtendMessageTimeoutNotFound) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *ExtendMessageTimeoutNotFound) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode implements json.Marshaler.
func (s *ExtendMessageTimeoutOK) Encode(e *jx.Encoder) {
	e.ObjStart()
	s.encodeFields(e)
	e.ObjEnd()
}

// encodeFields encodes fields.
func (s *ExtendMessageTimeoutOK) encodeFields(e *jx.Encoder) {
	{
		e.FieldStart("result")
		e.Str(s.Result)
	}
	{
		e.FieldStart("message")
		s.Message.Encode(e)
	}
}

var jsonFieldsNameOfExtendMessageTimeoutOK = [2]string{
	0: "result",
	1: "message",
}

// Decode decodes ExtendMessageTimeoutOK from json.
func (s *ExtendMessageTimeoutOK) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode ExtendMessageTimeoutOK to nil")
	}
	var requiredBitSet [1]uint8

	if err := d.ObjBytes(func(d *jx.Decoder, k []byte) error {
		switch string(k) {
		case "result":
			requiredBitSet[0] |= 1 << 0
			if err := func() error {
				v, err := d.Str()
				s.Result = string(v)
				if err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"result\"")
			}
		case "message":
			requiredBitSet[0] |= 1 << 1
			if err := func() error {
				if err := s.Message.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"message\"")
			}
		default:
			return d.Skip()
		}
		return nil
	}); err != nil {
		return errors.Wrap(err, "decode ExtendMessageTimeoutOK")
	}
	// Validate required fields.
	var failures []validate.FieldError
	for i, mask := range [1]uint8{
		0b00000011,
	} {
		if result := (requiredBitSet[i] & mask) ^ mask; result != 0 {
			// Mask only required fields and check equality to mask using XOR.
			//
			// If XOR result is not zero, result is not equal to expected, so some fields are missed.
			// Bits of fields which would be set are actually bits of missed fields.
			missed := bits.OnesCount8(result)
			for bitN := 0; bitN < missed; bitN++ {
				bitIdx := bits.TrailingZeros8(result)
				fieldIdx := i*8 + bitIdx
				var name string
				if fieldIdx < len(jsonFieldsNameOfExtendMessageTimeoutOK) {
					name = jsonFieldsNameOfExtendMessageTimeoutOK[fieldIdx]
				} else {
					name = strconv.Itoa(fieldIdx)
				}
				failures = append(failures, validate.FieldError{
					Name:  name,
					Error: validate.ErrFieldRequired,
				})
				// Reset bit.
				result &^= 1 << bitIdx
			}
		}
	}
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}

	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s *ExtendMessageTimeoutOK) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *ExtendMessageTimeoutOK) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode encodes ExtendMessageTimeoutUnauthorized as json.
func (s *ExtendMessageTimeoutUnauthorized) Encode(e *jx.Encoder) {
	unwrapped := (*Error)(s)

	unwrapped.Encode(e)
}

// Decode decodes ExtendMessageTimeoutUnauthorized from json.
func (s *ExtendMessageTimeoutUnauthorized) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode ExtendMessageTimeoutUnauthorized to nil")
	}
	var unwrapped Error
	if err := func() error {
		if err := unwrapped.Decode(d); err != nil {
			return err
		}
		return nil
	}(); err != nil {
		return errors.Wrap(err, "alias")
	}
	*s = ExtendMessageTimeoutUnauthorized(unwrapped)
	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s *ExtendMessageTimeoutUnauthorized) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *ExtendMessageTimeoutUnauthorized) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode implements json.Marshaler.
func (s *Message) Encode(e *jx.Encoder) {
	e.ObjStart()
	s.encodeFields(e)
	e.ObjEnd()
}

// encodeFields encodes fields.
func (s *Message) encodeFields(e *jx.Encoder) {
	{
		e.FieldStart("id")
		s.ID.Encode(e)
	}
	{
		e.FieldStart("content")
		s.Content.Encode(e)
	}
	{
		e.FieldStart("created_at")
		e.Int64(s.CreatedAt)
	}
	{
		e.FieldStart("updated_at")
		e.Int64(s.UpdatedAt)
	}
	{
		e.FieldStart("expires_at")
		e.Int64(s.ExpiresAt)
	}
	{
		e.FieldStart("acquired_at")
		e.Int64(s.AcquiredAt)
	}
	{
		e.FieldStart("visibility_timeout_at")
		e.Int64(s.VisibilityTimeoutAt)
	}
}

var jsonFieldsNameOfMessage = [7]string{
	0: "id",
	1: "content",
	2: "created_at",
	3: "updated_at",
	4: "expires_at",
	5: "acquired_at",
	6: "visibility_timeout_at",
}

// Decode decodes Message from json.
func (s *Message) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode Message to nil")
	}
	var requiredBitSet [1]uint8

	if err := d.ObjBytes(func(d *jx.Decoder, k []byte) error {
		switch string(k) {
		case "id":
			requiredBitSet[0] |= 1 << 0
			if err := func() error {
				if err := s.ID.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"id\"")
			}
		case "content":
			requiredBitSet[0] |= 1 << 1
			if err := func() error {
				if err := s.Content.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"content\"")
			}
		case "created_at":
			requiredBitSet[0] |= 1 << 2
			if err := func() error {
				v, err := d.Int64()
				s.CreatedAt = int64(v)
				if err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"created_at\"")
			}
		case "updated_at":
			requiredBitSet[0] |= 1 << 3
			if err := func() error {
				v, err := d.Int64()
				s.UpdatedAt = int64(v)
				if err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"updated_at\"")
			}
		case "expires_at":
			requiredBitSet[0] |= 1 << 4
			if err := func() error {
				v, err := d.Int64()
				s.ExpiresAt = int64(v)
				if err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"expires_at\"")
			}
		case "acquired_at":
			requiredBitSet[0] |= 1 << 5
			if err := func() error {
				v, err := d.Int64()
				s.AcquiredAt = int64(v)
				if err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"acquired_at\"")
			}
		case "visibility_timeout_at":
			requiredBitSet[0] |= 1 << 6
			if err := func() error {
				v, err := d.Int64()
				s.VisibilityTimeoutAt = int64(v)
				if err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"visibility_timeout_at\"")
			}
		default:
			return d.Skip()
		}
		return nil
	}); err != nil {
		return errors.Wrap(err, "decode Message")
	}
	// Validate required fields.
	var failures []validate.FieldError
	for i, mask := range [1]uint8{
		0b01111111,
	} {
		if result := (requiredBitSet[i] & mask) ^ mask; result != 0 {
			// Mask only required fields and check equality to mask using XOR.
			//
			// If XOR result is not zero, result is not equal to expected, so some fields are missed.
			// Bits of fields which would be set are actually bits of missed fields.
			missed := bits.OnesCount8(result)
			for bitN := 0; bitN < missed; bitN++ {
				bitIdx := bits.TrailingZeros8(result)
				fieldIdx := i*8 + bitIdx
				var name string
				if fieldIdx < len(jsonFieldsNameOfMessage) {
					name = jsonFieldsNameOfMessage[fieldIdx]
				} else {
					name = strconv.Itoa(fieldIdx)
				}
				failures = append(failures, validate.FieldError{
					Name:  name,
					Error: validate.ErrFieldRequired,
				})
				// Reset bit.
				result &^= 1 << bitIdx
			}
		}
	}
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}

	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s *Message) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *Message) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode encodes MessageContent as json.
func (s MessageContent) Encode(e *jx.Encoder) {
	unwrapped := string(s)

	e.Str(unwrapped)
}

// Decode decodes MessageContent from json.
func (s *MessageContent) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode MessageContent to nil")
	}
	var unwrapped string
	if err := func() error {
		v, err := d.Str()
		unwrapped = string(v)
		if err != nil {
			return err
		}
		return nil
	}(); err != nil {
		return errors.Wrap(err, "alias")
	}
	*s = MessageContent(unwrapped)
	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s MessageContent) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *MessageContent) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode encodes MessageId as json.
func (s MessageId) Encode(e *jx.Encoder) {
	unwrapped := string(s)

	e.Str(unwrapped)
}

// Decode decodes MessageId from json.
func (s *MessageId) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode MessageId to nil")
	}
	var unwrapped string
	if err := func() error {
		v, err := d.Str()
		unwrapped = string(v)
		if err != nil {
			return err
		}
		return nil
	}(); err != nil {
		return errors.Wrap(err, "alias")
	}
	*s = MessageId(unwrapped)
	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s MessageId) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *MessageId) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode implements json.Marshaler.
func (s *NewMessage) Encode(e *jx.Encoder) {
	e.ObjStart()
	s.encodeFields(e)
	e.ObjEnd()
}

// encodeFields encodes fields.
func (s *NewMessage) encodeFields(e *jx.Encoder) {
	{
		e.FieldStart("id")
		s.ID.Encode(e)
	}
	{
		e.FieldStart("content")
		s.Content.Encode(e)
	}
	{
		e.FieldStart("created_at")
		e.Int64(s.CreatedAt)
	}
	{
		e.FieldStart("updated_at")
		e.Int64(s.UpdatedAt)
	}
	{
		e.FieldStart("expires_at")
		e.Int64(s.ExpiresAt)
	}
}

var jsonFieldsNameOfNewMessage = [5]string{
	0: "id",
	1: "content",
	2: "created_at",
	3: "updated_at",
	4: "expires_at",
}

// Decode decodes NewMessage from json.
func (s *NewMessage) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode NewMessage to nil")
	}
	var requiredBitSet [1]uint8

	if err := d.ObjBytes(func(d *jx.Decoder, k []byte) error {
		switch string(k) {
		case "id":
			requiredBitSet[0] |= 1 << 0
			if err := func() error {
				if err := s.ID.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"id\"")
			}
		case "content":
			requiredBitSet[0] |= 1 << 1
			if err := func() error {
				if err := s.Content.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"content\"")
			}
		case "created_at":
			requiredBitSet[0] |= 1 << 2
			if err := func() error {
				v, err := d.Int64()
				s.CreatedAt = int64(v)
				if err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"created_at\"")
			}
		case "updated_at":
			requiredBitSet[0] |= 1 << 3
			if err := func() error {
				v, err := d.Int64()
				s.UpdatedAt = int64(v)
				if err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"updated_at\"")
			}
		case "expires_at":
			requiredBitSet[0] |= 1 << 4
			if err := func() error {
				v, err := d.Int64()
				s.ExpiresAt = int64(v)
				if err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"expires_at\"")
			}
		default:
			return d.Skip()
		}
		return nil
	}); err != nil {
		return errors.Wrap(err, "decode NewMessage")
	}
	// Validate required fields.
	var failures []validate.FieldError
	for i, mask := range [1]uint8{
		0b00011111,
	} {
		if result := (requiredBitSet[i] & mask) ^ mask; result != 0 {
			// Mask only required fields and check equality to mask using XOR.
			//
			// If XOR result is not zero, result is not equal to expected, so some fields are missed.
			// Bits of fields which would be set are actually bits of missed fields.
			missed := bits.OnesCount8(result)
			for bitN := 0; bitN < missed; bitN++ {
				bitIdx := bits.TrailingZeros8(result)
				fieldIdx := i*8 + bitIdx
				var name string
				if fieldIdx < len(jsonFieldsNameOfNewMessage) {
					name = jsonFieldsNameOfNewMessage[fieldIdx]
				} else {
					name = strconv.Itoa(fieldIdx)
				}
				failures = append(failures, validate.FieldError{
					Name:  name,
					Error: validate.ErrFieldRequired,
				})
				// Reset bit.
				result &^= 1 << bitIdx
			}
		}
	}
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}

	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s *NewMessage) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *NewMessage) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode encodes int64 as json.
func (o OptInt64) Encode(e *jx.Encoder) {
	if !o.Set {
		return
	}
	e.Int64(int64(o.Value))
}

// Decode decodes int64 from json.
func (o *OptInt64) Decode(d *jx.Decoder) error {
	if o == nil {
		return errors.New("invalid: unable to decode OptInt64 to nil")
	}
	o.Set = true
	v, err := d.Int64()
	if err != nil {
		return err
	}
	o.Value = int64(v)
	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s OptInt64) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *OptInt64) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode encodes string as json.
func (o OptString) Encode(e *jx.Encoder) {
	if !o.Set {
		return
	}
	e.Str(string(o.Value))
}

// Decode decodes string from json.
func (o *OptString) Decode(d *jx.Decoder) error {
	if o == nil {
		return errors.New("invalid: unable to decode OptString to nil")
	}
	o.Set = true
	v, err := d.Str()
	if err != nil {
		return err
	}
	o.Value = string(v)
	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s OptString) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *OptString) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode encodes ReceiveMessageBadRequest as json.
func (s *ReceiveMessageBadRequest) Encode(e *jx.Encoder) {
	unwrapped := (*Error)(s)

	unwrapped.Encode(e)
}

// Decode decodes ReceiveMessageBadRequest from json.
func (s *ReceiveMessageBadRequest) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode ReceiveMessageBadRequest to nil")
	}
	var unwrapped Error
	if err := func() error {
		if err := unwrapped.Decode(d); err != nil {
			return err
		}
		return nil
	}(); err != nil {
		return errors.Wrap(err, "alias")
	}
	*s = ReceiveMessageBadRequest(unwrapped)
	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s *ReceiveMessageBadRequest) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *ReceiveMessageBadRequest) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode encodes ReceiveMessageInternalServerError as json.
func (s *ReceiveMessageInternalServerError) Encode(e *jx.Encoder) {
	unwrapped := (*Error)(s)

	unwrapped.Encode(e)
}

// Decode decodes ReceiveMessageInternalServerError from json.
func (s *ReceiveMessageInternalServerError) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode ReceiveMessageInternalServerError to nil")
	}
	var unwrapped Error
	if err := func() error {
		if err := unwrapped.Decode(d); err != nil {
			return err
		}
		return nil
	}(); err != nil {
		return errors.Wrap(err, "alias")
	}
	*s = ReceiveMessageInternalServerError(unwrapped)
	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s *ReceiveMessageInternalServerError) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *ReceiveMessageInternalServerError) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode implements json.Marshaler.
func (s *ReceiveMessageOK) Encode(e *jx.Encoder) {
	e.ObjStart()
	s.encodeFields(e)
	e.ObjEnd()
}

// encodeFields encodes fields.
func (s *ReceiveMessageOK) encodeFields(e *jx.Encoder) {
	{
		e.FieldStart("result")
		e.Str(s.Result)
	}
	{
		e.FieldStart("messages")
		e.ArrStart()
		for _, elem := range s.Messages {
			elem.Encode(e)
		}
		e.ArrEnd()
	}
}

var jsonFieldsNameOfReceiveMessageOK = [2]string{
	0: "result",
	1: "messages",
}

// Decode decodes ReceiveMessageOK from json.
func (s *ReceiveMessageOK) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode ReceiveMessageOK to nil")
	}
	var requiredBitSet [1]uint8

	if err := d.ObjBytes(func(d *jx.Decoder, k []byte) error {
		switch string(k) {
		case "result":
			requiredBitSet[0] |= 1 << 0
			if err := func() error {
				v, err := d.Str()
				s.Result = string(v)
				if err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"result\"")
			}
		case "messages":
			requiredBitSet[0] |= 1 << 1
			if err := func() error {
				s.Messages = make([]Message, 0)
				if err := d.Arr(func(d *jx.Decoder) error {
					var elem Message
					if err := elem.Decode(d); err != nil {
						return err
					}
					s.Messages = append(s.Messages, elem)
					return nil
				}); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"messages\"")
			}
		default:
			return d.Skip()
		}
		return nil
	}); err != nil {
		return errors.Wrap(err, "decode ReceiveMessageOK")
	}
	// Validate required fields.
	var failures []validate.FieldError
	for i, mask := range [1]uint8{
		0b00000011,
	} {
		if result := (requiredBitSet[i] & mask) ^ mask; result != 0 {
			// Mask only required fields and check equality to mask using XOR.
			//
			// If XOR result is not zero, result is not equal to expected, so some fields are missed.
			// Bits of fields which would be set are actually bits of missed fields.
			missed := bits.OnesCount8(result)
			for bitN := 0; bitN < missed; bitN++ {
				bitIdx := bits.TrailingZeros8(result)
				fieldIdx := i*8 + bitIdx
				var name string
				if fieldIdx < len(jsonFieldsNameOfReceiveMessageOK) {
					name = jsonFieldsNameOfReceiveMessageOK[fieldIdx]
				} else {
					name = strconv.Itoa(fieldIdx)
				}
				failures = append(failures, validate.FieldError{
					Name:  name,
					Error: validate.ErrFieldRequired,
				})
				// Reset bit.
				result &^= 1 << bitIdx
			}
		}
	}
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}

	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s *ReceiveMessageOK) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *ReceiveMessageOK) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode encodes ReceiveMessageUnauthorized as json.
func (s *ReceiveMessageUnauthorized) Encode(e *jx.Encoder) {
	unwrapped := (*Error)(s)

	unwrapped.Encode(e)
}

// Decode decodes ReceiveMessageUnauthorized from json.
func (s *ReceiveMessageUnauthorized) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode ReceiveMessageUnauthorized to nil")
	}
	var unwrapped Error
	if err := func() error {
		if err := unwrapped.Decode(d); err != nil {
			return err
		}
		return nil
	}(); err != nil {
		return errors.Wrap(err, "alias")
	}
	*s = ReceiveMessageUnauthorized(unwrapped)
	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s *ReceiveMessageUnauthorized) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *ReceiveMessageUnauthorized) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode encodes SendMessageBadRequest as json.
func (s *SendMessageBadRequest) Encode(e *jx.Encoder) {
	unwrapped := (*Error)(s)

	unwrapped.Encode(e)
}

// Decode decodes SendMessageBadRequest from json.
func (s *SendMessageBadRequest) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode SendMessageBadRequest to nil")
	}
	var unwrapped Error
	if err := func() error {
		if err := unwrapped.Decode(d); err != nil {
			return err
		}
		return nil
	}(); err != nil {
		return errors.Wrap(err, "alias")
	}
	*s = SendMessageBadRequest(unwrapped)
	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s *SendMessageBadRequest) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *SendMessageBadRequest) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode encodes SendMessageInternalServerError as json.
func (s *SendMessageInternalServerError) Encode(e *jx.Encoder) {
	unwrapped := (*Error)(s)

	unwrapped.Encode(e)
}

// Decode decodes SendMessageInternalServerError from json.
func (s *SendMessageInternalServerError) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode SendMessageInternalServerError to nil")
	}
	var unwrapped Error
	if err := func() error {
		if err := unwrapped.Decode(d); err != nil {
			return err
		}
		return nil
	}(); err != nil {
		return errors.Wrap(err, "alias")
	}
	*s = SendMessageInternalServerError(unwrapped)
	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s *SendMessageInternalServerError) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *SendMessageInternalServerError) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode implements json.Marshaler.
func (s *SendMessageOK) Encode(e *jx.Encoder) {
	e.ObjStart()
	s.encodeFields(e)
	e.ObjEnd()
}

// encodeFields encodes fields.
func (s *SendMessageOK) encodeFields(e *jx.Encoder) {
	{
		e.FieldStart("result")
		e.Str(s.Result)
	}
	{
		e.FieldStart("message")
		s.Message.Encode(e)
	}
}

var jsonFieldsNameOfSendMessageOK = [2]string{
	0: "result",
	1: "message",
}

// Decode decodes SendMessageOK from json.
func (s *SendMessageOK) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode SendMessageOK to nil")
	}
	var requiredBitSet [1]uint8

	if err := d.ObjBytes(func(d *jx.Decoder, k []byte) error {
		switch string(k) {
		case "result":
			requiredBitSet[0] |= 1 << 0
			if err := func() error {
				v, err := d.Str()
				s.Result = string(v)
				if err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"result\"")
			}
		case "message":
			requiredBitSet[0] |= 1 << 1
			if err := func() error {
				if err := s.Message.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"message\"")
			}
		default:
			return d.Skip()
		}
		return nil
	}); err != nil {
		return errors.Wrap(err, "decode SendMessageOK")
	}
	// Validate required fields.
	var failures []validate.FieldError
	for i, mask := range [1]uint8{
		0b00000011,
	} {
		if result := (requiredBitSet[i] & mask) ^ mask; result != 0 {
			// Mask only required fields and check equality to mask using XOR.
			//
			// If XOR result is not zero, result is not equal to expected, so some fields are missed.
			// Bits of fields which would be set are actually bits of missed fields.
			missed := bits.OnesCount8(result)
			for bitN := 0; bitN < missed; bitN++ {
				bitIdx := bits.TrailingZeros8(result)
				fieldIdx := i*8 + bitIdx
				var name string
				if fieldIdx < len(jsonFieldsNameOfSendMessageOK) {
					name = jsonFieldsNameOfSendMessageOK[fieldIdx]
				} else {
					name = strconv.Itoa(fieldIdx)
				}
				failures = append(failures, validate.FieldError{
					Name:  name,
					Error: validate.ErrFieldRequired,
				})
				// Reset bit.
				result &^= 1 << bitIdx
			}
		}
	}
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}

	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s *SendMessageOK) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *SendMessageOK) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode encodes SendMessageUnauthorized as json.
func (s *SendMessageUnauthorized) Encode(e *jx.Encoder) {
	unwrapped := (*Error)(s)

	unwrapped.Encode(e)
}

// Decode decodes SendMessageUnauthorized from json.
func (s *SendMessageUnauthorized) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode SendMessageUnauthorized to nil")
	}
	var unwrapped Error
	if err := func() error {
		if err := unwrapped.Decode(d); err != nil {
			return err
		}
		return nil
	}(); err != nil {
		return errors.Wrap(err, "alias")
	}
	*s = SendMessageUnauthorized(unwrapped)
	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s *SendMessageUnauthorized) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *SendMessageUnauthorized) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode implements json.Marshaler.
func (s *SendRequest) Encode(e *jx.Encoder) {
	e.ObjStart()
	s.encodeFields(e)
	e.ObjEnd()
}

// encodeFields encodes fields.
func (s *SendRequest) encodeFields(e *jx.Encoder) {
	{
		e.FieldStart("content")
		s.Content.Encode(e)
	}
}

var jsonFieldsNameOfSendRequest = [1]string{
	0: "content",
}

// Decode decodes SendRequest from json.
func (s *SendRequest) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode SendRequest to nil")
	}
	var requiredBitSet [1]uint8

	if err := d.ObjBytes(func(d *jx.Decoder, k []byte) error {
		switch string(k) {
		case "content":
			requiredBitSet[0] |= 1 << 0
			if err := func() error {
				if err := s.Content.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"content\"")
			}
		default:
			return d.Skip()
		}
		return nil
	}); err != nil {
		return errors.Wrap(err, "decode SendRequest")
	}
	// Validate required fields.
	var failures []validate.FieldError
	for i, mask := range [1]uint8{
		0b00000001,
	} {
		if result := (requiredBitSet[i] & mask) ^ mask; result != 0 {
			// Mask only required fields and check equality to mask using XOR.
			//
			// If XOR result is not zero, result is not equal to expected, so some fields are missed.
			// Bits of fields which would be set are actually bits of missed fields.
			missed := bits.OnesCount8(result)
			for bitN := 0; bitN < missed; bitN++ {
				bitIdx := bits.TrailingZeros8(result)
				fieldIdx := i*8 + bitIdx
				var name string
				if fieldIdx < len(jsonFieldsNameOfSendRequest) {
					name = jsonFieldsNameOfSendRequest[fieldIdx]
				} else {
					name = strconv.Itoa(fieldIdx)
				}
				failures = append(failures, validate.FieldError{
					Name:  name,
					Error: validate.ErrFieldRequired,
				})
				// Reset bit.
				result &^= 1 << bitIdx
			}
		}
	}
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}

	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s *SendRequest) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *SendRequest) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}
