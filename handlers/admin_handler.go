package handlers

import (
	"accounting-api-with-go/internal/eventstore"
	"accounting-api-with-go/internal/services"
	"accounting-api-with-go/internal/utils"
	"net/http"
	"os"
)

type AdminHandler struct {
    ReplaySvc *services.ReplayService
    EventStore eventstore.EventStore
}

func NewAdminHandler(rs *services.ReplayService, es eventstore.EventStore) *AdminHandler {
    return &AdminHandler{ReplaySvc: rs, EventStore: es}
}

func (h *AdminHandler) Replay(w http.ResponseWriter, r *http.Request) {
	   token := r.Header.Get("X-Admin-Replay-Token")
    if token == "" || token != os.Getenv("ADMIN_REPLAY_TOKEN") {
        http.Error(w, "Unauthorized", http.StatusUnauthorized)
        return
    }
	
	err := h.EventStore.ReplayEvents(r.Context(), func(e eventstore.Event) error {
    	return h.ReplaySvc.ApplyEvent(r.Context(), e)
	})

    if err != nil {
		utils.WriteErrorResponse(w, "Replay failed:"+err.Error(), http.StatusInternalServerError)
        return
    }

	utils.WriteSuccessResponse(w, "Replay completed", http.StatusOK)
}