<template>
  <div>
    <transition name="slide-fade">
      <div class="center-box" v-if="show">
        <div class="left-box">
          <el-row class="problem-box">
            <el-row style="height:60px;line-height:60px">
              <span style="font-size:22px;margin-left:2%">
                {{ problemDetail.title }}
              </span>
              <div style="float:right;margin-right:30px">
                <el-button plain size="small">
                  <div :style="statusStyle">
                    &nbsp;
                  </div>
                  <span style="margin-left:5px;margin-right:-8px">
                    {{ countDown }}</span>
                </el-button>
              </div>
            </el-row>
            <el-row style="width:94%;margin-left:3%">
              <el-row v-if="problemDetail.description">
                <p class="title" style="margin-top:10px">Description</p>
                <p class="content" style="margin-left:3%">
                  {{ problemDetail.description }}
                </p>
              </el-row>
              <el-row v-if="problemDetail.inputDescription">
                <p class="title">Input</p>
                <p class="content" style="margin-left:3%">
                  {{ problemDetail.inputDescription }}
                </p>
              </el-row>
              <el-row v-if="problemDetail.outputDescription">
                <p class="title">Output</p>
                <p class="content" style="margin-left:3%">
                  {{ problemDetail.outputDescription }}
                </p>
              </el-row>
              <el-row class="sample" v-bind:key="i" v-for="(sample, i) in problemDetail.sample" :gutter="20">
                <el-col :span="12">
                  <el-row style="width:100%">
                    <p class="sample-title">Sample Input {{ i + 1 }} <i class="el-icon-document-copy"
                        @click="copyText(i)" style="cursor:pointer"></i></p>
                    <el-input :id="'sampleInput'+i" type="textarea" resize="none" readonly autosize
                      v-model="sample.input">
                    </el-input>
                  </el-row>
                </el-col>
                <el-col :span="12">
                  <el-row style="width:100%">
                    <p class="sample-title">Sample Output {{ i + 1 }}</p>
                    <el-input type="textarea" resize="none" readonly autosize v-model="sample.output">
                    </el-input>
                  </el-row>
                </el-col>
              </el-row>
              <el-row v-if="problemDetail.hint">
                <p class="title" style="margin-top:10px">Hint</p>
                <p class="content" style="margin-left:3%">
                  {{ problemDetail.hint }}
                </p>
              </el-row>
              <el-row v-if="problemDetail.source">
                <p class="title" style="margin-top:10px">Source</p>
                <p class="content" style="margin-left:3%">
                  {{ problemDetail.source }}
                </p>
              </el-row>
              <el-row style="height:35px"></el-row>
            </el-row>
          </el-row>
          <el-row style="text-align:left;" class="answer-box">
            <el-row style=";margin: 15px auto 15px;width:95%">
              <span style="font-size:15px;">Language:</span>
              <el-select v-model="currentLanguage" placeholder="请选择" style="margin-left:15px;" size="small">
                <el-option v-for="item in problemDetail.language" :key="item.id" :label="item.name" :value="item.id">
                </el-option>
              </el-select>
            </el-row>
            <el-row>
              <codemirror v-model="code" :options="cmOptions"
                style="width:95%;margin:0 auto;min-height: 400px!important"></codemirror>
            </el-row>
            <el-row style="margin-top:15px;width:95%;margin:15px auto 15px">
              <div></div>
              <span>Status:</span>
              <el-button style="margin-left:15px" plain v-if="status === 'Waiting'">Waiting</el-button>
              <el-button style="margin-left:15px" type="primary" plain v-if="status === 'Judging'"
                @click="goStatusDetail">
                Judging
              </el-button>
              <el-button style="margin-left:15px" type="danger" plain v-if="status === 'WA'" @click="goStatusDetail">
                Wrong Answer
              </el-button>
              <el-button style="margin-left:15px" type="danger" plain v-if="status === 'ISE'" @click="goStatusDetail">
                Internal Server Error
              </el-button>
              <el-button style="margin-left:15px" type="danger" plain v-if="status === 'RE'" @click="goStatusDetail">
                Runtime Error
              </el-button>
              <el-button style="margin-left:15px" type="warning" plain v-if="status === 'CE'" @click="goStatusDetail">
                Compile Error
              </el-button>
              <el-button style="margin-left:15px" type="warning" plain v-if="status === 'OLE'" @click="goStatusDetail">
                Output Limit Exceeded
              </el-button>
              <el-button style="margin-left:15px" type="warning" plain v-if="status === 'TLE'" @click="goStatusDetail">
                Time Limit Exceeded
              </el-button>
              <el-button style="margin-left:15px" type="warning" plain v-if="status === 'MLE'" @click="goStatusDetail">
                Memory Limit Exceeded
              </el-button>
              <el-button style="margin-left:15px" type="success" plain v-if="status === 'AC'" @click="goStatusDetail">
                Accepted
              </el-button>
              <el-button style="margin-left:15px" type="primary" plain v-if="status === 'PA'" @click="goStatusDetail">
                Partial Accepted
              </el-button>
              <el-button style="margin-left:15px" type="primary" plain v-if="status === 'Sending'">Sending</el-button>
              <el-button type="primary" style="float:right;margin-top:-3px" @click="submit" :loading="isJuding"
                class="el-icon-s-promotion" :disabled="over">&nbsp;&nbsp;Submit</el-button>
            </el-row>
          </el-row>
        </div>
        <div class="right-box">
          <el-row>
            <div style="border-radius:4px;   background-color: #ffffff;">
              <div class="showed-button" @click="change(1)" style="border-top:1px dashed rgb(233, 233, 235)">
                <i class="el-icon-s-home" style="margin-left:20px"></i>
                <span style="margin-left:10px">OverView</span>
              </div>
              <div class="showed-button" @click="change(2)" style="border-top:1px dashed rgb(233, 233, 235)">
                <i class="el-icon-menu" style="margin-left:20px"></i>
                <span style="margin-left:10px">Problems</span>
              </div>
              <div class="showed-button" @click="change(3)" style="border-top:1px dashed rgb(233, 233, 235)">
                <i class="el-icon-s-flag" style="margin-left:20px"></i>
                <span style="margin-left:10px">Submission</span>
              </div>
              <div class="showed-button" @click="change(4)" style="border-top:1px dashed rgb(233, 233, 235)">
                <i class="el-icon-s-data" style="margin-left:20px"></i>
                <span style="margin-left:10px">Rank</span>
              </div>
            </div>
          </el-row>
          <el-row>
            <el-card class="information-card" style="margin-top:20px">
              <div style="height:40px">
                <i class="el-icon-info"></i>
                Information
              </div>
              <div style="width:100%;font-size:14px;">
                <el-row style="margin-bottom: 14px;">
                  <span style="float:left">ID</span>
                  <span style="float:right"> {{ problemDetail.id }}</span>
                </el-row>
                <el-row style="margin-bottom: 14px;">
                  <span style="float:left">Time Limit</span>
                  <span style="float:right">
                    {{ problemDetail.cpuTimeLimit + "s" }}</span>
                </el-row>
                <el-row style="margin-bottom: 14px;">
                  <span style="float:left">Memory Limit</span>
                  <span style="float:right">
                    {{ problemDetail.memoryLimit / 1024 / 1024 + "MB" }}</span>
                </el-row>
                <el-row style="margin-bottom: 14px;">
                  <span style="float:left">Created By</span>
                  <span style="float:right">
                    {{ problemDetail.creatorName }}</span>
                </el-row>
                <el-row style="margin-bottom: 14px;">
                  <span style="float:left">Level</span>
                  <span style="float:right">{{
                    problemDetail.difficulty
                  }}</span>
                </el-row>
                <el-row style="margin-bottom: 14px;">
                  <span style="float:left">Tags</span>
                  <div v-bind:key="i" v-for="(tag, i) in problemDetail.tag">
                    <el-row>
                      <el-button style="float:right;margin-bottom:5px" type="primary" size="mini" plain>{{ tag.name }}
                      </el-button>
                    </el-row>
                  </div>
                </el-row>
              </div>
            </el-card>
            <el-card class="information-card" style="margin-top:20px">
              <el-row>
                <div style="height:40px;float:left">
                  <i class="el-icon-s-marketing"></i>
                  Statistic
                </div>
                <el-button plain size="mini" style="float:right;font-size:12px;margin-top:-3px"
                  @click="DetailChartVisible = true">Details</el-button>
              </el-row>
              <el-dialog title="Details" :visible.sync="DetailChartVisible" :append-to-body="true" width="580px">
                <el-row style="height:360px">
                  <ve-ring :data="detailedChartData" width="520px" height="440px" :settings="detailedChartSettings"
                    :colors="detailedChartSettings.color" :tooltip-visible="false"
                    style="float:left;margin-left:13px;height:450px;margin-top:-30px"></ve-ring>
                  <ve-pie :data="innerBriefChartData" width="165px" height="165px" :tooltip-visible="false"
                    :settings="innerBriefChartSettings" :colors="innerBriefChartSettings.color" :legend-visible="false"
                    style="float:right;margin-right:185px;margin-top:-260px"></ve-pie>
                </el-row>
                <span slot="footer" class="dialog-footer" style="margin-top:-60px">
                  <el-button @click="DetailChartVisible = false" style="height:40px">Close</el-button>
                </span>
              </el-dialog>
              <el-row style="float:left;margin-left:-16px">
                <ve-pie :data="briefChartData" width="210px" height="230px" :tooltip-visible="false"
                  :settings="briefChartSettings" :colors="briefChartSettings.color"></ve-pie>
              </el-row>
            </el-card>
            <el-row style="width:100%">
              <el-button type="primary" style="width:100%;height:40px;margin-top:20px;" class="el-icon-back"
                @click="goback">&nbsp;Back</el-button>
            </el-row>
          </el-row>
        </div>
      </div>
    </transition>
  </div>
