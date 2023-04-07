import service from '@/utils/request'

// @Tags SheBeiZuLinBiao
// @Summary 创建SheBeiZuLinBiao
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.SheBeiZuLinBiao true "创建SheBeiZuLinBiao"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /sheBeiZuLinBiao/createSheBeiZuLinBiao [post]
export const createSheBeiZuLinBiao = (data) => {
  return service({
    url: '/sheBeiZuLinBiao/createSheBeiZuLinBiao',
    method: 'post',
    data
  })
}

// @Tags SheBeiZuLinBiao
// @Summary 删除SheBeiZuLinBiao
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.SheBeiZuLinBiao true "删除SheBeiZuLinBiao"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /sheBeiZuLinBiao/deleteSheBeiZuLinBiao [delete]
export const deleteSheBeiZuLinBiao = (data) => {
  return service({
    url: '/sheBeiZuLinBiao/deleteSheBeiZuLinBiao',
    method: 'delete',
    data
  })
}

// @Tags SheBeiZuLinBiao
// @Summary 删除SheBeiZuLinBiao
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除SheBeiZuLinBiao"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /sheBeiZuLinBiao/deleteSheBeiZuLinBiao [delete]
export const deleteSheBeiZuLinBiaoByIds = (data) => {
  return service({
    url: '/sheBeiZuLinBiao/deleteSheBeiZuLinBiaoByIds',
    method: 'delete',
    data
  })
}

// @Tags SheBeiZuLinBiao
// @Summary 更新SheBeiZuLinBiao
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.SheBeiZuLinBiao true "更新SheBeiZuLinBiao"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /sheBeiZuLinBiao/updateSheBeiZuLinBiao [put]
export const updateSheBeiZuLinBiao = (data) => {
  return service({
    url: '/sheBeiZuLinBiao/updateSheBeiZuLinBiao',
    method: 'put',
    data
  })
}

// @Tags SheBeiZuLinBiao
// @Summary 用id查询SheBeiZuLinBiao
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.SheBeiZuLinBiao true "用id查询SheBeiZuLinBiao"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /sheBeiZuLinBiao/findSheBeiZuLinBiao [get]
export const findSheBeiZuLinBiao = (params) => {
  return service({
    url: '/sheBeiZuLinBiao/findSheBeiZuLinBiao',
    method: 'get',
    params
  })
}

// @Tags SheBeiZuLinBiao
// @Summary 分页获取SheBeiZuLinBiao列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取SheBeiZuLinBiao列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /sheBeiZuLinBiao/getSheBeiZuLinBiaoList [get]
export const getSheBeiZuLinBiaoList = (params) => {
  return service({
    url: '/sheBeiZuLinBiao/getSheBeiZuLinBiaoList',
    method: 'get',
    params
  })
}
