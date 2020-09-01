<template>
    <div>
        <transition name="slide-fade">
            <div class="center-box" v-if="show">
                <div class="left-box">
                    <div style="width:100%;height:200px">
                        <el-avatar style="width:150px;height:150px;margin-left:50px;margin-top:30px" :size="50"
                                   :src="this.$http.defaults.baseURL+detail.iconPath"></el-avatar>
                    </div>
                    <div :class="{ 'left-box-item' :true ,'left-box-item-active':c==1}" @click="c=1">
                        Profile
                    </div>
                    <div :class="{ 'left-box-item' :true ,'left-box-item-active':c==2}" @click="c=2">
                        <span> Account</span>
                    </div>
                </div>
                <div class="right-box">
                    <transition name="slide-fade">
                        <div v-if="c === 1">
                            <el-row style="font-size: 21px;font-weight: 500;margin-top:20px">
                                Avatar Setting
                            </el-row>
                            <el-upload v-if="!showCropper" style="margin-top:20px" drag action="" :auto-upload="false"
                                       :show-file-list="false" :on-change="changeUpload">
                                <i class="el-icon-upload"></i>
                                <div class="el-upload__text">Drop Here, or <em>click to select</em></div>
                                <div class="el-upload__tip" slot="tip">Only JPG / PNG files available, and no more than
                                    2 MB
                                </div>
                            </el-upload>
                            <el-row v-if="showCropper" style="margin-top:20px">
                                <el-col :span="12" class="cropper-content">
                                    <div class="cropper" style="display:flex">
                                        <vueCropper ref="cropper" style="float:left" :img="option.img"
                                                    :outputSize="option.size"
                                                    :outputType="option.outputType" :info="true" :full="option.full"
                                                    :canMove="option.canMove"
                                                    :canMoveBox="option.canMoveBox" :original="option.original"
                                                    :autoCrop="option.autoCrop"
                                                    :fixed="option.fixed" :fixedNumber="option.fixedNumber"
                                                    :centerBox="option.centerBox"
                                                    :infoTrue="option.infoTrue" :fixedBox="option.fixedBox"
                                                    @realTime="realTime"></vueCropper>
                                        <div style="flex:1;margin-left:20px;">
                                            <div style="position:relative;top:50%;transform:translateY(-50%);">
                                                <div>
                                                    <el-button type="primary" icon="el-icon-refresh-left"
                                                               @click="rotateLeft"></el-button>
                                                </div>
                                                <div>
                                                    <el-button type="primary" icon="el-icon-refresh-right"
                                                               @click="rotateRight" class="btn">
                                                    </el-button>
                                                </div>
                                                <div>
                                                    <el-button type="primary" icon="el-icon-close" @click="close"
                                                               class="btn"></el-button>
                                                </div>
                                                <div>
                                                    <el-button type="primary" icon="el-icon-check" @click="beforeUpload"
                                                               class="btn"></el-button>
                                                </div>
                                            </div>
                                        </div>
                                    </div>
                                </el-col>
                                <el-col :span="12"
                                        :style="{'width': previews.w+2 + 'px', 'height': previews.h+2 + 'px' , 'margin-left':'20px'}"
                                        style="overflow:hidden;">
                                    <div :style="previews.div" style="border:1px solid #DCDFE6;border-radius:5px">
                                        <el-avatar shape="square" :src="option.img" :style="previews.img"/>
                                    </div>
                                </el-col>
                            </el-row>

                            <el-row style="font-size:21px;font-weight:500;margin-top:20px">
                                Profile Setting
                            </el-row>
                            <el-row :gutter="40" style="margin-top:20px">
                                <el-col :span="12">
                                    <div class="input-label">Real Name</div>
                                    <el-input size="small" v-model="detail.realName">
                                    </el-input>
                                </el-col>
                                <el-col :span="12">
                                    <div class="input-label">Signature</div>
                                    <el-input size="small" v-model="detail.signature">
                                    </el-input>
                                </el-col>
                            </el-row>
                            <el-row :gutter="40" style="margin-top:20px">
                                <el-col :span="12">
                                    <div class="input-label">School</div>
                                    <el-input size="small" v-model="detail.school">
                                    </el-input>
                                </el-col>
                                <el-col :span="12">
                                    <div class="input-label">Major</div>
                                    <el-input size="small" v-model="detail.major">
                                    </el-input>
                                </el-col>
                            </el-row>
                            <el-row :gutter="40" style="margin-top:20px">
                                <el-col :span="12">
                                    <div class="input-label">Github</div>
                                    <el-input size="small" v-model="detail.github">
                                    </el-input>
                                </el-col>
                                <el-col :span="12">
                                    <div class="input-label">Blog</div>
                                    <el-input size="small" v-model="detail.blog">
                                    </el-input>
                                </el-col>
                            </el-row>
                            <el-button type="primary" size="small" style="margin:30px 0px" @click="saveProfile">Save
                                All
                            </el-button>
                            <el-dialog title="Your Avater will be set to" :visible.sync="dialogVisible" width="600px">
                                <div style="display:flex;text-align:center">
                                    <div style="flex:1">
                                        <el-avatar :size="200" :src="img"/>
                                    </div>
                                    <div style="flex:1">
                                        <el-avatar shape="square" :size="200" :src="img"/>
                                    </div>
                                </div>
                                <span slot="footer" class="dialog-footer">
                  <el-button type="primary" @click="upload" :loading="loading">Upload</el-button>
                </span>
                            </el-dialog>
                        </div>
                    </transition>
                    <transition name="slide-fade">
                        <div v-if="c === 2">
                            <div style="display:flex">
                                <div style="flex:1;border-right:1px dotted #DCDFE6">
                                    <el-row style="font-size:21px;font-weight:500;margin-top:20px">
                                        Change Password
                                    </el-row>
                                    <el-row style="margin-top:20px;margin-right:40px">
                                        <div class="input-label"><span style="color:red">*&nbsp;</span>Current Password
                                        </div>
                                        <el-input size="small" v-model="curPassword" show-password>
                                        </el-input>
                                        <div class="input-label" style="margin-top:20px"><span
                                            style="color:red">*&nbsp;</span>New Password
                                        </div>
                                        <el-input size="small" v-model="newPassword" show-password>
                                        </el-input>
                                        <div class="input-label" style="margin-top:20px"><span
                                            style="color:red">*&nbsp;</span>Confirm New
                                            Password
                                        </div>
                                        <el-input size="small" v-model="newPassword1" show-password>
                                        </el-input>
                                    </el-row>
                                    <el-button @click="savePassword" type="primary" size="small"
                                               style="margin:30px 0px">Update Password
                                    </el-button>
                                </div>
                                <div style="flex:1;">
                                    <el-row style="margin-left:40px">
                                        <el-row style="font-size:21px;font-weight:500;margin-top:20px">
                                            Change Email
                                        </el-row>
                                        <el-row style="margin-top:20px;margin-right:10px">
                                            <div class="input-label">Current Email</div>
                                            <el-input size="small" v-model="detail.email" disabled>
                                            </el-input>
                                            <div class="input-label" style="margin-top:20px"><span style="color:red">*&nbsp;</span>Current
                                                Password
                                            </div>
                                            <el-input size="small" v-model="curPassword1" show-password>
                                            </el-input>
                                            <div class="input-label" style="margin-top:20px"><span style="color:red">*&nbsp;</span>New
                                                Email
                                            </div>
                                            <el-input size="small" v-model="newEmail">
                                            </el-input>
                                        </el-row>
                                        <el-button @click="saveEmail" type="primary" size="small"
                                                   style="margin:30px 0px">Update Email
                                        </el-button>
                                    </el-row>
                                </div>
                            </div>
                        </div>
                    </transition>
                </div>
            </div>
        </transition>
    </div>
