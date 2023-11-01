<template>
  <div class="dashboard-container">
    <div class="dashboard-text">name: {{ name }}</div>
    <img src="http://img.suzone.net/biaoqing/201805/16977d32191cf4740604d37485803aba.gif"/>
  </div>
</template>

<script>
import { mapGetters } from 'vuex'

export default {
  name: 'Dashboard',
  computed: {
    ...mapGetters([
      'name'
    ])
  }
}
</script>

<style lang="scss" scoped>
.dashboard {
  &-container {
    margin: 30px;
  }
  &-text {
    font-size: 30px;
    line-height: 46px;
  }
}
</style>
