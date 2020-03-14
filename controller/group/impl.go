package group

import (
	"net/http"

	"github.com/acm-uiuc/core/context"
	"github.com/acm-uiuc/core/service"
)

type GroupController struct {
	svc *service.Service
}

func New(svc *service.Service) *GroupController {
	return &GroupController{
		svc: svc,
	}
}

func (controller *GroupController) GetGroups(ctx *context.Context) error {
	groups, err := controller.svc.Group.GetGroups()
	if err != nil {
		return ctx.String(http.StatusBadRequest, "Failed Group Lookup")
	}

	return ctx.JSON(http.StatusOK, &groups)
}

func (controller *GroupController) VerifyMembership(ctx *context.Context) error {
	req := struct {
		Username  string `json:"username"`
		GroupType string `json:"group_type`
		GroupName string `json:"group_name`
	}{}

	err := ctx.Bind(&req)
	if err != nil {
		return ctx.String(http.StatusBadRequest, "Failed Bind")
	}

	isMember, err := controller.svc.Group.VerifyMembership(req.Username, req.GroupType, req.GroupName)
	if err != nil {
		return ctx.String(http.StatusBadRequest, "Failed Membership Verification")
	}

	return ctx.JSON(http.StatusOK, &struct {
		IsMember bool `json:"is_member"`
	}{
		IsMember: isMember,
	})
}
