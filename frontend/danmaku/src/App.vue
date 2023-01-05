<template>
  <Danmaku class="main"/>
</template>

<script lang="ts">
import {defineComponent} from 'vue';
import Danmaku from './components/Danmaku.vue';

export default defineComponent({
  name: 'App',
  components: {
    Danmaku
  }
});
</script>

<style>
#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  color: #2c3e50;
}

#app >>> div >>>  .ivu-notice-notice {
  margin-bottom: 10px;
  padding: 16px;
  border-radius: 4px;
  box-shadow: 0 1px 6px rgb(0 0 0 / 20%);
  background: #ffffffd9!important;
  line-height: 1;
  position: relative;
  overflow: hidden;
}
</style>
