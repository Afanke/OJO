<template>
    <div>
        <transition name="slide-fade">
            <div class="center-box" v-if="show">
                <el-row style="height:60px;line-height:60px;">
                    <span style="font-size:20px;margin-left:20px">Judge Server</span>
                </el-row>
                <el-table
                        :data="judgeServer"
                        style="width: 100%">
                    <el-table-column
                            label="#"
                            type="index"
                            width="60"
                            align="center"
                    >
                    </el-table-column>
                    <el-table-column
                            prop="name"
                            label="Name"
                            align="center"
                            min-width="100">
                    </el-table-column>
                    <el-table-column
                            prop="address"
                            label="Address"
                            align="center"
                            min-width="60">
                    </el-table-column>
                    <el-table-column
                            prop="port"
                            label="Port"
                            align="center"
                            min-width="60">
                    </el-table-column>
                    <el-table-column
                            prop="weight"
                            label="Weight"
                            align="center"
                            min-width="60">
                    </el-table-column>
                    <el-table-column
                            label="State"
                            align="center"
                            min-width="100">
                        <template slot-scope="scope">
                            <el-button v-if="scope.row.connected&&scope.row.enabled" size="mini" type="success" plain>
                                Connected
                            </el-button>
                            <el-button v-if="!scope.row.enabled" size="mini" type="danger" plain>Disabled</el-button>
                            <el-button v-if="!scope.row.connected&&scope.row.enabled&&scope.row.message" size="mini"
                                       type="warning" plain>
                                {{scope.row.message}}
                            </el-button>
                            <el-button v-if="!scope.row.connected&&scope.row.enabled&&!scope.row.message" size="mini"
                                       type="warning" plain>
                                Connecting
                            </el-button>
                        </template>
                    </el-table-column>
                    <el-table-column
                            label="Enabled"
                            align="center"
                            min-width="60">
                        <template slot-scope="scope">
                            <el-switch @change="switchEnabled(scope.row)" v-model="scope.row.enabled"
                                       active-color="#409eff" inactive-color="#dcdfe6">
                            </el-switch>
                        </template>
                    </el-table-column>
                    <el-table-column label="Option" width="120" align="center">
                        <template slot-scope="scope">
                            <el-row>
                                <el-tooltip content="Edit" placement="top">
                                    <el-button size="mini" class="el-icon-edit-outline"
                                               @click="edit(shallowCopy(scope.row))"></el-button>
                                </el-tooltip>
                                <el-tooltip content="Delete" placement="top">
                                    <el-button size="mini" class="el-icon-delete" style="color:red"
                                               @click="del(scope.row.id)">
                                    </el-button>
                                </el-tooltip>
                            </el-row>
                        </template>
                    </el-table-column>
                </el-table>
                <el-row style="height:50px;margin-top:20px">
                    <el-button @click="createVisible=true" style="margin-left:30px;float:left" type="primary"
                               size="small" class="el-icon-plus">
                        Add
                    </el-button>
                </el-row>
                <el-dialog title="Add Judge Server" :visible.sync="createVisible" width="800px"
                           :before-close="handleCreateClose">
                    <el-row :gutter="20">
                        <el-col :span="3" style="line-height:40px">
                            <span style="color:red">*</span>
                            <span>&nbsp;Name</span>
                        </el-col>
                        <el-col :span="9">
                            <el-input v-model="addForm.name" placeholder="Name"></el-input>
                        </el-col>
                        <el-col :span="3" style="line-height:40px">
                            <span style="color:red">*</span>
                            <span>&nbsp;Password</span>
                        </el-col>
                        <el-col :span="9">
                            <el-input v-model="addForm.password" placeholder="Password"></el-input>
                        </el-col>
                    </el-row>
                    <el-row :gutter="20" style="margin-top: 20px;">
                        <el-col :span="3" style="line-height:40px">
                            <span style="color:red">*</span>
                            <span>&nbsp;Address</span>
                        </el-col>
                        <el-col :span="9">
                            <el-input v-model="addForm.address" placeholder="Address"></el-input>
                        </el-col>
                        <el-col :span="3" style="line-height:40px">
                            <span style="color:red">*</span>
                            <span>&nbsp;Port</span>
                        </el-col>
                        <el-col :span="9">
                            <el-input v-model="addForm.port" placeholder="Port"></el-input>
                        </el-col>
                    </el-row>
                    <el-row :gutter="20" style="margin-top: 20px;">
                        <el-col :span="3" style="line-height:40px">
                            <span style="color:red">*</span>
                            <span>&nbsp;Weight</span>
                        </el-col>
                        <el-col :span="9">
                            <el-input-number v-model="addForm.weight" controls-position="right"
                                             :min="1"
                                             :max="30000" :step="1">
                            </el-input-number>
                        </el-col>
                        <el-col :span="3" style="line-height:40px">
                            <span style="color:red">*</span>
                            <span>&nbsp;Enabled</span>
                        </el-col>
                        <el-col :span="9">
                            <el-switch style="margin-top:10px" v-model="addForm.enabled" active-color="#13ce66"
                                       inactive-color="#ff4949"></el-switch>
                        </el-col>
                    </el-row>
                    <el-row style="margin-top:30px">
                        <el-button style="float:right" type="primary" @click="add">Save</el-button>
                    </el-row>
                </el-dialog>
                <el-dialog title="Edit Judge Server" :visible.sync="editVisible" width="800px"
                           :before-close="handleEditClose">
                    <el-row :gutter="20">
                        <el-col :span="3" style="line-height:40px">
                            <span style="color:red">*</span>
                            <span>&nbsp;Name</span>
                        </el-col>
                        <el-col :span="9">
                            <el-input v-model="editForm.name" placeholder="Name"></el-input>
                        </el-col>
                        <el-col :span="3" style="line-height:40px">
                            <span style="color:red">*</span>
                            <span>&nbsp;Password</span>
                        </el-col>
                        <el-col :span="9">
                            <el-input v-model="editForm.password" placeholder="Password"></el-input>
                        </el-col>
                    </el-row>
                    <el-row :gutter="20" style="margin-top: 20px;">
                        <el-col :span="3" style="line-height:40px">
                            <span style="color:red">*</span>
                            <span>&nbsp;Address</span>
                        </el-col>
                        <el-col :span="9">
                            <el-input v-model="editForm.address" placeholder="Address"></el-input>
                        </el-col>
                        <el-col :span="3" style="line-height:40px">
                            <span style="color:red">*</span>
                            <span>&nbsp;Port</span>
                        </el-col>
                        <el-col :span="9">
                            <el-input v-model="editForm.port" placeholder="Port"></el-input>
                        </el-col>
                    </el-row>
                    <el-row :gutter="20" style="margin-top: 20px;">
                        <el-col :span="3" style="line-height:40px">
                            <span style="color:red">*</span>
                            <span>&nbsp;Weight</span>
                        </el-col>
                        <el-col :span="9">
                            <el-input-number v-model="editForm.weight" controls-position="right"
                                             :min="1"
                                             :max="30000" :step="1">
                            </el-input-number>
                        </el-col>
                        <el-col :span="3" style="line-height:40px">
                            <span style="color:red">*</span>
                            <span>&nbsp;Enabled</span>
                        </el-col>
                        <el-col :span="9">
                            <el-switch style="margin-top:10px" v-model="editForm.enabled" active-color="#13ce66"
                                       inactive-color="#ff4949"></el-switch>
                        </el-col>
                    </el-row>
                    <el-row style="margin-top:30px">
                        <el-button style="float:right" type="primary" @click="update">Save</el-button>
                    </el-row>
                </el-dialog>
            </div>
        </transition>
    </div>

