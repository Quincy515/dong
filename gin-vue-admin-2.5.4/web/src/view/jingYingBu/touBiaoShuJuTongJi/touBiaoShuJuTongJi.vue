<template>
  <div>
    <div class="gva-search-box">
      <el-form :inline="true" :model="searchInfo" class="demo-form-inline">
      <el-form-item label="创建时间">
      <el-date-picker v-model="searchInfo.startCreatedAt" type="datetime" placeholder="开始时间"></el-date-picker>
       —
      <el-date-picker v-model="searchInfo.endCreatedAt" type="datetime" placeholder="结束时间"></el-date-picker>
      </el-form-item>
           <el-form-item label="等级要求" prop="dengJiYaoQiu">
            <el-select v-model="searchInfo.dengJiYaoQiu" clearable placeholder="请选择" @clear="()=>{searchInfo.dengJiYaoQiu=undefined}">
              <el-option v-for="(item,key) in deng_ji_yao_qiuOptions" :key="key" :label="item.label" :value="item.value" />
            </el-select>
            </el-form-item>
           <el-form-item label="所属地区" prop="suoShuDiQu">
            <el-select v-model="searchInfo.suoShuDiQu" clearable placeholder="请选择" @clear="()=>{searchInfo.suoShuDiQu=undefined}">
              <el-option v-for="(item,key) in suo_shu_di_quOptions" :key="key" :label="item.label" :value="item.value" />
            </el-select>
            </el-form-item>
        <el-form-item label="项目名称">
         <el-input v-model="searchInfo.xiangMuMingCheng" placeholder="搜索条件" />

        </el-form-item>
           <el-form-item label="资质要求" prop="ziZhiYaoQiu">
            <el-select v-model="searchInfo.ziZhiYaoQiu" clearable placeholder="请选择" @clear="()=>{searchInfo.ziZhiYaoQiu=undefined}">
              <el-option v-for="(item,key) in zi_zhi_yao_qiuOptions" :key="key" :label="item.label" :value="item.value" />
            </el-select>
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
        <el-table-column align="left" label="等级要求" prop="dengJiYaoQiu" width="120">
            <template #default="scope">
            {{ filterDict(scope.row.dengJiYaoQiu,deng_ji_yao_qiuOptions) }}
            </template>
        </el-table-column>
        <el-table-column align="left" label="第二报价" prop="diErBaoJia" width="120" />
        <el-table-column align="left" label="第二下浮率" prop="diErXiaFuLu" width="120" />
        <el-table-column align="left" label="第二中标候选人" prop="diErZhongBiaoHouXuanRen" width="120" />
        <el-table-column align="left" label="第三报价" prop="diSanBaoJia" width="120" />
        <el-table-column align="left" label="第三下浮率" prop="diSanXiaFuLu" width="120" />
        <el-table-column align="left" label="第三中标候选人" prop="diSanZhongBiaoHouXuanRen" width="120" />
        <el-table-column align="left" label="第一报价" prop="diYiBaoJia" width="120" />
        <el-table-column align="left" label="第一下浮率" prop="diYiXiaFuLu" width="120" />
        <el-table-column align="left" label="第一中标候选人" prop="diYiZhongBiaoHouXuanRen" width="120" />
        <el-table-column align="left" label="所属" prop="group" width="120" />
         <el-table-column align="left" label="开标时间" width="180">
            <template #default="scope">{{ formatDate(scope.row.kaiBiaoShiJian) }}</template>
         </el-table-column>
        <el-table-column align="left" label="所属地区" prop="suoShuDiQu" width="120">
            <template #default="scope">
            {{ filterDict(scope.row.suoShuDiQu,suo_shu_di_quOptions) }}
            </template>
        </el-table-column>
        <el-table-column align="left" label="投标单位名称" prop="touBiaoDanWeiMingCheng" width="120" />
        <el-table-column align="left" label="项目名称" prop="xiangMuMingCheng" width="120" />
        <el-table-column align="left" label="招标规模" prop="zhaoBiaoGuiMo" width="120" />
        <el-table-column align="left" label="资质要求" prop="ziZhiYaoQiu" width="120">
            <template #default="scope">
            {{ filterDict(scope.row.ziZhiYaoQiu,zi_zhi_yao_qiuOptions) }}
            </template>
        </el-table-column>
        <el-table-column align="left" label="按钮组">
            <template #default="scope">
            <el-button type="primary" link icon="edit" size="small" class="table-button" @click="updateTouBiaoShuJuTongJiFunc(scope.row)">变更</el-button>
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
        <el-form-item label="等级要求:"  prop="dengJiYaoQiu" >
          <el-select v-model="formData.dengJiYaoQiu" placeholder="请选择" style="width:100%" :clearable="true" >
            <el-option v-for="(item,key) in deng_ji_yao_qiuOptions" :key="key" :label="item.label" :value="item.value" />
          </el-select>
        </el-form-item>
        <el-form-item label="第二报价:"  prop="diErBaoJia" >
          <el-input v-model="formData.diErBaoJia" :clearable="true"  placeholder="请输入" />
        </el-form-item>
        <el-form-item label="第二下浮率:"  prop="diErXiaFuLu" >
          <el-input v-model="formData.diErXiaFuLu" :clearable="true"  placeholder="请输入" />
        </el-form-item>
        <el-form-item label="第二中标候选人:"  prop="diErZhongBiaoHouXuanRen" >
          <el-input v-model="formData.diErZhongBiaoHouXuanRen" :clearable="true"  placeholder="请输入" />
        </el-form-item>
        <el-form-item label="第三报价:"  prop="diSanBaoJia" >
          <el-input v-model="formData.diSanBaoJia" :clearable="true"  placeholder="请输入" />
        </el-form-item>
        <el-form-item label="第三下浮率:"  prop="diSanXiaFuLu" >
          <el-input v-model="formData.diSanXiaFuLu" :clearable="true"  placeholder="请输入" />
        </el-form-item>
        <el-form-item label="第三中标候选人:"  prop="diSanZhongBiaoHouXuanRen" >
          <el-input v-model="formData.diSanZhongBiaoHouXuanRen" :clearable="true"  placeholder="请输入" />
        </el-form-item>
        <el-form-item label="第一报价:"  prop="diYiBaoJia" >
          <el-input v-model="formData.diYiBaoJia" :clearable="true"  placeholder="请输入" />
        </el-form-item>
        <el-form-item label="第一下浮率:"  prop="diYiXiaFuLu" >
          <el-input v-model="formData.diYiXiaFuLu" :clearable="true"  placeholder="请输入" />
        </el-form-item>
        <el-form-item label="第一中标候选人:"  prop="diYiZhongBiaoHouXuanRen" >
          <el-input v-model="formData.diYiZhongBiaoHouXuanRen" :clearable="true"  placeholder="请输入" />
        </el-form-item>
        <el-form-item label="所属:"  prop="group" >
          <el-input v-model="formData.group" :clearable="true"  placeholder="请输入" />
        </el-form-item>
        <el-form-item label="开标时间:"  prop="kaiBiaoShiJian" >
          <el-date-picker v-model="formData.kaiBiaoShiJian" type="date" style="width:100%" placeholder="选择日期" :clearable="true"  />
        </el-form-item>
        <el-form-item label="所属地区:"  prop="suoShuDiQu" >
          <el-select v-model="formData.suoShuDiQu" placeholder="请选择" style="width:100%" :clearable="true" >
            <el-option v-for="(item,key) in suo_shu_di_quOptions" :key="key" :label="item.label" :value="item.value" />
          </el-select>
        </el-form-item>
        <el-form-item label="投标单位名称:"  prop="touBiaoDanWeiMingCheng" >
          <el-input v-model="formData.touBiaoDanWeiMingCheng" :clearable="true"  placeholder="请输入" />
        </el-form-item>
        <el-form-item label="项目名称:"  prop="xiangMuMingCheng" >
          <el-input v-model="formData.xiangMuMingCheng" :clearable="true"  placeholder="请输入" />
        </el-form-item>
        <el-form-item label="招标规模:"  prop="zhaoBiaoGuiMo" >
          <el-input v-model="formData.zhaoBiaoGuiMo" :clearable="true"  placeholder="请输入" />
        </el-form-item>
        <el-form-item label="资质要求:"  prop="ziZhiYaoQiu" >
          <el-select v-model="formData.ziZhiYaoQiu" placeholder="请选择" style="width:100%" :clearable="true" >
            <el-option v-for="(item,key) in zi_zhi_yao_qiuOptions" :key="key" :label="item.label" :value="item.value" />
          </el-select>
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
  name: 'TouBiaoShuJuTongJi'
}
</script>

