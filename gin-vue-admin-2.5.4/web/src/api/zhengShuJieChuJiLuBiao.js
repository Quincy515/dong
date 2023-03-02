import service from '@/utils/request'

// @Tags ZhengShuJieChuJiLuBiao
// @Summary 创建ZhengShuJieChuJiLuBiao
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.ZhengShuJieChuJiLuBiao true "创建ZhengShuJieChuJiLuBiao"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /zhengShuJieChuJiLuBiao/createZhengShuJieChuJiLuBiao [post]
export const createZhengShuJieChuJiLuBiao = (data) => {
  return service({
    url: '/zhengShuJieChuJiLuBiao/createZhengShuJieChuJiLuBiao',
    method: 'post',
    data
  })
}

// @Tags ZhengShuJieChuJiLuBiao
// @Summary 删除ZhengShuJieChuJiLuBiao
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.ZhengShuJieChuJiLuBiao true "删除ZhengShuJieChuJiLuBiao"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /zhengShuJieChuJiLuBiao/deleteZhengShuJieChuJiLuBiao [delete]
export const deleteZhengShuJieChuJiLuBiao = (data) => {
  return service({
    url: '/zhengShuJieChuJiLuBiao/deleteZhengShuJieChuJiLuBiao',
    method: 'delete',
    data
  })
}

// @Tags ZhengShuJieChuJiLuBiao
// @Summary 删除ZhengShuJieChuJiLuBiao
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除ZhengShuJieChuJiLuBiao"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /zhengShuJieChuJiLuBiao/deleteZhengShuJieChuJiLuBiao [delete]
export const deleteZhengShuJieChuJiLuBiaoByIds = (data) => {
  return service({
    url: '/zhengShuJieChuJiLuBiao/deleteZhengShuJieChuJiLuBiaoByIds',
    method: 'delete',
    data
  })
}

// @Tags ZhengShuJieChuJiLuBiao
// @Summary 更新ZhengShuJieChuJiLuBiao
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.ZhengShuJieChuJiLuBiao true "更新ZhengShuJieChuJiLuBiao"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /zhengShuJieChuJiLuBiao/updateZhengShuJieChuJiLuBiao [put]
export const updateZhengShuJieChuJiLuBiao = (data) => {
  return service({
    url: '/zhengShuJieChuJiLuBiao/updateZhengShuJieChuJiLuBiao',
    method: 'put',
    data
  })
}

// @Tags ZhengShuJieChuJiLuBiao
// @Summary 用id查询ZhengShuJieChuJiLuBiao
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.ZhengShuJieChuJiLuBiao true "用id查询ZhengShuJieChuJiLuBiao"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /zhengShuJieChuJiLuBiao/findZhengShuJieChuJiLuBiao [get]
export const findZhengShuJieChuJiLuBiao = (params) => {
  return service({
    url: '/zhengShuJieChuJiLuBiao/findZhengShuJieChuJiLuBiao',
    method: 'get',
    params
  })
}

// @Tags ZhengShuJieChuJiLuBiao
// @Summary 分页获取ZhengShuJieChuJiLuBiao列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取ZhengShuJieChuJiLuBiao列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /zhengShuJieChuJiLuBiao/getZhengShuJieChuJiLuBiaoList [get]
export const getZhengShuJieChuJiLuBiaoList = (params) => {
  return service({
    url: '/zhengShuJieChuJiLuBiao/getZhengShuJieChuJiLuBiaoList',
    method: 'get',
    params
  })
}
