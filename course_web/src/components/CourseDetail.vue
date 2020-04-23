<template xmlns:text-align="http://www.w3.org/1999/xhtml" xmlns:list-style-position="http://www.w3.org/1999/xhtml">
    <div  :data="courseDetail">
        <h2>课程信息</h2>

        <ul>
            <li><label>课程id：</label><span>{{courseDetail.course.cid}}</span></li> <br>
            <li><label>课程名称：</label><span>{{courseDetail.course.name}}</span></li>
            <li><label>所属科目：</label><span>{{courseDetail.course.subject}}</span></li>
            <li><label>原价：</label><span>{{courseDetail.course.preAmount / 100}}</span></li>
            <li><label>折后价：</label><span>{{courseDetail.course.afAmount / 100}}</span></li>
            <li><label>封面：</label><span><img :src="courseDetail.course.coverUrl" width="50px" height="50px"/></span></li>
        </ul>

        <h2>教师信息</h2>
        <ul v-for=" item in courseDetail.teacherList">
            <li><label>名称:</label><span>{{item.name}}</span></li>
            <li><label>头像:</label><span><img :src="item.coverUrl" width="50px" height="50px"/></span></li>
            <li><label>介绍:</label><span>{{item.introduce}}</span></li>
            <li><label>职位:</label><span>{{item.tutor == 0 ? "教师" : "辅导"}}</span></li>
        </ul>

    </div>
</template>
<script>
    export default {
        props: {
            courseDetail: {
                type: Object,
                default: null
            }
        },
        data() {
            return {
                loading: false
            }
        }
    }
</script>
<style>

</style>