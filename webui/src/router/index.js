import {createRouter, createWebHashHistory} from 'vue-router'
import HomeView from '../views/HomeView.vue'
import LoginView from '../views/LoginView.vue'
import ProfileView from '../views/ProfileView.vue'
import SettingsView from '../views/SettingsView.vue'
import SearchView from '../views/SearchView.vue'

const router = createRouter({
	history: createWebHashHistory(import.meta.env.BASE_URL),
	routes: [
		{path: '/home', component: HomeView},
		{path: '/login', component: LoginView},
		{path: '/', redirect: '/login'},
		{path: '/users/:user_id', component: ProfileView},
		{path: '/search', component: SearchView},
		{path: '/users/:user_id/settings', component: SettingsView}
	]
})




export default router
