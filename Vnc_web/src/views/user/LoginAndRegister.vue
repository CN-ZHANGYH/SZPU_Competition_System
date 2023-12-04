<script setup>
import {h, reactive, toRef, toRefs} from "vue";
import { register,login } from "@/api/user";
import router from "@/router";
const data = reactive({
  form: {},
  username: null,
  company: null,
  password: null,
  email: null,
  role: null
});

const {form} = toRefs(data)

function userlogin(){
  login(form.value).then(res => {
    if (res.code == 0){
      localStorage.setItem("user",JSON.stringify(res.data))
      window.$message.success(res.msg)
      router.push({
        path: '/vnc'
      })
    }else {
      window.$message.error(res.msg)
    }
  })
}

function userRegister(){
  register(form.value).then(res => {
    if (res.code == 0){
      window.$message.success(res.msg)
    }else {
      window.$message.error(res.msg)
    }
  })
}

</script>

<template>
  <div class="container">
    <n-card title="深圳职业技术大学区块链训练平台" hoverable embedded>
      <n-tabs
          class="card-tabs"
          default-value="signin"
          size="large"
          animated
          pane-wrapper-style="margin: 0 -4px"
          pane-style="padding-left: 4px; padding-right: 4px; box-sizing: border-box;"
      >
        <n-tab-pane name="signin" tab="登录">
          <n-form :model="form"  label-placement="left" label-width="auto">
            <div style="text-align: center">
              <img src="" >
            </div>
            <n-form-item-row
                label="用户名"
                path="username"
                :rule="{
                    required: true,
                    message: '请输入用户名',
                    trigger: ['input', 'blur']
                  }"
            >
              <n-input v-model:value="form.username" placeholder="请输入用户名"/>
            </n-form-item-row>
            <n-form-item-row
                label="密码"
                path="password"
                :rule="{
                    required: true,
                    message: '请输入用户密码',
                    trigger: ['input', 'blur']
                  }"
            >
              <n-input v-model:value="form.password" placeholder="请输入用户密码"/>
            </n-form-item-row>
          </n-form>
          <n-button type="primary" block strong @click="userlogin">
            登录
          </n-button>
        </n-tab-pane>
        <n-tab-pane name="signup" tab="注册">
          <n-form :model="form" label-placement="left" label-width="auto">
            <n-form-item-row
                label="用户名"
                path="username"
                :rule="{
                    required: true,
                    message: '请输入用户名',
                    trigger: ['input', 'blur']
                  }"

            >
              <n-input v-model:value="form.username" placeholder="请输入用户名"/>
            </n-form-item-row>
            <n-form-item-row
                label="密码"
                path="password"
                :rule="{
                    required: true,
                    message: '请输入用户密码',
                    trigger: ['input', 'blur']
                  }"
            >
              <n-input v-model:value="form.password"  placeholder="请输入用户密码"/>
            </n-form-item-row>
            <n-form-item-row
                label="邮箱"
                path="email"
                :rule="{
                    required: true,
                    message: '请输入邮箱',
                    trigger: ['input', 'blur']
                  }"
            >
              <n-input v-model:value="form.email" placeholder="请输入邮箱"/>
            </n-form-item-row>
            <n-form-item-row
                label="手机"
                path="phone"
                :rule="{
                    required: true,
                    message: '请输入手机号',
                    trigger: ['input', 'blur']
                  }"
            >
              <n-input v-model:value="form.phone" placeholder="请输入手机号"/>
            </n-form-item-row>
          </n-form>
          <n-button type="primary" block strong @click="userRegister">
            注册
          </n-button>
        </n-tab-pane>
      </n-tabs>
    </n-card>
  </div>
</template>

<style scoped>
.card-tabs .n-tabs-nav--bar-type {
  padding-left: 4px;
}

.container {
  display: flex;
  justify-content: center;
  align-items: center;
  height: 100vh;
  background-image: url("@/assets/KIT-21x.png");
  background-size: 100% 100%;
}

.n-card {
  width: 400px;
  height: 430px;
  border-radius: 10px;
  background-color: #f7faff;
  animation: slideIn 0.5s ease-out;
}

@keyframes slideIn {
  0% {
    transform: translateY(-50px);
    opacity: 0;
  }
  100% {
    transform: translateY(0);
    opacity: 1;
  }
}

.card-tabs .n-tabs-nav--bar-type {
  padding-left: 4px;
}

.card-tabs .n-tabs-nav--bar-type .n-tabs-nav__item {
  transition: all 0.3s ease-out;
}

.card-tabs .n-tabs-nav--bar-type .n-tabs-nav__item:hover {
  transform: scale(1.1);
}

.n-tabs-content {
  animation: fadeIn 0.5s ease-out;
}

@keyframes fadeIn {
  0% {
    opacity: 0;
  }
  100% {
    opacity: 1;
  }
}

</style>