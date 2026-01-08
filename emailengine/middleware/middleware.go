package middleware

import (
	"github.com/stax-ai/emailengine/emailengine/internal/httptransport"
)

type Middleware = httptransport.RoundTripperMiddleware
