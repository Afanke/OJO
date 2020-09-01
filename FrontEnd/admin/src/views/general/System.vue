<template>
    <div>
        <transition name="slide-fade">
            <div v-if="show">
                <div class="content">
                    <el-row style="height:60px;line-height:60px;">
                        <span style="font-size:20px;margin-left:20px">SMTP Config</span>
                    </el-row>
                    <el-row style="height:1px;float:top;border-top:1px solid rgb(233, 233, 235);"></el-row>
                    <el-row :gutter="20" style="margin-top:20px;margin-left:20px;margin-right:20px">
                        <el-col :span="3" style="line-height:40px">
                            <span style="color:red">*</span>
                            <span>&nbsp;Server</span>
                        </el-col>
                        <el-col :span="9">
                            <el-input v-model="SMTP.server" placeholder="Server"></el-input>
                        </el-col>
                        <el-col :span="3" style="line-height:40px">
                            <span style="color:red">*</span>
                            <span>&nbsp;Port</span>
                        </el-col>
                        <el-col :span="9">
                            <el-input v-model="SMTP.port" placeholder="Port"></el-input>
                        </el-col>
                    </el-row>
                    <el-row :gutter="20" style="margin-top:20px;margin-left:20px;margin-right:20px">
                        <el-col :span="3" style="line-height:40px">
                            <span style="color:red">*</span>
                            <span>&nbsp;Email</span>
                        </el-col>
                        <el-col :span="9">
                            <el-input v-model="SMTP.email" placeholder="Email"></el-input>
                        </el-col>
                        <el-col :span="3" style="line-height:40px">
                            <span style="color:red">*</span>
                            <span>&nbsp;Password</span>
                        </el-col>
                        <el-col :span="9">
                            <el-input v-model="SMTP.password" placeholder="Password"></el-input>
                        </el-col>
                    </el-row>
                    <el-row style="margin-top:10px">
                        <el-button style="float:left;margin-left:30px;margin-top:10px" type="primary" @click="saveSMTP">
                            Save
                        </el-button>
                    </el-row>
                    <el-row style="height:20px">
                    </el-row>
                </div>
                <div class="content" style="margin-top:20px">
                    <el-row style="height:60px;line-height:60px;">
                        <span style="font-size:20px;margin-left:20px">Web Config</span>
                    </el-row>
                    <el-row style="height:1px;float:top;border-top:1px solid rgb(233, 233, 235);"></el-row>
                    <el-row :gutter="20" style="margin-top:20px;margin-left:20px;margin-right:20px">
                        <el-col :span="3" style="line-height:40px">
                            <span style="color:red">*</span>
                            <span>&nbsp;Name</span>
                        </el-col>
                        <el-col :span="9">
                            <el-input v-model="Web.name" placeholder="Name"></el-input>
                        </el-col>
                    </el-row>
                    <el-row :gutter="20" style="margin-top:20px;margin-left:20px;margin-right:20px">
                        <el-col :span="3" style="line-height:40px">
                            <span style="color:red">*</span>
                            <span>&nbsp;Footer</span>
                        </el-col>
                        <el-col :span="21">
                            <el-input type="textarea" resize="none" :autosize="{ minRows: 2}"
                                      placeholder="<a href='xxx'>Footer</a>"
                                      v-model="Web.footer">
                            </el-input>
                        </el-col>
                    </el-row>
                    <el-row :gutter="20" style="margin-top:20px;margin-left:20px;margin-right:20px">
                        <el-col :span="3" style="line-height:40px">
                            <span style="color:red">*</span>
                            <span>&nbsp;Allow Register</span>
                        </el-col>
                        <el-col :span="21">
                            <el-switch style="margin-top:10px" v-model="Web.allowRegister" active-color="#13ce66"
                                       inactive-color="#ff4949">
                            </el-switch>
                        </el-col>
                    </el-row>
                    <el-row style="margin-top:10px">
                        <el-button @click="saveWeb" style="float:left;margin-left:30px;margin-top:10px" type="primary">
                            Save
                        </el-button>
                    </el-row>
                    <el-row style="height:20px">
                    </el-row>
                </div>
            </div>

        </transition>
    </div>

</template>
<script>
export default {
    data() {
        return {
            show: false,
            SMTP: {
                server: "",
                port: "",
                email: "",
                password: "",
            },
            Web: {
                name: "",
                footer: "",
                allowRegister: true,
            },
            judgeServer: [],
            timeout: null
        }
    },
    created() {
        this.$bus.emit("changeHeader", "2-4")
        this.show = false
    },
    mounted() {
        this.show = true
        this.getAll()
    },
    methods: {
        async getAll() {
            try {
                const {
                    data: res
                } = await this.$http.get('/admin/sys/getAll');
                if (res.error) {
                    this.$message.error(res.error)
                    return
                }
                this.SMTP.password = res.data.password
                this.SMTP.server = res.data.server
                this.SMTP.email = res.data.email
                this.SMTP.port = res.data.port
                this.Web.name = res.data.name
                this.Web.footer = res.data.footer
                this.Web.allowRegister = res.data.allowRegister
            } catch (err) {
                console.log(err);
            }
        },
        async saveSMTP() {
            try {
                const {
                    data: res
                } = await this.$http.post('/admin/sys/updateSMTP', {
                    password: this.SMTP.password,
                    server: this.SMTP.server,
                    email: this.SMTP.email,
                    port: Number(this.SMTP.port),
                });
                if (res.error) {
                    this.$message.error(res.error)
                    return
                }
                this.$message({
                    message: res.data,
                    type: 'success'
                });
            } catch (err) {
                console.log(err);
            }
        },
        async saveWeb() {
            try {
                const {
                    data: res
                } = await this.$http.post('/admin/sys/updateWeb', {
                    name: this.Web.name,
                    footer: this.Web.footer,
                    allowRegister: this.Web.allowRegister,
                });
                if (res.error) {
                    this.$message.error(res.error)
                    return
                }
                this.$message({
                    message: res.data,
                    type: 'success'
                });
            } catch (err) {
                console.log(err);
            }
        }
    },
    components: {}
};
</script>

<style scoped>
.content {
    background-color: #ffffff;
    border-radius: 10px;
    /* min-height: 600px; */
    box-shadow: 0 2px 12px 0 rgba(0, 0, 0, 0.1);
}

.slide-fade-enter-active {
    transition: all 0.8s ease;
}

.slide-fade-leave-active {
    transition: all .8s cubic-bezier(1.0, 0.5, 0.8, 1.0);
}

.slide-fade-enter,
.slide-fade-leave-to

    /* .slide-fade-leave-active for below version 2.1.8 */
{
    transform: translateY(40px);
    opacity: 0;
}
</style>