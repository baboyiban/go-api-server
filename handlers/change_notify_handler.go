package handlers

import (
	"net/http"

	"github.com/baboyiban/go-api-server/service"
	"github.com/gin-gonic/gin"
)

type ChangeNotifyHandler struct {
	service *service.ChangeNotifyService
}

func NewChangeNotifyHandler(s *service.ChangeNotifyService) *ChangeNotifyHandler {
	return &ChangeNotifyHandler{service: s}
}

// ListChangeNotifies godoc
// @Summary      변경 이벤트 목록 조회
// @Description  change_notify 테이블의 모든 변경 이벤트를 반환합니다.
// @Tags         change_notify
// @Produce      json
// @Success      200  {array}   dto.ChangeNotifyResponse
// @Failure      500  {object}  dto.ErrorResponse
// @Router       /api/change-notify [get]
func (h *ChangeNotifyHandler) ListChangeNotifies(c *gin.Context) {
	notifies, err := h.service.ListChangeNotifies(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to list change_notify", "details": err.Error()})
		return
	}
	c.JSON(http.StatusOK, notifies)
}
