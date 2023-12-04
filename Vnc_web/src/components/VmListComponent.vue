<template>
  <n-data-table
      :bordered="false"
      :columns="columns"
      :data="createData"
      style="margin-top: 20px;margin-left: 5px;"
  />
</template>


<script setup>
import {NButton, NTag} from "naive-ui";
import {h, onMounted, reactive, ref} from "vue";
import {getVmList, removeVm} from "@/api/manager";
import router from "@/router";

const columns = reactive([
  {
    title: "主机名称",
    key: "Name",
    render(row) {
      return h(
          NTag,
          {
            style: {
              marginRight: "6px"
            },
            type: "info",
            bordered: true
          },
          {
            default: () => row.Name
          }
      );
    }
  },
  {
    title: "IP/域名",
    key: "Host",
    render(row) {
      return h(
          NTag,
          {
            style: {
              marginRight: "6px"
            },
            type: "error",
            bordered: false
          },
          {
            default: () => row.Host
          }
      );
    }
  },
  {
    title: "VNC密码",
    key: "Password",
    render(row) {
      return h(
          NTag,
          {
            style: {
              marginRight: "6px"
            },
            type: "warning",
            bordered: false
          },
          {
            default: () => row.Password
          }
      );
    }
  },
  {
    title: "VNC端口",
    key: "Port",
    render(row) {
      return h(
          NTag,
          {
            style: {
              marginRight: "6px"
            },
            type: "info",
            bordered: false
          },
          {
            default: () => row.Port
          }
      );
    }
  },
  {
    title: "访问链接",
    key: "Url",
    render(row) {
      return h(
          NTag,
          {
            style: {
              marginRight: "6px"
            },
            type: "primary",
            bordered: false
          },
          {
            default: () => row.Url
          }
      );
    }
  },
  {
    title: "操作",
    key: "action",
    render(row) {
      return h(
          NButton,
          {
            style: {
              marginRight: "6px"
            },
            type: "error",
            bordered: false,
            onClick: () => deleteVm(row)
          },
          {
            default: () => "删除虚拟机"
          }
      );
    }
  }
])
const createData = ref([])



onMounted(() => {
  getList()
})


function getList(){
  getVmList().then(res => {
    console.log(res)
    createData.value = res.data.data
  })
}

function deleteVm(row){
  removeVm({name: row.Name}).then(() => {
    window.$message.success("删除成功")
    getVmList();
    router.go(0)
  })
}

</script>