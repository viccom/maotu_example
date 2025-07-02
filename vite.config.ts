import { defineConfig } from 'vite';
import vue from '@vitejs/plugin-vue';
import * as path from 'path';
import { fileURLToPath } from 'url';

const __dirname = path.dirname(fileURLToPath(import.meta.url));

export default defineConfig({
  build: {
    outDir: './goserver/dist', // 指定输出目录（默认是 'dist'）
    emptyOutDir: true,     // 构建前清空目录（可选）
  },
  plugins: [vue()],
  resolve: {
    alias: {
      '@': path.resolve(__dirname, 'src'),
    },
  },
  server: {
    port: 5673,
        proxy: {
      '/api': {
        target: 'http://127.0.0.1:3000',
        changeOrigin: true,
        // rewrite: (path) => path.replace(/^\/api/, '/api'), // 通常无需重写
      }
    },
    host: '0.0.0.0' // 明确监听所有地址
  }
});
