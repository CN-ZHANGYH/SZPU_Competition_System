<script setup>
import {computed, h, reactive, ref, toRefs} from "vue";
import {NAvatar, NButton, NIcon, NTag} from "naive-ui";
import {
  createDockerContainer,
  getContainerList,
  getDockerInfo, pullDockerImageByName,
  removeDockerContainer, searchContainer, searchDockerImagesList, selectDockerImagesLabel,
  startDockerContainer,
  stopDockerContainer
} from "@/api/docker";
import {CheckmarkCircle,Navigate,RocketSharp,LogoDocker} from "@vicons/ionicons5";
import {formatTime} from "@/utils/formatDate"
const createColumns = reactive([
    {
      title: "镜像ID",
      key: "Id",
      render(row) {
        return h(
            NTag,
            {
              style: {
                with: "15px",
                marginRight: "6px",
                marginTop: "10px"
              },
              type: "success",
              bordered: false,
              round: true
            },
            {
              icon: () => h(NIcon,{component: Navigate}),
              default: () => row.Id
            },

        );
      }
    },
    {
      title: "镜像名称",
      key: "Image",
      render(row) {
        return h(
            NTag,
            {
              style: {
                with: "500px",
                marginRight: "6px",
                marginTop: "10px"
              },
              type: "info",
              bordered: false,
              round: true
            },
            {
              icon: () => h(NIcon,{component: RocketSharp}),
              default: () => row.Image
            },

        );
      }
    },
    {
      title: "容器名称",
      key: "Names"
    },
    {
      title: "端口",
      key: "Ports",
      render(row) {
        const maxTagsPerRow = 3;  // 每行最多显示的标签数量
        const tags = [];
        for (let i = 0; i < row.Ports.length; i += maxTagsPerRow) {
          const slicedPorts = row.Ports.slice(i, i + maxTagsPerRow);
          const tagElements = slicedPorts.map((tagKey) => {
            return h(
                NTag,
                {
                  style: {
                    marginRight: "6px",
                    marginTop: "10px"
                  },
                  type: "info",
                },
                {
                  default: () =>
                      tagKey.PublicPort != undefined
                          ? tagKey.PrivatePort + "->" + tagKey.PublicPort
                          : tagKey.PrivatePort
                }
            );
          });
          tags.push(h("div", {}, tagElements));  // 将每行的标签放入一个div中
        }
        return tags;
      }
    },
    {
      title: "状态",
      key: "Status",
      render(row) {
        return h(
            NTag,
            {
              style: {
                with: "15px",
                marginRight: "6px",
                marginTop: "10px"
              },
              type: row.Status == "running" ? "success" : "error",
              bordered: false,
            },
            {
              default: () => row.Status
            },

        );
      }
    },
    {
      title: "操作",
      key: "actions",
      render(row) {
        let actions = ["开启","停止","删除"]
        const tags = actions.map((tagKey) => {
          return h(
              NButton,
              {
                style: {
                  marginRight: "6px",
                  marginTop: "10px"
                },
                type: tagKey == "开启" ? "success" : tagKey == "停止" ? "info" : "error",
                bordered: false,
                onClick: () => handleDockerContainer(tagKey,row.Id)
              },
              {
                default: () => tagKey
              },

          );
        });
        return tags;
      }
    },
])


function getList() {
  getContainerList().then(res => {
    containerList.value = res.data || []
  })

  getDockerInfo().then(res => {
    dockerInfo.value = res.data
  })

  selectDockerImagesLabel().then(res => {
    imagesOptions.value = res.data
  })
}

function handleDockerContainer(value,id){
  if (value == "开启") {
    startDockerContainer({containerId: id}).then(res => {
      if (res.code == 0) {
        window.$message.success(res.msg)
        getList()
      }else {
        window.$message.error(res.msg)
      }
    })
  }else if (value == "停止") {
    stopDockerContainer({containerId: id}).then(res => {
      if (res.code == 0) {
        window.$message.success(res.msg)
        getList()
      }else {
        window.$message.error(res.msg)
      }
    })
  }else if (value == "删除") {
    removeDockerContainer({containerId: id}).then(res => {
      if (res.code == 0) {
        window.$message.success(res.msg)
        getList()
      }else {
        window.$message.error(res.msg)
      }
    })
  }

}


getList()
const data = ref([])
const dockerInfo = ref({})
const searchKeyword = ref("")
const active = ref(false);
const fromData =  reactive({
  form: {},
  imagesOptions: [],
  Port_List: [],
  Host_Port_List: [],
  Container_Port_List: [],
  Volume_List: [],
  Host_Volume_List: [],
  Container_Volume_List: []
})

const {form,imagesOptions,Port_List,Host_Port_List,Container_Port_List,Volume_List,Host_Volume_List,Container_Volume_List} = toRefs(fromData)

