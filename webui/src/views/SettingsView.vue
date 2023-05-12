<script>
export default {
	data: function () {
		return {
			errormsg: null,
			nickname: "",
		}
	},

	methods:{
		async modifyusername(){
			try{
				let resp = await this.$axios.put("/users/"+this.$route.params.id,{
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
	<div class="container-fluid">
		<div class="row">
			<div class="col d-flex justify-content-center mb-2">
				<h1>{{ this.$route.params.id }}'s Settings</h1>
			</div>
		</div>


		<div class="row mt-2">
			<div class="col d-flex justify-content-center">
				<div class="input-group mb-3 w-25">
					<input
						type="text"
						class="form-control w-25"
						placeholder="Your new username..."
						maxlength="15"
						minlength="1"
						v-model="username"
					/>
					<div class="input-group-append">
						<button class="btn btn-outline-secondary" 
						@click="modifyusername"
						:disabled="username === null || nickname.length >= 15 || nickname.length < 1 || nickname.trim().length===0">
						Modify</button>
					</div>
				</div>
			</div>
		</div>

		<div class="row" >
			<div v-if="username.trim().length>0" class="col d-flex justify-content-center">
				Preview: {{nickname}} @{{ this.$route.params.id }}
			</div>
		</div>

		<div class="row">
			<ErrorMsg v-if="errormsg" :msg="errormsg"></ErrorMsg>
		</div>
	</div>
	
</template>
