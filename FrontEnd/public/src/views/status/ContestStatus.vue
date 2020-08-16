<template>
    <div>
        <transition name="slide-fade">
            <div v-if="show">
                <div class="center-box">
                    <el-row style="height:60px">
                        <span style="float:left;font-size:20px;margin-left:20px;margin-top:15px">Contest Status</span>
                        <el-button
                                style="float:right;margin-top:15px;margin-right:20px;"
                                class="el-icon-refresh"
                                type="primary"
                                size="small"
                                @click="reset">
                            &nbsp;Reset
                        </el-button>
                    </el-row>
                    <el-row>
                        <el-table stripe size="small" :data="status" style="width:100%;border-radius:10px"
                                  v-loading="loading">
                            <el-table-column align="center" label="Submit Time" width="180">
                                <template slot-scope="scope">
                                    <i class="el-icon-time"></i>
                                    <span style="margin-left: 10px">
                                        {{scope.row.submitTime | formatDateTime}}</span>
                                </template>
                            </el-table-column>
                            <el-table-column align="center" label="Id" width="180">
                                <template slot-scope="scope">
                                    <el-link
                                            type="primary"
                                            :underline="false"
                                            @click="gotoResult(scope.row.eid)">
                                        {{ scope.row.eid.slice(0,16) }}
                                    </el-link>
                                </template>
                            </el-table-column>
                            <el-table-column align="center" label="Status" width="240">
                                <template slot-scope="scope">
                                    <el-button
                                            size="mini"
                                            :type="scope.row.flag | formatType ">
                                        {{scope.row.flag |formatFlags}}
                                    </el-button>
                                </template>
                            </el-table-column>
                            <el-table-column align="center" label="Contest">
                                <template slot-scope="scope">
                                    <el-link
                                            :underline="false"
                                            @click="gotoContest(scope.row.cid)">
                                        {{ scope.row.contestName }}
                                    </el-link>
                                </template>
                            </el-table-column>
                            <el-table-column align="center" label="Problem">
                                <template slot-scope="scope">
                                    <el-link
                                            :underline="false"
                                            @click="gotoAnswer(scope.row.cid,scope.row.pid)">
                                        {{ scope.row.problemName }}
                                    </el-link>
                                </template>
                            </el-table-column>
                            <el-table-column align="center" label="Language" min-width="80">
                                <template slot-scope="scope">
                                    <span>
                                        {{getLang(scope.row.lid) }}
                                    </span>
                                </template>
                            </el-table-column>
                            <el-table-column align="center" prop="totalScore" label="Score">
                            </el-table-column>
                        </el-table>
                    </el-row>
                </div>
                <el-row style="width:80%;margin:20px auto 0">
                    <el-pagination
                            style="float:right;"
                            hide-on-single-page
                            background=""
                            layout="prev, pager, next"
                            :page-size="pageSize"
                            @current-change="handlePageChange"
                            :current-page="page"
                            :total="count">
                    </el-pagination>
                </el-row>
            </div>
        </transition>
    </div>
