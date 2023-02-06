package delivery

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/win30221/core/http/ctx"
	"github.com/win30221/core/http/response"
	"github.com/win30221/core/http/validate"
	"github.com/win30221/ichiban-kuji/domain"
	"github.com/win30221/ichiban-kuji/service/usecase"
)

type Delivery struct {
	useCase *usecase.UseCase
}

func New(r *gin.RouterGroup, useCase *usecase.UseCase) {
	d := &Delivery{
		useCase: useCase,
	}

	r.GET("/users/:userID", d.GetUsersIchibanKuji)

	r.GET("/:ichibanKujiID", d.GetIchibanKuji)
	r.POST("", d.CreateIchibanKuji)
	r.PATCH("/:ichibanKujiID", d.UpdateIchibanKuji)
	r.DELETE("/:ichibanKujiID", d.DeleteIchibanKuji)

	r.POST("/:ichibanKujiID/rewards", d.CreateIchibanKujiReward)
	r.GET("/:ichibanKujiID/rewards", d.GetIchibanKujiRewards)
	r.GET("/rewards/:ichibanKujiRewardID", d.GetIchibanKujiReward)
	r.DELETE("/:ichibanKujiID/rewards", d.DeleteIchibanKujiRewards)
	r.PATCH("/rewards/:ichibanKujiRewardID", d.UpdateIchibanKujiReward)
	r.DELETE("/rewards/:ichibanKujiRewardID", d.DeleteIchibanKujiReward)
}

// GetUsersIchibanKuji godoc
// @Summary 取得用戶一番賞列表
// @Security Systoken
// @Param userID path string true " "
// @Success 200 {object} []domain.IchibanKuji
// @Router /ichiban-kuji/{userID} [get]
func (d *Delivery) GetUsersIchibanKuji(c *gin.Context) {
	ctx := ctx.New(c, c.Request.Context())

	userID := c.Param("userID")

	res, err := d.useCase.GetUsersIchibanKuji(ctx, userID)
	if err != nil {
		response.Error(ctx, http.StatusBadRequest, err)
		return
	}

	response.OK(ctx, res)
}

// CreateIchibanKuji godoc
// @Summary 建立一番賞
// @Security Systoken
// @Param - formData domain.CreateIchibanKujiReq true " "
// @Success 200 {string} success
// @Router /ichiban-kuji/ [post]
func (d *Delivery) CreateIchibanKuji(c *gin.Context) {
	ctx := ctx.New(c, c.Request.Context())
	req := domain.CreateIchibanKujiReq{}

	err := c.ShouldBind(&req)
	if err != nil {
		response.BindParameterError(ctx, err)
		return
	}

	err = validate.Struct(req)
	if err != nil {
		response.ValidParameterError(ctx, err)
		return
	}

	res, err := d.useCase.CreateIchibanKuji(ctx, req)
	if err != nil {
		response.Error(ctx, http.StatusBadRequest, err)
		return
	}

	response.OK(ctx, res)
}

// GetIchibanKuji godoc
// @Summary 取得一番賞資訊
// @Security Systoken
// @Param ichibanKujiID path string true " "
// @Success 200 {object} domain.IchibanKuji
// @Router /ichiban-kuji/{userID} [get]
func (d *Delivery) GetIchibanKuji(c *gin.Context) {
	ctx := ctx.New(c, c.Request.Context())

	ichibanKujiID := c.Param("ichibanKujiID")

	res, err := d.useCase.GetIchibanKuji(ctx, ichibanKujiID)
	if err != nil {
		response.Error(ctx, http.StatusBadRequest, err)
		return
	}

	response.OK(ctx, res)
}

// UpdateIchibanKuji godoc
// @Summary 修改一番賞
// @Security Systoken
// @Param - formData domain.UpdateIchibanKujiReq true " "
// @Success 200 {string} success
// @Router /ichiban-kuji/{ichibanKujiID} [patch]
func (d *Delivery) UpdateIchibanKuji(c *gin.Context) {
	ctx := ctx.New(c, c.Request.Context())
	req := domain.UpdateIchibanKujiReq{}

	err := c.ShouldBind(&req)
	if err != nil {
		response.BindParameterError(ctx, err)
		return
	}

	err = validate.Struct(req)
	if err != nil {
		response.ValidParameterError(ctx, err)
		return
	}

	ichibanKujiID := c.Param("ichibanKujiID")

	err = d.useCase.UpdateIchibanKuji(ctx, ichibanKujiID, req)
	if err != nil {
		response.Error(ctx, http.StatusBadRequest, err)
		return
	}

	response.OK(ctx, nil)
}

// DeleteIchibanKuji godoc
// @Summary 刪除一番賞
// @Security Systoken
// @Param - formData domain.UpdateIchibanKujiRewardReq true " "
// @Success 200 {string} success
// @Router /ichiban-kuji/rewards/{ichibanKujiRewardID} [patch]
func (d *Delivery) DeleteIchibanKuji(c *gin.Context) {
	ctx := ctx.New(c, c.Request.Context())

	ichibanKujiID := c.Param("ichibanKujiID")

	err := d.useCase.DeleteIchibanKuji(ctx, ichibanKujiID)
	if err != nil {
		response.Error(ctx, http.StatusBadRequest, err)
		return
	}

	response.OK(ctx, nil)
}

