package router

import (
	"net/http"
	"strconv"
	"strings"

	"git.containerum.net/ch/permissions/pkg/errors"
	"git.containerum.net/ch/permissions/pkg/model"
	"git.containerum.net/ch/permissions/pkg/server"
	"github.com/containerum/cherry/adaptors/gonic"
	"github.com/containerum/utils/httputil"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

type volumeHandlers struct {
	tv   *TranslateValidate
	acts server.VolumeActions
}

func (vh *volumeHandlers) createVolumeHandler(ctx *gin.Context) {
	var req model.VolumeCreateRequest
	if err := ctx.ShouldBindWith(&req, binding.JSON); err != nil {
		ctx.AbortWithStatusJSON(vh.tv.BadRequest(ctx, err))
		return
	}
	if err := vh.acts.CreateVolume(ctx.Request.Context(), req); err != nil {
		ctx.AbortWithStatusJSON(vh.tv.HandleError(err))
		return
	}

	ctx.Status(http.StatusCreated)
}

func (vh *volumeHandlers) getVolumeHandler(ctx *gin.Context) {
	ret, err := vh.acts.GetVolume(ctx.Request.Context(), ctx.Param("label"))
	if err != nil {
		ctx.AbortWithStatusJSON(vh.tv.HandleError(err))
		return
	}

	httputil.MaskForNonAdmin(ctx, &ret)

	ctx.JSON(http.StatusOK, ret)
}

func (vh *volumeHandlers) getUserVolumesHandler(ctx *gin.Context) {
	ret, err := vh.acts.GetUserVolumes(ctx.Request.Context(), strings.Split(ctx.Param("filter"), ",")...)

	if err != nil {
		ctx.AbortWithStatusJSON(vh.tv.HandleError(err))
		return
	}

	for i := range ret {
		httputil.MaskForNonAdmin(ctx, &ret[i])
	}

	ctx.JSON(http.StatusOK, ret)
}

func (vh *volumeHandlers) getAllVolumesHandler(ctx *gin.Context) {
	page, err := strconv.Atoi(ctx.Query("page"))
	if err != nil {
		gonic.Gonic(errors.ErrRequestValidationFailed().AddDetailF("page number not integer"), ctx)
		return
	}
	perPage, err := strconv.Atoi(ctx.Query("per_page"))
	if err != nil {
		gonic.Gonic(errors.ErrRequestValidationFailed().AddDetailF("per page limit not integer"), ctx)
		return
	}

	ret, err := vh.acts.GetAllVolumes(ctx.Request.Context(), page, perPage, strings.Split(ctx.Param("filter"), ",")...)

	if err != nil {
		ctx.AbortWithStatusJSON(vh.tv.HandleError(err))
		return
	}

	ctx.JSON(http.StatusOK, ret)
}

func (r *Router) SetupVolumeHandlers(acts server.VolumeActions) {
	handlers := &volumeHandlers{tv: r.tv, acts: acts}

	// swagger:operation POST /volumes Volumes CreateVolume
	//
	// Create Volume for User by Tariff.
	// Should be chosen first storage, where free space allows to create volume with provided capacity.
	//
	// ---
	// parameters:
	//  - $ref: '#/parameters/UserIDHeader'
	//  - $ref: '#/parameters/UserRoleHeader'
	//  - $ref: '#/parameters/SubstitutedUserID'
	//  - name: body
	//    in: body
	//    required: true
	//    schema:
	//      $ref: '#/definitions/VolumeCreateRequest'
	// responses:
	//   '201':
	//     description: volume created
	//   default:
	//     $ref: '#/responses/error'
	r.engine.POST("/volumes", handlers.createVolumeHandler)

	// swagger:operation GET /volumes/{label} Volumes GetVolume
	//
	// Get volume.
	//
	// ---
	// parameters:
	//  - $ref: '#/parameters/UserIDHeader'
	//  - $ref: '#/parameters/UserRoleHeader'
	//  - $ref: '#/parameters/SubstitutedUserID'
	//  - name: label
	//    in: path
	//    required: true
	//    type: string
	// responses:
	//   '200':
	//     description: volume response
	//     schema:
	//       $ref: '#/definitions/VolumeWithPermissions'
	//   default:
	//     $ref: '#/responses/error'
	r.engine.GET("/volumes/:volume", handlers.getVolumeHandler)

	// swagger:operation GET /volumes Volumes GetUserVolumes
	//
	// Get user volumes.
	//
	// ---
	// parameters:
	//  - $ref: '#/parameters/UserIDHeader'
	//  - $ref: '#/parameters/UserRoleHeader'
	//  - $ref: '#/parameters/SubstitutedUserID'
	//  - name: label
	//    in: path
	//    required: true
	//    type: string
	//  - $ref: '#/parameters/Filters'
	// responses:
	//   '200':
	//     description: volumes response
	//     schema:
	//       type: array
	//       items:
	//         $ref: '#/definitions/VolumeWithPermissions'
	//   default:
	//     $ref: '#/responses/error'
	r.engine.GET("/volumes", handlers.getUserVolumesHandler)

	// swagger:operation GET /admin/volumes Volumes GetAllVolumes
	//
	// Get all volumes (admin only).
	//
	// ---
	// parameters:
	//  - $ref: '#/parameters/UserIDHeader'
	//  - $ref: '#/parameters/UserRoleHeader'
	//  - $ref: '#/parameters/SubstitutedUserID'
	//  - name: label
	//    in: path
	//    required: true
	//    type: string
	//  - $ref: '#/parameters/Filters'
	//  - $ref: '#/parameters/PageNum'
	//  - $ref: '#/parameters/PerPageLimit'
	// responses:
	//   '200':
	//     description: volumes response
	//     schema:
	//       type: array
	//       items:
	//         $ref: '#/definitions/VolumeWithPermissions'
	//   default:
	//     $ref: '#/responses/error'
	r.engine.GET("/admin/volumes", httputil.RequireAdminRole(errors.ErrAdminRequired), handlers.getAllVolumesHandler)
}
