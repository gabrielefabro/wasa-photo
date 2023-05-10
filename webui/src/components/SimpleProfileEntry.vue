<script>
export default {
    emits: ["exit-list-from-entry", 'error-occurred', 'data-update'],
    props: {
        data: { type: Object, required: true },
    },
    data() {
        return {
            username: this.$props.data["username"],
            userID: this.$props.data["userID"],
        }
    },
    methods: {
        goToProfile() {
            this.$router.push(`/users/${this.userID}`);
            this.$emit('exit-list-from-entry')
        },
    },
}
</script>


<template>
    <div class="profile-entries" @click="goToProfile">
        <span class="profile-entries-username">{{ username }}</span>
    </div>
</template>


<style>
.profile-entries {
    display: flex;
    align-items: center;
    margin-bottom: 0.5em;

    cursor: pointer;

    border-top: 0.01em solid rgb(0, 0, 0, 0.3);

    padding: 0.5em 0 0 0.5em;
}


.profile-entries-username {
    font-size: 1em;
    font-weight: 600;
}
</style>