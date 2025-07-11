// Code generated by ogen, DO NOT EDIT.

package message

// SetFake set fake values.
func (s *DeleteMessageBadRequest) SetFake() {
	var unwrapped Error
	{
		unwrapped.SetFake()
	}
	*s = DeleteMessageBadRequest(unwrapped)
}

// SetFake set fake values.
func (s *DeleteMessageInternalServerError) SetFake() {
	var unwrapped Error
	{
		unwrapped.SetFake()
	}
	*s = DeleteMessageInternalServerError(unwrapped)
}

// SetFake set fake values.
func (s *DeleteMessageNotFound) SetFake() {
	var unwrapped Error
	{
		unwrapped.SetFake()
	}
	*s = DeleteMessageNotFound(unwrapped)
}

// SetFake set fake values.
func (s *DeleteMessageOK) SetFake() {
	{
		{
			s.Result = "string"
		}
	}
}

// SetFake set fake values.
func (s *DeleteMessageUnauthorized) SetFake() {
	var unwrapped Error
	{
		unwrapped.SetFake()
	}
	*s = DeleteMessageUnauthorized(unwrapped)
}

// SetFake set fake values.
func (s *Error) SetFake() {
	{
		{
			s.Code.SetFake()
		}
	}
	{
		{
			s.Message.SetFake()
		}
	}
}

// SetFake set fake values.
func (s *ExtendMessageTimeoutBadRequest) SetFake() {
	var unwrapped Error
	{
		unwrapped.SetFake()
	}
	*s = ExtendMessageTimeoutBadRequest(unwrapped)
}

// SetFake set fake values.
func (s *ExtendMessageTimeoutInternalServerError) SetFake() {
	var unwrapped Error
	{
		unwrapped.SetFake()
	}
	*s = ExtendMessageTimeoutInternalServerError(unwrapped)
}

// SetFake set fake values.
func (s *ExtendMessageTimeoutNotFound) SetFake() {
	var unwrapped Error
	{
		unwrapped.SetFake()
	}
	*s = ExtendMessageTimeoutNotFound(unwrapped)
}

// SetFake set fake values.
func (s *ExtendMessageTimeoutOK) SetFake() {
	{
		{
			s.Result = "string"
		}
	}
	{
		{
			s.Message.SetFake()
		}
	}
}

// SetFake set fake values.
func (s *ExtendMessageTimeoutUnauthorized) SetFake() {
	var unwrapped Error
	{
		unwrapped.SetFake()
	}
	*s = ExtendMessageTimeoutUnauthorized(unwrapped)
}

// SetFake set fake values.
func (s *Message) SetFake() {
	{
		{
			s.ID.SetFake()
		}
	}
	{
		{
			s.Content.SetFake()
		}
	}
	{
		{
			s.CreatedAt = int64(0)
		}
	}
	{
		{
			s.UpdatedAt = int64(0)
		}
	}
	{
		{
			s.ExpiresAt = int64(0)
		}
	}
	{
		{
			s.AcquiredAt = int64(0)
		}
	}
	{
		{
			s.VisibilityTimeoutAt = int64(0)
		}
	}
}

// SetFake set fake values.
func (s *MessageContent) SetFake() {
	var unwrapped string
	{
		unwrapped = "string"
	}
	*s = MessageContent(unwrapped)
}

// SetFake set fake values.
func (s *MessageId) SetFake() {
	var unwrapped string
	{
		unwrapped = "string"
	}
	*s = MessageId(unwrapped)
}

// SetFake set fake values.
func (s *NewMessage) SetFake() {
	{
		{
			s.ID.SetFake()
		}
	}
	{
		{
			s.Content.SetFake()
		}
	}
	{
		{
			s.CreatedAt = int64(0)
		}
	}
	{
		{
			s.UpdatedAt = int64(0)
		}
	}
	{
		{
			s.ExpiresAt = int64(0)
		}
	}
}

// SetFake set fake values.
func (s *OptInt64) SetFake() {
	var elem int64
	{
		elem = int64(0)
	}
	s.SetTo(elem)
}

// SetFake set fake values.
func (s *OptString) SetFake() {
	var elem string
	{
		elem = "string"
	}
	s.SetTo(elem)
}

// SetFake set fake values.
func (s *ReceiveMessageBadRequest) SetFake() {
	var unwrapped Error
	{
		unwrapped.SetFake()
	}
	*s = ReceiveMessageBadRequest(unwrapped)
}

// SetFake set fake values.
func (s *ReceiveMessageInternalServerError) SetFake() {
	var unwrapped Error
	{
		unwrapped.SetFake()
	}
	*s = ReceiveMessageInternalServerError(unwrapped)
}

// SetFake set fake values.
func (s *ReceiveMessageOK) SetFake() {
	{
		{
			s.Result = "string"
		}
	}
	{
		{
			s.Messages = nil
			for i := 0; i < 0; i++ {
				var elem Message
				{
					elem.SetFake()
				}
				s.Messages = append(s.Messages, elem)
			}
		}
	}
}

// SetFake set fake values.
func (s *ReceiveMessageUnauthorized) SetFake() {
	var unwrapped Error
	{
		unwrapped.SetFake()
	}
	*s = ReceiveMessageUnauthorized(unwrapped)
}

// SetFake set fake values.
func (s *SendMessageBadRequest) SetFake() {
	var unwrapped Error
	{
		unwrapped.SetFake()
	}
	*s = SendMessageBadRequest(unwrapped)
}

// SetFake set fake values.
func (s *SendMessageInternalServerError) SetFake() {
	var unwrapped Error
	{
		unwrapped.SetFake()
	}
	*s = SendMessageInternalServerError(unwrapped)
}

// SetFake set fake values.
func (s *SendMessageOK) SetFake() {
	{
		{
			s.Result = "string"
		}
	}
	{
		{
			s.Message.SetFake()
		}
	}
}

// SetFake set fake values.
func (s *SendMessageUnauthorized) SetFake() {
	var unwrapped Error
	{
		unwrapped.SetFake()
	}
	*s = SendMessageUnauthorized(unwrapped)
}

// SetFake set fake values.
func (s *SendRequest) SetFake() {
	{
		{
			s.Content.SetFake()
		}
	}
}
