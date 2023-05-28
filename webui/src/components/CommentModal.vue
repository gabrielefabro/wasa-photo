<script>
export default {	
	data(){
		return{
			comment:"",
		}
	},
	props:['modal_id','comments_list','owner_user_id','post_id'],

	methods: {
		async addComment(){
			try{
				let response = await this.$axios.post("/users/"+ this.owner_user_id +"/posts/"+this.post_id+"/comments",{
					user_id: localStorage.getItem('token'),
					post_id: this.post_id,
					text: this.comment
				},{
					headers:{
						'Content-Type': 'application/json'
					}
				})

				this.$emit('addComment',{
					comment_id: response.data.comment_id, 
					post_id: this.photo_id, 
					user_id: localStorage.getItem('token'), 
					text: this.comment}
				)
				this.comment = ""
				
			}catch(e){
				console.log(e.toString())
			}
		},

		eliminateCommentToParent(value){
			this.$emit('eliminateComment',value)
		},

		addCommentToParent(newCommentJSON){
			this.$emit('addComment',newCommentJSON)
		},
	},
}
</script>

<template>
    <div class="modal fade my-modal-disp-none" :id="modal_id" tabindex="-1" aria-hidden="true">
        <div class="modal-dialog modal-dialog-centered modal-dialog modal-dialog-scrollable ">
            <div class="modal-content">

                <div class="modal-header">
                    <h1 class="modal-title fs-5" :id="modal_id">Comments</h1>
                    <button type="button" class="btn-close" data-bs-dismiss="modal" aria-label="Close"></button>
                </div>

                <div class="modal-body">
                    <PostComment v-for="(comm,index) in comments_list" 
					:key="index" 
					:author="comm.user_id" 
					:username="comm.username"
					:comment_id="comm.comment_id"
					:post_id="comm.post_id"
					:text="comm.text"
					:photo_owner="owner_user_id"
					

					@eliminateComment="eliminateCommentToParent"
					/>

                </div>
                <div class="modal-footer d-flex justify-content-center w-100">
                    <div class="row w-100 ">
                        <div class="col-10">
                            <div class="mb-3 me-auto">
                                
                                <textarea class="form-control" id="exampleFormControlTextarea1" 
								placeholder="Add a comment..." rows="1" maxLength="50" v-model="comment"></textarea>
                            </div>
                        </div>

                        <div class="col-2 d-flex align-items-center">
                            <button type="button" class="btn btn-primary" 
							@click.prevent="addComment" 
							:disabled="comment.length < 1 || comment.length > 50">
							Send
							</button>
                        </div>
                    </div>
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
