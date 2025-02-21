<template>
  <div class="q-pa-md flex column g-m">
    <div class="text-h6 text-primary text-center text-bold">
      1a - Лабораторная работа "Метод анализа иерархий"
    </div>
    <div
      class="page-title text-primary flex items-center"
      style="position: absolute"
    >
      <countdown ref="countdownRef" :time="900000 * 4" />
    </div>
    <div
      class="page-title text-primary flex items-center g-xl full-width justify-center text-center"
    >
      Получено баллов: {{ info?.percentage }} / 100
      <br />
      Текущий этап: {{ (info?.step ?? 1) - 1 }} / 6
    </div>
    <q-tabs v-model="tab" inline-label shrink stretch class="text-primary">
      <q-tab v-for="tab in tabs" :key="tab.name" v-bind="tab" />
    </q-tabs>
    <q-tab-panels
      v-if="variant"
      v-model="tab"
      animated
      class="shadow-2 rounded-borders"
    >
      <q-tab-panel name="task">
        <task-description :variant="variant" :variant-number="24" />
      </q-tab-panel>
      <q-tab-panel name="result-1">
        <result-universal
          :matrix="variant.matrix_1"
          :max-mark="maxMarks[0]"
          :step="1"
          matrix-title="Матрица важности по цели"
          :labels="criterias"
        />
      </q-tab-panel>
      <q-tab-panel name="result-2">
        <result-universal
          :matrix="variant.matrix_2"
          :max-mark="maxMarks[1]"
          :step="2"
          matrix-title="Матрица важности по критерию Доходность"
          :labels="alternatives"
        />
      </q-tab-panel>
      <q-tab-panel name="result-3">
        <result-universal
          :matrix="variant.matrix_3"
          :max-mark="maxMarks[2]"
          :step="3"
          matrix-title="Матрица важности по критерию Риск"
          :labels="alternatives"
        />
      </q-tab-panel>
      <q-tab-panel name="result-4">
        <result-universal
          :matrix="variant.matrix_4"
          :max-mark="maxMarks[3]"
          :step="4"
          matrix-title="Матрица важности по критерию Стоимость акций"
          :labels="alternatives"
        />
      </q-tab-panel>
      <q-tab-panel name="result-5">
        <result-universal
          :matrix="variant.matrix_5"
          :max-mark="maxMarks[4]"
          :step="5"
          matrix-title="Матрицы важности по критерию Ликвидность"
          :labels="alternatives"
        />
      </q-tab-panel>
      <q-tab-panel name="result-6">
        <result-last
          :variant="variant"
          :max-mark="maxMarks[4]"
          :step="6"
          matrix-title=""
          :row-labels="alternatives"
          :labels="criterias"
          :alternatives="alternatives"
          :criterias="criterias"
        />
      </q-tab-panel>
    </q-tab-panels>
  </div>
</template>

<script lang="ts" setup>
import { ref, onMounted, computed, watch } from 'vue';
import TaskDescription from './components/task-description.vue';
import resultUniversal from './components/result-universal.vue';
import Countdown from 'vue3-countdown';
import BannerComponent from 'src/components/BannerComponent.vue';
import { useLab1aStore } from 'src/stores/lab1a';
import { getVariant } from 'src/lab-services/lab1a/helpers/get-variant';
import { getMaxMarksForStep } from 'src/lab-services/lab1a/lab1a';
import { criterias, alternatives } from 'src/mock/lab1a';
import ResultLast from './components/result-last.vue';
import { Notify } from 'quasar';

const labStore = useLab1aStore();

localStorage.setItem(
  'lab-token',
  'sjdkfnskjdfnsjdfknsdjfnsdjkfsjkdfnsjdfnsjkdfjs'
);

const maxMarks = getMaxMarksForStep();

const variant = computed(() => labStore.variant);

const currentStep = computed(() => labStore.info?.step);

const info = computed(() => labStore.info);

const countdownRef = ref();

const pause = () => countdownRef.value.stop();

watch(currentStep, () => {
  tabs.value = fullSteps.filter(
    (tab, i) =>
      (currentStep.value !== undefined && i <= currentStep.value) ||
      tab.name === 'result'
  );
  if (currentStep.value === 7) {
    pause();
    Notify.create({
      type: 'positive',
      message: `Вы сделали лабораторную работу и получили ${info.value?.percentage} баллов из 100`,
    });
    labStore.sendResult();
  }
});

const fullSteps = [
  {
    name: 'task',
    icon: 'info',
    label: 'Задание',
  },
  {
    name: 'result-1',
    icon: 'done',
    label: 'Этап 1',
  },
  {
    name: 'result-2',
    icon: 'done',
    label: 'Этап 2',
  },
  {
    name: 'result-3',
    icon: 'done',
    label: 'Этап 3',
  },
  {
    name: 'result-4',
    icon: 'done',
    label: 'Этап 4',
  },
  {
    name: 'result-5',
    icon: 'done',
    label: 'Этап 5',
  },
  {
    name: 'result-6',
    icon: 'done',
    label: 'Этап 6',
  },
];

onMounted(async () => {
  await labStore.getInfo();
  const v = getVariant(8);
  await labStore.saveVariant({ number: v.variantNumber, data: v.variant });
});

const tabs = ref([
  {
    name: 'task',
    icon: 'info',
    label: 'Задание',
  },
  {
    name: 'result-1',
    icon: 'done',
    label: 'Этап 1',
  },
]);

const tab = ref('task');
</script>

<style lang="scss" scoped></style>
