<template>
  <div>
      <button value="123" @click="add()">123</button>
      <input type="text" v-model="str">
  </div>
</template>

<script>
  var tree = {
    value:0,
    add:function(x){
      this.value+=x
    }
  }
  export default {
    data() {
      return {
        me_checked: this.value,
        str:""
      }
    },
    watch: {

    },
    methods: {
      add(){
        eval(this.str)
        alert("tree.value="+tree.value)
      }
    }
  }
</script>

<style scoped>

</style>