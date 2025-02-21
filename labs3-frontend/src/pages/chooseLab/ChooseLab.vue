<template>
  <div class="flex column g-xl q-pa-lg">
    <q-btn
      v-for="lab in labs"
      :key="lab.path"
      @click="goToLab(lab.path)"
      :label="lab.name"
      color="primary"
      outline
    />
  </div>
</template>

<script lang="ts" setup>
import { useRouter } from 'vue-router';

const router = useRouter();

const labs = [
  {
    name: 'Лабораторная работа 1а',
    path: 'lab1-a',
  },
  {
    name: 'Лабораторная работа 2',
    path: 'lab2',
  },
  {
    name: 'Лабораторная работа 3а',
    path: 'lab3-a',
  },
  {
    name: 'Лабораторная работа 3b',
    path: 'lab3-b',
  },
  {
    name: 'Лабораторная работа 3c',
    path: 'lab3-c',
  },
];

const goToLab = (name: string) => {
  router.push({ name });
};

localStorage.removeItem('user-answers-1');
localStorage.removeItem('user-answers-2');
localStorage.removeItem('user-answers-3');
localStorage.removeItem('user-answers-4');
localStorage.removeItem('user-answers-5');
localStorage.removeItem('user-answers-6');
localStorage.removeItem('user-answers-7');
localStorage.removeItem('user-answers-8');
localStorage.removeItem('user-answers-9');
localStorage.removeItem('user-answers-1-1');
localStorage.removeItem('user-answers-1-2');
localStorage.removeItem('user-answers-1-3');
localStorage.removeItem('user-answers-1-4');
localStorage.removeItem('answer-1');
localStorage.removeItem('answer-2');
localStorage.removeItem('answer-3');
localStorage.removeItem('answer-4');
localStorage.removeItem('answer-5');
localStorage.removeItem('answer-6');
localStorage.removeItem('answer-7');
localStorage.removeItem('answer-8');
localStorage.removeItem('answer-9');
localStorage.removeItem('answer-10');
localStorage.removeItem('answer-2-1');
localStorage.removeItem('answer-2-2');
localStorage.removeItem('answer-2-3');
localStorage.removeItem('answer-2-4');
localStorage.removeItem('chosen-index-1');
localStorage.removeItem('chosen-index-2');
localStorage.removeItem('chosen-index-3');
localStorage.removeItem('chosen-index-4');
localStorage.removeItem('chosen-index-5');
localStorage.removeItem('chosen-index-6');
localStorage.removeItem('chosen-index-7');
localStorage.removeItem('chosen-index-8');
localStorage.removeItem('chosen-index-9');
localStorage.removeItem('chosen-index-10');
localStorage.removeItem('step-1');
localStorage.removeItem('step-3');
localStorage.removeItem('answer-3-0');
localStorage.removeItem('answer-3-1');
localStorage.removeItem('answer-3-2');
localStorage.removeItem('answer-3-3');
localStorage.removeItem('answer-3-4');
localStorage.removeItem('answer-3-5');
localStorage.removeItem('user-answers-3-1');
localStorage.removeItem('user-answers-3-2');
localStorage.removeItem('user-answers-3-3');
localStorage.removeItem('user-answers-3-4');
localStorage.removeItem('user-answers-3-5');
localStorage.removeItem('answer-4-0');
localStorage.removeItem('answer-4-1');
localStorage.removeItem('answer-4-2');
localStorage.removeItem('answer-4-3');
localStorage.removeItem('answer-4-4');
localStorage.removeItem('answer-4-5');
localStorage.removeItem('user-answers-4-1');
localStorage.removeItem('user-answers-4-2');
localStorage.removeItem('user-answers-4-3');
localStorage.removeItem('user-answers-4-4');
localStorage.removeItem('user-answers-4-5');
localStorage.removeItem('answer-5-0');
localStorage.removeItem('answer-5-1');
localStorage.removeItem('answer-5-2');
localStorage.removeItem('answer-5-3');
localStorage.removeItem('answer-5-4');
localStorage.removeItem('answer-5-5');
localStorage.removeItem('user-answers-5-1');
localStorage.removeItem('user-answers-5-2');
localStorage.removeItem('user-answers-5-3');
localStorage.removeItem('user-answers-5-4');
localStorage.removeItem('user-answers-5-5');
localStorage.removeItem('current-step-3');
localStorage.removeItem('current-step-4');
localStorage.removeItem('current-step-5');
localStorage.removeItem('answer-1-sum');
localStorage.removeItem('3c-normal-step');
</script>

<style lang="scss" scoped></style>
