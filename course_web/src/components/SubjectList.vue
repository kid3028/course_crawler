<template>

    <div class="deit">
        <div class="crumbs">
            <div class="cantainer">
                <div class="block">
                    <span class="demonstration">当前日期</span>
                    <el-date-picker
                            v-model="recordTime"
                            type="date"
                            @change="handleDateChange"
                            placeholder="选择日期">
                    </el-date-picker>
                </div>
                <!--<el-table style="width: 100%;"-->
                          <!--:data="statistic.slice((current-1)*size,current*size)">-->
                <el-table style="width: 100%;"
                          :data="statistic"
                        @row-click="handleRowClick"
                >
                    <el-table-column type="index" width="50">
                    </el-table-column>
                    <el-table-column label="日期" prop="recordTime" align="center">
                    </el-table-column>
                    <el-table-column label="科目" prop="subject"  align="center">
                    </el-table-column>
                    <el-table-column label="数量" prop="cnt"  align="center">
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
    </div>

</template>

<script>
    import {dayStatistic} from "../api/course";
    import {formatDate} from "../lib/date";
    import router from "../router";

    export default {
        created() {
            this.handleDayStatistic()
        },
        methods: {
            handleSizeChange: function (size) {
                this.size = size
                this.current = 1
                this.handleDayStatistic()
            },
            handleCurrentChange: function(currentPage){
                this.current = currentPage
                this.handleDayStatistic()
            },
            handleDateChange(event) {
                this.current = 1
                this.recordTime = event
                this.handleDayStatistic()
            },
            handleDayStatistic() {
                dayStatistic(formatDate(this.recordTime, "yyyy-MM-dd"), this.current, this.size).then(res => {
                    for(let data of res.data.record) {
                        data.recordTime = data.recordTime.substring(0,10)
                    }
                    this.statistic = res.data.record
                    this.total = res.data.total
                })
            },
            handleRowClick(val) {
                this.$router.push({
                    path: "/subject",
                    query:{
                        recordTime: val.recordTime,
                        subject: val.subject
                    }
                })
            }
        },
        data() {
            return {
                current: 1,
                size: 10,
                statistic: [],
                total: 0,
                recordTime: new Date()
            }
        }
    }
</script>