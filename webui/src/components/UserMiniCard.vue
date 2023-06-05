
<script>
export default {
  props: ['user_id', 'username'],

  methods: {
    clickedUser() {
      this.$emit('clickedUser', this.user_id);
    },
  },
};
</script>

<template>
  <div class="row mb-2 mt-2">
    <div class="col d-flex justify-content-center">
      <div class="user-mini-card card bg-light border-start">
        <div class="card-body text-center"> 
          <div class="d-flex align-items-center">
            <img
              class="profile-pic"
              src="../assets/photo/vettore-icona-profilo-avatar-predefinito-foto-utente-social-media-183042379.jpg" 
              alt="Profile Picture"
            />
            <h5 @click="clickedUser" class="user-mini-card-title ml-2">
              <span class="username">{{ username }}</span> @{{ user_id }}
            </h5>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<style>
.user {
  color: var(--color-text);
}

.user-mini-card {
  color: black;
  width: 40%; 
  height: auto; 
  border: none;
  border-color: black;
}

.user-mini-card-title:hover {
  cursor: pointer;
  font-weight: bold;
}

.profile-pic {
  width: 40px;
  height: 40px;
  border-radius: 50%;
  object-fit: cover;
}
</style>
