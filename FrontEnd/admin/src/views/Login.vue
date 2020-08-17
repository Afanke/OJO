<template>
  <div class="container">
    <div class="login-wrapper">
      <div class="header">Admin Login</div>
      <div class="form-warpper">
        <input type="text" v-model="username" name="username" placeholder="username" class="input-item">
        <input type="password" v-model="password" name="password" placeholder="password" class="input-item">
        <div class="btn" @click="login">Login</div>
      </div>
      <div class="msg">
      </div>
    </div>

  </div>
</template>

<script>
  export default {
    data() {
      return {
        username:"",
        password:"",
      }
    },
    methods: {
      async login(){
         try {
          const {
            data: res
          } = await this.$http.post('/user/adminLogin', {
            username: this.username,
            password: this.password
          });
          if (res.error) {
            this.$message.error(res.error)
            return
          }
          this.$bus.emit("freshUserStatus")
        } catch (err) {
          console.log(err);
          alert(err)
        }
      },
    },
    components: {

    }
  }
</script>

<style scoped>
  
  .container {
    width: 100%;
    height: 100%;
    margin: 0;
    padding: 0;
    background-image: linear-gradient(125deg, #2c3e50, #27ae60, #2980b9, #e74c3c, #8e44ad);
    background-size: 400%;
    animation: bganimation 15s infinite;
  }

  @keyframes bganimation {
    0% {
      background-position: 0 50%;
    }

    50% {
      background-position: 100% 50%;
    }

    100% {
      background-position: 0 50%;
    }

  }

  .login-wrapper {
    background-color: rgba(255, 255, 255,1);
    width: 250px;
    height: 500px;
    border-radius: 15px;
    padding: 0 50px;
    position: relative;
    left: 50%;
    top: 50%;
    transform: translate(-50%, -50%);
  }

  .login-wrapper .header {
    font-size: 30px;
    font-weight: bold;
    text-align: center;
    line-height: 200px;
  }

  .login-wrapper .form-warpper .input-item {
    display: block;
    width: 100%;
    margin-left:-9px;
    margin-bottom: 20px;
    border: 0;
    padding: 10px;
    border-bottom: 1px solid rgb(128, 125, 125);
    font-size: 15px;
    outline: none;
  }

  .login-wrapper .form-warpper .input-item::placeholder {
    text-transform: uppercase;
  }

  .login-wrapper .form-warpper .btn {
    cursor: pointer;
    text-align: center;
    margin-left:-9px;
    padding: 10px;
    width: 100%;
    margin-top: 40px;
    background-image: linear-gradient(to right, #a6c1ee, #fbc2eb);
    color: #fff;
  }

  .login-wrapper .msg {
    text-align: center;
    line-height: 80px;
  }

  .login-wrapper .msg a {
    text-decoration-line: none;
    color: #a6c1ee;
  }
</style>