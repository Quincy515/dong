import service from '@/utils/request'

// @Tags GongZhangDengJi
// @Summary 创建GongZhangDengJi
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.GongZhangDengJi true "创建GongZhangDengJi"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /gongZhangDengJi/createGongZhangDengJi [post]
export const createGongZhangDengJi = (data) => {
  return service({
    url: '/gongZhangDengJi/createGongZhangDengJi',
    method: 'post',
    data
  })
}

// @Tags GongZhangDengJi
// @Summary 删除GongZhangDengJi
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.GongZhangDengJi true "删除GongZhangDengJi"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /gongZhangDengJi/deleteGongZhangDengJi [delete]
export const deleteGongZhangDengJi = (data) => {
  return service({
    url: '/gongZhangDengJi/deleteGongZhangDengJi',
    method: 'delete',
    data
  })
}

// @Tags GongZhangDengJi
// @Summary 删除GongZhangDengJi
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除GongZhangDengJi"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /gongZhangDengJi/deleteGongZhangDengJi [delete]
export const deleteGongZhangDengJiByIds = (data) => {
  return service({
    url: '/gongZhangDengJi/deleteGongZhangDengJiByIds',
    method: 'delete',
    data
  })
}

// @Tags GongZhangDengJi
// @Summary 更新GongZhangDengJi
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.GongZhangDengJi true "更新GongZhangDengJi"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /gongZhangDengJi/updateGongZhangDengJi [put]
export const updateGongZhangDengJi = (data) => {
  return service({
    url: '/gongZhangDengJi/updateGongZhangDengJi',
    method: 'put',
    data
  })
}

// @Tags GongZhangDengJi
// @Summary 用id查询GongZhangDengJi
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.GongZhangDengJi true "用id查询GongZhangDengJi"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /gongZhangDengJi/findGongZhangDengJi [get]
export const findGongZhangDengJi = (params) => {
  return service({
    url: '/gongZhangDengJi/findGongZhangDengJi',
    method: 'get',
    params
  })
}

// @Tags GongZhangDengJi
// @Summary 分页获取GongZhangDengJi列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取GongZhangDengJi列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /gongZhangDengJi/getGongZhangDengJiList [get]
export const getGongZhangDengJiList = (params) => {
  return service({
    url: '/gongZhangDengJi/getGongZhangDengJiList',
    method: 'get',
    params
  })
}
