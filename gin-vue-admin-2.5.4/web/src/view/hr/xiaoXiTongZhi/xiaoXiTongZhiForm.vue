<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" ref="elFormRef" label-position="right" :rules="rule" label-width="80px">
        <el-form-item label="备注:" prop="beiZhu">
          <el-input v-model="formData.beiZhu" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="内容:" prop="content">
          <el-input v-model="formData.content" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="所属:" prop="group">
          <el-input v-model="formData.group" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="接收方 id:" prop="receiverId">
          <el-input v-model.number="formData.receiverId" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="跳转地址:" prop="redirectUrl">
          <el-input v-model="formData.redirectUrl" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="发送方 id:" prop="senderId">
          <el-input v-model.number="formData.senderId" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="状态 0.未读 1.已读:" prop="status">
          <el-input v-model.number="formData.status" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="标题:" prop="title">
          <el-input v-model="formData.title" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="类型:" prop="type">
          <el-input v-model="formData.type" :clearable="true" placeholder="请输入" />
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
  name: 'XiaoXiTongZhi'
}
</script>

<script setup>
import {
  createXiaoXiTongZhi,
  updateXiaoXiTongZhi,
  findXiaoXiTongZhi
} from '@/api/xiaoXiTongZhi'

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
            content: '',
            group: '',
            receiverId: 0,
            redirectUrl: '',
            senderId: 0,
            status: 0,
            title: '',
            type: '',
        })
// 验证规则
const rule = reactive({
})

const elFormRef = ref()

// 初始化方法
const init = async () => {
 // 建议通过url传参获取目标数据ID 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
    if (route.query.id) {
      const res = await findXiaoXiTongZhi({ ID: route.query.id })
      if (res.code === 0) {
        formData.value = res.data.rexiaoXiTongZhi
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
               res = await createXiaoXiTongZhi(formData.value)
               break
             case 'update':
               res = await updateXiaoXiTongZhi(formData.value)
               break
             default:
               res = await createXiaoXiTongZhi(formData.value)
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
