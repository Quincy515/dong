import service from '@/utils/request'

// @Tags WenJianZiLiaoJieDianBiao
// @Summary 创建WenJianZiLiaoJieDianBiao
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.WenJianZiLiaoJieDianBiao true "创建WenJianZiLiaoJieDianBiao"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /wenJianZiLiaoJieDianBiao/createWenJianZiLiaoJieDianBiao [post]
export const createWenJianZiLiaoJieDianBiao = (data) => {
  return service({
    url: '/wenJianZiLiaoJieDianBiao/createWenJianZiLiaoJieDianBiao',
    method: 'post',
    data
  })
}

// @Tags WenJianZiLiaoJieDianBiao
// @Summary 删除WenJianZiLiaoJieDianBiao
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.WenJianZiLiaoJieDianBiao true "删除WenJianZiLiaoJieDianBiao"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /wenJianZiLiaoJieDianBiao/deleteWenJianZiLiaoJieDianBiao [delete]
export const deleteWenJianZiLiaoJieDianBiao = (data) => {
  return service({
    url: '/wenJianZiLiaoJieDianBiao/deleteWenJianZiLiaoJieDianBiao',
    method: 'delete',
    data
  })
}

// @Tags WenJianZiLiaoJieDianBiao
// @Summary 删除WenJianZiLiaoJieDianBiao
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除WenJianZiLiaoJieDianBiao"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /wenJianZiLiaoJieDianBiao/deleteWenJianZiLiaoJieDianBiao [delete]
export const deleteWenJianZiLiaoJieDianBiaoByIds = (data) => {
  return service({
    url: '/wenJianZiLiaoJieDianBiao/deleteWenJianZiLiaoJieDianBiaoByIds',
    method: 'delete',
    data
  })
}

// @Tags WenJianZiLiaoJieDianBiao
// @Summary 更新WenJianZiLiaoJieDianBiao
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.WenJianZiLiaoJieDianBiao true "更新WenJianZiLiaoJieDianBiao"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /wenJianZiLiaoJieDianBiao/updateWenJianZiLiaoJieDianBiao [put]
export const updateWenJianZiLiaoJieDianBiao = (data) => {
  return service({
    url: '/wenJianZiLiaoJieDianBiao/updateWenJianZiLiaoJieDianBiao',
    method: 'put',
    data
  })
}

// @Tags WenJianZiLiaoJieDianBiao
// @Summary 用id查询WenJianZiLiaoJieDianBiao
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.WenJianZiLiaoJieDianBiao true "用id查询WenJianZiLiaoJieDianBiao"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /wenJianZiLiaoJieDianBiao/findWenJianZiLiaoJieDianBiao [get]
export const findWenJianZiLiaoJieDianBiao = (params) => {
  return service({
    url: '/wenJianZiLiaoJieDianBiao/findWenJianZiLiaoJieDianBiao',
    method: 'get',
    params
  })
}

// @Tags WenJianZiLiaoJieDianBiao
// @Summary 分页获取WenJianZiLiaoJieDianBiao列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取WenJianZiLiaoJieDianBiao列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /wenJianZiLiaoJieDianBiao/getWenJianZiLiaoJieDianBiaoList [get]
export const getWenJianZiLiaoJieDianBiaoList = (params) => {
  return service({
    url: '/wenJianZiLiaoJieDianBiao/getWenJianZiLiaoJieDianBiaoList',
    method: 'get',
    params
  })
}
