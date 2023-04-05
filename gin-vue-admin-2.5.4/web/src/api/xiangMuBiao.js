import service from '@/utils/request'

// @Tags XiangMuBiao
// @Summary 创建XiangMuBiao
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.XiangMuBiao true "创建XiangMuBiao"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /xiangMuBiao/createXiangMuBiao [post]
export const createXiangMuBiao = (data) => {
  return service({
    url: '/xiangMuBiao/createXiangMuBiao',
    method: 'post',
    data
  })
}

// @Tags XiangMuBiao
// @Summary 删除XiangMuBiao
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.XiangMuBiao true "删除XiangMuBiao"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /xiangMuBiao/deleteXiangMuBiao [delete]
export const deleteXiangMuBiao = (data) => {
  return service({
    url: '/xiangMuBiao/deleteXiangMuBiao',
    method: 'delete',
    data
  })
}

// @Tags XiangMuBiao
// @Summary 删除XiangMuBiao
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除XiangMuBiao"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /xiangMuBiao/deleteXiangMuBiao [delete]
export const deleteXiangMuBiaoByIds = (data) => {
  return service({
    url: '/xiangMuBiao/deleteXiangMuBiaoByIds',
    method: 'delete',
    data
  })
}

// @Tags XiangMuBiao
// @Summary 更新XiangMuBiao
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.XiangMuBiao true "更新XiangMuBiao"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /xiangMuBiao/updateXiangMuBiao [put]
export const updateXiangMuBiao = (data) => {
  return service({
    url: '/xiangMuBiao/updateXiangMuBiao',
    method: 'put',
    data
  })
}

// @Tags XiangMuBiao
// @Summary 用id查询XiangMuBiao
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.XiangMuBiao true "用id查询XiangMuBiao"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /xiangMuBiao/findXiangMuBiao [get]
export const findXiangMuBiao = (params) => {
  return service({
    url: '/xiangMuBiao/findXiangMuBiao',
    method: 'get',
    params
  })
}

// @Tags XiangMuBiao
// @Summary 分页获取XiangMuBiao列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取XiangMuBiao列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /xiangMuBiao/getXiangMuBiaoList [get]
export const getXiangMuBiaoList = (params) => {
  return service({
    url: '/xiangMuBiao/getXiangMuBiaoList',
    method: 'get',
    params
  })
}
