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
import { ref, onMounted } from 'vue';
import { MtEdit, type IExportJson } from 'maotu';
import { useRouter } from 'vue-router';

const router = useRouter();

// 加载设备变量信息（通过接口获取）
const deviceInfo = ref([]);

onMounted(async () => {
  try {
    const res = await fetch('/api/devices');
    if (res.ok) {
      deviceInfo.value = await res.json();
    } else {
      console.error('获取设备信息失败', res.status);
    }
  } catch (err) {
    console.error('获取设备信息异常', err);
  }
});

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

