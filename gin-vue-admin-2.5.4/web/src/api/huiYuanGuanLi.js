import service from '@/utils/request'

// @Tags HuiYuanGuanLi
// @Summary 创建HuiYuanGuanLi
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.HuiYuanGuanLi true "创建HuiYuanGuanLi"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /huiYuanGuanLi/createHuiYuanGuanLi [post]
export const createHuiYuanGuanLi = (data) => {
  return service({
    url: '/huiYuanGuanLi/createHuiYuanGuanLi',
    method: 'post',
    data
  })
}

// @Tags HuiYuanGuanLi
// @Summary 删除HuiYuanGuanLi
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.HuiYuanGuanLi true "删除HuiYuanGuanLi"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /huiYuanGuanLi/deleteHuiYuanGuanLi [delete]
export const deleteHuiYuanGuanLi = (data) => {
  return service({
    url: '/huiYuanGuanLi/deleteHuiYuanGuanLi',
    method: 'delete',
    data
  })
}

// @Tags HuiYuanGuanLi
// @Summary 删除HuiYuanGuanLi
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除HuiYuanGuanLi"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /huiYuanGuanLi/deleteHuiYuanGuanLi [delete]
export const deleteHuiYuanGuanLiByIds = (data) => {
  return service({
    url: '/huiYuanGuanLi/deleteHuiYuanGuanLiByIds',
    method: 'delete',
    data
  })
}

// @Tags HuiYuanGuanLi
// @Summary 更新HuiYuanGuanLi
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.HuiYuanGuanLi true "更新HuiYuanGuanLi"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /huiYuanGuanLi/updateHuiYuanGuanLi [put]
export const updateHuiYuanGuanLi = (data) => {
  return service({
    url: '/huiYuanGuanLi/updateHuiYuanGuanLi',
    method: 'put',
    data
  })
}

// @Tags HuiYuanGuanLi
// @Summary 用id查询HuiYuanGuanLi
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.HuiYuanGuanLi true "用id查询HuiYuanGuanLi"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /huiYuanGuanLi/findHuiYuanGuanLi [get]
export const findHuiYuanGuanLi = (params) => {
  return service({
    url: '/huiYuanGuanLi/findHuiYuanGuanLi',
    method: 'get',
    params
  })
}

// @Tags HuiYuanGuanLi
// @Summary 分页获取HuiYuanGuanLi列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取HuiYuanGuanLi列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /huiYuanGuanLi/getHuiYuanGuanLiList [get]
export const getHuiYuanGuanLiList = (params) => {
  return service({
    url: '/huiYuanGuanLi/getHuiYuanGuanLiList',
    method: 'get',
    params
  })
}
