<script>
export default {
	data: function() {
		return {
			errorMsg: null,
			identifier: "",
			disabled: true,
		}
	},
	methods: {
		async login() {
			this.errormsg = null;
			try {
				let response = await this.$axios.post("/session",{
					user_id: this.identifier.trim()
				});

				localStorage.setItem('token',response.data.userId);
				this.$router.replace("/home")
				this.$emit('updatedLoggedChild',true)
				
			} catch (e) {
				this.errorMsg = this.$utils.errorToString(e);;
			}
		},
	},
	mounted(){
		if (localStorage.getItem('token')){
			this.$router.replace("/home")
		}
	},
	
}
</script>

<template>
	<div class="container-fluid h-100 m-0 p-0 login">

		<div class="row ">
			<div class="col">
				<ErrorMsg v-if="errorMsg" :msg="errorMsg" @close-error="errorMsg = ''"></ErrorMsg>
			</div>
		</div>

		<div class="row h-100 w-100 m-0">
			
			<form @submit.prevent="login" class="d-flex flex-column align-items-center justify-content-center p-0">

				<div class="row mt-2 mb-3 border-bottom">
					<div class="col">
						<h2 class="login-title">WASAPhoto Login</h2>
					</div>
				</div>

				<div class="row mt-2 mb-3">
					<div class="col">
						<input 
						type="text" 
						class="form-control" 
						v-model="identifier" 
						maxlength="16"
						minlength="1"
						placeholder="Your identifier" />
					</div>
				</div>

				<div class="row mt-2 mb-5 ">
					<div class="col ">
						<button class="btn btn-dark" :disabled="identifier == null || identifier.length >16 || identifier.length <1 || identifier.trim().length<1"> 
						Register/Login 
						</button>
					</div>
				</div>
			</form>
		</div>
	</div>
</template>

<style>
.login {
    height: 100vh;
}

.login-title {
    color: black;
}
</style>