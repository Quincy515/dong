import service from '@/utils/request'

// @Tags ShiGongJinDuBiao
// @Summary 创建ShiGongJinDuBiao
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.ShiGongJinDuBiao true "创建ShiGongJinDuBiao"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /shiGongJinDuBiao/createShiGongJinDuBiao [post]
export const createShiGongJinDuBiao = (data) => {
  return service({
    url: '/shiGongJinDuBiao/createShiGongJinDuBiao',
    method: 'post',
    data
  })
}

// @Tags ShiGongJinDuBiao
// @Summary 删除ShiGongJinDuBiao
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.ShiGongJinDuBiao true "删除ShiGongJinDuBiao"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /shiGongJinDuBiao/deleteShiGongJinDuBiao [delete]
export const deleteShiGongJinDuBiao = (data) => {
  return service({
    url: '/shiGongJinDuBiao/deleteShiGongJinDuBiao',
    method: 'delete',
    data
  })
}

// @Tags ShiGongJinDuBiao
// @Summary 删除ShiGongJinDuBiao
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除ShiGongJinDuBiao"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /shiGongJinDuBiao/deleteShiGongJinDuBiao [delete]
export const deleteShiGongJinDuBiaoByIds = (data) => {
  return service({
    url: '/shiGongJinDuBiao/deleteShiGongJinDuBiaoByIds',
    method: 'delete',
    data
  })
}

// @Tags ShiGongJinDuBiao
// @Summary 更新ShiGongJinDuBiao
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.ShiGongJinDuBiao true "更新ShiGongJinDuBiao"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /shiGongJinDuBiao/updateShiGongJinDuBiao [put]
export const updateShiGongJinDuBiao = (data) => {
  return service({
    url: '/shiGongJinDuBiao/updateShiGongJinDuBiao',
    method: 'put',
    data
  })
}

// @Tags ShiGongJinDuBiao
// @Summary 用id查询ShiGongJinDuBiao
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.ShiGongJinDuBiao true "用id查询ShiGongJinDuBiao"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /shiGongJinDuBiao/findShiGongJinDuBiao [get]
export const findShiGongJinDuBiao = (params) => {
  return service({
    url: '/shiGongJinDuBiao/findShiGongJinDuBiao',
    method: 'get',
    params
  })
}

// @Tags ShiGongJinDuBiao
// @Summary 分页获取ShiGongJinDuBiao列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取ShiGongJinDuBiao列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /shiGongJinDuBiao/getShiGongJinDuBiaoList [get]
export const getShiGongJinDuBiaoList = (params) => {
  return service({
    url: '/shiGongJinDuBiao/getShiGongJinDuBiaoList',
    method: 'get',
    params
  })
}
