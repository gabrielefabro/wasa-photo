<script>
export default {
	data: function () {
		return {
			errormsg: "",
			posts: [],
            search: "",
		}
	},

    methods: {
        getID(post) {
			return `${post.user.userId}` + `${post.postId}`;
		},
        async getMyStream() {
            try {
				this.errormsg = null
				
				let response = await this.$axios.get("/users/" + localStorage.userId + "/home")
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

    async mounted() {
        await this.getMyStream()
    }
    }

}
</script>

<template>

<ErrorMsg v-if="errormsg" :msg="errormsg"></ErrorMsg>
</template>