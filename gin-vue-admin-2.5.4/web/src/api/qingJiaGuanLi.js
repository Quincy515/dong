import service from '@/utils/request'

// @Tags QingJiaGuanLi
// @Summary 创建QingJiaGuanLi
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.QingJiaGuanLi true "创建QingJiaGuanLi"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /qingJiaGuanLi/createQingJiaGuanLi [post]
export const createQingJiaGuanLi = (data) => {
  return service({
    url: '/qingJiaGuanLi/createQingJiaGuanLi',
    method: 'post',
    data
  })
}

// @Tags QingJiaGuanLi
// @Summary 删除QingJiaGuanLi
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.QingJiaGuanLi true "删除QingJiaGuanLi"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /qingJiaGuanLi/deleteQingJiaGuanLi [delete]
export const deleteQingJiaGuanLi = (data) => {
  return service({
    url: '/qingJiaGuanLi/deleteQingJiaGuanLi',
    method: 'delete',
    data
  })
}

// @Tags QingJiaGuanLi
// @Summary 删除QingJiaGuanLi
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除QingJiaGuanLi"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /qingJiaGuanLi/deleteQingJiaGuanLi [delete]
export const deleteQingJiaGuanLiByIds = (data) => {
  return service({
    url: '/qingJiaGuanLi/deleteQingJiaGuanLiByIds',
    method: 'delete',
    data
  })
}

// @Tags QingJiaGuanLi
// @Summary 更新QingJiaGuanLi
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.QingJiaGuanLi true "更新QingJiaGuanLi"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /qingJiaGuanLi/updateQingJiaGuanLi [put]
export const updateQingJiaGuanLi = (data) => {
  return service({
    url: '/qingJiaGuanLi/updateQingJiaGuanLi',
    method: 'put',
    data
  })
}

// @Tags QingJiaGuanLi
// @Summary 用id查询QingJiaGuanLi
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.QingJiaGuanLi true "用id查询QingJiaGuanLi"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /qingJiaGuanLi/findQingJiaGuanLi [get]
export const findQingJiaGuanLi = (params) => {
  return service({
    url: '/qingJiaGuanLi/findQingJiaGuanLi',
    method: 'get',
    params
  })
}

// @Tags QingJiaGuanLi
// @Summary 分页获取QingJiaGuanLi列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取QingJiaGuanLi列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /qingJiaGuanLi/getQingJiaGuanLiList [get]
export const getQingJiaGuanLiList = (params) => {
  return service({
    url: '/qingJiaGuanLi/getQingJiaGuanLiList',
    method: 'get',
    params
  })
}
