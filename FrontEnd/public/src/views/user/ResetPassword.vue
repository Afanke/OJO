<template>
    <div>
        <transition name="slide-fade">
            <div class="center-box" v-if="show">
                <transition  name="slide-fade">
                    <div class="inner-box" v-if="c === 1">
                        <h2 style="text-align: center;padding-top:20px">Reset Password</h2>
                        <el-steps style="padding-top:10px" align-center  :active="c">
                            <el-step title="Enter Your Email" ></el-step>
                            <el-step title="步骤 2" ></el-step>
                            <el-step title="Success"></el-step>
                        </el-steps>
                        <el-form :label-position="left" style="width:80%;margin:40px auto" :model="form1">
                            <el-form-item>
                                <el-input placeholder="Your Email Address" prefix-icon="el-icon-message" v-model="form1.email"></el-input>
                            </el-form-item>
                            <el-form-item>
                                <el-col :span="16">
                                    <el-input v-model="form1.captcha" prefix-icon="el-icon-question" clearable placeholder="Captcha"></el-input>
                                </el-col>
                                <el-col class="line" :span="7">
                                    <el-image style="width: 100%; height: 42px;margin-left:19px" :src="captchaUrl" fit="contain"
                                              @click="changeCaptcha"></el-image>
                                </el-col>
                            </el-form-item>
                            <el-form-item>
                                <el-button style="width: 100%" type="primary">Next Step</el-button>
                            </el-form-item>
                        </el-form>
                        <div style="height: 10px;margin-top:10px"></div>
                    </div>
                </transition>
                <transition name="slide-fade">
                    <div v-if="c === 2">

                    </div>
                </transition>
                <transition name="slide-fade">
                    <div v-if="c === 3">

                    </div>
                </transition>
            </div>
        </transition>
    </div>

</template>
<script>
    export default {
        created() {
        },
        mounted() {
            this.show = true
        },
        data() {
            return {
                c:1,
                show: false,
                form1:{
                    email:"",
                    captcha:"",
                },
                captchaUrl: this.$url + "/user/captcha",
            }
        },
        methods: {
            changeCaptcha() {
                this.captchaUrl = this.$url + "/user/captcha?k=" + Math.random();
            },
        },
        components: {}
    };
</script>

<style scoped>
    .center-box {
        width: 100%;
    }

    .inner-box{
        width: 500px;
        margin: 20px auto 0;
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
        /* .slide-fade-leave-active for below version 2.1.8 */
    {
        transform: translateY(40px);
        opacity: 0;
    }
</style>


