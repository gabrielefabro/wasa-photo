<script>
export default {
	data: function () {
		return {
			errorMsg: null,
			posts: [],
		}
	},

	methods: {
		
		async getMyStream() {  
			try {
				this.errormsg = null
				let response = await this.$axios.get("/users/" + localStorage.getItem('token') + "/home")

				if (response.data != null){
					this.posts = response.data
				}
				
			} catch (e) {
				this.errormsg = e.toString()
			}
		}
	},

	mounted() {
		if (!localStorage.token) {
			this.$router.push('/login');
			return
		}
		this.getMyStream();
	}

}
</script>

<template>
	<div class="container-fluid">
		<ErrorMsg v-if="errorMsg" :msg="errorMsg" @close-error="errorMsg = ''"></ErrorMsg>
		<div class="row">
			<Post
				v-for="(post,index) in posts"
				:key="index"
				:user_id="post.user_id"
				:post_id="post.post_id"
				:comments="post.comment != nil ? post.comment : []"
				:likes="post.like != nil ? post.like : []"
				:pubblication_time="post.pubblication_time"
			/>
		</div>
		<span v-if="posts.length == 0" class="no-posts-text"> There are no posts yet, </span>
		<span v-if="posts.length == 0" class="no-posts-text fw-500 fs-6"> start to follow someone!</span>
	</div>
</template>

<style>
</style>