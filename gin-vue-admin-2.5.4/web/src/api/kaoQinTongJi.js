import service from '@/utils/request'

// @Tags KaoQinTongJi
// @Summary 创建KaoQinTongJi
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.KaoQinTongJi true "创建KaoQinTongJi"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /kaoQinTongJi/createKaoQinTongJi [post]
export const createKaoQinTongJi = (data) => {
  return service({
    url: '/kaoQinTongJi/createKaoQinTongJi',
    method: 'post',
    data
  })
}

// @Tags KaoQinTongJi
// @Summary 删除KaoQinTongJi
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.KaoQinTongJi true "删除KaoQinTongJi"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /kaoQinTongJi/deleteKaoQinTongJi [delete]
export const deleteKaoQinTongJi = (data) => {
  return service({
    url: '/kaoQinTongJi/deleteKaoQinTongJi',
    method: 'delete',
    data
  })
}

// @Tags KaoQinTongJi
// @Summary 删除KaoQinTongJi
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除KaoQinTongJi"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /kaoQinTongJi/deleteKaoQinTongJi [delete]
export const deleteKaoQinTongJiByIds = (data) => {
  return service({
    url: '/kaoQinTongJi/deleteKaoQinTongJiByIds',
    method: 'delete',
    data
  })
}

// @Tags KaoQinTongJi
// @Summary 更新KaoQinTongJi
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.KaoQinTongJi true "更新KaoQinTongJi"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /kaoQinTongJi/updateKaoQinTongJi [put]
export const updateKaoQinTongJi = (data) => {
  return service({
    url: '/kaoQinTongJi/updateKaoQinTongJi',
    method: 'put',
    data
  })
}

// @Tags KaoQinTongJi
// @Summary 用id查询KaoQinTongJi
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.KaoQinTongJi true "用id查询KaoQinTongJi"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /kaoQinTongJi/findKaoQinTongJi [get]
export const findKaoQinTongJi = (params) => {
  return service({
    url: '/kaoQinTongJi/findKaoQinTongJi',
    method: 'get',
    params
  })
}

// @Tags KaoQinTongJi
// @Summary 分页获取KaoQinTongJi列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取KaoQinTongJi列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /kaoQinTongJi/getKaoQinTongJiList [get]
export const getKaoQinTongJiList = (params) => {
  return service({
    url: '/kaoQinTongJi/getKaoQinTongJiList',
    method: 'get',
    params
  })
}
