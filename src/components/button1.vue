<template>
    <div style="width: 100%; height: 100%">
      <button class="my-button" :data-custom-info="props.customInfo" style="width: 100%; height: 100%">{{ props.text }}</button>
    </div>
</template>
<!-- 导出属性 -->
<script setup lang="ts">
  const props = defineProps({
    text: String,
    bgColor: String,
    fontFamily: String,
    customInfo: String
  });
  </script>
  <!-- 样式 -->
  <style scoped>
  .my-button {
    background-color: v-bind('props.bgColor');
    font-family: v-bind('props.fontFamily');
  }
  </style>