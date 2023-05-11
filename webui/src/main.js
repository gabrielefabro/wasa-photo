import {createApp, reactive} from 'vue'
import App from './App.vue'
import router from './router'
import axios from './services/axios.js';
import LoadingSpinner from './components/LoadingSpinner.vue'
import ErrorMsg from './components/ErrorMsg.vue'




import './assets/main.css'
import './assets/login_style.css'
import './assets/dashboard.css'


const app = createApp(App)
app.config.globalProperties.$axios = axios;






app.use(router)
app.mount('#app')