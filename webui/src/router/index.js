import {createRouter, createWebHashHistory} from 'vue-router'
import HomeView from '../views/HomeView.vue'
import LoginView from '../views/LoginView.vue'
import ProfileView from '../views/ProfileView.vue'

const router = createRouter({
	history: createWebHashHistory(import.meta.env.BASE_URL),
	routes: [
		{path: '/home', component: HomeView},
		{path: '/login', component: LoginView},
		{path: '/', redirect: '/login'},
		{path: '/users/:user_id', component: ProfileView},
	]
})




export default router
