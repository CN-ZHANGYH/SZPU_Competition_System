<script setup>
// 引入createDiscreteApi
import { createDiscreteApi } from 'naive-ui'

// 获取所有的消息组件
const { message, notification, dialog, loadingBar } = createDiscreteApi(
    ['message', 'dialog', 'notification', 'loadingBar']
)

// 全局挂载windows
window.$message = message;
window.$notification = notification;
window.$dialog = dialog;
window.$loadingBar = loadingBar;
</script>

<template>
  <div></div>
</template>