import {
    createApp,
    reactive
} from 'vue'
import App from './App.vue'
import router from './router'
import axios from './services/axios.js';
import ErrorMsg from './components/ErrorMsg.vue'
import LoadingSpinner from './components/LoadingSpinner.vue'
import Navbar from './components/Navbar.vue'
import Post from './components/Post.vue'
import UserMiniCard from './components/UserMiniCard.vue'
import LikeModal from './components/LikeModal.vue'
import CommentModal from './components/CommentModal.vue'
import PostComment from './components/PostComment.vue'

import './assets/dashboard.css'
import './assets/main.css'

const app = createApp(App)
app.config.globalProperties.$axios = axios;

app.component("ErrorMsg", ErrorMsg);
app.component("LoadingSpinner", LoadingSpinner);
app.component("Navbar", Navbar);
app.component("Post", Post);
app.component("UserMiniCard", UserMiniCard);
app.component("LikeModal", LikeModal);
app.component("CommentModal", CommentModal);
app.component("PostComment", PostComment);
app.use(router)
app.mount('#app')