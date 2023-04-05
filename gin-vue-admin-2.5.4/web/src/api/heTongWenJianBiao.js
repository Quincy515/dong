import service from '@/utils/request'

// @Tags HeTongWenJianBiao
// @Summary 创建HeTongWenJianBiao
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.HeTongWenJianBiao true "创建HeTongWenJianBiao"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /heTongWenJianBiao/createHeTongWenJianBiao [post]
export const createHeTongWenJianBiao = (data) => {
  return service({
    url: '/heTongWenJianBiao/createHeTongWenJianBiao',
    method: 'post',
    data
  })
}

// @Tags HeTongWenJianBiao
// @Summary 删除HeTongWenJianBiao
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.HeTongWenJianBiao true "删除HeTongWenJianBiao"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /heTongWenJianBiao/deleteHeTongWenJianBiao [delete]
export const deleteHeTongWenJianBiao = (data) => {
  return service({
    url: '/heTongWenJianBiao/deleteHeTongWenJianBiao',
    method: 'delete',
    data
  })
}

// @Tags HeTongWenJianBiao
// @Summary 删除HeTongWenJianBiao
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除HeTongWenJianBiao"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /heTongWenJianBiao/deleteHeTongWenJianBiao [delete]
export const deleteHeTongWenJianBiaoByIds = (data) => {
  return service({
    url: '/heTongWenJianBiao/deleteHeTongWenJianBiaoByIds',
    method: 'delete',
    data
  })
}

// @Tags HeTongWenJianBiao
// @Summary 更新HeTongWenJianBiao
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.HeTongWenJianBiao true "更新HeTongWenJianBiao"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /heTongWenJianBiao/updateHeTongWenJianBiao [put]
export const updateHeTongWenJianBiao = (data) => {
  return service({
    url: '/heTongWenJianBiao/updateHeTongWenJianBiao',
    method: 'put',
    data
  })
}

// @Tags HeTongWenJianBiao
// @Summary 用id查询HeTongWenJianBiao
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.HeTongWenJianBiao true "用id查询HeTongWenJianBiao"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /heTongWenJianBiao/findHeTongWenJianBiao [get]
export const findHeTongWenJianBiao = (params) => {
  return service({
    url: '/heTongWenJianBiao/findHeTongWenJianBiao',
    method: 'get',
    params
  })
}

// @Tags HeTongWenJianBiao
// @Summary 分页获取HeTongWenJianBiao列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取HeTongWenJianBiao列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /heTongWenJianBiao/getHeTongWenJianBiaoList [get]
export const getHeTongWenJianBiaoList = (params) => {
  return service({
    url: '/heTongWenJianBiao/getHeTongWenJianBiaoList',
    method: 'get',
    params
  })
}
