<template>
    <div>
        <transition name="slide-fade">
            <div class="center-box" v-if="show">
                <div class="left-box">
                    <el-row class="problem-box">
                        <el-row>
                            <p style="font-size:22px;margin-left:2%">
                                {{ detail.title }}
                            </p>
                        </el-row>
                        <el-row style="width:94%;margin-left:3%">
                            <el-row v-if="detail.description">
                                <p class="title" style="margin-top:10px;">Description</p>
                                <p class="content" style="margin-left:3%" v-html="detail.description">
                                </p>
                            </el-row>
                            <el-row v-if="detail.inputDescription">
                                <p class="title">Input</p>
                                <p class="content" style="margin-left:3%" v-html="detail.inputDescription">
                                </p>
                            </el-row>
                            <el-row v-if="detail.outputDescription">
                                <p class="title">Output</p>
                                <p class="content" style="margin-left:3%" v-html="detail.outputDescription">
                                </p>
                            </el-row>
                            <el-row class="sample" v-bind:key="i" v-for="(sample, i) in detail.sample" :gutter="20">
                                <el-col :span="12">
                                    <el-row style="width:100%">
                                        <p class="sample-title">Sample Input {{ i + 1 }} <i
                                            class="el-icon-document-copy"
                                            @click="copyText(i)" style="cursor:pointer"
                                        ></i></p>
                                        <el-input type="textarea" :id="'sampleInput'+i" resize="none" readonly autosize
                                                  v-model="sample.input">
                                        </el-input>
                                    </el-row>
                                </el-col>
                                <el-col :span="12">
                                    <el-row style="width:100%">
                                        <p class="sample-title">Sample Output {{ i + 1 }}</p>

                                        <el-input type="textarea" resize="none" readonly autosize
                                                  v-model="sample.output">
                                        </el-input>
                                    </el-row>
                                </el-col>
                            </el-row>
                            <el-row v-if="detail.hint">
                                <p class="title" style="margin-top:10px">Hint</p>
                                <p class="content" style="margin-left:3%" v-html="detail.hint">
                                </p>
                            </el-row>
                            <el-row v-if="detail.source">
                                <p class="title" style="margin-top:10px">Source</p>
                                <p class="content" style="margin-left:3%" v-html="detail.source">
                                </p>
                            </el-row>
                            <el-row style="height:35px"></el-row>
                        </el-row>
                    </el-row>
                    <el-row style="text-align:left;" class="answer-box">
                        <el-row style=";margin: 15px auto 15px;width:95%">
                            <span style="font-size:15px;">Language:</span>
                            <el-select v-model="currentLanguage" placeholder="Language" style="margin-left:15px;"
                                       size="small" @change="langOptions.mode = getMIMEType(currentLanguage)">
                                <el-option v-for="item in detail.language" :key="item.id" :label="item.name"
                                           :value="item.name">
                                </el-option>
                            </el-select>
                            <el-tooltip effect="dark" content="Reset Code" placement="top-start">
                                <el-button @click="resetCode" plain style="margin-left:20px" size="small"
                                           class="el-icon-refresh"></el-button>
                            </el-tooltip>
                            <div style="float:right">
                                <span style="font-size:15px;">Theme:</span>
                                <el-select v-model="currentTheme" placeholder="Theme" style="margin-left:15px;"
                                           size="small" @change="setTheme">
                                    <el-option v-for="item in theme" :key="item" :label="item"
                                               :value="item">
                                    </el-option>
                                </el-select>
                            </div>

                        </el-row>
                        <el-row>
                            <codemirror v-model="code[currentLanguage]" :options="langOptions"
                                        style="width:95%;margin:0 auto;min-height: 400px!important"></codemirror>
                        </el-row>
                        <el-row style="width:95%;margin:15px auto 15px">
                            <div></div>
                            <span v-if="flag">Status:</span>
                            <el-button style="margin-left:15px" type="primary" plain v-if="flag === 'JUG'"
                                       @click="goStatusDetail">Judging
                            </el-button>
                            <el-button style="margin-left:15px" type="danger" plain v-if="flag === 'WA'"
                                       @click="goStatusDetail">
                                Wrong Answer
                            </el-button>
                            <el-button style="margin-left:15px" type="danger" plain v-if="flag === 'ISE'"
                                       @click="goStatusDetail">
                                Internal Server Error
                            </el-button>
                            <el-button style="margin-left:15px" type="danger" plain v-if="flag === 'RE'"
                                       @click="goStatusDetail">
                                Runtime Error
                            </el-button>
                            <el-button style="margin-left:15px" type="warning" plain v-if="flag === 'CE'"
                                       @click="goStatusDetail">
                                Compile Error
                            </el-button>
                            <el-button style="margin-left:15px" type="warning" plain v-if="flag === 'OLE'"
                                       @click="goStatusDetail">
                                Output Limit Exceeded
                            </el-button>
                            <el-button style="margin-left:15px" type="warning" plain v-if="flag === 'TLE'"
                                       @click="goStatusDetail">
                                Time Limit Exceeded
                            </el-button>
                            <el-button style="margin-left:15px" type="warning" plain v-if="flag === 'MLE'"
                                       @click="goStatusDetail">
                                Memory Limit Exceeded
                            </el-button>
                            <el-button style="margin-left:15px" type="success" plain v-if="flag === 'AC'"
                                       @click="goStatusDetail">
                                Accepted
                            </el-button>
                            <el-button style="margin-left:15px" type="primary" plain v-if="flag === 'PA'"
                                       @click="goStatusDetail">
                                Partial Accepted
                            </el-button>
                            <el-button style="margin-left:15px" type="primary" plain v-if="flag === 'Sending'"
                                       @click="goStatusDetail">Sending
                            </el-button>
                            <el-button type="primary" style="float:right;" @click="submit"
                                       :loading="isJudging"
                                       class="el-icon-s-promotion">&nbsp;&nbsp;Submit
                            </el-button>

                            <!-- <el-progress :text-inside="true" :stroke-width="20" :percentage="70"></el-progress> -->
                        </el-row>
                    </el-row>
                </div>
                <div class="right-box">
                    <el-row>
                        <el-button :plain="true" @click="goStatus"
                                   style="width:100%;height:50px;text-align:center"><span
                            class="el-icon-tickets" style="margin-left:-30%;font-size:15px">&nbsp;Submissions</span>
                        </el-button>
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
                                    <span style="float:right"> {{ detail.id }}</span>
                                </el-row>
                                <el-row style="margin-bottom: 14px;">
                                    <span style="float:left">CPU Time Limit</span>
                                    <span style="float:right">
                    {{ detail.limit[getIndexByLang(detail.limit, currentLanguage)].maxCpuTime + 'ms' }}</span>
                                </el-row>
                                <el-row style="margin-bottom: 14px;">
                                    <span style="float:left">Real Time Limit</span>
                                    <span style="float:right">
                    {{ detail.limit[getIndexByLang(detail.limit, currentLanguage)].maxRealTime + 'ms' }}</span>
                                </el-row>
                                <el-row style="margin-bottom: 14px;">
                                    <span style="float:left">Memory Limit</span>
                                    <span style="float:right">
                    {{ detail.limit[getIndexByLang(detail.limit, currentLanguage)].maxMemory + 'KB' }}</span>
                                </el-row>
                                <el-row style="margin-bottom: 14px;">
                                    <span style="float:left">Created By</span>
                                    <span style="float:right">
                    {{ detail.creatorName }}</span>
                                </el-row>
                                <el-row style="margin-bottom: 14px;">
                                    <span style="float:left">Level</span>
                                    <span style="float:right">{{
                                            detail.difficulty
                                        }}</span>
                                </el-row>
                                <el-row style="margin-bottom: 14px;">
                                    <span style="float:left">Tags</span>
                                    <div v-bind:key="i" v-for="(tag, i) in detail.tag">
                                        <el-row>
                                            <el-button style="float:right;margin-bottom:5px" type="primary" size="mini"
                                                       plain>{{ tag.name }}
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
                                           @click="DetailChartVisible = true">Details
                                </el-button>
                            </el-row>
                            <el-dialog title="Details" :visible.sync="DetailChartVisible" :append-to-body="true"
                                       width="580px">
                                <el-row style="height:360px">
                                    <ve-ring :data="detailedChartData" width="520px" height="440px"
                                             :settings="detailedChartSettings"
                                             :colors="detailedChartSettings.color" :tooltip-visible="false"
                                             style="float:left;margin-left:13px;height:450px;margin-top:-30px"></ve-ring>
                                    <!-- style="float:left;margin-left:13px;height:500px" -->
                                    <ve-pie :data="innerBriefChartData" width="165px" height="165px"
                                            :tooltip-visible="false"
                                            :settings="innerBriefChartSettings" :colors="innerBriefChartSettings.color"
                                            :legend-visible="false"
                                            style="float:right;margin-right:185px;margin-top:-260px"></ve-pie>
                                    <!-- style="float:right;margin-right:160px;margin-top:-344px" -->
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
                            <el-button type="primary" style="width:100%;height:40px;margin-top:20px;"
                                       class="el-icon-back"
                                       @click="goBack">&nbsp;Back
                            </el-button>
                        </el-row>
                    </el-row>
                </div>
            </div>
        </transition>
    </div>
