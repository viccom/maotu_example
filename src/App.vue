<script setup lang="ts">
import { RouterView } from 'vue-router';
import { leftAsideStore } from 'maotu';
import 'maotu/dist/style.css';
import { getCurrentInstance } from 'vue';
import { defineAsyncComponent } from 'vue'
const SwitchComponent = defineAsyncComponent(() => import('./components/custom-switch.vue'))
const threeCircle = defineAsyncComponent(  () => import('./components/custom-demo.vue'))
const instance = getCurrentInstance();
instance?.appContext.app.component('custom-s1', SwitchComponent);
instance?.appContext.app.component('threeCircle', threeCircle);
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
      }
    },
    // 可选：添加其他属性

  })
}

leftAsideStore.registerConfig('svg文件', register_config);
leftAsideStore.registerConfig('自定义svg', [

// 可选：添加其他组件
  {
    id: 'custom-s1',
    title: '自定义开关',
    type: 'custom-svg',
    thumbnail: '/svgs/switch.svg', // 可选缩略图
     props: {
      isOn: {
        type: 'switch',
        val: false,
        title: '开关状态'
      },
      trackColor: {
        type: 'color',
        val: '#CCCCCC',
        title: '轨道颜色'
      },
      thumbColor: {
        type: 'color',
        val: '#42b983',
        title: '按钮颜色'
      }
    }
  },
{
    id: 'threeCircle',
    title: '三色线段圆形',
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
  },

]);
// leftAsideStore.registerConfig('自定义svg', [
// {
//     id: 'custom-demo',
//     title: '三色线段圆形',
//     type: 'custom-svg',
//     thumbnail: '/svgs/demo.svg',
//     props: {
//       circleFill: {
//         type: 'color',
//         val: '#FF0000',
//         title: '圆颜色'
//       },
//       pathFill1: {
//         type: 'color',
//         val: '#00FF00',
//         title: '线1颜色'
//       },
//       pathFill2: {
//         type: 'color',
//         val: '#0000FF',
//         title: '线2颜色'
//       },
//       pathFill3: {
//         type: 'color',
//         val: '#FFFF00',
//         title: '线3颜色'
//       },
//       showLine2: {
//         type: 'switch',
//         val: true,
//         title: '显示线2'
//       }
//     }
//   },
// // 可选：添加其他组件
//   {
//     id: 'custom-switch',
//     title: '自定义开关',
//     type: 'custom-svg',
//     thumbnail: '/svgs/custom-switch.svg', // 可选缩略图
//     props: {
//       isOn: {
//         type: 'switch',
//         val: false,
//         title: '开关状态'
//       },
//       trackColor: {
//         type: 'color',
//         val: '#CCCCCC',
//         title: '轨道颜色'
//       },
//       thumbColor: {
//         type: 'color',
//         val: '#42b983',
//         title: '按钮颜色'
//       }
//     }
//   }

// ]);
</script>

<template>
  <RouterView />
</template>
