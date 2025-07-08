// src/shims-vue.d.ts
declare module '*.svg' {
  const content: string;
  export default content;
}

// 支持 ?raw 导入
declare module '*.svg?raw' {
  const content: string;
  export default content;
}
// 支持 ?url 导入
declare module '*.vue' {
  import { DefineComponent } from 'vue'
  const component: DefineComponent<{}, {}, any>
  export default component
}