</template>

<script>
// require component
import {codemirror} from 'vue-codemirror';

import 'codemirror/lib/codemirror.css';

import 'codemirror/theme/idea.css';
import 'codemirror/theme/darcula.css';

import 'codemirror/mode/clike/clike';
import 'codemirror/mode/go/go';
import 'codemirror/mode/python/python';
import 'codemirror/addon/scroll/annotatescrollbar.js'
import 'codemirror/addon/search/matchesonscrollbar.js'
import 'codemirror/addon/search/match-highlighter.js'
import 'codemirror/addon/search/jump-to-line.js'
import 'codemirror/addon/dialog/dialog.js'
import 'codemirror/addon/dialog/dialog.css'
import 'codemirror/addon/search/searchcursor.js'
import 'codemirror/addon/search/search.js'
import 'codemirror/addon/fold/foldgutter.css'
import 'codemirror/addon/fold/foldcode'
import 'codemirror/addon/fold/foldgutter'
import 'codemirror/addon/fold/brace-fold'
import 'codemirror/addon/fold/comment-fold'


import 'codemirror/addon/selection/active-line';
/* MIME Type
C: text/x-csrc
C++: text/x-c++src
Java: text/x-java

*/
import VePie from 'v-charts/lib/pie.common'
import VeRing from 'v-charts/lib/ring.common'