</template>
<script>
    export default {
        data() {
            return {
                show: false,
                judgeServer: [],
                timeout: null,
                createVisible: false,
                editVisible: false,
                addForm: {
                    address: "",
                    connected: false,
                    enabled: false,
                    id: 0,
                    name: "",
                    password: "",
                    port: "",
                    weight: 1,
                },
                editForm: {
                    address: "",
                    connected: false,
                    enabled: false,
                    id: 0,
                    name: "",
                    password: "",
                    port: "",
                    weight: 1,
                },
            }
        },
        beforeDestroy() {
            clearTimeout(this.timeout)
        },
        created() {
            this.$bus.emit("changeHeader", "2-3")
            this.show = false
        },
        mounted() {
            this.show = true
            this.getAllInfo()
        },

        methods: {
            check(obj) {
                if (obj.name === "") {
                    this.$message.error("name can't be empty")
                    return false
                }
                if (obj.address === "") {
                    this.$message.error("address can't be empty")
                    return false
                }
                if (obj.port === "") {
                    this.$message.error("port can't be empty")
                    return false
                }
                return true
            },
            edit(obj) {
                this.editForm = obj
                this.editVisible = true
            },
            async getAllInfo() {
                try {
                    const {
                        data: res
                    } = await this.$http.get('/admin/jsp/getAllInfo');
                    if (res.error) {
                        this.$message.error(res.error)
                    } else {
                        this.judgeServer = res.data
                        this.timeout = setTimeout(
                            this.getAllInfo
                            , 5000);
                    }
                } catch (err) {
                    console.log(err);
                    alert(err)
                }
            },
            async add() {
                try {
                    if (!this.check(this.addForm)) {
                        return
                    }
                    this.addForm.port = Number(this.addForm.port)
                    const {
                        data: res
                    } = await this.$http.post('/admin/jsp/addJudgeServer',
                        this.addForm);
                    if (res.error) {
                        this.$message.error(res.error)
                    } else {
                        clearTimeout(this.timeout)
                        await this.getAllInfo()
                        this.createVisible = false
                        this.handleCreateClose()
                        this.$message.success("save successfully")
                    }
                } catch (err) {
                    console.log(err);
                    alert(err)
                }
            },
            async del(id) {
                try {
                    const {
                        data: res
                    } = await this.$http.post('/admin/jsp/deleteJudgeServer',
                        {
                            id: Number(id)
                        });
                    if (res.error) {
                        this.$message.error(res.error)
                    } else {
                        clearTimeout(this.timeout)
                        await this.getAllInfo()
                        this.$message.success("delete successfully")
                    }
                } catch (err) {
                    console.log(err);
                    alert(err)
                }
            },
            async update() {
                try {
                    if (!this.check(this.editForm)) {
                        return
                    }
                    this.editForm.port = Number(this.editForm.port)
                    const {
                        data: res
                    } = await this.$http.post('/admin/jsp/updateJudgeServer',
                        this.editForm);
                    if (res.error) {
                        this.$message.error(res.error)
                    } else {
                        clearTimeout(this.timeout)
                        await this.getAllInfo()
                        this.editVisible = false
                        this.handleEditClose()
                        this.$message.success("update successfully")
                    }
                } catch (err) {
                    console.log(err);
                    alert(err)
                }
            },
            async switchEnabled(obj) {
                console.log(obj)
                if (obj.enabled) {
                    try {
                        const {
                            data: res
                        } = await this.$http.post('/admin/jsp/setEnabledTrue', {
                            id: obj.id
                        });
                        if (res.error) {
                            this.$message.error(res.error)
                            obj.enabled = false
                            return
                        }
                        obj.enabled = true
                    } catch (err) {
                        console.log(err);
                        // alert(err)
                    }
                } else {
                    try {
                        const {
                            data: res
                        } = await this.$http.post('/admin/jsp/setEnabledFalse', {
                            id: obj.id
                        });
                        if (res.error) {
                            this.$message.error(res.error)
                            obj.enabled = true
                            return
                        }
                        obj.enabled = false
                    } catch (err) {
                        console.log(err);
                        // alert(err)
                    }
                }
            },
            handleCreateClose(done) {
                this.addForm = {
                    address: "",
                    enabled: false,
                    id: 0,
                    name: "",
                    password: "",
                    port: 2333,
                    weight: 1,
                }
                if (done) {
                    done()
                }
            },
            handleEditClose(done) {
                this.editForm = {
                    address: "",
                    enabled: false,
                    id: 0,
                    name: "",
                    password: "",
                    port: "",
                    weight: 1,
                }
                if (done) {
                    done()
                }
            },
            shallowCopy(src) {
                const dst = {};
                for (const prop in src) {
                    if (src.hasOwnProperty(prop)) {
                        dst[prop] = src[prop];
                    }
                }
                return dst;
            }
        },
        components: {}
    };
</script>

<style scoped>


    .center-box {
        background-color: #ffffff;
        border-radius: 10px;
        /* min-height: 600px; */
        box-shadow: 0 2px 12px 0 rgba(0, 0, 0, 0.1);
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