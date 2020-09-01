<template>
    <div>
        <transition name="slide-fade">
            <div class="center-box" v-if="show">
                <el-row>
                    <el-col :span="16">
                        <el-row>
                            <el-col :span="24">
                                <img src="../assets/images/coding.png" alt=""/>
                            </el-col>
                            <el-col :span="24">
                                <p style="font-weight:500;font-size: 60px">功能开发中</p>
                            </el-col>
                        </el-row>
                    </el-col>
                    <el-col :span="8">
                        <el-row>
                            <el-col :span="24">
                                <p style="font-weight:550;font-size: 30px">当前进度</p>
                            </el-col>
                            <el-col :span="24">
                                <el-progress type="dashboard" :percentage="fepct" :color="colors"></el-progress>
                            </el-col>
                            <el-col :span="24">
                                <p style="font-weight:500;font-size: 15px">Frontend</p>
                            </el-col>
                            <el-col :span="24">
                                <el-progress type="dashboard" :percentage="bepct" :color="colors"></el-progress>
                            </el-col>
                            <el-col :span="24">
                                <p style="font-weight:500;font-size: 15px">Backend</p>
                            </el-col>
                            <el-col :span="24">
                                <el-progress type="dashboard" :percentage="jspct" :color="colors"></el-progress>
                            </el-col>
                            <el-col :span="24">
                                <p style="font-weight:500;font-size: 15px">JudgeServer</p>
                            </el-col>

                        </el-row>
                    </el-col>

                </el-row>
            </div>
        </transition>
    </div>

</template>
<script>
export default {
    created() {
        this.$bus.emit("changeHeader", "6")
        // this.practiseListLoading=true
        this.show = false
    },
    mounted() {
        this.show = true
        setTimeout(async () => {
            const res = await this.$http.post('/getProgress')
            console.log(res)
            this.fepct = res.data.data[0].progress
            this.bepct = res.data.data[1].progress
            this.jspct = res.data.data[2].progress
        }, 1000)
    },
    data() {
        return {
            show: false,
            fepct: 0,
            bepct: 0,
            jspct: 0,
            colors: [
                {color: '#f56c6c', percentage: 20},
                {color: '#e6a23c', percentage: 40},
                {color: '#1989fa', percentage: 60},
                {color: '#6f7ad3', percentage: 80},
                {color: '#5cb87a', percentage: 100},
            ]
        }
    },
    methods: {
        logout() {
            this.$router.push("/login");
        },
        say(x) {
            alert(x)
        }
    },
    components: {}
};
</script>

<style scoped>
.center-box {
    min-width: 600px;
    margin: 20px auto 0;
    width: 80%;
    background-color: #ffffff;
    border-radius: 10px;
    box-shadow: 0 2px 12px 0 rgba(0, 0, 0, 0.1);
}

.el-col {
    text-align: center;
    display: flex;
    justify-content: center;
    align-items: center;
}

img {
    /* margin: 0 auto; */
    margin-top: 100px;
    width: 300px;
    height: 300px;
    /* margin: 0 auto; */
    /* border-radius: 50%;  */
    /* left: 50%; */
    /* transform: translate(-5%, 0); */
    background-color: #ffffff;
    /* display: flex;
    justify-content: center;
    align-items: center; */
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



   