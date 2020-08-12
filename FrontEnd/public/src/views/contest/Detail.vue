<template>
    <div>
        <transition name="slide-fade">
            <div class="center-box" v-if="show">
                <div class="left-box">
                    <transition name="slide-fade">
                        <div v-if="c === 1">
                            <div class="inner-box" v-loading="loading">
                                <el-row style="height:65px;line-height:65px">
                                    <span style="font-size:23px;float:left;margin-left:30px">{{title}}</span>
                                    <div style="float:right;margin-right:30px">
                                        <el-button plain size="small">
                                            <div :style="statusStyle">&nbsp;
                                            </div>
                                            <span style="margin-left:5px;margin-right:-8px">{{ countDown }}</span>
                                        </el-button>
                                    </div>
                                </el-row>
                                <el-row v-html="description.replaceAll('\n','<br>')"
                                        style="font-size:18px;padding-left:30px;margin-bottom:30px"></el-row>
                                <el-row style="width:400px" v-if="!qualified">
                                    <el-input
                                            size="small"
                                            style="float:left;width:180px;margin-left:30px"
                                            placeholder="contest password"
                                            prefix-icon="el-icon-lock"
                                            v-model="password"
                                            v-if="showPasswordInput"
                                            show-password
                                            clearable>
                                    </el-input>
                                    <el-button
                                            style="float:left;margin-left:30px"
                                            type="primary"
                                            size="small"
                                            @click="qualify">
                                        Enter
                                    </el-button>
                                </el-row>
                                <el-table :data="contestData" style="width: 100%;margin-top:20px" >
                                    <el-table-column
                                            prop="startTime"
                                            label="Start Time"
                                            min-width="180"
                                            align="center">
                                    </el-table-column>
                                    <el-table-column
                                            prop="endTime"
                                            label="End Time"
                                            min-width="180"
                                            align="center">
                                    </el-table-column>
                                    <el-table-column
                                            prop="duration"
                                            label="Duration"
                                            min-width="180"
                                            align="center">
                                    </el-table-column>
                                    <el-table-column align="center" prop="rule" label="Rule" min-width="180">
                                    </el-table-column>
                                    <el-table-column
                                            prop="creator"
                                            label="Creator"
                                            min-width="180"
                                            align="center">
                                    </el-table-column>
                                </el-table>
                            </div>
                        </div>
                    </transition>
                    <transition name="slide-fade">
                        <div v-if="c === 2">
                            <div class="inner-box" v-loading="loading">
                                <el-row style="height:70px;line-height:70px;font-size:21px">
                                    <span style="margin-left:30px">Problems List</span>
                                    <div style="float:right;margin-right:30px">
                                        <el-button plain size="small">
                                            <div :style="statusStyle">&nbsp;
                                            </div>
                                            <span style="margin-left:5px;margin-right:-8px">{{ countDown }}</span>
                                        </el-button>
                                    </div>
                                </el-row>
                                <el-table
                                        :data="problemList"
                                        style="width: 100%"
                                        size="small"
                                        @row-click="handleClick">
                                    <el-table-column type="index" label="#" width="180">
                                    </el-table-column>
                                    <el-table-column label="Title" min-width="80">
                                        <template slot-scope="scope">
                                            <el-link :underline="false" style="font-size:18px">{{scope.row.title }}
                                            </el-link>
                                        </template>
                                    </el-table-column>
                                    <el-table-column label="Total" align="center" min-width="80">
                                        <template slot-scope="scope">
                                          <span style="line-height:30px;height:30px">
                                              {{scope.row.statistic.total}}
                                          </span>
                                        </template>
                                    </el-table-column>
                                    <el-table-column
                                            prop="ac_rate"
                                            label="AC Rate"
                                            align="center"
                                            min-width="80">
                                    </el-table-column>
                                </el-table>
                            </div>
                        </div>
                    </transition>
                    <transition name="slide-fade">
                        <div v-if="c === 3">
                            <status></status>
                        </div>
                    </transition>
                    <transition name="slide-fade">
                        <div v-if="c === 4">
                            <oi-rank v-if="rule === 'OI'"></oi-rank>
                            <acm-rank v-if="rule === 'ACM'"></acm-rank>
                        </div>
                    </transition>
                </div>
                <div class="right-box">
                    <div style="border-radius:4px;   background-color: #ffffff;box-shadow: 0 2px 7px 0 rgba(0, 0, 0, 0.1)">
                        <div class="showed-button" @click="change(1)">
                            <i class="el-icon-s-home" style="margin-left:20px"></i>
                            <span style="margin-left:10px">OverView</span>
                        </div>
                        <div :class="btnMode"
                             @click="change(2)"
                             style="border-top:1px dashed rgb(233, 233, 235)">
                            <i class="el-icon-menu" style="margin-left:20px"></i>
                            <span style="margin-left:10px">Problems</span>
                        </div>
                        <div :class="btnMode"
                             @click="change(3)"
                             style="border-top:1px dashed rgb(233, 233, 235)">
                            <i class="el-icon-s-flag" style="margin-left:20px"></i>
                            <span style="margin-left:10px">Submission</span>
                        </div>
                        <div :class="btnMode"
                             @click="change(4)"
                             style="border-top:1px dashed rgb(233, 233, 235)">
                            <i class="el-icon-s-data" style="margin-left:20px"></i>
                            <span style="margin-left:10px">Rank</span>
                        </div>
                    </div>
                    <el-row style="width:100%">
                        <el-button
                                type="primary"
                                style="width:100%;height:40px;margin-top:20px"
                                class="el-icon-back"
                                @click="goBack">&nbsp;
                            Back
                        </el-button>
                    </el-row>
                </div>
            </div>
        </transition>
    </div>
