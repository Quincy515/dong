<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" ref="elFormRef" label-position="right" :rules="rule" label-width="80px">
        <el-form-item label="备注:" prop="beiZhu">
          <el-input v-model="formData.beiZhu" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="公章名称:" prop="gongZhangMingCheng">
          <el-input v-model="formData.gongZhangMingCheng" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="所属:" prop="group">
          <el-input v-model="formData.group" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="经手人:" prop="jingShouRen">
          <el-input v-model="formData.jingShouRen" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="经手时间:" prop="jingShouShiJian">
          <el-date-picker v-model="formData.jingShouShiJian" type="date" placeholder="选择日期" :clearable="true"></el-date-picker>
        </el-form-item>
        <el-form-item label="申请人:" prop="shenQingRen">
          <el-input v-model="formData.shenQingRen" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="申请时间:" prop="shenQingShiJian">
          <el-date-picker v-model="formData.shenQingShiJian" type="date" placeholder="选择日期" :clearable="true"></el-date-picker>
        </el-form-item>
        <el-form-item label="事项:" prop="shiXiang">
          <el-input v-model="formData.shiXiang" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="数量:" prop="shuLiang">
          <el-input v-model.number="formData.shuLiang" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="详情:" prop="xiangQing">
          <el-input v-model="formData.xiangQing" :clearable="true" placeholder="请输入" />
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
  name: 'GongZhangDengJi'
}
</script>

<script setup>
import {
  createGongZhangDengJi,
  updateGongZhangDengJi,
  findGongZhangDengJi
} from '@/api/gongZhangDengJi'

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
            gongZhangMingCheng: '',
            group: '',
            jingShouRen: '',
            jingShouShiJian: new Date(),
            shenQingRen: '',
            shenQingShiJian: new Date(),
            shiXiang: '',
            shuLiang: 0,
            xiangQing: '',
        })
// 验证规则
const rule = reactive({
})

const elFormRef = ref()

// 初始化方法
const init = async () => {
 // 建议通过url传参获取目标数据ID 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
    if (route.query.id) {
      const res = await findGongZhangDengJi({ ID: route.query.id })
      if (res.code === 0) {
        formData.value = res.data.regongZhangDengJi
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
               res = await createGongZhangDengJi(formData.value)
               break
             case 'update':
               res = await updateGongZhangDengJi(formData.value)
               break
             default:
               res = await createGongZhangDengJi(formData.value)
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
