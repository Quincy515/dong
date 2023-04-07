import service from '@/utils/request'

// @Tags CaiGouJiHuaMingXiBiao
// @Summary 创建CaiGouJiHuaMingXiBiao
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.CaiGouJiHuaMingXiBiao true "创建CaiGouJiHuaMingXiBiao"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /caiGouJiHuaMingXiBiao/createCaiGouJiHuaMingXiBiao [post]
export const createCaiGouJiHuaMingXiBiao = (data) => {
  return service({
    url: '/caiGouJiHuaMingXiBiao/createCaiGouJiHuaMingXiBiao',
    method: 'post',
    data
  })
}

// @Tags CaiGouJiHuaMingXiBiao
// @Summary 删除CaiGouJiHuaMingXiBiao
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.CaiGouJiHuaMingXiBiao true "删除CaiGouJiHuaMingXiBiao"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /caiGouJiHuaMingXiBiao/deleteCaiGouJiHuaMingXiBiao [delete]
export const deleteCaiGouJiHuaMingXiBiao = (data) => {
  return service({
    url: '/caiGouJiHuaMingXiBiao/deleteCaiGouJiHuaMingXiBiao',
    method: 'delete',
    data
  })
}

// @Tags CaiGouJiHuaMingXiBiao
// @Summary 删除CaiGouJiHuaMingXiBiao
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除CaiGouJiHuaMingXiBiao"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /caiGouJiHuaMingXiBiao/deleteCaiGouJiHuaMingXiBiao [delete]
export const deleteCaiGouJiHuaMingXiBiaoByIds = (data) => {
  return service({
    url: '/caiGouJiHuaMingXiBiao/deleteCaiGouJiHuaMingXiBiaoByIds',
    method: 'delete',
    data
  })
}

// @Tags CaiGouJiHuaMingXiBiao
// @Summary 更新CaiGouJiHuaMingXiBiao
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.CaiGouJiHuaMingXiBiao true "更新CaiGouJiHuaMingXiBiao"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /caiGouJiHuaMingXiBiao/updateCaiGouJiHuaMingXiBiao [put]
export const updateCaiGouJiHuaMingXiBiao = (data) => {
  return service({
    url: '/caiGouJiHuaMingXiBiao/updateCaiGouJiHuaMingXiBiao',
    method: 'put',
    data
  })
}

// @Tags CaiGouJiHuaMingXiBiao
// @Summary 用id查询CaiGouJiHuaMingXiBiao
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.CaiGouJiHuaMingXiBiao true "用id查询CaiGouJiHuaMingXiBiao"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /caiGouJiHuaMingXiBiao/findCaiGouJiHuaMingXiBiao [get]
export const findCaiGouJiHuaMingXiBiao = (params) => {
  return service({
    url: '/caiGouJiHuaMingXiBiao/findCaiGouJiHuaMingXiBiao',
    method: 'get',
    params
  })
}

// @Tags CaiGouJiHuaMingXiBiao
// @Summary 分页获取CaiGouJiHuaMingXiBiao列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取CaiGouJiHuaMingXiBiao列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /caiGouJiHuaMingXiBiao/getCaiGouJiHuaMingXiBiaoList [get]
export const getCaiGouJiHuaMingXiBiaoList = (params) => {
  return service({
    url: '/caiGouJiHuaMingXiBiao/getCaiGouJiHuaMingXiBiaoList',
    method: 'get',
    params
  })
}
