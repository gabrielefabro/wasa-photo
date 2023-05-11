<script>
import FloatingNavbar from '@/components/FloattingNavbar.vue'
export default {
	data: function () {
		return {
			errormsg: "",
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
			
			} catch (error) {
				this.errormsg = error.toString()
			}
        },

		addLike(data) {
			this.posts.forEach(post => {
				if (post.postId == data.postId) {
					post.liked = data.liked;
					post.likesCount++;
				}
			});
		},

		async deletePost(postID) {
			const index = this.posts.findIndex(post => post.postId == postId && post.user.userId == localStorage.userId);
			try {
				await this.$axios.delete(`users/${localStorage.userId}/posts/${postId}`)
				this.posts.splice(index, 1);
			} catch (e) {
				this.errorMsg = error.toString();
			}
		},

    async mounted() {
        await this.getMyStream()
    }
    }

}
</script>

<template>
	<div class="container-fluid">
		<ErrorMsg v-if="errormsg" :msg="errormsg"></ErrorMsg>

		<div class="row">
			<post
				v-for="(post,index) in posts"
				:key="index"
				:username="post.username"
				:postId="post.postId"
				:comments="post.comments != nil ? post.comments : []"
				:likes="post.likes != nil ? post.likes : []"
				:upload_date="post.pubblicationTime"
			/>
		</div>
    <span v-if="posts.length == 0" class="no-posts-text"> There are no posts yet </span>
	<span v-if="posts.length == 0" class="no-posts-text fw-500 fs-6"> Start to follow someone!</span>
	</div>
</template>