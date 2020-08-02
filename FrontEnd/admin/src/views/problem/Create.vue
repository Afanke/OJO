<template>
    <div>
        <transition name="slide-fade">
            <div class="center-box" v-if="show">
                <el-page-header style="height:60px;line-height:60px;margin-left:20px" title="Back" @back="goBack"
                                content="Create Problem">
                </el-page-header>
                <el-row style="height:1px;float:top;border-top:1px solid rgb(233, 233, 235);">
                </el-row>
                <el-row style="margin-left:20px;margin-right:20px">
                    <el-row :gutter="20" style="margin-top:20px">
                        <el-col :span="6">
                            <span style="color:red">*</span>
                            <span>&nbsp;Display ID</span>
                            <el-input v-model="ref" placeholder="Display ID" style="margin-top:20px"></el-input>
                        </el-col>
                        <el-col :span="18">
                            <span style="color:red">*</span>
                            <span>&nbsp;Title</span>
                            <el-input v-model="title" placeholder="Title" style="margin-top:20px"></el-input>
                        </el-col>
                    </el-row>
                    <el-row style="margin-top:30px">
                        <span style="color:red">*</span>
                        <span>&nbsp;Description</span>
                        <editor style="margin-top:20px" v-model="description"></editor>
                    </el-row>
                    <el-row style="margin-top:30px">
                        <span style="color:red">*</span>
                        <span>&nbsp;Input Description</span>
                        <editor style="margin-top:20px" v-model="inputDescription"></editor>
                    </el-row>
                    <el-row style="margin-top:30px">
                        <span style="color:red">*</span>
                        <span>&nbsp;Output Description</span>
                        <editor style="margin-top:20px" v-model="outputDescription"></editor>
                    </el-row>

                    <el-row :gutter="20" style="margin-top:40px">
                        <el-col :span="2">
                            <span style="color:red">*</span>
                            <span>&nbsp;Visible</span>
                            <el-row style="margin-top:25px;margin-left:10px">
                                <el-switch v-model="visible" active-color="#13ce66" inactive-color="#ff4949">
                                </el-switch>
                            </el-row>
                        </el-col>
                        <el-col :span="5">
                            <span style="color:red">*</span>
                            <span>&nbsp;Difficulty</span>
                            <el-row class="small-element">
                                <el-select v-model="difficulty" size="small" placeholder="请选择" style="width:100px">
                                    <el-option v-for="(item,index1) in difficultyOptions" :key="index1+'index1'"
                                               :label="item.label"
                                               :value="item.value">
                                    </el-option>
                                </el-select>
                            </el-row>
                        </el-col>
                        <el-col :span="7" :offset="1">
                            <span style="color:#ffffff">*</span>
                            <span>&nbsp;Tags</span>
                            <el-row class="small-element">
                                <el-popover placement="bottom" width="400" trigger="click">
                                    <el-checkbox-group v-model="tags">
                                        <el-checkbox style="min-width:100px" :key="index2+'index2'" :label="item.name"
                                                     v-for="(item,index2) in allTags"></el-checkbox>
                                    </el-checkbox-group>
                                    <el-button size="small" slot="reference">Click To Select</el-button>
                                </el-popover>
                            </el-row>
                        </el-col>
                        <el-col :span="8">
                            <span style="color:red">*</span>
                            <span>&nbsp;Languages</span>
                            <el-row class="small-element">
                                <el-checkbox label="C" v-model="useLang.C"></el-checkbox>
                                <el-checkbox label="C++" v-model="useLang.Cpp"></el-checkbox>
                                <el-checkbox label="Java" v-model="useLang.Java"></el-checkbox>
                                <el-checkbox label="Python" v-model="useLang.Python"></el-checkbox>
                                <el-checkbox label="Go" v-model="useLang.Go"></el-checkbox>
                            </el-row>
                        </el-col>
                    </el-row>
                    <el-row v-for="(item,name,index) in limit" v-if="useLang[name]" :key="'limit'+index"
                            style="margin-top:40px;border:1px solid rgb(233, 233, 235);border-radius:5px;white-space: nowrap;text-overflow :ellipsis">
                        <el-row :gutter="10" style="margin-bottom:9px;border-bottom:1px solid rgb(233, 233, 235)">
                            <p style="padding-left: 25px;font-weight:bold; "> Resource limit of {{name}}</p>
                        </el-row>
                        <el-row :gutter="30"
                                style="font-size:15px;margin-bottom: 15px;margin-top:15px">
                            <el-col :span="4" offset="2">
                                <span style="color:red">*</span>
                                <span>&nbsp;CPU Time Limit (ms)</span>
                                <el-row class="small-element">
                                    <el-input-number size="small" v-model="item.maxCpuTime" controls-position="right"
                                                     :min="10"
                                                     :max="30000" :step="100">
                                    </el-input-number>
                                </el-row>
                            </el-col>
                            <el-col :span="4">
                                <span style="color:red">*</span>
                                <span>&nbsp;Real Time Limit (ms)</span>
                                <el-row class="small-element">
                                    <el-input-number size="small" v-model="item.maxRealTime" controls-position="right"
                                                     :min="10"
                                                     :max="60000" :step="100">
                                    </el-input-number>
                                </el-row>
                            </el-col>
                            <el-col :span="4">
                                <span style="color:red">*</span>
                                <span>&nbsp;Memory Limit (KB)</span>
                                <el-row class="small-element">
                                    <el-input-number size="small" v-model="item.maxMemory" controls-position="right"
                                                     :min="1024" :max="2097152" :step="1024">
                                    </el-input-number>
                                </el-row>
                            </el-col>
                            <el-col :span="4">
                                <span style="color:red">*</span>
                                <span>&nbsp;Compile Multiple</span>
                                <el-row class="small-element">
                                    <el-input-number size="small" v-model="item.compMp" controls-position="right"
                                                     :min="1" :max="100" :step="1">
                                    </el-input-number>
                                </el-row>
                            </el-col>
                            <el-col :span="4">
                                <span style="color:red">*</span>
                                <span>&nbsp;Special Judge Multiple</span>
                                <el-row class="small-element">
                                    <el-input-number size="small" v-model="item.SPJMp" controls-position="right"
                                                     :min="1" :max="100" :step="1">
                                    </el-input-number>
                                </el-row>
                            </el-col>
                        </el-row>
                    </el-row>
                    <el-row style="border:1px solid rgb(233, 233, 235);border-radius:5px;margin-top:30px"
                            v-for="(item,index3) in sample" :key="index3+'index3'">
                        <el-row style="height:50px;line-height:50px">
                            <span style="margin-left:10px;font-weight:bold">Sample {{index3+1}}</span>
                            <el-button class="el-icon-delete" type="warning"
                                       style="margin-right:40px;float:right;margin-top:9px" size="small"
                                       @click="deleteSample(index3)">&nbsp;Delete
                            </el-button>
                        </el-row>
                        <el-row style="background-color:#FAFDFF">
                            <el-row :gutter="20" style="margin:0 6px 20px 6px">
                                <el-col :span="12">
                                    <el-row style="height:50px;line-height:50px">
                                        <span style="color:red">*</span>
                                        <span>&nbsp;Input Sample</span>
                                    </el-row>
                                    <el-row>
                                        <el-input resize="none" :autosize="{ minRows: 4}" type="textarea" :rows="2"
                                                  placeholder="Input Sample" v-model="item.input">
                                        </el-input>
                                    </el-row>
                                </el-col>
                                <el-col :span="12">
                                    <el-row style="height:50px;line-height:50px">
                                        <span style="color:red">*</span>
                                        <span>&nbsp;Output Sample</span>
                                    </el-row>
                                    <el-row>
                                        <el-input resize="none" :autosize="{ minRows: 4}" type="textarea" :rows="2"
                                                  placeholder="Output Sample" v-model="item.output">
                                        </el-input>
                                    </el-row>
                                </el-col>
                            </el-row>
                        </el-row>
                    </el-row>
                    <el-row class="add-button">
                        <div @click="addSample">
                            <i class="el-icon-plus"></i>
                            <span>&nbsp;Add Sample</span>
                        </div>
                    </el-row>
                    <el-row style="border:1px solid rgb(233, 233, 235);border-radius:5px;margin-top:30px"
                            v-for="(item,index4) in problemCase" :key="index4+'index4'">
                        <el-row style="height:50px;line-height:50px">
                            <span style="margin-left:10px;font-weight:bold">Test Case {{index4+1}}</span>
                            <el-button class="el-icon-delete" type="warning"
                                       style="margin-right:40px;float:right;margin-top:9px" size="small"
                                       @click="deleteProblemCase(index4)">&nbsp;Delete
                            </el-button>
                            <el-input-number :step="10" size="small" :min="0"
                                             style="float:right;margin-right:40px;margin-top:9px"
                                             v-model="item.score"></el-input-number>
                            <span style="float:right;margin-right:10px;">Score:</span>
                        </el-row>
                        <el-row style="background-color:#FAFDFF">
                            <el-row :gutter="20" style="margin:0 6px 20px 6px">
                                <el-col :span="12">
                                    <el-row style="height:50px;line-height:50px">
                                        <span style="color:red">*</span>
                                        <span>&nbsp;Input</span>
                                    </el-row>
                                    <el-row>
                                        <el-input resize="none" :autosize="{ minRows: 4}" type="textarea" :rows="2"
                                                  placeholder="Input"
                                                  v-model="item.input">
                                        </el-input>
                                    </el-row>
                                </el-col>
                                <el-col :span="12">
                                    <el-row style="height:50px;line-height:50px">
                                        <span style="color:red">*</span>
                                        <span>&nbsp;Output</span>
                                    </el-row>
                                    <el-row>
                                        <el-input resize="none" :autosize="{ minRows: 4}" type="textarea" :rows="2"
                                                  placeholder="Output"
                                                  v-model="item.output">
                                        </el-input>
                                    </el-row>
                                </el-col>
                            </el-row>
                        </el-row>
                    </el-row>
                    <el-row class="add-button">
                        <div @click="addProblemCase">
                            <i class="el-icon-plus"></i>
                            <span>&nbsp;Add Test Case</span>
                        </div>
                    </el-row>
                    <el-row style="margin-top:30px">
                        <span>&nbsp;Template</span>
                        <el-row class="template-checkbox">
                            <el-checkbox v-model="useTmpl.C" :disabled="!useLang.C">C</el-checkbox>
                            <el-button class="el-icon-refresh" v-if="useTmpl.C" type="primary"
                                       style="float:right;margin-right:20px"
                                       size="small"
                                       @click="resetTmpl('C')">
                                &nbsp;Reset
                            </el-button>
                            <codemirror v-if="useTmpl.C" v-model="tmpl.C" :options="COptions"
                                        style="width:100%;margin: 20px auto 0;border:1px solid rgb(233, 233, 235)">
                            </codemirror>
                        </el-row>
                        <el-row class="template-checkbox">
                            <el-checkbox v-model="useTmpl.Cpp" :disabled="!useLang.Cpp">C++</el-checkbox>
                            <el-button class="el-icon-refresh" v-if="useTmpl.Cpp" type="primary"
                                       style="float:right;margin-right:20px"
                                       size="small" @click="resetTmpl('Cpp')">
                                &nbsp;Reset
                            </el-button>
                            <codemirror v-if="useTmpl.Cpp" v-model="tmpl.Cpp" :options="CppOptions"
                                        style="width:100%;margin: 20px auto 0;border:1px solid rgb(233, 233, 235)">
                            </codemirror>
                        </el-row>
                        <el-row class="template-checkbox">
                            <el-checkbox v-model="useTmpl.Java" :disabled="!useLang.Java">Java</el-checkbox>
                            <el-button class="el-icon-refresh" v-if="useTmpl.Java" type="primary"
                                       style="float:right;margin-right:20px"
                                       size="small" @click="resetTmpl('Java')">
                                &nbsp;Reset
                            </el-button>
                            <codemirror v-if="useTmpl.Java" v-model="tmpl.Java" :options="JavaOptions"
                                        style="width:100%;margin: 20px auto 0;border:1px solid rgb(233, 233, 235)">
                            </codemirror>
                        </el-row>
                        <el-row class="template-checkbox">
                            <el-checkbox v-model="useTmpl.Python" :disabled="!useLang.Python">Python</el-checkbox>
                            <el-button class="el-icon-refresh" v-if="useTmpl.Python" type="primary"
                                       style="float:right;margin-right:20px"
                                       size="small"
                                       @click="resetTmpl('Python')">
                                &nbsp;Reset
                            </el-button>
                            <codemirror v-if="useTmpl.Python" v-model="tmpl.Python" :options="PythonOptions"
                                        style="width:100%;margin: 20px auto 0;border:1px solid rgb(233, 233, 235)">
                            </codemirror>
                        </el-row>
                        <el-row class="template-checkbox">
                            <el-checkbox v-model="useTmpl.Go" :disabled="!useLang.Go">Go</el-checkbox>
                            <el-button class="el-icon-refresh" v-if="useTmpl.Go" type="primary"
                                       style="float:right;margin-right:20px"
                                       size="small" @click="resetTmpl('Go')">
                                &nbsp;Reset
                            </el-button>
                            <codemirror v-if="useTmpl.Go" v-model="tmpl.Go" :options="GoOptions"
                                        style="width:100%;margin: 20px auto 0;border:1px solid rgb(233, 233, 235)">
                            </codemirror>
                        </el-row>
                    </el-row>
                    <el-row style="margin-top:30px">
                        <span>&nbsp;Special Judge</span>
                        <el-row style="margin-top:20px">
                            <el-checkbox v-model="useSPJ">Use Special Judge</el-checkbox>
                        </el-row>
                        <el-row class="spj-checkbox" v-if="useSPJ">
                            <el-row>
                                <span style="margin-top:17px;float:left;margin-left:10px;font-size:14px">Lang:</span>
                                <el-radio-group @change="changeLangOptions1" v-model="SPJLang"
                                                style="margin-top:19px;float:left;margin-left:10px">
                                    <el-radio label="C">C</el-radio>
                                    <el-radio label="Cpp">C++</el-radio>
                                    <el-radio label="Java">Java</el-radio>
                                    <el-radio label="Python">Python</el-radio>
                                    <el-radio label="Go">Go</el-radio>
                                </el-radio-group>
                                <el-button class="el-icon-refresh" style="float:right;margin-top:10px;margin-right:20px"
                                           type="primary"
                                           size="small"
                                           @click="resetSPJ">
                                    &nbsp;Reset
                                </el-button>
                            </el-row>
                            <codemirror v-model="SPJCode" :options="LangOptions1"
                                        style="width:100%;margin:0 auto;margin-top:10px;border:1px solid rgb(233, 233, 235)">
                            </codemirror>
                        </el-row>
                    </el-row>
                    <el-row style="margin-top:30px">
                        <span>&nbsp;Local Test</span>
                        <el-row class="local-test-checkbox">
                            <el-row>
                                <span style="margin-top:17px;float:left;margin-left:10px;font-size:14px">Lang:</span>
                                <el-radio-group @change="changeLangOptions2" v-model="ltLang"
                                                style="margin-top:19px;float:left;margin-left:10px">
                                    <el-radio label="C" :disabled="!useLang.C">C</el-radio>
                                    <el-radio label="Cpp" :disabled="!useLang.Cpp">C++</el-radio>
                                    <el-radio label="Java" :disabled="!useLang.Java">Java</el-radio>
                                    <el-radio label="Python" :disabled="!useLang.Python">Python</el-radio>
                                    <el-radio label="Go" :disabled="!useLang.Go">Go</el-radio>
                                </el-radio-group>
                                <el-button class="el-icon-position"
                                           style="float:right;margin-top:10px;margin-right:20px" type="primary"
                                           size="small"
                                           @click="localTest">&nbsp;Test
                                </el-button>
                                <el-button v-if="ltRes.flag==='AC'" type="success" plain size="small"
                                           class="local-test-result"
                                           @click="dialogVisible=true">Accepted
                                </el-button>
                                <el-button v-if="ltRes.flag==='Judging'" type="primary" plain size="small"
                                           class="local-test-result">
                                    Judging
                                </el-button>
                                <el-button v-if="ltRes.flag==='WA'" type="danger" plain size="small"
                                           class="local-test-result"
                                           @click="dialogVisible=true">Wrong Answer
                                </el-button>
                                <el-button v-if="ltRes.flag==='PA'" type="primary" plain size="small"
                                           class="local-test-result"
                                           @click="dialogVisible=true">Partcial Accepted
                                </el-button>
                                <el-button v-if="ltRes.flag==='TLE'" type="warning" plain size="small"
                                           class="local-test-result"
                                           @click="dialogVisible=true">Time Limit Exceeded
                                </el-button>
                                <el-button v-if="ltRes.flag==='MLE'" type="warning" plain size="small"
                                           class="local-test-result"
                                           @click="dialogVisible=true">Memory Limit Exceeded
                                </el-button>
                                <el-button v-if="ltRes.flag==='OLE'" type="warning" plain size="small"
                                           class="local-test-result"
                                           @click="dialogVisible=true">Output Limit Exceeded
                                </el-button>
                                <el-button v-if="ltRes.flag==='ISE'" type="danger" plain size="small"
                                           class="local-test-result"
                                           @click="dialogVisible=true">Internal Server Error
                                </el-button>
                                <el-button v-if="ltRes.flag==='RE'" type="danger" plain size="small"
                                           class="local-test-result"
                                           @click="dialogVisible=true">Runtime Error
                                </el-button>
                                <el-button v-if="ltRes.flag==='CE'" type="warning" plain size="small"
                                           class="local-test-result"
                                           @click="dialogVisible=true">Compile Error
                                </el-button>
                            </el-row>
                            <codemirror v-model="ltCode" :options="LangOptions2"
                                        style="width:100%;margin:0 auto;margin-top:10px;border-top:1px solid rgb(233, 233, 235)">
                            </codemirror>
                        </el-row>
                    </el-row>
                    <el-row style="margin-top:30px">
                        <span>&nbsp;Hint</span>
                        <editor style="margin-top:20px" v-model="hint"></editor>
                    </el-row>
                    <el-row style="margin-top:30px">
                        <span>&nbsp;Source</span>
                        <editor style="margin-top:20px" v-model="source"></editor>
                    </el-row>
                    <el-row style="margin-top:20px;margin-bottom:20px;">
                        <div style="text-align: center;">
                            <el-button type="primary" style="width:200px" @click="save">Save</el-button>
                        </div>
                    </el-row>
                </el-row>
            </div>
        </transition>
        <el-dialog title="Test Case Result" :visible.sync="dialogVisible" class="dialog" width="600px">
            <div class="page-container" v-if="ltRes.errorMsg">
                <p v-html="ltRes.errorMsg.replaceAll('\n','<br>')"></p>
            </div>
            <el-table :data="ltRes.testCase" style="margin-top:-25px">
                <el-table-column type="expand">
                    <template slot-scope="props">
                        <el-form label-position="left" inline class="demo-table-expand">
                            <el-form-item label="Flag">
                                <span>{{ props.row.flag }}</span>
                            </el-form-item>
                            <el-form-item label="Score">
                                <span>{{ props.row.score }}</span>
                            </el-form-item>
                            <el-form-item label="CpuTime">
                                <span>{{ props.row.actualCpuTime }}&nbsp;ms</span>
                            </el-form-item>
                            <el-form-item label="RealTime">
                                <span>{{ props.row.actualRealTime }}&nbsp;ms</span>
                            </el-form-item>
                            <el-form-item label="RealMemory">
                                <span>{{ props.row.realMemory }}&nbsp;KB</span>
                            </el-form-item>
                            <el-form-item label="Input">
                                <span v-html="props.row.input.replaceAll('\n','<br>')"></span>
                            </el-form-item>
                            <el-form-item label="ExpectedOutput">
                                <span v-html="props.row.expectedOutput.replaceAll('\n','<br>')"></span>
                            </el-form-item>
                            <el-form-item label="RealOutput">
                                <span v-html="props.row.realOutput.replaceAll('\n','<br>')"></span>
                            </el-form-item>
                            <el-form-item label="ErrorOutput">
                                <span v-html="props.row.errorOutput.replaceAll('\n','<br>')"></span>
                            </el-form-item>
                            <el-form-item label="SPJOutput">
                                <span v-html="props.row.SPJOutput.replaceAll('\n','<br>')"></span>
                            </el-form-item>
                            <el-form-item label="SPJErrorOutput">
                                <span v-html="props.row.SPJErrorOutput.replaceAll('\n','<br>')" ></span>
                            </el-form-item>
                        </el-form>
                    </template>
                </el-table-column>
                <el-table-column label="Index" width="200" type="index" align="center"></el-table-column>
                <el-table-column label="Result" min-width="200" align="center">
                    <template slot-scope="scope">
                        <el-button v-if="scope.row.flag==='AC'" type="success" size="small">Accepted
                        </el-button>
                        <el-button v-if="scope.row.flag==='WA'" type="danger" size="small">Wrong Answer
                        </el-button>
                        <el-button v-if="scope.row.flag==='PA'" type="primary" size="small">Partcial
                            Accepted
                        </el-button>
                        <el-button v-if="scope.row.flag==='TLE'" type="warning" size="small">Time Limit
                            Exceeded
                        </el-button>
                        <el-button v-if="scope.row.flag==='MLE'" type="warning" size="small">Memory Limit
                            Exceeded
                        </el-button>
                        <el-button v-if="scope.row.flag==='OLE'" type="warning" size="small">Output Limit
                            Exceeded
                        </el-button>
                        <el-button v-if="scope.row.flag==='ISE'" type="danger" size="small">Internal
                            Server Error
                        </el-button>
                        <el-button v-if="scope.row.flag==='RE'" type="danger" size="small">Runtime Error
                        </el-button>
                        <el-button v-if="scope.row.flag==='CE'" type="warning" size="small">Compile Error
                        </el-button>
                    </template>
                </el-table-column>
            </el-table>
        </el-dialog>
    </div>
