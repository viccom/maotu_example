<template>
  <div style="margin: 40px">
    <TemperatureMeter
      :width="180"
      :height="300"
      :min-temp="-20"
      :max-temp="100"
      :current-temp="currentTemp"
      unit="C"
    />
    <div style="margin-top: 40px;">
      <input type="number" style="width: 60px;" v-model.number="inputTemp" placeholder="输入温度值" @change="updateCurrentTemp"/>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue';
import TemperatureMeter from '@/components/TemperatureMeter.vue';

const currentTemp = ref(50);
const inputTemp = ref<number | null>(null);

const updateCurrentTemp = () => {
  if (inputTemp.value !== null && !isNaN(inputTemp.value)) {
    currentTemp.value = inputTemp.value;
  }
};
</script>