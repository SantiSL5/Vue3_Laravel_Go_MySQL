<template>
  <div class="container-fluid padding-fix">
    <Header />
    <router-view></router-view>
  </div>
</template>

<script>
import "bootstrap/dist/css/bootstrap.min.css";
import "bootstrap/dist/js/bootstrap.min.js";
import Header from './components/Header';

export default {
  components: { Header }
}
</script>

<style>
.padding-fix {
  padding: 0 !important;
}
</style>