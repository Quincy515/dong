<template>
  <div>
    <div class="gva-search-box">
      <el-form :inline="true" :model="searchInfo" class="demo-form-inline">
      <el-form-item label="创建时间">
      <el-date-picker v-model="searchInfo.startCreatedAt" type="datetime" placeholder="开始时间"></el-date-picker>
       —
      <el-date-picker v-model="searchInfo.endCreatedAt" type="datetime" placeholder="结束时间"></el-date-picker>
      </el-form-item>
           <el-form-item label="发票类型" prop="faPiaoLeiXing">
            <el-select v-model="searchInfo.faPiaoLeiXing" clearable placeholder="请选择" @clear="()=>{searchInfo.faPiaoLeiXing=undefined}">
              <el-option v-for="(item,key) in faPiaoLeiXingOptions" :key="key" :label="item.label" :value="item.value" />
            </el-select>
            </el-form-item>
           <el-form-item label="付款方式" prop="fuKuanFangShi">
            <el-select v-model="searchInfo.fuKuanFangShi" clearable placeholder="请选择" @clear="()=>{searchInfo.fuKuanFangShi=undefined}">
              <el-option v-for="(item,key) in fuKuanFangShiOptions" :key="key" :label="item.label" :value="item.value" />
            </el-select>
            </el-form-item>
        <el-form-item label="合同编号">
         <el-input v-model="searchInfo.heTongBianHao" placeholder="搜索条件" />

        </el-form-item>
           <el-form-item label="合同类别 1.主合同 2.补充合同" prop="heTongLeiBie">
            <el-select v-model="searchInfo.heTongLeiBie" clearable placeholder="请选择" @clear="()=>{searchInfo.heTongLeiBie=undefined}">
              <el-option v-for="(item,key) in heTongLeiBieOptions" :key="key" :label="item.label" :value="item.value" />
            </el-select>
            </el-form-item>
           <el-form-item label="合同类型 1.承包 2.分包 3.采购 4.租聘" prop="heTongLeiXing">
            <el-select v-model="searchInfo.heTongLeiXing" clearable placeholder="请选择" @clear="()=>{searchInfo.heTongLeiXing=undefined}">
              <el-option v-for="(item,key) in heTongLeiXingOptions" :key="key" :label="item.label" :value="item.value" />
            </el-select>
            </el-form-item>
        <el-form-item label="合同名称">
         <el-input v-model="searchInfo.heTongMingCheng" placeholder="搜索条件" />

        </el-form-item>
           <el-form-item label="合同状态" prop="heTongZhuangTai">
            <el-select v-model="searchInfo.heTongZhuangTai" clearable placeholder="请选择" @clear="()=>{searchInfo.heTongZhuangTai=undefined}">
              <el-option v-for="(item,key) in heTongZhuangTaiOptions" :key="key" :label="item.label" :value="item.value" />
            </el-select>
            </el-form-item>
        <el-form-item label="甲方公司">
            
             <el-input v-model.number="searchInfo.jiaFangGongSi" placeholder="搜索条件" />

        </el-form-item>
           <el-form-item label="结算方式" prop="jieSuanFangShi">
            <el-select v-model="searchInfo.jieSuanFangShi" clearable placeholder="请选择" @clear="()=>{searchInfo.jieSuanFangShi=undefined}">
              <el-option v-for="(item,key) in jieSuanFangShiOptions" :key="key" :label="item.label" :value="item.value" />
            </el-select>
            </el-form-item>
        <el-form-item label="签订人">
            
             <el-input v-model.number="searchInfo.qianDingRen" placeholder="搜索条件" />

        </el-form-item>
        <el-form-item label="添加人">
            
             <el-input v-model.number="searchInfo.tianJiaRen" placeholder="搜索条件" />

        </el-form-item>
        <el-form-item label="项目ID">
            
             <el-input v-model.number="searchInfo.xiangMuId" placeholder="搜索条件" />

        </el-form-item>
        <el-form-item label="乙方公司">
            
             <el-input v-model.number="searchInfo.yiFangGongSi" placeholder="搜索条件" />

        </el-form-item>
        <el-form-item>
          <el-button size="small" type="primary" icon="search" @click="onSubmit">查询</el-button>
          <el-button size="small" icon="refresh" @click="onReset">重置</el-button>
        </el-form-item>
      </el-form>
    </div>
    <div class="gva-table-box">
        <div class="gva-btn-list">
            <el-button size="small" type="primary" icon="plus" @click="openDialog">新增</el-button>
            <el-popover v-model:visible="deleteVisible" placement="top" width="160">
            <p>确定要删除吗？</p>
            <div style="text-align: right; margin-top: 8px;">
                <el-button size="small" type="primary" link @click="deleteVisible = false">取消</el-button>
                <el-button size="small" type="primary" @click="onDelete">确定</el-button>
            </div>
            <template #reference>
                <el-button icon="delete" size="small" style="margin-left: 10px;" :disabled="!multipleSelection.length" @click="deleteVisible = true">删除</el-button>
            </template>
            </el-popover>
        </div>
        <el-table
        ref="multipleTable"
        style="width: 100%"
        tooltip-effect="dark"
        :data="tableData"
        row-key="ID"
        @selection-change="handleSelectionChange"
        >
        <el-table-column type="selection" width="55" />
        <el-table-column align="left" label="日期" width="180">
            <template #default="scope">{{ formatDate(scope.row.CreatedAt) }}</template>
        </el-table-column>
        <el-table-column align="left" label="备注" prop="beiZhu" width="120" />
        <el-table-column align="left" label="发票类型" prop="faPiaoLeiXing" width="120">
            <template #default="scope">
            {{ filterDict(scope.row.faPiaoLeiXing,faPiaoLeiXingOptions) }}
            </template>
        </el-table-column>
        <el-table-column align="left" label="付款方式" prop="fuKuanFangShi" width="120">
            <template #default="scope">
            {{ filterDict(scope.row.fuKuanFangShi,fuKuanFangShiOptions) }}
            </template>
        </el-table-column>
        <el-table-column align="left" label="所属" prop="group" width="120" />
        <el-table-column align="left" label="合同编号" prop="heTongBianHao" width="120" />
        <el-table-column align="left" label="合同金额" prop="heTongJinE" width="120" />
        <el-table-column align="left" label="合同类别 1.主合同 2.补充合同" prop="heTongLeiBie" width="120">
            <template #default="scope">
            {{ filterDict(scope.row.heTongLeiBie,heTongLeiBieOptions) }}
            </template>
        </el-table-column>
        <el-table-column align="left" label="合同类型 1.承包 2.分包 3.采购 4.租聘" prop="heTongLeiXing" width="120">
            <template #default="scope">
            {{ filterDict(scope.row.heTongLeiXing,heTongLeiXingOptions) }}
            </template>
        </el-table-column>
        <el-table-column align="left" label="合同名称" prop="heTongMingCheng" width="120" />
        <el-table-column align="left" label="合同下载地址" prop="heTongXiaZaiDiZhi" width="120" />
        <el-table-column align="left" label="合同状态" prop="heTongZhuangTai" width="120">
            <template #default="scope">
            {{ filterDict(scope.row.heTongZhuangTai,heTongZhuangTaiOptions) }}
            </template>
        </el-table-column>
        <el-table-column align="left" label="甲方公司" prop="jiaFangGongSi" width="120" />
        <el-table-column align="left" label="结算方式" prop="jieSuanFangShi" width="120">
            <template #default="scope">
            {{ filterDict(scope.row.jieSuanFangShi,jieSuanFangShiOptions) }}
            </template>
        </el-table-column>
        <el-table-column align="left" label="开户行" prop="kaiHuHang" width="120" />
        <el-table-column align="left" label="开票单位" prop="kaiPiaoDanWei" width="120" />
        <el-table-column align="left" label="签订人" prop="qianDingRen" width="120" />
         <el-table-column align="left" label="签订日期" width="180">
            <template #default="scope">{{ formatDate(scope.row.qianDingRiQi) }}</template>
         </el-table-column>
        <el-table-column align="left" label="收付款条件" prop="shouFuKuanTiaoJian" width="120" />
        <el-table-column align="left" label="收款账号" prop="shouKuanZhangHao" width="120" />
        <el-table-column align="left" label="添加人" prop="tianJiaRen" width="120" />
         <el-table-column align="left" label="添加时间" width="180">
            <template #default="scope">{{ formatDate(scope.row.tianJiaShiJian) }}</template>
         </el-table-column>
        <el-table-column align="left" label="项目ID" prop="xiangMuId" width="120" />
        <el-table-column align="left" label="乙方公司" prop="yiFangGongSi" width="120" />
        <el-table-column align="left" label="银行名称" prop="yinHangMingCheng" width="120" />
        <el-table-column align="left" label="主要条款" prop="zhuYaoTiaoKuan" width="120" />
        <el-table-column align="left" label="按钮组">
            <template #default="scope">
            <el-button type="primary" link icon="edit" size="small" class="table-button" @click="updateHeTongWenJianBiaoFunc(scope.row)">变更</el-button>
            <el-button type="primary" link icon="delete" size="small" @click="deleteRow(scope.row)">删除</el-button>
            </template>
        </el-table-column>
        </el-table>
        <div class="gva-pagination">
            <el-pagination
            layout="total, sizes, prev, pager, next, jumper"
            :current-page="page"
            :page-size="pageSize"
            :page-sizes="[10, 30, 50, 100]"
            :total="total"
            @current-change="handleCurrentChange"
            @size-change="handleSizeChange"
            />
        </div>
    </div>
    <el-dialog v-model="dialogFormVisible" :before-close="closeDialog" title="弹窗操作">
      <el-form :model="formData" label-position="right" ref="elFormRef" :rules="rule" label-width="80px">
        <el-form-item label="备注:"  prop="beiZhu" >
          <el-input v-model="formData.beiZhu" :clearable="true"  placeholder="请输入" />
        </el-form-item>
        <el-form-item label="发票类型:"  prop="faPiaoLeiXing" >
          <el-select v-model="formData.faPiaoLeiXing" placeholder="请选择" style="width:100%" :clearable="true" >
            <el-option v-for="(item,key) in faPiaoLeiXingOptions" :key="key" :label="item.label" :value="item.value" />
          </el-select>
        </el-form-item>
        <el-form-item label="付款方式:"  prop="fuKuanFangShi" >
          <el-select v-model="formData.fuKuanFangShi" placeholder="请选择" style="width:100%" :clearable="true" >
            <el-option v-for="(item,key) in fuKuanFangShiOptions" :key="key" :label="item.label" :value="item.value" />
          </el-select>
        </el-form-item>
        <el-form-item label="所属:"  prop="group" >
          <el-input v-model="formData.group" :clearable="true"  placeholder="请输入" />
        </el-form-item>
        <el-form-item label="合同编号:"  prop="heTongBianHao" >
          <el-input v-model="formData.heTongBianHao" :clearable="true"  placeholder="请输入" />
        </el-form-item>
        <el-form-item label="合同金额:"  prop="heTongJinE" >
          <el-input-number v-model="formData.heTongJinE"  style="width:100%" :precision="2" :clearable="true"  />
        </el-form-item>
        <el-form-item label="合同类别 1.主合同 2.补充合同:"  prop="heTongLeiBie" >
          <el-select v-model="formData.heTongLeiBie" placeholder="请选择" style="width:100%" :clearable="true" >
            <el-option v-for="(item,key) in heTongLeiBieOptions" :key="key" :label="item.label" :value="item.value" />
          </el-select>
        </el-form-item>
        <el-form-item label="合同类型 1.承包 2.分包 3.采购 4.租聘:"  prop="heTongLeiXing" >
          <el-select v-model="formData.heTongLeiXing" placeholder="请选择" style="width:100%" :clearable="true" >
            <el-option v-for="(item,key) in heTongLeiXingOptions" :key="key" :label="item.label" :value="item.value" />
          </el-select>
        </el-form-item>
        <el-form-item label="合同名称:"  prop="heTongMingCheng" >
          <el-input v-model="formData.heTongMingCheng" :clearable="true"  placeholder="请输入" />
        </el-form-item>
        <el-form-item label="合同下载地址:"  prop="heTongXiaZaiDiZhi" >
          <el-input v-model="formData.heTongXiaZaiDiZhi" :clearable="true"  placeholder="请输入" />
        </el-form-item>
        <el-form-item label="合同状态:"  prop="heTongZhuangTai" >
          <el-select v-model="formData.heTongZhuangTai" placeholder="请选择" style="width:100%" :clearable="true" >
            <el-option v-for="(item,key) in heTongZhuangTaiOptions" :key="key" :label="item.label" :value="item.value" />
          </el-select>
        </el-form-item>
        <el-form-item label="甲方公司:"  prop="jiaFangGongSi" >
          <el-input v-model.number="formData.jiaFangGongSi" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="结算方式:"  prop="jieSuanFangShi" >
          <el-select v-model="formData.jieSuanFangShi" placeholder="请选择" style="width:100%" :clearable="true" >
            <el-option v-for="(item,key) in jieSuanFangShiOptions" :key="key" :label="item.label" :value="item.value" />
          </el-select>
        </el-form-item>
        <el-form-item label="开户行:"  prop="kaiHuHang" >
          <el-input v-model="formData.kaiHuHang" :clearable="true"  placeholder="请输入" />
        </el-form-item>
        <el-form-item label="开票单位:"  prop="kaiPiaoDanWei" >
          <el-input v-model="formData.kaiPiaoDanWei" :clearable="true"  placeholder="请输入" />
        </el-form-item>
        <el-form-item label="签订人:"  prop="qianDingRen" >
          <el-input v-model.number="formData.qianDingRen" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="签订日期:"  prop="qianDingRiQi" >
          <el-date-picker v-model="formData.qianDingRiQi" type="date" style="width:100%" placeholder="选择日期" :clearable="true"  />
        </el-form-item>
        <el-form-item label="收付款条件:"  prop="shouFuKuanTiaoJian" >
          <el-input v-model="formData.shouFuKuanTiaoJian" :clearable="true"  placeholder="请输入" />
        </el-form-item>
        <el-form-item label="收款账号:"  prop="shouKuanZhangHao" >
          <el-input v-model="formData.shouKuanZhangHao" :clearable="true"  placeholder="请输入" />
        </el-form-item>
        <el-form-item label="添加人:"  prop="tianJiaRen" >
          <el-input v-model.number="formData.tianJiaRen" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="添加时间:"  prop="tianJiaShiJian" >
          <el-date-picker v-model="formData.tianJiaShiJian" type="date" style="width:100%" placeholder="选择日期" :clearable="true"  />
        </el-form-item>
        <el-form-item label="项目ID:"  prop="xiangMuId" >
          <el-input v-model.number="formData.xiangMuId" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="乙方公司:"  prop="yiFangGongSi" >
          <el-input v-model.number="formData.yiFangGongSi" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="银行名称:"  prop="yinHangMingCheng" >
          <el-input v-model="formData.yinHangMingCheng" :clearable="true"  placeholder="请输入" />
        </el-form-item>
        <el-form-item label="主要条款:"  prop="zhuYaoTiaoKuan" >
          <el-input v-model="formData.zhuYaoTiaoKuan" :clearable="true"  placeholder="请输入" />
        </el-form-item>
      </el-form>
      <template #footer>
        <div class="dialog-footer">
          <el-button size="small" @click="closeDialog">取 消</el-button>
          <el-button size="small" type="primary" @click="enterDialog">确 定</el-button>
        </div>
      </template>
    </el-dialog>
  </div>
