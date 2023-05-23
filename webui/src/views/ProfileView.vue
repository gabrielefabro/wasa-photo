<script>
export default {
  data: function () {
    return {
      errormsg: null,
      userExists: false,
      banStatus: false,
      userName: "",
      followStatus: false,
      currentIsBanned: false,
      followerCnt: 0,
      followingCnt: 0,
      postCnt: 0,
      posts: [],
      following: [],
      followers: [],
    };
  },

  watch: {
    currentPath(newid, oldid) {
      if (newid !== oldid) {
        this.loadInfo();
      }
    },
  },

  computed: {
    currentPath() {
      return this.$route.params.user_id;
    },

    sameUser() {
      return this.$route.params.user_id === localStorage.getItem("token");
    },
  },

  methods: {
    async uploadFile() {
      let fileInput = document.getElementById("fileUploader");

      const file = fileInput.files[0];
      const reader = new FileReader();

      reader.readAsArrayBuffer(file);

      reader.onload = async () => {
        // Post photo: /users/:id/photos
        let response = await this.$axios.post(
          "/users/" + this.$route.params.user_id + "/posts",
          reader.result,
          {
            headers: {
              "Content-Type": file.type,
            },
          }
        );
        this.posts.unshift(response.data);
        this.postCnt += 1;
      };
    },

    async followClick() {
      try {
        if (this.followStatus) {
          await this.$axios.delete(
            "/users/" +
              this.$route.params.user_id +
              "/followers/" +
              localStorage.getItem("token")
          );
          this.followerCnt -= 1;
        } else {
          await this.$axios.put(
            "/users/" +
              this.$route.params.user_id +
              "/followers/" +
              localStorage.getItem("token")
          );
          this.followerCnt += 1;
        }
        this.followStatus = !this.followStatus;
      } catch (e) {
        this.errormsg = e.toString();
      }
    },

    async banClick() {
      try {
        if (this.banStatus) {
          await this.$axios.delete(
            "/users/" +
              localStorage.getItem("token") +
              "/banned_users/" +
              this.$route.params.user_id
          );
          this.loadInfo();
        } else {
          await this.$axios.put(
            "/users/" +
              localStorage.getItem("token") +
              "/banned_users/" +
              this.$route.params.user_id
          );
          this.followStatus = false;
        }
        this.banStatus = !this.banStatus;
      } catch (e) {
        this.errormsg = e.toString();
      }
    },

    async loadInfo() {
      if (this.$route.params.user_id === undefined) {
        return;
      }

      try {
        let response = await this.$axios.get(
          "/users/" + this.$route.params.user_id
        );

        this.banStatus = false;
        this.userExists = true;
        this.currentIsBanned = false;

        if (response.status === 206) {
          this.banStatus = true;
          return;
        }

        if (response.status === 204) {
          this.userExists = false;
        }

        this.userName = response.data.userName;
        this.followerCnt = response.data.followers != null
          ? response.data.followers.length
          : 0;
        this.followingCnt = response.data.following != null
          ? response.data.following.length
          : 0;
        this.postCnt = response.data.posts != null
          ? response.data.posts.length
          : 0;
        this.followStatus = response.data.followers != null
          ? response.data.followers.find(
              (obj) => obj.user_id === localStorage.getItem("token")
            )
          : false;
        this.posts = response.data.posts != null ? response.data.posts : [];
        this.followers = response.data.followers != null
          ? response.data.followers
          : [];
        this.following = response.data.following != null
          ? response.data.following
          : [];
      } catch (e) {
        this.currentIsBanned = true;
      }
    },

    goToSettings() {
      this.$router.push(this.$route.params.user_id + "/settings");
    },

    removePhotoFromList(post_id) {
      this.posts = this.posts.filter((item) => item.post_id !== post_id);
    },
  },

  async mounted() {
    await this.loadInfo();
  },
};
</script>
<template>
  <div class="container-fluid profile-page" v-if="!currentIsBanned && userExists">
    <div class="row">
      <div class="col-12 d-flex justify-content-center">
        <div class="card profile-card">
          <div class="row">
            <div class="col">
              <div class="card-body d-flex justify-content-between align-items-center">
                <div class="profile-info">
                  <h5 class="profile-username">{{ userName }}</h5>
                  <h6 class="profile-userid">@{{ this.$route.params.user_id }}</h6>
                </div>
                <div>
                  <button
                    v-if="!sameUser && !banStatus"
                    @click="followClick"
                    class="btn btn-follow"
                  >
                    {{ followStatus ? "Following" : "Follow" }}
                  </button>
                  <button
                    v-if="!sameUser"
                    @click="banClick"
                    class="btn btn-ban"
                  >
                    {{ banStatus ? "Unban" : "Ban" }}
                  </button>
                  <button
                    v-else
                    class="btn btn-settings"
                    @click="goToSettings"
                  >
                    <i class="fa-solid fa-gear"></i>
                  </button>
                </div>
              </div>
            </div>
          </div>

          <div v-if="!banStatus" class="row profile-stats">
            <div class="col-4">
              <h6 class="stat-label">Posts</h6>
              <h6 class="stat-count">{{ postCnt }}</h6>
            </div>
            <div class="col-4">
              <h6 class="stat-label">Followers</h6>
              <h6 class="stat-count">{{ followerCnt }}</h6>
            </div>
            <div class="col-4">
              <h6 class="stat-label">Following</h6>
              <h6 class="stat-count">{{ followingCnt }}</h6>
            </div>
          </div>
        </div>
      </div>
    </div>

    <div class="row">
      <div class="container-fluid mt-3">
        <div class="row">
          <div class="col-12 d-flex justify-content-center">
            <h2>Posts</h2>
            <input
              id="fileUploader"
              type="file"
              class="profile-file-upload"
              @change="uploadFile"
              accept=".jpg, .png"
            />
            <label
              v-if="sameUser"
              class="btn btn-add-photo"
              for="fileUploader"
            >
              Add Photo
            </label>
          </div>
        </div>

        <div class="row">
          <div class="col-3"></div>
          <div class="col-6">
            <hr class="border border-dark" />
          </div>
          <div class="col-3"></div>
        </div>
      </div>
    </div>

    <div class="row">
      <div class="col">
        <div v-if="!banStatus && postCnt > 0">
          <Post
            v-for="(post,index) in posts"
            :key="index"
            :user_id="this.$route.params.user_id"
            :post_id="post.post_id"
            :comments="post.comments"
            :likes="post.likes"
            :publication_time="post.publication_time"
            @removePhoto="removePhotoFromList"
          />
        </div>

        <div v-else class="mt-5 no-posts-message">
          <h2>No posts yet</h2>
        </div>
      </div>
    </div>

    <ErrorMsg v-if="errormsg" :msg="errormsg" />
  </div>
