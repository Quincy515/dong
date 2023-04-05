<template>
  <div>
    <div class="gva-search-box">
      <el-form :inline="true" :model="searchInfo" class="demo-form-inline">
      <el-form-item label="创建时间">
      <el-date-picker v-model="searchInfo.startCreatedAt" type="datetime" placeholder="开始时间"></el-date-picker>
       —
      <el-date-picker v-model="searchInfo.endCreatedAt" type="datetime" placeholder="结束时间"></el-date-picker>
      </el-form-item>
        <el-form-item label="项目编号">
            
             <el-input v-model.number="searchInfo.xiangMuBianHao" placeholder="搜索条件" />

        </el-form-item>
           <el-form-item label="保障金状态" prop="baoZhangJinZhuangTai">
            <el-select v-model="searchInfo.baoZhangJinZhuangTai" clearable placeholder="请选择" @clear="()=>{searchInfo.baoZhangJinZhuangTai=undefined}">
              <el-option v-for="(item,key) in baoZhangJinZhuangTaiOptions" :key="key" :label="item.label" :value="item.value" />
            </el-select>
            </el-form-item>
           <el-form-item label="保证金状态" prop="baoZhengJinZhuangTai">
            <el-select v-model="searchInfo.baoZhengJinZhuangTai" clearable placeholder="请选择" @clear="()=>{searchInfo.baoZhengJinZhuangTai=undefined}">
              <el-option v-for="(item,key) in baoZhengJinZhuangTaiOptions" :key="key" :label="item.label" :value="item.value" />
            </el-select>
            </el-form-item>
           <el-form-item label="付款方式" prop="fuKuanFangShi">
            <el-select v-model="searchInfo.fuKuanFangShi" clearable placeholder="请选择" @clear="()=>{searchInfo.fuKuanFangShi=undefined}">
              <el-option v-for="(item,key) in fuKuanFangShiOptions" :key="key" :label="item.label" :value="item.value" />
            </el-select>
            </el-form-item>
           <el-form-item label="结算状态 1.未结算 2.结算完成" prop="jieSuanZhuangTai">
            <el-select v-model="searchInfo.jieSuanZhuangTai" clearable placeholder="请选择" @clear="()=>{searchInfo.jieSuanZhuangTai=undefined}">
              <el-option v-for="(item,key) in jieSuanZhuangTaiOptions" :key="key" :label="item.label" :value="item.value" />
            </el-select>
            </el-form-item>
           <el-form-item label="项目状态 1.建设中 2.未完成 3.竣工" prop="xiangMuZhuangTai">
            <el-select v-model="searchInfo.xiangMuZhuangTai" clearable placeholder="请选择" @clear="()=>{searchInfo.xiangMuZhuangTai=undefined}">
              <el-option v-for="(item,key) in xiangMuZhuangTaiOptions" :key="key" :label="item.label" :value="item.value" />
            </el-select>
            </el-form-item>
        <el-form-item label="项目经理">
            
             <el-input v-model.number="searchInfo.xiangMuJingLi" placeholder="搜索条件" />

        </el-form-item>
        <el-form-item label="项目主管">
            
             <el-input v-model.number="searchInfo.xiangMuZhuGuan" placeholder="搜索条件" />

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
        <el-table-column align="left" label="所属" prop="group" width="120" />
        <el-table-column align="left" label="项目编号" prop="xiangMuBianHao" width="120" />
         <el-table-column align="left" label="计划开工日期" width="180">
            <template #default="scope">{{ formatDate(scope.row.jiHuaKaiGongRiQi) }}</template>
         </el-table-column>
         <el-table-column align="left" label="计划结束日期" width="180">
            <template #default="scope">{{ formatDate(scope.row.jiHuaJieShuRiQi) }}</template>
         </el-table-column>
        <el-table-column align="left" label="工程造价" prop="gongChengZaoJia" width="120" />
        <el-table-column align="left" label="农民工工资保障金" prop="nongMinGongGongZiBaoZhangJin" width="120" />
        <el-table-column align="left" label="保障金状态" prop="baoZhangJinZhuangTai" width="120">
            <template #default="scope">
            {{ filterDict(scope.row.baoZhangJinZhuangTai,baoZhangJinZhuangTaiOptions) }}
            </template>
        </el-table-column>
        <el-table-column align="left" label="履约保证金" prop="luYueBaoZhengJin" width="120" />
        <el-table-column align="left" label="保证金状态" prop="baoZhengJinZhuangTai" width="120">
            <template #default="scope">
            {{ filterDict(scope.row.baoZhengJinZhuangTai,baoZhengJinZhuangTaiOptions) }}
            </template>
        </el-table-column>
        <el-table-column align="left" label="工作保险金额" prop="gongZuoBaoXianJinE" width="120" />
        <el-table-column align="left" label="付款方式" prop="fuKuanFangShi" width="120">
            <template #default="scope">
            {{ filterDict(scope.row.fuKuanFangShi,fuKuanFangShiOptions) }}
            </template>
        </el-table-column>
        <el-table-column align="left" label="预期利润" prop="yuQiLiRun" width="120" />
         <el-table-column align="left" label="实际开工日期" width="180">
            <template #default="scope">{{ formatDate(scope.row.shiJiKaiGongRiQi) }}</template>
         </el-table-column>
         <el-table-column align="left" label="实际结束日期" width="180">
            <template #default="scope">{{ formatDate(scope.row.shiJiJieShuRiQi) }}</template>
         </el-table-column>
        <el-table-column align="left" label="结算状态 1.未结算 2.结算完成" prop="jieSuanZhuangTai" width="120">
            <template #default="scope">
            {{ filterDict(scope.row.jieSuanZhuangTai,jieSuanZhuangTaiOptions) }}
            </template>
        </el-table-column>
        <el-table-column align="left" label="项目状态 1.建设中 2.未完成 3.竣工" prop="xiangMuZhuangTai" width="120">
            <template #default="scope">
            {{ filterDict(scope.row.xiangMuZhuangTai,xiangMuZhuangTaiOptions) }}
            </template>
        </el-table-column>
        <el-table-column align="left" label="项目经理" prop="xiangMuJingLi" width="120" />
        <el-table-column align="left" label="项目主管" prop="xiangMuZhuGuan" width="120" />
        <el-table-column align="left" label="备注" prop="beiZhu" width="120" />
        <el-table-column align="left" label="按钮组">
            <template #default="scope">
            <el-button type="primary" link icon="edit" size="small" class="table-button" @click="updateXiangMuBiaoFunc(scope.row)">变更</el-button>
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
        <el-form-item label="所属:"  prop="group" >
          <el-input v-model="formData.group" :clearable="true"  placeholder="请输入" />
        </el-form-item>
        <el-form-item label="项目编号:"  prop="xiangMuBianHao" >
          <el-input v-model.number="formData.xiangMuBianHao" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="计划开工日期:"  prop="jiHuaKaiGongRiQi" >
          <el-date-picker v-model="formData.jiHuaKaiGongRiQi" type="date" style="width:100%" placeholder="选择日期" :clearable="true"  />
        </el-form-item>
        <el-form-item label="计划结束日期:"  prop="jiHuaJieShuRiQi" >
          <el-date-picker v-model="formData.jiHuaJieShuRiQi" type="date" style="width:100%" placeholder="选择日期" :clearable="true"  />
        </el-form-item>
        <el-form-item label="工程造价:"  prop="gongChengZaoJia" >
          <el-input-number v-model="formData.gongChengZaoJia"  style="width:100%" :precision="2" :clearable="true"  />
        </el-form-item>
        <el-form-item label="农民工工资保障金:"  prop="nongMinGongGongZiBaoZhangJin" >
          <el-input-number v-model="formData.nongMinGongGongZiBaoZhangJin"  style="width:100%" :precision="2" :clearable="true"  />
        </el-form-item>
        <el-form-item label="保障金状态:"  prop="baoZhangJinZhuangTai" >
          <el-select v-model="formData.baoZhangJinZhuangTai" placeholder="请选择" style="width:100%" :clearable="true" >
            <el-option v-for="(item,key) in baoZhangJinZhuangTaiOptions" :key="key" :label="item.label" :value="item.value" />
          </el-select>
        </el-form-item>
        <el-form-item label="履约保证金:"  prop="luYueBaoZhengJin" >
          <el-input-number v-model="formData.luYueBaoZhengJin"  style="width:100%" :precision="2" :clearable="true"  />
        </el-form-item>
        <el-form-item label="保证金状态:"  prop="baoZhengJinZhuangTai" >
          <el-select v-model="formData.baoZhengJinZhuangTai" placeholder="请选择" style="width:100%" :clearable="true" >
            <el-option v-for="(item,key) in baoZhengJinZhuangTaiOptions" :key="key" :label="item.label" :value="item.value" />
          </el-select>
        </el-form-item>
        <el-form-item label="工作保险金额:"  prop="gongZuoBaoXianJinE" >
          <el-input-number v-model="formData.gongZuoBaoXianJinE"  style="width:100%" :precision="2" :clearable="true"  />
        </el-form-item>
        <el-form-item label="付款方式:"  prop="fuKuanFangShi" >
          <el-select v-model="formData.fuKuanFangShi" placeholder="请选择" style="width:100%" :clearable="true" >
            <el-option v-for="(item,key) in fuKuanFangShiOptions" :key="key" :label="item.label" :value="item.value" />
          </el-select>
        </el-form-item>
        <el-form-item label="预期利润:"  prop="yuQiLiRun" >
          <el-input-number v-model="formData.yuQiLiRun"  style="width:100%" :precision="2" :clearable="true"  />
        </el-form-item>
        <el-form-item label="实际开工日期:"  prop="shiJiKaiGongRiQi" >
          <el-date-picker v-model="formData.shiJiKaiGongRiQi" type="date" style="width:100%" placeholder="选择日期" :clearable="true"  />
        </el-form-item>
        <el-form-item label="实际结束日期:"  prop="shiJiJieShuRiQi" >
          <el-date-picker v-model="formData.shiJiJieShuRiQi" type="date" style="width:100%" placeholder="选择日期" :clearable="true"  />
        </el-form-item>
        <el-form-item label="结算状态 1.未结算 2.结算完成:"  prop="jieSuanZhuangTai" >
          <el-select v-model="formData.jieSuanZhuangTai" placeholder="请选择" style="width:100%" :clearable="true" >
            <el-option v-for="(item,key) in jieSuanZhuangTaiOptions" :key="key" :label="item.label" :value="item.value" />
          </el-select>
        </el-form-item>
        <el-form-item label="项目状态 1.建设中 2.未完成 3.竣工:"  prop="xiangMuZhuangTai" >
          <el-select v-model="formData.xiangMuZhuangTai" placeholder="请选择" style="width:100%" :clearable="true" >
            <el-option v-for="(item,key) in xiangMuZhuangTaiOptions" :key="key" :label="item.label" :value="item.value" />
          </el-select>
        </el-form-item>
        <el-form-item label="项目经理:"  prop="xiangMuJingLi" >
          <el-input v-model.number="formData.xiangMuJingLi" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="项目主管:"  prop="xiangMuZhuGuan" >
          <el-input v-model.number="formData.xiangMuZhuGuan" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="备注:"  prop="beiZhu" >
          <el-input v-model="formData.beiZhu" :clearable="true"  placeholder="请输入" />
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
  name: 'XiangMuBiao'
}
</script>

