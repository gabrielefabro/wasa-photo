<script>
export default {
    data(){
        return {
            user: "",
        }
    },
	props: ['text','author','photo_owner','comment_id','post_id','username'],

    methods:{
        async deleteComment(){
            try{
                await this.$axios.delete("/users/"+this.photo_owner+"/posts/"+this.post_id+"/comments/"+this.comment_id)

                this.$emit('eliminateComment',this.comment_id)

            }catch (e){
                console.log(e.toString())
            }
        },
    },

    mounted(){
        this.user = localStorage.getItem('token')
    }

}
</script>

<template>
	<div class="container-fluid">

        <hr>
        <div class="row">
            <div class="col-10">
                <h5>{{username}} @{{author}}</h5>
            </div>

            <div class="col-2">
                <button v-if="user === author || user === photo_owner" class="btn my-btn-comm" @click="deleteComment">
                    <i class="fa-regular fa-trash-can my-trash-icon"></i>
                </button>
            </div>

        </div>

        <div class="row">
            <div class="col-12">
                {{text}}
            </div>

        </div>
        <hr>
    </div>
</template>

<style>
.my-btn-comm{
    border: none;
}
.my-btn-comm:hover{
    border: none;
    color: var(--color-red-danger);
    transform: scale(1.1);
}

</style>
