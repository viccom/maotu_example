<script setup lang="ts">
import { RouterView } from 'vue-router';
import { leftAsideStore } from 'maotu';

// 注册组件库（可根据实际需求扩展）
leftAsideStore.registerConfig('设备组件', [
  {
    id: 'demo',
    title: '演示svg文件',
    type: 'svg',
    thumbnail: '/svgs/demo.svg',
    props: {
      fill: {
        type: 'color',
        val: '#FF0000',
        title: '填充色'
      }
    }
  }
  // ...可继续添加自定义组件
]);
</script>

<template>
  <RouterView />
</template>
