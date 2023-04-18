<script>
export default {
	data: function() {
		return {
			errormsg: null,
			posts: [],
		}
	},
	methods: {
		async loadStream() {
			this.errormsg = null;
			try {
				this.errormsg = null
				// Home get: "/users/:id/home"
				let response = await this.$axios.get("/users/" + localStorage.getItem('token') + "/home")
				if (response.data != null){
					this.posts = response.data
				}
				
			} catch (e) {
				this.errormsg = e.toString()
		}
	},
	mounted() {
		this.loadStream()
	}
}
}
</script>

<template>
	<div class="container-fluid">
		<ErrorMsg v-if="errormsg" :msg="errormsg"></ErrorMsg>

		<div class="row">
			<Photo
				v-for="(post,index) in posts"
				:key="index"
				:owner="post.owner"
				:photo_id="post.photo_id"
				:comments="post.comments != nil ? post.comments : []"
				:likes="post.likes != nil ? post.likes : []"
				:upload_date="post.date"
			/>
		</div>

		<div v-if="photos.length === 0" class="row ">
			<h1 class="d-flex justify-content-center mt-5" style="color: white;">There's no content yet, follow somebody!</h1>
		</div>
	</div>
</template>

<style>
</style>