</template>

<script>
export default {
  name: 'HeTongWenJianBiao'
}
</script>

<script setup>
import {
  createHeTongWenJianBiao,
  deleteHeTongWenJianBiao,
  deleteHeTongWenJianBiaoByIds,
  updateHeTongWenJianBiao,
  findHeTongWenJianBiao,
  getHeTongWenJianBiaoList
} from '@/api/heTongWenJianBiao'

// 全量引入格式化工具 请按需保留
import { getDictFunc, formatDate, formatBoolean, filterDict } from '@/utils/format'
import { ElMessage, ElMessageBox } from 'element-plus'
import { ref, reactive } from 'vue'

// 自动化生成的字典（可能为空）以及字段
const faPiaoLeiXingOptions = ref([])
const fuKuanFangShiOptions = ref([])
const heTongLeiBieOptions = ref([])
const heTongLeiXingOptions = ref([])
const heTongZhuangTaiOptions = ref([])
const jieSuanFangShiOptions = ref([])
const formData = ref({
        beiZhu: '',
        faPiaoLeiXing: undefined,
        fuKuanFangShi: undefined,
        group: '',
        heTongBianHao: '',
        heTongJinE: 0,
        heTongLeiBie: undefined,
        heTongLeiXing: undefined,
        heTongMingCheng: '',
        heTongXiaZaiDiZhi: '',
        heTongZhuangTai: undefined,
        jiaFangGongSi: 0,
        jieSuanFangShi: undefined,
        kaiHuHang: '',
        kaiPiaoDanWei: '',
        qianDingRen: 0,
        qianDingRiQi: new Date(),
        shouFuKuanTiaoJian: '',
        shouKuanZhangHao: '',
        tianJiaRen: 0,
        tianJiaShiJian: new Date(),
        xiangMuId: 0,
        yiFangGongSi: 0,
        yinHangMingCheng: '',
        zhuYaoTiaoKuan: '',
        })

