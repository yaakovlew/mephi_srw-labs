<template>
  <div>
    <BannerComponent>
      <q-table
        flat
        bordered
        title="Результаты"
        :rows="results"
        :columns="columns"
        row-key="name"
        binary-state-sort
        hide-bottom
        :separator="'cell'"
        :pagination="pagination"
      >
      </q-table>
    </BannerComponent>
  </div>
</template>

<script lang="ts" setup>
import BannerComponent from 'src/components/BannerComponent.vue';
import { Ref } from 'vue';
import { ref, onMounted } from 'vue';

const results: Ref<any[]> = ref([]);

const pagination = {
  rowsPerPage: 0,
};

const answersLocalStorageKeys = [
  'answer-1-sum',
  'answer-2',
  'answer-3-sum',
  'answer-4',
  'answer-7',
];

onMounted(() => {
  results.value = answersLocalStorageKeys
    .map((key, i) =>
      localStorage.getItem(key) !== null
        ? {
            ...JSON.parse(localStorage.getItem(key)),
            step: i + 1,
          }
        : null
    )
    .filter((item) => item);
});

const columns = [
  {
    name: 'step',
    label: 'Этап',
    field: 'step',
    style: 'width: 33%;',
  },
  {
    name: 'percentage',
    label: 'Баллы',
    field: 'percentage',
    style: 'width: 33%;',
  },
  {
    name: 'max_mark',
    label: 'Макс. балл',
    field: 'max_mark',
    style: 'width: 33%;',
  },
];
</script>

<style lang="scss" scoped></style>
