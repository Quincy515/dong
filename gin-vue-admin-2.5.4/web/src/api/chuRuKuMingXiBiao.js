import service from '@/utils/request'

// @Tags ChuRuKuMingXiBiao
// @Summary 创建ChuRuKuMingXiBiao
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.ChuRuKuMingXiBiao true "创建ChuRuKuMingXiBiao"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /chuRuKuMingXiBiao/createChuRuKuMingXiBiao [post]
export const createChuRuKuMingXiBiao = (data) => {
  return service({
    url: '/chuRuKuMingXiBiao/createChuRuKuMingXiBiao',
    method: 'post',
    data
  })
}

// @Tags ChuRuKuMingXiBiao
// @Summary 删除ChuRuKuMingXiBiao
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.ChuRuKuMingXiBiao true "删除ChuRuKuMingXiBiao"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /chuRuKuMingXiBiao/deleteChuRuKuMingXiBiao [delete]
export const deleteChuRuKuMingXiBiao = (data) => {
  return service({
    url: '/chuRuKuMingXiBiao/deleteChuRuKuMingXiBiao',
    method: 'delete',
    data
  })
}

// @Tags ChuRuKuMingXiBiao
// @Summary 删除ChuRuKuMingXiBiao
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除ChuRuKuMingXiBiao"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /chuRuKuMingXiBiao/deleteChuRuKuMingXiBiao [delete]
export const deleteChuRuKuMingXiBiaoByIds = (data) => {
  return service({
    url: '/chuRuKuMingXiBiao/deleteChuRuKuMingXiBiaoByIds',
    method: 'delete',
    data
  })
}

// @Tags ChuRuKuMingXiBiao
// @Summary 更新ChuRuKuMingXiBiao
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.ChuRuKuMingXiBiao true "更新ChuRuKuMingXiBiao"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /chuRuKuMingXiBiao/updateChuRuKuMingXiBiao [put]
export const updateChuRuKuMingXiBiao = (data) => {
  return service({
    url: '/chuRuKuMingXiBiao/updateChuRuKuMingXiBiao',
    method: 'put',
    data
  })
}

// @Tags ChuRuKuMingXiBiao
// @Summary 用id查询ChuRuKuMingXiBiao
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.ChuRuKuMingXiBiao true "用id查询ChuRuKuMingXiBiao"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /chuRuKuMingXiBiao/findChuRuKuMingXiBiao [get]
export const findChuRuKuMingXiBiao = (params) => {
  return service({
    url: '/chuRuKuMingXiBiao/findChuRuKuMingXiBiao',
    method: 'get',
    params
  })
}

// @Tags ChuRuKuMingXiBiao
// @Summary 分页获取ChuRuKuMingXiBiao列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取ChuRuKuMingXiBiao列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /chuRuKuMingXiBiao/getChuRuKuMingXiBiaoList [get]
export const getChuRuKuMingXiBiaoList = (params) => {
  return service({
    url: '/chuRuKuMingXiBiao/getChuRuKuMingXiBiaoList',
    method: 'get',
    params
  })
}
