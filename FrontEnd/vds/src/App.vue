<template>
  <div id="app">
    <van-nav-bar title="VDS" left-text="返回" left-arrow @click-left="onClickLeft" @click-right="onClickRight" />
    <router-view></router-view>
  </div>
</template>

<script>
  // import $ from 'jquery'
  // // 1、定义这两个函数
  // function touchEventToMouseEvent(event, eventType) {
  //   if (!event.originalEvent || !event.originalEvent.targetTouches || event.originalEvent.targetTouches.length != 1)
  //     return false;
  //   var te = event.originalEvent.targetTouches[0];
  //   var clientX = te.clientX,
  //     clientY = te.clientY,
  //     screenX = te.screenX,
  //     screenY = te.screenY;

  //   var simEvent = new MouseEvent(eventType, {
  //     clientX: clientX,
  //     clientY: clientY,
  //     screenX: screenX,
  //     screenY: screenY,
  //     button: 0,
  //     buttons: 0
  //   });
  //   return simEvent;
  // }

  // function findElm(targetElement) {
  //   targetElement.on('touchstart', function (e) {
  //     console.log('touchstart');
  //     console.log(e);
  //     var simEvent = touchEventToMouseEvent(e, 'mousedown');
  //     if (simEvent != null) {
  //       $(this)[0].dispatchEvent(simEvent);
  //     }
  //   });

  //   targetElement.on('touchmove', function (e) {
  //     e.preventDefault();
  //     console.log('touchmove');
  //     var simEvent = touchEventToMouseEvent(e, 'mousemove');
  //     if (simEvent != null) {
  //       $(this)[0].dispatchEvent(simEvent);
  //     }
  //   });

  //   targetElement.on('touchend', function (e) {
  //     console.log('touchend');
  //     console.log(e);
  //     var simEvent = touchEventToMouseEvent(e, 'mouseup');
  //     if (simEvent != null) {
  //       $(this)[0].dispatchEvent(simEvent);
  //     }
  //   });
  // }

  // // 2、执行 findElm(selectorElement) 即可将移动端的touch
  // findElm($("#app"));
  import {
    Toast
  } from 'vant';
  export default {
    name: 'app',
    components: {},
    methods: {
      onClickLeft() {
        this.$router.go(-1)
      },
      onClickRight() {
        Toast('按钮');
      }
    }
  }
</script>

<style>
  body {
    background-color: #fbf9fe;
    height: 100%;
    overflow: hidden
  }

  #app {
    height: 100%;
  }
</style>