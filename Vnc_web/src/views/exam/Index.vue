<template>
  <n-card title="试题管理">
    <n-data-table
        :columns="columns"
        :data="data"
        :pagination="pagination"
        :bordered="false"
    />
  </n-card>
  <n-drawer v-model:show="active" :width="900" :placement="placement">
    <n-drawer-content title="赛题详细内容">
      <n-card hoverable style="border-radius: 100px;">
        <span>类型：</span><n-tag type="success" :bordered="false">{{contentInfo.Type}}</n-tag>
      </n-card>
      <v-md-preview :text="text"></v-md-preview>
    </n-drawer-content>
  </n-drawer>
</template>


<script setup>

import {h, onMounted, ref} from "vue";
import {NButton, NTag} from "naive-ui";
import {getExamList, removeContentInfo} from "@/api/markdown";
const text = ref(``)
const contentInfo = ref({})

const columns = [
  {
    title: "试题ID",
    key: "Id"
  },
  {
    title: "试题名称",
    key: "Name"
  },
  {
    title: "试题环境",
    key: "VmList",
    render(row) {
      const tags = JSON.parse(row.VmList).map((tagKey) => {
        return h(
            NTag,
            {
              style: {
                marginRight: "6px"
              },
              strong: true,
              type: "success",
              bordered: false
            },
            {
              default: () => tagKey
            }
        );
      });
      return tags;
    }
  },
  {
    title: "试题内容",
    key: "Content",
    render(row) {
      return h(
          NButton,
          {
            type: "info",
            strong: true,
            tertiary: true,
            size: "small",
          },
          { default: () => row.Content.substring(0,50)
          }
      );
    }
  },
  {
    title: "查看",
    key: "action",
    render(row) {
      return h(
          NButton,
          {
            type: "info",
            strong: true,
            size: "small",
            onClick: () => viewContent(row)
          },
          { default: () => "查看" }
      );
    }
  },
  {
    title: "删除",
    key: "action",
    render(row) {
      return h(
          NButton,
          {
            type: "error",
            strong: true,
            size: "small",
            onClick: () => deleteContent(row)
          },
          { default: () => "删除" }
      );
    }
  }
];

const data = ref([]);
const active = ref(false);
const placement = ref("left");



function getExamAllList(){
  getExamList().then(res => {
    data.value = res.data.data
  })
}

function viewContent(row) {
  active.value = true;
  placement.value = "left";
  text.value = row.Content
  contentInfo.value = row
}

function deleteContent(row){
  removeContentInfo({name: row.Name}).then(res => {
    if (res.code == 0) {
      getExamAllList()
      window.$message.success(res.msg)
    }else  {
      window.$message.error(res.msg)
    }
  })
}
onMounted(() => {
  getExamAllList()
})

getExamAllList()
</script>
