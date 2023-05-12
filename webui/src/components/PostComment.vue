<script>
export default {
    data(){
        return {
            user: "",
        }
    },
	props: ['content','author','postOwner','commentId','postId','username'],

    methods:{
        async deleteComment(){
            try{
                await this.$axios.delete("/users/"+this.postOwner+"/posts/"+this.postId+"/comments/"+this.commentId)

                this.$emit('eliminateComment',this.commentId)

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
                <button v-if="user === author || user === postOwner" class="btn my-btn-comm" @click="deleteComment">
                    <i class="fa-regular fa-trash-can my-trash-icon"></i>
                </button>
            </div>

        </div>

        <div class="row">
            <div class="col-12">
                {{content}}
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