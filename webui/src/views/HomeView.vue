<script>
import router from '../router';

export default {
	data: function () {
		return {
			errorMsg: "",
			posts: [],
		}
	},

	methods: {
		async getMyStream() {
			try {
				this.errormsg = null;

				let response = await this.$axios.get("/users/" + localStorage.getItem('token') + "/home");

				if (response.data != null) {
					this.posts = response.data;
				}
			} catch (error) {
				this.errorMsg = this.$utils.errorToString(e);
			}
		},
		async mounted() {
			if (!localStorage.token) {
				this.$router.push('/login');
				return
			}
			await this.getMyStream();
		},

		goToSettingsPage() {
			this.$router.push("/users/" + localStorage.getItem('token') + "/settings");
		}
	}
}
</script>

<template>
	<div class="container-fluid">
		<ErrorMsg v-if="errorMsg" :msg="errorMsg" @close-error="errorMsg = ''"></ErrorMsg>
		<div class="row">
			<post
				v-for="(post,index) in posts"
				:key="index"
				:username="post.username"
				:postId="post.postId"
				:comments="post.comments != nil ? post.comments : []"
				:likes="post.likes != nil ? post.likes : []"
				:pubblication_time="post.pubblication_time"
			/>
		</div>
		<span v-if="posts.length == 0" class="no-posts-text"> There are no posts yet, </span>
		<span v-if="posts.length == 0" class="no-posts-text fw-500 fs-6"> start to follow someone!</span>

		<!-- Aggiunta del pulsante delle impostazioni -->
		<button class="settings-button" @click="goToSettingsPage">
			<i class="fa fa-cog"></i>
		</button>
	</div>
</template>

<style>
.settings-button {
	position: fixed;
	bottom: 20px;
	right: 20px;
	background-color: #fff;
	border: none;
	border-radius: 50%;
	width: 40px;
	height: 40px;
	box-shadow: 0 2px 4px rgba(0, 0, 0, 0.1);
	cursor: pointer;
}

.settings-button i {
	font-size: 20px;
}
</style>