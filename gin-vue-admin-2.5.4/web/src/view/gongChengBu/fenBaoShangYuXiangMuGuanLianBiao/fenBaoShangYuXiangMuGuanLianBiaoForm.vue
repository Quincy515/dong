<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" ref="elFormRef" label-position="right" :rules="rule" label-width="80px">
        <el-form-item label="备注:" prop="beiZhu">
          <el-input v-model="formData.beiZhu" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="分包商ID:" prop="fenBaoShangId">
          <el-input v-model.number="formData.fenBaoShangId" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="所属:" prop="group">
          <el-input v-model="formData.group" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="合同编号:" prop="heTongBianHao">
          <el-input v-model="formData.heTongBianHao" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="合同名称:" prop="heTongMingCheng">
          <el-input v-model="formData.heTongMingCheng" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="进场日期:" prop="jinChangRiQi">
          <el-date-picker v-model="formData.jinChangRiQi" type="date" placeholder="选择日期" :clearable="true"></el-date-picker>
        </el-form-item>
        <el-form-item label="退场日期:" prop="tuiChangRiQi">
          <el-date-picker v-model="formData.tuiChangRiQi" type="date" placeholder="选择日期" :clearable="true"></el-date-picker>
        </el-form-item>
        <el-form-item label="委托人:" prop="weiTuoRen">
          <el-input v-model="formData.weiTuoRen" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="委托人手机:" prop="weiTuoRenShouJi">
          <el-input v-model="formData.weiTuoRenShouJi" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="项目ID:" prop="xiangMuId">
          <el-input v-model.number="formData.xiangMuId" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="状态:" prop="zhuangTai">
          <el-input v-model.number="formData.zhuangTai" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item>
          <el-button size="small" type="primary" @click="save">保存</el-button>
          <el-button size="small" type="primary" @click="back">返回</el-button>
        </el-form-item>
      </el-form>
    </div>
  </div>
</template>

<script>
export default {
  name: 'FenBaoShangYuXiangMuGuanLianBiao'
}
</script>

<script setup>
import {
  createFenBaoShangYuXiangMuGuanLianBiao,
  updateFenBaoShangYuXiangMuGuanLianBiao,
  findFenBaoShangYuXiangMuGuanLianBiao
} from '@/api/fenBaoShangYuXiangMuGuanLianBiao'

// 自动获取字典
import { getDictFunc } from '@/utils/format'
import { useRoute, useRouter } from "vue-router"
import { ElMessage } from 'element-plus'
import { ref, reactive } from 'vue'
const route = useRoute()
const router = useRouter()

const type = ref('')
const formData = ref({
            beiZhu: '',
            fenBaoShangId: 0,
            group: '',
            heTongBianHao: '',
            heTongMingCheng: '',
            jinChangRiQi: new Date(),
            tuiChangRiQi: new Date(),
            weiTuoRen: '',
            weiTuoRenShouJi: '',
            xiangMuId: 0,
            zhuangTai: 0,
        })
// 验证规则
const rule = reactive({
})

const elFormRef = ref()

// 初始化方法
const init = async () => {
 // 建议通过url传参获取目标数据ID 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
    if (route.query.id) {
      const res = await findFenBaoShangYuXiangMuGuanLianBiao({ ID: route.query.id })
      if (res.code === 0) {
        formData.value = res.data.refenBaoShangYuXiangMuGuanLianBiao
        type.value = 'update'
      }
    } else {
      type.value = 'create'
    }
}

init()
// 保存按钮
const save = async() => {
      elFormRef.value?.validate( async (valid) => {
         if (!valid) return
            let res
           switch (type.value) {
             case 'create':
               res = await createFenBaoShangYuXiangMuGuanLianBiao(formData.value)
               break
             case 'update':
               res = await updateFenBaoShangYuXiangMuGuanLianBiao(formData.value)
               break
             default:
               res = await createFenBaoShangYuXiangMuGuanLianBiao(formData.value)
               break
           }
           if (res.code === 0) {
             ElMessage({
               type: 'success',
               message: '创建/更改成功'
             })
           }
       })
}

// 返回按钮
const back = () => {
    router.go(-1)
}

</script>

<style>
</style>
