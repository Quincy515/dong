<template>
  <div>
    <div class="gva-search-box">
      <el-form :inline="true" :model="searchInfo" class="demo-form-inline">
      <el-form-item label="创建时间">
      <el-date-picker v-model="searchInfo.startCreatedAt" type="datetime" placeholder="开始时间"></el-date-picker>
       —
      <el-date-picker v-model="searchInfo.endCreatedAt" type="datetime" placeholder="结束时间"></el-date-picker>
      </el-form-item>
        <el-form-item label="法人">
         <el-input v-model="searchInfo.faRen" placeholder="搜索条件" />

        </el-form-item>
           <el-form-item label="分包商类型" prop="fenBaoShangLeiXing">
            <el-select v-model="searchInfo.fenBaoShangLeiXing" clearable placeholder="请选择" @clear="()=>{searchInfo.fenBaoShangLeiXing=undefined}">
              <el-option v-for="(item,key) in fenBaoShangLeiXingOptions" :key="key" :label="item.label" :value="item.value" />
            </el-select>
            </el-form-item>
        <el-form-item label="分包商名称">
         <el-input v-model="searchInfo.fenBaoShangMingCheng" placeholder="搜索条件" />

        </el-form-item>
        <el-form-item label="添加人">
            
             <el-input v-model.number="searchInfo.tianJiaRen" placeholder="搜索条件" />

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
        <el-table-column align="left" label="法人" prop="faRen" width="120" />
        <el-table-column align="left" label="法人手机" prop="faRenShouJi" width="120" />
        <el-table-column align="left" label="分包商类型" prop="fenBaoShangLeiXing" width="120">
            <template #default="scope">
            {{ filterDict(scope.row.fenBaoShangLeiXing,fenBaoShangLeiXingOptions) }}
            </template>
        </el-table-column>
        <el-table-column align="left" label="分包商名称" prop="fenBaoShangMingCheng" width="120" />
        <el-table-column align="left" label="所属" prop="group" width="120" />
        <el-table-column align="left" label="开户行" prop="kaiHuHang" width="120" />
        <el-table-column align="left" label="社会信用代码" prop="sheHuiXinYongDaiMa" width="120" />
        <el-table-column align="left" label="添加人" prop="tianJiaRen" width="120" />
         <el-table-column align="left" label="添加时间" width="180">
            <template #default="scope">{{ formatDate(scope.row.tianJiaShiJian) }}</template>
         </el-table-column>
        <el-table-column align="left" label="银行名称" prop="yinHangMingCheng" width="120" />
        <el-table-column align="left" label="账号" prop="zhangHao" width="120" />
        <el-table-column align="left" label="状态" prop="zhuangTai" width="120" />
        <el-table-column align="left" label="资质证照" prop="ziZhiZhengZhao" width="120" />
        <el-table-column align="left" label="按钮组">
            <template #default="scope">
            <el-button type="primary" link icon="edit" size="small" class="table-button" @click="updateFenBaoShangBiaoFunc(scope.row)">变更</el-button>
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
        <el-form-item label="法人:"  prop="faRen" >
          <el-input v-model="formData.faRen" :clearable="true"  placeholder="请输入" />
        </el-form-item>
        <el-form-item label="法人手机:"  prop="faRenShouJi" >
          <el-input v-model="formData.faRenShouJi" :clearable="true"  placeholder="请输入" />
        </el-form-item>
        <el-form-item label="分包商类型:"  prop="fenBaoShangLeiXing" >
          <el-select v-model="formData.fenBaoShangLeiXing" placeholder="请选择" style="width:100%" :clearable="true" >
            <el-option v-for="(item,key) in fenBaoShangLeiXingOptions" :key="key" :label="item.label" :value="item.value" />
          </el-select>
        </el-form-item>
        <el-form-item label="分包商名称:"  prop="fenBaoShangMingCheng" >
          <el-input v-model="formData.fenBaoShangMingCheng" :clearable="true"  placeholder="请输入" />
        </el-form-item>
        <el-form-item label="所属:"  prop="group" >
          <el-input v-model="formData.group" :clearable="true"  placeholder="请输入" />
        </el-form-item>
        <el-form-item label="开户行:"  prop="kaiHuHang" >
          <el-input v-model="formData.kaiHuHang" :clearable="true"  placeholder="请输入" />
        </el-form-item>
        <el-form-item label="社会信用代码:"  prop="sheHuiXinYongDaiMa" >
          <el-input v-model="formData.sheHuiXinYongDaiMa" :clearable="true"  placeholder="请输入" />
        </el-form-item>
        <el-form-item label="添加人:"  prop="tianJiaRen" >
          <el-input v-model.number="formData.tianJiaRen" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="添加时间:"  prop="tianJiaShiJian" >
          <el-date-picker v-model="formData.tianJiaShiJian" type="date" style="width:100%" placeholder="选择日期" :clearable="true"  />
        </el-form-item>
        <el-form-item label="银行名称:"  prop="yinHangMingCheng" >
          <el-input v-model="formData.yinHangMingCheng" :clearable="true"  placeholder="请输入" />
        </el-form-item>
        <el-form-item label="账号:"  prop="zhangHao" >
          <el-input v-model="formData.zhangHao" :clearable="true"  placeholder="请输入" />
        </el-form-item>
        <el-form-item label="状态:"  prop="zhuangTai" >
          <el-input v-model.number="formData.zhuangTai" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="资质证照:"  prop="ziZhiZhengZhao" >
          <el-input v-model="formData.ziZhiZhengZhao" :clearable="true"  placeholder="请输入" />
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
  name: 'FenBaoShangBiao'
}
</script>

