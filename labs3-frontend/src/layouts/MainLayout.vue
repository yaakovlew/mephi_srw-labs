<template>
  <q-layout view="lHh Lpr lFf" class="main-layout bg-grey-2">
    <q-header elevated class="main-layout__header">
      <q-toolbar>
        <q-btn
          flat
          dense
          round
          icon="menu"
          aria-label="Menu"
          @click="toggleLeftDrawer"
        />
        <div style="flex-grow: 100"></div>
        <q-btn to="/" @click="deleteToken" flat round dense icon="logout" />
      </q-toolbar>
    </q-header>
    <q-page-container>
      <router-view />
    </q-page-container>
  </q-layout>
</template>

<script setup lang="ts">
import { useAuthStore } from 'src/stores/auth';
import { ref } from 'vue';
import { useRouter } from 'vue-router';

const router = useRouter();

const leftDrawerOpen = ref(false);

const authStore = useAuthStore();

const toMain = () => {
  router.push('/');
};

const deleteToken = () => {
  authStore.logout();
  router.push('/');
};

function toggleLeftDrawer() {
  leftDrawerOpen.value = !leftDrawerOpen.value;
}
</script>

<style scoped lang="scss">
$max-width: 1440px;
.main-layout {
  max-width: $max-width;
  margin: 0 auto;

  &__header {
    max-width: $max-width;
    margin: 0 auto;
  }
}
</style>
