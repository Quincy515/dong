import service from '@/utils/request'

// @Tags GongSiZhangHaoMiMa
// @Summary 创建GongSiZhangHaoMiMa
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.GongSiZhangHaoMiMa true "创建GongSiZhangHaoMiMa"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /gongSiZhangHaoMiMa/createGongSiZhangHaoMiMa [post]
export const createGongSiZhangHaoMiMa = (data) => {
  return service({
    url: '/gongSiZhangHaoMiMa/createGongSiZhangHaoMiMa',
    method: 'post',
    data
  })
}

// @Tags GongSiZhangHaoMiMa
// @Summary 删除GongSiZhangHaoMiMa
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.GongSiZhangHaoMiMa true "删除GongSiZhangHaoMiMa"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /gongSiZhangHaoMiMa/deleteGongSiZhangHaoMiMa [delete]
export const deleteGongSiZhangHaoMiMa = (data) => {
  return service({
    url: '/gongSiZhangHaoMiMa/deleteGongSiZhangHaoMiMa',
    method: 'delete',
    data
  })
}

// @Tags GongSiZhangHaoMiMa
// @Summary 删除GongSiZhangHaoMiMa
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除GongSiZhangHaoMiMa"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /gongSiZhangHaoMiMa/deleteGongSiZhangHaoMiMa [delete]
export const deleteGongSiZhangHaoMiMaByIds = (data) => {
  return service({
    url: '/gongSiZhangHaoMiMa/deleteGongSiZhangHaoMiMaByIds',
    method: 'delete',
    data
  })
}

// @Tags GongSiZhangHaoMiMa
// @Summary 更新GongSiZhangHaoMiMa
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.GongSiZhangHaoMiMa true "更新GongSiZhangHaoMiMa"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /gongSiZhangHaoMiMa/updateGongSiZhangHaoMiMa [put]
export const updateGongSiZhangHaoMiMa = (data) => {
  return service({
    url: '/gongSiZhangHaoMiMa/updateGongSiZhangHaoMiMa',
    method: 'put',
    data
  })
}

// @Tags GongSiZhangHaoMiMa
// @Summary 用id查询GongSiZhangHaoMiMa
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.GongSiZhangHaoMiMa true "用id查询GongSiZhangHaoMiMa"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /gongSiZhangHaoMiMa/findGongSiZhangHaoMiMa [get]
export const findGongSiZhangHaoMiMa = (params) => {
  return service({
    url: '/gongSiZhangHaoMiMa/findGongSiZhangHaoMiMa',
    method: 'get',
    params
  })
}

// @Tags GongSiZhangHaoMiMa
// @Summary 分页获取GongSiZhangHaoMiMa列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取GongSiZhangHaoMiMa列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /gongSiZhangHaoMiMa/getGongSiZhangHaoMiMaList [get]
export const getGongSiZhangHaoMiMaList = (params) => {
  return service({
    url: '/gongSiZhangHaoMiMa/getGongSiZhangHaoMiMaList',
    method: 'get',
    params
  })
}
