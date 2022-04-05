package dashboard

import (
	"net/http"

	"github.com/avatardev/ipos-mblb-backend/pkg/errors"
	"github.com/avatardev/ipos-mblb-backend/pkg/util/privutil"
	"github.com/avatardev/ipos-mblb-backend/pkg/util/responseutil"
)

type DashboardHandler struct {
	Service DashboardService
}

func NewDashboardHandler(service DashboardService) DashboardHandler {
	return DashboardHandler{Service: service}
}

func (d *DashboardHandler) GetStatistics() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if !privutil.CheckUserPrivilege(r.Context(), privutil.USER_ADMIN, privutil.USER_CHECKER) {
			responseutil.WriteErrorResponse(w, errors.ErrUserPriv)
			return
		}

		res, err := d.Service.GetStatistics(r.Context())
		if err != nil {
			responseutil.WriteErrorResponse(w, err)
			return
		}

		responseutil.WriteSuccessResponse(w, http.StatusOK, res)
	}
}
