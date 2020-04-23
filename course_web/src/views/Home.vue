<template>
  <div class="home">
    <SubjectList/>
  </div>
</template>

<script>
// @ is an alias to /src
import SubjectList from '@/components/SubjectList.vue'


export default {
  name: 'Home',
  components: {
      SubjectList
  }
}
</script>
