<script>
export default {
    data() {
        return {
            file64: "",

            errorMsg: '',
        }
    },
    methods: {
        onChange() {
            const file = this.$refs.file.files[0];

            // Check file type, only jpg and jpeg are allowed
            const fileType = file.type;
            if (fileType !== "image/jpeg") {
                this.errorMsg = "File type not supported, only jpg and jpeg are allowed";
                document.querySelector('.drag-drop-area').style.backgroundColor = "#FF8989";
                return
            }

            // Check file size, max 5MB
            const fileSize = file.size;
            if (fileSize > 5242880) {
                console.log("cioa");
                this.errorMsg = "File size is too big. Max size is 5MB";
                document.querySelector('.drag-drop-area').style.backgroundColor = "#FF8989";
                return
            }

            // Convert file to base64
            this.file64 = URL.createObjectURL(file);
        },
        async createPost(postData) {
            const formData = new FormData();
            formData.append("image", postData['imageFile']);
            formData.append("caption", postData['caption']);

            try {
                let _ = await this.$axios.post(`profiles/${localStorage.userID}/posts`, formData, {
                    headers: {
                        'Authorization': `${localStorage.token}`,
                        'content-type': 'multipart/form-data'
                    }
                });
                this.$emit('refresh-data')
                setTimeout(this.$emit('exit-upload-form'), 1000);
            } catch (e) {
                this.errorMsg = this.$utils.errorToString(e);;
                this.$emit('error-occurred', this.errorMsg);
            }
        },
    }
}