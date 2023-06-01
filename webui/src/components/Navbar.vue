<script>
export default {
  data() {
    return {
      textVar: "",
      iconProfile: "fa-regular",
    };
  },
  methods: {
    logout() {
      localStorage.removeItem("token");
      this.$emit("logoutNavbar", false);
    },
    goBackHome() {
      this.$emit("requestUpdateView", "/home");
    },
    searchFunc() {
      this.$emit("searchNavbar", this.textVar);
      this.textVar = "";
    },
    myProfile() {
      this.$emit("requestUpdateView", "/users/" + localStorage.getItem("token"));
    },
  },
};

</script>

<template>
  <nav class="navbar navbar-expand-lg bg-light d-flex justify-content-between sticky-top mb-3 my-nav bg-transparent">
    <div class="col-4">
      <button @click="goBackHome" class="my-trnsp-btn me-2" type="button">
        <i class="my-nav-icon-profile fas fa-home"></i>
      </button>
    </div>

    <!-- -->
    <div class="col-4">
      <form class="form-inline my-2 my-lg-0 d-flex justify-content-center m-auto">
        <input class="form-control mr-sm-2 w-75 search-input" v-model="textVar" type="search" placeholder="Search users">
        <button class="btn btn-light ms-2" type="submit" @click.prevent="searchFunc" style="display: none;">Search</button>
      </form>
    </div>

    <div class="col-4 d-flex justify-content-end">
      <button @click="myProfile" class="my-trnsp-btn me-2" type="button">
        <!--Profile-->
        <i :class="'my-nav-icon-profile me-1 w-100 h-100 '+iconProfile+ ' fa-user'"></i>
      </button>

      <button @click="logout" class="my-trnsp-btn me-2" type="button">
        <i class="my-nav-icon-quit me-1 w-100 h-100 fa-solid fa-right-from-bracket"></i>
      </button>
    </div>
  </nav>
</template>

<style>
.my-trnsp-btn {
  background-color: transparent;
  border: none;
  padding: 0;
}

.my-nav-icon-profile {
  color: black;
  font-size: 1.5rem;
  transition: transform 0.3s;
}

.my-nav-icon-profile:hover {
  transform: scale(1.2);
}

.my-nav-icon-quit {
  color: black;
  font-size: 1.5rem;
}

.my-nav-icon-quit:hover {
  color: var(--color-red-danger);
  transform: scale(1.2);
}

.navbar-elements {
  color: black;
  font-size: 0.875rem;
}

.navbar-elements:hover {
  color: var(--color-red-danger);
}

.brand-img {
  height: 30px;
  width: 30px;
}

.search-input {
  width: 100%;
  height: 40px;
  font-size: 1rem;
}
.my-nav-icon-profile {
  color: black;
  font-size: 1.5rem;
  transition: transform 0.3s;
}

.my-nav-icon-profile:hover {
  transform: scale(1.2);
}
</style>