<template>
    <div class="box" v-if="show" v-loading="loading">
        <el-row v-if="showRank" style="height:60px">
            <span style="float:left;font-size:20px;margin-left:20px;margin-top:15px">Rank</span>
            <div  style="float:right;margin-right:30px;margin-top:15px">
                <span style="font-size: 12px;color: gray;padding-right: 10px">Auto Refresh:</span>
                <el-button plain size="small">
                    <div :style="statusStyle">
                        &nbsp;
                    </div>
                    <span style="margin-left:5px;margin-right:-8px">
                        {{ countDown }}
                    </span>
                </el-button>
            </div>
        </el-row>
        <el-row v-if="!showRank">
         <p style="color: gray;font-size: 30px;text-align:center;line-height:144px">Rank Closed</p>
        </el-row>
        <div v-if="showRank" style="width:85%;float:left;margin-left:7.5%">
            <ve-line style="width:100%" :settings="chartSettings" :legend-visible="true" :extend="chartExtend">
            </ve-line>
        </div>
        <div v-if="showRank" style="width:100%">
            <el-table :data="tableData" style="width: 100%" v-loading="rankLoading" :cell-style="cellStyle"  size="mini">
                <el-table-column type="index" label="#" min-width="10" align="center" :index="indexMethod">
                </el-table-column>
                <el-table-column prop="username" label="Username" align="center" min-width="10">
                </el-table-column>
                <el-table-column label="AC/Total" align="center" min-width="10">
                    <template slot-scope="scope">
                        <span>{{scope.row.ac}}&nbsp;/&nbsp;{{scope.row.total}}</span>
                    </template>
                </el-table-column>
                <el-table-column label="TotalTime" align="center" min-width="10">
                    <template slot-scope="scope">
                        <span v-if="scope.row.totalTime">{{countDuration(startTime,stringToDate(scope.row.totalTime))}}</span>
                    </template>
                </el-table-column>
                <el-table-column :label="item.title" min-width="10" v-bind:key="i" v-for="(item, i) in problemList"
                                 align="center">
                    <template slot="header">
                        <el-link :underline="false" @click="goProblem(item.id)">{{item.title}}</el-link>
                    </template>
                    <template slot-scope="scope">
                        <span v-if="scope.row.detail[i].total">{{countDuration(startTime,scope.row.detail[i].lastSubmitTime)}}</span><br>
                        <span v-if="scope.row.detail[i] && scope.row.detail[i].total-scope.row.detail[i].ac">(-{{scope.row.detail[i].total-scope.row.detail[i].ac}})</span>
                    </template>
                </el-table-column>
            </el-table>
            <el-pagination style="float:right;margin-top:20px" background layout="prev, pager, next" :page-size="pageSize"
                           @current-change="handlePageChange" :current-page="page" :total="count">
            </el-pagination>
        </div>

    </div>
