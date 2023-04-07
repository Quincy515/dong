import service from '@/utils/request'

// @Tags XiangMuZhuCaiBiao
// @Summary 创建XiangMuZhuCaiBiao
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.XiangMuZhuCaiBiao true "创建XiangMuZhuCaiBiao"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /xiangMuZhuCaiBiao/createXiangMuZhuCaiBiao [post]
export const createXiangMuZhuCaiBiao = (data) => {
  return service({
    url: '/xiangMuZhuCaiBiao/createXiangMuZhuCaiBiao',
    method: 'post',
    data
  })
}

// @Tags XiangMuZhuCaiBiao
// @Summary 删除XiangMuZhuCaiBiao
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.XiangMuZhuCaiBiao true "删除XiangMuZhuCaiBiao"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /xiangMuZhuCaiBiao/deleteXiangMuZhuCaiBiao [delete]
export const deleteXiangMuZhuCaiBiao = (data) => {
  return service({
    url: '/xiangMuZhuCaiBiao/deleteXiangMuZhuCaiBiao',
    method: 'delete',
    data
  })
}

// @Tags XiangMuZhuCaiBiao
// @Summary 删除XiangMuZhuCaiBiao
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除XiangMuZhuCaiBiao"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /xiangMuZhuCaiBiao/deleteXiangMuZhuCaiBiao [delete]
export const deleteXiangMuZhuCaiBiaoByIds = (data) => {
  return service({
    url: '/xiangMuZhuCaiBiao/deleteXiangMuZhuCaiBiaoByIds',
    method: 'delete',
    data
  })
}

// @Tags XiangMuZhuCaiBiao
// @Summary 更新XiangMuZhuCaiBiao
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.XiangMuZhuCaiBiao true "更新XiangMuZhuCaiBiao"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /xiangMuZhuCaiBiao/updateXiangMuZhuCaiBiao [put]
export const updateXiangMuZhuCaiBiao = (data) => {
  return service({
    url: '/xiangMuZhuCaiBiao/updateXiangMuZhuCaiBiao',
    method: 'put',
    data
  })
}

// @Tags XiangMuZhuCaiBiao
// @Summary 用id查询XiangMuZhuCaiBiao
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.XiangMuZhuCaiBiao true "用id查询XiangMuZhuCaiBiao"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /xiangMuZhuCaiBiao/findXiangMuZhuCaiBiao [get]
export const findXiangMuZhuCaiBiao = (params) => {
  return service({
    url: '/xiangMuZhuCaiBiao/findXiangMuZhuCaiBiao',
    method: 'get',
    params
  })
}

// @Tags XiangMuZhuCaiBiao
// @Summary 分页获取XiangMuZhuCaiBiao列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取XiangMuZhuCaiBiao列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /xiangMuZhuCaiBiao/getXiangMuZhuCaiBiaoList [get]
export const getXiangMuZhuCaiBiaoList = (params) => {
  return service({
    url: '/xiangMuZhuCaiBiao/getXiangMuZhuCaiBiaoList',
    method: 'get',
    params
  })
}
