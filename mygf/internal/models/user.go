package models

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
	"mygf/internal/dao"
	"mygf/internal/model/entity"
)

type user struct {
}

var (
	User = user{}
)

//列表
func (u *user) GetList(ctx context.Context, isTmp int, page int, rows int) (data []*entity.User, err error) {
	err = dao.User.Ctx(ctx).Where(dao.User.Columns().IsTmp, isTmp).Limit(page, rows).Scan(&data)
	return
}

//详情
func (u *user) GetOne(ctx context.Context, uuid int) (data *entity.User, err error) {
	err = dao.User.Ctx(ctx).Where(dao.User.Columns().Uuid, uuid).Scan(&data)

	return
}

//添加
func (u *user) Add(ctx context.Context, params interface{}) (err error) {
	lastId, err := dao.User.Ctx(ctx).InsertAndGetId(params)
	if err != nil {
		return
	}

	_, err = dao.User.Ctx(ctx).Where(dao.User.Columns().Id, lastId).Data(g.Map{
		dao.User.Columns().Uuid: "101" + gconv.String(lastId),
	}).Update()

	return
}

//编辑
func (u *user) Save(ctx context.Context, where interface{}, params interface{}) (err error) {
	_, err = dao.User.Ctx(ctx).Where(where).Data(params).Update()
	return
}

//删除
func (u *user) Delete(ctx context.Context, where interface{}) (err error) {
	_, err = dao.User.Ctx(ctx).Where(where).Delete()
	return
}
