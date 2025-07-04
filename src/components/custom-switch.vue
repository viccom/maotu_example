<template>
  <svg
    width="120" 
    height="60"
    xmlns="http://www.w3.org/2000/svg"
    xmlns:xlink="http://www.w3.org/1999/xlink"
    version="1.1"
  >
    <!-- 背景轨道 -->
    <rect 
      id="rect1"
      :fill="props.trackColor"      
      rx="30" 
      ry="30"
      width="100%" 
      height="100%"
    />
    <!-- 开关按钮（动态位置） -->
    <circle
      id="circle1"
      :fill="props.thumbColor"
      :cx="props.isOn ? 90 : 30"
      cy="30"
      r="25"
    />
    <!-- 状态文字 -->
    <text
      id="text1"
      x="60"
      y="35"
      text-anchor="middle"
      font-size="14"
      fill="#FFF"
    >
      {{ props.isOn ? 'ON' : 'OFF' }}
    </text>
  </svg>
</template>

<script setup lang="ts">
const props = defineProps({
  id: String,
  isOn: Boolean,  // 正确写法
  trackColor: String, // 正确写法
  thumbColor: String,
})
</script>


