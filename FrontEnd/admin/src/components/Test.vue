<template>
  <div>
<div class="weui-div" @click.stop="toggle">
    <span class="weui-switch" :class="{'weui-switch-on' : me_checked}" :value="value" ></span>
    <div v-if='!value' class="weui-div-2">All</div>
    <div v-if='value' class="weui-div-1">Mine</div>
  </div>
  </div>
</template>

<script>
 export default {
    props: {
      value: {
        type: Boolean,
        default: true
      },
    },
    data() {
      return {
        me_checked: this.value
      }
    },
    watch: {
      value(){
        this.me_checked = this.value;
      }
    },
    methods: {
      toggle() {
          this.me_checked = !this.me_checked;
          this.$emit('changeSwitch', this.me_checked);
      }
    }
  }
</script>

<style scoped>
 .weui-div{
    position:relative;
    font-weight: bold;
    cursor:pointer;
    width: 60px;
    height: 26px;
  }
  .weui-div-1{
    position:absolute;
    left:8px;
    top:0;
    line-height:28px;
    font-size:10px;
    color:#FFFFFF;;
  }
  .weui-div-2{
    position:absolute;
    right:14px;
    top:0;
    line-height:28px;
    font-size:10px;
    color:#999999;
  }
  .weui-switch {
    display: block;
    position: relative;
    width: 60px;
    height: 26px;
    border: 1px solid #DFDFDF;
    outline: 0;
    border-radius: 12px;
    box-sizing: border-box;
    background-color: #DFDFDF;
    transition: background-color 0.1s, border 0.1s;
    cursor: pointer;
  }
  .weui-switch:before {
    content: " ";
    position: absolute;
    top: 0;
    left: 0;
    width: 24px;
    height: 24px;
    border-radius: 12px;
    background-color: #E6E6E6;
    transition: transform 0.35s cubic-bezier(0.45, 1, 0.4, 1);
  }
  .weui-switch:after {
    content: " ";
    position: absolute;
    top: 0;
    left: 0;
    width: 24px;
    height: 24px;
    border-radius: 12px;
    background-color: #FFFFFF;
    box-shadow: 0 1px 3px rgba(0, 0, 0, 0.4);
    transition: transform 0.35s cubic-bezier(0.4, 0.4, 0.25, 1.35);
  }
  .weui-switch-on {
    border-color: #409eff;
    background-color: #409eff;
  }
  .weui-switch-on:before {
    border-color: #409eff;
    background-color: #409eff;
  }
  .weui-switch-on:after {
    transform: translateX(35px);
  }
</style>