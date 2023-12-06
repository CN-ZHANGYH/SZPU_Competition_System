<template>
  <div>
    <div style="display: flex;justify-content: space-between">
      <n-breadcrumb>
        <n-breadcrumb-item>
          <n-icon :component="Planet" :size="30"/> <span style="font-weight: bold;font-size: 20px;margin-left: 10px">试题板块</span>
        </n-breadcrumb-item>
      </n-breadcrumb>
      <div style="display: flex;justify-content: center">
        <n-select :options="contentOptions" v-model:value="searchValue" style="width: 250px" placeholder="请选择试题内容"></n-select>
        <span><n-button type="info" @click="changeContentInfo" style="margin-right: 5px">切换</n-button></span>
      </div>
    </div>
    <div>
      <n-tabs type="line" animated>
        <n-tab-pane name="oasis" tab="试题">
          <v-md-preview v-if="text != ''" :text="text"></v-md-preview>
          <n-result status="success" title="试题板块" description="可以切换试题进行作答" v-if="text == ''" style="margin-top: 50px;height: 850px">
          </n-result>
        </n-tab-pane>
        <n-tab-pane name="the beatles" tab="笔记" style="height: 850px">
          <v-md-editor  v-model="note" @save="saveMd" height="800px"></v-md-editor>
        </n-tab-pane>
      </n-tabs>

    </div>
  </div>
</template>

<script setup>
import {
  Planet as Planet,


} from "@vicons/ionicons5"
import {getCurrentInstance, onMounted, ref} from "vue";
import {getCountExam, getExamInfo} from "@/api/markdown";

const vmList = ref([])
const text = ref("")
const note = ref("")
const contentOptions = ref([])
const searchValue = ref("")
const instance = getCurrentInstance()
const vmListOptions = ref([])
function changeContentInfo(){
  getExamInfo({name: searchValue.value}).then(res => {
    if (searchValue.value == ""){
      return window.$message.error("试题为空")
    }

    text.value = res.data.Content
    vmList.value = JSON.parse(res.data.VmList)

    vmListOptions.value = []
    for (let index = 0; index < vmList.value.length; index++) {
      const vm_url = vmList.value[index];
      var obj = {
        label: "环境" + (index + 1) + ":" + vm_url,
        value: vm_url
      }
      vmListOptions.value.push(obj)
    }
    console.log(vmListOptions)

    instance.emit('myEvent',vmListOptions.value)
  })
}

onMounted(() => {
  getCountExam().then(res => {
    contentOptions.value = res.data
  })

  note.value = JSON.parse(localStorage.getItem("note")) || "";
})

function saveMd(text,html){
  localStorage.setItem("note",JSON.stringify(text))
}
</script>


<style>
</style>
