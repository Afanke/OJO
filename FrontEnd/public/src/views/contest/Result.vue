<template>
  <div>
    <transition name="slide-fade">
      <div class="center-box" v-if="show">
        <el-row v-if="bigFlag.show" :style="bigFlag.bgCss">
          <el-col style="width:50px;">
            <div style="height:100%;width:20%;font-size:30px">
              <span :class="bigFlag.iconCls" :style="bigFlag.iconCss"></span>
            </div>
          </el-col>
          <el-col style="width:800px;margin-left:-55px">
            <div style="height:100%;width:80%;text-align:left">
              <p style="font-size:20px;width:400px">{{ bigFlag.flag }}</p>
              <p style="font-size:14px">
                <!-- <span style="margin-right:20px">Author:</span> -->
                <span style="margin-right:20px">Problem:&nbsp;{{ result.problemName }}</span>
                <span style="margin-right:20px">Author:&nbsp;{{ result.username }}</span>
                <span style="margin-right:20px">Language:&nbsp;{{ getLang(result.lid)  }}</span>
                <span style="margin-right:20px">Score:&nbsp;{{ result.totalScore }}</span>
              </p>
            </div>
          </el-col>
        </el-row>
        <el-row v-if="result.errorMsg" style="margin-top:20px" :class="getBgClass()">
          <p v-html="result.errorMsg.replaceAll('\n','<br>')"></p>
        </el-row>
        <el-row style="background-color:#ffffff;margin-top:20px;border-radius: 4px;overflow: hidden">
          <el-row>
            <span style="font-size:20px;margin-top:15px;margin-left:20px;float:left" v-if="!option">Status Detail</span>
            <span style="font-size:20px;margin-top:15px;margin-left:20px;float:left" v-if="option">Code</span>
            <el-switch style="display: block;float:right;margin-top:15px;margin-right:2%;" v-model="option"
                       active-color="#409eff" inactive-color="#409eff" :active-value="true" :inactive-value="false"
                       active-text="Code" inactive-text="Status Detail">
            </el-switch>
          </el-row>

          <el-table v-if="!option" :data="resultDetail" size="mini" style="width: 100%;margin-top:15px;">
            <el-table-column type="expand">
              <template slot-scope="props">
                <el-form label-position="left" inline class="demo-table-expand">
                  <el-form-item label="Flag:">
                    <span>{{ props.row.flag }}</span>
                  </el-form-item>
                  <el-form-item label="CpuTime:">
                    <span>{{ props.row.cpuTime }}&nbsp;ms</span>
                  </el-form-item>
                  <el-form-item label="RealTime:">
                    <span>{{ props.row.realTime }}&nbsp;ms</span>
                  </el-form-item>
                  <el-form-item label="Memory(RSS):">
                    <span>{{ props.row.realMemory }}&nbsp;KB</span>
                  </el-form-item>
                  <el-form-item label="Score:">
                    <span>{{ props.row.score }}</span>
                  </el-form-item>
                  <el-form-item label="Output:" v-if="props.row.realOutput">
                    <span  v-html="props.row.realOutput.replaceAll('\n','<br>')"></span>
                  </el-form-item>
                  <el-form-item v-if="props.row.errorOutput" label="ErrorOutput:" >
                    <span style="white-space: pre-wrap;" v-html="props.row.errorOutput.replaceAll('\n','<br>')"></span>
                  </el-form-item>
                </el-form>
              </template>
            </el-table-column>
            <el-table-column label="#">
              <template slot-scope="scope">
                <p>{{ scope.$index + 1 }}</p>
              </template>
            </el-table-column>
            <el-table-column label="Status" align="center">
              <template slot-scope="scope">
                <el-button size="mini" :type="scope.row.flag | formatType">
                  {{ scope.row.flag | formatFlags }}
                </el-button>
              </template>
            </el-table-column>
            <el-table-column label="CpuTime" align="center">
              <template slot-scope="scope">
                <p>{{ scope.row.cpuTime }}ms</p>
              </template>
            </el-table-column>
            <el-table-column label="Memory" align="center">
              <template slot-scope="scope">
                <p>{{ scope.row.realMemory  }}KB</p>
              </template>
            </el-table-column>
            <el-table-column label="Score" prop="score" align="center">
            </el-table-column>
          </el-table>
          <div class="resultCode">
            <codemirror ref="myCm" style="margin-top:15px;margin-left:2%;width:98%;"
                        v-if="option" :value="result.code" :options="cmOptions">
            </codemirror>
          </div>
        </el-row>
        <el-row style="width:auto;margin-top:20px;">
          <el-button type="primary" style="float:right" class="el-icon-back" @click="goBack">Back</el-button>
        </el-row>
      </div>
    </transition>
  </div>
