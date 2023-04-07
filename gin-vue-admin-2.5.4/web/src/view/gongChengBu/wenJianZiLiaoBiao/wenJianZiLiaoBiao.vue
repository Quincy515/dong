<template>
  <div>
    <div class="gva-search-box">
      <el-form :inline="true" :model="searchInfo" class="demo-form-inline">
      <el-form-item label="创建时间">
      <el-date-picker v-model="searchInfo.startCreatedAt" type="datetime" placeholder="开始时间"></el-date-picker>
       —
      <el-date-picker v-model="searchInfo.endCreatedAt" type="datetime" placeholder="结束时间"></el-date-picker>
      </el-form-item>
        <el-form-item label="节点ID只保存最近的节点">
            
             <el-input v-model.number="searchInfo.jieDianId" placeholder="搜索条件" />

        </el-form-item>
        <el-form-item label="是否审核">
            
             <el-input v-model.number="searchInfo.shiFouShenHe" placeholder="搜索条件" />

        </el-form-item>
        <el-form-item label="文件id">
            
             <el-input v-model.number="searchInfo.wenJianId" placeholder="搜索条件" />

        </el-form-item>
        <el-form-item label="文件名称">
         <el-input v-model="searchInfo.wenJianMingCheng" placeholder="搜索条件" />

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
        <el-table-column align="left" label="所属" prop="group" width="120" />
        <el-table-column align="left" label="节点ID只保存最近的节点" prop="jieDianId" width="120" />
        <el-table-column align="left" label="上传人" prop="shangChuanRen" width="120" />
         <el-table-column align="left" label="上传日期" width="180">
            <template #default="scope">{{ formatDate(scope.row.shangChuanRiQi) }}</template>
         </el-table-column>
        <el-table-column align="left" label="审核人" prop="shenHeRen" width="120" />
         <el-table-column align="left" label="审核日期" width="180">
            <template #default="scope">{{ formatDate(scope.row.shenHeRiQi) }}</template>
         </el-table-column>
        <el-table-column align="left" label="是否审核" prop="shiFouShenHe" width="120" />
        <el-table-column align="left" label="文件id" prop="wenJianId" width="120" />
        <el-table-column align="left" label="文件类型" prop="wenJianLeiXing" width="120" />
        <el-table-column align="left" label="文件名称" prop="wenJianMingCheng" width="120" />
        <el-table-column align="left" label="阅读次数" prop="yueDuCiShu" width="120" />
        <el-table-column align="left" label="状态" prop="zhuangTai" width="120" />
        <el-table-column align="left" label="按钮组">
            <template #default="scope">
            <el-button type="primary" link icon="edit" size="small" class="table-button" @click="updateWenJianZiLiaoBiaoFunc(scope.row)">变更</el-button>
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
        <el-form-item label="所属:"  prop="group" >
          <el-input v-model="formData.group" :clearable="true"  placeholder="请输入" />
        </el-form-item>
        <el-form-item label="节点ID只保存最近的节点:"  prop="jieDianId" >
          <el-input v-model.number="formData.jieDianId" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="上传人:"  prop="shangChuanRen" >
          <el-input v-model.number="formData.shangChuanRen" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="上传日期:"  prop="shangChuanRiQi" >
          <el-date-picker v-model="formData.shangChuanRiQi" type="date" style="width:100%" placeholder="选择日期" :clearable="true"  />
        </el-form-item>
        <el-form-item label="审核人:"  prop="shenHeRen" >
          <el-input v-model.number="formData.shenHeRen" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="审核日期:"  prop="shenHeRiQi" >
          <el-date-picker v-model="formData.shenHeRiQi" type="date" style="width:100%" placeholder="选择日期" :clearable="true"  />
        </el-form-item>
        <el-form-item label="是否审核:"  prop="shiFouShenHe" >
          <el-input v-model.number="formData.shiFouShenHe" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="文件id:"  prop="wenJianId" >
          <el-input v-model.number="formData.wenJianId" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="文件类型:"  prop="wenJianLeiXing" >
          <el-input v-model.number="formData.wenJianLeiXing" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="文件名称:"  prop="wenJianMingCheng" >
          <el-input v-model="formData.wenJianMingCheng" :clearable="true"  placeholder="请输入" />
        </el-form-item>
        <el-form-item label="阅读次数:"  prop="yueDuCiShu" >
          <el-input v-model.number="formData.yueDuCiShu" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="状态:"  prop="zhuangTai" >
          <el-input v-model.number="formData.zhuangTai" :clearable="true" placeholder="请输入" />
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
  name: 'WenJianZiLiaoBiao'
}
</script>

<script setup>
import {
  createWenJianZiLiaoBiao,
  deleteWenJianZiLiaoBiao,
  deleteWenJianZiLiaoBiaoByIds,
  updateWenJianZiLiaoBiao,
  findWenJianZiLiaoBiao,
  getWenJianZiLiaoBiaoList
} from '@/api/wenJianZiLiaoBiao'

// 全量引入格式化工具 请按需保留
import { getDictFunc, formatDate, formatBoolean, filterDict } from '@/utils/format'
import { ElMessage, ElMessageBox } from 'element-plus'
import { ref, reactive } from 'vue'

// 自动化生成的字典（可能为空）以及字段
const formData = ref({
        beiZhu: '',
        group: '',
        jieDianId: 0,
        shangChuanRen: 0,
        shangChuanRiQi: new Date(),
        shenHeRen: 0,
        shenHeRiQi: new Date(),
        shiFouShenHe: 0,
        wenJianId: 0,
        wenJianLeiXing: 0,
        wenJianMingCheng: '',
        yueDuCiShu: 0,
        zhuangTai: 0,
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
  const table = await getWenJianZiLiaoBiaoList({ page: page.value, pageSize: pageSize.value, ...searchInfo.value })
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
            deleteWenJianZiLiaoBiaoFunc(row)
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
      const res = await deleteWenJianZiLiaoBiaoByIds({ ids })
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
const updateWenJianZiLiaoBiaoFunc = async(row) => {
    const res = await findWenJianZiLiaoBiao({ ID: row.ID })
    type.value = 'update'
    if (res.code === 0) {
        formData.value = res.data.rewenJianZiLiaoBiao
        dialogFormVisible.value = true
    }
}


// 删除行
const deleteWenJianZiLiaoBiaoFunc = async (row) => {
    const res = await deleteWenJianZiLiaoBiao({ ID: row.ID })
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
        group: '',
        jieDianId: 0,
        shangChuanRen: 0,
        shangChuanRiQi: new Date(),
        shenHeRen: 0,
        shenHeRiQi: new Date(),
        shiFouShenHe: 0,
        wenJianId: 0,
        wenJianLeiXing: 0,
        wenJianMingCheng: '',
        yueDuCiShu: 0,
        zhuangTai: 0,
        }
}
// 弹窗确定
const enterDialog = async () => {
     elFormRef.value?.validate( async (valid) => {
             if (!valid) return
              let res
              switch (type.value) {
                case 'create':
                  res = await createWenJianZiLiaoBiao(formData.value)
                  break
                case 'update':
                  res = await updateWenJianZiLiaoBiao(formData.value)
                  break
                default:
                  res = await createWenJianZiLiaoBiao(formData.value)
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
