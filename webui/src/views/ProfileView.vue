<script>
export default {
  data: function () {
    return {
      errormsg: null,
      userExists: false,
      banStatus: false,
      username: "",
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

  watch:{
        currentPath(newid,oldid){
            if (newid !== oldid){
                this.loadInfo()
            }
        },
    },

	computed:{

        currentPath(){
            return this.$route.params.user_id
        },
        

		sameUser(){
			return this.$route.params.user_id === localStorage.getItem('token')
		},
	},
  methods: {

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

        this.username = response.data.username;
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
      this.$router.push(this.$route.params.user_id+'/settings')
    },

    removePhotoFromList(post_id) {
      this.posts = this.posts.filter((item) => item.post_id !== post_id);
    },
  },

  mounted() {
    this.loadInfo();
  },
};
</script>
<template>
  <div class="container-fluid" v-if="!currentIsBanned && userExists">
    <div class="row">
      <div class="col-12 d-flex justify-content-center">
        <div class="card w-50 container-fluid">
          <div class="row">
            <div class="col">
              <div class="card-body d-flex justify-content-between align-items-center">
                <h5 class="card-title p-0 me-auto mt-auto">{{ username }} @{{ this.$route.params.user_id }}</h5>

                <button v-if="!sameUser && !banStatus" @click="followClick" class="btn btn-success ms-2">
                  {{ followStatus ? "Unfollow" : "Follow" }}
                </button>

                <button v-if="!sameUser" @click="banClick" class="btn btn-danger ms-2">
                  {{ banStatus ? "Unban" : "Ban" }}
                </button>

                <button v-else class="my-trnsp-btn ms-2" @click="goToSettings">
                  <i class="my-nav-icon-gear fa-solid fa-gear"></i>
                </button>
              </div>
            </div>
          </div>

          <div v-if="!banStatus" class="row mt-1 mb-1">
            <div class="col-4 d-flex justify-content-start">
              <h6 class="ms-3 p-0">Posts: {{ postCnt }}</h6>
            </div>

            <div class="col-4 d-flex justify-content-center">
              <h6 class="p-0">Followers: {{ followerCnt }}</h6>
            </div>

            <div class="col-4 d-flex justify-content-end">
              <h6 class="p-0 me-3">Following: {{ followingCnt }}</h6>
            </div>
          </div>
        </div>
      </div>
    </div>

    <div class="upload-form-background" @click.self="this.$emit('exit-upload-form')">
        <div class="upload-form-container" v-if="!file64">
            <div class="drag-drop-area-container">
                <button class="drag-drop-area" @click="this.$refs.file.click()">
                    <input type="file" ref="file" accept=".jpg,.jpeg" @change="onChange" hidden />
                    <span class="drag-drop-area-text">
                        Drop your photo here
                    </span>
                    <span class="drag-drop-area-subtext">
                        max size 5MB, only jpg, jpeg
                    </span>

                </button>
            </div>
            <div class="bottom-area">
                <button @click="this.$refs.file.click()" class="upload-button">Choose File
                    <input type="file" ref="file" accept=".jpg,.jpeg" @change="onChange" hidden />
                </button>

            </div>
        </div>

        <EditorPost :image64="file64" :editorType="this.$props.photoType" v-if="file64"
            @exit-upload-form="this.$emit('exit-upload-form')" @save-upload-form="saveData" />
    </div>

    <div class="row">
      <div class="col">
        <div v-if="!banStatus && postCnt > 0">
          <Post v-for="(post,index) in posts" 
                :key="index" 
                :user_id="this.$route.params.user_id" 
                :post_id="post.post_id" 
                :comments="post.comments" 
                :likes="post.likes" 
                :publication_time="post.publication_time" 
                :isOwner="sameUser" 

                @removePhoto="removePhotoFromList"
          />
        </div>

        <div v-else class="mt-5">
          <h2 class="d-flex justify-content-center" style="color: white;">No posts yet</h2>
        </div>
      </div>
    </div>

    <ErrorMsg v-if="errormsg" :msg="errormsg"></ErrorMsg>
  </div>
</template>

<style>
.profile-file-upload {
  display: none;
}

.my-nav-icon-gear {
  color: #6c757d;
  transition: transform 0.3s;
}

.my-nav-icon-gear:hover {
  transform: scale(1.3);
}

.my-btn-add-photo {
  background-color: #28a745;
  border-color: #6c757d;
  transition: color 0.3s, background-color 0.3s, border-color 0.3s;
}

.my-btn-add-photo:hover {
  color: #fff;
  background-color: #218838;
  border-color: #6c757d;
}
</style>


