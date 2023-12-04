<template>
  <div style="display: flex;justify-content: center;">
      <n-card title="添加试题" style="width: 1000px;margin-top: 100px;border-radius: 20px" hoverable>
        <div style="display: flex;justify-content: center;margin-left: 200px">
          <n-row >
            <n-col :span="12">
              <n-statistic label="统计数据" :value="99">
                <template #prefix>
                  <n-icon>
                    <Watch/>
                  </n-icon>
                </template>
                <template #suffix>
                  / 100
                </template>
              </n-statistic>
            </n-col>
            <n-col :span="12">
              <n-statistic label="活跃用户" :value="80">
                <template #prefix>
                  <n-icon>
                    <ShareSocialSharp/>
                  </n-icon>
                </template>
                <template #suffix>
                  / 100
                </template>
              </n-statistic>
            </n-col>
          </n-row>
        </div>
        <n-form :model="markForm"  style="width: 800px;margin-top: 50px" label-placement="left" label-width="auto">
          <n-form-item label="标题" required>
            <n-input v-model:value="markForm.name" placeholder="请输入试题标题"></n-input>
          </n-form-item>
          <n-form-item label="虚拟机" required>
            <n-select placeholder="请选择对应模块虚拟机"  multiple :options="vmoptions" v-model:value="vmList"></n-select>
          </n-form-item>
          <n-form-item label="试题类型" required>
            <n-select placeholder="请选择试题类型" :options="typeoptions" v-model:value="markForm.type"></n-select>
          </n-form-item>
          <n-form-item label="题目" required>
            <v-md-editor v-model="markdownContent" @save="saveMd"  style="width: 800px;height:600px;"></v-md-editor>
          </n-form-item>
        </n-form>
        <div>
          <n-button type="success" @click="submitAddMd" style="margin-right: 10px" secondary>添加试题</n-button>
          <n-button type="error" secondary>取消试题</n-button>
        </div>
      </n-card>
  </div>
</template>


<script setup>
import {onMounted, reactive, ref, toRefs} from "vue";
import {addExam} from "@/api/markdown";
import {
  Watch as Watch,
  ShareSocialSharp as ShareSocialSharp
} from "@vicons/ionicons5"
import {getVmsLabel} from "@/api/manager";
import router from "@/router";
const markdownContent = ref("")
const markForm =  ref({
  name: "",
  content: "",
  type: "选拔",
  vmList: ""
})

const typeoptions = reactive([
  {
    label: "选拔",
    value: "选拔"
  },
  {
    label: "省赛",
    value: "省赛"
  },
  {
    label: "国赛",
    value: "国赛"
  }
])
const vmoptions = ref([])
const vmList = ref([])

function saveMd(text,html) {
  localStorage.setItem("text",text)
  localStorage.setItem("html",html)

}
function submitAddMd(){
  var content = JSON.stringify(markdownContent._rawValue)
  markForm.value.content = JSON.parse(content)
  markForm.value.vmList = JSON.stringify(vmList.value)
  addExam(markForm.value).then(res => {
    router.go(0)
    return window.$message.success("添加成功")
  })

}

function getAllHostLabel() {
  getVmsLabel().then(res => {
    vmoptions.value = res.data
  })
}
onMounted(() => {
  markForm.value.content = localStorage.getItem("text") || " "
  getAllHostLabel()
})

getAllHostLabel()
</script>


<style>

</style>