<script>
export default{
	data() {
		return {
			errorMsg: "",
			username: "",

			testUsername: new RegExp('^[a-z0-9]*$')
		}
	},

	methods: {
		async doLogin() {
			this.errorMsg = null;
			try {
				if(!this.testUsername.test(this.username)) throw "Invalid user, it should contain only lowcase letters or numbers"
				if (this.username.length < 1 || this.username.length > 15) throw "Invalid username, it must contains mininum 1 characters and maximum 15 characters"
				let response = await this.$axios.post('/session', {
					username: this.username,
				});
				this.$router.replace("/home")
				localStorage.userId = response.data.user.userId;
                localStorage.username = response.data.user.username;
			} catch (error) {
				this.errorMsg = error.toString();
				
			}
		},
		mounted() {
        localStorage.clear();
    },
	}
	
}

</script>

<template>
	<div class="container-fluid h-100 m-0 p-0 login">

		<div class="row ">
			<div class="col">
				<ErrorMsg v-if="errorMsg" :msg="errorMsg"></ErrorMsg>
			</div>
		</div>

		<div class="row h-100 w-100 m-0">
			
			<form @submit.prevent="doLogin" class="d-flex flex-column align-items-center justify-content-center p-0">

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
						maxlength="15"
						minlength="1"
						placeholder="Your username" />
					</div>
				</div>

				<div class="row mt-2 mb-5 ">
					<div class="col ">
						<button class="btn btn-dark" :disabled="identifier == null || identifier.length >16 || identifier.length <3 || identifier.trim().length<3"> 
						Register/Login 
						</button>
					</div>
				</div>
			</form>
		</div>
	</div>
</template>

<style>

.login-title {
    color: black;
}
</style>