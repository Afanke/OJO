<template>

<div >
  <transition name="slide-fade">
 
  <div class="center-box" v-if="show">
    <pre style="border:1px solid black">123213
    </pre>
  </div>
  </transition>
</div>

</template>
<script>
export default {
    data() {
      return {
          show:false,
      }
    },
     created() {
      this.$bus.emit("changeHeader", "3-1")
      this.show = false
    },
    mounted() {
      this.show = true
    },
  methods: {
  
  },
  components: {}
};
</script>

<style scoped>
.center-box {
  background-color: #ffffff;
  border-radius: 10px;
  box-shadow: 0 2px 12px 0 rgba(0, 0, 0, 0.1);
}

.slide-fade-enter-active {
  transition: all 0.8s ease;
}
.slide-fade-leave-active {
  transition: all .8s cubic-bezier(1.0, 0.5, 0.8, 1.0);
}
.slide-fade-enter, .slide-fade-leave-to
/* .slide-fade-leave-active for below version 2.1.8 */ {
  transform: translateY(40px);
  opacity: 0;
}
</style>



   