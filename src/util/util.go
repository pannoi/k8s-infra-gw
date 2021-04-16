package util

import (
	"net/http"
	"os"
	"crypto/rand"
	"encoding/base64"

	"github.com/gofrs/uuid"
	"go.uber.org/zap"

	c "infra-gw/src/cont"
)

// HomeDir returns homedir path
func HomeDir() string {
	if h := os.Getenv("HOME"); h != "" {
		return h
	}
	return os.Getenv("USERPROFILE") // windows
}

// LoggerWithRequestID returns a logger which has the requestID attached to it
func LoggerWithRequestID(log *zap.SugaredLogger, r *http.Request) *zap.SugaredLogger {
	requestID := r.Context().Value(c.KeyReqID)

	if requestID == nil {
		log.Debugw("Request has no requestID attached", "path", r.URL.Path)
		return log
	}

	return log.With("requestID", requestID.(uuid.UUID).String())
}

func MustGenerateRandomAsByte(len int) []byte {
	buff := make([]byte, len/2)
	_, err := rand.Read(buff)
	if err != nil {
		panic(err)
	}

	return []byte(base64.RawStdEncoding.Strict().EncodeToString(buff))
}
