<template>
    <div id="app">
        <el-container v-if="isAdmin && show">
            <el-header style="padding:0;background-color:rgb(102, 177, 255);color:#fff;" height="50px">
                <div style="float:left;margin-left:20px;line-height:50px;font-size:25px">
                    Hello OJO
                </div>
                <el-dropdown style="float:right;margin-right:40px;margin-top:7px" @command="handleCommand">
                    <el-button size="medium">
                        {{ username }}<i class="el-icon-arrow-down el-icon--right"></i>
                    </el-button>
                    <el-dropdown-menu slot="dropdown">
                        <!-- <el-dropdown-item>Go Back</el-dropdown-item> -->
                        <el-dropdown-item command="logout">Log Out</el-dropdown-item>
                    </el-dropdown-menu>
                </el-dropdown>
                <div class="full-screen" @click="changeScreen">
                    <i class="el-icon-full-screen"></i>
                </div>
            </el-header>
            <el-container style="height:100%">
                <el-aside width="230px" style="height:100%;border-right: solid 1px #e6e6e6;background-color:#fff">
                    <app-header></app-header>
                </el-aside>
                <el-main style="min-width:1050px">
                    <router-view>
                    </router-view>
                </el-main>
                <!-- <el-footer> -->
                <!-- </el-footer> -->
            </el-container>
        </el-container>
        <login v-if="!isAdmin && show"></login>
    </div>
</template>
<script>
import Header from '@/views/Header.vue';
import Login from '@/views/Login.vue';
import Scrollbar from "element-ui/lib/scrollbar"


export default {
    name: 'app',
    components: {
        appHeader: Header,
        elScrollbar: Scrollbar,
        login: Login
    },
    data() {
        return {
            username: "",
            isFullScreen: false,
            isAdmin: false,
            show: false,
        };
    },
    created() {
        this.getUserInfo()
        this.$bus.on("freshUserStatus", this.getUserInfo)
    },
    methods: {
        handleCommand(command) {
            switch (command) {
                case "logout":
                    this.logout()
                    break
            }
        },
        async getUserInfo() {
            try {
                const {
                    data: res0
                } = await this.$http.post('/admin/user/getDetail', {});
                if (res0.error) {
                    // this.$message.error(res0.error)
                    this.isAdmin = false
                    return
                }
                if (res0.data.type < 2) {
                    this.isAdmin = false
                    return
                }
                this.isAdmin = true
                this.username = res0.data.username
            } catch (err) {
                console.log(err);
            } finally {
                this.show = true
            }
        },
        async logout() {
            try {
                const {
                    data: res0
                } = await this.$http.post('/user/logout', {});
                if (res0.error) {
                    this.$message.error(res0.error)
                    return
                }
                this.$bus.emit("freshUserStatus")
            } catch (err) {
                console.log(err);
            }
        },
        changeScreen() {
            if (this.isFullScreen) {
                this.isFullScreen = false
                this.exitScreen()
                return
            }
            this.isFullScreen = true
            this.fullScreen()
        },
        fullScreen() {
            const el = document.documentElement;
            const rfs = el.requestFullScreen || el.webkitRequestFullScreen || el.mozRequestFullScreen || el
                .msRequestFullScreen;
            if (rfs) {
                rfs.call(el);
            } else if (typeof window.ActiveXObject !== "undefined") {
                const wscript = new ActiveXObject("WScript.Shell");
                if (wscript != null) {
                    wscript.SendKeys("{F11}");
                }
            }
        },
        exitScreen() {
            const el = document;
            const cfs = el.cancelFullScreen || el.webkitCancelFullScreen || el.mozCancelFullScreen || el.exitFullScreen;
            if (cfs) {
                cfs.call(el);
            } else if (typeof window.ActiveXObject !== "undefined") {
                const wscript = new ActiveXObject("WScript.Shell");
                if (wscript != null) {
                    wscript.SendKeys("{F11}");
                }
            }
        }
    }
}
</script>
<style>
.full-screen {
    float: right;
    font-size: 20px;
    margin-top: 14px;
    margin-right: 20px;
    cursor: pointer;
}

html,
body {
    height: 100%;
    width: 100%;
    padding: 0;
    margin: 0;
    border: 0;
    background-color: rgb(244, 244, 245);
    font-family: "Helvetica Neue",
    Helvetica,
    "PingFang SC",
    "Hiragino Sans GB",
    "Microsoft YaHei",
    "微软雅黑",
    Arial,
    sans-serif;
}

#app {
    width: 100%;
    height: 100%;
    display: flex;
}

.el-aside {
    height: 100%;
    background-color: #ffffff;
}

.el-scrollbar__wrap.default-scrollbar__wrap {
    overflow-x: auto;
}
</style>