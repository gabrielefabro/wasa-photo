<script>
export default {	

	props:['modal_id','like'],

}
</script>

<template>
	<div class="modal fade modal-dialog modal-dialog-centered modal-dialog-scrollable my-modal-disp-none" :id="modal_id" tabindex="-1" aria-hidden="true">
	<div class="modal-dialog">
		<div class="modal-content">
		<div class="modal-header">
			<h1 class="modal-title fs-5" :id="modal_id">Likes</h1>
			<button type="button" class="btn-close" data-bs-dismiss="modal" aria-label="Close"></button>
		</div>
		<div class="modal-body">
			<UserMiniCard v-for="(user,index) in like" 
			:key="index" 
			:user_id="user.user_id"
			:username="user.username"
			/>
		</div>
		</div>
	</div>
	</div>
</template>

<style>
.my-modal-disp-none{
	display: none;
}

</style>