// 分页管理
const itemsPerPage = 9;
const currentPage = ref(1)
const totalPages = computed(() => Math.ceil(containerList.value.length / itemsPerPage));
const containerList = ref([])

const displayContainerList = computed(() => {
  const startIndex = (currentPage.value - 1) * itemsPerPage;
  const endIndex = startIndex + itemsPerPage;
  return containerList.value.slice(startIndex, endIndex);
});
function handlePageChange(page) {
  currentPage.value = page;
  getList()
}

function search(){
  if (searchKeyword.value == "") {
    getList()
  }else  {
    searchContainer({containerId: searchKeyword.value}).then(res => {
      containerList.value = res.data
    })
  }
}

const restartPolicyOptions = reactive([
  {
    label: "无重启策略",
    value: "no"
  },
  {
    label: "容器总是自动重启",
    value: "always"
  },
  {
    label: "容器在非零退出状态时重启（默认最多重启3次）",
    value: "on-failure"
  },  {
    label: "除非手动停止，否则容器总是自动重启",
    value: "unless-stopped"
  }
])

const imageKeyWord = ref("")

const searchImageListData = ref([])
const searchImageListColumns =  reactive([
  {
    title: "操作",
    key: "action",
    width: 100,
    render(row) {
      return h(
          NButton,
          {
            type: "success",
            bordered: false,
            strong: true,
            secondary: true,
            onClick: () => pullDockerImage(row)
          },
          {
            default: () => "拉取"
          },
      );
    }
  },
  {
    title: "镜像名称",
    key: "Name",
    render(row) {
      return h(
          NTag,
          {
            type: "info",
            bordered: false,
            round: true
          },
          {
            icon: () => h(NIcon,{component: RocketSharp}),
            default: () => row.Name
          },

      );
    }
  },
  {
    title: "镜像描述",
    key: "Desc",
    width: 200,
  }
])

const addItem = () => {
  Port_List.value.push({
    port: "",
    label: "HOST"
  });
  Port_List.value.push({
    port: "",
    label: "容器"
  });
};

const removeItem = (index) => {
  Port_List.value.splice(index, 1);
  Port_List.value.splice(index -1, 1);
};
function createVm(){
  active.value = true
  form.value = {}
  Port_List.value = []
  Volume_List.value = []
  addItem()
  addVolume()
}

function addVolume(){
  Volume_List.value.push({
    path: "",
    label: "宿主机"
  });
  Volume_List.value.push({
    path: "",
    label: "容器"
  });
}

function removeVolume(index) {
  Volume_List.value.splice(index, 1);
  Volume_List.value.splice(index -1, 1);
}


function searchDockerImage(){
  if (imageKeyWord.value == "") {
    window.$message.error("请输入镜像名称")
    return
  }
  searchDockerImagesList({name: imageKeyWord.value}).then(res => {
    if (res.code == 0) {
      searchImageListData.value = res.data
      window.$message.success(res.msg)
    }else {
      window.$message.error(res.msg)

    }
  })
}
function submit(){
  var date = new Date();
  for (let i = 0; i < Port_List.value.length; i++) {
    if (Port_List.value[i].port == "") {
      window.$message.error("请填写端口")
      return
    }
    if (Port_List.value[i].label == "HOST") {
      Host_Port_List.value.push(Port_List.value[i].port)
    }else if (Port_List.value[i].label == "容器") {
      Container_Port_List.value.push(Port_List.value[i].port)
    }
  }

  for (let i = 0; i < Volume_List.value.length; i++) {
    if (Volume_List.value[i].label == "宿主机") {
      Host_Volume_List.value.push(Volume_List.value[i].path)
    }else if (Volume_List.value[i].label == "容器") {
      Container_Volume_List.value.push(Volume_List.value[i].path)
    }
  }
  form.value.Host_Port_List = Host_Port_List.value
  form.value.Container_Port_List = Container_Port_List.value
  form.value.Host_Volume_List = Host_Volume_List.value
  form.value.Container_Volume_List = Container_Volume_List.value

  createDockerContainer(form.value).then(res => {
    if (res.code == 0) {
      window.$message.success(res.msg)
      getList()
      active.value = false
      window.$notification.create({
        title: `创建${form.value.ContainerName}容器成功`,
        description: "容器正在创建中,请稍等片刻",
        content:
            `
            当前的容器名称： ${form.value.ContainerName}
            当前的容器镜像： ${form.value.ContainerImage}
            当前的容器端口： ${form.value.Container_Port_List}
            当前的主机端口： ${form.value.Host_Port_List}
            `,
        meta: formatTime(date),
        avatar: () => h(NIcon, {
          component: LogoDocker
        }),
      })
    }else {
      window.$message.error(res.msg)
      active.value = false
    }
  })

}