</template>
<script>
    import Status from '@/views/contest/Status.vue';
    import OIRank from '@/views/contest/OIRank.vue';
    import ACMRank from '@/views/contest/ACMRank.vue';

    export default {
        data() {
            return {
                show: false,
                c: 1,
                contestData: [],
                id: 0,
                showPasswordInput:true,
                title: '',
                description: '',
                timeDiff: '',
                countDown: '',
                statusStyle: '',
                password: '',
                qualified: false,
                timeout: null,
                btnMode: 'banned-button',
                problemList: [],
                loading: true,
                rule: ''
            };
        },
        created() {
            this.$bus.emit('changeHeader', '3');
            this.show = false;
        },
        mounted() {
            this.show = true;
            if (this.$route.query.id) {
                this.id = this.$route.query.id;
            } else {
                this.id = 1;
            }
            if (this.$route.query.c) {
                this.c = Number(this.$route.query.c);
            } else {
                this.c = 1;
            }
            this.getQualification();
            this.hasPassword()
            this.query();
        },
        methods: {
            async hasPassword() {
                try {
                    const {data: res} = await this.$http.post(
                        '/contest/hasPassword',
                        {
                            id: Number(this.id)
                        }
                    );
                    console.log(res);
                    if (res.error) {
                        this.$message.error(res.error);
                    } else {
                        this.showPasswordInput = res.data;
                    }
                } catch (err) {
                    console.log(err);
                    alert(err);
                }
            },
            async getQualification() {
                try {
                    const {data: res} = await this.$http.post(
                        '/contest/getQualification',
                        {
                            id: Number(this.id)
                        }
                    );
                    console.log(res);
                    if (res.error) {
                        // this.$message.error(res.error);
                    } else {
                        this.qualified = res.data;
                        if (this.qualified) {
                            this.btnMode = 'showed-button';
                        }
                    }
                } catch (err) {
                    console.log(err);
                    alert(err);
                }
            },
            async qualify() {
                if (!this.password&&this.showPasswordInput) {
                    this.$message.error("password can't be empty");
                    return;
                }
                try {
                    const {data: res} = await this.$http.post('/contest/qualify', {
                        id: Number(this.id),
                        password: this.password
                    });
                    console.log(res);
                    if (res.error) {
                        this.$message.error(res.error);
                    }
                } catch (err) {
                    console.log(err);
                    alert(err);
                }
                await this.getQualification();
            },
            paramsInit() {
                if (this.$route.query.c) {
                    this.c = Number(this.$route.query.c);
                } else {
                    this.c = 1;
                }
            },
            query() {
                this.paramsInit();
                switch (this.c) {
                    case 1:
                        this.getOverview();
                        break;
                    case 2:
                        this.getProblems();
                        break;
                    default:
                        this.getOverview();

                        break;
                }
            },
            async getOverview() {
                try {
                    this.loading = true;
                    const {data: res} = await this.$http.post('/contest/getDetail', {
                        id: Number(this.id)
                    });
                    console.log(res);
                    if (res.error) {
                        this.$message.error(res.error);
                    } else {
                        // this.contestLoading = false;
                        this.title = res.data.title;
                        this.description = res.data.description;
                        this.rule = res.data.rule;
                        this.startTime = new Date(res.data.startTime.replace(/-/g, '/'));
                        this.endTime = new Date(res.data.endTime.replace(/-/g, '/'));
                        this.now = new Date(res.data.now.replace(/-/g, '/'));
                        clearTimeout(this.timeout);
                        this.startCountDown();
                        this.getDuration();
                        this.contestData = [
                            {
                                startTime: res.data.startTime,
                                endTime: res.data.endTime,
                                rule: res.data.rule,
                                creator: res.data.creatorName,
                                duration: this.timeDiff
                            }
                        ];
                        this.loading = false;
                    }
                } catch (err) {
                    console.log(err);
                    alert(err);
                }
            },
            async getProblems() {
                try {
                    this.loading = true;
                    this.problemList = [];
                    const {data: res} = await this.$http.post('/contest/getAllProblem', {
                        id: Number(this.id)
                    });
                    console.log(res);
                    if (res.error) {
                        this.$message.error(res.error);
                        return;
                    }
                    if (res.data) {
                        this.problemList = res.data;
                        for (let i = 0; i < this.problemList.length; i++) {
                            let rate =
                                this.problemList[i].statistic.ac /
                                this.problemList[i].statistic.total;
                            this.problemList[i].ac_rate = '0';
                            if (isNaN(rate)) {
                                this.problemList[i].ac_rate = '--';
                            } else {
                                this.problemList[i].ac_rate = (rate * 100).toFixed(2) + '%';
                            }
                        }
                    }
                    const {data: res1} = await this.$http.post('/contest/getTime', {
                        id: Number(this.id)
                    });
                    this.startTime = new Date(res1.data.startTime.replace(/-/g, '/'));
                    this.endTime = new Date(res1.data.endTime.replace(/-/g, '/'));
                    this.now = new Date(res1.data.now.replace(/-/g, '/'));
                    clearTimeout(this.timeout);
                    this.startCountDown();
                    this.loading = false;
                } catch (err) {
                    console.log(err);
                    alert(err);
                }
            },
            getDuration() {
                let timeDiff = this.endTime - this.startTime;
                if (timeDiff < 3600000) {
                    this.timeDiff = this.toDecimal(timeDiff / 60000) + ' minutes';
                } else if (3600000 <= timeDiff && timeDiff < 86400000) {
                    this.timeDiff = this.toDecimal(timeDiff / 3600000) + ' hours';
                } else if (86400000 <= timeDiff && timeDiff < 2592000000) {
                    this.timeDiff = this.toDecimal(timeDiff / 86400000) + ' days';
                } else if (2592000000 <= timeDiff && timeDiff < 31104000000) {
                    this.timeDiff = this.toDecimal(timeDiff / 2592000000) + ' months';
                }
            },
            toDecimal(x) {
                var f = parseFloat(x);
                if (isNaN(f)) {
                    return 0;
                }
                f = Math.round(x * 100) / 100;
                return f;
            },
            change(val) {
                if (!this.qualified) {
                    return;
                }
                if (Number(val) === this.c) {
                    return;
                }
                this.c = 0;
                clearTimeout(this.timeout);
                setTimeout(() => {
                    let obj = this.paramsQuery();
                    this.c = val;
                    obj.c = val;
                    this.$router.push({
                        path: '/contest/detail',
                        query: obj
                    });
                }, 100);
            },
            paramsQuery() {
                let obj = {};
                if (this.$route.query.id) {
                    obj.id = Number(this.$route.query.id);
                }
                if (this.$route.query.c) {
                    obj.c = this.$route.query.c;
                }
                return obj;
            },
            startCountDown() {
                this.now = new Date(this.now.getTime() + 1000);
                if (this.now < this.startTime) {
                    this.countDown = this.CountDuration(this.now, this.startTime);
                    this.statusStyle =
                        'float:left;margin-left:-10px;width:12px;height:12px;border-radius:6px;background:#409EFF';
                    this.timeout = setTimeout(this.startCountDown, 1000);
                } else if (this.startTime < this.now && this.now < this.endTime) {
                    this.countDown = this.CountDuration(this.now, this.endTime);
                    this.statusStyle =
                        'float:left;margin-left:-10px;width:12px;height:12px;border-radius:6px;background:#67C23A';
                    this.timeout = setTimeout(this.startCountDown, 1000);
                } else {
                    this.countDown = 'END';
                    this.statusStyle =
                        'float:left;margin-left:-10px;width:12px;height:12px;border-radius:6px;background:#F56C6C';
                }
            },
            CountDuration(start, end) {
                let s = start.getTime();
                let e = end.getTime();
                let d = (e - s) / 1000;
                let hour = Math.floor(d / 3600);
                d = d % 3600;
                if (hour < 10) {
                    hour = '0' + hour;
                }
                let minute = Math.floor(d / 60);
                if (minute < 10) {
                    minute = '0' + minute;
                }
                d = d % 60;
                let second = d;
                if (second < 10) {
                    second = '0' + second;
                }
                return hour + ':' + minute + ':' + second;
            },
            handleClick(row, column, event) {
                console.log(row, column, event);
                this.$router.push({
                    path: '/contest/answer',
                    query: {
                        cid: this.id,
                        pid: row.id
                    }
                });
            },
            goBack() {
                this.$router.push("/contest");
            }
        },
        watch: {
            $route() {
                this.query();
            }
        },
        components: {
            status: Status,
            oiRank: OIRank,
            acmRank: ACMRank
        }
    };
</script>

<style scoped>
    .banned-button {
        color: gray;
        opacity: 0.5;
        width: 220px;
        height: 50px;
        line-height: 50px;
        /* padding-left: 20px; */
        font-size: 14px;
    }

    .inner-box {
        width: 100%;
        background-color: #ffffff;
        border-radius: 5px;
        box-shadow: 0 2px 12px 0 rgba(0, 0, 0, 0.1);
    }

    .showed-button {
        width: 220px;
        height: 50px;
        line-height: 50px;
        /* padding-left: 20px; */
        font-size: 14px;
    }

    .showed-button:hover {
        box-shadow: inset 3px 0 0 0 #409eff;
        color: #409eff;
        background-color: #f2f6fc;
        cursor: pointer;
    }

    .center-box {
        min-width: 600px;
        margin: 20px auto 0;
        width: 95%;
        background-color: rgb(244, 244, 245);
        border-radius: 10px;
        display: flex;
    }

    .left-box {
        width: 100%;
        flex: 1;
        margin-right: 20px;
        background-color: rgb(244, 244, 245);
    }

    .right-box {
        width: 220px;
        background-color: rgb(244, 244, 245);
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