// 验证规则
const rule = reactive({
})

const elFormRef = ref()


// =========== 表格控制部分 ===========
const page = ref(1)
const total = ref(0)
const pageSize = ref(10)
const tableData = ref([])
const searchInfo = ref({})

// 重置
const onReset = () => {
  searchInfo.value = {}
}

// 搜索
const onSubmit = () => {
  page.value = 1
  pageSize.value = 10
  getTableData()
}

// 分页
const handleSizeChange = (val) => {
  pageSize.value = val
  getTableData()
}

// 修改页面容量
const handleCurrentChange = (val) => {
  page.value = val
  getTableData()
}

// 查询
const getTableData = async() => {
  const table = await getHeTongWenJianBiaoList({ page: page.value, pageSize: pageSize.value, ...searchInfo.value })
  if (table.code === 0) {
    tableData.value = table.data.list
    total.value = table.data.total
    page.value = table.data.page
    pageSize.value = table.data.pageSize
  }
}

getTableData()

// ============== 表格控制部分结束 ===============

// 获取需要的字典 可能为空 按需保留
const setOptions = async () =>{
    faPiaoLeiXingOptions.value = await getDictFunc('faPiaoLeiXing')
    fuKuanFangShiOptions.value = await getDictFunc('fuKuanFangShi')
    heTongLeiBieOptions.value = await getDictFunc('heTongLeiBie')
    heTongLeiXingOptions.value = await getDictFunc('heTongLeiXing')
    heTongZhuangTaiOptions.value = await getDictFunc('heTongZhuangTai')
    jieSuanFangShiOptions.value = await getDictFunc('jieSuanFangShi')
}

