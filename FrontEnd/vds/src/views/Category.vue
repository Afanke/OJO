<template>
  <div>
    <div class="vux-demo">
      <!-- <img class="logo" src="../assets/vux_logo.png"> -->
      <van-cell-group title="graph">
        <van-cell style="text-align:left" :center="false" title="DFS深度优先遍历(完全)" is-link to="/graph/dfs" />
      </van-cell-group>
      <van-cell style="text-align:left" title="DFS深度优先遍历(节点至多遍历一次)" is-link to="/graph/dfs-once" />
    </div>
  </div>
</template>

<script>
  export default {
    components: {},
    data() {
      return {
        msg: 'Hello World!'
      }
    },
    methods: {
      goto() {
        console.log(123)
        this.$router.push("/graph/dfs")
      }
    }
  }
</script>

<style>
  .vux-demo {
    text-align: center;
  }

  .logo {
    width: 100px;
    height: 100px
  }
</style>