<script setup>
import {
  createFenBaoShangBiao,
  deleteFenBaoShangBiao,
  deleteFenBaoShangBiaoByIds,
  updateFenBaoShangBiao,
  findFenBaoShangBiao,
  getFenBaoShangBiaoList
} from '@/api/fenBaoShangBiao'

// 全量引入格式化工具 请按需保留
import { getDictFunc, formatDate, formatBoolean, filterDict } from '@/utils/format'
import { ElMessage, ElMessageBox } from 'element-plus'
import { ref, reactive } from 'vue'

// 自动化生成的字典（可能为空）以及字段
const fenBaoShangLeiXingOptions = ref([])
const formData = ref({
        beiZhu: '',
        faRen: '',
        faRenShouJi: '',
        fenBaoShangLeiXing: undefined,
        fenBaoShangMingCheng: '',
        group: '',
        kaiHuHang: '',
        sheHuiXinYongDaiMa: '',
        tianJiaRen: 0,
        tianJiaShiJian: new Date(),
        yinHangMingCheng: '',
        zhangHao: '',
        zhuangTai: 0,
        ziZhiZhengZhao: '',
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
  const table = await getFenBaoShangBiaoList({ page: page.value, pageSize: pageSize.value, ...searchInfo.value })
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
    fenBaoShangLeiXingOptions.value = await getDictFunc('fenBaoShangLeiXing')
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
            deleteFenBaoShangBiaoFunc(row)
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
      const res = await deleteFenBaoShangBiaoByIds({ ids })
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
const updateFenBaoShangBiaoFunc = async(row) => {
    const res = await findFenBaoShangBiao({ ID: row.ID })
    type.value = 'update'
    if (res.code === 0) {
        formData.value = res.data.refenBaoShangBiao
        dialogFormVisible.value = true
    }
}


// 删除行
const deleteFenBaoShangBiaoFunc = async (row) => {
    const res = await deleteFenBaoShangBiao({ ID: row.ID })
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
        faRen: '',
        faRenShouJi: '',
        fenBaoShangLeiXing: undefined,
        fenBaoShangMingCheng: '',
        group: '',
        kaiHuHang: '',
        sheHuiXinYongDaiMa: '',
        tianJiaRen: 0,
        tianJiaShiJian: new Date(),
        yinHangMingCheng: '',
        zhangHao: '',
        zhuangTai: 0,
        ziZhiZhengZhao: '',
        }
}
// 弹窗确定
const enterDialog = async () => {
     elFormRef.value?.validate( async (valid) => {
             if (!valid) return
              let res
              switch (type.value) {
                case 'create':
                  res = await createFenBaoShangBiao(formData.value)
                  break
                case 'update':
                  res = await updateFenBaoShangBiao(formData.value)
                  break
                default:
                  res = await createFenBaoShangBiao(formData.value)
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