// 获取需要的字典 可能为空 按需保留
setOptions()


// 多选数据
const multipleSelection = ref([])
// 多选
const handleSelectionChange = (val) => {
    multipleSelection.value = val
}

// 删除行
const deleteRow = (row) => {
    ElMessageBox.confirm('确定要删除吗?', '提示', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning'
    }).then(() => {
            deleteHeTongWenJianBiaoFunc(row)
        })
    }


// 批量删除控制标记
const deleteVisible = ref(false)

// 多选删除
const onDelete = async() => {
      const ids = []
      if (multipleSelection.value.length === 0) {
        ElMessage({
          type: 'warning',
          message: '请选择要删除的数据'
        })
        return
      }
      multipleSelection.value &&
        multipleSelection.value.map(item => {
          ids.push(item.ID)
        })
      const res = await deleteHeTongWenJianBiaoByIds({ ids })
      if (res.code === 0) {
        ElMessage({
          type: 'success',
          message: '删除成功'
        })
        if (tableData.value.length === ids.length && page.value > 1) {
          page.value--
        }
        deleteVisible.value = false
        getTableData()
      }
    }

// 行为控制标记（弹窗内部需要增还是改）
const type = ref('')

// 更新行
const updateHeTongWenJianBiaoFunc = async(row) => {
    const res = await findHeTongWenJianBiao({ ID: row.ID })
    type.value = 'update'
    if (res.code === 0) {
        formData.value = res.data.reheTongWenJianBiao
        dialogFormVisible.value = true
    }
}


