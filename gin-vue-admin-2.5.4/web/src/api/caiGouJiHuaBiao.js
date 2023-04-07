import service from '@/utils/request'

// @Tags CaiGouJiHuaBiao
// @Summary 创建CaiGouJiHuaBiao
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.CaiGouJiHuaBiao true "创建CaiGouJiHuaBiao"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /caiGouJiHuaBiao/createCaiGouJiHuaBiao [post]
export const createCaiGouJiHuaBiao = (data) => {
  return service({
    url: '/caiGouJiHuaBiao/createCaiGouJiHuaBiao',
    method: 'post',
    data
  })
}

// @Tags CaiGouJiHuaBiao
// @Summary 删除CaiGouJiHuaBiao
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.CaiGouJiHuaBiao true "删除CaiGouJiHuaBiao"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /caiGouJiHuaBiao/deleteCaiGouJiHuaBiao [delete]
export const deleteCaiGouJiHuaBiao = (data) => {
  return service({
    url: '/caiGouJiHuaBiao/deleteCaiGouJiHuaBiao',
    method: 'delete',
    data
  })
}

// @Tags CaiGouJiHuaBiao
// @Summary 删除CaiGouJiHuaBiao
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除CaiGouJiHuaBiao"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /caiGouJiHuaBiao/deleteCaiGouJiHuaBiao [delete]
export const deleteCaiGouJiHuaBiaoByIds = (data) => {
  return service({
    url: '/caiGouJiHuaBiao/deleteCaiGouJiHuaBiaoByIds',
    method: 'delete',
    data
  })
}

// @Tags CaiGouJiHuaBiao
// @Summary 更新CaiGouJiHuaBiao
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.CaiGouJiHuaBiao true "更新CaiGouJiHuaBiao"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /caiGouJiHuaBiao/updateCaiGouJiHuaBiao [put]
export const updateCaiGouJiHuaBiao = (data) => {
  return service({
    url: '/caiGouJiHuaBiao/updateCaiGouJiHuaBiao',
    method: 'put',
    data
  })
}

// @Tags CaiGouJiHuaBiao
// @Summary 用id查询CaiGouJiHuaBiao
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.CaiGouJiHuaBiao true "用id查询CaiGouJiHuaBiao"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /caiGouJiHuaBiao/findCaiGouJiHuaBiao [get]
export const findCaiGouJiHuaBiao = (params) => {
  return service({
    url: '/caiGouJiHuaBiao/findCaiGouJiHuaBiao',
    method: 'get',
    params
  })
}

// @Tags CaiGouJiHuaBiao
// @Summary 分页获取CaiGouJiHuaBiao列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取CaiGouJiHuaBiao列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /caiGouJiHuaBiao/getCaiGouJiHuaBiaoList [get]
export const getCaiGouJiHuaBiaoList = (params) => {
  return service({
    url: '/caiGouJiHuaBiao/getCaiGouJiHuaBiaoList',
    method: 'get',
    params
  })
}
