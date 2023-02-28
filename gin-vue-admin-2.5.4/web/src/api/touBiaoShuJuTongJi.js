import service from '@/utils/request'

// @Tags TouBiaoShuJuTongJi
// @Summary 创建TouBiaoShuJuTongJi
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.TouBiaoShuJuTongJi true "创建TouBiaoShuJuTongJi"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /touBiaoShuJuTongJi/createTouBiaoShuJuTongJi [post]
export const createTouBiaoShuJuTongJi = (data) => {
  return service({
    url: '/touBiaoShuJuTongJi/createTouBiaoShuJuTongJi',
    method: 'post',
    data
  })
}

// @Tags TouBiaoShuJuTongJi
// @Summary 删除TouBiaoShuJuTongJi
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.TouBiaoShuJuTongJi true "删除TouBiaoShuJuTongJi"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /touBiaoShuJuTongJi/deleteTouBiaoShuJuTongJi [delete]
export const deleteTouBiaoShuJuTongJi = (data) => {
  return service({
    url: '/touBiaoShuJuTongJi/deleteTouBiaoShuJuTongJi',
    method: 'delete',
    data
  })
}

// @Tags TouBiaoShuJuTongJi
// @Summary 删除TouBiaoShuJuTongJi
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除TouBiaoShuJuTongJi"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /touBiaoShuJuTongJi/deleteTouBiaoShuJuTongJi [delete]
export const deleteTouBiaoShuJuTongJiByIds = (data) => {
  return service({
    url: '/touBiaoShuJuTongJi/deleteTouBiaoShuJuTongJiByIds',
    method: 'delete',
    data
  })
}

// @Tags TouBiaoShuJuTongJi
// @Summary 更新TouBiaoShuJuTongJi
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.TouBiaoShuJuTongJi true "更新TouBiaoShuJuTongJi"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /touBiaoShuJuTongJi/updateTouBiaoShuJuTongJi [put]
export const updateTouBiaoShuJuTongJi = (data) => {
  return service({
    url: '/touBiaoShuJuTongJi/updateTouBiaoShuJuTongJi',
    method: 'put',
    data
  })
}

// @Tags TouBiaoShuJuTongJi
// @Summary 用id查询TouBiaoShuJuTongJi
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.TouBiaoShuJuTongJi true "用id查询TouBiaoShuJuTongJi"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /touBiaoShuJuTongJi/findTouBiaoShuJuTongJi [get]
export const findTouBiaoShuJuTongJi = (params) => {
  return service({
    url: '/touBiaoShuJuTongJi/findTouBiaoShuJuTongJi',
    method: 'get',
    params
  })
}

// @Tags TouBiaoShuJuTongJi
// @Summary 分页获取TouBiaoShuJuTongJi列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取TouBiaoShuJuTongJi列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /touBiaoShuJuTongJi/getTouBiaoShuJuTongJiList [get]
export const getTouBiaoShuJuTongJiList = (params) => {
  return service({
    url: '/touBiaoShuJuTongJi/getTouBiaoShuJuTongJiList',
    method: 'get',
    params
  })
}
