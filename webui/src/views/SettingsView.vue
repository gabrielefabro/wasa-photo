<script>
export default {
	data: function () {
		return {
			errormsg: null,
			username: "",
		}
	},

	methods:{
		async modifyNickname(){
			try{
				let resp = await this.$axios.put("/users/"+this.$route.params.user_id,{
					username: this.username,
				})

				this.username=""
			}catch (e){
				this.errormsg = e.toString();
			}
		},
	},

}
</script>

<template>
	<div class="container">
	  <div class="row">
		<div class="col d-flex justify-content-center mb-4">
		  <h1>{{ this.$route.params.user_id }}'s Settings</h1>
		</div>
	  </div>
  
	  <div class="row">
		<div class="col-12 d-flex justify-content-center">
		  <p class="text-danger me-1">[Disclaimer]</p>
		  <p>A user has the following structure:</p>
		  <p class="text-success ms-1 me-1">username</p>
		  <p>@identifier.</p>
		</div>
		<div class="col-12 d-flex justify-content-center">
		  <p>You can only modify the part before the @ symbol (the <span class="text-success">Username</span>) and not the identifier of the user.</p>
		</div>
		<div class="col-12 d-flex justify-content-center">
		  <p>The term "Username" refers to a nickname.</p>
		</div>
	  </div>
  
	  <div class="row mt-3">
		<div class="col d-flex justify-content-center">
		  <div class="input-group mb-3 w-50">
			<input
			  type="text"
			  class="form-control"
			  placeholder="Enter your new username..."
			  maxlength="15"
			  minlength="1"
			  v-model="username"
			/>
			<div class="input-group-append">
			  <button
				class="btn btn-primary"
				@click="modifyNickname"
				:disabled="!username || username.length > 15 || username.length < 1 || username.trim().length === 0"
			  >
				Modify
			  </button>
			</div>
		  </div>
		</div>
	  </div>
  
	  <div class="row">
		<div v-if="username.trim().length > 0" class="col d-flex justify-content-center mt-2">
		  Preview: {{ username }}@{{ this.$route.params.user_id }}
		</div>
	  </div>
  
	  <div class="row mt-4">
		<div class="col d-flex justify-content-center">
		  <ErrorMsg v-if="errormsg" :msg="errormsg"></ErrorMsg>
		</div>
	  </div>
	</div>
  </template>
  

<style>
</style>