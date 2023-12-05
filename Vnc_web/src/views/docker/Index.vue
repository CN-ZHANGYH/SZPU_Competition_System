<script setup>
import {h, reactive, ref} from "vue";
import {NButton, NTag} from "naive-ui";
import {getContainerList, getDockerInfo} from "@/api/docker";
import {CheckmarkCircle} from "@vicons/ionicons5";

const createColumns = reactive([
    {
      title: "镜像ID",
      key: "Id"
    },
    {
      title: "镜像名称",
      key: "Image"
    },
    {
      title: "容器名称",
      key: "Names"
    },
    {
      title: "端口",
      key: "Ports",
      render(row) {
        const tags = row.Ports.map((tagKey) => {
          return h(
              NTag,
              {
                style: {
                  marginRight: "6px",
                  marginTop: "10px"
                },
                type: "info",
                bordered: false
              },
              {
                default: () => tagKey.PublicPort != undefined ? tagKey.PrivatePort + "->" + tagKey.PublicPort : tagKey.PrivatePort
              }
          );
        });
        return tags;
      }
    },
    {
      title: "Action",
      key: "actions",
      render(row) {
        return h(
            NButton,
            {
              size: "small",
              onClick: () => sendMail(row)
            },
            { default: () => "Send Email" }
        );
      }
    }
])


function getList() {
  getContainerList().then(res => {
    console.log(res)
    data.value = res.data
  })

  getDockerInfo().then(res => {
    dockerInfo.value = res.data
  })
}


getList()
const data = ref([])
const dockerInfo = ref({})
</script>

<template>
<n-card>
<div style="display: flex;">
  <div style="justify-content: flex-start">
    <n-card style="width: 450px;height: 800px;border-radius: 20px;margin-right: 5px" hoverable>
      <div>
        <n-descriptions label-placement="left" title="系统环境详细信息" :column="1" style="margin-bottom: 20px">
          <n-descriptions-item label="容器版本">
            <n-tag type="success" round :bordered="false">
              {{dockerInfo.id}}
              <template #icon>
                <n-icon :component="CheckmarkCircle" />
              </template>
            </n-tag>
          </n-descriptions-item>
          <n-descriptions-item label="系统版本">
            <n-tag type="error" round :bordered="false">
              {{ dockerInfo.system }}
              <template #icon>
                <n-icon :component="CheckmarkCircle" />
              </template>
            </n-tag>
          </n-descriptions-item>
          <n-descriptions-item label="运行状态">
            <n-tag type="warning" round :bordered="false">
              {{dockerInfo.status}}
              <template #icon>
                <n-icon :component="CheckmarkCircle" />
              </template>
            </n-tag>
          </n-descriptions-item>
          <n-descriptions-item label="容器版本">
            <n-tag type="info" round :bordered="false">
              {{dockerInfo.version}}
              <template #icon>
                <n-icon :component="CheckmarkCircle" />
              </template>
            </n-tag>          </n-descriptions-item>
        </n-descriptions>
        <div style="display: flex;justify-content: space-between">
          <span>内存使用率:</span>
          <n-progress
              type="line"
              :percentage="dockerInfo.memory_percent"
              :indicator-placement="'inside'"
              processing
              status="success"
              :height="24"
              style="margin-bottom: 10px;width: 300px"
          />
        </div>
        <div style="display: flex;justify-content: space-between">
          <span>已使用内存:</span>
          <n-progress
              type="line"
              :percentage="dockerInfo.memory_total"
              :indicator-placement="'inside'"
              processing
              status="default"
              :height="24"
              style="margin-bottom: 10px;width: 300px"
          />
        </div>
        <div style="display: flex;justify-content: space-between">
          <span>活跃的内存:</span>
          <n-progress
              type="line"
              :percentage="dockerInfo.memory_used"
              :indicator-placement="'inside'"
              processing
              status="warning"
              :height="24"
              style="margin-bottom: 10px;width: 300px"
          />
        </div>
        <div style="display: flex;justify-content: space-between">
          <span>空闲的内存:</span>
          <n-progress
              type="line"
              :percentage="dockerInfo.memory_ava"
              :indicator-placement="'inside'"
              processing
              status="error"
              :height="24"
              style="margin-bottom: 10px;width: 300px"
          />
        </div>
      </div>
    </n-card>
  </div>
  <div style="justify-content: flex-end;width: 100%">
    <n-card style="border-radius: 20px" hoverable >
      <n-data-table
          :columns="createColumns"
          :data="data"
          :bordered="false"
          :pagination="pagination"
          default-expand-all
      />
      <n-pagination v-model:page="page" :page-count="100" style="margin-left: 68%;margin-top: 50px"/>
    </n-card>
  </div>
</div>
</n-card>
</template>

<style scoped>

</style>