</template>
<script>
import {VueCropper} from 'vue-cropper'

export default {
    created() {
        this.$bus.emit('changeHeader', '1');
        this.show = false;
    },

    data() {
        return {
            show: false,
            dialogVisible: false,
            c: 1,
            input: "",
            loading: false,
            showCropper: false,
            img: '',
            option: {
                img: '', // 裁剪图片的地址
                info: true, // 裁剪框的大小信息
                outputSize: 1, // 裁剪生成图片的质量
                outputType: 'png', // 裁剪生成图片的格式
                canScale: true, // 图片是否允许滚轮缩放
                autoCrop: true, // 是否默认生成截图框
                fixedBox: false, // 固定截图框大小 不允许改变
                fixed: true, // 是否开启截图框宽高固定比例
                fixedNumber: [1, 1], // 截图框的宽高比例
                full: true, // 是否输出原图比例的截图
                canMoveBox: true, // 截图框能否拖动
                original: false, // 上传图片按照原始比例渲染
                centerBox: false, // 截图框是否被限制在图片里面
                infoTrue: false // true 为展示真实输出图片宽高 false 展示看到的截图框宽高
            },
            previews: {
                w: 50,
                h: 50
            },
            fileName: "",
            detail: {},
            curPassword: "",
            newPassword: "",
            newPassword1: "",
            newEmail: "",
            curPassword1: "",
        };
    },
    mounted() {
        this.show = true;
        this.getDetail();
    },
    methods: {
        realTime(data) {
            this.previews = data
        },
        changeUpload(file, fileList) {
            const isLt2M = file.size / 1024 / 1024 < 2
            if (!isLt2M) {
                this.$message.error('上传文件大小不能超过 2MB!')
                return false
            }
            this.fileName = file.name
            let reader = new FileReader();
            reader.readAsDataURL(file.raw);
            reader.onload = () => {
                this.option.img = reader.result
                this.showCropper = true
            };
        },
        params_init() {
            if (this.$route.query.c) {
                this.c = Number(this.$route.query.c);
            } else {
                this.c = 1;
            }
        },
        params_query() {
            let obj = {};
            if (this.$route.query.c) {
                obj.c = this.$route.query.c;
            }
            return obj;
        },
        rotateRight() {
            this.$refs.cropper.rotateRight()
        },
        rotateLeft() {
            this.$refs.cropper.rotateLeft()
        },
        close() {
            this.showCropper = false
        },
        beforeUpload() {
            this.$refs.cropper.getCropBlob((data) => {
                let reader = new FileReader();
                reader.readAsDataURL(data);
                reader.onload = () => {
                    this.img = reader.result
                    this.dialogVisible = true
                };
            })
        },
        upload() {
            this.loading = true
            this.$refs.cropper.getCropBlob(async (data) => {
                try {
                    let formData = new FormData()
                    formData.append("file", data, this.fileName)
                    const {
                        data: res
                    } = await this.$http.post("/user/uploadImg", formData, {
                        contentType: false,
                        processData: false,
                        headers: {
                            'Content-Type': 'application/x-www-form-urlencoded'
                        }
                    })
                    if (res.error) {
                        this.$message.error(res.error);
                        return;
                    }
                    this.showCropper = false
                    await this.getDetail()
                    this.$bus.emit("changeUserIcon")
                } catch (err) {
                    console.log(err);
                } finally {
                    this.dialogVisible = false
                    this.loading = false
                }
            })
        },
        async getDetail() {
            try {
                const {
                    data: res
                } = await this.$http.post("/user/getDetail", {})
                if (res.error) {
                    this.$message.error(res.error);
                    return;
                }
                this.detail = res.data
            } catch (err) {
                console.log(err);
            }
        },
        async saveProfile() {
            try {
                const {
                    data: res
                } = await this.$http.post("/user/updateProfile", this.detail)
                if (res.error) {
                    this.$message.error(res.error);
                    return;
                }
                this.$message.success(res.data);
                await this.getDetail()
            } catch (err) {
                console.log(err);
            }
        },
        async savePassword() {
            try {
                if (this.newPassword !== this.newPassword1) {
                    this.$message.error("new and confirmed passwords are not the same");
                    return
                }
                if (this.newPassword.length < 8) {
                    this.$message.error("new password cannot be less than 8 characters");
                    return
                }
                const {
                    data: res
                } = await this.$http.post("/user/updatePassword", {
                    password: this.curPassword,
                    new: this.newPassword
                })
                if (res.error) {
                    this.$message.error(res.error);
                    return;
                }
                this.curPassword = ""
                this.newPassword = ""
                this.newPassword1 = ""
                this.$message.success(res.data);
                await this.getDetail()
            } catch (err) {
                console.log(err);
            }
        },
        async saveEmail() {
            try {
                let reg = /^\w+((-\w+)|(\.\w+))*\@[A-Za-z0-9]+((\.|-)[A-Za-z0-9]+)*\.[A-Za-z0-9]+$/;
                let isok = reg.test(this.newEmail);
                if (!isok) {
                    this.$message.error("email format not correct");
                    return
                }
                const {
                    data: res
                } = await this.$http.post("/user/updateEmail", {
                    password: this.curPassword1,
                    new: this.newEmail
                })
                if (res.error) {
                    this.$message.error(res.error);
                    return;
                }
                this.$message.success(res.data);
                this.newEmail = ""
                this.curPassword1 = ""
                await this.getDetail()
            } catch (err) {
                console.log(err);
            }
        },
    },
    watch: {
        $route() {
            this.query();
        }
    },
    components: {
        VueCropper
    }
};
</script>

<style scoped>
.center-box {
    margin: 20px auto 0;
    width: 85%;
    background-color: #fff;
    border-radius: 10px;
    display: flex;
    box-shadow: 0 2px 12px 0 rgba(0, 0, 0, 0.1)
}

.left-box {
    width: 250px;
    border-right: 1px solid #DCDFE6
}

.btn {
    margin-top: 20px
}

.input-label {
    line-height: 32px;
    font-size: 12px;
}

.cropper {
    width: auto;
    height: 300px;
}

.left-box-item {
    width: 100%;
    height: 50px;
    line-height: 50px;
    text-align: center;
    border-right: 2px solid transparent;
    box-sizing: border-box;
    transition: all .5s ease;
}

.left-box-item-active {
    border-right: #409EFF 2px solid;
    color: #409EFF;
}

.left-box-item:hover {
    background-color: rgb(236, 245, 255);
    cursor: pointer;
}

.right-box {
    flex: 1;
    padding: 10px 40px
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