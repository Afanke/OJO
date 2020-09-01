<template>

    <div>
        <transition name="slide-fade">

            <div class="center-box" v-if="show">
                <el-row style="height:60px;line-height:60px">
                    <span style="font-size:20px;margin-left:20px">Tag List</span>
                    <el-button style="float:right;margin-top:15px;margin-right:20px;" class="el-icon-refresh"
                               type="primary"
                               size="small" @click="reset">&nbsp;Reset
                    </el-button>
                    <el-input style="float:right;width:200px;margin-top:15px;margin-right:20px" placeholder="keywords"
                              v-model="keywords" size="small">
                        <el-button slot="append" icon="el-icon-search"
                                   style="color:#ffffff;background-color:#409EFF;border-top-left-radius:0;border-bottom-left-radius:0;margin-right:-21px;margin-top:-7px"
                                   size="small" @click="handleKeywordsChange"></el-button>
                    </el-input>
                    <div style="float:right;margin-top:16px;margin-right:80px">
                        <my-switch v-model="mine" @toggle="toggle"></my-switch>
                    </div>
                </el-row>
                <el-row style="height:1px;float:top;border-top:1px solid rgb(233, 233, 235);"></el-row>
                <el-table :data="tableData" style="width: 100%" v-loading="loading" size="small">
                    <el-table-column label="ID" prop="id" align="center" min-width="30">
                    </el-table-column>
                    <el-table-column label="Name" prop="name" min-width="60">
                    </el-table-column>
                    <el-table-column label="Creator" prop="creatorName" align="center" min-width="60">
                    </el-table-column>
                    <el-table-column label="Create Time" prop="createTime" align="center" min-width="60">
                    </el-table-column>
                    <el-table-column label="Last Update Time" prop="lastUpdateTime" align="center" min-width="60">
                    </el-table-column>
                    <el-table-column label="Visible" width="90" align="center">
                        <template slot-scope="scope">
                            <div @click="switchVisible(scope.row)">
                                <el-switch v-model="scope.row.visible" active-color="#409eff" inactive-color="#dcdfe6">
                                </el-switch>
                            </div>
                        </template>
                    </el-table-column>
                    <el-table-column label="Shared" width="90" align="center">
                        <template slot-scope="scope">
                            <div @click="switchShared(scope.row)">
                                <el-switch v-model="scope.row.shared" active-color="#409eff" inactive-color="#dcdfe6">
                                </el-switch>
                            </div>
                        </template>
                    </el-table-column>
                    <el-table-column label="Option" width="120" align="center">
                        <template slot-scope="scope">
                            <el-row>
                                <el-tooltip content="Edit" placement="top">
                                    <el-button size="mini" class="el-icon-edit-outline"
                                               @click="goEditTag(scope.row)"></el-button>
                                </el-tooltip>
                                <el-tooltip content="Delete" placement="top">
                                    <el-button size="mini" class="el-icon-delete" style="color:red"
                                               @click="deleteTag(scope.row.id)">
                                    </el-button>
                                </el-tooltip>
                            </el-row>
                        </template>
                    </el-table-column>
                </el-table>
                <el-row style="margin-top:20px;">
                    <el-button type="primary" @click="goCreate" size="small" class="el-icon-plus"
                               style="margin-left:30px;float:left"> Create
                    </el-button>
                    <el-pagination @current-change="handlePageChange" :page-size="10"
                                   style="float:right;margin-right:30px"
                                   layout="prev, pager, next" :total="count" :current-page.sync="page">
                    </el-pagination>
                </el-row>
                <el-row style="margin-top:20px;">
                </el-row>

            </div>
        </transition>
        <el-dialog title="Create Tag" :visible.sync="showCreateTag" width="500px" :before-close="handleCreateClose">
            <el-row :gutter="20">
                <el-col :span="4" style="line-height:40px">
                    <span style="color:red">*</span>
                    <span>&nbsp;Name</span>
                </el-col>
                <el-col :span="18">
                    <el-input v-model="tagName" placeholder="Name"></el-input>
                </el-col>
            </el-row>
            <el-row :gutter="20" style="margin-top:40px">
                <el-col :span="5" style="margin-top:2px">
                    <span style="color:#fff">*</span>
                    <span>&nbsp;Visible</span>
                </el-col>
                <el-col :span="6">
                    <el-switch v-model="visible" active-color="#13ce66" inactive-color="#ff4949">
                    </el-switch>
                </el-col>
                <el-col :span="5" style="margin-top:2px">
                    <span style="color:#fff">*</span>
                    <span>&nbsp;Shared</span>
                </el-col>
                <el-col :span="6">
                    <el-switch v-model="shared" active-color="#13ce66" inactive-color="#ff4949">
                    </el-switch>
                </el-col>
            </el-row>
            <span slot="footer" class="dialog-footer">
        <el-button @click="handleCreateClose">Cancel</el-button>
        <el-button type="primary" @click="createTag">Save</el-button>
      </span>
        </el-dialog>
        <el-dialog title="Edit Tag" :visible.sync="showEditTag" width="500px" :before-close="handleEditClose">
            <el-row :gutter="20">
                <el-col :span="4" style="line-height:40px">
                    <span style="color:red">*</span>
                    <span>&nbsp;Name</span>
                </el-col>
                <el-col :span="18">
                    <el-input v-model="tagName" placeholder="Name"></el-input>
                </el-col>
            </el-row>
            <span slot="footer" class="dialog-footer">
        <el-button @click="handleEditClose">Cancel</el-button>
        <el-button type="primary" @click="editTag">Save</el-button>
      </span>
        </el-dialog>
    </div>

