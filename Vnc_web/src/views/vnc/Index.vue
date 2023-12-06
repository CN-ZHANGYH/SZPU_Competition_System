<template>
  <div style="display: flex;justify-content: space-between;margin-top: 5px;">
    <n-card id="cardLeft"  style="height: 100%;width: auto;margin-left: 10px;border-radius: 20px;margin-right: 20px;"  hoverable>
    <!-- 题目组件 -->
      <n-layout has-sider>
        <n-layout-sider
            bordered
            show-trigger="bar"
            collapse-mode="transform"
            :collapsed-width="0"
            width="480"
            :native-scrollbar="false"
            :show-collapsed-content="false"
            style="max-height: 850px;"
        >
          <ContentIndex @myEvent="getVmList"></ContentIndex>
        </n-layout-sider>
        <n-layout style="max-height: 800px;margin-right: 20px" />
      </n-layout>

    </n-card>
    <div id="cardRight" style="width: 100%;">
      <n-card style="border-radius: 20px" hoverable>
        <n-card  style="width: 100%;border-radius: 20px;height: 60px" hoverable>
          <div style="display: flex;justify-content: space-between;margin-top: -10px">
            <!-- <n-dynamic-tags v-model:value="tags" :render-tag="renderTag" size="large" /> -->
            <n-select :options="vmListOptions" style="width: 400px;" placeholder="请选择题目配套的环境" @update:value="handleUpdateValueVmUrl"/>
            <div style="display:flex;justify-content: space-between;">
              <n-select placeholder="请选择自定义" :options="vmOptions" style="width: 200px;margin-right: 5px" @update:value="handleUpdateValue" ></n-select>
              <n-button type="success" @click="changeFISCO" size="small" strong style="margin-right: 5px">FISCO官网</n-button>
              <n-button type="info" @click="changeWEBASE" size="small" strong style="margin-right: 5px">WEBASE官网</n-button>
              <n-button type="error" @click="changeSolidity" size="small" strong style="margin-right: 5px">Solidity官网</n-button>
              <n-button type="warning" @click="showModal = true" size="small" strong style="margin-right: 50px">自定义</n-button>
            </div>
          </div>
        </n-card>
        <!-- 虚拟机组件 -->
        <n-spin :show="show">
          <VmComponent style="width: 100%" :vmUrl="vmUrl"></VmComponent>
        </n-spin>
      </n-card>
    </div>
    <n-modal v-model:show="showModal">
      <n-card
          style="width: 600px;"
          title="添加自定义虚拟机地址"
          :bordered="false"
          size="huge"
          role="dialog"
          aria-modal="true"
      >
        <n-input placeholder="请输入新的主机地址" v-model:value="newVmValue"></n-input>
        <n-button type="success" @click="addOwnerVm" style="margin-top: 10px">添加</n-button>
      </n-card>
    </n-modal>
  </div>
</template>


<script setup>
import {h, onMounted, ref, watch} from "vue";
import {NTag} from "naive-ui";
import ContentIndex from "@/views/content/Index.vue"
import VmComponent from "@/components/VmComponent.vue";

// const vmList = ref([])
const show = ref(true)
const tags = ref([]);
const vmUrl = ref("")
const newVmValue = ref("")
const showModal = ref(false)
const props = defineProps(['vmUrl'])
const vmListOptions = ref([])


function getVmList(arr) {
  vmListOptions.value = arr
  // tags.value = vmList.value
  vmUrl.value = vmListOptions.value[0].value
}

const vmOptions = ref([])
onMounted(() => {
  vmListOptions.value = vmListOptions.value || []
  if (tags.value.length == 0) {
    changeFISCO()
    setTimeout(() => {
      show.value = false
    },1000)
  }
})

function handleUpdateValueVmUrl(value,option){
  vmUrl.value = value
}

// const renderTag =  (tag, index) => {
//   return h(
//       NTag,
//       {
//         type: "success",
//         bordered: false,
//         closable: true,
//         onClose: () => {
//           tags.value.splice(index, 1);
//           localStorage.setItem("HOST",JSON.stringify(tags.value))
//         },
//         onclick: () => {
//           vmUrl.value = "http://" + tag + ":6082/vnc.html"
//           window.$message.success("切换成功")
//         }
//       },
//       {
//         default: () => tag
//       }
//   );
// }

watch(tags, (newTags, oldTags) => {
  if (vmUrl.value == "" && tags.value.length == 1) {
    vmUrl.value = newTags
  }
  localStorage.setItem("HOST",JSON.stringify(tags.value))
});

const changeFISCO = () => {
  show.value = true
  vmUrl.value = "https://fisco-bcos-documentation.readthedocs.io/zh_CN/latest/index.html"
  setTimeout(() => {
    show.value = false
  },1000)
}
const changeWEBASE = () => {
  show.value = true
  vmUrl.value = "https://webasedoc.readthedocs.io/zh_CN/latest/"
  setTimeout(() => {
    show.value = false
  },1000)
}

const changeSolidity = () => {
  show.value = true
  // vmUrl.value = "https://www.osgeo.cn/solidity/"
  vmUrl.value = "https://wtf.academy/solidity-start/"
  setTimeout(() => {
    show.value = false
  },1000)
}

const addOwnerVm = () => {
  if (newVmValue.value == ""){
    return window.$message.error("不能为空")
  }

  vmOptions.value.push({
    label: newVmValue.value,
    value: newVmValue.value
  })
  vmUrl.value = newVmValue.value
  showModal.value = false
  localStorage.setItem("vmOptions",JSON.stringify(vmOptions))
}

const handleUpdateValue = (value, option) => {
  console.log(option)
  if (vmOptions.value.length <= 0 || vmOptions.value == undefined) {
    return window.$message.error("当前为空")
  }
  vmUrl.value = value
}

var cardLeft = document.getElementById("cardLeft");
var cardRight = document.getElementById("cardRight");

// cardLeft.addEventListener("mousedown",() => {
//
// });

</script>


<style scoped>
.vnc-class {
  width: 1360px;
  height: 768px;
  border-radius: 15px;
  margin-top: 15px;
}
#cardRight{
  flex: 1;
}
</style>
