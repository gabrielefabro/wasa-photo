<script>
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
				this.errormsg = null
				
				let response = await this.$axios.get("/users/" + localStorage.getItem('token') + "/home")
				
                if (response.data != null){
					this.posts = response.data
				}
			
			} catch (error) {
				this.errorMsg = this.$utils.errorToString(e);;
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
		<ErrorMsg v-if="errorMsg" :msg="errorMsg" @close-error="errorMsg = ''"></ErrorMsg>
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