export default {
    components: {
        codemirror,
        VePie,
        VeRing,
    },
    data() {
        return {
            flag: "",
            statistic: {},
            isJudging: false,
            psid: -1,
            show: false,
            code: {
                C: "",
                Cpp: "",
                Java: "",
                Python: "",
                Go: "",
            },
            currentTheme: "",
            theme: ['idea', 'darcula'],
            waitTimes: 0,
            currentLanguage: '',
            DetailChartVisible: false,
            langOptions: {
                // codemirror options
                tabSize: 4,
                mode: '',
                autoRefresh: true,
                styleActiveLine: true,
                smartIndent: true,
                indentUnit: 4,
                theme: 'idea',
                lineNumbers: true,
                line: true,
                foldGutter: true,
                lineWrapping: true,
                gutters: ['CodeMirror-linenumbers', 'CodeMirror-foldgutter', 'CodeMirror-lint-markers'],
            },
            detail: {
                cpuTimeLimit: 0,
                createTime: '',
                creatorName: '',
                description: '',
                difficulty: '',
                hint: '',
                id: 0,
                inputDescription: '',
                outputDescription: '',
                ioMode: '',
                language: [],
                lastUpdateTime: '',
                memoryLimit: 0,
                realTimeLimit: 0,
                ref: '',
                ruleType: '',
                sample: [],
                tag: [],
                title: ''
            },
            detailedChartSettings: {
                label: {
                    show: true,
                    position: 'outside',
                    formatter: '{b}: {c}\n{d}%',
                    fontWeight: 'bold'
                },
                // labelLine: {
                //   show: true
                // },
                color: [
                    '#73d13d',
                    '#f5222d',
                    '#fa541c',
                    '#faad14',
                    '#722ed1',
                    '#13c2c2',
                    '#eb2f96'
                ],
                offsetY: 270,
                radius: [110, 160]
            },
            detailedChartData: {
                columns: ['Status', 'Statistic'],
                rows: [{
                    Status: 'AC',
                    Statistic: 0
                },
                    {
                        Status: 'WA',
                        Statistic: 0
                    },
                    {
                        Status: 'CE',
                        Statistic: 0
                    },
                    {
                        Status: 'RE',
                        Statistic: 0
                    },
                    {
                        Status: 'TLE',
                        Statistic: 0
                    },
                    {
                        Status: 'MLE',
                        Statistic: 0
                    },
                    {
                        Status: 'OLE',
                        Statistic: 0
                    }
                ]
            },
            innerBriefChartSettings: {
                label: {
                    show: true,
                    position: 'inner',
                    formatter: '{b}: {c}\n{d}%',
                    fontWeight: 'bold'
                },
                labelLine: {
                    show: true
                },
                color: ['#73d13d', '#f5222d'],
                radius: 70,
                offsetY: 80
            },
            innerBriefChartData: {
                columns: ['status', 'number'],
                rows: [{
                    status: 'AC',
                    number: 1393
                },
                    {
                        status: 'WA',
                        number: 3530
                    }
                ]
            },
            briefChartSettings: {
                label: {
                    show: true,
                    position: 'inner',
                    formatter: '{b}: {c}\n{d}%',
                    fontWeight: 'bold'
                },
                labelLine: {
                    show: false
                },
                color: ['#73d13d', '#f5222d'],
                radius: 80,
                offsetY: 120
            },
            briefChartData: {
                columns: ['status', 'number'],
                rows: [{
                    status: 'AC',
                    number: 0
                },
                    {
                        status: 'WA',
                        number: 0
                    }
                ]
            }
        };
    },
    mounted() {
        this.loadTheme()
        this.getDetail()
    },
    beforeCreate() {
        this.show = false;
        this.$bus.emit('changeHeader', '2');
    },
    methods: {
        async getDetail() {
            try {
                const {
                    data: res
                } = await this.$http.post('/practice/getDetail', {
                    id: Number(this.$route.query.id)
                });
                if (res.error) {
                    if (res.error === "sql: no rows in result set")
                        this.$message.error("the problem has benn hidden or deleted");
                    await this.$router.push("/practice")
                    return;
                }
                this.statistic = res.data.statistic;
                this.detail = res.data;
                this.currentLanguage = this.detail.language[0].name;
                this.langOptions.mode = this.getMIMEType(this.currentLanguage)
                this.refreshStatistic()
                this.setTemplate()
                this.show = true;
                const {
                    data: res1
                } = await this.$http.post(
                    '/practice/getCurrentStatus', {
                        id: Number(this.$route.query.id)
                    }
                );
                if (!res1.data) {
                    return
                }
                if (res1.error) {
                    this.$message.error(res1.error);
                    return;
                }
                let data = res1.data
                if (data.flag) {
                    this.flag = data.flag;
                }
                if (data.code) {
                    this.code[this.getLang(data.lid)] = data.code;
                }
                this.currentLanguage = this.getLang(data.lid)
                if (data.eid) {
                    this.psid = res1.data.eid;
                }
                this.langOptions.mode = this.getMIMEType(this.currentLanguage)
            } catch (err) {
                console.log(err);
            }
        },
        getMIMEType(lang) {
            switch (lang) {
                case "C":
                    return "text/x-csrc"
                case "Cpp":
                    return "text/x-c++src"
                case "Java":
                    return "text/x-java"
                case "Python":
                    return "python"
                case "Go":
                    return "go"
            }
        },
        getIndexByLang(obj, lang) {
            let lid = this.getLid(lang)
            if (obj) {
                for (let i = 0; i < obj.length; i++) {
                    if (obj[i].lid === lid) {
                        return i;
                    }
                }
            }
        },
        getLid(lang) {
            switch (lang) {
                case "C":
                    return 1
                case "Cpp":
                    return 2
                case "Java":
                    return 3
                case "Python":
                    return 4
                case "Go":
                    return 5
                default:
                    this.$message.error("no such language " + lang)
                    throw "no such language"
            }
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
        setTemplate() {
            let tmpl = this.detail.template
            if (tmpl) {
                for (let i = 0; i < tmpl.length; i++) {
                    this.code[this.getLang(tmpl[i].lid)] = tmpl[i].content;
                }
            }
        },
        loadTheme() {
            let storage = window.localStorage
            let theme = storage.getItem("theme")
            if (theme) {
                this.langOptions.theme = theme
                this.currentTheme = theme
            } else {
                this.langOptions.theme = this.theme[0]
                this.currentTheme = this.theme[0]
            }
        },
        setTheme() {
            this.langOptions.theme = this.currentTheme
            let storage = window.localStorage
            storage.setItem("theme", this.currentTheme)
        },
        resetCode() {
            let lang = this.currentLanguage
            let tmpl = this.detail.template
            this.code[lang] = tmpl[this.getIndexByLang(tmpl, lang)].content
        },
        copyToClipBoard(id) { //复制到剪切板
            if (document.execCommand) {
                let e = document.getElementById(id);
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
                status: 'AC',
                number: this.statistic.ac
            },
                {
                    status: 'WA',
                    number: this.statistic.total - this.statistic.ac
                }
            ];
            this.innerBriefChartData.rows = this.briefChartData.rows;
            this.detailedChartData.rows = [{
                Status: 'AC',
                Statistic: this.statistic.ac
            },
                {
                    Status: 'WA',
                    Statistic: this.statistic.wa
                },
                {
                    Status: 'CE',
                    Statistic: this.statistic.ce
                },
                {
                    Status: 'RE',
                    Statistic: this.statistic.re
                },
                {
                    Status: 'TLE',
                    Statistic: this.statistic.tle
                },
                {
                    Status: 'MLE',
                    Statistic: this.statistic.mle
                },
                {
                    Status: 'OLE',
                    Statistic: this.statistic.ole
                }
            ];
        },
        goBack() {
            this.$router.go(-1);
        },
        goStatusDetail() {
            this.$router.push({
                path: '/practice/result',
                query: {
                    id: this.psid
                }
            })
        },
        goStatus() {
            this.$router.push('/status/practice');
        },
        async getStatus() {
            if (this.flag === "JUG" && this.waitTimes < 50) {
                this.isJudging = true;
                const {
                    data: res
                } = await this.$http.post('/practice/getStatus', {
                    id: this.psid
                });
                if (res.error) {
                    this.$message.error(res.error);
                    return;
                }
                this.flag = res.data.flag
                this.waitTimes += 1
                if (this.flag === "JUG" && this.waitTimes < 50 && this.$route.path === "/practice/answer") {
                    setTimeout(this.getStatus, 500)
                } else {
                    this.isJudging = false;
                }
            }
        },
        async submit() {
            if (!this.code[this.currentLanguage]) {
                this.$message.error('Code can not be empty!');
                return;
            }
            try {
                this.isJudging = true;
                this.waitTimes = 0
                this.flag = 'Sending';
                const {
                    data: res
                } = await this.$http.post('/practice/submit', {
                    code: this.code[this.currentLanguage],
                    lid: this.getLid(this.currentLanguage),
                    pid: this.detail.id,
                });
                if (res.error) {
                    this.$message.error(res.error);
                    this.isJudging = false
                    this.flag = '';
                    return;
                }
                this.psid = res.data.eid;
                this.flag = res.data.flag;
                setTimeout(this.getStatus, 500)
            } catch (err) {
                console.log(err);
            }
        }
    },
};
</script>

<style scoped>
.center-box {
    min-width: 600px;
    margin: 20px auto 0 auto;
    width: 95%;
    background-color: rgb(244, 244, 245);
    border-radius: 10px;
    display: flex;
}

.center-box >>> .CodeMirror-scroll {
    overflow: scroll !important;
    min-height: 400px;
    height: auto;
}

.center-box >>> .CodeMirror {
    height: auto;
}

.sample {
    text-align: left;
    width: 100%;
}

.sample-title {
    color: #3091f2;
    font-size: 18px;
    font-weight: 400;
    float: left
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

.center-box >>> .CodeMirror-scroll {
    min-height: 400px !important;
}
</style>