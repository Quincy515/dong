<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" ref="elFormRef" label-position="right" :rules="rule" label-width="80px">
        <el-form-item label="所属:" prop="group">
          <el-input v-model="formData.group" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="证书名称:" prop="zhengShuMingCheng">
          <el-input v-model="formData.zhengShuMingCheng" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="证书编号:" prop="zhengShuBianHao">
          <el-input v-model="formData.zhengShuBianHao" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="证书类别:" prop="zhengShuLeiBie">
          <el-select v-model="formData.zhengShuLeiBie" placeholder="请选择" :clearable="true">
            <el-option v-for="(item,key) in zheng_shu_lei_bieOptions" :key="key" :label="item.label" :value="item.value" />
          </el-select>
        </el-form-item>
        <el-form-item label="所有人姓名:" prop="suoYouRenXingMing">
          <el-input v-model="formData.suoYouRenXingMing" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="身份证号:" prop="shenFenZhengHao">
          <el-input v-model="formData.shenFenZhengHao" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="手机号码:" prop="shouJiHaoMa">
          <el-input v-model="formData.shouJiHaoMa" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="发证部门:" prop="faZhengBuMen">
          <el-input v-model="formData.faZhengBuMen" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="证书所属城市:" prop="zhengShuSuoShuChengShi">
          <el-select v-model="formData.zhengShuSuoShuChengShi" placeholder="请选择" :clearable="true">
            <el-option v-for="(item,key) in zheng_shu_suo_shu_cheng_shiOptions" :key="key" :label="item.label" :value="item.value" />
          </el-select>
        </el-form-item>
        <el-form-item label="证书状态:" prop="zhengShuZhuangTai">
          <el-select v-model="formData.zhengShuZhuangTai" placeholder="请选择" :clearable="true">
            <el-option v-for="(item,key) in zheng_shu_zhuang_taiOptions" :key="key" :label="item.label" :value="item.value" />
          </el-select>
        </el-form-item>
        <el-form-item label="证书到期日期:" prop="zhengShuDaoQiRiQi">
          <el-date-picker v-model="formData.zhengShuDaoQiRiQi" type="date" placeholder="选择日期" :clearable="true"></el-date-picker>
        </el-form-item>
        <el-form-item label="证书年使用费:" prop="zhengShuNianShiYongFei">
          <el-input v-model="formData.zhengShuNianShiYongFei" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="支付日期:" prop="zhiFuRiQi">
          <el-date-picker v-model="formData.zhiFuRiQi" type="date" placeholder="选择日期" :clearable="true"></el-date-picker>
        </el-form-item>
        <el-form-item label="是否借出:" prop="shiFouJieChu">
          <el-switch v-model="formData.shiFouJieChu" active-color="#13ce66" inactive-color="#ff4949" active-text="是" inactive-text="否" clearable ></el-switch>
        </el-form-item>
        <el-form-item label="添加人:" prop="tianJiaRen">
          <el-input v-model.number="formData.tianJiaRen" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="备注:" prop="beiZhu">
          <el-input v-model="formData.beiZhu" :clearable="true" placeholder="请输入" />
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
  name: 'ZhengShuBiao'
}
</script>

<script setup>
import {
  createZhengShuBiao,
  updateZhengShuBiao,
  findZhengShuBiao
} from '@/api/zhengShuBiao'

// 自动获取字典
import { getDictFunc } from '@/utils/format'
import { useRoute, useRouter } from "vue-router"
import { ElMessage } from 'element-plus'
import { ref, reactive } from 'vue'
const route = useRoute()
const router = useRouter()

const type = ref('')
const zheng_shu_lei_bieOptions = ref([])
const zheng_shu_suo_shu_cheng_shiOptions = ref([])
const zheng_shu_zhuang_taiOptions = ref([])
const formData = ref({
            group: '',
            zhengShuMingCheng: '',
            zhengShuBianHao: '',
            zhengShuLeiBie: undefined,
            suoYouRenXingMing: '',
            shenFenZhengHao: '',
            shouJiHaoMa: '',
            faZhengBuMen: '',
            zhengShuSuoShuChengShi: undefined,
            zhengShuZhuangTai: undefined,
            zhengShuDaoQiRiQi: new Date(),
            zhengShuNianShiYongFei: '',
            zhiFuRiQi: new Date(),
            shiFouJieChu: false,
            tianJiaRen: 0,
            beiZhu: '',
        })
// 验证规则
const rule = reactive({
})

const elFormRef = ref()

// 初始化方法
const init = async () => {
 // 建议通过url传参获取目标数据ID 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
    if (route.query.id) {
      const res = await findZhengShuBiao({ ID: route.query.id })
      if (res.code === 0) {
        formData.value = res.data.rezhengShuBiao
        type.value = 'update'
      }
    } else {
      type.value = 'create'
    }
    zheng_shu_lei_bieOptions.value = await getDictFunc('zheng_shu_lei_bie')
    zheng_shu_suo_shu_cheng_shiOptions.value = await getDictFunc('zheng_shu_suo_shu_cheng_shi')
    zheng_shu_zhuang_taiOptions.value = await getDictFunc('zheng_shu_zhuang_tai')
}

init()
// 保存按钮
const save = async() => {
      elFormRef.value?.validate( async (valid) => {
         if (!valid) return
            let res
           switch (type.value) {
             case 'create':
               res = await createZhengShuBiao(formData.value)
               break
             case 'update':
               res = await updateZhengShuBiao(formData.value)
               break
             default:
               res = await createZhengShuBiao(formData.value)
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
