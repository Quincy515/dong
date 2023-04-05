<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" ref="elFormRef" label-position="right" :rules="rule" label-width="80px">
        <el-form-item label="备注:" prop="beiZhu">
          <el-input v-model="formData.beiZhu" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="法人:" prop="faRen">
          <el-input v-model="formData.faRen" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="法人手机:" prop="faRenShouJi">
          <el-input v-model="formData.faRenShouJi" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="分包商类型:" prop="fenBaoShangLeiXing">
          <el-select v-model="formData.fenBaoShangLeiXing" placeholder="请选择" :clearable="true">
            <el-option v-for="(item,key) in fenBaoShangLeiXingOptions" :key="key" :label="item.label" :value="item.value" />
          </el-select>
        </el-form-item>
        <el-form-item label="分包商名称:" prop="fenBaoShangMingCheng">
          <el-input v-model="formData.fenBaoShangMingCheng" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="所属:" prop="group">
          <el-input v-model="formData.group" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="开户行:" prop="kaiHuHang">
          <el-input v-model="formData.kaiHuHang" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="社会信用代码:" prop="sheHuiXinYongDaiMa">
          <el-input v-model="formData.sheHuiXinYongDaiMa" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="添加人:" prop="tianJiaRen">
          <el-input v-model.number="formData.tianJiaRen" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="添加时间:" prop="tianJiaShiJian">
          <el-date-picker v-model="formData.tianJiaShiJian" type="date" placeholder="选择日期" :clearable="true"></el-date-picker>
        </el-form-item>
        <el-form-item label="银行名称:" prop="yinHangMingCheng">
          <el-input v-model="formData.yinHangMingCheng" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="账号:" prop="zhangHao">
          <el-input v-model="formData.zhangHao" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="状态:" prop="zhuangTai">
          <el-input v-model.number="formData.zhuangTai" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="资质证照:" prop="ziZhiZhengZhao">
          <el-input v-model="formData.ziZhiZhengZhao" :clearable="true" placeholder="请输入" />
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
  name: 'FenBaoShangBiao'
}
</script>

<script setup>
import {
  createFenBaoShangBiao,
  updateFenBaoShangBiao,
  findFenBaoShangBiao
} from '@/api/fenBaoShangBiao'

// 自动获取字典
import { getDictFunc } from '@/utils/format'
import { useRoute, useRouter } from "vue-router"
import { ElMessage } from 'element-plus'
import { ref, reactive } from 'vue'
const route = useRoute()
const router = useRouter()

const type = ref('')
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

// 初始化方法
const init = async () => {
 // 建议通过url传参获取目标数据ID 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
    if (route.query.id) {
      const res = await findFenBaoShangBiao({ ID: route.query.id })
      if (res.code === 0) {
        formData.value = res.data.refenBaoShangBiao
        type.value = 'update'
      }
    } else {
      type.value = 'create'
    }
    fenBaoShangLeiXingOptions.value = await getDictFunc('fenBaoShangLeiXing')
}

init()
// 保存按钮
const save = async() => {
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
