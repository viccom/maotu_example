import { createApp } from 'vue';
import App from './App.vue';
import router from './router';
import 'element-plus/dist/index.css';
import 'maotu/dist/style.css';

const app = createApp(App);
app.use(router);
app.mount('#app');
console.log('Vue app mounted on #app');