</template>

<style>
.profile-page {
  background-color: #fafafa;
  padding-top: 20px;
}

.profile-card {
  width: 50%;
}

.profile-username {
  font-size: 24px;
  font-weight: bold;
}

.profile-userid {
  font-size: 16px;
  color: gray;
  margin-top: 5px;
}

.btn-follow {
  background-color: #3897f0;
  color: white;
  border: none;
  padding: 8px 16px;
  border-radius: 4px;
  font-size: 14px;
  font-weight: bold;
}

.btn-ban {
  background-color: #ed4956;
  color: white;
  border: none;
  padding: 8px 16px;
  border-radius: 4px;
  font-size: 14px;
  font-weight: bold;
}

.btn-settings {
  background-color: transparent;
  border: none;
  padding: 8px;
  cursor: pointer;
}

.profile-stats {
  margin-top: 10px;
  margin-bottom: 20px;
}

.stat-label {
  font-size: 14px;
  font-weight: bold;
  color: gray;
}

.stat-count {
  font-size: 18px;
  font-weight: bold;
}

.btn-add-photo {
  background-color: #3897f0;
  color: white;
  border: none;
  padding: 8px 16px;
  border-radius: 4px;
  font-size: 14px;
  font-weight: bold;
  cursor: pointer;
}

.no-posts-message {
  display: flex;
  justify-content: center;
  align-items: center;
  height: 200px;
  background-color: white;
  border-radius: 4px;
  color: gray;
  font-size: 20px;
  font-weight: bold;
}

.fa-gear {
  color: gray;
}

.fa-gear:hover {
  transform: scale(1.3);
}
</style>