function pullDockerImage(row) {
  pullDockerImageByName({name:row.Name}).then(res => {
    if (res.code == 0){
      window.$message.success(res.msg)
    }else {
      window.$message.error(res.msg)
    }
  })
}
</script>

<template>
<n-card style="height: 100%">
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
      <n-divider />
      <div class="search-container">
        <n-input v-model:value="imageKeyWord" placeholder="请输入镜像名称" size="large" status="success"/>
        <n-button type="info" size="large" @click="searchDockerImage" >搜索</n-button>
      </div>
        <n-data-table
            :columns="searchImageListColumns"
            :data="searchImageListData"
            :max-height="260"
            :min-height="260"
            :bordered="false"
            :scroll-x="600"
        />
    </n-card>
  </div>
  <div style="justify-content: flex-end;width: 100%">
    <n-card style="border-radius: 20px" hoverable  title="容器镜像管理">
      <div class="search-container">
        <n-input v-model:value="searchKeyword" placeholder="请输入镜像ID" size="large" status="success" style="width: 200px"/>
        <n-button type="primary" size="large" @click="search" >查询</n-button>
        <n-button type="info" size="large" @click="createVm" style="margin-left: 10px">创建容器</n-button>
      </div>
      <n-data-table
          :columns="createColumns"
          :data="displayContainerList"
          :bordered="false"
          default-expand-all
      />
      <n-pagination :v-model:page="currentPage" :page-count="totalPages"  @change="handlePageChange"  style="margin-left: 80%;margin-top: 50px"/>
    </n-card>
  </div>
  <n-drawer v-model:show="active" :width="800" placement="left">
    <n-drawer-content title="创建容器">
      <n-card title="初始化容器信息">
        <n-steps current="3" style="margin-bottom: 30px;margin-top: 20px">
          <n-step
              title="选择镜像"
              description="选择已有的Docker镜像"
          />
          <n-step
              title="暴露端口"
              description="暴露当前的服务端口"
          />
          <n-step
              title="配置数据卷"
              description="配置容器的数据持久化服务"
          />
        </n-steps>
        <n-form
            ref="formRef"
            :model="form"
            size="large"
            label-placement="left"
        >
          <n-grid :cols="24" :x-gap="24">
            <n-form-item-gi :span="24" label="容器名称" path="inputValue">
              <n-input v-model:value="form.ContainerName" placeholder="请输入容器名称" />
            </n-form-item-gi>
            <n-form-item-gi :span="24" label="容器镜像" path="textareaValue">
              <n-select :options="imagesOptions" v-model:value="form.ContainerImage" placeholder="请选择容器镜像"></n-select>
            </n-form-item-gi>
            <n-form-item-gi :span="24"  label="重启策略">
              <n-select :options="restartPolicyOptions" v-model:value="form.RestartPolicy" placeholder="请选择重启策略"> </n-select>
            </n-form-item-gi>
            <n-form-item-gi
                :span="12"
                v-for="(item, index) in Port_List"
                :key="index"
                :label="`${item.label}端口`"
            >
              <n-input v-model:value="Port_List[index].port" clearable  :placeholder="`请输入当前的${item.label}端口`"/>
            </n-form-item-gi>
            <n-form-item-gi :span="24" style="display: flex;justify-content: center">
              <n-button style="margin-right: 12px" @click="removeItem(index)" type="error" :bordered="false">
                删除
              </n-button>
              <n-button attr-type="button" @click="addItem" type="success" :bordered="false">
                增加
              </n-button>
            </n-form-item-gi>
            <n-form-item-gi
                :span="12"
                v-for="(item, index) in Volume_List"
                :key="index"
                :label="`${item.label}路径`"
            >
              <n-input v-model:value="Volume_List[index].path" clearable  :placeholder="`请输入当前的${item.label}端口`"/>
            </n-form-item-gi>
            <n-form-item-gi :span="24" style="display: flex;justify-content: center">
              <n-button style="margin-right: 12px" @click="removeVolume(index)" type="error" :bordered="false">
                删除
              </n-button>
              <n-button attr-type="button" @click="addVolume" type="success" :bordered="false">
                增加
              </n-button>
            </n-form-item-gi>
            <n-form-item-gi :span="24"  path="textareaValue">
              <n-button type="info" @click="submit" style="width: 100%;margin-top: 20px">提交</n-button>
            </n-form-item-gi>
          </n-grid>
        </n-form>
      </n-card>
    </n-drawer-content>
  </n-drawer>
</div>
</n-card>
</template>

<style scoped>
.search-container {
  display: flex;
  align-items: center;
  justify-content: flex-end;
  margin-bottom: 10px;
}

</style>
