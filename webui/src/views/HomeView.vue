<script>
export default {
  data: function () {
    return {
      errorMsg: "",
      posts: [],
    };
  },

  methods: {
    async getMyStream() {
      try {
        this.errorMsg = null;

        let response = await this.$axios.get(
          "/users/" + localStorage.getItem("token") + "/home"
        );

        if (response.data != null) {
          this.posts = response.data;
        }
      } catch (error) {
        this.errorMsg = this.$utils.errorToString(e);
      }
    },

    async mounted() {
      await this.getMyStream();
    },

    goToSettingsPage() {
      this.$router.push("/users/" + localStorage.getItem("token") + "/settings");
    },

    logout() {
      // Implement your logout logic here
    },
  },
};
</script>

<template>
  <div class="container-fluid">
    <ErrorMsg v-if="errorMsg" :msg="errorMsg" @close-error="errorMsg = ''"></ErrorMsg>
    <div class="row">
      <post
        v-for="(post, index) in posts"
        :key="index"
        :username="post.username"
        :postId="post.postId"
        :comments="post.comments != null ? post.comments : []"
        :likes="post.likes != null ? post.likes : []"
        :publication_time="post.publication_time"
      />
    </div>
    <span v-if="posts.length === 0" class="no-posts-text">There are no posts yet,</span>
    <span v-if="posts.length === 0" class="no-posts-text fw-500 fs-6">start to follow someone!</span>

    <!-- Pulsante per andare al profilo -->
    <button class="profile-button" @click="goToSettingsPage">
      <i class="fa fa-user"></i>
    </button>

    <!-- Pulsante di logout -->
    <button class="logout-button" @click="logout">
      <i class="fa fa-sign-out"></i>
    </button>
  </div>
</template>

<style>
.settings-button {
  position: fixed;
  bottom: 20px;
  right: 20px;
  background-color: #fff;
  border: none;
  border-radius: 50%;
  width: 40px;
  height: 40px;
  box-shadow: 0 2px 4px rgba(0, 0, 0, 0.1);
  cursor: pointer;
}

.settings-button i {
  font-size: 20px;
}

.profile-button,
.logout-button {
  position: fixed;
  bottom: 20px;
  right: 80px;
  background-color: #fff;
  border: none;
  border-radius: 50%;
  width: 40px;
  height: 40px;
  box-shadow: 0 2px 4px rgba(0, 0, 0, 0.1);
  cursor: pointer;
}

.profile-button i,
.logout-button i {
  font-size: 20px;
}

.profile-button {
  background-color: #3498db;
  color: #fff;
}

.logout-button {
  background-color: #e74c3c;
  color: #fff;
}

.no-posts-text {
  color: #555;
  display: block;
  text-align: center;
  margin-top: 20px;
}
</style>
