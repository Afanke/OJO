<template>
    <div>
        <transition name="slide-fade">
            <div v-if="showList">
                <div class="center-box">
                    <el-row style="height:60px">
                        <span style="float:left;font-size:20px;margin-left:20px;margin-top:15px">Announcement</span>
                        <el-button style="float:right;margin-top:15px;margin-right:20px;" class="el-icon-refresh"
                                   type="primary"
                                   size="small" @click="reset">&nbsp;Reset
                        </el-button>
                    </el-row>
                    <el-row>
                        <el-table :show-header="false" :data="data" style="width:100%;border-radius:10px"
                                  v-loading="loading">
                            <el-table-column align="center" prop="" min-width="10" label="Score">
                            </el-table-column>
                            <el-table-column align="left" label="Id" min-width="180">
                                <template slot-scope="scope">
                                    <el-link style="font-size:18px" @click="getDetail(scope.row.id)">{{
                                            scope.row.title
                                        }}
                                    </el-link>
                                </template>
                            </el-table-column>
                            <el-table-column align="center" prop="lastUpdateTime" min-width="50" label="Score">
                            </el-table-column>
                            <el-table-column align="center" prop="creatorName" min-width="50" label="Language">
                            </el-table-column>
                        </el-table>
                    </el-row>
                </div>
                <el-row style="width:95%;margin-top:20px">
                    <el-pagination style="float:right;" background="" layout="prev, pager, next"
                                   :page-size="10" @current-change="handlePageChange" :current-page="page"
                                   :total="count">
                    </el-pagination>
                </el-row>
            </div>
            <div v-if="showDetail">
                <div class="center-box">
                    <el-row style="height:60px">
                        <span style="float:left;font-size:20px;margin-left:20px;margin-top:15px">{{ title }}</span>
                        <el-button style="float:right;margin-top:15px;margin-right:20px;" class="el-icon-back"
                                   type="primary"
                                   size="small" @click="goback">&nbsp;Back
                        </el-button>
                    </el-row>
                    <el-row style="width:90%;margin-left:5%" v-html="content">
                    </el-row>
                </div>
            </div>
        </transition>
    </div>
</template>
<script>
export default {
    created() {
        this.$bus.emit('changeHeader', '1');
        this.showList = false;
        this.showDetail = false
    },
    async mounted() {
        this.showList = true;
        this.loading = true
        await this.queryList()
    },
    data() {
        return {
            loading: false,
            count: 0,
            page: 1,
            showList: false,
            showDetail: false,
            title: "",
            content: "",
            data: []
        };
    },
    methods: {
        params_init() {
            if (this.$route.query.page) {
                this.page = Number(this.$route.query.page);
            } else {
                this.page = 1;
            }
        },
        params_query() {
            let obj = {};
            if (this.$route.query.page) {
                obj.page = Number(this.$route.query.page);
            }
            return obj;
        },
        fresh(obj) {
            this.$router.push({
                path: '/home',
                query: obj
            })
        },
        reset() {
            this.$router.push({
                path: '/home'
            });
        },
        goback() {
            this.showDetail = false
            this.showList = true
        },
        handlePageChange(val) {
            let obj = this.params_query();
            obj.page = Number(val);
            this.fresh(obj);
        },
        async queryList() {
            this.loading = true
            this.params_init();
            try {
                const {
                    data: res
                } = await this.$http.post('/announcement/getAll', {
                    page: this.page
                });
                if (res.error) {
                    this.$message.error(res.error);
                    return
                }
                this.data = res.data;
                this.loading = false
                const {
                    data: res1
                } = await this.$http.post('/announcement/getCount', {});
                if (res1.error) {
                    this.$message.error(res1.error);
                    return
                }
                this.count = res1.data;
            } catch (err) {
                console.log(err);
            }
        },
        async getDetail(id) {
            try {
                const {
                    data: res
                } = await this.$http.post('/announcement/getDetail', {
                    id: Number(id)
                });
                if (res.error) {
                    this.$message.error(res.error);
                    return
                }
                this.title = res.data.title
                this.content = res.data.content
                this.showList = false
                this.showDetail = true
            } catch (err) {
                console.log(err);
            }
        }
    },
    watch: {
        $route() {
            this.queryList();
        }
    },
    components: {},
};
</script>

<style scoped>
.center-box {
    min-width: 600px;
    margin: 20px auto 0;
    width: 90%;
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

.slide-fade-enter,
.slide-fade-leave-to

    /* .slide-fade-leave-active for below version 2.1.8 */
{
    transform: translateY(40px);
    opacity: 0;
}
</style>