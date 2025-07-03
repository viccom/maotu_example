<script setup lang="ts">
import { RouterView } from 'vue-router';
import { leftAsideStore } from 'maotu';
import 'maotu/dist/style.css';
import { getCurrentInstance } from 'vue';
import CustomDemo from './components/custom-demo.vue';
const instance = getCurrentInstance();
instance?.appContext.app.component('custom-demo', CustomDemo);
const modulesFiles = import.meta.glob("/public/svgs/**.svg", { eager: true, as: 'raw' })
const register_config:any = [];
// import svgdemo from '/svgs/demo.svg?raw'; // 引入svg文件

// 注册组件库（可根据实际需求扩展）
for (const key in modulesFiles) {
  //根据路径获取svg文件名
  const name = key.split("/").pop().split(".")[0];
  register_config.push({
    id: name,
    title: name,
    type: 'svg',
    thumbnail: key.replace('/public', ''),
    svg: modulesFiles[key],
    props: {
      fill: {
        type: 'color',
        val: '#FF0000',
        title: '填充色'
      },
      point:{
        type:'number',
        val:0,
        title:'点位'
      },
      // 新增液柱高度控制
      liquidHeight: {
        type: 'number',
        val: 30,              // 默认高度（与 SVG 原始值一致）
        title: '液柱高度',
        min: 0,               // 最小高度（完全收缩）
        max: 140,             // 最大高度（180 - 40，根据容器计算）
      },
      // 新增开关
      isOpen: {
        type: 'switch',
        val: false, // 默认关闭状态
        title: '开关状态'
      },
    }
    // 可选：添加其他属性
  })
}

leftAsideStore.registerConfig('svg文件', register_config);
leftAsideStore.registerConfig('自定义svg', [
{
    id: 'custom-demo',
    title: '自定义svg',
    type: 'custom-svg',
    thumbnail: '/svgs/demo.svg',
    props: {
      circleFill: {
        type: 'color',
        val: '#FF0000',
        title: '圆颜色'
      },
      pathFill1: {
        type: 'color',
        val: '#00FF00',
        title: '线1颜色'
      },
      pathFill2: {
        type: 'color',
        val: '#0000FF',
        title: '线2颜色'
      },
      pathFill3: {
        type: 'color',
        val: '#FFFF00',
        title: '线3颜色'
      },
      showLine2: {
        type: 'switch',
        val: true,
        title: '显示线2'
      }
    }
  }
]);
</script>

<template>
  <RouterView />
</template>
