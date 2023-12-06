<template>
  <n-card style="height: 70px">
    <div style="display: flex;justify-content: space-between;">
      <span style="display: flex;flex-direction: column;font-size: 20px;font-family: 'Californian FB';font-weight: bold">VNC虚拟机管理</span>
    </div>
  </n-card>

  <div style="display: flex;justify-content: space-between ">
    <AddVmComponent/>
    <VmListComponent/>
  </div>
</template>


<script setup>
import router from "@/router";
import AddVmComponent from "@/components/AddVmComponent.vue";
import VmListComponent from "@/components/VmListComponent.vue";
const AddVm = () => {
  router.push("/manager/addVm")
}
</script>
