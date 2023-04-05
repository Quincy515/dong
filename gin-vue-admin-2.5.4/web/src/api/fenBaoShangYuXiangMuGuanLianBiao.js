import service from '@/utils/request'

// @Tags FenBaoShangYuXiangMuGuanLianBiao
// @Summary 创建FenBaoShangYuXiangMuGuanLianBiao
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.FenBaoShangYuXiangMuGuanLianBiao true "创建FenBaoShangYuXiangMuGuanLianBiao"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /fenBaoShangYuXiangMuGuanLianBiao/createFenBaoShangYuXiangMuGuanLianBiao [post]
export const createFenBaoShangYuXiangMuGuanLianBiao = (data) => {
  return service({
    url: '/fenBaoShangYuXiangMuGuanLianBiao/createFenBaoShangYuXiangMuGuanLianBiao',
    method: 'post',
    data
  })
}

// @Tags FenBaoShangYuXiangMuGuanLianBiao
// @Summary 删除FenBaoShangYuXiangMuGuanLianBiao
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.FenBaoShangYuXiangMuGuanLianBiao true "删除FenBaoShangYuXiangMuGuanLianBiao"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /fenBaoShangYuXiangMuGuanLianBiao/deleteFenBaoShangYuXiangMuGuanLianBiao [delete]
export const deleteFenBaoShangYuXiangMuGuanLianBiao = (data) => {
  return service({
    url: '/fenBaoShangYuXiangMuGuanLianBiao/deleteFenBaoShangYuXiangMuGuanLianBiao',
    method: 'delete',
    data
  })
}

// @Tags FenBaoShangYuXiangMuGuanLianBiao
// @Summary 删除FenBaoShangYuXiangMuGuanLianBiao
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除FenBaoShangYuXiangMuGuanLianBiao"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /fenBaoShangYuXiangMuGuanLianBiao/deleteFenBaoShangYuXiangMuGuanLianBiao [delete]
export const deleteFenBaoShangYuXiangMuGuanLianBiaoByIds = (data) => {
  return service({
    url: '/fenBaoShangYuXiangMuGuanLianBiao/deleteFenBaoShangYuXiangMuGuanLianBiaoByIds',
    method: 'delete',
    data
  })
}

// @Tags FenBaoShangYuXiangMuGuanLianBiao
// @Summary 更新FenBaoShangYuXiangMuGuanLianBiao
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.FenBaoShangYuXiangMuGuanLianBiao true "更新FenBaoShangYuXiangMuGuanLianBiao"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /fenBaoShangYuXiangMuGuanLianBiao/updateFenBaoShangYuXiangMuGuanLianBiao [put]
export const updateFenBaoShangYuXiangMuGuanLianBiao = (data) => {
  return service({
    url: '/fenBaoShangYuXiangMuGuanLianBiao/updateFenBaoShangYuXiangMuGuanLianBiao',
    method: 'put',
    data
  })
}

// @Tags FenBaoShangYuXiangMuGuanLianBiao
// @Summary 用id查询FenBaoShangYuXiangMuGuanLianBiao
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.FenBaoShangYuXiangMuGuanLianBiao true "用id查询FenBaoShangYuXiangMuGuanLianBiao"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /fenBaoShangYuXiangMuGuanLianBiao/findFenBaoShangYuXiangMuGuanLianBiao [get]
export const findFenBaoShangYuXiangMuGuanLianBiao = (params) => {
  return service({
    url: '/fenBaoShangYuXiangMuGuanLianBiao/findFenBaoShangYuXiangMuGuanLianBiao',
    method: 'get',
    params
  })
}

// @Tags FenBaoShangYuXiangMuGuanLianBiao
// @Summary 分页获取FenBaoShangYuXiangMuGuanLianBiao列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取FenBaoShangYuXiangMuGuanLianBiao列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /fenBaoShangYuXiangMuGuanLianBiao/getFenBaoShangYuXiangMuGuanLianBiaoList [get]
export const getFenBaoShangYuXiangMuGuanLianBiaoList = (params) => {
  return service({
    url: '/fenBaoShangYuXiangMuGuanLianBiao/getFenBaoShangYuXiangMuGuanLianBiaoList',
    method: 'get',
    params
  })
}
