package app

import (
	"encoding/json"
	"errors"
	"net/http"
	"time"

	"github.com/doktorupnos/crow/backend/internal/respond"
)

func HandlePanic(w http.ResponseWriter, r *http.Request) {
	panic("The server automatically recovers from panics")
}

func HandleError(w http.ResponseWriter, _ *http.Request) {
	const s = http.StatusInternalServerError
	respond.Error(w, s, errors.New(http.StatusText(s)))
}

type SleepRequest struct {
	Duration string `json:"duration"`
}

func HandleSleep(w http.ResponseWriter, r *http.Request) {
	req := &SleepRequest{}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		respond.Error(w, http.StatusBadRequest, err)
		return
	}

	dur, err := time.ParseDuration(req.Duration)
	if err != nil {
		respond.Error(w, http.StatusBadRequest, err)
		return
	}

	time.Sleep(dur)

	w.WriteHeader(http.StatusOK)
}
