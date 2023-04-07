import service from '@/utils/request'

// @Tags WenJianZiLiaoBiao
// @Summary 创建WenJianZiLiaoBiao
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.WenJianZiLiaoBiao true "创建WenJianZiLiaoBiao"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /wenJianZiLiaoBiao/createWenJianZiLiaoBiao [post]
export const createWenJianZiLiaoBiao = (data) => {
  return service({
    url: '/wenJianZiLiaoBiao/createWenJianZiLiaoBiao',
    method: 'post',
    data
  })
}

// @Tags WenJianZiLiaoBiao
// @Summary 删除WenJianZiLiaoBiao
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.WenJianZiLiaoBiao true "删除WenJianZiLiaoBiao"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /wenJianZiLiaoBiao/deleteWenJianZiLiaoBiao [delete]
export const deleteWenJianZiLiaoBiao = (data) => {
  return service({
    url: '/wenJianZiLiaoBiao/deleteWenJianZiLiaoBiao',
    method: 'delete',
    data
  })
}

// @Tags WenJianZiLiaoBiao
// @Summary 删除WenJianZiLiaoBiao
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除WenJianZiLiaoBiao"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /wenJianZiLiaoBiao/deleteWenJianZiLiaoBiao [delete]
export const deleteWenJianZiLiaoBiaoByIds = (data) => {
  return service({
    url: '/wenJianZiLiaoBiao/deleteWenJianZiLiaoBiaoByIds',
    method: 'delete',
    data
  })
}

// @Tags WenJianZiLiaoBiao
// @Summary 更新WenJianZiLiaoBiao
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.WenJianZiLiaoBiao true "更新WenJianZiLiaoBiao"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /wenJianZiLiaoBiao/updateWenJianZiLiaoBiao [put]
export const updateWenJianZiLiaoBiao = (data) => {
  return service({
    url: '/wenJianZiLiaoBiao/updateWenJianZiLiaoBiao',
    method: 'put',
    data
  })
}

// @Tags WenJianZiLiaoBiao
// @Summary 用id查询WenJianZiLiaoBiao
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.WenJianZiLiaoBiao true "用id查询WenJianZiLiaoBiao"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /wenJianZiLiaoBiao/findWenJianZiLiaoBiao [get]
export const findWenJianZiLiaoBiao = (params) => {
  return service({
    url: '/wenJianZiLiaoBiao/findWenJianZiLiaoBiao',
    method: 'get',
    params
  })
}

// @Tags WenJianZiLiaoBiao
// @Summary 分页获取WenJianZiLiaoBiao列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取WenJianZiLiaoBiao列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /wenJianZiLiaoBiao/getWenJianZiLiaoBiaoList [get]
export const getWenJianZiLiaoBiaoList = (params) => {
  return service({
    url: '/wenJianZiLiaoBiao/getWenJianZiLiaoBiaoList',
    method: 'get',
    params
  })
}
