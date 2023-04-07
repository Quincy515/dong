import service from '@/utils/request'

// @Tags ChuRuKuDanJuBiao
// @Summary 创建ChuRuKuDanJuBiao
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.ChuRuKuDanJuBiao true "创建ChuRuKuDanJuBiao"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /chuRuKuDanJuBiao/createChuRuKuDanJuBiao [post]
export const createChuRuKuDanJuBiao = (data) => {
  return service({
    url: '/chuRuKuDanJuBiao/createChuRuKuDanJuBiao',
    method: 'post',
    data
  })
}

// @Tags ChuRuKuDanJuBiao
// @Summary 删除ChuRuKuDanJuBiao
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.ChuRuKuDanJuBiao true "删除ChuRuKuDanJuBiao"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /chuRuKuDanJuBiao/deleteChuRuKuDanJuBiao [delete]
export const deleteChuRuKuDanJuBiao = (data) => {
  return service({
    url: '/chuRuKuDanJuBiao/deleteChuRuKuDanJuBiao',
    method: 'delete',
    data
  })
}

// @Tags ChuRuKuDanJuBiao
// @Summary 删除ChuRuKuDanJuBiao
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除ChuRuKuDanJuBiao"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /chuRuKuDanJuBiao/deleteChuRuKuDanJuBiao [delete]
export const deleteChuRuKuDanJuBiaoByIds = (data) => {
  return service({
    url: '/chuRuKuDanJuBiao/deleteChuRuKuDanJuBiaoByIds',
    method: 'delete',
    data
  })
}

// @Tags ChuRuKuDanJuBiao
// @Summary 更新ChuRuKuDanJuBiao
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.ChuRuKuDanJuBiao true "更新ChuRuKuDanJuBiao"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /chuRuKuDanJuBiao/updateChuRuKuDanJuBiao [put]
export const updateChuRuKuDanJuBiao = (data) => {
  return service({
    url: '/chuRuKuDanJuBiao/updateChuRuKuDanJuBiao',
    method: 'put',
    data
  })
}

// @Tags ChuRuKuDanJuBiao
// @Summary 用id查询ChuRuKuDanJuBiao
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.ChuRuKuDanJuBiao true "用id查询ChuRuKuDanJuBiao"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /chuRuKuDanJuBiao/findChuRuKuDanJuBiao [get]
export const findChuRuKuDanJuBiao = (params) => {
  return service({
    url: '/chuRuKuDanJuBiao/findChuRuKuDanJuBiao',
    method: 'get',
    params
  })
}

// @Tags ChuRuKuDanJuBiao
// @Summary 分页获取ChuRuKuDanJuBiao列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取ChuRuKuDanJuBiao列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /chuRuKuDanJuBiao/getChuRuKuDanJuBiaoList [get]
export const getChuRuKuDanJuBiaoList = (params) => {
  return service({
    url: '/chuRuKuDanJuBiao/getChuRuKuDanJuBiaoList',
    method: 'get',
    params
  })
}