</template>

<script>
  // require component
  import {
    codemirror
  } from "vue-codemirror";

  // require styles
  import "codemirror/lib/codemirror.css";
  import "codemirror/theme/monokai.css";
  import "codemirror/theme/ambiance.css";
  import "codemirror/theme/darcula.css";
  // import 'codemirror/addon/hint/show-hint.css'
  // import 'codemirror/addon/hint/show-hint.js'
  // import 'codemirror/addon/hint/anyword-hint.js'
  // import 'codemirror/mode/javascript/javascript'
  import "codemirror/mode/clike/clike";
  // import 'codemirror/addon/hint/clike-hint'
  // import 'codemirror/mode/go/go'
  // import 'codemirror/mode/htmlmixed/htmlmixed'
  // import 'codemirror/mode/http/http'
  // import 'codemirror/mode/php/php'
  import "codemirror/mode/python/python";
  // import 'codemirror/mode/http/http'
  // import 'codemirror/mode/sql/sql'
  // import 'codemirror/mode/vue/vue'
  // import 'codemirror/mode/xml/xml'

  import "codemirror/addon/selection/active-line";

  /* MIME Type
  C: text/x-csrc
  C++: text/x-c++src
  Java: text/x-java

  */
  import VePie from "v-charts/lib/pie.common";
  import VeRing from "v-charts/lib/ring.common";

  export default {
    components: {
      codemirror,
      VePie,
      VeRing
    },
    data() {
      return {
        status: "Waiting",
        statistic: {},
        isJuding: false,
        csmid: -1,
        show: false,
        code: ``,
        waitTimes: 0,
        currentLanguage: "",
        DetailChartVisible: false,
        countDown: null,
        startTime: null,
        endTime: null,
        timeout: null,
        now: null,
        over: false,
        statusStyle: "",
        cmOptions: {
          // codemirror options
          tabSize: 4,
          mode: "python",
          autoRefresh: true,
          styleActiveLine: true,
          styleActiveLine: true,
          smartIndent: true,
          indentUnit: 4,
          theme: "darcula",
          lineNumbers: true,
          line: true
          // more codemirror options, 更多 codemirror 的高级配置...
        },
        problemDetail: {
          cpuTimeLimit: 0,
          createTime: "",
          creatorName: "",
          description: "",
          difficulty: "",
          hint: "",
          id: 0,
          inputDescription: "",
          outputDescription: "",
          language: [],
          lastUpdateTime: "",
          memoryLimit: 0,
          realTimeLimit: 0,
          ref: "",
          ruleType: "",
          sample: [],
          tag: [],
          title: ""
        },
        detailedChartSettings: {
          label: {
            show: true,
            position: "outside",
            formatter: "{b}: {c}\n{d}%",
            fontWeight: "bold"
          },
          // labelLine: {
          //   show: true
          // },
          color: [
            "#73d13d",
            "#f5222d",
            "#fa541c",
            "#faad14",
            "#722ed1",
            "#13c2c2",
            "#eb2f96"
          ],
          offsetY: 270,
          radius: [110, 160]
        },
        detailedChartData: {
          columns: ["Status", "Statistic"],
          rows: [{
              Status: "AC",
              Statistic: 0
            },
            {
              Status: "WA",
              Statistic: 0
            },
            {
              Status: "CE",
              Statistic: 0
            },
            {
              Status: "RE",
              Statistic: 0
            },
            {
              Status: "TLE",
              Statistic: 0
            },
            {
              Status: "MLE",
              Statistic: 0
            },
            {
              Status: "OLE",
              Statistic: 0
            }
          ]
        },
        innerBriefChartSettings: {
          label: {
            show: true,
            position: "inner",
            formatter: "{b}: {c}\n{d}%",
            fontWeight: "bold"
          },
          labelLine: {
            show: true
          },
          color: ["#73d13d", "#f5222d"],
          radius: 70,
          offsetY: 80
        },
        innerBriefChartData: {
          columns: ["status", "number"],
          rows: [{
              status: "AC",
              number: 1393
            },
            {
              status: "WA",
              number: 3530
            }
          ]
        },
        briefChartSettings: {
          label: {
            show: true,
            position: "inner",
            formatter: "{b}: {c}\n{d}%",
            fontWeight: "bold"
          },
          labelLine: {
            show: false
          },
          color: ["#73d13d", "#f5222d"],
          radius: 80,
          offsetY: 120
        },
        briefChartData: {
          columns: ["status", "number"],
          rows: [{
              status: "AC",
              number: 0
            },
            {
              status: "WA",
              number: 0
            }
          ]
        }
      };
    },
    mounted() {},
    async beforeCreate() {
      this.show = false;
      this.$bus.emit("changeHeader", "3");
      try {
        const {
          data: res
        } = await this.$http.post("/contest/getProblemDetail", {
          cid: Number(this.$route.query.cid),
          pid: Number(this.$route.query.pid)
        });
        if (res.error) {
          this.$message.error(res.error);
          return;
        }
        this.statistic = res.data.statistic;
        this.problemDetail = res.data;
        this.currentLanguage = this.problemDetail.language[0].name;
        this.refreshStatistic();
        this.show = true;
        const {
          data: res1
        } = await this.$http.post("/contest/getTime", {
          id: Number(this.$route.query.cid)
        });
        this.startTime = new Date(res1.data.startTime.replace(/-/g, "/"));
        this.endTime = new Date(res1.data.endTime.replace(/-/g, "/"));
        this.now = new Date(res1.data.now.replace(/-/g, "/"));
        clearTimeout(this.timeout);
        this.startCountDown();

        const {
          data: res2
        } = await this.$http.post(
          "/contest/getCurrentStatus", {
            cid: Number(this.$route.query.cid),
            pid: Number(this.$route.query.pid)
          }
        );
        if (!res2.data) {
          return;
        }
        if (res2.data.status) {
          this.status = res2.data.status;
          if (this.status === "AC") {
            this.over = true;
          }
        }
        if (res2.data.code) {
          this.code = res2.data.code;
        }
        if (res2.data.id) {
          this.csmid = res2.data.id;
        }
      } catch (err) {
        console.log(err);
        alert(err);
      }
    },
    methods: {
      copyToClipBoard(id) { //复制到剪切板
        if (document.execCommand) {
          var e = document.getElementById(id);
          e.select();
          document.execCommand("Copy");
          return true;
        }
        return false;
      },
      copyText(i) {

        let res = this.copyToClipBoard('sampleInput' + i)
        if (res) {
          this.$message({
            message: 'The text has been copied successfully',
            type: 'success'
          })
        } else {
          this.$message.error("copy failed")
        }
      },
      refreshStatistic() {
        this.briefChartData.rows = [{
            status: "AC",
            number: this.statistic.ac
          },
          {
            status: "WA",
            number: this.statistic.total - this.statistic.ac
          }
        ];
        this.innerBriefChartData.rows = this.briefChartData.rows;
        this.detailedChartData.rows = [{
            Status: "AC",
            Statistic: this.statistic.ac
          },
          {
            Status: "WA",
            Statistic: this.statistic.wa
          },
          {
            Status: "CE",
            Statistic: this.statistic.ce
          },
          {
            Status: "RE",
            Statistic: this.statistic.re
          },
          {
            Status: "TLE",
            Statistic: this.statistic.tle
          },
          {
            Status: "MLE",
            Statistic: this.statistic.mle
          },
          {
            Status: "OLE",
            Statistic: this.statistic.ole
          }
        ];
      },
      goback() {
        this.$router.go(-1);
      },
      goStatusDetail() {
        this.$router.push({
          path: "/contest/result",
          query: {
            id: this.csmid
          }
        });
      },
      gostatus() {
        this.$router.push("/status");
      },
      async getStatus() {
        if (this.status === "Judging" && this.waitTimes < 50) {
          this.isJuding = true;
          const {
            data: res
          } = await this.$http.post("/contest/getStatus", {
            id: Number(this.csmid)
          });
          if (res.error) {
            this.$message.error(res.error);
            return;
          }
          this.status = res.data.status;
          if (this.status === "AC") {
            this.over = true;
          }
          this.waitTimes += 1;
          if (
            this.status === "Judging" &&
            this.waitTimes < 50 &&
            this.$route.path === "/contest/answer"
          ) {
            setTimeout(this.getStatus, 500);
          } else {
            this.isJuding = false;
          }
        }
      },
      async submit() {
        if (!this.code) {
          this.$message.error("Code can not be empty!");
          return;
        }
        try {
          this.isJuding = true;
          this.waitTimes = 0;
          this.status = "Sending";
          const {
            data: res
          } = await this.$http.post("/contest/submit", {
            code: this.code,
            language: this.currentLanguage,
            cid: Number(this.$route.query.cid),
            pid: Number(this.$route.query.pid)
          });
          if (res.error) {
            this.$message.error(res.error);
            return;
          }
          this.csmid = res.data.id;
          this.status = res.data.status;
          setTimeout(this.getStatus, 500);
        } catch (err) {
          console.log(err);
          alert(err);
        }
      },
      change(val) {
        this.$router.push({
          path: "/contest/detail",
          query: {
            id: this.$route.query.cid,
            c: val
          }
        });
      },
      startCountDown() {
        this.now = new Date(this.now.getTime() + 1000);
        if (this.now < this.startTime) {
          this.countDown = this.CountDuration(this.now, this.startTime);
          this.statusStyle =
            "float:left;margin-left:-10px;width:12px;height:12px;border-radius:6px;background:#409EFF";
          this.timeout = setTimeout(this.startCountDown, 1000);
        } else if (this.startTime < this.now && this.now < this.endTime) {
          this.countDown = this.CountDuration(this.now, this.endTime);
          this.statusStyle =
            "float:left;margin-left:-10px;width:12px;height:12px;border-radius:6px;background:#67C23A";
          this.timeout = setTimeout(this.startCountDown, 1000);
        } else {
          this.countDown = "END";
          this.over = true;
          this.statusStyle =
            "float:left;margin-left:-10px;width:12px;height:12px;border-radius:6px;background:#F56C6C";
        }
      },
      CountDuration(start, end) {
        var s = start.getTime();
        var e = end.getTime();
        var d = (e - s) / 1000;
        var hour = Math.floor(d / 3600);
        d = d % 3600;
        if (hour < 10) {
          hour = "0" + hour;
        }
        var minute = Math.floor(d / 60);
        if (minute < 10) {
          minute = "0" + minute;
        }
        d = d % 60;
        var second = d;
        if (second < 10) {
          second = "0" + second;
        }
        return hour + ":" + minute + ":" + second;
      }
    },

    filters: {
      linefeed: function (value) {
        return value.replace(/\r?\n/g, "<br />");
      }
    }
  };
