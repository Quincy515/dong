import service from '@/utils/request'

// @Tags XiaoXiTongZhi
// @Summary 创建XiaoXiTongZhi
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.XiaoXiTongZhi true "创建XiaoXiTongZhi"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /xiaoXiTongZhi/createXiaoXiTongZhi [post]
export const createXiaoXiTongZhi = (data) => {
  return service({
    url: '/xiaoXiTongZhi/createXiaoXiTongZhi',
    method: 'post',
    data
  })
}

// @Tags XiaoXiTongZhi
// @Summary 删除XiaoXiTongZhi
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.XiaoXiTongZhi true "删除XiaoXiTongZhi"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /xiaoXiTongZhi/deleteXiaoXiTongZhi [delete]
export const deleteXiaoXiTongZhi = (data) => {
  return service({
    url: '/xiaoXiTongZhi/deleteXiaoXiTongZhi',
    method: 'delete',
    data
  })
}

// @Tags XiaoXiTongZhi
// @Summary 删除XiaoXiTongZhi
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除XiaoXiTongZhi"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /xiaoXiTongZhi/deleteXiaoXiTongZhi [delete]
export const deleteXiaoXiTongZhiByIds = (data) => {
  return service({
    url: '/xiaoXiTongZhi/deleteXiaoXiTongZhiByIds',
    method: 'delete',
    data
  })
}

// @Tags XiaoXiTongZhi
// @Summary 更新XiaoXiTongZhi
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.XiaoXiTongZhi true "更新XiaoXiTongZhi"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /xiaoXiTongZhi/updateXiaoXiTongZhi [put]
export const updateXiaoXiTongZhi = (data) => {
  return service({
    url: '/xiaoXiTongZhi/updateXiaoXiTongZhi',
    method: 'put',
    data
  })
}

// @Tags XiaoXiTongZhi
// @Summary 用id查询XiaoXiTongZhi
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.XiaoXiTongZhi true "用id查询XiaoXiTongZhi"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /xiaoXiTongZhi/findXiaoXiTongZhi [get]
export const findXiaoXiTongZhi = (params) => {
  return service({
    url: '/xiaoXiTongZhi/findXiaoXiTongZhi',
    method: 'get',
    params
  })
}

// @Tags XiaoXiTongZhi
// @Summary 分页获取XiaoXiTongZhi列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取XiaoXiTongZhi列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /xiaoXiTongZhi/getXiaoXiTongZhiList [get]
export const getXiaoXiTongZhiList = (params) => {
  return service({
    url: '/xiaoXiTongZhi/getXiaoXiTongZhiList',
    method: 'get',
    params
  })
}
