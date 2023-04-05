import service from '@/utils/request'

// @Tags XiangMuShiGongBiao
// @Summary 创建XiangMuShiGongBiao
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.XiangMuShiGongBiao true "创建XiangMuShiGongBiao"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /xiangMuShiGongBiao/createXiangMuShiGongBiao [post]
export const createXiangMuShiGongBiao = (data) => {
  return service({
    url: '/xiangMuShiGongBiao/createXiangMuShiGongBiao',
    method: 'post',
    data
  })
}

// @Tags XiangMuShiGongBiao
// @Summary 删除XiangMuShiGongBiao
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.XiangMuShiGongBiao true "删除XiangMuShiGongBiao"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /xiangMuShiGongBiao/deleteXiangMuShiGongBiao [delete]
export const deleteXiangMuShiGongBiao = (data) => {
  return service({
    url: '/xiangMuShiGongBiao/deleteXiangMuShiGongBiao',
    method: 'delete',
    data
  })
}

// @Tags XiangMuShiGongBiao
// @Summary 删除XiangMuShiGongBiao
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除XiangMuShiGongBiao"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /xiangMuShiGongBiao/deleteXiangMuShiGongBiao [delete]
export const deleteXiangMuShiGongBiaoByIds = (data) => {
  return service({
    url: '/xiangMuShiGongBiao/deleteXiangMuShiGongBiaoByIds',
    method: 'delete',
    data
  })
}

// @Tags XiangMuShiGongBiao
// @Summary 更新XiangMuShiGongBiao
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.XiangMuShiGongBiao true "更新XiangMuShiGongBiao"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /xiangMuShiGongBiao/updateXiangMuShiGongBiao [put]
export const updateXiangMuShiGongBiao = (data) => {
  return service({
    url: '/xiangMuShiGongBiao/updateXiangMuShiGongBiao',
    method: 'put',
    data
  })
}

// @Tags XiangMuShiGongBiao
// @Summary 用id查询XiangMuShiGongBiao
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.XiangMuShiGongBiao true "用id查询XiangMuShiGongBiao"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /xiangMuShiGongBiao/findXiangMuShiGongBiao [get]
export const findXiangMuShiGongBiao = (params) => {
  return service({
    url: '/xiangMuShiGongBiao/findXiangMuShiGongBiao',
    method: 'get',
    params
  })
}

// @Tags XiangMuShiGongBiao
// @Summary 分页获取XiangMuShiGongBiao列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取XiangMuShiGongBiao列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /xiangMuShiGongBiao/getXiangMuShiGongBiaoList [get]
export const getXiangMuShiGongBiaoList = (params) => {
  return service({
    url: '/xiangMuShiGongBiao/getXiangMuShiGongBiaoList',
    method: 'get',
    params
  })
}
