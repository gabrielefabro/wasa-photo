<script>
export default {
    components: {
        UploadPhoto,
    },
    data() {
        return {
            errorMsg: "",
            // Profile data
            userID: parseInt(this.$route.params.userID),
            username: "",
            followersCount: 0,
            followingsCount: 0,
            postsCount: 0,
            isFollowed: false,

            isOwner: false,

            followTextButton: "Follow",

            // Other Data
            textCounter: 0,
            profilesArray: [],
            textHeader: "",
            typeList: "",

            // Posts data
            posts: [],
            showPost: false,
            postViewData: {},

            // Load more data
            busy: false,
            dataAvaible: true,

            // Follower data
            dataGetter: () => { },
            showList: false,

            // Options data
            showOptions: false,

            isLoading: false,
        }
    },
    methods: {
        async getProfile() {
            this.isLoading = true;
            try {
                let response = await this.$axios.get(`users/${this.userID}`)
                this.userID = response.data.user.userID;
                this.username = response.data.user.username;
                this.followersCount = response.data.followersCount;
                this.followingsCount = response.data.followingsCount;
                this.postsCount = response.data.postsCount;
                this.isFollowed = response.data.isFollowed;
                this.followTextButton = this.isFollowed ? "Unfollow" : "Follow";
                this.isOwner = localStorage.userID == this.userID;
            } catch (e) {
                this.errormsg = e.toString();
            }
            this.isLoading = false;
        },
        async getPosts() {
            this.isLoading = true;
            try {
                let response = await this.$axios.get(`/users/${this.userID}/posts`);
                if (response.data == null) {
                    this.dataAvaible = false;
                    this.isLoading = false;
                    return;
                }
                this.posts.push(...response.data);
            } catch (e) {
                this.errormsg = e.toString();            };
            this.isLoading = false;
        },
        editingUsername() {
            if (this.isOwner) {
                document.querySelectorAll(".top-body-profile-username")[0].style.outline = "auto";
                document.querySelectorAll(".top-body-profile-username")[0].style.outlineColor = "#03C988";
            }
        },
        async saveChangeUsername() {
            if (this.isOwner) {
                document.querySelectorAll(".top-body-profile-username")[0].style.outline = "none";
                if (this.username == "" | this.username.length < 1) {
                    this.username = localStorage.username;
                    return
                }
                this.isLoading = true;
                try {
                    let _ = await this.$axios.put(`/users/${this.userID}/username`, { username: this.username });
                    localStorage.username = this.username;
                } catch (e) {
                    this.errormsg = e.toString();
                    this.username = localStorage.username;
                }
                this.isLoading = false;
            }
        },
        getFollowers() {
            this.showList = true;
            this.textHeader = "Followers";
            this.typeList = "simple";
            this.dataGetter = async (profilesArray, dataAvaible) => {
                try {
                    let response = await this.$axios.get(`/users/${this.userID}/followers`);
                    if (response.data == null) {
                        dataAvaible = false;
                        return;
                    }
                    profilesArray.push(...response.data);
                } catch (e) {
                    this.errormsg = e.toString();
                }
            }
        },
        getFollowings() {
            this.showList = true;
            this.textHeader = "Followings";
            this.typeList = "simple";
            this.dataGetter = async (profilesArray, dataAvaible) => {
                try {
                    let response = await this.$axios.get(`/users/${this.userID}/followings`);
                    if (response.data == null) {
                        dataAvaible = false;
                        return;
                    }
                    profilesArray.push(...response.data);
                } catch (e) {
                    this.errormsg = e.toString();
                }
            }
        },
        freeLists() {
            this.showList = false;
            this.profilesArray = [];
            this.textHeader = "";
        },
        async follow() {
            if (this.isFollowed) {
                try {
                    let _ = await this.$axios.delete(`users/${localStorage.userID}/followings/${this.userID}`);
                    this.isFollowed = false;
                    this.followTextButton = "Follow";
                    this.followersCount--;
                } catch (e) {
                    this.errormsg = e.toString();                }
            } else {
                try {
                    let _ = await this.$axios.put(`users/${localStorage.userID}/followings/${this.userID}`, {});
                    this.isFollowed = true;
                    this.followTextButton = "Unfollow";
                    this.followersCount++;
                } catch (e) {
                    this.errormsg = e.toString();                }
            }
        },
        openPost(post) {
            this.showPost = true;
            this.postViewData = post;
        },
        exitPost() {
            this.showPost = false;
            this.postViewData = {};
        },
        updateLike(data) {
            this.posts.forEach(post => {
                if (post.postID == data.postID) {
                    post.liked = data.liked;
                    post.likesCount = data.liked ? post.likesCount + 1 : post.likesCount - 1;
                }
            });
        },
        doLogout() {
            localStorage.clear();
            this.$router.push('/login');
        },
        getBans() {
            this.showList = true;
            this.textHeader = "Bans";
            this.typeList = "ban";
            this.dataGetter = async (profilesArray, dataAvaible) => {
                try {
                    let response = await this.$axios.get(`/users/${this.userID}/bans`);
                    if (response.data == null) {
                        dataAvaible = false;
                        return;
                    }
                    profilesArray.push(...response.data);
                } catch (e) {
                    this.errorMsg = this.$utils.errorToString(e);;
                }
            }
        },
        closeOptions() {
            setTimeout(() => {
                this.showOptions = false;
            }, 500);
        },
        async banUser() {
            try {
                let _ = await this.$axios.put(`/users/${localStorage.userID}/bans/${this.userID}`, {});
                this.$router.push(`/users/${localStorage.userID}`);
            } catch (e) {
                this.errormsg = e.toString();            }
            this.showOptions = false;
        },
        async deletePost(postID) {
            this.isLoading = true;
            try {
                let _ = await this.$axios.delete(`users/${localStorage.userID}/posts/${postID}`);
                this.posts = this.posts.filter(post => post.postID != postID);
                this.postsCount--;
                this.exitPost();
            } catch (e) {
                this.errormsg = e.toString();            }
            this.isLoading = false;
        },
        async deleteProfile() {
            this.isLoading = true;
            try {
                let _ = await this.$axios.delete(`users/${localStorage.userID}`);
                this.doLogout();
            } catch (e) {
                this.errorMsg = this.$utils.errorToString(e);;
            }
            this.isLoading = false;
        },
    },
    beforeMount() {
        if (!localStorage.token) {
            this.$router.push('/login');
        }
        if (localStorage.userID === this.$route.params.userID) {
            this.isOwner = true;
        }
    },

    mounted() {
        this.getProfile();
        this.getPosts();

        if (this.isOwner) {
            document.querySelectorAll(".top-body-profile-bio-text")[0].style.cursor = "text";
            document.querySelectorAll(".top-body-profile-username")[0].style.cursor = "text";

            document.querySelectorAll(".top-profile-picture")[0].style.cursor = "pointer";
        }

        document.addEventListener('scroll');
    },

    beforeRouteUpdate(to, from) {
        this.posts = [];
        this.dataAvaible = true;

        this.userID = parseInt(to.params.userID);
        this.getProfile();
        this.getPosts();
    },
}