</template>
<script>
    import VeLine from "v-charts/lib/line.common.js";
    import "echarts/lib/component/markPoint";
    import "echarts/lib/component/title";

    export default {
        data() {
            return {
                show: false,
                showRank:true,
                res: {},
                res1: {},
                tableData: [],
                problemList: [],
                startTime: "",
                page: 1,
                pageSize:5,
                statusStyle:"",
                countDown:"",
                count: 5,
                updateTime:new Date(),
                loading: true,
                rankLoading: true,
                timeout:null,
                top10: {},
                chartExtend: {
                    title: {
                        text: "TOP10",
                        left: "5%",
                        textAlign: "center"
                    },
                    legend: {
                        data: []
                    },
                    tooltip: {
                        trigger: "axis",
                        axisPointer: {
                            type: "cross",
                            animation: true,
                            snap: true
                        }
                    },
                    toolbox: {
                        show: true,
                        feature: {
                            dataZoom: {
                                //  yAxisIndex: 'none'
                            },
                            dataView: {
                                readOnly: false
                            },
                            //magicType: {type: ['line']},
                            restore: {},
                            saveAsImage: {}
                        }
                    },
                    grid: {
                        // top: 70,
                        // bottom: 50
                    },
                    xAxis: {
                        type: "time",
                        scale: true,
                        boundaryGap: ["20%", "50%"]
                    },
                    yAxis: {
                        type: "value",
                        axisPointer: {
                            snap: true,
                            type: "none"
                        },
                        scale: true
                    },
                    series: []
                },
                chartData: {
                    columns: ["date", "11"],
                    rows: []
                },
                chartSettings: {
                    xAxisType: "time"
                }
            };
        },
        created() {
            this.$bus.emit("changeHeader", "3");
            this.show = false;
        },
        async mounted() {
            this.show = true;
            try {
                const {data: res1} = await this.$http.post(
                    "/contest/getAllProblemName", {
                        id: Number(this.$route.query.id)
                    });
                if (res1.error) {
                    this.$message.error(res1.error);
                    return;
                }
                this.problemList = res1.data;
                const {data: res2} = await this.$http.post("/contest/getTime", {
                    id: Number(this.$route.query.id)
                });
                if (res2.error) {
                    this.$message.error(res2.error);
                    return;
                }
                this.startTime = res2.data.startTime;
                this.startTime = new Date(this.startTime.replace(/-/g, "/"));
                await this.getTop10()
                await this.getCount()
                await this.getRank()
                this.rankLoading = false;
                this.loading = false;
            } catch (err) {
                console.log(err);
                alert(err);
            }
        },
        methods: {
            async getTop10(){
                try {
                    const {data: res} = await this.$http.post("/contest/getACMTop10", {
                        id: Number(this.$route.query.id)
                    });
                    if (res.error) {
                        if (res.error==="rank closed"){
                            this.showRank=false
                            return
                        }
                        this.$message.error(res.error);
                        return;
                    }
                    this.top10 = res.data.rank;
                    this.processData(this.top10)
                    this.prepareChart()
                } catch (err) {
                    console.log(err);
                    alert(err);
                }
            },
            cellStyle(tb) {
                if (tb.columnIndex > 3) {
                    if (!tb.row.detail[tb.columnIndex - 4]) {
                        return 'height:57px'
                    }
                    if (tb.row.detail[tb.columnIndex - 4].firstAC) {
                        return 'background-color:#67C23A;color:#fff;height:57px'
                    }
                    if (tb.row.detail[tb.columnIndex - 4].ac) {
                        return 'background-color:rgb(225, 243, 216);height:57px'
                    }
                    if (!tb.row.detail[tb.columnIndex - 4].ac && tb.row.detail[tb.columnIndex - 4].total) {
                        return 'background-color:rgb(253, 226, 226);height:57px'
                    }
                }
            },
            processData(obj){
                for (let i = 0; i < obj.length; i++) {
                    for (let j = 0; j < obj[i].detail.length; j++) {
                        obj[i].detail[j].lastSubmitTime=this.stringToDate(obj[i].detail[j].lastSubmitTime)
                    }
                }
            },
            prepareChart() {
                if (!this.top10) {
                    return
                }
                this.chartExtend.legend.data=[]
                this.chartExtend.series=[]
                for (let i = 0; i < this.top10.length; i++) {
                    let name = this.top10[i].username;
                    this.chartExtend.legend.data.push(name);
                    let ac = 0;
                    let obj = {
                        name: name,
                        type: "line",
                        smooth: true,
                        showSymbol: true,
                        data: [
                            [this.startTime.getTime(), 0]
                        ]
                    };
                    for (let j = 0; j < this.top10[i].detail.length; j++) {
                        if (this.top10[i].detail[j].ac) {
                            ac++;
                            obj.data.push([
                                this.top10[i].detail[j].lastSubmitTime.getTime() ,
                                ac
                            ]);
                        }
                    }

                    this.chartExtend.series.push(obj);
                }
            },
            prepareTable(data) {
                if (!data) {
                    return
                }
                let firstAC=data.firstAC
                for (let k in firstAC) if(firstAC.hasOwnProperty(k)) {
                    firstAC[k]=this.stringToDate(firstAC[k])
                }
                this.tableData = data.rank
                for (let i = 0; i < this.tableData.length; i++) {
                    for (let j = 0; j < this.problemList.length; j++) {
                        let hasPb=false
                        let hasFirstAC=false
                        for (let k= 0; k < this.tableData[i].detail.length; k++) {
                            if(this.tableData[i].detail[k].pid===this.problemList[j].id){
                                hasPb=true
                                if (!hasFirstAC&&this.tableData[i].detail[k].lastSubmitTime.getTime()===firstAC[this.problemList[j].id].getTime()){
                                    this.tableData[i].detail[k].firstAC=true
                                    hasFirstAC=true
                                }
                                break
                            }
                        }
                        if (!hasPb) {
                            this.tableData[i].detail.splice(j, 0, {
                                ac: false,
                                firstAC: false,
                                lastSubmitTime: this.startTime,
                                pid: 0,
                                total: 0,
                            })
                        }

                    }
                }
            },
            async getCount() {
                const {data: res} = await this.$http.post("/contest/getACMRankCount", {
                    id: Number(this.$route.query.id)
                });
                if (res.error) {
                    if (res.error==="rank closed"){
                        this.showRank=false
                        return
                    }
                    this.$message.error(res.error);
                    return;
                }
                this.count = Number(res.data);
                this.loading = false;
            },
            async getRank() {
                try {
                    this.rankLoading = true;
                    this.paramsInit();
                    const {
                        data: res
                    } = await this.$http.post("/contest/getACMRank", {
                        cid: Number(this.$route.query.id),
                        page: this.page
                    });
                    if (res.error) {
                        if (res.error==="rank closed"){
                            this.showRank=false
                            return
                        }
                        this.$message.error(res.error);
                        return;
                    }
                    this.processData(res.data.rank)
                    this.prepareTable(res.data)
                    this.updateTime=new Date(new Date(res.data.updateTime).getTime()+60000)
                    clearTimeout(this.timeout)
                    this.startCountDown()
                    this.rankLoading = false;
                } catch (err) {
                    console.log(err);
                    alert(err);
                }
            },
            goProblem(val) {
                this.$router.push({
                    path: "/contest/answer",
                    query: {
                        cid: this.$route.query.id,
                        pid: val
                    }
                });
            },
            startCountDown() {
                let now = new Date();
                if (now < this.updateTime) {
                    let temp=((this.updateTime.getTime()-now.getTime())/1000).toFixed(0)
                    this.countDown = temp>9?temp:"0"+temp;
                    this.statusStyle = "float:left;margin-left:-10px;width:12px;height:12px;border-radius:6px;background:#67C23A";
                    this.timeout = setTimeout(this.startCountDown, 1000);
                } else  {
                    this.countDown = "- -";
                    this.statusStyle = "float:left;margin-left:-10px;width:12px;height:12px;border-radius:6px;background:#409EFF";
                    this.getTop10()
                    this.getCount()
                    this.getRank()
                }
            },
            countDuration(t1,t2) {
                let d = (t2.getTime()/1000-t1.getTime()/1000)
                let hour = Math.floor(d / 3600);
                d = d % 3600;
                if (hour < 10) {
                    hour = "0" + hour;
                }
                let minute = Math.floor(d / 60);
                if (minute < 10) {
                    minute = "0" + minute;
                }
                d = d % 60;
                let second = d;
                if (second < 10) {
                    second = "0" + second;
                }
                return hour + ":" + minute + ":" + second;
            },
            handlePageChange(val) {
                let obj = this.paramsQuery();
                obj.page = Number(val);
                this.fresh(obj);
            },
            fresh(obj) {
                this.$router.push({
                    path: "/contest/detail",
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
            },
            stringToDate(str) {
                str=str.replace("T", " ").replace("C", "")
                let tempStr = str.split(" ");
                let dateStr = tempStr[0].split("-");
                let year = parseInt(dateStr[0], 10);
                let month = parseInt(dateStr[1], 10) - 1;
                let day = parseInt(dateStr[2], 10);
                let timeStr = tempStr[1].split(":");
                let hour = parseInt(timeStr [0], 10);
                let minute = parseInt(timeStr[1], 10);
                let second = parseInt(timeStr[2], 10);
                return new Date(year, month, day, hour, minute, second);

            }
        },
        components: {
            VeLine
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
        width: 100%;
        background-color: #ffffff;
        border-radius: 10px;
        /* margin-bottom: 20px; */
        box-shadow: 0 2px 12px 0 rgba(0, 0, 0, 0.1);
    }
</style>