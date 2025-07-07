<template>
  <div style="margin: 40px; display: flex; gap: 20px;">
    <!-- 左侧 TemperatureMeter 组件 -->
    <div style="flex: 1;">
      <TemperatureMeter
        :width="180"
        :height="360"
        :min-temp="-40"
        :max-temp="100"
        :current-temp="currentTemp"
        unit="C"
      />
      <div style="margin-top: 40px;">
        <input type="number" style="width: 60px;" v-model.number="inputTemp" placeholder="输入温度值" @change="updateCurrentTemp"/>
      </div>
    </div>
    <!-- 右侧 custom-switch 组件 -->
    <div style="flex: 1;">
      <custom-switch 
        v-model:isOn="switchState" 
        :track-color="trackColor" 
        :thumb-color="thumbColor"
      />
      <!-- 新增：测试输入框 -->
      <div style="margin-top: 20px;">
        <label>轨道颜色：</label>
        <input type="color" v-model="trackColor" />
      </div>
      <div style="margin-top: 10px;">
        <label>按钮颜色：</label>
        <input type="color" v-model="thumbColor" />
      </div>
      <div style="margin-top: 10px;">
        <label>当前值：</label>
        <input type="number" style="width: 60px;" v-model="switchState" />

      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue';
import { getCurrentInstance } from 'vue';
import { defineAsyncComponent } from 'vue';

// 动态导入 TemperatureMeter 组件
const instance = getCurrentInstance();
const TemperatureMeter = defineAsyncComponent(() => import('@/components/TemperatureMeter.vue'));
instance?.appContext.app.component('TemperatureMeter', TemperatureMeter);

// 新增：动态导入 custom-switch 组件
const CustomSwitch = defineAsyncComponent(() => import('@/components/custom-switch.vue'));
instance?.appContext.app.component('customSwitch', CustomSwitch);

const currentTemp = ref(28);
const inputTemp = ref<number | null>(null);

// 新增：custom-switch 相关状态和属性
const switchState = ref(false); // 开关状态
const trackColor = ref('#CCCCCC'); // 轨道颜色
const thumbColor = ref('#42b983'); // 按钮颜色

const updateCurrentTemp = () => {
  if (inputTemp.value !== null && !isNaN(inputTemp.value)) {
    currentTemp.value = inputTemp.value;
  }
};
</script>