</template>
<script>
    export default {
        created() {
            this.$bus.emit('changeHeader', '4-2');
        },
        async mounted() {
            this.show = false;
            this.show = true;
            this.loading = true
            await this.queryList()
        },
        data() {
            return {
                pageSize:12,
                loading: false,
                count: 0,
                page: 1,
                show: false,
                status: [],
            };
        },
        methods: {
            getLang(lid) {
                switch (lid) {
                    case 1:
                        return "C"
                    case 2:
                        return "Cpp"
                    case 3:
                        return "Java"
                    case 4:
                        return "Python"
                    case 5:
                        return "Go"
                    default:
                        this.$message.error("no such language id" + lid)
                        throw "no such language id"
                }
            },
            paramsInit() {
                if (this.$route.query.page) {
                    this.page = Number(this.$route.query.page);
                } else {
                    this.page = 1;
                }
            },
            paramsQuery() {
                let obj = {};
                if (this.$route.query.page) {
                    obj.page = Number(this.$route.query.page);
                }
                return obj;
            },
            fresh(obj) {
                this.$router.push({
                    path: '/status/contest',
                    query: obj
                })
            },
            reset() {
                this.$router.push({
                    path: '/status/contest'
                });
            },
            gotoAnswer(val1,val2) {
                this.$router.push({
                    path: '/contest/answer',
                    query: {cid: Number(val1),pid:Number(val2)}
                });
            },
            gotoContest(val) {
                this.$router.push({
                    path: '/contest/detail',
                    query: {id: Number(val)}
                });
            },
            gotoResult(val) {
                this.$router.push({
                    path: '/contest/result',
                    query: {id: val}
                });
            },
            handlePageChange(val) {
                let obj = this.paramsQuery();
                obj.page = Number(val);
                this.fresh(obj);
            },
            async queryList() {
                this.loading = true
                this.paramsInit();
                try {
                    const {data: res} = await this.$http.post('/contest/getAllStatus', {
                        page: this.page
                    });
                    if (res.error) {
                        this.$message.error(res.error);
                    } else {
                        this.status = res.data;
                        this.loading = false
                    }
                    const {data: res1} = await this.$http.post('/contest/getAllStatusCount');
                    if (res1.error) {
                        // this.$message.error(res1.error);
                    } else {
                        this.count = res1.data;
                    }
                } catch (err) {
                    console.log(err);
                    alert(err)
                }
            }
        },
        watch: {
            $route() {
                this.queryList();
            }
        },
        components: {},
        filters: {
            formatDateTime: function (value) {
                let d = new Date(value);
                return d.getFullYear() +
                    '-' +
                    (d.getMonth() + 1 < 10 ? '0' + (d.getMonth() + 1) : d.getMonth() + 1) +
                    '-' +
                    (d.getDate() < 10 ? '0' + d.getDate() : d.getDate()) +
                    ' ' +
                    (d.getHours() < 10 ? '0' + d.getDate() : d.getDate()) +
                    ':' +
                    (d.getMinutes() < 10 ? '0' + d.getMinutes() : d.getMinutes()) +
                    ':' +
                    (d.getSeconds() < 10 ? '0' + d.getSeconds() : d.getSeconds());
            },
            formatFlags: function (value) {
                switch (value) {
                    case "RE":
                        return "Runtime Error"
                    case "CE":
                        return "Compile Error"
                    case "WA":
                        return "Wrong Answer"
                    case "ISE":
                        return "Internal Server Error"
                    case "TLE":
                        return "Time Limit Exceeded"
                    case "MLE":
                        return "Memory Limit Exceeded"
                    case "OLE":
                        return "Output Limit Exceeded"
                    case "PA":
                        return "Partial Accepted"
                    case "Judging":
                        return "Judging"
                    case "Pending":
                        return "Pending"
                    case "AC":
                        return "Accepted"
                    default:
                        return "Internal Server Error"
                }
            },
            formatType: function (value) {
                switch (value) {
                    case "RE":
                    case "WA":
                    case "ISE":
                        return "danger"
                    case "TLE":
                    case "MLE":
                    case "OLE":
                    case "CE":
                        return "warning"
                    case "PA":
                    case "Judging":
                    case "Pending":
                        return "primary"
                    case "AC":
                        return "success"
                    default:
                        return "danger"
                }
            }
        }
    };
</script>

<style scoped>
    .center-box {
        min-width: 1000px;
        margin: 20px auto 0 ;
        width: 80%;
        background-color: #ffffff;
        border-radius: 10px;
        box-shadow: 0 2px 12px 0 rgba(0, 0, 0, 0.1);
    }

    .slide-fade-enter-active {
        transition: all 0.8s ease;
    }

    .slide-fade-leave-active {
        transition: all 0.8s cubic-bezier(1, 0.5, 0.8, 1);
    }

    .slide-fade-enter, .slide-fade-leave-to
        /* .slide-fade-leave-active for below version 2.1.8 */
    {
        transform: translateY(40px);
        opacity: 0;
    }
</style>
