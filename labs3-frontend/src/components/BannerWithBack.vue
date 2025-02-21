<template>
  <div class="flex text-primary page-title full-width g-m">
    <banner-component @click="router.back()">
      <slot name="icon">
        <q-icon
          name="arrow_back"
          color="primary"
          size="40px"
          class="cursor-pointer"
        />
      </slot>
    </banner-component>
    <banner-component style="flex-grow: 100">
      <slot />
    </banner-component>
  </div>
</template>

<script lang="ts" setup>
import { useRouter } from 'vue-router';
import BannerComponent from './BannerComponent.vue';

const router = useRouter();
</script>

<style lang="scss" scoped></style>
