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
				:like="post.like != nil ? post.like : []"
				:publication_time="post.publication_time"
			/>
		</div>
		<span v-if="posts.length == 0" class="box-no-posts-text"> No one of your followers has posted yet, or maybe you just need to follow someone!! </span>
	</div>
</template>

<style>
.box-no-posts-container {
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  height: 80vh;
  text-align: center;
}

.box-no-posts-text {
  font-size: 24px;
  font-weight: bold;
  font-family: Arial, sans-serif;
  color: #555555;
  margin-bottom: 8px;
}
</style>