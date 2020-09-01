<template>
    <div class="box" v-if="show" v-loading="loading">
        <el-row v-if="showRank" style="height:60px">
          <span style="float:left;font-size:20px;margin-left:20px;margin-top:15px">
            Rank
          </span>
        </el-row>
        <el-row v-if="!showRank">
            <p style="color: gray;font-size: 30px;text-align:center;line-height:144px">Rank Closed</p>
        </el-row>
        <div v-if="showRank" style="width:85%;float:left;margin-left:7.5%">
            <ve-histogram
                :title="chartTitle"
                :data="chartData"
                style="width:100%"
                :settings="chartSettings"
                :legend-visible="false"
                :mark-point="markPoint">
            </ve-histogram>
        </div>
        <div v-if="showRank" style="width:100%">
            <el-table :data="tableData" style="width: 100%;border-radius: 10px;" v-loading="rankLoading">
                <el-table-column type="index" label="#" min-width="10" align="center" :index="indexMethod">
                </el-table-column>
                <el-table-column prop="username" label="Username" align="center" min-width="10">
                </el-table-column>
                <el-table-column label="Last Submit" align="center" min-width="10">
                    <template slot-scope="scope">
                        {{ countDuration(scope.row.lastSubmitTime) }}
                    </template>
                </el-table-column>
                <el-table-column
                    prop="totalScore"
                    label="Total Score"
                    align="center"
                    min-width="10">
                </el-table-column>
                <el-table-column
                    :label="item.title"
                    min-width="10"
                    v-bind:key="i"
                    v-for="(item, i) in problemList"
                    align="center">
                    <template slot="header">
                        <el-link :underline="false" @click="goProblem(item.id)">
                            {{ item.title }}
                        </el-link>
                    </template>
                    <template slot-scope="scope">
            <span v-if="scope.row.OIDetail[i]">
              {{ scope.row.OIDetail[i].maxScore }}
            </span>
                    </template>
                </el-table-column>
            </el-table>
            <el-pagination
                style="float:right;margin-top:20px"
                background
                layout="prev, pager, next"
                :page-size="pageSize"
                @current-change="handlePageChange"
                :current-page="page"
                :total="count">
            </el-pagination>
        </div>
    </div>
</template>
<script>
import VeHistogram from 'v-charts/lib/histogram.common.js';
import 'echarts/lib/component/markPoint';
import 'echarts/lib/component/title';
// import "v-charts/lib/title.common";
export default {
    data() {
        return {
            show: false,
            showRank: true,
            pageSize: 20,
            res: {},
            res1: {},
            tableData: [],
            problemList: [],
            startTime: '',
            page: 1,
            count: 5,
            loading: true,
            rankLoading: true,
            markPoint: {
                data: [
                    {
                        name: 'max',
                        type: 'max'
                    }
                ]
            },
            chartTitle: {
                show: true,
                left: 'center',
                text: 'Top 10',
                align: 'right'
            },
            chartData: {
                columns: ['username', 'totalScore'],
                rows: [{username: 'None', totalScore: 0}]
            },
            chartSettings: {
                itemStyle: {
                    color: '#409EFF'
                },
                labelMap: {
                    totalScore: 'Score'
                },
                legendName: {
                    totalScore: 'Score'
                }
            }
        };
    },
    created() {
        this.$bus.emit('changeHeader', '3');
        this.show = false;
    },
    async mounted() {
        this.show = true;
        try {
            const {data: res1} = await this.$http.post(
                '/contest/getAllProblemName',
                {
                    id: Number(this.$route.query.id)
                }
            );
            if (res1.error) {
                this.$message.error(res1.error);
                return;
            }
            this.problemList = res1.data;
            const {data: res2} = await this.$http.post('/contest/getTime', {
                id: Number(this.$route.query.id)
            });
            if (res2.error) {
                this.$message.error(res2.error);
                return;
            }
            this.startTime = res2.data.startTime;
            await this.getTop10()
            await this.getCount()
            this.loading = false
            await this.getRank();
        } catch (err) {
            console.log(err);
        }
    },
    methods: {
        async getTop10() {
            try {
                const {data: res} = await this.$http.post('/contest/getOITop10', {
                    id: Number(this.$route.query.id)
                });
                if (res.error) {
                    if (res.error === "rank closed") {
                        this.showRank = false
                        return
                    }
                    this.$message.error(res.error);
                    return;
                }
                this.chartData.rows = res.data;

            } catch (err) {
                console.log(err);
            }
        },
        async getCount() {
            try {
                const {data: res} = await this.$http.post('/contest/getOIRankCount', {
                    id: Number(this.$route.query.id)
                });
                if (res.error) {
                    if (res.error === "rank closed") {
                        this.showRank = false
                        return
                    }
                    this.$message.error(res.error);
                    return;
                }
                this.count = Number(res.data);
                this.loading = false
                await this.getRank();
            } catch (err) {
                console.log(err);
            }
        },
        async getRank() {
            try {
                this.rankLoading = true
                this.paramsInit();
                const {data: res} = await this.$http.post('/contest/getOIRank', {
                    cid: Number(this.$route.query.id),
                    page: this.page
                });
                if (res.error) {
                    if (res.error === "rank closed") {
                        this.showRank = false
                        return
                    }
                    this.$message.error(res.error);
                    return;
                }
                this.tableData = res.data;
                this.rankLoading = false
            } catch (err) {
                console.log(err);
            }
        },
        goProblem(val) {
            this.$router.push({
                path: '/contest/answer',
                query: {
                    cid: this.$route.query.id,
                    pid: val
                }
            });
        },
        countDuration(val) {
            let s = new Date(this.startTime.replace(/-/g, '/'));
            let e = new Date(val.replace(/-/g, '/')).getTime();
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
        handlePageChange(val) {
            let obj = this.paramsQuery();
            obj.page = Number(val);
            this.fresh(obj);
        },
        fresh(obj) {
            this.$router.push({
                path: '/contest/detail',
                query: obj
            });
        },
        paramsQuery() {
            let obj = {};
            if (this.$route.query.page) {
                obj.page = Number(this.$route.query.page);
            }
            if (this.$route.query.c) {
                obj.c = this.$route.query.c;
            }
            if (this.$route.query.id) {
                obj.id = Number(this.$route.query.id);
            }
            return obj;
        },
        paramsInit() {
            if (this.$route.query.page) {
                this.page = Number(this.$route.query.page);
            } else {
                this.page = 1;
            }
        },
        indexMethod(index) {
            return 1 + index + (this.page - 1) * this.pageSize;
        }
    },
    components: {
        VeHistogram
    },
    watch: {
        $route() {
            this.getCount()
            this.getRank()
        }
    }
};
</script>

<style scoped>
.box {
    /* min-width: 600px; */
    width: 100%;
    background-color: #ffffff;
    border-radius: 10px;
    margin-bottom: 20px;
    box-shadow: 0 2px 12px 0 rgba(0, 0, 0, 0.1);
}
</style>
