import service from '@/utils/request'

// @Tags YingShouKuanBiao
// @Summary 创建YingShouKuanBiao
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.YingShouKuanBiao true "创建YingShouKuanBiao"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /yingShouKuanBiao/createYingShouKuanBiao [post]
export const createYingShouKuanBiao = (data) => {
  return service({
    url: '/yingShouKuanBiao/createYingShouKuanBiao',
    method: 'post',
    data
  })
}

// @Tags YingShouKuanBiao
// @Summary 删除YingShouKuanBiao
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.YingShouKuanBiao true "删除YingShouKuanBiao"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /yingShouKuanBiao/deleteYingShouKuanBiao [delete]
export const deleteYingShouKuanBiao = (data) => {
  return service({
    url: '/yingShouKuanBiao/deleteYingShouKuanBiao',
    method: 'delete',
    data
  })
}

// @Tags YingShouKuanBiao
// @Summary 删除YingShouKuanBiao
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除YingShouKuanBiao"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /yingShouKuanBiao/deleteYingShouKuanBiao [delete]
export const deleteYingShouKuanBiaoByIds = (data) => {
  return service({
    url: '/yingShouKuanBiao/deleteYingShouKuanBiaoByIds',
    method: 'delete',
    data
  })
}

// @Tags YingShouKuanBiao
// @Summary 更新YingShouKuanBiao
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.YingShouKuanBiao true "更新YingShouKuanBiao"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /yingShouKuanBiao/updateYingShouKuanBiao [put]
export const updateYingShouKuanBiao = (data) => {
  return service({
    url: '/yingShouKuanBiao/updateYingShouKuanBiao',
    method: 'put',
    data
  })
}

// @Tags YingShouKuanBiao
// @Summary 用id查询YingShouKuanBiao
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.YingShouKuanBiao true "用id查询YingShouKuanBiao"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /yingShouKuanBiao/findYingShouKuanBiao [get]
export const findYingShouKuanBiao = (params) => {
  return service({
    url: '/yingShouKuanBiao/findYingShouKuanBiao',
    method: 'get',
    params
  })
}

// @Tags YingShouKuanBiao
// @Summary 分页获取YingShouKuanBiao列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取YingShouKuanBiao列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /yingShouKuanBiao/getYingShouKuanBiaoList [get]
export const getYingShouKuanBiaoList = (params) => {
  return service({
    url: '/yingShouKuanBiao/getYingShouKuanBiaoList',
    method: 'get',
    params
  })
}
