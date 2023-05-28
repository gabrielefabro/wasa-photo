<script>
export default {
	data: function() {
		return {
			users: [],
			errormsg: null,
		}
	},

	props:['searchValue'],

	watch:{
		searchValue: function(){
			this.loadSearchedUsers() 
		},
	},

	methods:{
		async loadSearchedUsers(){
			this.errormsg = null;
			if (
				this.searchValue === undefined ||
				this.searchValue === "" || 
				this.searchValue.includes("?") ||
				this.searchValue.includes("_")){
				this.users = []
				return 
			}
			try {
				let response = await this.$axios.get("/users",{
						params: {
						user_id: this.searchValue,
						username: this.username
					},
				});
				this.users = response.data

			} catch (e) {
				this.errormsg = e.toString();
			}
		},

		goToProfile(user_id){
			this.$router.replace("/users/"+user_id)
		}
	},

	async mounted(){
		if (!localStorage.getItem('token')){
			this.$router.replace("/login")
		}
		await this.loadSearchedUsers()
		
	},
}
</script>

<template>
	<div class="container-fluid h-100 ">
		<UserMiniCard v-for="(user,index) in users" 
		:key="index"
		:user_id="user.user_id" 
		:username="user.nickname" 
		@clickedUser="goToProfile"/>

		<p v-if="users.length == 0" class="no-result-text d-flex justify-content-center"> No users found.</p>

		<ErrorMsg v-if="errormsg" :msg="errormsg"></ErrorMsg>
	</div>
</template>

<style>

.no-result-text{
	color: white;
	font-style: italic;
}
</style>