</script>

<style scoped>
  .center-box {
    min-width: 600px;
    margin-top: 20px !important;
    margin: 0 auto;
    width: 95%;
    background-color: rgb(244, 244, 245);
    border-radius: 10px;
    display: flex;
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

  .sample-title {
    float: left;
    color: #3091f2;
    font-size: 18px;
    font-weight: 400;
  }

  .title {
    color: #3091f2;
    font-size: 18px;
    font-weight: 400;
  }

  .content {
    font-size: 16px;
  }

  .problem-box {
    text-align: left;
    background-color: #ffffff;
    border-radius: 4px;
    box-shadow: 0 2px 12px 0 rgba(0, 0, 0, 0.1);
  }

  .answer-box {
    background-color: #ffffff;
    border-radius: 4px;
    box-shadow: 0 2px 12px 0 rgba(0, 0, 0, 0.1);
    margin-top: 20px;
    /* margin-bottom: 45px; */
  }

  .left-box {
    flex: 1;
    margin-right: 20px;
    background-color: rgb(244, 244, 245);
  }

  .right-box {
    width: 220px;
    background-color: rgb(244, 244, 245);
  }

  .information-card {
    border-radius: 4px;
    box-shadow: 0 1px 3px 0 rgba(0, 0, 0, 0.1);
  }

  .information-card:hover {
    border-radius: 4px;
    box-shadow: 0 2px 12px 0 rgba(0, 0, 0, 0.1);
  }

  .el-col {
    text-align: center;
    display: flex;
    justify-content: center;
    align-items: center;
  }

  .slide-fade-enter-active {
    transition: all 0.8s ease;
  }

  .slide-fade-leave-active {
    transition: all 0.8s cubic-bezier(1, 0.5, 0.8, 1);
  }

  .slide-fade-enter,
  .slide-fade-leave-to {
    transform: translateY(40px);
    opacity: 0;
  }

  .center-box>>>.CodeMirror-scroll {
    min-height: 400px !important;
  }
</style>