package app

import (
	"go-admin-x/internal/controllers/common"
	models "go-admin-x/internal/model"
	"go-admin-x/internal/model/app"

	"go-admin-x/internal/util/hash"

	"github.com/gin-gonic/gin"
)

type Version struct{}

// 分页数据
func (Version) List(c *gin.Context) {
	page := common.GetPageIndex(c)
	limit := common.GetPageLimit(c)
	sort := common.GetPageSort(c)
	key := common.GetPageKey(c)
	status := common.GetQueryToUint(c, "status")
	var whereOrder []models.PageWhereOrder
	order := "ID DESC"
	if len(sort) >= 2 {
		orderType := sort[0:1]
		order = sort[1:len(sort)]
		if orderType == "+" {
			order += " ASC"
		} else {
			order += " DESC"
		}
	}
	whereOrder = append(whereOrder, models.PageWhereOrder{Order: order})
	if key != "" {
		v := "%" + key + "%"
		var arr []interface{}
		arr = append(arr, v)
		arr = append(arr, v)
		whereOrder = append(whereOrder, models.PageWhereOrder{Where: "user_name like ? or real_name like ?", Value: arr})
	}
	if status > 0 {
		var arr []interface{}
		arr = append(arr, status)
		whereOrder = append(whereOrder, models.PageWhereOrder{Where: "status = ?", Value: arr})
	}
	var total uint64
	list := []app.Version{}
	err := models.GetPage(&app.Version{}, &app.Version{}, &list, page, limit, &total, whereOrder...)
	if err != nil {
		common.ResErrSrv(c, err)
		return
	}
	common.ResSuccessPage(c, total, &list)
}

// 详情
func (Version) Detail(c *gin.Context) {
	id := common.GetQueryToUint64(c, "id")
	var model app.Version
	where := app.Version{}
	where.ID = id
	_, err := models.First(&where, &model)
	if err != nil {
		common.ResErrSrv(c, err)
		return
	}
	common.ResSuccess(c, &model)
}

// 更新
func (Version) Update(c *gin.Context) {
	model := app.Version{}
	err := c.Bind(&model)
	if err != nil {
		common.ResErrSrv(c, err)
		return
	}
	where := app.Version{}
	where.ID = model.ID
	modelOld := app.Version{}
	_, err = models.First(&where, &modelOld)
	if err != nil {
		common.ResErrSrv(c, err)
		return
	}
	model.UserName = modelOld.UserName
	model.Password = modelOld.Password
	err = models.Save(&model)
	if err != nil {
		common.ResFail(c, "操作失败")
		return
	}
	common.ResSuccessMsg(c)
}

//新增
func (Admin) Create(c *gin.Context) {
	model := app.Version{}
	err := c.Bind(&model)
	if err != nil {
		common.ResErrSrv(c, err)
		return
	}
	model.Password = hash.Md5String(common.MD5_PREFIX + model.Password)
	err = models.Create(&model)
	if err != nil {
		common.ResFail(c, "操作失败")
		return
	}
	common.ResSuccess(c, gin.H{"id": model.ID})
}

// 删除数据
func (Admin) Delete(c *gin.Context) {
	var ids []uint64
	err := c.Bind(&ids)
	if err != nil || len(ids) == 0 {
		common.ResErrSrv(c, err)
		return
	}
	admin := app.Version{}
	err = admin.Delete(ids)
	if err != nil {
		common.ResErrSrv(c, err)
		return
	}
	common.ResSuccessMsg(c)
}
