import service from '@/utils/request'

// @Tags ZhengShuBiao
// @Summary 创建ZhengShuBiao
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.ZhengShuBiao true "创建ZhengShuBiao"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /zhengShuBiao/createZhengShuBiao [post]
export const createZhengShuBiao = (data) => {
  return service({
    url: '/zhengShuBiao/createZhengShuBiao',
    method: 'post',
    data
  })
}

// @Tags ZhengShuBiao
// @Summary 删除ZhengShuBiao
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.ZhengShuBiao true "删除ZhengShuBiao"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /zhengShuBiao/deleteZhengShuBiao [delete]
export const deleteZhengShuBiao = (data) => {
  return service({
    url: '/zhengShuBiao/deleteZhengShuBiao',
    method: 'delete',
    data
  })
}

// @Tags ZhengShuBiao
// @Summary 删除ZhengShuBiao
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除ZhengShuBiao"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /zhengShuBiao/deleteZhengShuBiao [delete]
export const deleteZhengShuBiaoByIds = (data) => {
  return service({
    url: '/zhengShuBiao/deleteZhengShuBiaoByIds',
    method: 'delete',
    data
  })
}

// @Tags ZhengShuBiao
// @Summary 更新ZhengShuBiao
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.ZhengShuBiao true "更新ZhengShuBiao"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /zhengShuBiao/updateZhengShuBiao [put]
export const updateZhengShuBiao = (data) => {
  return service({
    url: '/zhengShuBiao/updateZhengShuBiao',
    method: 'put',
    data
  })
}

// @Tags ZhengShuBiao
// @Summary 用id查询ZhengShuBiao
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.ZhengShuBiao true "用id查询ZhengShuBiao"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /zhengShuBiao/findZhengShuBiao [get]
export const findZhengShuBiao = (params) => {
  return service({
    url: '/zhengShuBiao/findZhengShuBiao',
    method: 'get',
    params
  })
}

// @Tags ZhengShuBiao
// @Summary 分页获取ZhengShuBiao列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取ZhengShuBiao列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /zhengShuBiao/getZhengShuBiaoList [get]
export const getZhengShuBiaoList = (params) => {
  return service({
    url: '/zhengShuBiao/getZhengShuBiaoList',
    method: 'get',
    params
  })
}
