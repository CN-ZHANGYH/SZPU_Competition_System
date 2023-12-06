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
  width: 100%;
  height: 900px;
}
</style>
