<template>
  <banner-component class="q-pa-md q-mt-xl container" style="margin-top: 100px">
    <auth-sign-in />
  </banner-component>
</template>

<script setup lang="ts">
import { ref } from 'vue';
import AuthSignIn from './components/AuthSignIn.vue';

import BannerComponent from 'src/components/BannerComponent.vue';

const tab = ref('sign-in');
</script>

<style lang="scss" scoped>
@import '../../css/quasar.variables.scss';
.container {
  max-width: 600px;
  margin: 0 auto;
}
.backgrond-primary {
  background-color: $primary;
}
</style>
