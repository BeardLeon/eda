package http_handler

import (
	"go.uber.org/zap"
	"golang-template/src/edaPkg"
	"golang-template/src/zaplog"
	"net/http"
)

type EdaHandler struct {
	EdaPkg edaPkg.EdaPkg
}

var Logger = zaplog.Logger

func (h *EdaHandler) ImportJsonFile(w http.ResponseWriter, r *http.Request) {
	err := h.EdaPkg.ImportJsonFile(r.PostForm.Get("file"))
	if err != nil {
		Logger.Error("ImportJsonFile", zap.Any("", err))
		return
	}
	Logger.Info("ImportJsonFile ", zap.Any("title: ", h.EdaPkg.Title))
}