<script setup>
import {
  createTouBiaoShuJuTongJi,
  deleteTouBiaoShuJuTongJi,
  deleteTouBiaoShuJuTongJiByIds,
  updateTouBiaoShuJuTongJi,
  findTouBiaoShuJuTongJi,
  getTouBiaoShuJuTongJiList
} from '@/api/touBiaoShuJuTongJi'

// 全量引入格式化工具 请按需保留
import { getDictFunc, formatDate, formatBoolean, filterDict } from '@/utils/format'
import { ElMessage, ElMessageBox } from 'element-plus'
import { ref, reactive } from 'vue'

// 自动化生成的字典（可能为空）以及字段
const deng_ji_yao_qiuOptions = ref([])
const suo_shu_di_quOptions = ref([])
const zi_zhi_yao_qiuOptions = ref([])
const formData = ref({
        beiZhu: '',
        dengJiYaoQiu: undefined,
        diErBaoJia: '',
        diErXiaFuLu: '',
        diErZhongBiaoHouXuanRen: '',
        diSanBaoJia: '',
        diSanXiaFuLu: '',
        diSanZhongBiaoHouXuanRen: '',
        diYiBaoJia: '',
        diYiXiaFuLu: '',
        diYiZhongBiaoHouXuanRen: '',
        group: '',
        kaiBiaoShiJian: new Date(),
        suoShuDiQu: undefined,
        touBiaoDanWeiMingCheng: '',
        xiangMuMingCheng: '',
        zhaoBiaoGuiMo: '',
        ziZhiYaoQiu: undefined,
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
  const table = await getTouBiaoShuJuTongJiList({ page: page.value, pageSize: pageSize.value, ...searchInfo.value })
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
    deng_ji_yao_qiuOptions.value = await getDictFunc('deng_ji_yao_qiu')
    suo_shu_di_quOptions.value = await getDictFunc('suo_shu_di_qu')
    zi_zhi_yao_qiuOptions.value = await getDictFunc('zi_zhi_yao_qiu')
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
            deleteTouBiaoShuJuTongJiFunc(row)
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
      const res = await deleteTouBiaoShuJuTongJiByIds({ ids })
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
const updateTouBiaoShuJuTongJiFunc = async(row) => {
    const res = await findTouBiaoShuJuTongJi({ ID: row.ID })
    type.value = 'update'
    if (res.code === 0) {
        formData.value = res.data.retouBiaoShuJuTongJi
        dialogFormVisible.value = true
    }
}


// 删除行
const deleteTouBiaoShuJuTongJiFunc = async (row) => {
    const res = await deleteTouBiaoShuJuTongJi({ ID: row.ID })
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
        dengJiYaoQiu: undefined,
        diErBaoJia: '',
        diErXiaFuLu: '',
        diErZhongBiaoHouXuanRen: '',
        diSanBaoJia: '',
        diSanXiaFuLu: '',
        diSanZhongBiaoHouXuanRen: '',
        diYiBaoJia: '',
        diYiXiaFuLu: '',
        diYiZhongBiaoHouXuanRen: '',
        group: '',
        kaiBiaoShiJian: new Date(),
        suoShuDiQu: undefined,
        touBiaoDanWeiMingCheng: '',
        xiangMuMingCheng: '',
        zhaoBiaoGuiMo: '',
        ziZhiYaoQiu: undefined,
        }
}
// 弹窗确定
const enterDialog = async () => {
     elFormRef.value?.validate( async (valid) => {
             if (!valid) return
              let res
              switch (type.value) {
                case 'create':
                  res = await createTouBiaoShuJuTongJi(formData.value)
                  break
                case 'update':
                  res = await updateTouBiaoShuJuTongJi(formData.value)
                  break
                default:
                  res = await createTouBiaoShuJuTongJi(formData.value)
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
