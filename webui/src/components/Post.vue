<script>
export default {
	data(){
		return{
			photoURL: "",
			liked: false,
			allComments: [],
			allLikes: [],
		}
	},

	props: ['user_id','like','comments',"publication_time","post_id","isOwner"], 

	methods:{
		loadPhoto(){
			this.photoURL = __API_URL__+ "/users/"+this.user_id+"/posts/"+this.post_id 
		},

		async deletePhoto(){
			try{
				await this.$axios.delete("/users/"+this.user_id+"/posts/"+this.post_id)
				this.$emit("removePhoto",this.post_id)
			}catch(e){
			}
		},

		photoOwnerClick: function(){
			this.$router.replace("/users/"+this.user_id)
		},

		async toggleLike() {

			if(this.isOwner){ 
				return
			}

			const bearer = localStorage.getItem('token')

			try{
				if (!this.liked){

					await this.$axios.put("/users/"+ this.user_id +"/posts/"+this.post_id+"/likes/"+ bearer)
					this.allLikes.push({
						user_id: bearer,
						username: bearer
					})

				}else{
					await this.$axios.delete("/users/"+ this.user_id  +"/posts/"+this.post_id+"/likes/"+ bearer)
					this.allLikes.pop()
				}

				this.liked = !this.liked;
			}catch(e){
				//
			}
      		
    	},

		removeCommentFromList(value){
			this.allComments = this.allComments.filter(item=> item.comment_id !== value)
		},

		addCommentToList(comment){
			this.allComments.push(comment)
		},
	},
	
	mounted(){
		this.loadPhoto()

		if (this.like != null){
			this.allLikes = this.like
		}

		if (this.like != null){
			this.liked = this.allLikes.some(obj => obj.user_id === localStorage.getItem('token'))
		}
		if (this.comments != null){
			this.allComments = this.comments
		}		
	},
	aftermounted() {
		this.photoURL = ""
	}

}
</script>

<template>
	<div class="container-fluid mt-3 mb-5">
	  <LikeModal :modal_id="'like_modal' + post_id" :like="allLikes" />
  
	  <CommentModal
		:modal_id="'comment_modal' + post_id"
		:comments_list="allComments"
		:owner_user_id="user_id"
		:post_id="post_id"
		@eliminateComment="removeCommentFromList"
		@addComment="addCommentToList"
	  />
  
	  <div class="d-flex flex-row justify-content-center">
		<div class="card my-card">
		  <div class="d-flex justify-content-end">
			<button v-if="isOwner" class="my-trnsp-btn my-dlt-btn me-2" @click="deletePhoto">
			  <i class="fa-solid fa-trash w-100 h-100"></i>
			</button>
		  </div>
		  <div class="d-flex justify-content-center photo-background-color">
			<img :src="photoURL" class="card-img-top img-fluid">
		  </div>
  
		  <div class="card-body">
			<div class="container">
			  <div class="d-flex flex-row justify-content-end align-items-center mb-2">
				<button class="my-trnsp-btn m-0 p-1 me-auto" @click="photoOwnerClick">
				  <i>From {{ user_id }}</i>
				</button>
  
				<button class="my-like-btn m-0 p-1 d-flex justify-content-center align-items-center">
				  <i
					@click="toggleLike"
					:class="'me-1 my-heart-color w-100 h-100 fa ' + (liked ? 'fa-heart' : 'fa-heart-o')"
				  ></i>
				  <i data-bs-toggle="modal" :data-bs-target="'#like_modal' + post_id" class="my-like-count">
					{{ allLikes.length }}
				  </i>
				</button>
  
				<button class="my-comment-btn m-0 p-1  d-flex justify-content-center align-items-center" data-bs-toggle="modal" :data-bs-target="'#comment_modal' + post_id">
				  <i class="my-comment-icon fa-regular fa-comment me-1" @click="commentClick"></i>
				  <i class="my-comment-count">{{ allComments != null ? allComments.length : 0 }}</i>
				</button>
			  </div>
  
			  <div class="d-flex flex-row justify-content-start align-items-center">
				<p>Uploaded on {{ publication_time }}</p>
			  </div>
			</div>
		  </div>
		</div>
	  </div>
	</div>
  </template>
  
  <style>
  .photo-background-color {
	background-color: grey;
  }
  
  .my-card {
	width: 27rem;
	border-color: black;
	border-width: thin;
  }
  
  .my-heart-color {
	color: #ff7171;
  }
  
  .my-heart-color:hover {
	color: #f44336;
  }
  
  .my-comment-btn {
	background-color: #ffffff;
	border: none;
	color: #808080;
  }
  
  .my-comment-btn:hover {
	color: #000000;
  }
  
  .my-comment-icon {
	color: #808080;
  }
  
  .my-comment-count {
	color: #808080;
  }
  
  .my-like-btn {
	background-color: #ffffff;
	border: none;
	color: #808080;
  }
  
  .my-like-btn:hover {
  color: #ff7171;
}

.my-like-count {
  color: #808080;
}

.my-dlt-btn {
  font-size: 19px;
  color: #808080;
}

.my-dlt-btn:hover {
  font-size: 19px;
  color: #ff7171;
}
</style>

