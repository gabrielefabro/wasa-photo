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
	<ErrorMsg v-if="errormsg" :msg="errormsg"></ErrorMsg>	<div class="login-container">
        <div class="title-login-container">
            <span class="login-title"> Login </span>
        </div>
        <div class="box-container-login">
            <span class="box-text-container-login">Username</span>
            <input :on-submit="doLogin" type="text" name="username-form"  v-model="username" maxlength="15" placeholder="Your username" >
        </div>
        <div class="bottom-login-container">
            <button @click="doLogin"> Login </button>
        </div>

    </div>
</template>