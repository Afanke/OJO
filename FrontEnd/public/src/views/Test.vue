<template>
  <div class="center-box" v-if="show">
   <el-button type="primary" @click="send">buttonCont</el-button>
  </div>
</template>
<script>
  import {
    VueCropper
  } from 'vue-cropper'

  export default {
    data() {
      return {
        show: false,

      };
    },
    created() {
      this.$bus.emit('changeHeader', '1');
      this.show = false;
    },
    async mounted() {
      this.show = true;

    },
    methods: {
     async send(){
       try {
         let {
           data: res
         } = await this.$http.post('/contest/getACMRank', {
           cid: 1
         });
         if (res.error) {
           this.$message.error(res.error);
           return;
         }
         console.log(res.data)
       } catch (err) {
         console.log(err);
         alert(err);
       }
     }
    },
    components: {}
  };
</script>

<style scoped>
  .center-box {
    min-width: 600px;
    margin: 0 auto;
    width: 90%;
    background-color: #fff;
    border-radius: 10px;
  }

  .cropper {
    width: auto;
    height: 300px;
  }
</style>