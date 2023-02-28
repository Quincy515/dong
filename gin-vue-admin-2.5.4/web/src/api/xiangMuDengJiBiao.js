import service from '@/utils/request'

// @Tags XiangMuDengJiBiao
// @Summary 创建XiangMuDengJiBiao
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.XiangMuDengJiBiao true "创建XiangMuDengJiBiao"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /xiangMuDengJiBiao/createXiangMuDengJiBiao [post]
export const createXiangMuDengJiBiao = (data) => {
  return service({
    url: '/xiangMuDengJiBiao/createXiangMuDengJiBiao',
    method: 'post',
    data
  })
}

// @Tags XiangMuDengJiBiao
// @Summary 删除XiangMuDengJiBiao
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.XiangMuDengJiBiao true "删除XiangMuDengJiBiao"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /xiangMuDengJiBiao/deleteXiangMuDengJiBiao [delete]
export const deleteXiangMuDengJiBiao = (data) => {
  return service({
    url: '/xiangMuDengJiBiao/deleteXiangMuDengJiBiao',
    method: 'delete',
    data
  })
}

// @Tags XiangMuDengJiBiao
// @Summary 删除XiangMuDengJiBiao
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除XiangMuDengJiBiao"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /xiangMuDengJiBiao/deleteXiangMuDengJiBiao [delete]
export const deleteXiangMuDengJiBiaoByIds = (data) => {
  return service({
    url: '/xiangMuDengJiBiao/deleteXiangMuDengJiBiaoByIds',
    method: 'delete',
    data
  })
}

// @Tags XiangMuDengJiBiao
// @Summary 更新XiangMuDengJiBiao
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.XiangMuDengJiBiao true "更新XiangMuDengJiBiao"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /xiangMuDengJiBiao/updateXiangMuDengJiBiao [put]
export const updateXiangMuDengJiBiao = (data) => {
  return service({
    url: '/xiangMuDengJiBiao/updateXiangMuDengJiBiao',
    method: 'put',
    data
  })
}

// @Tags XiangMuDengJiBiao
// @Summary 用id查询XiangMuDengJiBiao
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.XiangMuDengJiBiao true "用id查询XiangMuDengJiBiao"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /xiangMuDengJiBiao/findXiangMuDengJiBiao [get]
export const findXiangMuDengJiBiao = (params) => {
  return service({
    url: '/xiangMuDengJiBiao/findXiangMuDengJiBiao',
    method: 'get',
    params
  })
}

// @Tags XiangMuDengJiBiao
// @Summary 分页获取XiangMuDengJiBiao列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取XiangMuDengJiBiao列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /xiangMuDengJiBiao/getXiangMuDengJiBiaoList [get]
export const getXiangMuDengJiBiaoList = (params) => {
  return service({
    url: '/xiangMuDengJiBiao/getXiangMuDengJiBiaoList',
    method: 'get',
    params
  })
}
