<template>
  <div>
    <transition name="slide-fade">
      <div class="center-box" v-if="show">
        <el-page-header style="height:60px;line-height:60px;margin-left:20px" title="Back" @back="goBack"
          content="Edit Problem">
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
            <el-col :span="8">
              <span style="color:red">*</span>
              <span>&nbsp;CPU Time Limit(ms)</span>
              <el-row class="small-element">
                <el-input-number v-model="cpuTimeLimit" controls-position="right" :min="100" :max="10000" :step="100">
                </el-input-number>
              </el-row>

            </el-col>
            <el-col :span="8">
              <span style="color:red">*</span>
              <span>&nbsp;Real Time Limit(ms)</span>
              <el-row class="small-element">
                <el-input-number v-model="realTimeLimit" controls-position="right" :min="100" :max="20000" :step="100">
                </el-input-number>
              </el-row>

            </el-col>
            <el-col :span="8">
              <span style="color:red">*</span>
              <span>&nbsp;Memory Limit(KB)</span>
              <el-row class="small-element">
                <el-input-number v-model="memoryLimit" controls-position="right" :min="2" :max="262144" :step="1024">
                </el-input-number>
              </el-row>

            </el-col>
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
                  <el-option v-for="(item,index1) in difficultyOptions" :key="index1+'index1'" :label="item.label"
                    :value="item.value">
                  </el-option>
                </el-select>
              </el-row>

            </el-col>
            <el-col :span="8" :offset="1">
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
                <el-checkbox-group v-model="languages">
                  <el-checkbox label="C" disabled></el-checkbox>
                  <el-checkbox label="C++" disabled></el-checkbox>
                  <el-checkbox label="Java" disabled></el-checkbox>
                  <el-checkbox label="Python3"></el-checkbox>
                </el-checkbox-group>
              </el-row>
            </el-col>
          </el-row>
          <el-row style="border:1px solid rgb(233, 233, 235);border-radius:5px;margin-top:30px"
            v-for="(item,index3) in sample" :key="index3+'index3'">
            <el-row style="height:50px;line-height:50px">
              <span style="margin-left:10px;font-weight:bold">Sample {{index3+1}}</span>
              <el-button type="warning" style="margin-right:40px;float:right;margin-top:9px" size="small"
                @click="deleteSample(index3)">Delete
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

          <el-row id="add-button">
            <div @click="addSample">
              <i class="el-icon-plus"></i>
              <span>&nbsp;Add Sample</span>
            </div>
          </el-row>
          <el-row style="border:1px solid rgb(233, 233, 235);border-radius:5px;margin-top:30px"
            v-for="(item,index4) in problemCase" :key="index4+'index4'">
            <el-row style="height:50px;line-height:50px">
              <span style="margin-left:10px;font-weight:bold">Test Case {{index4+1}}</span>
              <el-button type="warning" style="margin-right:40px;float:right;margin-top:9px" size="small"
                @click="deleteProblemCase(index4)">Delete
              </el-button>
              <el-input-number :step="10" size="small" :min="0" style="float:right;margin-right:40px;margin-top:9px"
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
                    <el-input resize="none" :autosize="{ minRows: 4}" type="textarea" :rows="2" placeholder="Input"
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
                    <el-input resize="none" :autosize="{ minRows: 4}" type="textarea" :rows="2" placeholder="Output"
                      v-model="item.output">
                    </el-input>
                  </el-row>
                </el-col>
              </el-row>
            </el-row>
          </el-row>
          <el-row id="add-button">
            <div @click="addProblemCase">
              <i class="el-icon-plus"></i>
              <span>&nbsp;Add Test Case</span>
            </div>
          </el-row>
          <el-row style="margin-top:30px">
            <span style="color:#fff">*</span>
            <span>&nbsp;Hint</span>
            <editor style="margin-top:20px" v-model="hint"></editor>
          </el-row>
          <el-row style="margin-top:30px">
            <span style="color:#fff">*</span>
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
  </div>
