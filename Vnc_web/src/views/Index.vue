<template>
  <div class="home">
    <n-layout>
      <n-layout-header bordered style="display: flex;justify-content: space-between;padding: 5px;">
        <n-image width="50" height="50"  :src="LogoImage" style="margin-left: 20px"></n-image>
        <n-menu mode="horizontal"  :options="menuOptions" @update:value="handleUpdateValue" />
      </n-layout-header>
      <n-layout-content>
        <router-view></router-view>
      </n-layout-content>

      <n-modal v-model:show="showModal">
        <n-card
            style="width: 350px"
            title="😘😘😘😘😘W😘😘😘😘😘"
            :bordered="false"
            size="huge"
            role="dialog"
            aria-modal="true"
        >
          <n-image :src="PayImage" style="width: 300px;height: 300px;display: flex;justify-content: center"></n-image>
        </n-card>
      </n-modal>

    </n-layout>
  </div>
</template>

<script setup>
import {
  Terminal as Terminal,
  RocketSharp as RocketSharp,
  Newspaper as Newspaper,
  DesktopOutline as DesktopOutline,
  ShieldCheckmark as ShieldCheckmark,
  EyeSharp as EyeSharp,
  Build as Build,
  LogIn as LogIn,
  LogOut as LogOut,
  Person as Person,
  Card as Card
} from "@vicons/ionicons5";
import {h, reactive, ref} from "vue";
import {NIcon} from "naive-ui";
import LogoImage from "@/assets/logo.png"
import PayImage from "@/assets/pay.jpg"
import router from "@/router";
import { onMounted } from "vue";

const showModal = ref(false)
function renderIcon(icon) {
  return () => h(NIcon, null, { default: () => h(icon) });
}

const menuOptions = ref([])

// @ is an alias to /src
const userOptions = [
  {
    label: "首页",
    key: "home",
    icon: renderIcon(Terminal)
  },
  {
    label: "升级",
    key: "dance-dance-dance",
    icon: renderIcon(RocketSharp),
    children: [
      {
        type: "group",
        label: "版本V1.0",
        key: "people",
        children: [
          {
            label: "充值",
            key: "pay",
            icon: renderIcon(ShieldCheckmark)
          },
          {
            label: "文档",
            key: "sheep-man",
            icon: renderIcon(ShieldCheckmark)
          }
        ]
      },
    ]
  },
  {
    label: "个人",
    key: "pernson-data",
    icon: renderIcon(Person),
    children: [
          {
            label: "账户",
            key: "account",
            icon: renderIcon(Card)
          },
          {
            label: "退出",
            key: "logout",
            icon: renderIcon(LogOut)
          }
    ]
  }
];

const adminOptions = reactive([
{
    label: "首页",
    key: "home",
    icon: renderIcon(Terminal)
  },
  {
    label: "环境配置",
    key: "env",
    icon: renderIcon(DesktopOutline),
    children: [
      {
        label: "虚拟机管理",
        key: "manager",
        icon: renderIcon(Build)
      },
      {
        label: "容器管理",
        key: "docker",
        icon: renderIcon(EyeSharp)
      },
    ]
  },
  {
    label: "试题管理",
    key: "markdown",
    icon: renderIcon(Newspaper),
    children: [
      {
        label: "编辑试题",
        key: "md",
        icon: renderIcon(Build)
      },
      {
        label: "试题管理",
        key: "exam",
        icon: renderIcon(EyeSharp)
      },
    ]
  },
  {
    label: "升级",
    key: "dance-dance-dance",
    icon: renderIcon(RocketSharp),
    children: [
      {
        type: "group",
        label: "版本V1.0",
        key: "people",
        children: [
          {
            label: "充值",
            key: "pay",
            icon: renderIcon(ShieldCheckmark)
          },
          {
            label: "文档",
            key: "sheep-man",
            icon: renderIcon(ShieldCheckmark)
          }
        ]
      },
    ]
  },
  {
    label: "个人",
    key: "pernson-data",
    icon: renderIcon(Person),
    children: [
          {
            label: "账户",
            key: "account",
            icon: renderIcon(Card)
          },
          {
            label: "退出",
            key: "logout",
            icon: renderIcon(LogOut)
          }
    ]
  }
])




function handleUpdateValue(key, item) {
  if (key == "home") {
    router.push("/vnc")
  }

  if (key == "md") {
    router.push("/md")
  }

  if (key == "exam") {
    router.push("/exam")
  }

  if (key == "manager"){
    router.push("/manager")
  }

  if (key == "pay"){
    showModal.value = true
  }
  if (key == "logout") {
    router.push("/login")
    localStorage.removeItem("user")
  }
  if (key == "account") {
    router.push("/profile")
  }
  if (key == "docker") {
    router.push("/docker")
  }
}


onMounted(() => {
  const role = JSON.parse(localStorage.getItem("user")).Role
  console.log(role)
  if (role == 1){
    menuOptions.value = userOptions
  }else {
    menuOptions.value = adminOptions
  }
})
</script>