</template>
<script>
  import {
    codemirror
  } from 'vue-codemirror';

  // require styles
  import 'codemirror/lib/codemirror.css';
  import 'codemirror/mode/python/python';
  import 'codemirror/theme/idea.css';
  import 'codemirror/addon/selection/active-line';
  export default {
    data() {
      return {
        errorMsg:"",
        show: false,
        option: 0,
        waitTimes: 0,
        csmid: -1,
        bigFlag: {
          show: false,
          bgCss: '',
          flag: '',
          iconCls: '',
          iconCss: ''
        },
        result: {
          id: 0,
          pid: 0,
          flag: '',
          submitTime: '',
          totalScore: 0,
          uid: 0
        },
        resultDetail: [],
        cmOptions: {
          // codemirror options
          tabSize: 4,
          readOnly: true,
          mode: 'python',
          theme: 'idea',
          lineNumbers: false,
          line: true
          // more codemirror options, 更多 codemirror 的高级配置...
        }
      };
    },
    created() {
      this.$bus.emit('changeHeader', '4');
    },
    async beforeCreate() {
      try {
        let csmid = this.$route.query.id;
        let {
          data: res0
        } = await this.$http.post('/contest/getStatus', {
          id: csmid
        });
        if (res0.error) {
          this.$message.error(res0.error);
          return;
        }
        this.result = res0.data;
        this.refreshBigFlag();
        if (this.result.flag !== 'JUG') {
          await this.getStatusDetail();
        } else {
          await this.getStatus()
        }
      } catch (err) {
        console.log(err);
        alert(err);
      }
    },
    mounted() {
      this.show = false;
      this.show = true;
    },
    methods: {
      getBgClass(){
        switch (this.result.flag) {
          case 'RE':
          case 'WA':
          case 'ISE':
            return "red-container"
          case 'TLE':
          case 'MLE':
          case 'OLE':
          case 'CE':
            return "yellow-container"
          case 'PA':
          case 'JUG':
          case 'Pending':
          case 'AC':
          default:
            return "no-container"
        }
      },
      refreshBigFlag() {
        switch (this.result.flag) {
          case 'RE':
          case 'WA':
          case 'ISE':
            this.bigFlag.bgCss =
                    'width:100%;height:100px;background-color:rgb(253, 226, 226);border-radius: 4px';
            this.bigFlag.iconCls = 'el-icon-error';
            this.bigFlag.iconCss = 'margin:35px 5px 0px;color:red';
            break;
          case 'TLE':
          case 'MLE':
          case 'OLE':
          case 'CE':
            this.bigFlag.bgCss =
                    'width:100%;height:100px;background-color:rgb(250, 236, 216);border-radius: 4px';
            this.bigFlag.iconCls = 'el-icon-warning';
            this.bigFlag.iconCss = 'margin:35px 5px 0px;color:#E6A23C';
            break;
          case 'PA':
          case 'JUG':
          case 'Pending':
            this.bigFlag.bgCss =
                    'width:100%;height:100px;background-color:rgb(217, 236, 255);border-radius: 4px';
            this.bigFlag.iconCls = 'el-icon-s-help';
            this.bigFlag.iconCss = 'margin:35px 5px 0px;color:#409EFF';
            break;
          case 'AC':
            this.bigFlag.bgCss =
                    'width:100%;height:100px;background-color:rgb(225, 243, 216);border-radius: 4px';
            this.bigFlag.iconCls = 'el-icon-success';
            this.bigFlag.iconCss = 'margin:35px 5px 0px;color:rgb(103, 194, 58)';
            break;
          default:
            this.bigFlag.bgCss =
                    'width:100%;height:100px;background-color:rgb(253, 226, 226);border-radius: 4px';
            this.bigFlag.iconCls = 'el-icon-error';
            this.bigFlag.iconCss = 'margin:35px 5px 0px;color:red';
            break;
        }
        switch (this.result.flag) {
          case 'RE':
            this.bigFlag.flag = 'Runtime Error';
            break;
          case 'CE':
            this.bigFlag.flag = 'Compile Error';
            break;
          case 'WA':
            this.bigFlag.flag = 'Wrong Answer';
            break;
          case 'ISE':
            this.bigFlag.flag = 'Internal Server Error';
            break;
          case 'TLE':
            this.bigFlag.flag = 'Time Limit Exceeded';
            break;
          case 'MLE':
            this.bigFlag.flag = 'Memory Limit Exceeded';
            break;
          case 'OLE':
            this.bigFlag.flag = 'Output Limit Exceeded';
            break;
          case 'PA':
            this.bigFlag.flag = 'Partial Accepted';
            break;
          case 'JUG':
            this.bigFlag.flag = 'Judging';
            break;
          case 'Pending':
            this.bigFlag.flag = 'Pending';
            break;
          case 'AC':
            this.bigFlag.flag = 'Accepted';
            break;
          default:
            this.bigFlag.flag = 'Internal Server Error';
        }
        this.bigFlag.show = true;
      },
      log() {
        console.log(typeof this.option);
      },
      goBack() {
        this.$router.go(-1);
      },
      async getStatus() {
        if (this.result.flag === 'JUG' && this.waitTimes < 50) {
          const {
            data: res
          } = await this.$http.post('/contest/getStatus', {
            id: this.$route.query.id
          });
          if (res.error) {
            this.$message.error(res.error);
            return;
          }
          this.result = res.data;
          this.waitTimes += 1;
          if (
                  this.result.flag === 'JUG' &&
                  this.waitTimes < 50 &&
                  this.$route.path === '/contest/result'
          ) {
            setTimeout(this.getStatus, 500);
          } else {
            console.log(res.result);
            this.refreshBigFlag();
            await this.getStatusDetail();
          }
        }
      },
      async getStatusDetail() {
        const {
          data: res
        } = await this.$http.post('/contest/getStatusDetail', {
          id: this.$route.query.id
        });
        if (res.error) {
          this.$message.error(res.error);
          return;
        }
        this.resultDetail = res.data;
      },
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
    },
    components: {
      codemirror
    },
    filters: {
      formatFlags: function (value) {
        switch (value) {
          case 'RE':
            return 'Runtime Error';
          case 'CE':
            return 'Compile Error';
          case 'WA':
            return 'Wrong Answer';
          case 'ISE':
            return 'Internal Server Error';
          case 'TLE':
            return 'Time Limit Exceeded';
          case 'MLE':
            return 'Memory Limit Exceeded';
          case 'OLE':
            return 'Output Limit Exceeded';
          case 'PA':
            return 'Partial Accepted';
          case 'JUG':
            return 'Judging';
          case 'Pending':
            return 'Pending';
          case 'AC':
            return 'Accepted';
          default:
            return 'Internal Server Error'
        }
      },
      formatType: function (value) {
        switch (value) {
          case 'RE':
          case 'WA':
          case 'ISE':
            return 'danger';
          case 'TLE':
          case 'MLE':
          case 'OLE':
          case 'CE':
            return 'warning';
          case 'PA':
          case 'JUG':
          case 'Pending':
            return 'primary';
          case 'AC':
            return 'success';
          default:
            return 'danger'
        }
      }
    }
  };
