import service from '@/utils/request'

// @Tags WaiBuFeiYongQingKuanBiao
// @Summary 创建WaiBuFeiYongQingKuanBiao
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.WaiBuFeiYongQingKuanBiao true "创建WaiBuFeiYongQingKuanBiao"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /waiBuFeiYongQingKuanBiao/createWaiBuFeiYongQingKuanBiao [post]
export const createWaiBuFeiYongQingKuanBiao = (data) => {
  return service({
    url: '/waiBuFeiYongQingKuanBiao/createWaiBuFeiYongQingKuanBiao',
    method: 'post',
    data
  })
}

// @Tags WaiBuFeiYongQingKuanBiao
// @Summary 删除WaiBuFeiYongQingKuanBiao
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.WaiBuFeiYongQingKuanBiao true "删除WaiBuFeiYongQingKuanBiao"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /waiBuFeiYongQingKuanBiao/deleteWaiBuFeiYongQingKuanBiao [delete]
export const deleteWaiBuFeiYongQingKuanBiao = (data) => {
  return service({
    url: '/waiBuFeiYongQingKuanBiao/deleteWaiBuFeiYongQingKuanBiao',
    method: 'delete',
    data
  })
}

// @Tags WaiBuFeiYongQingKuanBiao
// @Summary 删除WaiBuFeiYongQingKuanBiao
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除WaiBuFeiYongQingKuanBiao"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /waiBuFeiYongQingKuanBiao/deleteWaiBuFeiYongQingKuanBiao [delete]
export const deleteWaiBuFeiYongQingKuanBiaoByIds = (data) => {
  return service({
    url: '/waiBuFeiYongQingKuanBiao/deleteWaiBuFeiYongQingKuanBiaoByIds',
    method: 'delete',
    data
  })
}

// @Tags WaiBuFeiYongQingKuanBiao
// @Summary 更新WaiBuFeiYongQingKuanBiao
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.WaiBuFeiYongQingKuanBiao true "更新WaiBuFeiYongQingKuanBiao"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /waiBuFeiYongQingKuanBiao/updateWaiBuFeiYongQingKuanBiao [put]
export const updateWaiBuFeiYongQingKuanBiao = (data) => {
  return service({
    url: '/waiBuFeiYongQingKuanBiao/updateWaiBuFeiYongQingKuanBiao',
    method: 'put',
    data
  })
}

// @Tags WaiBuFeiYongQingKuanBiao
// @Summary 用id查询WaiBuFeiYongQingKuanBiao
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.WaiBuFeiYongQingKuanBiao true "用id查询WaiBuFeiYongQingKuanBiao"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /waiBuFeiYongQingKuanBiao/findWaiBuFeiYongQingKuanBiao [get]
export const findWaiBuFeiYongQingKuanBiao = (params) => {
  return service({
    url: '/waiBuFeiYongQingKuanBiao/findWaiBuFeiYongQingKuanBiao',
    method: 'get',
    params
  })
}

// @Tags WaiBuFeiYongQingKuanBiao
// @Summary 分页获取WaiBuFeiYongQingKuanBiao列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取WaiBuFeiYongQingKuanBiao列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /waiBuFeiYongQingKuanBiao/getWaiBuFeiYongQingKuanBiaoList [get]
export const getWaiBuFeiYongQingKuanBiaoList = (params) => {
  return service({
    url: '/waiBuFeiYongQingKuanBiao/getWaiBuFeiYongQingKuanBiaoList',
    method: 'get',
    params
  })
}
