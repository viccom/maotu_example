<template>
  <div class="temperature-meter">
    <svg
      width="80"
      height="300"
      xmlns="http://www.w3.org/2000/svg"
      xmlns:xlink="http://www.w3.org/1999/xlink"
      version="1.1"
    >
      <!-- 温度计外壳（固化样式） -->
      <rect 
        x="10" 
        y="0" 
        width="70" 
        height="280" 
        rx="8" 
        ry="8" 
        fill="#ffffff" 
        stroke="#000000" 
        stroke-width="2"
      />
      <!-- 温度计内部背景 -->
      <rect 
        x="30" 
        y="30" 
        width="40" 
        height="240" 
        fill="#eeeeee"
      />
      <!-- 动态刻度系统（固化样式） -->
      <g 
        stroke="#000000" 
        stroke-width="1.5" 
        font-family="Arial" 
        font-size="12" 
        fill="#000000" 
        text-anchor="end"
      >
        <template v-for="tick in majorTicks" :key="tick.value">
          <line 
            x1="50" 
            :y1="scaleToY(tick.value)" 
            x2="38" 
            :y2="scaleToY(tick.value)"
          />
          <text 
            x="35" 
            :y="scaleToY(tick.value) + 4"
          >
            {{ tick.label }}
          </text>
        </template>
      </g>
      <!-- 动态次刻度 -->
      <g stroke="#000000" stroke-width="1">
        <line 
          v-for="tick in minorTicks" 
          :key="tick.value"
          x1="50" 
          :y1="scaleToY(tick.value)" 
          x2="44" 
          :y2="scaleToY(tick.value)"
        />
      </g>
      <!-- 动态温度指示条 -->
      <rect 
        x="50"
        :y="Math.min(scaleToY(currentTemp), scaleToY(-40))"
        width="20"
        :height="Math.abs(scaleToY(currentTemp) - scaleToY(-40))"
        fill="red"
      />
      <!-- 温度单位（固化摄氏度） -->
      <text 
        x="50" 
        y="20" 
        font-family="Arial" 
        font-size="14" 
        text-anchor="middle" 
        fill="#000000"
      >
        ℃
      </text>
    </svg>
  </div>
</template>

<script setup lang="ts">
import { computed } from 'vue';

// 只保留currentTemp作为可配置属性
const props = defineProps({
  currentTemp: { 
    type: Number, 
    default: 20,
    validator: (value: number) => {
      return value >= -40 && value <= 100;
    }
  }
});

// 固化温度范围
const minTemp = -40;
const maxTemp = 100;

// 生成主刻度（步长20°C）
const majorTicks = computed(() => {
  const result = [];
  for (let value = -40; value <= 100; value += 20) {
    result.push({ 
      value,
      label: value.toFixed(0)
    });
  }
  return result;
});

// 生成次刻度（步长10°C）
const minorTicks = computed(() => {
  const result = [];
  for (let value = -40; value <= 100; value += 10) {
    if (value % 20 !== 0) { // 避免与主刻度重复
      result.push({ value });
    }
  }
  return result;
});

// 温度值到Y坐标映射
const scaleToY = (temp: number) => {
  const ratio = (temp - minTemp) / (maxTemp - minTemp);
  return 30 + (1 - ratio) * 240;
};

 import { onMounted } from 'vue'
 onMounted(() => {
  console.log('组件挂载完成，当前温度:', props.currentTemp)
  console.log('主刻度:', majorTicks.value)   
  console.log('Y坐标计算:', scaleToY(props.currentTemp))
})
</script>

<style scoped>
.temperature-meter {
  display: inline-block;
  width: 80px;
  height: 300px;
  /* 防止样式冲突 */
  transform: translateZ(0);
  backface-visibility: hidden;
}
</style>
