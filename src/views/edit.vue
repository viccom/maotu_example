<template>
  <div class="w-1/1 h-100vh">
    <mt-edit
      :use-thumbnail="true"
      :device-info="deviceInfo"
      @on-preview-click="onPreviewClick"
      @on-return-click="onReturnClick"
      @on-save-click="onSaveClick"
      @on-item-resize-done="onItemResizeDone"
    ></mt-edit>
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue';
import { MtEdit, type IExportJson } from 'maotu';
import { useRouter } from 'vue-router';

const router = useRouter();

// 设备变量信息，参考 bind-device/edit.vue
const deviceInfo = ref([
  {
    label: '设备A',
    value: 'devA',
    children: [
      { label: '浮点变量', value: 'floatVar' },
      { label: '开关变量', value: 'boolVar' }
    ]
  },
  {
    label: '设备B',
    value: 'devB',
    children: [
      { label: '温度', value: 'temp' },
      { label: '湿度', value: 'humi' }
    ]
  }
]);

const onPreviewClick = (exportJson: IExportJson) => {
  sessionStorage.setItem('exportJson', JSON.stringify(exportJson));
  const routeUrl = router.resolve({
    name: 'preview'
  });
  window.open(routeUrl.href, '_blank');
};
const onSaveClick = (e: IExportJson) => {
  console.log(e, '这是要保存的数据');
};
const onReturnClick = () => {
  router.go(-1);
};

const onItemResizeDone = (e: any) => {
  console.log('onItemResizeDone', e);
};
</script>

<style scoped></style>
