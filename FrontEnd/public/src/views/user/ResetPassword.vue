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
                        <el-form label-position="left" :rules="rule1" style="width:80%;margin:40px auto" :model="form1">
                            <el-form-item prop="email">
                                <el-input placeholder="Your Email Address" prefix-icon="el-icon-message"
                                          v-model="form1.email"></el-input>
                            </el-form-item>
                            <el-form-item prop="captcha">
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
                                <el-button @click="sendEmail" :loading="loading" style="width: 100%" type="primary">Next
                                    Step
                                </el-button>
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
                            <captcha-box style="width: 75%;float:left;margin-left:12.5%" v-model="vcode"></captcha-box>
                            <!--              <el-button style="float:right">Resend</el-button>-->
                        </el-row>
                        <el-button @click="sendVCode" :loading="loading"
                                   style="width: 80%;margin-left:10%;margin-top:20px;margin-bottom:30px" type="primary">
                            Next Step
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
                        <el-form label-position="left" style="width:80%;margin:40px auto" :rules="rule2" :model="form2">
                            <el-form-item prop="password">
                                <el-input placeholder="New Password" prefix-icon="el-icon-lock"
                                          v-model="form2.password" show-password></el-input>
                            </el-form-item>
                            <el-form-item prop="password2">
                                <el-input placeholder="Password Again" prefix-icon="el-icon-lock"
                                          v-model="form2.password2" clearable show-password></el-input>
                            </el-form-item>
                            <el-form-item>
                                <el-button @click="sendPassword" :loading="loading" style="width: 100%" type="primary">
                                    Next Step
                                </el-button>
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
            vcode: "222",
            form1: {
                email: "",
                captcha: "",
            },
            loading: false,
            rule1: {
                email: [
                    {type: 'email', required: true, message: 'please enter your email address', trigger: 'blur'},

                ],
                captcha: [
                    {required: true, message: 'please enter captcha', trigger: 'blur'}
                ],
            },
            form2: {
                password: "",
                password2: "",
            },
            rule2: {
                password: [
                    {required: true, message: 'please enter your new password', trigger: 'blur'},
                    {
                        min: 8,
                        max: 32,
                        required: true,
                        message: 'The length of password must between 8 and 32',
                        trigger: 'blur'
                    },
                ],
                password2: [
                    {required: true, message: 'please enter your new password again', trigger: 'blur'},
                    {validator: this.getVPMethod(), trigger: 'blur'}
                ],
            },
            captchaUrl: this.$url + "/user/captcha",
        }
    },
    methods: {
        getVPMethod() {
            return (rule, value, callback) => {
                if (value !== this.form2.password) {
                    callback(new Error('The passwords are inconsistent'));
                } else {
                    callback();
                }
            }
        },
        changeCaptcha() {
            this.captchaUrl = this.$url + "/user/captcha?k=" + Math.random();
        },
        async sendEmail() {
            try {
                this.loading = true
                const {data: res} = await this.$http.post("/user/sendRPEmail", this.form1);
                if (res.error) {
                    this.$message.error(res.error);
                    this.changeCaptcha()
                    this.form1.captcha = ""
                    this.loading = false
                    return;
                }
                this.loading = false
                this.c++
            } catch (err) {
                console.log(err);
            }
        },
        async sendVCode() {
            try {
                this.loading = true
                const {
                    data: res
                } = await this.$http.post("/user/checkVCode", {
                    captcha: this.vcode
                });
                if (res.error) {
                    this.$message.error(res.error);
                    this.loading = false
                    return;
                }
                this.loading = false
                this.c++
            } catch (err) {
                console.log(err);
            }
        },

        async sendPassword() {
            try {
                this.loading = true
                const {
                    data: res
                } = await this.$http.post("/user/resetPassword", {
                    captcha: this.vcode,
                    password: this.form2.password
                });
                if (res.error) {
                    this.$message.error(res.error);
                    this.loading = false
                    return;
                }
                this.loading = false
                this.c++
            } catch (err) {
                console.log(err);
            }
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


