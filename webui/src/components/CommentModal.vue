<script>
export default {	
	data(){
		return{
			commentValue:"",
		}
	},
	props:['modalId','commentsList','postOwner','postId'],

	methods: {
		async addComment(){
			try{
				let response = await this.$axios.post("/users/"+ this.postOwner +"/posts/"+this.postId+"/comments",{
					userId: localStorage.getItem('token'),
					comment: this.commentValue
				},{
					headers:{
						'Content-Type': 'application/json'
					}
				})

				this.$emit('addComment',{
					commentId: response.data.commentId, 
					postId: this.postId, 
					userId: localStorage.getItem('token'), 
					comment: this.commentValue}
				)
				this.commentValue = ""
				
			}catch(e){
				console.log(e.toString())
			}
		},
	},
}
</script>

<template>
    <div class="modal fade my-modal-disp-none" :id="modalId" tabindex="-1" aria-hidden="true">
        <div class="modal-dialog modal-dialog-centered modal-dialog modal-dialog-scrollable ">
            <div class="modal-content">

                <div class="modal-header">
                    <h1 class="modal-title fs-5" :id="modalId">Comments</h1>
                    <button type="button" class="btn-close" data-bs-dismiss="modal" aria-label="Close"></button>
                </div>

                <div class="modal-body">
                    <PostComment v-for="(comm,index) in commentsList" 
					:key="index" 
					:author="comm.userId" 
					:username="comm.username"
					:commentId="comm.commentId"
					:postId="comm.postId"
					:content="comm.comment"
					:postOwner="postOwner"
					/>

                </div>
                <div class="modal-footer d-flex justify-content-center w-100">
                    <div class="row w-100 ">
                        <div class="col-10">
                            <div class="mb-3 me-auto">
                                
                                <textarea class="form-control" id="exampleFormControlTextarea1" 
								placeholder="Add a comment..." rows="1" maxLength="30" v-model="commentValue"></textarea>
                            </div>
                        </div>

                        <div class="col-2 d-flex align-items-center">
                            <button type="button" class="btn btn-primary" 
							@click.prevent="addComment" 
							:disabled="commentValue.length < 1 || commentValue.length > 50">
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