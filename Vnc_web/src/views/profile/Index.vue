<script setup>
import {reactive, ref, toRefs} from "vue";
import router from "@/router";
import image from '@/assets/avatar20.png'
import {GetUser} from "@/utils/getUser";
import {getUserInfo, UpdateUserInfo} from "@/api/user";


const data = reactive({
    form: {},
    username: null,
    password: null,
    email: null
})

const {form} = toRefs(data)

function logout(){
    router.push({
        path: "/"
    })
}

function update(){
  UpdateUserInfo(form.value).then(res => {
    if (res.code == 0){
      window.$message.success(res.msg)
    }else {
      window.$message.error(res.msg)
      return
    }
  }).catch( error => {
    console.log(error)
  })
  localStorage.removeItem("user")
  router.push("/login")
  window.$message.success("请重新登录")
}

const username = ref(GetUser().Username)
getUserInfo({username: username.value}).then(res => {
  console.log(form.value)
  form.value = res.data
})

</script>

<template>
<n-card hoverable>
    <n-form ref="formRef" :model="form" label-placement="left" label-width="auto">
    <n-form-item label="头像">
            <n-avatar
            :size="60"
            :src="image"
            />
        </n-form-item>
        <n-divider />
        <n-form-item label="账号">
            <n-input v-model:value="form.Username" disabled/>
        </n-form-item>
        <n-form-item label="密码">
            <n-input v-model:value="form.Password" disabled/>
        </n-form-item>
        <n-form-item label="邮箱">
            <n-input v-model:value="form.Email"/>
        </n-form-item>
      <n-form-item label="手机号码">
        <n-input v-model:value="form.Phone"/>
      </n-form-item>
      <n-form-item label="角色">
        <div><n-tag type="primary" v-if="form.Role == 1">普通用户</n-tag></div>
        <div><n-tag type="primary" v-if="form.Role == 2">管理员</n-tag></div>
      </n-form-item>
    </n-form>

    <div style="text-align: center;margin-top: 50px;">
        <n-button type="success" style="margin-right: 30px;" @click="update">更新个人信息</n-button>
        <n-button type="error" @click="logout">退出登录</n-button>
    </div>
</n-card>
</template>

<style scoped>

</style>