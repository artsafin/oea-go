package web

import (
	"context"
	"github.com/google/uuid"
	"go.uber.org/zap"
	"net/http"
	"oea-go/internal/common"
	"oea-go/resources"
	"strings"
)

func faviconMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		if request.RequestURI == "/favicon.ico" {
			writer.Header().Add("Content-Type", "image/png")
			writer.WriteHeader(http.StatusOK)
			writer.Write(resources.MustReadBytes("assets/icon.png"))
			return
		}

		next.ServeHTTP(writer, request)
	})
}

func cssMapRejectorMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		if strings.HasSuffix(request.RequestURI, ".css.map") {
			writer.WriteHeader(http.StatusNotFound)
			return
		}

		next.ServeHTTP(writer, request)
	})
}

type loggerMiddleware struct {
	logger *zap.SugaredLogger
}

func (m loggerMiddleware) MiddlewareFunc(next http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, r *http.Request) {
		logger := common.NewRequestLogger(r, m.logger)
		logger.Infof("Request: %v %v %v %v", r.Method, r.RequestURI, r.RemoteAddr, r.UserAgent())

		next.ServeHTTP(writer, r)
	})
}

func requestIdMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, r *http.Request) {
		requestId := r.Header.Get(RequestIdHeader)
		if requestId == "" {
			requestId = uuid.Must(uuid.NewRandom()).String()
		}

		ctx := context.WithValue(r.Context(), common.RequestIdContextKey, requestId)

		writer.Header().Set(RequestIdHeader, requestId)

		next.ServeHTTP(writer, r.WithContext(ctx))
	})
}
