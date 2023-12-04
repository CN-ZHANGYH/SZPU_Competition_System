<template>
  <iframe class="vnc-class" :src="vmUrl">
  </iframe>

</template>


<script setup>
defineProps({
  vmUrl: String
})

</script>


<style scoped>
.vnc-class {
  width: 1360px;
  height: 900px;
}
</style>
