<template>
    <div class="deit">
        <div class="crumbs">
            <div class="cantainer">
                <el-table style="width: 100%;"
                          :data="courseList"
                          @row-click="handleRowClick">
                    <el-table-column type="index" width="50">
                    </el-table-column>
                    <el-table-column label="日期" prop="recordTime" align="center">
                    </el-table-column>
                    <el-table-column label="科目" prop="subject"  align="center">
                    </el-table-column>

                    <el-table-column label="课程id" prop="cid"  align="center">
                    </el-table-column>
                    <el-table-column label="名称" prop="name"  align="center">
                    </el-table-column>
                    <el-table-column label="封面" align="center">
                        <template slot-scope="scope">
                            <img :src="scope.row.coverUrl" alt="" width="80px" height="80px">
                        </template>
                    </el-table-column>
                    <el-table-column label="原价" prop="preAmount"  align="center">
                    </el-table-column>
                    <el-table-column label="折后价" prop="afAmount"  align="center">
                    </el-table-column>

                </el-table>
                <el-pagination
                        @size-change="handleSizeChange"
                        @current-change="handleCurrentChange"
                        :current-page="current"
                        :page-sizes="[1, 5, 10, 20, 40]"
                        :page-size="size"
                        layout="total, sizes, prev, pager, next, jumper"
                        :total="total">
                </el-pagination>
            </div>
        </div>

        <div>
            <el-dialog v-if="visible" :visible.sync="visible" width="600px" :modal="true" @close="handleCloseDialog">
                <h2 slot="title">详情</h2>
                <CourseDetail ref="courseDetail" :courseDetail="courseDetail"></CourseDetail>
                <!--<div slot="footer" class="detail-wrap-bottom">-->
                    <!--<el-button type="primary">确认</el-button>-->
                    <!--<el-button type="default">退回</el-button>-->
                <!--</div>-->
            </el-dialog>
        </div>
    </div>


</template>

<script>
    import {courseList, courseDetail} from "../api/course";

    import CourseDetail from '@/components/CourseDetail.vue'

    export default {
        created() {
            this.handleCourseList()
        },
        methods: {
            handleSizeChange: function (size) {
                this.size = size
                this.current = 1
                this.handleCourseList()
            },
            handleCurrentChange: function(currentPage){
                this.current = currentPage
                this.handleCourseList()
            },
            handleCourseList() {
                courseList(this.recordTime, this.subject, this.current, this.size).then(res => {
                    for(let data of res.data.record) {
                        data.recordTime = data.recordTime.substring(0,10)
                        data.preAmount = data.preAmount / 100
                        data.afAmount = data.afAmount / 100
                    }
                    this.courseList = res.data.record
                    this.total = res.data.total
                })
            },
            handleRowClick(val) {
                courseDetail(val.cid).then(res => {
                    this.courseDetail = res.data
                    this.visible = !this.visible
                })
            },
            handleCloseDialog() {
                this.visible = false
                this.courseDetail = {}
            }
        },
        data() {
            return {
                current: 1,
                size: 10,
                courseList: [],
                total: 0,
                activeCid: null,
                visible: false,
                courseDetail: {}
            }
        },
        props:{
            recordTime: String,
            subject: String
        },
        components:{
            CourseDetail: CourseDetail
        }
    }
</script>