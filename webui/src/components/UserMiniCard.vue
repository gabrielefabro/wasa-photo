<script>
export default {
	props: ['user_id','username'],

    methods:{

        clickedUser(){
            this.$emit('clickedUser',this.user_id)
        }
    },
}
</script>
<template>
    <div class="row mb-2 mt-2">
      <div class="col d-flex justify-content-center">
        <div class="user-mini-card card bg-light border-start">
          <div class="card-body">
            <h5 @click="clickedUser" class="user-mini-card-title d-flex justify-content-center">{{ username }} @{{ user_id }}</h5>
          </div>
        </div>
      </div>
    </div>
  </template>
  
  <style>
  .user {
    color: #333;
  }
  
  .user-mini-card {
    color: #333;
    width: 25%;
    border: none;
    border-color: #333;
    transition: all 0.3s;
  }
  
  .user-mini-card:hover {
    transform: translateY(-3px);
    box-shadow: 0 4px 8px rgba(0, 0, 0, 0.1);
  }
  
  .user-mini-card-title {
    cursor: pointer;
    font-weight: bold;
    transition: all 0.3s;
  }
  
  .user-mini-card-title:hover {
    color: #007bff;
  }
  </style>
  