<template>
    <div>
        <transition name="slide-fade">
            <div class="center-box" v-if="show">
                <transition name="slide-fade">
                    <div class="inner-box" v-if="c === 1">
                        <h2 style="text-align: center;padding-top:20px">Reset Password</h2>
                        <el-steps style="padding-top:10px" finish-status="success" align-center :active="c-1">
                            <el-step title="Enter Email"></el-step>
                            <el-step title="Check Email"></el-step>
                            <el-step title="Set Password"></el-step>
                        </el-steps>
                        <el-form label-position="left" style="width:80%;margin:40px auto" :model="form1">
                            <el-form-item>
                                <el-input placeholder="Your Email Address" prefix-icon="el-icon-message"
                                          v-model="form1.email"></el-input>
                            </el-form-item>
                            <el-form-item>
                                <el-col :span="16">
                                    <el-input v-model="form1.captcha" prefix-icon="el-icon-question" clearable
                                              placeholder="Captcha"></el-input>
                                </el-col>
                                <el-col class="line" :span="7">
                                    <el-image style="width: 100%; height: 42px;margin-left:19px" :src="captchaUrl"
                                              fit="contain"
                                              @click="changeCaptcha"></el-image>
                                </el-col>
                            </el-form-item>
                            <el-form-item>
                                <el-button @click="c=2" style="width: 100%" type="primary">Next Step</el-button>
                            </el-form-item>
                        </el-form>
                        <div style="height: 10px;margin-top:10px"></div>
                    </div>
                </transition>
                <transition name="slide-fade">
                    <div class="inner-box" v-if="c === 2">
                        <h2 style="text-align: center;padding-top:20px">Reset Password</h2>
                        <el-steps style="padding-top:10px" finish-status="success" align-center :active="c-1">
                            <el-step title="Enter Email"></el-step>
                            <el-step title="Check Email"></el-step>
                            <el-step title="Set Password"></el-step>
                        </el-steps>
                        <p style="text-align: center;margin-top:20px">We have sent the verification code to your
                            email.</p>
                        <p style="text-align: center">Please check your email and enter the verification code below</p>
                        <el-row style="width: 80%;margin-left: 10%;margin-top:20px">
                            <captcha-box style="width: 75%;float:left"></captcha-box>
                            <el-button style="float:right">Resend</el-button>
                        </el-row>
                        <el-button @click="c=3" style="width: 80%;margin-left:10%;margin-top:20px;margin-bottom:30px"
                                   type="primary">Next Step
                        </el-button>
                    </div>
                </transition>
                <transition name="slide-fade">
                    <div class="inner-box" v-if="c === 3">
                        <h2 style="text-align: center;padding-top:20px">Reset Password</h2>
                        <el-steps style="padding-top:10px" finish-status="success" align-center :active="c-1">
                            <el-step title="Enter Email"></el-step>
                            <el-step title="Check Email"></el-step>
                            <el-step title="Set Password"></el-step>
                        </el-steps>
                        <p style="text-align: center;margin-top:20px">Please enter your new password</p>
                        <el-form label-position="left" style="width:80%;margin:40px auto" :model="form2">
                            <el-form-item>
                                <el-input placeholder="New Password" prefix-icon="el-icon-lock"
                                          v-model="form2.password"></el-input>
                            </el-form-item>
                            <el-form-item>
                                <el-input placeholder="Password Again"  prefix-icon="el-icon-lock"
                                          v-model="form2.password2" clearable></el-input>
                            </el-form-item>
                            <el-form-item>
                                <el-button @click="c=4" style="width: 100%" type="primary">Next Step</el-button>
                            </el-form-item>
                        </el-form>
                        <div style="height: 10px;margin-top:10px"></div>
                    </div>
                </transition>
                <transition name="slide-fade">
                    <div class="inner-box" v-if="c === 4">
                        <h2 style="text-align: center;padding-top:20px">Reset Password</h2>
                        <el-steps style="padding-top:10px" finish-status="success" align-center :active="c-1">
                            <el-step title="Enter Email"></el-step>
                            <el-step title="Check Email"></el-step>
                            <el-step title="Set Password"></el-step>
                        </el-steps>
                        <p style="text-align: center;margin-top:20px">Reset password successfully</p>
                        <p style="text-align: center">Please log in again</p>
                        <div style="height: 10px;margin-top:10px"></div>
                    </div>
                </transition>
            </div>
        </transition>
    </div>

</template>
<script>
    import captchaBox from '@/components/CaptchaBox.vue';

    export default {
        created() {
        },
        mounted() {
            this.show = true
        },
        data() {
            return {
                c: 1,
                show: false,
                form1: {
                    email: "",
                    captcha: "",
                },
                form2: {
                    password: "",
                    password2: "",
                },
                captchaUrl: this.$url + "/user/captcha",
            }
        },
        methods: {
            changeCaptcha() {
                this.captchaUrl = this.$url + "/user/captcha?k=" + Math.random();
            },
        },
        components: {
            captchaBox: captchaBox
        }
    };
</script>

<style scoped>
    .center-box {
        width: 100%;
    }

    .inner-box {
        width: 500px;
        margin: 20px auto 0;
        background-color: #ffffff;
        border-radius: 10px;
        box-shadow: 0 2px 12px 0 rgba(0, 0, 0, 0.1);
    }

    .slide-fade-enter-active {
        transition: all 0.8s ease;
    }

    .slide-fade-enter
        /* .slide-fade-leave-active for below version 2.1.8 */
    {
        transform: translateY(40px);
        opacity: 0;
    }
</style>


