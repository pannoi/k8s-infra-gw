package util

import (
	"net/http"
	"os"
	"crypto/rand"
	"encoding/base64"

	"github.com/gofrs/uuid"
	"go.uber.org/zap"

	c "infra-gw/src/cont"

	corev1 "k8s.io/api/core/v1"
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

//Int32 returns a pointer to an int32
func Int32(i int32) *int32 {
	return &i
}
var Int32Ptr = Int32

func SecretKeyRef(name, key string) *corev1.EnvVarSource {
	return &corev1.EnvVarSource{SecretKeyRef: &corev1.SecretKeySelector{LocalObjectReference: corev1.LocalObjectReference{Name: name}, Key: key}}
}
