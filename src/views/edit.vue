<template>
  <div class="w-1/1 h-100vh">
    <mt-edit
      ref="MtEditRef"
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
const MtEditRef = ref<InstanceType<typeof MtEdit>>(); // 新增

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

  // 加载画布配置
  const viewConfig = await loadViewConfig();
  if (viewConfig && MtEditRef.value) {
    const res = MtEditRef.value.setImportJson(viewConfig);
    if (res) {
      console.log('画布配置加载成功');
    }
  } else if (MtEditRef.value) {
    // 可选：加载默认配置
    const defaultCfg = {
      canvasCfg: {
        width: 1920,
        height: 1080,
        scale: 1,
        color: '',
        img: '',
        guide: true,
        adsorp: true,
        adsorp_diff: 5,
        transform_origin: { x: 0, y: 0 },
        drag_offset: { x: 0, y: 0 }
      },
      gridCfg: { enabled: true, align: true, size: 10 },
      json: [],
      eventList: []
    };
    MtEditRef.value.setImportJson(defaultCfg);
    console.log('加载默认画布配置');
  }
});

// 从后端接口 /api/loadview 获取画布配置
const loadViewConfig = async () => {
  try {
    const res = await fetch('/api/loadview');
    if (res.ok) {
      const viewConfig = await res.json();
      console.log('加载的画布配置:', viewConfig);
      return viewConfig;
    } else {
      console.error('加载画布配置失败', res.status);
      return null;
    }
  } catch (err) {
    console.error('加载画布配置异常', err);
    return null;
  }
};

const onPreviewClick = (exportJson: IExportJson) => {
  sessionStorage.setItem('exportJson', JSON.stringify(exportJson));
  const routeUrl = router.resolve({
    name: 'preview'
  });
  window.open(routeUrl.href, '_blank');
};
const onSaveClick = (e: IExportJson) => {
  console.log(e, '这是要保存的数据');
//后端接口 /api/saveview 代码
// app.post('/api/saveview', express.json(), (req, res) => {
//   fs.writeFileSync(VIEW_JSON_PATH, JSON.stringify(req.body, null, 2));
//   res.json({ success: true, message: '配置已保存' });
// });
  // 发送 POST 请求保存数据
  fetch('/api/saveview', {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json'
    },
    body: JSON.stringify(e)
  })
    .then(res => res.json())
    .then(data => {
      if (data.success) {
        console.log('保存成功');
      } else {
        console.error('保存失败:', data.message);
      }
    })
    .catch(err => {
      console.error('保存异常', err);
    });
};
const onReturnClick = () => {
  router.go(-1);
};

const onItemResizeDone = (e: any) => {
  console.log('onItemResizeDone', e);
};
</script>

