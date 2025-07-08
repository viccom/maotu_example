<template>
  <div class="temperature-meter">
    <svg :width="props.width" :height="props.height" viewBox="0 0 180 300" xmlns="http://www.w3.org/2000/svg">
      <!-- 温度计外壳 -->
      <rect 
        x="10" 
        y="0" 
        width="70" 
        height="280" 
        rx="8" 
        ry="8" 
        :fill="props.caseFill" 
        :stroke="props.caseStroke" 
        :stroke-width="props.caseStrokeWidth"
      />
      <!-- 温度计内部背景 -->
      <rect 
        x="30" 
        y="30" 
        width="40" 
        height="240" 
        :fill="props.bgFill"
      />
      <!-- 动态刻度系统 -->
      <g 
        :stroke="props.scaleStroke" 
        :stroke-width="props.majorScaleWidth" 
        font-family="Arial" 
        :font-size="props.scaleFontSize" 
        :fill="props.scaleTextColor" 
        text-anchor="end"
      >
        <template v-for="tick in majorTicks" :key="tick.value">
          <line 
            :x1="props.scaleLineX" 
            :y1="scaleToY(tick.value)" 
            :x2="props.scaleLineX - props.scaleLineLength" 
            :y2="scaleToY(tick.value)"
          />
          <text 
            :x="props.scaleTextX" 
            :y="scaleToY(tick.value) + 4"
          >
            {{ tick.label }}
          </text>
        </template>
      </g>
      <!-- 动态次刻度 -->
      <g :stroke="props.scaleStroke" :stroke-width="props.minorScaleWidth">
        <line 
          v-for="tick in minorTicks" 
          :key="tick.value"
          :x1="props.scaleLineX" 
          :y1="scaleToY(tick.value)" 
          :x2="props.scaleLineX - props.scaleLineLength/2" 
          :y2="scaleToY(tick.value)"
        />
      </g>
      <!-- 动态温度指示条 -->
      <rect 
        x="50"
        :y="Math.min(scaleToY(props.currentTemp), scaleToY(props.minTemp))"
        width="20"
        :height="Math.abs(scaleToY(props.currentTemp) - scaleToY(props.minTemp))"
        :fill="props.tempBarFill"
      />
      <!-- 动态温度单位 -->
      <text 
        x="50" 
        y="20" 
        font-family="Arial" 
        :font-size="props.unitFontSize" 
        text-anchor="middle" 
        :fill="props.unitTextColor"
      >
        {{ unitSymbol }}
      </text>
    </svg>
  </div>
</template>

<!--  -->
<script setup lang="ts">
import { computed } from 'vue';

const props = defineProps({
  // 尺寸控制
  width: { type: Number, default: 80 },
  height: { type: Number, default: 280 },
  
  // 温度范围
  minTemp: { type: Number, default: -20 },
  maxTemp: { type: Number, default: 100 },
  currentTemp: { type: Number, default: 20 },
  unit: { type: String, default: 'C' }, // 'C' or 'F'
  
  // 颜色配置
  caseFill: { type: String, default: '#fff' },
  caseStroke: { type: String, default: '#000' },
  caseStrokeWidth: { type: Number, default: 2 },
  bgFill: { type: String, default: '#eee' },
  tempBarFill: { type: String, default: 'red' },
  scaleStroke: { type: String, default: '#000' },
  scaleTextColor: { type: String, default: '#000' },
  unitTextColor: { type: String, default: '#000' },
  
  // 刻度样式
  majorScaleWidth: { type: Number, default: 1.5 },
  minorScaleWidth: { type: Number, default: 1 },
  scaleFontSize: { type: Number, default: 12 },
  unitFontSize: { type: Number, default: 14 },
  scaleLineX: { type: Number, default: 50 },
  scaleLineLength: { type: Number, default: 12 },
  scaleTextX: { type: Number, default: 35 }
});

const unitSymbol = computed(() => props.unit === 'C' ? '℃' : '℉');

const stepMajor = computed(() => (props.maxTemp - props.minTemp) / 5);
const stepMinor = computed(() => (props.maxTemp - props.minTemp) / 10);

// 修改 majorTicks 的生成方式为固定步长 20
const majorTicks = computed(() => {
  const result = [];
  // 从 minTemp 向下取整到最近的 20 的倍数开始，向上生成直到 maxTemp
  let start = Math.floor(props.minTemp / 20) * 20;
  for (let value = start; value <= props.maxTemp; value += 20) {
    let label = value.toFixed(0);
    if (props.unit === 'F') {
      label = ((value * 9/5) + 32).toFixed(0);
    }
    result.push({ value, label });
  }
  return result;
});

// 修改 minorTicks，这里设置为每 20 一个次刻度（也可设为 10）
const minorTicks = computed(() => {
  const result = [];
  let start = Math.floor(props.minTemp / 10) * 10;
  for (let value = start; value <= props.maxTemp; value += 10) {
    // 如果需要次刻度为 10，则可以在这里加 value + 10 等
    result.push({ value });
  }
  return result;
});

// 温度值到Y坐标映射
const scaleToY = (temp: number) => {
  const ratio = (temp - props.minTemp) / (props.maxTemp - props.minTemp);
  return 30 + (1 - ratio) * 240;
};
</script>

<style scoped>
.temperature-meter {
  display: inline-block;
}
</style>
