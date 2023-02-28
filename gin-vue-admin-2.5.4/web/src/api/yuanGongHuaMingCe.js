import service from '@/utils/request'

// @Tags YuanGongHuaMingCe
// @Summary 创建YuanGongHuaMingCe
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.YuanGongHuaMingCe true "创建YuanGongHuaMingCe"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /yuanGongHuaMingCe/createYuanGongHuaMingCe [post]
export const createYuanGongHuaMingCe = (data) => {
  return service({
    url: '/yuanGongHuaMingCe/createYuanGongHuaMingCe',
    method: 'post',
    data
  })
}

// @Tags YuanGongHuaMingCe
// @Summary 删除YuanGongHuaMingCe
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.YuanGongHuaMingCe true "删除YuanGongHuaMingCe"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /yuanGongHuaMingCe/deleteYuanGongHuaMingCe [delete]
export const deleteYuanGongHuaMingCe = (data) => {
  return service({
    url: '/yuanGongHuaMingCe/deleteYuanGongHuaMingCe',
    method: 'delete',
    data
  })
}

// @Tags YuanGongHuaMingCe
// @Summary 删除YuanGongHuaMingCe
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除YuanGongHuaMingCe"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /yuanGongHuaMingCe/deleteYuanGongHuaMingCe [delete]
export const deleteYuanGongHuaMingCeByIds = (data) => {
  return service({
    url: '/yuanGongHuaMingCe/deleteYuanGongHuaMingCeByIds',
    method: 'delete',
    data
  })
}

// @Tags YuanGongHuaMingCe
// @Summary 更新YuanGongHuaMingCe
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.YuanGongHuaMingCe true "更新YuanGongHuaMingCe"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /yuanGongHuaMingCe/updateYuanGongHuaMingCe [put]
export const updateYuanGongHuaMingCe = (data) => {
  return service({
    url: '/yuanGongHuaMingCe/updateYuanGongHuaMingCe',
    method: 'put',
    data
  })
}

// @Tags YuanGongHuaMingCe
// @Summary 用id查询YuanGongHuaMingCe
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.YuanGongHuaMingCe true "用id查询YuanGongHuaMingCe"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /yuanGongHuaMingCe/findYuanGongHuaMingCe [get]
export const findYuanGongHuaMingCe = (params) => {
  return service({
    url: '/yuanGongHuaMingCe/findYuanGongHuaMingCe',
    method: 'get',
    params
  })
}

// @Tags YuanGongHuaMingCe
// @Summary 分页获取YuanGongHuaMingCe列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取YuanGongHuaMingCe列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /yuanGongHuaMingCe/getYuanGongHuaMingCeList [get]
export const getYuanGongHuaMingCeList = (params) => {
  return service({
    url: '/yuanGongHuaMingCe/getYuanGongHuaMingCeList',
    method: 'get',
    params
  })
}
