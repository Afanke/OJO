<template>
<div >
  <transition name="slide-fade">
 
  <div class="center-box" v-if="show">
    <div class="star-night">
    <div class="stars">
      <div
        v-for="(item, index) in startsCount"
        :key="index"
        class="star"
        ref="star">
      </div>
    </div>
    <div class="main-box">
      <div class="carton-box">
        <h1 data-spotlight="OJO Online Judge" style="margin:200px 0">OJO Online Judge</h1>
      </div>
      <p style="width:100%;text-align:center;color:#e96f92"  >Under Construction</p>
    </div>
  </div>
    
  </div>
  </transition>
</div>
  
</template>

<script>
export default {
  data() {
    return {
      startsCount: 800,
      distance: 400,
      status:'',
      show:true,
       StatusOptions: [{
            value: 0,
            label: 'All'
          },
          // {
          //   value: 1,
          //   label: 'Not Started'
          // },
          // {
          //   value: 2,
          //   label: 'Under Way'
          // },
          // {
          //   value: 3,
          //   label: 'Ended'
          // }
        ],
    };
  },
  mounted() {
    let _this = this;
    let starArr = _this.$refs.star;
    starArr.forEach(item => {
      // alert("sad")
      var speed = 0.2 + Math.random() * 1;
      var thisDistance = _this.distance + Math.random() * 300;
      item.style.transformOrigin = '0 0 ' + thisDistance + 'px';
      item.style.transform =
        'translate3d(0,0,' +
        ((Math.random() * -thisDistance) +100)  +
        'px) rotateY(' +
        ((Math.random() * 360) +360) +
        'deg) rotateX(' +
        ((Math.random() * -100) -100) +
        'deg)'
    });
  },
  methods: {},
  components: {}
};
</script>

<style scoped>
.center-box {
  /* min-width:1100px; */
  margin: 20px auto 0;
  width: 80%;
  background-color: #ffffff;
  border-radius: 10px;
  box-shadow: 0 2px 12px 0 rgba(0, 0, 0, 0.1);
}
.main-box {
  /* position: absolute; */
  height:100%;
  top: 0;
  left: 0;
}

.star-night {
  /* position: relative; */
    border-radius: 10px;
  width: 100%;
  /* height:100%; */
  margin: 0;
  padding: 0;
  bottom: 0;
  background: radial-gradient(
    200% 100% at bottom center,
    #f7f7b6,
    #e96f92,
    #1b2947
  );
  background: radial-gradient(
    220% 105% at top center,
    #1b2947 10%,
    #77517d 40%,
    #e96f92 65%,
    #f7f7b6
  );
  /* background-attachment: fixed; */
  overflow: hidden;
}

.stars {
  transform: perspective(500px);
  transform-style: preserve-3d;
  position: absolute;
  perspective-origin: 50% 100%;
  left: 0%;
  width: 100%;
  animation: rotate 90s infinite linear;
  bottom: 0;
}
.star {
  width: 2px;
  height: 2px;
  background: #f7f7b8;
  -moz-border-radius: 1px;
  -webkit-border-radius: 1px;
  border-radius: 1px;
  position: absolute;
  top: 0;
  left: 0;
  /* backface-visibility: hidden; */
}
@keyframes rotate {
  0% {
    /* transform: perspective(300px) rotateZ(0deg) rotateX(0deg) rotateY(0); */
    transform: perspective(400px) rotateZ(20deg) rotateX(-40deg) rotateY(0);
  }
  100% {
    /* transform: perspective(300px) rotateZ(0deg) rotateX(0deg) */
    transform: perspective(400px) rotateZ(20deg) rotateX(-40deg)
      rotateY(-360deg);
  }
}

.carton-box {
  /* background-color:white; */
  margin-top:50px;
  display: flex;
  justify-content: center;
  align-items: center;
  min-height: 100px;
}
h1 {
  font-size: 100px;
  color: #e96f9200;
  font-family: Helvetica;
  margin: 0;
  padding: 0;
  /* letter-spacing: -0.3rem; */
  position: relative;
}
h1::after {
  content: attr(data-spotlight);
  color: transparent;
  position: absolute;
  top: 0;
  left: 0;
  -webkit-clip-path: ellipse(100px 100px at 0% 50%);
  clip-path: ellipse(100px 100px at 0% 50%);
  animation: spotlight 5s infinite;
  background-image: url('../assets/images/jianbian.jpg');
  background-size: 150%;
  background-position: center center;
  -webkit-background-clip: text;
  background-clip: text;
}
@keyframes spotlight {
  0% {
    -webkit-clip-path: ellipse(100px 100px at 0% 50%);
    clip-path: ellipse(100px 100px at 0% 50%);
  }

  50% {
    -webkit-clip-path: ellipse(100px 100px at 100% 50%);
    clip-path: ellipse(100px 100px at 100% 50%);
  }

  100% {
    -webkit-clip-path: ellipse(100px 100px at 0% 50%);
    clip-path: ellipse(100px 100px at 0% 50%);
  }
}
</style>
