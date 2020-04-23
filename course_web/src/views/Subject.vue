<template>
    <div class="home">
        <!--<img alt="Vue logo" src="../assets/logo.png">-->
        <CourseList :recordTime="recordTime" :subject="subject"/>
    </div>
</template>

<script>
    // @ is an alias to /src
    import CourseList from '@/components/CourseList.vue'

    export default {
        name: 'About',
        components: {
            CourseList
        },
        data(){
            return {
                recordTime: this.$route.query.recordTime,
                subject: this.$route.query.subject
        }
        }
    }
</script>
