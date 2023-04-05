import service from '@/utils/request'

// @Tags FenBaoShangBiao
// @Summary 创建FenBaoShangBiao
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.FenBaoShangBiao true "创建FenBaoShangBiao"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /fenBaoShangBiao/createFenBaoShangBiao [post]
export const createFenBaoShangBiao = (data) => {
  return service({
    url: '/fenBaoShangBiao/createFenBaoShangBiao',
    method: 'post',
    data
  })
}

// @Tags FenBaoShangBiao
// @Summary 删除FenBaoShangBiao
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.FenBaoShangBiao true "删除FenBaoShangBiao"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /fenBaoShangBiao/deleteFenBaoShangBiao [delete]
export const deleteFenBaoShangBiao = (data) => {
  return service({
    url: '/fenBaoShangBiao/deleteFenBaoShangBiao',
    method: 'delete',
    data
  })
}

// @Tags FenBaoShangBiao
// @Summary 删除FenBaoShangBiao
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除FenBaoShangBiao"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /fenBaoShangBiao/deleteFenBaoShangBiao [delete]
export const deleteFenBaoShangBiaoByIds = (data) => {
  return service({
    url: '/fenBaoShangBiao/deleteFenBaoShangBiaoByIds',
    method: 'delete',
    data
  })
}

// @Tags FenBaoShangBiao
// @Summary 更新FenBaoShangBiao
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.FenBaoShangBiao true "更新FenBaoShangBiao"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /fenBaoShangBiao/updateFenBaoShangBiao [put]
export const updateFenBaoShangBiao = (data) => {
  return service({
    url: '/fenBaoShangBiao/updateFenBaoShangBiao',
    method: 'put',
    data
  })
}

// @Tags FenBaoShangBiao
// @Summary 用id查询FenBaoShangBiao
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.FenBaoShangBiao true "用id查询FenBaoShangBiao"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /fenBaoShangBiao/findFenBaoShangBiao [get]
export const findFenBaoShangBiao = (params) => {
  return service({
    url: '/fenBaoShangBiao/findFenBaoShangBiao',
    method: 'get',
    params
  })
}

// @Tags FenBaoShangBiao
// @Summary 分页获取FenBaoShangBiao列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取FenBaoShangBiao列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /fenBaoShangBiao/getFenBaoShangBiaoList [get]
export const getFenBaoShangBiaoList = (params) => {
  return service({
    url: '/fenBaoShangBiao/getFenBaoShangBiaoList',
    method: 'get',
    params
  })
}
