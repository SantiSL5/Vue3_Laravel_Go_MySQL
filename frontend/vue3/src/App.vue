<template>
  <div class="container-fluid padding-fix">
    <Header />
    <router-view></router-view>
  </div>
</template>

<script>
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