</template>
<script>
    import Editor from '@/components/Editor'
    // require component
    import {codemirror} from 'vue-codemirror';
    // require styles
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

    export default {
        data() {
            return {
                dialogVisible: false,
                useLang: {
                    C: false,
                    Cpp: false,
                    Java: false,
                    Python: false,
                    Go: false,
                },
                limit: {
                    C: {
                        maxCpuTime: 1000,
                        maxRealTime: 1000,
                        maxMemory: 30720,
                        compMp: 2,
                        SPJMp: 2
                    },
                    Cpp: {
                        maxCpuTime: 1000,
                        maxRealTime: 1000,
                        maxMemory: 30720,
                        compMp: 2,
                        SPJMp: 2
                    },
                    Java: {
                        maxCpuTime: 1000,
                        maxRealTime: 1000,
                        maxMemory: 61440,
                        compMp: 2,
                        SPJMp: 2
                    },
                    Python: {
                        maxCpuTime: 1000,
                        maxRealTime: 1000,
                        maxMemory: 30720,
                        compMp: 2,
                        SPJMp: 2
                    },
                    Go: {
                        maxCpuTime: 1000,
                        maxRealTime: 1000,
                        maxMemory: 30720,
                        compMp: 2,
                        SPJMp: 2
                    },
                },
                useTmpl: {
                    C: false,
                    Cpp: false,
                    Java: false,
                    Python: false,
                    Go: false,
                },
                ltCode: "",
                ltLang: "",
                useSPJ: false,
                SPJCode: "",
                SPJLang: "",
                tmpl: {
                    C: "",
                    Cpp: "",
                    Java: "",
                    Python: "",
                    Go: "",
                },
                SPJExample: {
                    C: `#include <stdio.h>
#include <string.h>

char s1[100]={0};
char s2[100]={0};
char s3[100]={0};

int main(int argc, char *argv[]){
    FILE* f1=fopen(argv[1],"r");
    FILE* f2=fopen(argv[2],"r");
    FILE* f3=fopen(argv[3],"r");
    fgets(s1,100,f1); // Test Case Input
    fgets(s2,100,f2); // Test Case Expected Output
    fgets(s3,100,f3); // User Output
    if (strcmp(s2,s3)==0){
        printf("AC"); // Test Case Expected Output
    }else{
        printf("WA"); // You can print anything except "AC" to mark the answer "WA"
        // or do nothing...
    }
    fclose(f1);
    fclose(f2);
    fclose(f3);
}`,
                    Cpp: `#include <string>
#include <fstream>
#include <sstream>
#include <iostream>
#include <stdlib.h>
using namespace std;

string readFileString(char * filename)
{
    ifstream ifile(filename);
    ostringstream buf;
    char ch;
    while(buf&&ifile.get(ch))
    buf.put(ch);
    return buf.str();
}

int main(int argc, char *argv[]){
    string s1,s2,s3;
    s1=readFileString(argv[1]); // Test Case Input
    s2=readFileString(argv[2]); // Test Case Expected Output
    s3=readFileString(argv[3]); // User Output
    if (s2==s3){
        cout<<"AC"; // If the answer is right, just do this
    }else{
        cout<<"WA"; // You can print anything except "AC" to mark the answer "WA"
        // or do nothing...
    }
}`,
                    Java: `import java.io.*;

class SPJTest{

    public static void main(String args[]){
        String s1 = readToString(args[0]); // Test Case Input
        String s2 = readToString(args[1]); // Test Case Expected Output
        String s3 = readToString(args[2]); // User Output
        if (s2.equals(s3)) {
        System.out.printf("AC"); // If the answer is right, just do this
        } else {
        System.out.printf("WA"); // You can print anything except "AC" to mark the answer "WA"
        // or do nothing...
        }
    }

    public static String readToString(String fileName) {
        String encoding = "UTF-8";
        File file = new File(fileName);
        Long filelength = file.length();
        byte[] filecontent = new byte[filelength.intValue()];
        try {
            FileInputStream in = new FileInputStream(file);
            in.read(filecontent);
            in.close();
        } catch (FileNotFoundException e) {
            e.printStackTrace();
        } catch (IOException e) {
            e.printStackTrace();
        }
        try {
            return new String(filecontent, encoding);
        } catch (UnsupportedEncodingException e) {
            System.err.println("The OS does not support " + encoding);
            e.printStackTrace();
            return null;
        }
    }
}`,
                    Python: `import sys

if __name__=='__main__':
    s1=open(sys.argv[1],'r').read() # Test Case Input
    s2=open(sys.argv[2],'r').read() # Test Case Expected Output
    s3=open(sys.argv[3],'r').read() # User Output
    if s2==s3:
        print("AC",end="")  # If the answer is right, just do this
    else:
        print("WA",end="")  # You can print anything except "AC" to mark the answer "WA"
        # or do nothing...`,
                    Go: `package main

import(
    "io/ioutil"
    "os"
    "fmt"
)

func main(){
    // b1, _ := ioutil.ReadFile(os.Args[1])
    b2, _ := ioutil.ReadFile(os.Args[2])
    b3, _ := ioutil.ReadFile(os.Args[3])
    // s1:=string(b1)  // Test Case Input
    s2:=string(b2)  // Test Case Expected Output
    s3:=string(b3)  // User Output
    if s2==s3{
        fmt.Printf("AC")  // If the answer is right, just do this
        return
    }else{
        fmt.Printf("WA") // You can print anything except "AC" to mark the answer "WA"
        // or do nothing...
    }

}`,
                },
                tmplExample: {
                    C: `// PREPEND BEGIN
#include <stdio.h>

// PREPEND END

// TEMPLATE BEGIN
int add(int a,int b){
  // Please fill this blank
  return ___________;
}

// TEMPLATE END

// APPEND BEGIN
int main(){
    int a,b;
    scanf("%d %d",&a,&b);
    printf("%d",add(a,b));
}

// APPEND END`,
                    Cpp: `// PREPEND BEGIN
#include <iostream>
using namespace std;

// PREPEND END

// TEMPLATE BEGIN
int add(int a,int b){
  // Please fill this blank
  return ___________;
}

// TEMPLATE END

// APPEND BEGIN
int main(){
    int a,b;
    cin>>a;
    cin>>b;
    cout<<add(a,b);
}

// APPEND END
`,
                    Java: `
// PREPEND BEGIN
import java.util.Scanner;
class Test{

// PREPEND END

// TEMPLATE BEGIN
    public static int add(int a,int b){
        // Please fill this blank
        return ___________;
    }
// TEMPLATE END

// APPEND BEGIN
    public static void main(String args[]){
        Scanner sc = new Scanner(System.in);
        int a = sc.nextInt();
        int b = sc.nextInt();
        System.out.printf("%d",add(a,b));
    }
}
// APPEND END`,
                    Python: `# PREPEND BEGIN
import time

# PREPEND END

# TEMPLATE BEGIN
def my_add(a,b):
    # Please fill this blank
    return ___________

# TEMPLATE END

# APPEND BEGIN
if __name__=="__main__":
    print(my_add(int(input()),int(input())))

# APPEND END
`,
                    Go: `// PREPEND BEGIN
package main

import (
    "fmt"
)

// PREPEND END

// TEMPLATE BEGIN
func add(a,b int){
  // Please fill this blank
  return ___________
}

// TEMPLATE END

// APPEND BEGIN
func main(){
    var a,b int
    fmt.Scanf("%d %d",&a,&b)
    fmt.Printf("%d",add(a,b))
}

// APPEND END
                    `,
                },
                show: false,
                title: '',
                ref: '',
                allTags: [],
                textarea: '',
                realTags: [],
                realLanguages: [],
                tags: [],
                description: '',
                inputDescription: '',
                outputDescription: '',
                hint: '',
                source: '',
                ltRes: {
                    flag: "",
                    errorMsg: "",
                },
                visible: false,
                difficulty: 'Normal',
                COptions: {
                    // codemirror options
                    tabSize: 4,
                    mode: 'text/x-java',
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
                CppOptions: {
                    // codemirror options
                    tabSize: 4,
                    mode: 'text/x-csrc',
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
                JavaOptions: {
                    // codemirror options
                    tabSize: 4,
                    mode: 'text/x-csrc',
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
                PythonOptions: {
                    // codemirror options
                    tabSize: 4,
                    mode: 'python',
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
                GoOptions: {
                    // codemirror options
                    tabSize: 4,
                    mode: 'go',
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
                LangOptions1: {
                    // codemirror options
                    tabSize: 4,
                    mode: 'text/x-csrc',
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
                LangOptions2: {
                    // codemirror options
                    tabSize: 4,
                    mode: 'text/x-csrc',
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
                sample: [{
                    input: '',
                    output: '',
                }],
                problemCase: [{
                    input: '',
                    output: '',
                    score: 10
                }],
                difficultyOptions: [
                    {
                        value: 'Casual',
                        label: 'Casual'
                    }, {
                        value: 'Eazy',
                        label: 'Eazy'
                    }, {
                        value: 'Normal',
                        label: 'Normal'
                    }, {
                        value: 'Hard',
                        label: 'Hard'
                    }, {
                        value: 'Crazy',
                        label: 'Crazy'
                    }],
            }
        },
        created() {
            this.$bus.emit("changeHeader", "3-1")
            this.show = false
        },
        mounted() {
            this.show = true
            this.getAllTags()
        },
        methods: {
            resetTmpl(lang) {
                this.tmpl[lang] = this.tmplExample[lang]
            },
            resetSPJ() {
                this.SPJCode = this.SPJExample[this.SPJLang]
            },
            changeLangOptions1(lang) {
                this.changeLangOptions(lang, this.LangOptions1)
                this.SPJCode = this.SPJExample[lang]
            },
            changeLangOptions2(lang) {
                this.changeLangOptions(lang, this.LangOptions2)

            },
            changeLangOptions(lang, options) {
                switch (lang) {
                    case "C":
                        options.mode = "text/x-csrc"
                        break
                    case "Cpp":
                        options.mode = "text/x-c++src"
                        break
                    case "Java":
                        options.mode = "text/x-java"
                        break
                    case "Python":
                        options.mode = "python"
                        break
                    case "Go":
                        options.mode = "go"
                        break
                }
            },
            processTmpl(tmpl, lang) {
                let sign = ""
                let err = lang + " template is not available, please reset it"
                let obj = {
                    prepend: "",
                    content: "",
                    append: ""
                }
                switch (lang) {
                    case "C":
                    case "Cpp":
                    case "Java":
                    case "Go":
                        sign = "//"
                        break
                    case "Python":
                        sign = "#"
                        break
                    default:
                        this.$message.error("Error 1: " + err)
                        return null
                }
                let s1 = tmpl.split(sign + ` PREPEND BEGIN\n`)
                if (s1.length !== 2) {
                    this.$message.error("Error 2: " + err)
                    return null
                }
                let s2 = s1[1].split("\n" + sign + ` PREPEND END\n`)
                if (s2.length !== 2) {
                    this.$message.error("Error 3: " + err)
                    return null
                }
                obj.prepend = s2[0]
                if (obj.prepend.charAt(obj.prepend.length - 1) !== "\n") {
                    obj.prepend += "\n"
                }
                let s3 = s2[1].split(sign + ` TEMPLATE BEGIN\n`)
                if (s3.length !== 2) {
                    this.$message.error("Error 4: " + err)
                    return null
                }
                let s4 = s3[1].split("\n" + sign + ` TEMPLATE END\n`)
                if (s4.length !== 2) {
                    this.$message.error("Error 5: " + err)
                    return null
                }
                obj.content = s4[0]
                if (obj.content.charAt(obj.content.length - 1) !== "\n") {
                    obj.content += "\n"
                }
                let s5 = s4[1].split(sign + ` APPEND BEGIN\n`)
                if (s5.length !== 2) {
                    this.$message.error("Error 6: " + err)
                    return null
                }
                let s6 = s5[1].split("\n" + sign + ` APPEND END`)
                if (s6.length !== 2) {
                    this.$message.error("Error 7: " + err)
                    return null
                }
                obj.append = s6[0]
                if (obj.append.charAt(0) !== "\n") {
                    obj.append = "\n" + obj.append
                }
                if (obj.append.charAt(obj.append.length - 1) !== "\n") {
                    obj.append += "\n"
                }
                console.log(obj)
                return obj
            },
            async localTest() {
                let code = this.ltCode
                let lang = this.ltLang
                if (lang === "") {
                    this.$message.error("please select a local test language")
                    return
                }
                if (this.useTmpl[lang]) {
                    let r = this.processTmpl(this.tmpl[lang], lang)
                    if (r === null) {
                        return
                    }
                    code = r.prepend + code + r.append
                }
                let obj = {
                    maxMemory: this.limit[lang].maxMemory,
                    maxRealTime: this.limit[lang].maxRealTime,
                    maxCpuTime: this.limit[lang].maxCpuTime,
                    code: code,
                    lid: this.getLid(lang),
                    useSPJ: this.useSPJ,
                    SPJCode: this.SPJCode,
                    SPJLid: 1,
                    testCase: [],
                    compMp: this.limit[lang].compMp,
                    SPJMp: this.limit[lang].SPJMp
                }
                if (this.useSPJ) {
                    obj.SPJLid = this.getLid(this.SPJLang)
                }
                for (let i = 0; i < this.problemCase.length; i++) {
                    obj.testCase.push({
                        input: this.problemCase[i].input,
                        expectedOutput: this.problemCase[i].output,
                        score: this.problemCase[i].score,
                        id: i
                    })
                }
                try {
                    this.ltRes = {
                        flag: "Judging"
                    }
                    const {data: res} = await this.$http.post('/admin/problem/localTest', obj);
                    if (res.error) {
                        this.$message.error(res.error)
                        return
                    }
                    this.ltRes = res.data
                } catch (err) {
                    console.log(err);
                    alert(err)
                }
            },
            check() {
                if (this.title === "") {
                    this.$message.error("title is required")
                    return false
                }
                if (this.ref === "") {
                    this.$message.error("display id is required")
                    return false
                }
                if (this.description === "") {
                    this.$message.error("description is required")
                    return false
                }
                if (this.inputDescription === "") {
                    this.$message.error("input description is required")
                    return false
                }
                if (this.outputDescription === "") {
                    this.$message.error("output description is required")
                    return false
                }
                if (this.difficulty === "") {
                    this.$message.error("difficulty is required")
                    return false
                }
                if (!(this.useC || this.useCpp || this.useJava || this.usePython)) {
                    this.$message.error("language is required")
                    return false
                }
                for (let i = 0; i < this.sample.length; i++) {
                    if (this.sample[i].output === "") {
                        this.$message.error("sample " + (i + 1) + " output can't be empty")
                        return false
                    }
                }
                for (let i = 0; i < this.problemCase.length; i++) {
                    if (this.problemCase[i].output === "") {
                        this.$message.error("test case " + (i + 1) + " output can't be empty")
                        return false
                    }
                    if (this.problemCase[i].score === 0) {
                        this.$message.error("test case " + (i + 1) + " score can't be 0")
                        return false
                    }
                }
                return true
            },
            goBack() {
                this.$router.go(-1)
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
            getRealTags() {
                this.realTags = []
                for (let i = 0; i < this.tags.length; i++) {
                    for (let j = 0; j < this.allTags.length; j++) {
                        if (this.tags[i] === this.allTags[j].name) {
                            this.realTags.push({
                                id: this.allTags[j].id
                            })
                            break
                        }
                    }
                }
            },
            getRealLanguages() {
                this.realLanguages = []
                if (this.useC) {
                    this.realLanguages.push({
                        id: 1
                    })
                }
                if (this.useCpp) {
                    this.realLanguages.push({
                        id: 2
                    })
                }
                if (this.useJava) {
                    this.realLanguages.push({
                        id: 3
                    })
                }
                if (this.usePython) {
                    this.realLanguages.push({
                        id: 4
                    })
                }
            },
            async save() {
                if (!this.check()) {
                    return
                }
                this.getRealTags()
                this.getRealLanguages()
                let obj = {
                    title: this.title,
                    ref: this.ref,
                    description: this.description,
                    inputDescription: this.inputDescription,
                    outputDescription: this.outputDescription,
                    hint: this.hint,
                    source: this.source,
                    tag: this.realTags,
                    memoryLimit: this.memoryLimit,
                    realTimeLimit: this.realTimeLimit,
                    cpuTimeLimit: this.cpuTimeLimit,
                    visible: this.visible,
                    difficulty: this.difficulty,
                    language: this.realLanguages,
                    sample: this.sample,
                    problemCase: this.problemCase,
                    useSPJ: this.useSPJ,
                    template: [],
                    spj: []
                }
                if (!this.handleTmpl(obj)) {
                    return
                }
                this.handleSPJ(obj)
                console.log(obj)
                try {
                    const {
                        data: res
                    } = await this.$http.post('/admin/problem/addProblem', obj);
                    if (res.error) {
                        this.$message.error(res.error)
                        return
                    }
                    this.$message({
                        message: res.data,
                        type: 'success'
                    });
                    this.$router.push("/problem")
                } catch (err) {
                    console.log(err);
                    alert(err)
                }
            },
            handleTmpl(obj) {
                if (this.useCTmpl) {
                    let r = this.processTmpl(this.CTmpl, "C")
                    if (r === null) {
                        return false
                    }
                    obj.template.push(r)
                }
                if (this.useCppTmpl) {
                    let r = this.processTmpl(this.CppTmpl, "Cpp")
                    if (r === null) {
                        return false
                    }
                    obj.template.push(r)
                }
                if (this.useJavaTmpl) {
                    let r = this.processTmpl(this.JavaTmpl, "Java")
                    if (r === null) {
                        return false
                    }
                    obj.template.push(r)
                }
                if (this.usePythonTmpl) {
                    let r = this.processTmpl(this.PythonTmpl, "Python")
                    if (r === null) {
                        return false
                    }
                    obj.template.push(r)
                }
                return true
            },
            handleSPJ(obj) {
                if (this.useC) {
                    obj.spj.push({
                        lid: 1,
                        code: this.CSPJ,
                    })
                }
                if (this.useCpp) {
                    obj.spj.push({
                        lid: 2,
                        code: this.CppSPJ,
                    })
                }
                if (this.useJava) {
                    obj.spj.push({
                        lid: 3,
                        code: this.JavaSPJ,
                    })
                }
                if (this.usePython) {
                    obj.spj.push({
                        lid: 4,
                        code: this.PythonSPJ,
                    })
                }
            },
            addProblemCase() {
                this.problemCase.push({
                    input: '',
                    output: '',
                    score: 10
                })
            },
            addSample() {
                this.sample.push({
                    input: '',
                    output: ''
                })
            },
            deleteProblemCase(val) {
                if (this.problemCase.length === 1) {
                    this.$message.error("test case is required")
                    return
                }
                this.problemCase.splice(val, 1)
            },
            deleteSample(val) {
                if (this.sample.length === 1) {
                    this.$message.error("sample input and ouput is required")
                    return
                }
                this.sample.splice(val, 1)
            },
            async getAllTags() {
                try {
                    const {
                        data: res
                    } = await this.$http.get('/admin/tag/getAllShared');
                    if (res.error) {
                        this.$message.error(res.error)
                        return
                    }
                    this.allTags = res.data;
                } catch (err) {
                    console.log(err);
                    // alert(err)
                }
            },
        },
        components: {
            editor: Editor,
            codemirror,
        },
        watch: {
            useLang: {
                deep: true,
                handler(obj) {
                    for (const k in obj) if (obj.hasOwnProperty(k)) {
                        if (!obj[k]) {
                            this.useTmpl[k] = false
                            if (this.ltLang === k) {
                                this.ltLang = ""
                            }
                        }
                    }
                }
            },
            useSPJ() {
                if (this.SPJLang === "") {
                    this.SPJLang = "C"
                    this.SPJCode = this.SPJExample.C
                }
            },
            useTmpl: {
                deep: true,
                handler(obj) {
                    for (const k in obj) if (obj.hasOwnProperty(k)) {
                        if (obj[k]) {
                            if (this.tmpl[k] === '') {
                                this.tmpl[k] = this.tmplExample[k]
                            }
                        }
                    }
                }
            }
        }
    };
</script>

<style scoped>
    .center-box {
        width: 100%;
        background-color: #ffffff;
        border-radius: 10px;
        box-shadow: 0 2px 12px 0 rgba(0, 0, 0, 0.1);
    }

    .center-box >>> .w-e-text {
        overflow: visible !important
    }

    .center-box >>> .CodeMirror {
        height: auto;
    }

    .center-box >>> .CodeMirror-scroll {
        overflow: scroll !important;
        min-height: 200px;
        height: auto;
    }

    .template-checkbox {
        margin-top: 20px
    }

    .local-test-result {
        float: right;
        margin-right: 20px;
        margin-top: 10px;
    }

    .spj-checkbox {
        margin-top: 10px;
        border: 1px solid rgb(233, 233, 235);
    }

    .local-test-checkbox {
        margin-top: 10px;
        border: 1px solid rgb(233, 233, 235);
    }

    .page-container {
        padding: 8px 16px;
        background-color: #fff6f7;
        border-radius: 4px;
        border-left: 5px solid #fe6c6f;
        margin: 0 0 30px 0;
    }

    .dialog >>> .el-form-item__label {
        text-align: left;
        vertical-align: middle;
        float: left;
        font-size: 14px;
        min-width: 280px;
        color: #99a9bf;
        line-height: 40px;
        padding: 0 12px 0 0;
        box-sizing: border-box;
    }

    .add-button {
        border: 1px solid rgb(233, 233, 235);
        border-radius: 5px;
        margin-top: 30px;
        text-align: center;
        height: 40px;
        line-height: 40px;
    }

    .add-button:hover {
        background-color: #FAFDFF;
        cursor: pointer;
    }

    .small-element {
        margin-top: 20px;
        margin-left: 10px
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