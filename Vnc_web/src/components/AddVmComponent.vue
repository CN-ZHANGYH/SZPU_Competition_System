<template>
  <n-card title="添加虚拟机信息" style="width: 450px;height: 760px;border-radius: 20px;margin-top: 20px" hoverable>
    <div class="form-vm">
      <n-carousel style="border-radius: 20px" autoplay>
        <div class="image-container">
          <img src="https://blog-1304715799.cos.ap-nanjing.myqcloud.com/imgs/202310060216469.png" alt="" class="img">
          <div class="image-overlay"></div>
        </div>
      </n-carousel>
      <div style="display: flex;justify-content: space-between">
        <span>活跃:</span>
        <n-progress
            type="line"
            :percentage="60"
            :indicator-placement="'inside'"
            processing
            status="success"
            :height="24"
            style="margin-bottom: 10px;width: 270px"
        />
      </div>
      <div style="display: flex;justify-content: space-between">
        <span>空闲:</span>
        <n-progress
            type="line"
            :percentage="60"
            :indicator-placement="'inside'"
            processing
            status="default"
            :height="24"
            style="margin-bottom: 10px;width: 270px"
        />
      </div>
      <div style="display: flex;justify-content: space-between">
        <span>压力:</span>
        <n-progress
            type="line"
            :percentage="60"
            :indicator-placement="'inside'"
            processing
            status="warning"
            :height="24"
            style="margin-bottom: 10px;width: 270px"
        />
      </div>
      <div style="display: flex;justify-content: space-between">
        <span>异常:</span>
        <n-progress
            type="line"
            :percentage="60"
            :indicator-placement="'inside'"
            processing
            status="error"
            :height="24"
            style="margin-bottom: 10px;width: 270px"
        />
      </div>
    </div>
    <n-divider />
    <div style="display: flex;justify-content: left">
      <n-form
          ref="formRef"
          :model="model"
          label-placement="left"
          label-width="100px"
          size="medium"
          style="display: flex;flex-direction: column"
      >
        <n-form-item label="名称" path="name" required
                     :rule="{
                      required: true,
                      trigger: ['blur', 'input'],
                      message: '请输入主机名称或备注'
                  }"
        >
          <n-input v-model:value="model.name" placeholder="请输入主机名称"  />
        </n-form-item>
        <n-form-item label="HOST主机" path="host" required
                     :rule="{
                      required: true,
                      trigger: ['blur', 'input'],
                      message: '请输入主机IP地址'
                  }"

        >
          <n-input v-model:value="model.host" placeholder="请输入主机IP地址"  />
        </n-form-item>
        <n-form-item label="VNC密码" path="password" required
                     :rule="{
                      required: true,
                      trigger: ['blur', 'input'],
                      message: '请输入主机IP的VNC密码'
                  }"

        >
          <n-input
              v-model:value="model.password"
              placeholder="请输入主机IP的VNC密码"
          />
        </n-form-item>
        <n-form-item label="VNC端口" path="port" required
                    
        >
          <n-select
              v-model:value="model.port"
              placeholder="请选择远程端口"
              :options="generalOptions"
          />
        </n-form-item>
        <div style="display: flex;justify-content: center;">
          <n-button  @click="handleValidateClick" type="primary" secondary style="margin-right: 20px">添加虚拟机</n-button>
        </div>
      </n-form>
    </div>
  </n-card>
</template>


<script setup>
import {reactive, ref, toRefs} from "vue";
import router from "@/router";
import {addVm, getVmList} from "@/api/manager";
const formRef = ref(null);


const data = reactive({
  generalOptions: [
    {
      label: "6081",
      value: 6081
    },
    {
      label: "6082",
      value: 6082
    },
    {
      label: "6083",
      value: 6083
    },
    {
      label: "6084",
      value: 6084
    },
    {
      label: "6085",
      value: 6085
    },
    {
      label: "6086",
      value: 6086
    },    {
      label: "6087",
      value: 6087
    }
  ]
})

const {generalOptions} = toRefs(data)
const model = ref({
  name: "",
  host: "",
  password: "",
  port: null,
  url: ""
})

const handleValidateClick = () => {
  if (model.value.name == "" || model.value.host == "" || model.value.password == "" || model.value.port == ""){
    return window.$message.error("请填写虚拟机信息")
  }
  addVm(model.value).then(res => {
    console.log(res)
    if (res.code == 200) {
      window.$message.success(res.msg)
      router.go(0)
    }else  {
      window.$message.error("添加失败")
    }
  })
}

const goBack = () => {
  router.push("/manager/vmList.txt")
}

</script>


<style>
.form-vm span {
  font-weight: bold;

}

.image-container {
  margin-left: 18px;
  margin-bottom: 15px;
  position: relative;
  width: 280px; /* 调整为实际大小 */
  height: 200px; /* 调整为实际大小 */
  overflow: hidden;
  border-radius: 25px; /* 添加圆角 */
}

.image-container img {
  width: 280px;
  height: 200px;
  border-bottom: 1px solid;
  border-radius: inherit; /* 继承圆角 */
  transition: transform 0.3s ease;
}

.image-container img:hover {
  transform: scale(1.2);
}

.image-overlay {
  position: absolute;
  top: 0;
  left: 0;
  width: 100%;
  height: 100%;
  pointer-events: none;
  /* 这里可以添加动画样式 */
  animation: your-animation 1s infinite;
}
</style>