<script setup>
import {
  createXiangMuBiao,
  deleteXiangMuBiao,
  deleteXiangMuBiaoByIds,
  updateXiangMuBiao,
  findXiangMuBiao,
  getXiangMuBiaoList
} from '@/api/xiangMuBiao'

// 全量引入格式化工具 请按需保留
import { getDictFunc, formatDate, formatBoolean, filterDict } from '@/utils/format'
import { ElMessage, ElMessageBox } from 'element-plus'
import { ref, reactive } from 'vue'

// 自动化生成的字典（可能为空）以及字段
const jieSuanZhuangTaiOptions = ref([])
const xiangMuZhuangTaiOptions = ref([])
const baoZhangJinZhuangTaiOptions = ref([])
const baoZhengJinZhuangTaiOptions = ref([])
const fuKuanFangShiOptions = ref([])
const formData = ref({
        group: '',
        xiangMuBianHao: 0,
        jiHuaKaiGongRiQi: new Date(),
        jiHuaJieShuRiQi: new Date(),
        gongChengZaoJia: 0,
        nongMinGongGongZiBaoZhangJin: 0,
        baoZhangJinZhuangTai: undefined,
        luYueBaoZhengJin: 0,
        baoZhengJinZhuangTai: undefined,
        gongZuoBaoXianJinE: 0,
        fuKuanFangShi: undefined,
        yuQiLiRun: 0,
        shiJiKaiGongRiQi: new Date(),
        shiJiJieShuRiQi: new Date(),
        jieSuanZhuangTai: undefined,
        xiangMuZhuangTai: undefined,
        xiangMuJingLi: 0,
        xiangMuZhuGuan: 0,
        beiZhu: '',
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
  const table = await getXiangMuBiaoList({ page: page.value, pageSize: pageSize.value, ...searchInfo.value })
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
    jieSuanZhuangTaiOptions.value = await getDictFunc('jieSuanZhuangTai')
    xiangMuZhuangTaiOptions.value = await getDictFunc('xiangMuZhuangTai')
    baoZhangJinZhuangTaiOptions.value = await getDictFunc('baoZhangJinZhuangTai')
    baoZhengJinZhuangTaiOptions.value = await getDictFunc('baoZhengJinZhuangTai')
    fuKuanFangShiOptions.value = await getDictFunc('fuKuanFangShi')
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
            deleteXiangMuBiaoFunc(row)
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
      const res = await deleteXiangMuBiaoByIds({ ids })
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
const updateXiangMuBiaoFunc = async(row) => {
    const res = await findXiangMuBiao({ ID: row.ID })
    type.value = 'update'
    if (res.code === 0) {
        formData.value = res.data.rexiangMuBiao
        dialogFormVisible.value = true
    }
}


// 删除行
const deleteXiangMuBiaoFunc = async (row) => {
    const res = await deleteXiangMuBiao({ ID: row.ID })
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
        group: '',
        xiangMuBianHao: 0,
        jiHuaKaiGongRiQi: new Date(),
        jiHuaJieShuRiQi: new Date(),
        gongChengZaoJia: 0,
        nongMinGongGongZiBaoZhangJin: 0,
        baoZhangJinZhuangTai: undefined,
        luYueBaoZhengJin: 0,
        baoZhengJinZhuangTai: undefined,
        gongZuoBaoXianJinE: 0,
        fuKuanFangShi: undefined,
        yuQiLiRun: 0,
        shiJiKaiGongRiQi: new Date(),
        shiJiJieShuRiQi: new Date(),
        jieSuanZhuangTai: undefined,
        xiangMuZhuangTai: undefined,
        xiangMuJingLi: 0,
        xiangMuZhuGuan: 0,
        beiZhu: '',
        }
}
// 弹窗确定
const enterDialog = async () => {
     elFormRef.value?.validate( async (valid) => {
             if (!valid) return
              let res
              switch (type.value) {
                case 'create':
                  res = await createXiangMuBiao(formData.value)
                  break
                case 'update':
                  res = await updateXiangMuBiao(formData.value)
                  break
                default:
                  res = await createXiangMuBiao(formData.value)
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
