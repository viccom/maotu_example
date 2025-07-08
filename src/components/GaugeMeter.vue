<template>
  <div class="gauge-meter">
    <svg
      width="200"
      height="200"
      xmlns="http://www.w3.org/2000/svg"
      xmlns:xlink="http://www.w3.org/1999/xlink"
      version="1.1"
    >
      <!-- 外圆 -->
      <circle cx="100" cy="100" r="95" fill="#ffffff" stroke="#000000" stroke-width="2" />
      
      <!-- 刻度背景 -->
      <path
        d="M 20 100 A 80 80 0 0 1 180 100"
        fill="none"
        stroke="#eeeeee"
        stroke-width="15"
      />
      
      <!-- 颜色分区 -->
      <path
        d="M 20 100 A 80 80 0 0 1 100 20"
        fill="none"
        stroke="red"
        stroke-width="15"
      />
      <path
        d="M 100 20 A 80 80 0 0 1 180 100"
        fill="none"
        stroke="green"
        stroke-width="15"
      />
      
      <!-- 动态刻度 -->
      <path
        :d="`M 20 100 A 80 80 0 0 1 ${calculateArcX(currentValue)} ${calculateArcY(currentValue)}`"
        fill="none"
        stroke="#42b983"
        stroke-width="15"
      />
      
      <!-- 刻度线 -->
      <g stroke="#000000" stroke-width="2">
        <line
          v-for="tick in ticks"
          :key="tick.value"
          :x1="calculateTickX(tick.value, 90)"
          :y1="calculateTickY(tick.value, 90)"
          :x2="calculateTickX(tick.value, 80)"
          :y2="calculateTickY(tick.value, 80)"
        />
      </g>
      
      <!-- 刻度数字 -->
      <g font-family="Arial" font-size="12" text-anchor="middle" fill="#000000">
        <text
          v-for="tick in ticks"
          :key="tick.value"
          :x="calculateTickX(tick.value, 70)"
          :y="calculateTickY(tick.value, 70)"
          dy="4"
        >
          {{ tick.value }}
        </text>
      </g>
      
      <!-- 指针 -->
      <line
        x1="100"
        y1="100"
        :x2="calculateTickX(currentValue, 70)"
        :y2="calculateTickY(currentValue, 70)"
        stroke="black"
        stroke-width="3"
      />
      
      <!-- 中心圆 -->
      <circle cx="100" cy="100" r="5" fill="black" />
      
      <!-- 数值显示 -->
      <text
        x="100"
        y="80"
        font-family="Arial"
        font-size="16"
        text-anchor="middle"
        fill="#000000"
      >
        {{ currentValue }}
      </text>
    </svg>
  </div>
</template>

<script setup lang="ts">
import { computed } from 'vue';

const props = defineProps({
  currentValue: {
    type: Number,
    default: 50,
    validator: (value: number) => value >= 0 && value <= 100,
  },
});

const ticks = computed(() => {
  const result = [];
  for (let value = 0; value <= 100; value += 25) { // 调整刻度间隔为25
    result.push({ value });
  }
  return result;
});

const calculateArcX = (value: number) => {
  const angle = ((value / 100) * Math.PI) + Math.PI; // 调整角度起始位置为左边
  return 100 + Math.cos(angle) * 80;
};

const calculateArcY = (value: number) => {
  const angle = ((value / 100) * Math.PI) + Math.PI; // 调整角度起始位置为左边
  return 100 + Math.sin(angle) * 80;
};

const calculateTickX = (value: number, radius: number) => {
  const angle = ((value / 100) * Math.PI) + Math.PI; // 调整角度起始位置为左边
  return 100 + Math.cos(angle) * radius;
};

const calculateTickY = (value: number, radius: number) => {
  const angle = ((value / 100) * Math.PI) + Math.PI; // 调整角度起始位置为左边
  return 100 + Math.sin(angle) * radius;
};
</script>

<style scoped>
.gauge-meter {
  display: inline-block;
  width: 200px;
  height: 200px;
  transform: translateZ(0);
  backface-visibility: hidden;
}
</style>
