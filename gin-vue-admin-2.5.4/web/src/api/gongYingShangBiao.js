import service from '@/utils/request'

// @Tags GongYingShangBiao
// @Summary 创建GongYingShangBiao
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.GongYingShangBiao true "创建GongYingShangBiao"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /gongYingShangBiao/createGongYingShangBiao [post]
export const createGongYingShangBiao = (data) => {
  return service({
    url: '/gongYingShangBiao/createGongYingShangBiao',
    method: 'post',
    data
  })
}

// @Tags GongYingShangBiao
// @Summary 删除GongYingShangBiao
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.GongYingShangBiao true "删除GongYingShangBiao"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /gongYingShangBiao/deleteGongYingShangBiao [delete]
export const deleteGongYingShangBiao = (data) => {
  return service({
    url: '/gongYingShangBiao/deleteGongYingShangBiao',
    method: 'delete',
    data
  })
}

// @Tags GongYingShangBiao
// @Summary 删除GongYingShangBiao
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除GongYingShangBiao"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /gongYingShangBiao/deleteGongYingShangBiao [delete]
export const deleteGongYingShangBiaoByIds = (data) => {
  return service({
    url: '/gongYingShangBiao/deleteGongYingShangBiaoByIds',
    method: 'delete',
    data
  })
}

// @Tags GongYingShangBiao
// @Summary 更新GongYingShangBiao
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.GongYingShangBiao true "更新GongYingShangBiao"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /gongYingShangBiao/updateGongYingShangBiao [put]
export const updateGongYingShangBiao = (data) => {
  return service({
    url: '/gongYingShangBiao/updateGongYingShangBiao',
    method: 'put',
    data
  })
}

// @Tags GongYingShangBiao
// @Summary 用id查询GongYingShangBiao
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.GongYingShangBiao true "用id查询GongYingShangBiao"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /gongYingShangBiao/findGongYingShangBiao [get]
export const findGongYingShangBiao = (params) => {
  return service({
    url: '/gongYingShangBiao/findGongYingShangBiao',
    method: 'get',
    params
  })
}

// @Tags GongYingShangBiao
// @Summary 分页获取GongYingShangBiao列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取GongYingShangBiao列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /gongYingShangBiao/getGongYingShangBiaoList [get]
export const getGongYingShangBiaoList = (params) => {
  return service({
    url: '/gongYingShangBiao/getGongYingShangBiaoList',
    method: 'get',
    params
  })
}
