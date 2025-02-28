<template>
  <div id="app">
    <VideoList />
  </div>
</template>

<script>
import VideoList from './components/VideoList.vue';

export default {
  name: 'App',
  components: {
    VideoList,
  },
};
</script>

<style>
@import './assets/base.css';
/* #app{
  background: black;
} */
</style>
