import service from '@/utils/request'

// @Tags SheBeiMingXiBiao
// @Summary 创建SheBeiMingXiBiao
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.SheBeiMingXiBiao true "创建SheBeiMingXiBiao"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /sheBeiMingXiBiao/createSheBeiMingXiBiao [post]
export const createSheBeiMingXiBiao = (data) => {
  return service({
    url: '/sheBeiMingXiBiao/createSheBeiMingXiBiao',
    method: 'post',
    data
  })
}

// @Tags SheBeiMingXiBiao
// @Summary 删除SheBeiMingXiBiao
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.SheBeiMingXiBiao true "删除SheBeiMingXiBiao"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /sheBeiMingXiBiao/deleteSheBeiMingXiBiao [delete]
export const deleteSheBeiMingXiBiao = (data) => {
  return service({
    url: '/sheBeiMingXiBiao/deleteSheBeiMingXiBiao',
    method: 'delete',
    data
  })
}

// @Tags SheBeiMingXiBiao
// @Summary 删除SheBeiMingXiBiao
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除SheBeiMingXiBiao"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /sheBeiMingXiBiao/deleteSheBeiMingXiBiao [delete]
export const deleteSheBeiMingXiBiaoByIds = (data) => {
  return service({
    url: '/sheBeiMingXiBiao/deleteSheBeiMingXiBiaoByIds',
    method: 'delete',
    data
  })
}

// @Tags SheBeiMingXiBiao
// @Summary 更新SheBeiMingXiBiao
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.SheBeiMingXiBiao true "更新SheBeiMingXiBiao"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /sheBeiMingXiBiao/updateSheBeiMingXiBiao [put]
export const updateSheBeiMingXiBiao = (data) => {
  return service({
    url: '/sheBeiMingXiBiao/updateSheBeiMingXiBiao',
    method: 'put',
    data
  })
}

// @Tags SheBeiMingXiBiao
// @Summary 用id查询SheBeiMingXiBiao
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.SheBeiMingXiBiao true "用id查询SheBeiMingXiBiao"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /sheBeiMingXiBiao/findSheBeiMingXiBiao [get]
export const findSheBeiMingXiBiao = (params) => {
  return service({
    url: '/sheBeiMingXiBiao/findSheBeiMingXiBiao',
    method: 'get',
    params
  })
}

// @Tags SheBeiMingXiBiao
// @Summary 分页获取SheBeiMingXiBiao列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取SheBeiMingXiBiao列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /sheBeiMingXiBiao/getSheBeiMingXiBiaoList [get]
export const getSheBeiMingXiBiaoList = (params) => {
  return service({
    url: '/sheBeiMingXiBiao/getSheBeiMingXiBiaoList',
    method: 'get',
    params
  })
}