// 删除行
const deleteHeTongWenJianBiaoFunc = async (row) => {
    const res = await deleteHeTongWenJianBiao({ ID: row.ID })
    if (res.code === 0) {
        ElMessage({
                type: 'success',
                message: '删除成功'
            })
            if (tableData.value.length === 1 && page.value > 1) {
            page.value--
        }
        getTableData()
    }
}

// 弹窗控制标记
const dialogFormVisible = ref(false)

// 打开弹窗
const openDialog = () => {
    type.value = 'create'
    dialogFormVisible.value = true
}

// 关闭弹窗
const closeDialog = () => {
    dialogFormVisible.value = false
    formData.value = {
        beiZhu: '',
        faPiaoLeiXing: undefined,
        fuKuanFangShi: undefined,
        group: '',
        heTongBianHao: '',
        heTongJinE: 0,
        heTongLeiBie: undefined,
        heTongLeiXing: undefined,
        heTongMingCheng: '',
        heTongXiaZaiDiZhi: '',
        heTongZhuangTai: undefined,
        jiaFangGongSi: 0,
        jieSuanFangShi: undefined,
        kaiHuHang: '',
        kaiPiaoDanWei: '',
        qianDingRen: 0,
        qianDingRiQi: new Date(),
        shouFuKuanTiaoJian: '',
        shouKuanZhangHao: '',
        tianJiaRen: 0,
        tianJiaShiJian: new Date(),
        xiangMuId: 0,
        yiFangGongSi: 0,
        yinHangMingCheng: '',
        zhuYaoTiaoKuan: '',
        }
}
// 弹窗确定
const enterDialog = async () => {
     elFormRef.value?.validate( async (valid) => {
             if (!valid) return
              let res
              switch (type.value) {
                case 'create':
                  res = await createHeTongWenJianBiao(formData.value)
                  break
                case 'update':
                  res = await updateHeTongWenJianBiao(formData.value)
                  break
                default:
                  res = await createHeTongWenJianBiao(formData.value)
                  break
              }
              if (res.code === 0) {
                ElMessage({
                  type: 'success',
                  message: '创建/更改成功'
                })
                closeDialog()
                getTableData()
              }
      })
}
</script>

<style>
</style>
