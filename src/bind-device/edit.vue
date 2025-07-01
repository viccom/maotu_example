<template>
    <div class="w-1/1 h-100vh">
        <mt-edit ref="MtEditRef" :device-info="test_device_info" @on-return-click="onReturnClick"
            @on-preview-click="onPreviewClick"></mt-edit>
    </div>
</template>
<script setup lang="ts">
import type { IExportJson } from 'maotu';
import { MtEdit } from 'maotu';
import { onMounted, ref } from 'vue';
import { useRouter } from 'vue-router';
const router = useRouter();
const MtEditRef = ref<InstanceType<typeof MtEdit>>();
const test_device_info = ref([
    {
      label: '设备1',
      value: 'd1',
      children: [
        {
          label: '测试点位1',
          value: 'abc'
        },
        {
          label: '测试点位2',
          value: 'def'
        }
      ]
    },
    {
      label: '设备2',
      value: 'd2',
      children: [
        {
          label: '测试点位1',
          value: 'qwe'
        },
        {
          label: '测试点位2',
          value: 'asd'
        }
      ]
    }
  ]);
const onReturnClick = () => {
    router.go(-1);
};
const onPreviewClick = (exportJson: IExportJson) => {
    sessionStorage.setItem('exportJson', JSON.stringify(exportJson));
    const routeUrl = router.resolve({
        name: 'bind-device-preview'
    });
    window.open(routeUrl.href, '_blank');
};
onMounted(() => {
    const res = MtEditRef.value?.setImportJson({
        canvasCfg: {
            width: 1920,
            height: 1080,
            scale: 1,
            color: '',
            img: '',
            guide: true,
            adsorp: true,
            adsorp_diff: 5,
            transform_origin: {
                x: 0,
                y: 0
            },
            drag_offset: {
                x: 0,
                y: 0
            }
        },
        gridCfg: {
            enabled: true,
            align: true,
            size: 10
        },
        json: [],
        eventList: []
    });
    if (res) {
        console.log('加载成功');
    }
});
</script>

<style scoped></style>