// CreateIchibanKuji godoc
// @Summary 建立一番賞獎品
// @Security Systoken
// @Param - formData domain.CreateIchibanKujiRewardReq true " "
// @Success 200 {string} success
// @Router /ichiban-kuji/reward/{ichibanKujiID} [post]
func (d *Delivery) CreateIchibanKujiReward(c *gin.Context) {
	ctx := ctx.New(c, c.Request.Context())
	req := domain.CreateIchibanKujiRewardReq{}

	err := c.ShouldBind(&req)
	if err != nil {
		response.BindParameterError(ctx, err)
		return
	}

	err = validate.Struct(req)
	if err != nil {
		response.ValidParameterError(ctx, err)
		return
	}

	ichibanKujiID := c.Param("ichibanKujiID")

	err = d.useCase.CreateIchibanKujiReward(ctx, ichibanKujiID, req)
	if err != nil {
		response.Error(ctx, http.StatusBadRequest, err)
		return
	}

	response.OK(ctx, nil)
}

// GetIchibanKujiRewards godoc
// @Summary 取得一番賞獎品列表
// @Security Systoken
// @Param userID path string true " "
// @Success 200 {object} []domain.IchibanKuji
// @Router /ichiban-kuji/{userID} [get]
func (d *Delivery) GetIchibanKujiRewards(c *gin.Context) {
	ctx := ctx.New(c, c.Request.Context())

	ichibanKujiID := c.Param("ichibanKujiID")

	res, err := d.useCase.GetIchibanKujiRewards(ctx, ichibanKujiID)
	if err != nil {
		response.Error(ctx, http.StatusBadRequest, err)
		return
	}

	response.OK(ctx, res)
}

// GetIchibanKujiReward godoc
// @Summary 取得一番賞獎品
// @Security Systoken
// @Param userID path string true " "
// @Success 200 {object} []domain.IchibanKuji
// @Router /ichiban-kuji/{userID} [get]
func (d *Delivery) GetIchibanKujiReward(c *gin.Context) {
	ctx := ctx.New(c, c.Request.Context())

	ichibanKujiRewardID := c.Param("ichibanKujiRewardID")

	res, err := d.useCase.GetIchibanKujiReward(ctx, ichibanKujiRewardID)
	if err != nil {
		response.Error(ctx, http.StatusBadRequest, err)
		return
	}

	response.OK(ctx, res)
}

// DeleteIchibanKujiRewards godoc
// @Summary 刪除一番賞獎品列表
// @Security Systoken
// @Param - formData domain.UpdateIchibanKujiRewardReq true " "
// @Success 200 {string} success
// @Router /ichiban-kuji/rewards/{ichibanKujiRewardID} [patch]
func (d *Delivery) DeleteIchibanKujiRewards(c *gin.Context) {
	ctx := ctx.New(c, c.Request.Context())

	ichibanKujiID := c.Param("ichibanKujiID")

	err := d.useCase.DeleteIchibanKujiRewards(ctx, ichibanKujiID)
	if err != nil {
		response.Error(ctx, http.StatusBadRequest, err)
		return
	}

	response.OK(ctx, nil)
}

// UpdateIchibanKujiReward godoc
// @Summary 修改一番賞獎品內容
// @Security Systoken
// @Param - formData domain.UpdateIchibanKujiRewardReq true " "
// @Success 200 {string} success
// @Router /ichiban-kuji/{ichibanKujiRewardID} [patch]
func (d *Delivery) UpdateIchibanKujiReward(c *gin.Context) {
	ctx := ctx.New(c, c.Request.Context())
	req := domain.UpdateIchibanKujiRewardReq{}

	err := c.ShouldBind(&req)
	if err != nil {
		response.BindParameterError(ctx, err)
		return
	}

	err = validate.Struct(req)
	if err != nil {
		response.ValidParameterError(ctx, err)
		return
	}

	ichibanKujiRewardID := c.Param("ichibanKujiRewardID")

	err = d.useCase.UpdateIchibanKujiReward(ctx, ichibanKujiRewardID, req)
	if err != nil {
		response.Error(ctx, http.StatusBadRequest, err)
		return
	}

	response.OK(ctx, nil)
}

// DeleteIchibanKujiReward godoc
// @Summary 刪除一番賞獎品
// @Security Systoken
// @Param - formData domain.UpdateIchibanKujiRewardReq true " "
// @Success 200 {string} success
// @Router /ichiban-kuji/rewards/{ichibanKujiRewardID} [patch]
func (d *Delivery) DeleteIchibanKujiReward(c *gin.Context) {
	ctx := ctx.New(c, c.Request.Context())

	ichibanKujiRewardID := c.Param("ichibanKujiRewardID")

	err := d.useCase.DeleteIchibanKujiReward(ctx, ichibanKujiRewardID)
	if err != nil {
		response.Error(ctx, http.StatusBadRequest, err)
		return
	}

	response.OK(ctx, nil)
}
