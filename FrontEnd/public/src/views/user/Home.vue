<template>
    <div>
        <transition name="slide-fade">
            <div v-if="show">
                <div class="avatar">
                    <el-avatar icon="el-icon-user-solid" :size="140"
                               :src="this.$http.defaults.baseURL+detail.iconPath"></el-avatar>
                </div>
                <div class="center-box">
                    <el-row style="height:80px"></el-row>
                    <div class="text-container">
                        <el-row style="font-weight:600;font-size:30px">
                            {{ detail.username }}
                        </el-row>
                        <el-row>
                            {{ detail.realName }}
                        </el-row>
                        <el-row>
                            <div>
                                <span>{{ detail.school }}</span>
                                <el-divider v-if="detail.major&&detail.school" direction="vertical"></el-divider>
                                <span>{{ detail.major }}</span>
                            </div>
                        </el-row>
                        <el-row>
                            {{ detail.signature }}
                        </el-row>
                    </div>
                    <el-row style="width:80%;margin-left:10%">
                        <el-divider></el-divider>
                    </el-row>
                    <el-row class="text-container2">
                        <div class="item" style="border-right:1px #DCDFE6 solid">
                            <p>Solved</p>
                            <p class="emphasize">{{ stat.ac }}</p>
                        </div>
                        <div class="item" style="border-right:1px #DCDFE6 solid">
                            <p>Submissions</p>
                            <p class="emphasize">{{ stat.submission }}</p>
                        </div>
                        <div class="item">
                            <p>Score</p>
                            <p class="emphasize">{{ stat.score }}</p>
                        </div>
                    </el-row>
                    <el-row v-if="stat.solvedList" style="text-align:center;font-size:18px;margin-top:20px">
                        List of solved problems
                    </el-row>
                    <el-row v-if="stat.solvedList" style="width:70%;margin-left:15%;margin-top:20px;text-align:center">
                        <el-button size="small" @click="gotoProblem(item)" v-for="(item,index) in stat.solvedList"
                                   :key="index">
                            {{ put0(item) }}
                        </el-button>
                    </el-row>
                    <el-row style="text-align:center;margin-top:30px">
                        <img @click="gotoGithub" style="width:30px;margin-right:20px;cursor:pointer"
                             src="../../assets/images/social-github-outline.svg" alt="">
                        <img @click="copyEmail" style="width:30px;cursor:pointer" src="../../assets/images/email.svg"
                             alt="">
                        <img @click="gotoBlog" style="width:30px;margin-left:20px;cursor:pointer"
                             src="../../assets/images/microblog.svg" alt="">
                        <textarea id="email" style="display:none">{{ detail.email }}</textarea>
                    </el-row>
                    <el-row style="margin-top:20px"></el-row>
                </div>
            </div>
        </transition>
    </div>
</template>
<script>
export default {
    created() {
        this.$bus.emit("changeHeader", "1")
        this.show = false
    },
    mounted() {
        this.show = true
        this.getDetail()
    },
    data() {
        return {
            show: false,
            username: "",
            ac: "",
            detail: {},
            stat: {}
        }
    },
    methods: {
        async getDetail() {
            try {
                const {
                    data: res0
                } = await this.$http.post("/user/getDetail", {
                    id: Number(this.$route.query.id)
                });
                if (res0.error) {
                    this.$message.error(res0.error);
                    return;
                }
                this.detail = res0.data;
                const {
                    data: res1
                } = await this.$http.post("/user/getStatistic", {
                    id: Number(this.$route.query.id)
                });
                if (res1.error) {
                    this.$message.error(res1.error);
                    return;
                }
                this.stat = res1.data
            } catch (err) {
                console.log(err);
            }
        },
        gotoProblem(id) {
            this.$router.push({
                path: "/practice/answer",
                query: {
                    id: id
                }
            });
        },
        put0(item) {
            if (item < 10) {
                return "00" + item
            }
            if (item < 100) {
                return "0" + item
            }
        },
        gotoGithub() {
            if (!this.detail.github) {
                this.$message("He/She didn't fill in GitHub")
                return
            }
            window.open(this.detail.github)
        },
        gotoBlog() {
            if (!this.detail.blog) {
                this.$message("He/She didn't fill in Blog")
                return
            }
            window.open(this.detail.blog)
        },
        copyEmail() {
            if (!this.detail.email) {
                this.$message("He/She didn't fill in Email")
                return
            }
            this.copyText("email")
        },
        copyToClipBoard(id) {
            if (document.execCommand) {
                let e = document.getElementById(id);
                e.select();
                document.execCommand("Copy");
                return true;
            }
            return false;
        },
        copyText(id) {
            let res = this.copyToClipBoard(id)
            if (res) {
                this.$message({
                    message: 'copy ' + this.detail.email + ' successfully',
                    type: 'success'
                })
            } else {
                this.$message.error("copy failed")
            }
        },
    },
    components: {}
};
</script>

<style scoped>
.avatar {
    display: flex;
    justify-content: center;
    margin-top: 10px;
    height: 0px
}

.text-container {
    text-align: center;
}

.text-container2 {
    margin-left: 10%;
    text-align: center;
    width: 80%;
    display: flex;
    flex-wrap: nowrap;
    justify-content: center;
    /* text-align: center; */
}

.emphasize {
    font-size: 20px;
    font-weight: 600;
}

.item {
    flex-grow: 1;
}

.center-box {
    margin: 90px auto 0;
    width: 75%;
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

.slide-fade-enter,
.slide-fade-leave-to

    /* .slide-fade-leave-active for below version 2.1.8 */
{
    transform: translateY(40px);
    opacity: 0;
}
</style>