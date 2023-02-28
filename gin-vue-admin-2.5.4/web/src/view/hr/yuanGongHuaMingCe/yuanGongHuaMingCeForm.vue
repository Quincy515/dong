<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" ref="elFormRef" label-position="right" :rules="rule" label-width="80px">
        <el-form-item label="备注:" prop="beiZhu">
          <el-input v-model="formData.beiZhu" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="部门:" prop="buMen">
          <el-input v-model="formData.buMen" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="出生日期:" prop="chuShengRiQi">
          <el-input v-model="formData.chuShengRiQi" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="从事岗位或工种:" prop="congShiGangWeiHuoGongZhong">
          <el-input v-model="formData.congShiGangWeiHuoGongZhong" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="个人社保密码:" prop="geRenSheBaoMiMa">
          <el-input v-model="formData.geRenSheBaoMiMa" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="个人社保账号:" prop="geRenSheBaoZhangHao">
          <el-input v-model="formData.geRenSheBaoZhangHao" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="所属:" prop="group">
          <el-input v-model="formData.group" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="合同约定工资:" prop="heTongYueDingGongZi">
          <el-input v-model="formData.heTongYueDingGongZi" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="户口所在地:" prop="huKouSuoZaiDi">
          <el-input v-model="formData.huKouSuoZaiDi" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="劳动合同截止日期:" prop="laoDongHeTongJieZhiRiQi">
          <el-input v-model="formData.laoDongHeTongJieZhiRiQi" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="劳动合同开始日期:" prop="laoDongHeTongKaiShiRiQi">
          <el-input v-model="formData.laoDongHeTongKaiShiRiQi" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="离职日期:" prop="liZhiRiQi">
          <el-input v-model="formData.liZhiRiQi" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="联系电话:" prop="lianXiDianHua">
          <el-input v-model="formData.lianXiDianHua" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="门禁编号:" prop="menJinBianHao">
          <el-input v-model="formData.menJinBianHao" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="年龄:" prop="nianLing">
          <el-input v-model="formData.nianLing" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="入职时间:" prop="ruZhiShiJian">
          <el-input v-model="formData.ruZhiShiJian" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="身份证号码:" prop="shenFenZhengHaoMa">
          <el-input v-model="formData.shenFenZhengHaoMa" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="文化程度:" prop="wenHuaChengDu">
          <el-input v-model="formData.wenHuaChengDu" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="现居住地址:" prop="xianJuZhuDiZhi">
          <el-input v-model="formData.xianJuZhuDiZhi" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="性别:" prop="xingBie">
          <el-input v-model="formData.xingBie" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="姓名:" prop="xingMing">
          <el-input v-model="formData.xingMing" :clearable="true" placeholder="请输入" />
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
  name: 'YuanGongHuaMingCe'
}
</script>

<script setup>
import {
  createYuanGongHuaMingCe,
  updateYuanGongHuaMingCe,
  findYuanGongHuaMingCe
} from '@/api/yuanGongHuaMingCe'

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
            buMen: '',
            chuShengRiQi: '',
            congShiGangWeiHuoGongZhong: '',
            geRenSheBaoMiMa: '',
            geRenSheBaoZhangHao: '',
            group: '',
            heTongYueDingGongZi: '',
            huKouSuoZaiDi: '',
            laoDongHeTongJieZhiRiQi: '',
            laoDongHeTongKaiShiRiQi: '',
            liZhiRiQi: '',
            lianXiDianHua: '',
            menJinBianHao: '',
            nianLing: '',
            ruZhiShiJian: '',
            shenFenZhengHaoMa: '',
            wenHuaChengDu: '',
            xianJuZhuDiZhi: '',
            xingBie: '',
            xingMing: '',
        })
// 验证规则
const rule = reactive({
})

const elFormRef = ref()

// 初始化方法
const init = async () => {
 // 建议通过url传参获取目标数据ID 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
    if (route.query.id) {
      const res = await findYuanGongHuaMingCe({ ID: route.query.id })
      if (res.code === 0) {
        formData.value = res.data.reyuanGongHuaMingCe
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
               res = await createYuanGongHuaMingCe(formData.value)
               break
             case 'update':
               res = await updateYuanGongHuaMingCe(formData.value)
               break
             default:
               res = await createYuanGongHuaMingCe(formData.value)
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
