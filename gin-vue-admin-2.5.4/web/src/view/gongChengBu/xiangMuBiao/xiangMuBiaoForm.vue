<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" ref="elFormRef" label-position="right" :rules="rule" label-width="80px">
        <el-form-item label="所属:" prop="group">
          <el-input v-model="formData.group" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="项目编号:" prop="xiangMuBianHao">
          <el-input v-model.number="formData.xiangMuBianHao" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="计划开工日期:" prop="jiHuaKaiGongRiQi">
          <el-date-picker v-model="formData.jiHuaKaiGongRiQi" type="date" placeholder="选择日期" :clearable="true"></el-date-picker>
        </el-form-item>
        <el-form-item label="计划结束日期:" prop="jiHuaJieShuRiQi">
          <el-date-picker v-model="formData.jiHuaJieShuRiQi" type="date" placeholder="选择日期" :clearable="true"></el-date-picker>
        </el-form-item>
        <el-form-item label="工程造价:" prop="gongChengZaoJia">
          <el-input-number v-model="formData.gongChengZaoJia" :precision="2" :clearable="true"></el-input-number>
        </el-form-item>
        <el-form-item label="农民工工资保障金:" prop="nongMinGongGongZiBaoZhangJin">
          <el-input-number v-model="formData.nongMinGongGongZiBaoZhangJin" :precision="2" :clearable="true"></el-input-number>
        </el-form-item>
        <el-form-item label="保障金状态:" prop="baoZhangJinZhuangTai">
          <el-select v-model="formData.baoZhangJinZhuangTai" placeholder="请选择" :clearable="true">
            <el-option v-for="(item,key) in baoZhangJinZhuangTaiOptions" :key="key" :label="item.label" :value="item.value" />
          </el-select>
        </el-form-item>
        <el-form-item label="履约保证金:" prop="luYueBaoZhengJin">
          <el-input-number v-model="formData.luYueBaoZhengJin" :precision="2" :clearable="true"></el-input-number>
        </el-form-item>
        <el-form-item label="保证金状态:" prop="baoZhengJinZhuangTai">
          <el-select v-model="formData.baoZhengJinZhuangTai" placeholder="请选择" :clearable="true">
            <el-option v-for="(item,key) in baoZhengJinZhuangTaiOptions" :key="key" :label="item.label" :value="item.value" />
          </el-select>
        </el-form-item>
        <el-form-item label="工作保险金额:" prop="gongZuoBaoXianJinE">
          <el-input-number v-model="formData.gongZuoBaoXianJinE" :precision="2" :clearable="true"></el-input-number>
        </el-form-item>
        <el-form-item label="付款方式:" prop="fuKuanFangShi">
          <el-select v-model="formData.fuKuanFangShi" placeholder="请选择" :clearable="true">
            <el-option v-for="(item,key) in fuKuanFangShiOptions" :key="key" :label="item.label" :value="item.value" />
          </el-select>
        </el-form-item>
        <el-form-item label="预期利润:" prop="yuQiLiRun">
          <el-input-number v-model="formData.yuQiLiRun" :precision="2" :clearable="true"></el-input-number>
        </el-form-item>
        <el-form-item label="实际开工日期:" prop="shiJiKaiGongRiQi">
          <el-date-picker v-model="formData.shiJiKaiGongRiQi" type="date" placeholder="选择日期" :clearable="true"></el-date-picker>
        </el-form-item>
        <el-form-item label="实际结束日期:" prop="shiJiJieShuRiQi">
          <el-date-picker v-model="formData.shiJiJieShuRiQi" type="date" placeholder="选择日期" :clearable="true"></el-date-picker>
        </el-form-item>
        <el-form-item label="结算状态 1.未结算 2.结算完成:" prop="jieSuanZhuangTai">
          <el-select v-model="formData.jieSuanZhuangTai" placeholder="请选择" :clearable="true">
            <el-option v-for="(item,key) in jieSuanZhuangTaiOptions" :key="key" :label="item.label" :value="item.value" />
          </el-select>
        </el-form-item>
        <el-form-item label="项目状态 1.建设中 2.未完成 3.竣工:" prop="xiangMuZhuangTai">
          <el-select v-model="formData.xiangMuZhuangTai" placeholder="请选择" :clearable="true">
            <el-option v-for="(item,key) in xiangMuZhuangTaiOptions" :key="key" :label="item.label" :value="item.value" />
          </el-select>
        </el-form-item>
        <el-form-item label="项目经理:" prop="xiangMuJingLi">
          <el-input v-model.number="formData.xiangMuJingLi" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="项目主管:" prop="xiangMuZhuGuan">
          <el-input v-model.number="formData.xiangMuZhuGuan" :clearable="true" placeholder="请输入" />
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
  name: 'XiangMuBiao'
}
</script>

<script setup>
import {
  createXiangMuBiao,
  updateXiangMuBiao,
  findXiangMuBiao
} from '@/api/xiangMuBiao'

// 自动获取字典
import { getDictFunc } from '@/utils/format'
import { useRoute, useRouter } from "vue-router"
import { ElMessage } from 'element-plus'
import { ref, reactive } from 'vue'
const route = useRoute()
const router = useRouter()

const type = ref('')
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

// 初始化方法
const init = async () => {
 // 建议通过url传参获取目标数据ID 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
    if (route.query.id) {
      const res = await findXiangMuBiao({ ID: route.query.id })
      if (res.code === 0) {
        formData.value = res.data.rexiangMuBiao
        type.value = 'update'
      }
    } else {
      type.value = 'create'
    }
    jieSuanZhuangTaiOptions.value = await getDictFunc('jieSuanZhuangTai')
    xiangMuZhuangTaiOptions.value = await getDictFunc('xiangMuZhuangTai')
    baoZhangJinZhuangTaiOptions.value = await getDictFunc('baoZhangJinZhuangTai')
    baoZhengJinZhuangTaiOptions.value = await getDictFunc('baoZhengJinZhuangTai')
    fuKuanFangShiOptions.value = await getDictFunc('fuKuanFangShi')
}

init()
// 保存按钮
const save = async() => {
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
