<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" ref="elFormRef" label-position="right" :rules="rule" label-width="80px">
        <el-form-item label="备注:" prop="beiZhu">
          <el-input v-model="formData.beiZhu" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="所属:" prop="group">
          <el-input v-model="formData.group" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="会费:" prop="huiFei">
          <el-input v-model="formData.huiFei" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="会员名称:" prop="huiYuanMingCheng">
          <el-input v-model="formData.huiYuanMingCheng" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="级别:" prop="jiBie">
          <el-input v-model="formData.jiBie" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="缴费时间:" prop="jiaoFeiShiJian">
          <el-date-picker v-model="formData.jiaoFeiShiJian" type="date" placeholder="选择日期" :clearable="true"></el-date-picker>
        </el-form-item>
        <el-form-item label="结束有效期:" prop="jieShuYouXiaoQi">
          <el-date-picker v-model="formData.jieShuYouXiaoQi" type="date" placeholder="选择日期" :clearable="true"></el-date-picker>
        </el-form-item>
        <el-form-item label="开始有效期:" prop="kaiShiYouXiaoQi">
          <el-date-picker v-model="formData.kaiShiYouXiaoQi" type="date" placeholder="选择日期" :clearable="true"></el-date-picker>
        </el-form-item>
        <el-form-item label="入会联系:" prop="ruHuiLianXi">
          <el-input v-model="formData.ruHuiLianXi" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="网址:" prop="wangZhi">
          <el-input v-model="formData.wangZhi" :clearable="true" placeholder="请输入" />
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
  name: 'HuiYuanGuanLi'
}
</script>

<script setup>
import {
  createHuiYuanGuanLi,
  updateHuiYuanGuanLi,
  findHuiYuanGuanLi
} from '@/api/huiYuanGuanLi'

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
            group: '',
            huiFei: '',
            huiYuanMingCheng: '',
            jiBie: '',
            jiaoFeiShiJian: new Date(),
            jieShuYouXiaoQi: new Date(),
            kaiShiYouXiaoQi: new Date(),
            ruHuiLianXi: '',
            wangZhi: '',
        })
// 验证规则
const rule = reactive({
})

const elFormRef = ref()

// 初始化方法
const init = async () => {
 // 建议通过url传参获取目标数据ID 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
    if (route.query.id) {
      const res = await findHuiYuanGuanLi({ ID: route.query.id })
      if (res.code === 0) {
        formData.value = res.data.rehuiYuanGuanLi
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
               res = await createHuiYuanGuanLi(formData.value)
               break
             case 'update':
               res = await updateHuiYuanGuanLi(formData.value)
               break
             default:
               res = await createHuiYuanGuanLi(formData.value)
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
