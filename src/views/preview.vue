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
let tagKeys: string[] = [];

// 事件回调（如有需要可自定义处理）
const onEventCallBack = (type: string, item_id: string) => {
  console.log('事件回调:', type, item_id);
  // 可根据实际需求处理事件
};

async function fetchDeviceTags() {
  try {
    const res = await fetch('/api/devices');
    if (!res.ok) {
      throw new Error('获取设备信息失败: ' + res.status);
    }
    const devices = await res.json();
    // 收集所有 children.value
    tagKeys = [];
    devices.forEach((dev: any) => {
      if (Array.isArray(dev.children)) {
        dev.children.forEach((child: any) => {
          if (child.value) tagKeys.push(child.value);
        });
      }
    });
  } catch (err) {
    console.error('获取设备信息异常', err);
  }
}

function connectWebSocket() {
  ws = new WebSocket('ws://localhost:3001');
  ws.onopen = () => {
    console.log('WebSocket connected');
  };
  ws.onmessage = (event) => {
    console.log('WebSocket message:', event.data);
    const data = JSON.parse(event.data);
    // 遍历所有标签变量，赋值
    if (MtPreviewRef.value) {
      tagKeys.forEach(tagkey => {
        if (tagkey in data) {
          MtPreviewRef.value.setDevicePointByID(tagkey, data[tagkey]);
        }
      });
    }
  };
  ws.onerror = (err) => {
    console.error('WebSocket error:', err);
  };
  ws.onclose = () => {
    console.log('WebSocket closed');
  };
}

onMounted(async () => {
  // 加载编辑时保存的组态数据
  const json = sessionStorage.getItem('exportJson');
  if (json) {
    MtPreviewRef.value?.setImportJson(JSON.parse(json));
  }
  // 先获取设备标签变量，再连接 WebSocket
  await fetchDeviceTags();
  connectWebSocket();
});
onUnmounted(() => {
  ws?.close();
});
</script>