</template>
<script>
  import Editor from '@/components/Editor'
  export default {
    data() {
      return {
        id: 0,
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
        memoryLimit: 2048,
        realTimeLimit: 2000,
        cpuTimeLimit: 1000,
        visible: false,
        difficulty: 'Normal',
        languages: [],
        sample: [{
          input: '',
          output: '',
        }],
        problemCase: [{
          input: '',
          output: '',
          score: 10
        }],
        difficultyOptions: [{
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
      this.getDetail()
    },

    methods: {
      async getDetail() {
        try {
          const {
            data: res
          } = await this.$http.post('/admin/problem/getDetail', {
            id: Number(this.$route.query.id)
          });
          if (res.error) {
            this.$message.error(res.error)
            return
          }
          this.id = res.data.id
          this.title = res.data.title
          this.ref = res.data.ref
          this.description = res.data.description
          this.inputDescription = res.data.inputDescription
          this.outputDescription = res.data.outputDescription
          this.hint = res.data.hint
          this.source = res.data.source
          this.memoryLimit = res.data.memoryLimit
          this.realTimeLimit = res.data.realTimeLimit
          this.cpuTimeLimit = res.data.cpuTimeLimit
          this.visible = res.data.visible
          this.difficulty = res.data.difficulty
          this.problemCase = res.data.problemCase
          this.sample = res.data.sample
          for (let i = 0; i < res.data.language.length; i++) {
            this.languages.push("" + res.data.language[i].name)
          }
          if (res.data.tag) {
            for (let i = 0; i < res.data.tag.length; i++) {
              this.tags.push("" + res.data.tag[i].name)
            }
          }

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
        if (this.languages.length === 0) {
          this.$message.error("language is required")
          return false
        }
        for (let i = 0; i < this.sample.length; i++) {
          if (this.sample[i].output == "") {
            this.$message.error("sample " + (i + 1) + " output can't be empty")
            return false
          }
        }
        for (let i = 0; i < this.problemCase.length; i++) {
          if (this.problemCase[i].output == "") {
            this.$message.error("test case " + (i + 1) + " output can't be empty")
            return false
          }
          if (this.problemCase[i].score == 0) {
            this.$message.error("test case " + (i + 1) + " score can't be 0")
            return false
          }
        }
        return true
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
        for (let i = 0; i < this.languages.length; i++) {
          switch (this.languages[i]) {
            case "C":
              this.realLanguages.push({
                id: 1
              })
              break
            case "C++":
              this.realLanguages.push({
                id: 2
              })
              break
            case "Java":
              this.realLanguages.push({
                id: 3
              })
              break
            case "Python3":
              this.realLanguages.push({
                id: 4
              })
              break
          }
        }
      },
      async save() {
        if (!this.check()) {
          return
        }
        this.getRealTags()
        this.getRealLanguages()
        let obj = {
          id: this.id,
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
          problemCase: this.problemCase
        }
        console.log(obj)
        try {
          const {
            data: res
          } = await this.$http.post('/admin/problem/updateProblem', obj);
          if (res.error) {
            this.$message.error(res.error)
            return
          }
          this.$message({
            message: res.data,
            type: 'success'
          });
        } catch (err) {
          console.log(err);
          alert(err)
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
        console.log(this.sample)
        this.sample.push({
          input: '',
          output: ''
        })
        console.log(this.sample)
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
          // console.log(res);
          if (res.error) {
            this.$message.error(res.error)
            return
          }
          this.allTags = res.data;
          console.log(this.allTags)
        } catch (err) {
          console.log(err);
          // alert(err)
        }
      },
      goBack() {
        this.$router.go(-1)
      }
    },
    components: {
      editor: Editor
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


  .center-box>>>.w-e-text {
    overflow: visible !important
  }

  #add-button {
    border: 1px solid rgb(233, 233, 235);
    border-radius: 5px;
    margin-top: 30px;
    text-align: center;
    height: 40px;
    line-height: 40px;
  }

  #add-button:hover {
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