</script>


<template>
	<LoadingSpinner :loading=isLoading />
    <ErrorMsg v-if="errorMsg" :msg="errorMsg" @close-error="errorMsg = ''"></ErrorMsg>
    <UploadPhoto v-if="isEditingPropic" :photoType="'proPic'" @exit-upload-form="isEditingPropic = false"
        @refresh-data="updateProfile" @error-occurred="(value) => { errorMsg = value }" />
    <div class="top-profile-container">
        <div class="top-body-profile-container">
            <div class="profile-options-button-container">
                <button class="profile-options-button" @click="showOptions = true" @focusout="closeOptions">
                    <font-awesome-icon icon="fa-solid fa-ellipsis" />
                </button>
                <div v-if="showOptions && isOwner" class="profile-options-menu">
                    <div class="options-menu">
                        <div class="options-menu-item" @click="getBans">
                            <span>Bans list</span>
                        </div>
                        <div class="options-menu-item" @click="deleteProfile">
                            <span>Delete profile</span>
                        </div>
                        <div class="options-menu-item" @click="doLogout">
                            <span>Logout</span>
                        </div>
                    </div>
                </div>
                <div v-else-if="showOptions" class="profile-options-menu">
                    <div class="options-menu-item" @click="banUser">
                        <span>Ban this user</span>
                    </div>
                </div>
            </div>
            <input :readonly="!isOwner" v-model="username" class="top-body-profile-username" @focusin="editingUsername"
                @focusout="saveChangeUsername" @input="checkUsername" maxlength="15" spellcheck="false">
            <div class="top-body-profile-bio-container">
                <span class="top-body-profile-bio-text-counter">{{ textCounter }}/100</span>
                <span class="top-body-profile-bio-label">Bio</span>
                <textarea :readonly="!isOwner" @focusin="editingBio" @focusout="saveChangeBio" @input="countChar"
                    v-model="bio" class="top-body-profile-bio-text" spellcheck="false" maxlength="100" rows="2"></textarea>
            </div>
            <div class="top-body-profile-stats-container">
                <div class="profile-stats" @click="goToPost">
                    <span class="profile-stats-text">posts</span>
                    <span class="profile-stats-number">{{ postsCount }}</span>
                </div>
                <div class="profile-stats" @click="getFollowers">
                    <span class="profile-stats-text">followers</span>
                    <span class="profile-stats-number">{{ followersCount }}</span>
                </div>
                <div class="profile-stats" @click="getFollowings">
                    <span class="profile-stats-text">followings</span>
                    <span class="profile-stats-number">{{ followingsCount }}</span>
                </div>
            </div>
            <div class="top-body-profile-actions" v-if="!isOwner">
                <button class="profile-actions-button follow-button" @click="follow()"> {{ followTextButton }} </button>
            </div>
        </div>
    </div>

    <ProfilesList v-if="showList" :dataGetter="dataGetter" :textHeader="textHeader" :typeList="typeList"
        @exit-list="freeLists" />
        <div v-if="showPost" class="post-view" @click.self="exitPost">
            <Post :postData="postViewData" @delete-post="deletePost" @error-occurred="(value) => { errorMsg = value }" />
        </div>

    <div class="posts-container">
        <span v-if="(posts.length === 0)" class="posts-grid-nopost-text"> There are no posts yet </span>
        <div class="posts-grid-container" v-if="posts.length > 0">
            <div v-for="post in posts" :key="post.postID" class="posts-grid-post" @click="openPost(post)">
                <img :src="`data:image/jpg;base64,${post.image}`" loading="lazy" class="posts-grid-post-image"
                    :id="post.postID">
            </div>
        </div>

    </div>
</template>