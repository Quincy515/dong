<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" ref="elFormRef" label-position="right" :rules="rule" label-width="80px">
        <el-form-item label="备注:" prop="beiZhu">
          <el-input v-model="formData.beiZhu" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="等级要求:" prop="dengJiYaoQiu">
          <el-select v-model="formData.dengJiYaoQiu" placeholder="请选择" :clearable="true">
            <el-option v-for="(item,key) in deng_ji_yao_qiuOptions" :key="key" :label="item.label" :value="item.value" />
          </el-select>
        </el-form-item>
        <el-form-item label="第二报价:" prop="diErBaoJia">
          <el-input v-model="formData.diErBaoJia" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="第二下浮率:" prop="diErXiaFuLu">
          <el-input v-model="formData.diErXiaFuLu" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="第二中标候选人:" prop="diErZhongBiaoHouXuanRen">
          <el-input v-model="formData.diErZhongBiaoHouXuanRen" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="第三报价:" prop="diSanBaoJia">
          <el-input v-model="formData.diSanBaoJia" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="第三下浮率:" prop="diSanXiaFuLu">
          <el-input v-model="formData.diSanXiaFuLu" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="第三中标候选人:" prop="diSanZhongBiaoHouXuanRen">
          <el-input v-model="formData.diSanZhongBiaoHouXuanRen" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="第一报价:" prop="diYiBaoJia">
          <el-input v-model="formData.diYiBaoJia" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="第一下浮率:" prop="diYiXiaFuLu">
          <el-input v-model="formData.diYiXiaFuLu" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="第一中标候选人:" prop="diYiZhongBiaoHouXuanRen">
          <el-input v-model="formData.diYiZhongBiaoHouXuanRen" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="所属:" prop="group">
          <el-input v-model="formData.group" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="开标时间:" prop="kaiBiaoShiJian">
          <el-date-picker v-model="formData.kaiBiaoShiJian" type="date" placeholder="选择日期" :clearable="true"></el-date-picker>
        </el-form-item>
        <el-form-item label="所属地区:" prop="suoShuDiQu">
          <el-select v-model="formData.suoShuDiQu" placeholder="请选择" :clearable="true">
            <el-option v-for="(item,key) in suo_shu_di_quOptions" :key="key" :label="item.label" :value="item.value" />
          </el-select>
        </el-form-item>
        <el-form-item label="投标单位名称:" prop="touBiaoDanWeiMingCheng">
          <el-input v-model="formData.touBiaoDanWeiMingCheng" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="项目名称:" prop="xiangMuMingCheng">
          <el-input v-model="formData.xiangMuMingCheng" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="招标规模:" prop="zhaoBiaoGuiMo">
          <el-input v-model="formData.zhaoBiaoGuiMo" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="资质要求:" prop="ziZhiYaoQiu">
          <el-select v-model="formData.ziZhiYaoQiu" placeholder="请选择" :clearable="true">
            <el-option v-for="(item,key) in zi_zhi_yao_qiuOptions" :key="key" :label="item.label" :value="item.value" />
          </el-select>
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
  name: 'TouBiaoShuJuTongJi'
}
</script>

<script setup>
import {
  createTouBiaoShuJuTongJi,
  updateTouBiaoShuJuTongJi,
  findTouBiaoShuJuTongJi
} from '@/api/touBiaoShuJuTongJi'

// 自动获取字典
import { getDictFunc } from '@/utils/format'
import { useRoute, useRouter } from "vue-router"
import { ElMessage } from 'element-plus'
import { ref, reactive } from 'vue'
const route = useRoute()
const router = useRouter()

const type = ref('')
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

// 初始化方法
const init = async () => {
 // 建议通过url传参获取目标数据ID 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
    if (route.query.id) {
      const res = await findTouBiaoShuJuTongJi({ ID: route.query.id })
      if (res.code === 0) {
        formData.value = res.data.retouBiaoShuJuTongJi
        type.value = 'update'
      }
    } else {
      type.value = 'create'
    }
    deng_ji_yao_qiuOptions.value = await getDictFunc('deng_ji_yao_qiu')
    suo_shu_di_quOptions.value = await getDictFunc('suo_shu_di_qu')
    zi_zhi_yao_qiuOptions.value = await getDictFunc('zi_zhi_yao_qiu')
}

init()
// 保存按钮
const save = async() => {
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
