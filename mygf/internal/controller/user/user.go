package user

import (
	"context"
	"fmt"
	"github.com/gogf/gf/v2/frame/g"
	"mygf/api/v1"
	"mygf/internal/controller"
	"mygf/internal/dao"
	"mygf/internal/models"
)

//cmd.go调用接口
var (
	User    = user{}
	List    = uList{}
	Details = uDetails{}
)

type user struct {
}
type uList struct {
}
type uDetails struct {
}

//人员列表
func (c *uList) List(ctx context.Context, req *v1.UserListReq) (res *v1.UserListRes, err error) {
	data, err := models.User.GetList(ctx, req.IsTmp, req.Page, req.Rows)
	if err != nil {
		return
	}

	controller.BController.ReturnJson(ctx, controller.ReturnData{
		Code: 1,
		Msg:  "操作成功",
		Data: data,
	})

	return
}

//详情
func (c *uDetails) Details(ctx context.Context, req *v1.UserDetailsReq) (res *v1.UserDetailsRes, err error) {
	data, err := models.User.GetOne(ctx, req.Uuid)
	if err != nil {
		return
	}

	controller.BController.ReturnJson(ctx, controller.ReturnData{
		Code: 1,
		Msg:  "操作成功",
		Data: data,
	})

	return
}

//添加
func (c *user) Add(ctx context.Context, req *v1.UserAddReq) (res *v1.UserRes, err error) {
	params := g.Map{
		dao.User.Columns().Account:  "account",
		dao.User.Columns().Username: "username",
		dao.User.Columns().Password: "password",
		dao.User.Columns().Unit:     "unit",
	}

	err = models.User.Add(ctx, params)
	if err != nil {
		return
	}

	controller.BController.ReturnJson(ctx, controller.ReturnData{
		Code: 1,
		Msg:  "添加成功",
		Data: g.Array{},
	})

	return
}

//编辑
func (c *user) Edit(ctx context.Context, req *v1.UserEditReq) (res *v1.UserRes, err error) {
	err = models.User.Save(ctx, g.Map{dao.User.Columns().Uuid: req.Uuid}, g.Map{dao.User.Columns().Username: req.Username})
	if err != nil {
		return
	}

	controller.BController.ReturnJson(ctx, controller.ReturnData{
		Code: 1,
		Msg:  "保存成功",
		Data: g.Array{},
	})

	return
}

//删除
func (c *user) Delete(ctx context.Context, req *v1.UserDeleteReq) (res *v1.UserRes, err error) {
	fmt.Println("req:", req)
	err = models.User.Delete(ctx, g.Map{dao.User.Columns().Uuid: req.Uuid})
	if err != nil {
		return
	}

	controller.BController.ReturnJson(ctx, controller.ReturnData{
		Code: 1,
		Msg:  "删除成功",
		Data: g.Array{},
	})

	return
}
