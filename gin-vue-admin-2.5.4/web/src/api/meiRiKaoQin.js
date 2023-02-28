import service from '@/utils/request'

// @Tags MeiRiKaoQin
// @Summary 创建MeiRiKaoQin
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.MeiRiKaoQin true "创建MeiRiKaoQin"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /meiRiKaoQin/createMeiRiKaoQin [post]
export const createMeiRiKaoQin = (data) => {
  return service({
    url: '/meiRiKaoQin/createMeiRiKaoQin',
    method: 'post',
    data
  })
}

// @Tags MeiRiKaoQin
// @Summary 删除MeiRiKaoQin
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.MeiRiKaoQin true "删除MeiRiKaoQin"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /meiRiKaoQin/deleteMeiRiKaoQin [delete]
export const deleteMeiRiKaoQin = (data) => {
  return service({
    url: '/meiRiKaoQin/deleteMeiRiKaoQin',
    method: 'delete',
    data
  })
}

// @Tags MeiRiKaoQin
// @Summary 删除MeiRiKaoQin
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除MeiRiKaoQin"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /meiRiKaoQin/deleteMeiRiKaoQin [delete]
export const deleteMeiRiKaoQinByIds = (data) => {
  return service({
    url: '/meiRiKaoQin/deleteMeiRiKaoQinByIds',
    method: 'delete',
    data
  })
}

// @Tags MeiRiKaoQin
// @Summary 更新MeiRiKaoQin
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.MeiRiKaoQin true "更新MeiRiKaoQin"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /meiRiKaoQin/updateMeiRiKaoQin [put]
export const updateMeiRiKaoQin = (data) => {
  return service({
    url: '/meiRiKaoQin/updateMeiRiKaoQin',
    method: 'put',
    data
  })
}

// @Tags MeiRiKaoQin
// @Summary 用id查询MeiRiKaoQin
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.MeiRiKaoQin true "用id查询MeiRiKaoQin"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /meiRiKaoQin/findMeiRiKaoQin [get]
export const findMeiRiKaoQin = (params) => {
  return service({
    url: '/meiRiKaoQin/findMeiRiKaoQin',
    method: 'get',
    params
  })
}

// @Tags MeiRiKaoQin
// @Summary 分页获取MeiRiKaoQin列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取MeiRiKaoQin列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /meiRiKaoQin/getMeiRiKaoQinList [get]
export const getMeiRiKaoQinList = (params) => {
  return service({
    url: '/meiRiKaoQin/getMeiRiKaoQinList',
    method: 'get',
    params
  })
}
