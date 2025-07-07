<script setup lang="ts">
import { RouterView } from 'vue-router';
import { leftAsideStore } from 'maotu';
import 'maotu/dist/style.css';
import { getCurrentInstance } from 'vue';
import { defineAsyncComponent } from 'vue'
const SwitchComponent = defineAsyncComponent(() => import('./components/custom-switch.vue'))
const threeCircle = defineAsyncComponent(  () => import('./components/custom-demo.vue'))
const mybutton = defineAsyncComponent(() => import('./components/button1.vue'))
const TemperatureMeter = defineAsyncComponent(() => import('./components/TemperatureMeter.vue'))
const instance = getCurrentInstance();
instance?.appContext.app.component('customSwitch', SwitchComponent);
instance?.appContext.app.component('threeCircle', threeCircle);
instance?.appContext.app.component('button1', mybutton);
instance?.appContext.app.component('TemperatureMeter', TemperatureMeter);
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
    id: 'customSwitch',
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
  {
    id: 'TemperatureMeter',
    title: '温度计',
    type: 'custom-svg',
    thumbnail: '/svgs/TemperatureMeter.svg',
    props: {
      // width: {
      //   type: 'number',
      //   val: 180,
      //   title: '宽度'
      // },
      // height: {
      //   type: 'number',
      //   val: 300,
      //   title: '高度'
      // },
      // minTemp: {
      //   type: 'number',
      //   val: -20,
      //   title: '最小温度'
      // },
      // maxTemp: {
      //   type: 'number',
      //   val: 100,
      //   title: '最大温度'
      // },
      // currentTemp: {
      //   type: 'number',
      //   val: 20,
      //   title: '当前温度'
      // },
      // unit: {
      //   type: 'select',
      //   val: 'C',
      //   options: [
      //     { value: 'C', label: '摄氏度' },
      //     { value: 'F', label: '华氏度' }
      //   ],
      //   title: '温度单位'
      // },
      // caseFill: {
      //   type: 'color',
      //   val: '#fff',
      //   title: '外壳填充色'
      // },
      // caseStroke: {
      //   type: 'color',
      //   val: '#000',
      //   title: '外壳边框色'
      // },
      // caseStrokeWidth: {
      //   type: 'number',
      //   val: 2,
      //   title: '外壳边框宽度'
      // },
      // bgFill: {
      //   type: 'color',
      //   val: '#eee',
      //   title: '背景填充色'
      // },
      // tempBarFill:{
      //     type:'color',
      //     val:'red',
      //     title:'温度条颜色'
      // },
      // scaleStroke:{
      //     type:'color',
      //     val:'#000',
      //     title:'刻度线颜色'
      // },
      // scaleTextColor:{
      //     type:'color',
      //     val:'#000',
      //     title:'刻度文字颜色'
      // },
      // unitTextColor:{
      //     type:'color',
      //     val:'#000',
      //     title:'单位文字颜色'
      // },
    }
  }

]);

leftAsideStore.registerConfig('vue组件', [
    {
        id: 'button1',
        title: '按钮1',
        type: 'vue',
        thumbnail: '/svgs/button.svg',//缩略图请自己指定一个 图片还是svg都可以
        props: {
            text: {
                type: 'input',
                val: '按钮',
                title: '文本'
            },
            bgColor: {
                type: 'color',
                val: '#44B6E7',
                title: '背景色'
            },
            fontFamily: {
                title: '字体',
                type: 'select',
                val: '黑体',
                options: [
                    {
                        value: '黑体',
                        label: '黑体'
                    },
                    {
                        value: '宋体',
                        label: '宋体'
                    }
                ]
            },
            customInfo: {
                title: '自定义信息',
                type: 'select',
                val: '自定义信息',
                    options: [
                    {
                        value: '页面1',
                        label: '页面1'
                    },
                    {
                        value: '页面2',
                        label: '页面2'
                    }
                ]
            }
        }
    }
]);
</script>

<template>
  <RouterView />
</template>