</script>

<style scoped>
  .center-box {
    min-width: 600px;
    margin: 20px auto 0;
    width: 80%;
    background-color: rgb(244, 244, 245);
  }

  .center-box >>> .CodeMirror {
    height: auto;
  }

  .center-box >>> .CodeMirror-scroll {
    overflow: scroll !important;
    min-height: 100px;
    height: auto;
  }

  .center-box >>> .el-form-item__label {
    text-align: left;
    vertical-align: middle;
    float: left;
    font-size: 14px;
    min-width: 120px;
    color: #99a9bf;
    line-height: 40px;
    padding: 0 12px 0 0;
    box-sizing: border-box;
  }

  .red-container {
    padding: 8px 16px;
    background-color: #fff6f7;
    border-radius: 4px;
    border-left: 5px solid #fe6c6f;
    margin: 0 0 30px 0;
  }

  .yellow-container {
    padding: 8px 16px;
    background-color: rgb(253, 246, 236);
    border-radius: 4px;
    border-left: 5px solid #E6A23C;
    margin: 0 0 30px 0;

  }

  .no-container{
    display: none;
  }

  .el-col {
    text-align: center;
    display: flex;
    justify-content: center;
    align-items: center;
  }

  .demo-table-expand {
    font-size: 0;
  }

  .table-label {
    width: 90px !important;
    color: #515458 !important;
  }

  .demo-table-expand .el-form-item {
    margin-right: 0;
    margin-bottom: 0;
    width: 100%;
  }

  .slide-fade-enter-active {
    transition: all 0.8s ease;
  }

  .slide-fade-leave-active {
    transition: all 0.8s cubic-bezier(1, 0.5, 0.8, 1);
  }

  .slide-fade-enter,
  .slide-fade-leave-to

    /* .slide-fade-leave-active for below version 2.1.8 */
  {
    transform: translateY(40px);
    opacity: 0;
  }
</style>