</template>
<script>
import SwitchButton from "@/components/Switch.vue"

export default {
    data() {
        return {
            show: false,
            id: 0,
            showEditTag: false,
            tableData: [],
            mine: false,
            count: 0,
            loading: false,
            keywords: '',
            showCreateTag: false,
            tagName: '',
            visible: false,
            shared: false
        }
    },
    created() {
        this.$bus.emit("changeHeader", "3-2")
        this.show = false
    },
    async mounted() {
        this.show = true
        this.queryList()
    },
    methods: {
        async createTag() {
            if (this.tagName === '') {
                this.$message.error("Name can't be empty")
                return
            }
            try {
                const {
                    data: res
                } = await this.$http.post('/admin/tag/addTag', {
                    name: this.tagName,
                    shared: this.shared,
                    visible: this.visible
                });
                if (res.error) {
                    this.$message.error(res.error)
                    return
                }
                this.$message({
                    message: res.data,
                    type: 'success'
                });
                await this.queryList()
                this.handleCreateClose()
            } catch (err) {
                console.log(err);
            }
        },
        handleCreateClose(done) {
            this.tagName = '',
                this.visible = false,
                this.shared = false
            this.showCreateTag = false

            if (done) {
                done()
            }
        },
        handleEditClose(done) {
            this.tagName = ''
            this.id = 0
            this.showEditTag = false
            if (done) {
                done()
            }
        },
        goEditTag(obj) {
            this.tagName = obj.name
            this.id = obj.id
            this.showEditTag = true
        },
        toggle(checked) {
            this.mine = checked
            let obj = this.paramsQuery();
            obj.mine = this.mine;
            obj.page = 1;
            this.fresh(obj);
        },
        async editTag() {
            if (this.tagName === '') {
                this.$message.error("Name can't be empty")
                return
            }
            try {
                const {
                    data: res
                } = await this.$http.post('/admin/tag/updateTag', {
                    id: this.id,
                    name: this.tagName
                });
                if (res.error) {
                    this.$message.error(res.error)
                    return
                }
                this.$message({
                    message: res.data,
                    type: 'success'
                });
                await this.queryList()
                this.handleEditClose()
            } catch (err) {
                console.log(err);
            }
        },
        async deleteTag(id) {
            try {
                const {
                    data: res
                } = await this.$http.post('/admin/tag/deleteTag', {
                    id: id,
                });
                if (res.error) {
                    this.$message.error(res.error)
                    return
                }
                this.$message({
                    message: res.data,
                    type: 'success'
                });
                await this.queryList()
            } catch (err) {
                console.log(err);
            }
        },
        paramsInit() {
            if (this.$route.query.page) {
                this.page = Number(this.$route.query.page);
            } else {
                this.page = 1;
            }
            if (this.$route.query.keywords) {
                this.keywords = this.$route.query.keywords;
            } else {
                this.keywords = '';
            }
            if (this.$route.query.mine) {
                this.mine = true
            } else {
                this.mine = false
            }
        },
        paramsQuery() {
            let obj = {};
            if (this.$route.query.page) {
                obj.page = Number(this.$route.query.page);
            }
            if (this.$route.query.keywords) {
                obj.keywords = this.$route.query.keywords;
            }
            if (this.$route.query.mine) {
                obj.mine = true
            } else {
                obj.mine = false
            }
            return obj;
        },
        async queryList() {
            this.loading = true
            this.paramsInit()
            try {
                const {
                    data: res
                } = await this.$http.post('/admin/tag/getAll', {
                    page: this.page,
                    mine: this.mine,
                    keywords: this.keywords
                });
                if (res.error) {
                    this.$message.error(res.error)
                    return
                }
                this.tableData = res.data
                const {
                    data: res1
                } = await this.$http.post('/admin/tag/getCount', {
                    page: this.page,
                    mine: this.mine,
                    keywords: this.keywords
                });
                if (res1.error) {
                    this.$message.error(res1.error)
                    return
                }
                this.count = res1.data
                this.loading = false
            } catch (err) {
                console.log(err);
            }
        },
        async switchVisible(obj) {
            if (obj.visible) {
                try {
                    const {
                        data: res
                    } = await this.$http.post('/admin/tag/setVisibleTrue', {
                        id: obj.id
                    });
                    if (res.error) {
                        this.$message.error(res.error)
                        obj.visible = false
                        return
                    }
                    obj.visible = true
                    obj.lastUpdateTime = res.data
                } catch (err) {
                    console.log(err);
                }
            } else {
                try {
                    const {
                        data: res
                    } = await this.$http.post('/admin/tag/setVisibleFalse', {
                        id: obj.id
                    });
                    if (res.error) {
                        this.$message.error(res.error)
                        obj.visible = true
                        return
                    }
                    obj.visible = false
                    obj.lastUpdateTime = res.data
                } catch (err) {
                    console.log(err);
                }
            }
        },
        async switchShared(obj) {
            if (obj.shared) {
                try {
                    const {
                        data: res
                    } = await this.$http.post('/admin/tag/setSharedTrue', {
                        id: obj.id
                    });
                    if (res.error) {
                        this.$message.error(res.error)
                        obj.shared = false
                        return
                    }
                    obj.shared = true
                    obj.lastUpdateTime = res.data
                } catch (err) {
                    console.log(err);
                }
            } else {
                try {
                    const {
                        data: res
                    } = await this.$http.post('/admin/tag/setSharedFalse', {
                        id: obj.id
                    });
                    if (res.error) {
                        this.$message.error(res.error)
                        obj.shared = true
                        return
                    }
                    obj.shared = false
                    obj.lastUpdateTime = res.data
                } catch (err) {
                    console.log(err);
                }
            }
        },
        handlePageChange(val) {
            let obj = this.paramsQuery();
            obj.page = Number(val);
            this.fresh(obj);
        },
        handleKeywordsChange() {
            let obj = this.paramsQuery();
            obj.keywords = this.keywords;
            obj.page = 1;
            this.fresh(obj);
        },
        fresh(obj) {
            this.$router.push({
                path: '/problem/tag',
                query: obj
            });
        },
        reset() {
            this.$router.push({
                path: '/problem/tag'
            });
        },
        goCreate() {
            this.showCreateTag = true
        },
    },
    watch: {
        $route() {
            this.queryList();
        }
    },
    components: {
        mySwitch: SwitchButton
    }
};
</script>

<style scoped>
.center-box {
    background-color: #ffffff;
    border-radius: 10px;
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