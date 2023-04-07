import service from '@/utils/request'

// @Tags NeiBuBaoXiaoBiao
// @Summary 创建NeiBuBaoXiaoBiao
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.NeiBuBaoXiaoBiao true "创建NeiBuBaoXiaoBiao"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /neiBuBaoXiaoBiao/createNeiBuBaoXiaoBiao [post]
export const createNeiBuBaoXiaoBiao = (data) => {
  return service({
    url: '/neiBuBaoXiaoBiao/createNeiBuBaoXiaoBiao',
    method: 'post',
    data
  })
}

// @Tags NeiBuBaoXiaoBiao
// @Summary 删除NeiBuBaoXiaoBiao
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.NeiBuBaoXiaoBiao true "删除NeiBuBaoXiaoBiao"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /neiBuBaoXiaoBiao/deleteNeiBuBaoXiaoBiao [delete]
export const deleteNeiBuBaoXiaoBiao = (data) => {
  return service({
    url: '/neiBuBaoXiaoBiao/deleteNeiBuBaoXiaoBiao',
    method: 'delete',
    data
  })
}

// @Tags NeiBuBaoXiaoBiao
// @Summary 删除NeiBuBaoXiaoBiao
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除NeiBuBaoXiaoBiao"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /neiBuBaoXiaoBiao/deleteNeiBuBaoXiaoBiao [delete]
export const deleteNeiBuBaoXiaoBiaoByIds = (data) => {
  return service({
    url: '/neiBuBaoXiaoBiao/deleteNeiBuBaoXiaoBiaoByIds',
    method: 'delete',
    data
  })
}

// @Tags NeiBuBaoXiaoBiao
// @Summary 更新NeiBuBaoXiaoBiao
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.NeiBuBaoXiaoBiao true "更新NeiBuBaoXiaoBiao"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /neiBuBaoXiaoBiao/updateNeiBuBaoXiaoBiao [put]
export const updateNeiBuBaoXiaoBiao = (data) => {
  return service({
    url: '/neiBuBaoXiaoBiao/updateNeiBuBaoXiaoBiao',
    method: 'put',
    data
  })
}

// @Tags NeiBuBaoXiaoBiao
// @Summary 用id查询NeiBuBaoXiaoBiao
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.NeiBuBaoXiaoBiao true "用id查询NeiBuBaoXiaoBiao"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /neiBuBaoXiaoBiao/findNeiBuBaoXiaoBiao [get]
export const findNeiBuBaoXiaoBiao = (params) => {
  return service({
    url: '/neiBuBaoXiaoBiao/findNeiBuBaoXiaoBiao',
    method: 'get',
    params
  })
}

// @Tags NeiBuBaoXiaoBiao
// @Summary 分页获取NeiBuBaoXiaoBiao列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取NeiBuBaoXiaoBiao列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /neiBuBaoXiaoBiao/getNeiBuBaoXiaoBiaoList [get]
export const getNeiBuBaoXiaoBiaoList = (params) => {
  return service({
    url: '/neiBuBaoXiaoBiao/getNeiBuBaoXiaoBiaoList',
    method: 'get',
    params
  })
}
