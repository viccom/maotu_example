<template>
  <div class="w-1/1 h-100vh">
    <mt-preview
      ref="MtPreviewRef"
      @on-event-call-back="onEventCallBack"
    ></mt-preview>
  </div>
</template>
<script setup lang="ts">
import { ref, onMounted, onUnmounted } from 'vue';
import { MtPreview } from 'maotu';

const MtPreviewRef = ref<InstanceType<typeof MtPreview>>();
let ws: WebSocket | null = null;

// 事件回调（如有需要可自定义处理）
const onEventCallBack = (type: string, item_id: string) => {
  console.log('事件回调:', type, item_id);
  // 可根据实际需求处理事件
};

onMounted(() => {
  // 加载编辑时保存的组态数据
  const json = sessionStorage.getItem('exportJson');
  if (json) {
    MtPreviewRef.value?.setImportJson(JSON.parse(json));
  }
  // 连接后端WebSocket服务
  ws = new WebSocket('ws://localhost:3001');
  ws.onopen = () => {
    console.log('WebSocket connected');
  };
  ws.onmessage = (event) => {
    console.log('WebSocket message:', event.data);
    const data = JSON.parse(event.data);
    // 通过maotu API更新组件属性，变量名需与组态绑定一致
    if (MtPreviewRef.value) {
      if ('floatVar' in data) {
        MtPreviewRef.value.setDevicePointByID('floatVar', data.floatVar);
      }
      if ('boolVar' in data) {
        MtPreviewRef.value.setDevicePointByID('boolVar', data.boolVar);
      }
      // 如有其它变量可继续扩展
    }
  };
  ws.onerror = (err) => {
    console.error('WebSocket error:', err);
  };
  ws.onclose = () => {
    console.log('WebSocket closed');
  };
});
onUnmounted(() => {
  ws?.close();
});
</script>
