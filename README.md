# myTest - 基于maotu的可视化组态编辑器项目

## 项目简介

本项目基于 [maotu](https://www.npmjs.com/package/maotu) 可视化组态编辑器，前端采用 Vue3 + Vite + TypeScript，后端采用 Node.js 或 Bun 实现设备数据实时变化模拟。支持浮点变量每3秒产生一次随机数，布尔变量每60秒切换一次状态。

## 主要功能

- 可视化组态页面编辑态（edit）：支持组件拖拽、属性/事件/数据绑定、导入导出、保存、预览等功能。
- 组态页面运行态（preview）：实时展示绑定数据的变化，组件状态/颜色/数据显示实时联动。
- 后端数据服务：提供WebSocket接口，推送模拟设备变量数据。

## 目录结构

```
myTest/
├── README.md
├── package.json
├── vite.config.ts
├── tsconfig.json
├── src/
│   ├── main.ts
│   ├── App.vue
│   ├── router/
│   │   └── index.ts
│   ├── views/
│   │   ├── edit.vue
│   │   └── preview.vue
│   └── components/
│       └── ...（自定义组件）
└── server/
    └── index.js
```

## 启动方式

### 安装依赖

```sh
pnpm install
```

### 启动后端数据服务

```sh
pnpm run server
```

### 启动前端

```sh
pnpm run dev
```

## 参考

- maotu官方文档
- example/vue3 项目结构
