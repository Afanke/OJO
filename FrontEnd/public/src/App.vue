<template>
    <div id="app">
        <div class="container">

            <el-scrollbar style="height:100%;">
                <div class="header">
                    <app-header></app-header>
                </div>
                <div class="main">
                    <router-view></router-view>
                </div>
                <transition name="fade">
                    <div class="footer" v-if="showFooter">
                        <div v-html="footer"></div>
                        <div style="color:#C0C4CC">Power By OJO</div>
                    </div>
                </transition>

            </el-scrollbar>
        </div>
    </div>
</template>
<script>
import NProgress from 'nprogress';
import 'nprogress/nprogress.css';
import AppHeader from '@/views/Header.vue';
import Scrollbar from "element-ui/lib/scrollbar"

NProgress.start();
window.onload = function () {
    NProgress.done();
};
export default {
    name: 'app',
    components: {
        appHeader: AppHeader,
        elScrollbar: Scrollbar
    },
    data() {
        return {
            footer: "",
            name: "",
            showFooter: true
        };
    },
    async beforeCreate() {
        try {
            const {
                data: res
            } = await this.$http.get("/sys/getWebConfig");
            if (res.error) {
                this.$message.error(res.error)
                return
            }
            this.name = res.data.name
            this.footer = res.data.footer
        } catch (err) {
            console.log(err);
        }
        this.$bus.emit("OJName", this.name)
    },
    watch: {
        $route() {
            this.showFooter = false
            setTimeout(() => {
                this.showFooter = true
            }, 280)
        }
    },
};
</script>
<style>
html,
body {
    height: 100%;
    width: 100%;
    padding: 0;
    margin: 0;
    border: 0;
    background-color: rgb(244, 244, 245);
    font-family: Arial,
    "微软雅黑",
    "Microsoft YaHei",
    "PingFang SC",
    "Hiragino Sans GB",
    "Helvetica Neue",
    Helvetica,
    sans-serif;
}

#nprogress .bar {
    background: #409eff !important;
    height: 2px;
}

#app {
    width: 100%;
    height: 100%;
}

.container {
    height: 100%;
}

.header {
    width: 100%;
    min-width: 1300px;
}

.main {
    margin-top: 20px;
    min-width: 1300px;
    /* float:right;
    width:auto;
    right:20px;
    left:250px; */
    background-color: rgb(244, 244, 245);
    /* position: relative; */
}

.footer {
    width: 100%;
    height: 40px;
    margin: 30px auto 0;
    text-align: center;
    font-size: small;
    /* transition: all ease 0.5s; */
}

.fade-enter-active {
    transition: opacity .3s;
}

.fade-enter,
.fade-leave-to

    /* .fade-leave-active below version 2.1.8 */
{
    opacity: 0;
}

.el-scrollbar__wrap {
    /*overflow-x: hidden !important;*/
}

.el-select-dropdown__list {
    margin: 0 0 15px 0 !important;
}
</style>